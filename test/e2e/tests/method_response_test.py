# Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License"). You may
# not use this file except in compliance with the License. A copy of the
# License is located at
#
#     http://aws.amazon.com/apache2.0/
#
# or in the "license" file accompanying this file. This file is distributed
# on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
# express or implied. See the License for the specific language governing
# permissions and limitations under the License.

"""Integration tests for the Method Response resource with a focus on patching operations
"""

import logging
import time
from typing import Dict, Tuple
from functools import partial

import boto3
import pytest
from acktest.k8s import resource as k8s
from acktest.resources import random_suffix_name
from e2e import service_marker, CRD_GROUP, CRD_VERSION, SERVICE_NAME, load_apigateway_resource
from e2e.common.waiter import wait_until_deleted, safe_get
from e2e.replacement_values import REPLACEMENT_VALUES
from .resource_test import simple_resource

pytest_plugins = ["e2e.tests.rest_api_test"]

METHOD_RESPONSE_RESOURCE_PLURAL = "apimethodresponses"
MODIFY_WAIT_AFTER_SECONDS = 60


@pytest.fixture(scope='module')
def apigateway_client():
    return boto3.client(SERVICE_NAME)


@pytest.fixture
def method_with_response(simple_resource, apigateway_client):
    resource_ref, resource_cr, rest_api_cr, resource_query = simple_resource
    method_response_name = random_suffix_name('methodresponse', 32)

    method_query = {
        'restApiId': resource_query['restApiId'],
        'resourceId': resource_query['resourceId'],
        'httpMethod': 'GET'
    }

    apigateway_client.put_method(**method_query, authorizationType='NONE')

    replacements = REPLACEMENT_VALUES.copy()
    replacements["METHOD_RESPONSE_NAME"] = method_response_name
    replacements["RESOURCE_REF_NAME"] = resource_cr['metadata']['name']
    replacements["REST_API_REF_NAME"] = rest_api_cr['metadata']['name']

    method_response_data = load_apigateway_resource(
        "method_response_simple",
        additional_replacements=replacements
    )

    ref = k8s.CustomResourceReference(
        CRD_GROUP, CRD_VERSION, METHOD_RESPONSE_RESOURCE_PLURAL,
        method_response_name, namespace='default'
    )

    k8s.create_custom_resource(ref, method_response_data)
    cr = k8s.wait_resource_consumed_by_controller(ref, wait_periods=15)

    assert cr is not None
    assert k8s.get_resource_exists(ref)

    method_response_query = {
        **method_query,
        'statusCode': '200'
    }

    yield ref, cr, method_response_query

    _, deleted = k8s.delete_custom_resource(ref, 3, 10)
    assert deleted
    wait_until_deleted(
        partial(apigateway_client.get_method_response, **method_response_query))

    try:
        apigateway_client.delete_method(**method_query)
    except Exception as e:
        logging.error(f"Failed to delete method: {str(e)}")


@service_marker
class TestMethodResponsePatch:
    def test_method_response_with_patch(self, method_with_response, apigateway_client):
        ref, cr, method_response_query = method_with_response
        get_method_response = partial(
            apigateway_client.get_method_response, **method_response_query)

        method_response = safe_get(get_method_response)
        assert method_response is not None

        patch_data = {
            "spec": {
                "responseParameters": {
                    "method.response.header.Content-Type": True,
                    "method.response.header.X-Powered-By": False,
                    "method.response.header.X-Custom-Header": True
                },
                "responseModels": {
                    "application/json": "Empty"
                }
            }
        }

        k8s.patch_custom_resource(ref, patch_data)
        time.sleep(MODIFY_WAIT_AFTER_SECONDS)

        method_response = safe_get(get_method_response)

        assert 'responseParameters' in method_response, "responseParameters field is missing in the method response"
        assert method_response['responseParameters']['method.response.header.Content-Type'] is True, "Content-Type parameter should be True"
        assert method_response['responseParameters']['method.response.header.X-Powered-By'] is False, "X-Powered-By parameter should be False"

        assert 'responseModels' in method_response, "responseModels field is missing in the method response"
        assert method_response['responseModels']['application/json'] == 'Empty', "application/json model should be Empty"

        # Verify using boto3 API directly
        aws_method_response = apigateway_client.get_method_response(
            **method_response_query)

        expected_response_params = {
            'method.response.header.Content-Type': True,
            'method.response.header.X-Powered-By': False,
            'method.response.header.X-Custom-Header': True
        }

        expected_response_models = {
            'application/json': 'Empty'
        }

        assert aws_method_response['responseParameters'] == expected_response_params, "AWS API responseParameters don't match expected values"
        assert aws_method_response['responseModels'] == expected_response_models, "AWS API responseModels don't match expected values"
