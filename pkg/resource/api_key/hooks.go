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
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
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

	// Tags are managed through separate TagResource/UntagResource APIs,
	// not through patch operations in UpdateApiKey

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

// syncTags keeps tags in sync by calling TagResource and UntagResource APIs
func (rm *resourceManager) syncTags(
	ctx context.Context,
	desired *resource,
	latest *resource,
) (err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.syncTags")
	defer func() {
		exit(err)
	}()

	if latest.ko.Status.ACKResourceMetadata == nil || latest.ko.Status.ACKResourceMetadata.ARN == nil {
		return fmt.Errorf("resource ARN is nil")
	}

	resourceARN := aws.String(string(*latest.ko.Status.ACKResourceMetadata.ARN))

	desiredTagsMap := desired.ko.Spec.Tags
	latestTagsMap := latest.ko.Spec.Tags

	desiredTags, _ := convertToOrderedACKTags(desiredTagsMap)
	latestTags, _ := convertToOrderedACKTags(latestTagsMap)

	added, updated, removed := ackcompare.GetTagsDifference(latestTags, desiredTags)

	// Combine added and updated tags
	toAdd := make(map[string]string)
	for k, v := range added {
		toAdd[k] = v
	}
	for k, v := range updated {
		toAdd[k] = v
	}

	var toRemoveTagKeys []string
	for k := range removed {
		toRemoveTagKeys = append(toRemoveTagKeys, k)
	}

	// Remove tags using UntagResource API:
	if len(toRemoveTagKeys) > 0 {
		rlog.Debug("removing tags from resource", "tags", toRemoveTagKeys)
		_, err = rm.sdkapi.UntagResource(
			ctx,
			&svcsdk.UntagResourceInput{
				ResourceArn: resourceARN,
				TagKeys:     toRemoveTagKeys,
			},
		)
		rm.metrics.RecordAPICall("UPDATE", "UntagResource", err)
		if err != nil {
			return err
		}
	}

	// Add tags using TagResource API
	if len(toAdd) > 0 {
		rlog.Debug("adding tags to resource", "tags", toAdd)
		_, err = rm.sdkapi.TagResource(
			ctx,
			&svcsdk.TagResourceInput{
				ResourceArn: resourceARN,
				Tags:        toAdd,
			},
		)
		rm.metrics.RecordAPICall("UPDATE", "TagResource", err)
		if err != nil {
			return err
		}
	}

	return nil
}

func compareTags(
	delta *ackcompare.Delta,
	a *resource,
	b *resource,
) {
	if len(a.ko.Spec.Tags) != len(b.ko.Spec.Tags) {
		delta.Add("Spec.Tags", a.ko.Spec.Tags, b.ko.Spec.Tags)
	} else if len(a.ko.Spec.Tags) > 0 {
		// Convert map[string]*string to acktags.Tags for GetTagsDifference
		aTagsConverted, _ := convertToOrderedACKTags(a.ko.Spec.Tags)
		bTagsConverted, _ := convertToOrderedACKTags(b.ko.Spec.Tags)

		added, updated, removed := ackcompare.GetTagsDifference(bTagsConverted, aTagsConverted)
		if len(added) != 0 || len(updated) != 0 || len(removed) != 0 {
			delta.Add("Spec.Tags", a.ko.Spec.Tags, b.ko.Spec.Tags)
		}
	}
}

// fetchCurrentTags returns the current tags for the resource
// using the GetTags API: GET /tags/resource_arn
func (rm *resourceManager) fetchCurrentTags(
	ctx context.Context,
	resourceARN *string,
) (map[string]string, error) {
	output, err := rm.sdkapi.GetTags(
		ctx,
		&svcsdk.GetTagsInput{
			ResourceArn: resourceARN,
		},
	)

	if err != nil {
		return nil, err
	}

	return output.Tags, nil
}

// CustomResourcesDifference helps return differences in custom resources
func (rm *resourceManager) CustomResourcesDifference(
	a *resource,
	b *resource,
) *ackcompare.Delta {
	delta := ackcompare.NewDelta()
	compareTags(delta, a, b)
	return delta
}
