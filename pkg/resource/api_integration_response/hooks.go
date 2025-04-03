// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package api_integration_response

import (
	"github.com/aws-controllers-k8s/runtime/pkg/compare"
	svcsdk "github.com/aws/aws-sdk-go-v2/service/apigateway"

	"github.com/aws-controllers-k8s/apigateway-controller/pkg/util/patch"
)

func updateIntegrationResponseInput(desired, latest *resource, input *svcsdk.UpdateIntegrationResponseInput, delta *compare.Delta) {
	latestSpec := latest.ko.Spec
	desiredSpec := desired.ko.Spec

	var patchSet patch.Set
	if delta.DifferentAt("Spec.ContentHandling") {
		patchSet.Replace("/contentHandling", desiredSpec.ContentHandling)
	}
	if delta.DifferentAt("Spec.ResponseParameters") {
		patchSet.ForMap("/responseParameters", latestSpec.ResponseParameters, desiredSpec.ResponseParameters, true)
	}
	if delta.DifferentAt("Spec.ResponseTemplates") {
		patchSet.ForMap("/responseTemplates", latestSpec.ResponseTemplates, desiredSpec.ResponseTemplates, true)
	}
	if delta.DifferentAt("Spec.SelectionPattern") {
		patchSet.Replace("/selectionPattern", desiredSpec.SelectionPattern)
	}

	input.PatchOperations = patchSet.GetPatchOperations()
}
