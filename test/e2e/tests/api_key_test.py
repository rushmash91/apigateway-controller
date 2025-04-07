import pytest
import logging
import time
from typing import Dict, Tuple

from acktest.k8s import resource as k8s
from acktest.resources import random_suffix_name
from e2e import service_marker, CRD_GROUP, CRD_VERSION, load_apigateway_resource
from e2e.replacement_values import REPLACEMENT_VALUES
from e2e.bootstrap_resources import get_bootstrap_resources
from e2e.service_bootstrap import APIGW

API_KEY_RESOURCE_PLURAL = "apikeys"
MAX_WAIT_FOR_SYNCED_MINUTES = 10


@pytest.fixture
def simple_stage(simple_rest_api, apigateway_client) -> Tuple[k8s.CustomResourceReference, Dict]:
    stage_name = random_suffix_name("simple-stage", 32)

    replacements = REPLACEMENT_VALUES.copy()
    replacements["STAGE_NAME"] = stage_name
    (ref, rest_api_cr) = simple_rest_api
    replacements["REST_API_ID"] = rest_api_cr["status"]["id"]

    resource_data = load_apigateway_resource(
        "stage_simple",
        additional_replacements=replacements,
    )

    ref = k8s.CustomResourceReference(
        CRD_GROUP,
        CRD_VERSION,
        "stages",
        stage_name,
        namespace="default",
    )
    k8s.create_custom_resource(ref, resource_data)
    cr = k8s.wait_resource_consumed_by_controller(ref)

    assert cr is not None
    assert k8s.get_resource_exists(ref)

    yield ref, cr

    # Delete the stage
    _, deleted = k8s.delete_custom_resource(ref, 3, 10)
    assert deleted


@pytest.fixture
def simple_api_key(simple_stage, apigateway_client) -> Tuple[k8s.CustomResourceReference, Dict]:
    api_key_name = random_suffix_name("simple-api-key", 32)

    replacements = REPLACEMENT_VALUES.copy()
    replacements["API_KEY_NAME"] = api_key_name

    # Get Rest API ID and Stage Name from the stage fixture
    (stage_ref, stage_cr) = simple_stage
    (rest_api_ref, rest_api_cr) = simple_stage[0]._replacement

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

        assert aws_api_key["tags"]["k1"] == "v1"
        assert aws_api_key["tags"]["k2"] == "v2"


        update_description = "Updated API key description"
        update_tags = {"k1": "updated-v1", "k3": "v3"}

        cr["spec"]["description"] = update_description
        cr["spec"]["tags"] = update_tags
        cr["spec"]["enabled"] = False

        k8s.patch_custom_resource(ref, cr)
        updated_cr = k8s.wait_resource_consumed_by_controller(ref)

        updated_aws_api_key = apigateway_client.get_api_key(
            apiKey=api_key_id,
            includeValue=True
        )

        assert updated_aws_api_key["description"] == update_description
        assert updated_aws_api_key["enabled"] == False
        assert updated_aws_api_key["tags"]["k1"] == "updated-v1"
        assert updated_aws_api_key["tags"]["k3"] == "v3"
        assert "k2" not in updated_aws_api_key["tags"]
