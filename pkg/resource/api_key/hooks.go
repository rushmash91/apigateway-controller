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
	"context"
	"fmt"
	"strings"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	"github.com/aws/aws-sdk-go-v2/aws"
	svcsdk "github.com/aws/aws-sdk-go-v2/service/apigateway"

	svcapitypes "github.com/aws-controllers-k8s/apigateway-controller/apis/v1alpha1"
	"github.com/aws-controllers-k8s/apigateway-controller/pkg/tags"
	"github.com/aws-controllers-k8s/apigateway-controller/pkg/util/patch"
)

func updateApiKeyInput(desired *resource, input *svcsdk.UpdateApiKeyInput, delta *ackcompare.Delta) {
	desiredSpec := desired.ko.Spec
	var patchSet patch.Set

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
		updateStageKeyPatches(&patchSet, desiredSpec.StageKeys, desiredSpec.StageKeys)
	}

	input.PatchOperations = patchSet.GetPatchOperations()
}

// updateStageKeyPatches adds patch operations for stage keys, handling both additions and removals.
// Each StageKey represents a REST API stage in the format "restApiId/stageName".
// The path needs to be JSON Pointer encoded (RFC 6901) where "/" becomes "~1"
// to avoid conflicts with path separators.
//
// Example:
//
//	StageKey{RestAPIID: "abc123", StageName: "prod"} becomes "/stages/abc123~1prod"
func updateStageKeyPatches(patchSet *patch.Set, latest, desired []*svcapitypes.StageKey) {
	latestMap := make(map[string]bool)
	desiredMap := make(map[string]bool)

	// Build desired stage keys map
	for _, sk := range desired {
		if sk.RestAPIID != nil && sk.StageName != nil {
			key := fmt.Sprintf("%s/%s", *sk.RestAPIID, *sk.StageName)
			desiredMap[key] = true
		}
	}

	// Build latest stage keys map (only for those that are still desired)
	for _, sk := range latest {
		if sk.RestAPIID != nil && sk.StageName != nil {
			key := fmt.Sprintf("%s/%s", *sk.RestAPIID, *sk.StageName)
			if desiredMap[key] {
				latestMap[key] = true
			}
		}
	}

	// Add new stage keys
	for key := range desiredMap {
		if !latestMap[key] {
			encodedKey := strings.Replace(key, "/", "~1", -1)
			patchSet.Add(fmt.Sprintf("/stages/%s", encodedKey), aws.String(""))
		}
	}

	// Remove stage keys that are no longer desired
	for _, sk := range latest {
		if sk.RestAPIID != nil && sk.StageName != nil {
			key := fmt.Sprintf("%s/%s", *sk.RestAPIID, *sk.StageName)
			if !desiredMap[key] {
				encodedKey := strings.Replace(key, "/", "~1", -1)
				patchSet.Remove(fmt.Sprintf("/stages/%s", encodedKey), nil)
			}
		}
	}
}

// syncApiKeyTags synchronizes tags between desired and latest resources
func updateTags(
	ctx context.Context,
	rm *resourceManager,
	desired *resource,
	latest *resource,
) error {
	resourceARN := fmt.Sprintf(
		"arn:aws:apigateway:%s::/apikeys/%s",
		*desired.ko.Status.ACKResourceMetadata.Region,
		*desired.ko.Status.ID,
	)
	return tags.SyncTags(ctx, rm.sdkapi, rm.metrics, resourceARN, desired.ko.Spec.Tags, latest.ko.Spec.Tags)
}

// getStageKeysFromStrings converts a slice of stage key strings in the format "restAPIID/stageName"
// to a slice of StageKey objects
func getStageKeysFromStrings(stageKeyStrings []string) []*svcapitypes.StageKey {
	stageKeys := make([]*svcapitypes.StageKey, 0, len(stageKeyStrings))
	for _, stageKeyStr := range stageKeyStrings {
		parts := strings.Split(stageKeyStr, "/")
		stageKeys = append(stageKeys, &svcapitypes.StageKey{
			RestAPIID: &parts[0],
			StageName: &parts[1],
		})
	}
	return stageKeys
}
