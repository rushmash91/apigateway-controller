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

package api_key

import (
	"fmt"
	"strings"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	"github.com/aws/aws-sdk-go-v2/aws"
	svcsdk "github.com/aws/aws-sdk-go-v2/service/apigateway"
	svcsdktypes "github.com/aws/aws-sdk-go-v2/service/apigateway/types"

	"github.com/aws-controllers-k8s/apigateway-controller/pkg/util/patch"
)

func updateApiKeyInput(desired *resource, input *svcsdk.UpdateApiKeyInput, delta *ackcompare.Delta) {
	desiredSpec := desired.ko.Spec
	var patchSet patch.Set
	var stageKeyPatches []svcsdktypes.PatchOperation

	if delta.DifferentAt("Spec.Name") {
		patchSet.Replace("/name", desiredSpec.Name)
	}
	if delta.DifferentAt("Spec.Description") {
		patchSet.Replace("/description", desiredSpec.Description)
	}
	if delta.DifferentAt("Spec.Enabled") {
		patchSet.Replace("/enabled", aws.String(fmt.Sprintf("%t", *desiredSpec.Enabled)))
	}
	if delta.DifferentAt("Spec.CustomerID") {
		patchSet.Replace("/customerId", desiredSpec.CustomerID)
	}

	// Handle StageKeys with add/remove operations
	if delta.DifferentAt("Spec.StageKeys") && desiredSpec.StageKeys != nil {
		// Convert StageKey objects to strings in the format "restApiId/stageName"
		for _, sk := range desiredSpec.StageKeys {
			if sk.RestAPIID != nil && sk.StageName != nil {
				// Format: restApiId/stageName
				stageKeyStr := fmt.Sprintf("%s/%s", *sk.RestAPIID, *sk.StageName)
				// Encode the path - replace / with ~1 as per JSON Patch spec
				encodedStageKey := strings.Replace(stageKeyStr, "/", "~1", -1)
				// For stages, we need to use add operation to /stages/{encoded-stage-key}
				stageKeyPatches = append(stageKeyPatches, svcsdktypes.PatchOperation{
					Op:    svcsdktypes.OpAdd,
					Path:  aws.String(fmt.Sprintf("/stages/%s", encodedStageKey)),
					Value: aws.String(""),
				})
			}
		}
	}

	patchOps := patchSet.GetPatchOperations()
	input.PatchOperations = append(patchOps, stageKeyPatches...)
}
