# Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License"). You may
# not use this file except in compliance with the License. A copy of the
# License is located at
#
# http://aws.amazon.com/apache2.0/
#
# or in the "license" file accompanying this file. This file is distributed
# on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
# express or implied. See the License for the specific language governing
# permissions and limitations under the License.

"""Integration tests for the APIIntegrationResponse resource
"""

import logging
import time
from typing import Dict, Tuple
from functools import partial

import boto3
import pytest
from acktest.k8s import resource as k8s
from acktest.k8s import condition
from acktest.resources import random_suffix_name
from e2e import service_marker, CRD_GROUP, CRD_VERSION, SERVICE_NAME, load_apigateway_resource
from e2e.replacement_values import REPLACEMENT_VALUES
from e2e.common.waiter import wait_until_deleted, safe_get
from .rest_api_test import simple_rest_api
from .resource_test import simple_resource
from .integration_test import simple_integration

API_INTEGRATION_RESPONSE_RESOURCE_PLURAL = "apiintegrationresponses"
MODIFY_WAIT_AFTER_SECONDS = 60
MAX_WAIT_FOR_SYNCED_MINUTES = 1


@pytest.fixture(scope='module')
def apigateway_client():
    return boto3.client(SERVICE_NAME)


@pytest.fixture
def simple_api_integration_response(simple_integration, apigateway_client) -> Tuple[k8s.CustomResourceReference, Dict, Dict]:
    api_integration_response_name = random_suffix_name(
        'integration-response', 32)

    (int_ref, int_cr, resource_query) = simple_integration
    rest_api_id = resource_query['restApiId']
    resource_id = resource_query['resourceId']
    http_method = resource_query['httpMethod']

    method_response_params = {
        'restApiId': rest_api_id,
        'resourceId': resource_id,
        'httpMethod': http_method,
        'statusCode': '200'
    }
    apigateway_client.put_method_response(
        **method_response_params,
        responseParameters={
            'method.response.header.Content-Type': False,
            'method.response.header.X-Powered-By': False
        }
    )

    replacements = REPLACEMENT_VALUES.copy()
    replacements['API_INTEGRATION_RESPONSE_NAME'] = api_integration_response_name
    replacements['REST_API_REF_NAME'] = int_cr['spec']['restAPIRef']['from']['name']
    replacements['RESOURCE_REF_NAME'] = int_cr['spec']['resourceRef']['from']['name']

    resource_data = load_apigateway_resource(
        'api_integration_response_simple',
        additional_replacements=replacements,
    )
    logging.debug(resource_data)

    ref = k8s.CustomResourceReference(
        CRD_GROUP, CRD_VERSION, API_INTEGRATION_RESPONSE_RESOURCE_PLURAL,
        api_integration_response_name, namespace='default',
    )

    k8s.create_custom_resource(ref, resource_data)
    cr = k8s.wait_resource_consumed_by_controller(ref, wait_periods=30)

    assert cr is not None
    assert k8s.get_resource_exists(ref)

    # Create parameters for API Gateway queries
    integration_response_query = {
        'restApiId': rest_api_id,
        'resourceId': resource_id,
        'httpMethod': http_method,
        'statusCode': '200'
    }

    yield ref, cr, integration_response_query

    # Cleanup
    _, deleted = k8s.delete_custom_resource(ref, 10, 30)
    assert deleted
    wait_until_deleted(partial(
        apigateway_client.get_integration_response, **integration_response_query))

    apigateway_client.delete_method_response(**method_response_params)


@service_marker
@pytest.mark.canary
class TestAPIIntegrationResponse:
    def test_create_update_api_integration_response(self, simple_api_integration_response, apigateway_client):
        (ref, cr, integration_response_query) = simple_api_integration_response
        get_integration_response = partial(
            apigateway_client.get_integration_response, **integration_response_query)

        assert safe_get(get_integration_response) is not None

        updates = {
            'responseTemplates': {
                'application/json': '#set($inputRoot = $input.path(\'$\'))\n{\n  "data": $input.json(\'$\'),\n  "method": "$context.httpMethod",\n  "resourcePath": "$context.resourcePath",\n  "version": "$input.params(\'version\')",\n  "filter": "$input.params(\'filter\')"\n}'
            },
            'responseParameters': {
                'method.response.header.Content-Type': 'integration.response.header.Content-Type',
                'method.response.header.X-Powered-By': '\'AWS API Gateway\''
            }
        }
        k8s.patch_custom_resource(ref, {'spec': updates})
        time.sleep(MODIFY_WAIT_AFTER_SECONDS)

        # Verify updates in AWS
        aws_resource = get_integration_response()

        assert 'responseTemplates' in aws_resource
        assert aws_resource['responseTemplates'] == updates['responseTemplates']

        assert 'responseParameters' in aws_resource
        assert aws_resource['responseParameters'] == updates['responseParameters']

        # For completeness, also verify the entire update
        updated_fields = {field: aws_resource[field]
                          for field in updates.keys()}
        assert updated_fields == updates
