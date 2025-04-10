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

"""Integration tests for the API Key resource
"""

import pytest
import logging
import time
from typing import Dict, Tuple

from acktest.k8s import resource as k8s
from acktest.resources import random_suffix_name
from acktest import tags
from e2e import service_marker, CRD_GROUP, CRD_VERSION, load_apigateway_resource
from e2e.replacement_values import REPLACEMENT_VALUES
from e2e.bootstrap_resources import get_bootstrap_resources
from .rest_api_test import simple_rest_api
from .resource_test import simple_resource
from .integration_test import simple_integration
from .stage_test import simple_stage

API_KEY_RESOURCE_PLURAL = "apikeys"
MODIFY_WAIT_AFTER_SECONDS = 60


@pytest.fixture(scope='module')
def simple_api_key(simple_stage, simple_rest_api, apigateway_client) -> Tuple[k8s.CustomResourceReference, Dict]:
    api_key_name = random_suffix_name("simple-api-key", 32)

    replacements = REPLACEMENT_VALUES.copy()
    replacements["API_KEY_NAME"] = api_key_name

    (stage_ref, stage_cr, rest_api_id) = simple_stage
    (rest_api_ref, rest_api_cr) = simple_rest_api

    replacements["REST_API_ID"] = rest_api_cr["status"]["id"]
    replacements["STAGE_NAME"] = stage_cr["spec"]["stageName"]

    resource_data = load_apigateway_resource(
        "api_key_simple",
        additional_replacements=replacements,
    )

    ref = k8s.CustomResourceReference(
        CRD_GROUP,
        CRD_VERSION,
        API_KEY_RESOURCE_PLURAL,
        api_key_name,
        namespace="default",
    )
    k8s.create_custom_resource(ref, resource_data)
    cr = k8s.wait_resource_consumed_by_controller(ref)

    assert cr is not None
    assert k8s.get_resource_exists(ref)

    cr = k8s.get_resource(ref)
    yield ref, cr

    # Delete the API key
    _, deleted = k8s.delete_custom_resource(ref, 3, 10)
    assert deleted


@service_marker
@pytest.mark.canary
class TestAPIKey:
    def test_create_update_api_key(self, simple_api_key, apigateway_client):
        (ref, cr) = simple_api_key

        api_key_id = cr["status"]["id"]
        assert api_key_id is not None

        aws_api_key = apigateway_client.get_api_key(
            apiKey=api_key_id,
            includeValue=True
        )

        assert aws_api_key["name"] == cr["spec"]["name"]
        assert aws_api_key["description"] == "API Key for testing"
        assert aws_api_key["enabled"] == True

        # Validate tags
        latest_tags = aws_api_key["tags"]
        tags.assert_ack_system_tags(
            tags=latest_tags,
        )

        assert 'tags' in cr['spec']
        user_tags = cr["spec"]["tags"]
        tags.assert_equal_without_ack_tags(
            expected=user_tags,
            actual=latest_tags,
        )

        update_tags = {"k1": "updated-v1", "k3": "v3"}

        cr["spec"]["tags"] = update_tags
        cr["spec"]["enabled"] = False

        k8s.patch_custom_resource(ref, cr)
        updated_cr = k8s.wait_resource_consumed_by_controller(ref)

        time.sleep(MODIFY_WAIT_AFTER_SECONDS)
        updated_aws_api_key = apigateway_client.get_api_key(
            apiKey=api_key_id,
            includeValue=True
        )

        # Validate updated tags
        latest_tags = updated_aws_api_key["tags"]
        tags.assert_ack_system_tags(
            tags=latest_tags,
        )

        user_tags = updated_cr["spec"]["tags"]
        tags.assert_equal_without_ack_tags(
            expected=user_tags,
            actual=latest_tags,
        )

        assert updated_aws_api_key["enabled"] == False
