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

"""Integration tests for the Method resource with a focus on patching operations
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

# Import the rest_api_test module to make its fixtures available
pytest_plugins = ["e2e.tests.rest_api_test"]

METHOD_RESOURCE_PLURAL = "methods"
MODIFY_WAIT_AFTER_SECONDS = 60


@pytest.fixture(scope='module')
def apigateway_client():
    return boto3.client(SERVICE_NAME)


@pytest.fixture
def custom_method_with_patch(simple_resource, apigateway_client) -> Tuple[k8s.CustomResourceReference, Dict, Dict]:
    method_name = random_suffix_name('patch-method', 32)

    resource_ref, resource_cr, rest_api_cr, resource_query = simple_resource

    resource_id = resource_query['resourceId']
    rest_api_id = resource_query['restApiId']

    replacements = REPLACEMENT_VALUES.copy()
    replacements["METHOD_NAME"] = method_name
    replacements["RESOURCE_REF_NAME"] = resource_cr['metadata']['name']
    replacements["REST_API_REF_NAME"] = rest_api_cr['metadata']['name']

    method_data = load_apigateway_resource(
        "method_simple",
        additional_replacements=replacements
    )

    method_query = {
        'restApiId': rest_api_id,
        'resourceId': resource_id,
        'httpMethod': method_data['spec']['httpMethod']
    }

    ref = k8s.CustomResourceReference(
        CRD_GROUP, CRD_VERSION, METHOD_RESOURCE_PLURAL,
        method_name, namespace='default',
    )
    k8s.create_custom_resource(ref, method_data)
    cr = k8s.wait_resource_consumed_by_controller(ref, wait_periods=15)

    assert cr is not None
    assert k8s.get_resource_exists(ref)

    yield ref, cr, method_query

    _, deleted = k8s.delete_custom_resource(ref, 3, 10)
    assert deleted
    wait_until_deleted(partial(apigateway_client.get_method, **method_query))


@service_marker
class TestMethodPatch:
    def test_method_with_patch(self, custom_method_with_patch, apigateway_client):
        (ref, cr, method_query) = custom_method_with_patch
        get_method = partial(apigateway_client.get_method, **method_query)

        # Verify the initial method configuration is applied
        method = safe_get(get_method)
        assert method is not None
        assert method['operationName'] == 'NewTestOperation'

        # Apply patch to add request parameters
        patch_data = {
            'spec': {
                'requestParameters': {
                    'method.request.querystring.version': False,
                    'method.request.querystring.filter': False,
                    'method.request.header.X-Api-Key': True,
                    'method.request.header.Accept': False
                }
            }
        }

        k8s.patch_custom_resource(ref, patch_data)
        time.sleep(MODIFY_WAIT_AFTER_SECONDS)

        method = safe_get(get_method)

        assert 'requestParameters' in method, "requestParameters field is missing in the method"
        assert method['requestParameters']['method.request.querystring.version'] is False, "querystring.version parameter should be False"
        assert method['requestParameters']['method.request.querystring.filter'] is False, "querystring.filter parameter should be False"
        assert method['requestParameters']['method.request.header.X-Api-Key'] is True, "header.X-Api-Key parameter should be True"
        assert method['requestParameters']['method.request.header.Accept'] is False, "header.Accept parameter should be False"

        # Verify using boto3 API directly
        aws_method = apigateway_client.get_method(
            restApiId=method_query['restApiId'],
            resourceId=method_query['resourceId'],
            httpMethod=method_query['httpMethod']
        )

        expected_request_params = {
            'method.request.querystring.version': False,
            'method.request.querystring.filter': False,
            'method.request.header.X-Api-Key': True,
            'method.request.header.Accept': False
        }

        assert aws_method['requestParameters'] == expected_request_params, "AWS API requestParameters don't match expected values"

        assert method['requestParameters'] == aws_method['requestParameters'], "requestParameters don't match between cached and direct API calls"
