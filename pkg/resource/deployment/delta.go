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

// Code generated by ack-generate. DO NOT EDIT.

package deployment

import (
	"bytes"
	"reflect"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	acktags "github.com/aws-controllers-k8s/runtime/pkg/tags"
)

// Hack to avoid import errors during build...
var (
	_ = &bytes.Buffer{}
	_ = &reflect.Method{}
	_ = &acktags.Tags{}
)

// newResourceDelta returns a new `ackcompare.Delta` used to compare two
// resources
func newResourceDelta(
	a *resource,
	b *resource,
) *ackcompare.Delta {
	delta := ackcompare.NewDelta()
	if (a == nil && b != nil) ||
		(a != nil && b == nil) {
		delta.Add("", a, b)
		return delta
	}

	if ackcompare.HasNilDifference(a.ko.Spec.CacheClusterEnabled, b.ko.Spec.CacheClusterEnabled) {
		delta.Add("Spec.CacheClusterEnabled", a.ko.Spec.CacheClusterEnabled, b.ko.Spec.CacheClusterEnabled)
	} else if a.ko.Spec.CacheClusterEnabled != nil && b.ko.Spec.CacheClusterEnabled != nil {
		if *a.ko.Spec.CacheClusterEnabled != *b.ko.Spec.CacheClusterEnabled {
			delta.Add("Spec.CacheClusterEnabled", a.ko.Spec.CacheClusterEnabled, b.ko.Spec.CacheClusterEnabled)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.CacheClusterSize, b.ko.Spec.CacheClusterSize) {
		delta.Add("Spec.CacheClusterSize", a.ko.Spec.CacheClusterSize, b.ko.Spec.CacheClusterSize)
	} else if a.ko.Spec.CacheClusterSize != nil && b.ko.Spec.CacheClusterSize != nil {
		if *a.ko.Spec.CacheClusterSize != *b.ko.Spec.CacheClusterSize {
			delta.Add("Spec.CacheClusterSize", a.ko.Spec.CacheClusterSize, b.ko.Spec.CacheClusterSize)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.CanarySettings, b.ko.Spec.CanarySettings) {
		delta.Add("Spec.CanarySettings", a.ko.Spec.CanarySettings, b.ko.Spec.CanarySettings)
	} else if a.ko.Spec.CanarySettings != nil && b.ko.Spec.CanarySettings != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.CanarySettings.PercentTraffic, b.ko.Spec.CanarySettings.PercentTraffic) {
			delta.Add("Spec.CanarySettings.PercentTraffic", a.ko.Spec.CanarySettings.PercentTraffic, b.ko.Spec.CanarySettings.PercentTraffic)
		} else if a.ko.Spec.CanarySettings.PercentTraffic != nil && b.ko.Spec.CanarySettings.PercentTraffic != nil {
			if *a.ko.Spec.CanarySettings.PercentTraffic != *b.ko.Spec.CanarySettings.PercentTraffic {
				delta.Add("Spec.CanarySettings.PercentTraffic", a.ko.Spec.CanarySettings.PercentTraffic, b.ko.Spec.CanarySettings.PercentTraffic)
			}
		}
		if len(a.ko.Spec.CanarySettings.StageVariableOverrides) != len(b.ko.Spec.CanarySettings.StageVariableOverrides) {
			delta.Add("Spec.CanarySettings.StageVariableOverrides", a.ko.Spec.CanarySettings.StageVariableOverrides, b.ko.Spec.CanarySettings.StageVariableOverrides)
		} else if len(a.ko.Spec.CanarySettings.StageVariableOverrides) > 0 {
			if !ackcompare.MapStringStringPEqual(a.ko.Spec.CanarySettings.StageVariableOverrides, b.ko.Spec.CanarySettings.StageVariableOverrides) {
				delta.Add("Spec.CanarySettings.StageVariableOverrides", a.ko.Spec.CanarySettings.StageVariableOverrides, b.ko.Spec.CanarySettings.StageVariableOverrides)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.CanarySettings.UseStageCache, b.ko.Spec.CanarySettings.UseStageCache) {
			delta.Add("Spec.CanarySettings.UseStageCache", a.ko.Spec.CanarySettings.UseStageCache, b.ko.Spec.CanarySettings.UseStageCache)
		} else if a.ko.Spec.CanarySettings.UseStageCache != nil && b.ko.Spec.CanarySettings.UseStageCache != nil {
			if *a.ko.Spec.CanarySettings.UseStageCache != *b.ko.Spec.CanarySettings.UseStageCache {
				delta.Add("Spec.CanarySettings.UseStageCache", a.ko.Spec.CanarySettings.UseStageCache, b.ko.Spec.CanarySettings.UseStageCache)
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Description, b.ko.Spec.Description) {
		delta.Add("Spec.Description", a.ko.Spec.Description, b.ko.Spec.Description)
	} else if a.ko.Spec.Description != nil && b.ko.Spec.Description != nil {
		if *a.ko.Spec.Description != *b.ko.Spec.Description {
			delta.Add("Spec.Description", a.ko.Spec.Description, b.ko.Spec.Description)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.RestAPIID, b.ko.Spec.RestAPIID) {
		delta.Add("Spec.RestAPIID", a.ko.Spec.RestAPIID, b.ko.Spec.RestAPIID)
	} else if a.ko.Spec.RestAPIID != nil && b.ko.Spec.RestAPIID != nil {
		if *a.ko.Spec.RestAPIID != *b.ko.Spec.RestAPIID {
			delta.Add("Spec.RestAPIID", a.ko.Spec.RestAPIID, b.ko.Spec.RestAPIID)
		}
	}
	if !reflect.DeepEqual(a.ko.Spec.RestAPIRef, b.ko.Spec.RestAPIRef) {
		delta.Add("Spec.RestAPIRef", a.ko.Spec.RestAPIRef, b.ko.Spec.RestAPIRef)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.StageDescription, b.ko.Spec.StageDescription) {
		delta.Add("Spec.StageDescription", a.ko.Spec.StageDescription, b.ko.Spec.StageDescription)
	} else if a.ko.Spec.StageDescription != nil && b.ko.Spec.StageDescription != nil {
		if *a.ko.Spec.StageDescription != *b.ko.Spec.StageDescription {
			delta.Add("Spec.StageDescription", a.ko.Spec.StageDescription, b.ko.Spec.StageDescription)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.StageName, b.ko.Spec.StageName) {
		delta.Add("Spec.StageName", a.ko.Spec.StageName, b.ko.Spec.StageName)
	} else if a.ko.Spec.StageName != nil && b.ko.Spec.StageName != nil {
		if *a.ko.Spec.StageName != *b.ko.Spec.StageName {
			delta.Add("Spec.StageName", a.ko.Spec.StageName, b.ko.Spec.StageName)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.TracingEnabled, b.ko.Spec.TracingEnabled) {
		delta.Add("Spec.TracingEnabled", a.ko.Spec.TracingEnabled, b.ko.Spec.TracingEnabled)
	} else if a.ko.Spec.TracingEnabled != nil && b.ko.Spec.TracingEnabled != nil {
		if *a.ko.Spec.TracingEnabled != *b.ko.Spec.TracingEnabled {
			delta.Add("Spec.TracingEnabled", a.ko.Spec.TracingEnabled, b.ko.Spec.TracingEnabled)
		}
	}
	if len(a.ko.Spec.Variables) != len(b.ko.Spec.Variables) {
		delta.Add("Spec.Variables", a.ko.Spec.Variables, b.ko.Spec.Variables)
	} else if len(a.ko.Spec.Variables) > 0 {
		if !ackcompare.MapStringStringPEqual(a.ko.Spec.Variables, b.ko.Spec.Variables) {
			delta.Add("Spec.Variables", a.ko.Spec.Variables, b.ko.Spec.Variables)
		}
	}

	return delta
}
