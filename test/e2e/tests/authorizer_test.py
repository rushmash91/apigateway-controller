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

"""Integration tests for the Authorizer resource
"""

import logging
import time
import pytest
import boto3
from functools import partial

from acktest.resources import random_suffix_name
from acktest.k8s import resource as k8s
from e2e.bootstrap_resources import get_bootstrap_resources
from e2e import service_marker, CRD_GROUP, CRD_VERSION, load_apigateway_resource
from e2e.replacement_values import REPLACEMENT_VALUES
from .rest_api_test import simple_rest_api
from e2e.common.waiter import wait_until_deleted, safe_get

RESOURCE_PLURAL = "authorizers"
DEFAULT_WAIT_SECS = 10


@pytest.fixture(scope="function")
def authorizer_test_resources(simple_rest_api):
    _, rest_api_cr = simple_rest_api
    rest_api_id = rest_api_cr["status"].get("id")
    assert rest_api_id is not None

    bootstrap_resources = get_bootstrap_resources()

    user_pool_arn_1 = bootstrap_resources.AuthorizerUserPool1.user_pool_arn
    user_pool_arn_2 = bootstrap_resources.AuthorizerUserPool2.user_pool_arn

    authorizer_name = random_suffix_name("authorizer", 32)

    replacements_authorizer = REPLACEMENT_VALUES.copy()
    replacements_authorizer["AUTHORIZER_NAME"] = authorizer_name
    replacements_authorizer["REST_API_ID"] = rest_api_id
    replacements_authorizer["USER_POOL_ARN_1"] = user_pool_arn_1
    resource_data_authorizer = load_apigateway_resource(
        "authorizer_simple",
        additional_replacements=replacements_authorizer,
    )
    authorizer_ref = k8s.CustomResourceReference(
        CRD_GROUP, CRD_VERSION, RESOURCE_PLURAL,
        authorizer_name, namespace='default',
    )
    k8s.create_custom_resource(authorizer_ref, resource_data_authorizer)
    authorizer_cr = k8s.wait_resource_consumed_by_controller(authorizer_ref)
    assert authorizer_cr is not None
    assert k8s.get_resource_exists(authorizer_ref)
    authorizer_id = authorizer_cr["status"].get("id", None)
    assert authorizer_id is not None

    yield (authorizer_ref, authorizer_cr, authorizer_name, user_pool_arn_1, user_pool_arn_2)

    _, deleted_auth = k8s.delete_custom_resource(
        authorizer_ref, wait_periods=3, period_length=DEFAULT_WAIT_SECS)
    assert deleted_auth


@service_marker
@pytest.mark.canary
def test_authorizer_crud(authorizer_test_resources, apigateway_client):
    (authorizer_ref, authorizer_cr, authorizer_name,
     user_pool_arn_1, user_pool_arn_2) = authorizer_test_resources

    rest_api_id = authorizer_cr["spec"]["restAPIID"]
    authorizer_id = authorizer_cr["status"].get("id", None)
    assert authorizer_id is not None

    get_aws_authorizer = partial(
        apigateway_client.get_authorizer, restApiId=rest_api_id, authorizerId=authorizer_id)

    assert authorizer_cr is not None
    assert k8s.get_resource_exists(authorizer_ref)
    assert authorizer_cr["spec"].get("name") == authorizer_name
    assert authorizer_cr["spec"].get("type") == "COGNITO_USER_POOLS"
    assert authorizer_cr["spec"].get(
        "identitySource") == "method.request.header.Authorization"
    assert authorizer_cr["spec"].get("providerARNs") == [user_pool_arn_1]

    aws_authorizer = safe_get(get_aws_authorizer)
    assert aws_authorizer is not None, f"Authorizer {authorizer_id} not found in AWS API"
    assert aws_authorizer["name"] == authorizer_name
    assert aws_authorizer["type"] == "COGNITO_USER_POOLS"
    assert aws_authorizer["identitySource"] == "method.request.header.Authorization"
    assert aws_authorizer["providerARNs"] == [user_pool_arn_1]

    updated_authorizer_name = random_suffix_name("authorizer-updated", 32)
    updates = {
        "spec": {
            "name": updated_authorizer_name,
            "providerARNs": [user_pool_arn_1, user_pool_arn_2],
        },
    }
    k8s.patch_custom_resource(authorizer_ref, updates)
    time.sleep(DEFAULT_WAIT_SECS * 2)

    # Verify updated state
    updated_cr = k8s.get_resource(authorizer_ref)
    assert updated_cr["spec"].get("name") == updated_authorizer_name
    provider_arns_cr = updated_cr["spec"].get("providerARNs", [])
    assert len(provider_arns_cr) == 2
    assert user_pool_arn_1 in provider_arns_cr
    assert user_pool_arn_2 in provider_arns_cr

    aws_authorizer_updated = safe_get(get_aws_authorizer)
    assert aws_authorizer_updated is not None, f"Updated Authorizer {authorizer_id} not found in AWS API"
    assert aws_authorizer_updated["name"] == updated_authorizer_name
    provider_arns_aws = aws_authorizer_updated.get("providerARNs", [])
    assert len(provider_arns_aws) == 2
    assert user_pool_arn_1 in provider_arns_aws
    assert user_pool_arn_2 in provider_arns_aws
