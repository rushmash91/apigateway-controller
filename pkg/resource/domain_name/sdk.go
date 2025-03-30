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

package domain_name

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackcondition "github.com/aws-controllers-k8s/runtime/pkg/condition"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	ackrequeue "github.com/aws-controllers-k8s/runtime/pkg/requeue"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	"github.com/aws/aws-sdk-go-v2/aws"
	svcsdk "github.com/aws/aws-sdk-go-v2/service/apigateway"
	svcsdktypes "github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	smithy "github.com/aws/smithy-go"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws-controllers-k8s/apigateway-controller/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = strings.ToLower("")
	_ = &svcsdk.Client{}
	_ = &svcapitypes.DomainName{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
	_ = &ackcondition.NotManagedMessage
	_ = &reflect.Value{}
	_ = fmt.Sprintf("")
	_ = &ackrequeue.NoRequeue{}
	_ = &aws.Config{}
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkFind")
	defer func() {
		exit(err)
	}()
	// If any required fields in the input shape are missing, AWS resource is
	// not created yet. Return NotFound here to indicate to callers that the
	// resource isn't yet created.
	if rm.requiredFieldsMissingFromReadOneInput(r) {
		return nil, ackerr.NotFound
	}

	input, err := rm.newDescribeRequestPayload(r)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.GetDomainNameOutput
	resp, err = rm.sdkapi.GetDomainName(ctx, input)
	rm.metrics.RecordAPICall("READ_ONE", "GetDomainName", err)
	if err != nil {
		var awsErr smithy.APIError
		if errors.As(err, &awsErr) && awsErr.ErrorCode() == "NotFoundException" {
			return nil, ackerr.NotFound
		}
		return nil, err
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.CertificateArn != nil {
		ko.Spec.CertificateARN = resp.CertificateArn
	} else {
		ko.Spec.CertificateARN = nil
	}
	if resp.CertificateName != nil {
		ko.Spec.CertificateName = resp.CertificateName
	} else {
		ko.Spec.CertificateName = nil
	}
	if resp.CertificateUploadDate != nil {
		ko.Status.CertificateUploadDate = &metav1.Time{*resp.CertificateUploadDate}
	} else {
		ko.Status.CertificateUploadDate = nil
	}
	if resp.DistributionDomainName != nil {
		ko.Status.DistributionDomainName = resp.DistributionDomainName
	} else {
		ko.Status.DistributionDomainName = nil
	}
	if resp.DistributionHostedZoneId != nil {
		ko.Status.DistributionHostedZoneID = resp.DistributionHostedZoneId
	} else {
		ko.Status.DistributionHostedZoneID = nil
	}
	if resp.DomainName != nil {
		ko.Spec.DomainName = resp.DomainName
	} else {
		ko.Spec.DomainName = nil
	}
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.DomainNameArn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.DomainNameArn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}
	if resp.DomainNameId != nil {
		ko.Status.DomainNameID = resp.DomainNameId
	} else {
		ko.Status.DomainNameID = nil
	}
	if resp.DomainNameStatus != "" {
		ko.Status.DomainNameStatus = aws.String(string(resp.DomainNameStatus))
	} else {
		ko.Status.DomainNameStatus = nil
	}
	if resp.DomainNameStatusMessage != nil {
		ko.Status.DomainNameStatusMessage = resp.DomainNameStatusMessage
	} else {
		ko.Status.DomainNameStatusMessage = nil
	}
	if resp.EndpointConfiguration != nil {
		f10 := &svcapitypes.EndpointConfiguration{}
		if resp.EndpointConfiguration.Types != nil {
			f10f0 := []*string{}
			for _, f10f0iter := range resp.EndpointConfiguration.Types {
				var f10f0elem *string
				f10f0elem = aws.String(string(f10f0iter))
				f10f0 = append(f10f0, f10f0elem)
			}
			f10.Types = f10f0
		}
		if resp.EndpointConfiguration.VpcEndpointIds != nil {
			f10.VPCEndpointIDs = aws.StringSlice(resp.EndpointConfiguration.VpcEndpointIds)
		}
		ko.Spec.EndpointConfiguration = f10
	} else {
		ko.Spec.EndpointConfiguration = nil
	}
	if resp.ManagementPolicy != nil {
		ko.Status.ManagementPolicy = resp.ManagementPolicy
	} else {
		ko.Status.ManagementPolicy = nil
	}
	if resp.MutualTlsAuthentication != nil {
		f12 := &svcapitypes.MutualTLSAuthenticationInput{}
		if resp.MutualTlsAuthentication.TruststoreUri != nil {
			f12.TruststoreURI = resp.MutualTlsAuthentication.TruststoreUri
		}
		if resp.MutualTlsAuthentication.TruststoreVersion != nil {
			f12.TruststoreVersion = resp.MutualTlsAuthentication.TruststoreVersion
		}
		ko.Spec.MutualTLSAuthentication = f12
	} else {
		ko.Spec.MutualTLSAuthentication = nil
	}
	if resp.OwnershipVerificationCertificateArn != nil {
		ko.Spec.OwnershipVerificationCertificateARN = resp.OwnershipVerificationCertificateArn
	} else {
		ko.Spec.OwnershipVerificationCertificateARN = nil
	}
	if resp.Policy != nil {
		ko.Spec.Policy = resp.Policy
	} else {
		ko.Spec.Policy = nil
	}
	if resp.RegionalCertificateArn != nil {
		ko.Spec.RegionalCertificateARN = resp.RegionalCertificateArn
	} else {
		ko.Spec.RegionalCertificateARN = nil
	}
	if resp.RegionalCertificateName != nil {
		ko.Spec.RegionalCertificateName = resp.RegionalCertificateName
	} else {
		ko.Spec.RegionalCertificateName = nil
	}
	if resp.RegionalDomainName != nil {
		ko.Status.RegionalDomainName = resp.RegionalDomainName
	} else {
		ko.Status.RegionalDomainName = nil
	}
	if resp.RegionalHostedZoneId != nil {
		ko.Status.RegionalHostedZoneID = resp.RegionalHostedZoneId
	} else {
		ko.Status.RegionalHostedZoneID = nil
	}
	if resp.SecurityPolicy != "" {
		ko.Spec.SecurityPolicy = aws.String(string(resp.SecurityPolicy))
	} else {
		ko.Spec.SecurityPolicy = nil
	}
	if resp.Tags != nil {
		ko.Spec.Tags = aws.StringMap(resp.Tags)
	} else {
		ko.Spec.Tags = nil
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// requiredFieldsMissingFromReadOneInput returns true if there are any fields
// for the ReadOne Input shape that are required but not present in the
// resource's Spec or Status
func (rm *resourceManager) requiredFieldsMissingFromReadOneInput(
	r *resource,
) bool {
	return r.ko.Spec.DomainName == nil

}

// newDescribeRequestPayload returns SDK-specific struct for the HTTP request
// payload of the Describe API call for the resource
func (rm *resourceManager) newDescribeRequestPayload(
	r *resource,
) (*svcsdk.GetDomainNameInput, error) {
	res := &svcsdk.GetDomainNameInput{}

	if r.ko.Spec.DomainName != nil {
		res.DomainName = r.ko.Spec.DomainName
	}
	if r.ko.Status.DomainNameID != nil {
		res.DomainNameId = r.ko.Status.DomainNameID
	}

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a copy of the resource with resource fields (in both Spec and
// Status) filled in with values from the CREATE API operation's Output shape.
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	desired *resource,
) (created *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkCreate")
	defer func() {
		exit(err)
	}()
	input, err := rm.newCreateRequestPayload(ctx, desired)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.CreateDomainNameOutput
	_ = resp
	resp, err = rm.sdkapi.CreateDomainName(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "CreateDomainName", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if resp.CertificateArn != nil {
		ko.Spec.CertificateARN = resp.CertificateArn
	} else {
		ko.Spec.CertificateARN = nil
	}
	if resp.CertificateName != nil {
		ko.Spec.CertificateName = resp.CertificateName
	} else {
		ko.Spec.CertificateName = nil
	}
	if resp.CertificateUploadDate != nil {
		ko.Status.CertificateUploadDate = &metav1.Time{*resp.CertificateUploadDate}
	} else {
		ko.Status.CertificateUploadDate = nil
	}
	if resp.DistributionDomainName != nil {
		ko.Status.DistributionDomainName = resp.DistributionDomainName
	} else {
		ko.Status.DistributionDomainName = nil
	}
	if resp.DistributionHostedZoneId != nil {
		ko.Status.DistributionHostedZoneID = resp.DistributionHostedZoneId
	} else {
		ko.Status.DistributionHostedZoneID = nil
	}
	if resp.DomainName != nil {
		ko.Spec.DomainName = resp.DomainName
	} else {
		ko.Spec.DomainName = nil
	}
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.DomainNameArn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.DomainNameArn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}
	if resp.DomainNameId != nil {
		ko.Status.DomainNameID = resp.DomainNameId
	} else {
		ko.Status.DomainNameID = nil
	}
	if resp.DomainNameStatus != "" {
		ko.Status.DomainNameStatus = aws.String(string(resp.DomainNameStatus))
	} else {
		ko.Status.DomainNameStatus = nil
	}
	if resp.DomainNameStatusMessage != nil {
		ko.Status.DomainNameStatusMessage = resp.DomainNameStatusMessage
	} else {
		ko.Status.DomainNameStatusMessage = nil
	}
	if resp.EndpointConfiguration != nil {
		f10 := &svcapitypes.EndpointConfiguration{}
		if resp.EndpointConfiguration.Types != nil {
			f10f0 := []*string{}
			for _, f10f0iter := range resp.EndpointConfiguration.Types {
				var f10f0elem *string
				f10f0elem = aws.String(string(f10f0iter))
				f10f0 = append(f10f0, f10f0elem)
			}
			f10.Types = f10f0
		}
		if resp.EndpointConfiguration.VpcEndpointIds != nil {
			f10.VPCEndpointIDs = aws.StringSlice(resp.EndpointConfiguration.VpcEndpointIds)
		}
		ko.Spec.EndpointConfiguration = f10
	} else {
		ko.Spec.EndpointConfiguration = nil
	}
	if resp.ManagementPolicy != nil {
		ko.Status.ManagementPolicy = resp.ManagementPolicy
	} else {
		ko.Status.ManagementPolicy = nil
	}
	if resp.MutualTlsAuthentication != nil {
		f12 := &svcapitypes.MutualTLSAuthenticationInput{}
		if resp.MutualTlsAuthentication.TruststoreUri != nil {
			f12.TruststoreURI = resp.MutualTlsAuthentication.TruststoreUri
		}
		if resp.MutualTlsAuthentication.TruststoreVersion != nil {
			f12.TruststoreVersion = resp.MutualTlsAuthentication.TruststoreVersion
		}
		ko.Spec.MutualTLSAuthentication = f12
	} else {
		ko.Spec.MutualTLSAuthentication = nil
	}
	if resp.OwnershipVerificationCertificateArn != nil {
		ko.Spec.OwnershipVerificationCertificateARN = resp.OwnershipVerificationCertificateArn
	} else {
		ko.Spec.OwnershipVerificationCertificateARN = nil
	}
	if resp.Policy != nil {
		ko.Spec.Policy = resp.Policy
	} else {
		ko.Spec.Policy = nil
	}
	if resp.RegionalCertificateArn != nil {
		ko.Spec.RegionalCertificateARN = resp.RegionalCertificateArn
	} else {
		ko.Spec.RegionalCertificateARN = nil
	}
	if resp.RegionalCertificateName != nil {
		ko.Spec.RegionalCertificateName = resp.RegionalCertificateName
	} else {
		ko.Spec.RegionalCertificateName = nil
	}
	if resp.RegionalDomainName != nil {
		ko.Status.RegionalDomainName = resp.RegionalDomainName
	} else {
		ko.Status.RegionalDomainName = nil
	}
	if resp.RegionalHostedZoneId != nil {
		ko.Status.RegionalHostedZoneID = resp.RegionalHostedZoneId
	} else {
		ko.Status.RegionalHostedZoneID = nil
	}
	if resp.SecurityPolicy != "" {
		ko.Spec.SecurityPolicy = aws.String(string(resp.SecurityPolicy))
	} else {
		ko.Spec.SecurityPolicy = nil
	}
	if resp.Tags != nil {
		ko.Spec.Tags = aws.StringMap(resp.Tags)
	} else {
		ko.Spec.Tags = nil
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.CreateDomainNameInput, error) {
	res := &svcsdk.CreateDomainNameInput{}

	if r.ko.Spec.CertificateARN != nil {
		res.CertificateArn = r.ko.Spec.CertificateARN
	}
	if r.ko.Spec.CertificateBody != nil {
		res.CertificateBody = r.ko.Spec.CertificateBody
	}
	if r.ko.Spec.CertificateChain != nil {
		res.CertificateChain = r.ko.Spec.CertificateChain
	}
	if r.ko.Spec.CertificateName != nil {
		res.CertificateName = r.ko.Spec.CertificateName
	}
	if r.ko.Spec.CertificatePrivateKey != nil {
		res.CertificatePrivateKey = r.ko.Spec.CertificatePrivateKey
	}
	if r.ko.Spec.DomainName != nil {
		res.DomainName = r.ko.Spec.DomainName
	}
	if r.ko.Spec.EndpointConfiguration != nil {
		f6 := &svcsdktypes.EndpointConfiguration{}
		if r.ko.Spec.EndpointConfiguration.Types != nil {
			f6f0 := []svcsdktypes.EndpointType{}
			for _, f6f0iter := range r.ko.Spec.EndpointConfiguration.Types {
				var f6f0elem string
				f6f0elem = string(*f6f0iter)
				f6f0 = append(f6f0, svcsdktypes.EndpointType(f6f0elem))
			}
			f6.Types = f6f0
		}
		if r.ko.Spec.EndpointConfiguration.VPCEndpointIDs != nil {
			f6.VpcEndpointIds = aws.ToStringSlice(r.ko.Spec.EndpointConfiguration.VPCEndpointIDs)
		}
		res.EndpointConfiguration = f6
	}
	if r.ko.Spec.MutualTLSAuthentication != nil {
		f7 := &svcsdktypes.MutualTlsAuthenticationInput{}
		if r.ko.Spec.MutualTLSAuthentication.TruststoreURI != nil {
			f7.TruststoreUri = r.ko.Spec.MutualTLSAuthentication.TruststoreURI
		}
		if r.ko.Spec.MutualTLSAuthentication.TruststoreVersion != nil {
			f7.TruststoreVersion = r.ko.Spec.MutualTLSAuthentication.TruststoreVersion
		}
		res.MutualTlsAuthentication = f7
	}
	if r.ko.Spec.OwnershipVerificationCertificateARN != nil {
		res.OwnershipVerificationCertificateArn = r.ko.Spec.OwnershipVerificationCertificateARN
	}
	if r.ko.Spec.Policy != nil {
		res.Policy = r.ko.Spec.Policy
	}
	if r.ko.Spec.RegionalCertificateARN != nil {
		res.RegionalCertificateArn = r.ko.Spec.RegionalCertificateARN
	}
	if r.ko.Spec.RegionalCertificateName != nil {
		res.RegionalCertificateName = r.ko.Spec.RegionalCertificateName
	}
	if r.ko.Spec.SecurityPolicy != nil {
		res.SecurityPolicy = svcsdktypes.SecurityPolicy(*r.ko.Spec.SecurityPolicy)
	}
	if r.ko.Spec.Tags != nil {
		res.Tags = aws.ToStringMap(r.ko.Spec.Tags)
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	delta *ackcompare.Delta,
) (updated *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkUpdate")
	defer func() {
		exit(err)
	}()
	input, err := rm.newUpdateRequestPayload(ctx, desired, delta)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.UpdateDomainNameOutput
	_ = resp
	resp, err = rm.sdkapi.UpdateDomainName(ctx, input)
	rm.metrics.RecordAPICall("UPDATE", "UpdateDomainName", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if resp.CertificateArn != nil {
		ko.Spec.CertificateARN = resp.CertificateArn
	} else {
		ko.Spec.CertificateARN = nil
	}
	if resp.CertificateName != nil {
		ko.Spec.CertificateName = resp.CertificateName
	} else {
		ko.Spec.CertificateName = nil
	}
	if resp.CertificateUploadDate != nil {
		ko.Status.CertificateUploadDate = &metav1.Time{*resp.CertificateUploadDate}
	} else {
		ko.Status.CertificateUploadDate = nil
	}
	if resp.DistributionDomainName != nil {
		ko.Status.DistributionDomainName = resp.DistributionDomainName
	} else {
		ko.Status.DistributionDomainName = nil
	}
	if resp.DistributionHostedZoneId != nil {
		ko.Status.DistributionHostedZoneID = resp.DistributionHostedZoneId
	} else {
		ko.Status.DistributionHostedZoneID = nil
	}
	if resp.DomainName != nil {
		ko.Spec.DomainName = resp.DomainName
	} else {
		ko.Spec.DomainName = nil
	}
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.DomainNameArn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.DomainNameArn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}
	if resp.DomainNameId != nil {
		ko.Status.DomainNameID = resp.DomainNameId
	} else {
		ko.Status.DomainNameID = nil
	}
	if resp.DomainNameStatus != "" {
		ko.Status.DomainNameStatus = aws.String(string(resp.DomainNameStatus))
	} else {
		ko.Status.DomainNameStatus = nil
	}
	if resp.DomainNameStatusMessage != nil {
		ko.Status.DomainNameStatusMessage = resp.DomainNameStatusMessage
	} else {
		ko.Status.DomainNameStatusMessage = nil
	}
	if resp.EndpointConfiguration != nil {
		f10 := &svcapitypes.EndpointConfiguration{}
		if resp.EndpointConfiguration.Types != nil {
			f10f0 := []*string{}
			for _, f10f0iter := range resp.EndpointConfiguration.Types {
				var f10f0elem *string
				f10f0elem = aws.String(string(f10f0iter))
				f10f0 = append(f10f0, f10f0elem)
			}
			f10.Types = f10f0
		}
		if resp.EndpointConfiguration.VpcEndpointIds != nil {
			f10.VPCEndpointIDs = aws.StringSlice(resp.EndpointConfiguration.VpcEndpointIds)
		}
		ko.Spec.EndpointConfiguration = f10
	} else {
		ko.Spec.EndpointConfiguration = nil
	}
	if resp.ManagementPolicy != nil {
		ko.Status.ManagementPolicy = resp.ManagementPolicy
	} else {
		ko.Status.ManagementPolicy = nil
	}
	if resp.MutualTlsAuthentication != nil {
		f12 := &svcapitypes.MutualTLSAuthenticationInput{}
		if resp.MutualTlsAuthentication.TruststoreUri != nil {
			f12.TruststoreURI = resp.MutualTlsAuthentication.TruststoreUri
		}
		if resp.MutualTlsAuthentication.TruststoreVersion != nil {
			f12.TruststoreVersion = resp.MutualTlsAuthentication.TruststoreVersion
		}
		ko.Spec.MutualTLSAuthentication = f12
	} else {
		ko.Spec.MutualTLSAuthentication = nil
	}
	if resp.OwnershipVerificationCertificateArn != nil {
		ko.Spec.OwnershipVerificationCertificateARN = resp.OwnershipVerificationCertificateArn
	} else {
		ko.Spec.OwnershipVerificationCertificateARN = nil
	}
	if resp.Policy != nil {
		ko.Spec.Policy = resp.Policy
	} else {
		ko.Spec.Policy = nil
	}
	if resp.RegionalCertificateArn != nil {
		ko.Spec.RegionalCertificateARN = resp.RegionalCertificateArn
	} else {
		ko.Spec.RegionalCertificateARN = nil
	}
	if resp.RegionalCertificateName != nil {
		ko.Spec.RegionalCertificateName = resp.RegionalCertificateName
	} else {
		ko.Spec.RegionalCertificateName = nil
	}
	if resp.RegionalDomainName != nil {
		ko.Status.RegionalDomainName = resp.RegionalDomainName
	} else {
		ko.Status.RegionalDomainName = nil
	}
	if resp.RegionalHostedZoneId != nil {
		ko.Status.RegionalHostedZoneID = resp.RegionalHostedZoneId
	} else {
		ko.Status.RegionalHostedZoneID = nil
	}
	if resp.SecurityPolicy != "" {
		ko.Spec.SecurityPolicy = aws.String(string(resp.SecurityPolicy))
	} else {
		ko.Spec.SecurityPolicy = nil
	}
	if resp.Tags != nil {
		ko.Spec.Tags = aws.StringMap(resp.Tags)
	} else {
		ko.Spec.Tags = nil
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// newUpdateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Update API call for the resource
func (rm *resourceManager) newUpdateRequestPayload(
	ctx context.Context,
	r *resource,
	delta *ackcompare.Delta,
) (*svcsdk.UpdateDomainNameInput, error) {
	res := &svcsdk.UpdateDomainNameInput{}

	if r.ko.Spec.DomainName != nil {
		res.DomainName = r.ko.Spec.DomainName
	}
	if r.ko.Status.DomainNameID != nil {
		res.DomainNameId = r.ko.Status.DomainNameID
	}

	return res, nil
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkDelete")
	defer func() {
		exit(err)
	}()
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return nil, err
	}
	var resp *svcsdk.DeleteDomainNameOutput
	_ = resp
	resp, err = rm.sdkapi.DeleteDomainName(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "DeleteDomainName", err)
	return nil, err
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteDomainNameInput, error) {
	res := &svcsdk.DeleteDomainNameInput{}

	if r.ko.Spec.DomainName != nil {
		res.DomainName = r.ko.Spec.DomainName
	}
	if r.ko.Status.DomainNameID != nil {
		res.DomainNameId = r.ko.Status.DomainNameID
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.DomainName,
) {
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if ko.Status.ACKResourceMetadata.Region == nil {
		ko.Status.ACKResourceMetadata.Region = &rm.awsRegion
	}
	if ko.Status.ACKResourceMetadata.OwnerAccountID == nil {
		ko.Status.ACKResourceMetadata.OwnerAccountID = &rm.awsAccountID
	}
	if ko.Status.Conditions == nil {
		ko.Status.Conditions = []*ackv1alpha1.Condition{}
	}
}

// updateConditions returns updated resource, true; if conditions were updated
// else it returns nil, false
func (rm *resourceManager) updateConditions(
	r *resource,
	onSuccess bool,
	err error,
) (*resource, bool) {
	ko := r.ko.DeepCopy()
	rm.setStatusDefaults(ko)

	// Terminal condition
	var terminalCondition *ackv1alpha1.Condition = nil
	var recoverableCondition *ackv1alpha1.Condition = nil
	var syncCondition *ackv1alpha1.Condition = nil
	for _, condition := range ko.Status.Conditions {
		if condition.Type == ackv1alpha1.ConditionTypeTerminal {
			terminalCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeRecoverable {
			recoverableCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeResourceSynced {
			syncCondition = condition
		}
	}
	var termError *ackerr.TerminalError
	if rm.terminalAWSError(err) || err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound || errors.As(err, &termError) {
		if terminalCondition == nil {
			terminalCondition = &ackv1alpha1.Condition{
				Type: ackv1alpha1.ConditionTypeTerminal,
			}
			ko.Status.Conditions = append(ko.Status.Conditions, terminalCondition)
		}
		var errorMessage = ""
		if err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound || errors.As(err, &termError) {
			errorMessage = err.Error()
		} else {
			awsErr, _ := ackerr.AWSError(err)
			errorMessage = awsErr.Error()
		}
		terminalCondition.Status = corev1.ConditionTrue
		terminalCondition.Message = &errorMessage
	} else {
		// Clear the terminal condition if no longer present
		if terminalCondition != nil {
			terminalCondition.Status = corev1.ConditionFalse
			terminalCondition.Message = nil
		}
		// Handling Recoverable Conditions
		if err != nil {
			if recoverableCondition == nil {
				// Add a new Condition containing a non-terminal error
				recoverableCondition = &ackv1alpha1.Condition{
					Type: ackv1alpha1.ConditionTypeRecoverable,
				}
				ko.Status.Conditions = append(ko.Status.Conditions, recoverableCondition)
			}
			recoverableCondition.Status = corev1.ConditionTrue
			awsErr, _ := ackerr.AWSError(err)
			errorMessage := err.Error()
			if awsErr != nil {
				errorMessage = awsErr.Error()
			}
			recoverableCondition.Message = &errorMessage
		} else if recoverableCondition != nil {
			recoverableCondition.Status = corev1.ConditionFalse
			recoverableCondition.Message = nil
		}
	}
	// Required to avoid the "declared but not used" error in the default case
	_ = syncCondition
	if terminalCondition != nil || recoverableCondition != nil || syncCondition != nil {
		return &resource{ko}, true // updated
	}
	return nil, false // not updated
}

// terminalAWSError returns awserr, true; if the supplied error is an aws Error type
// and if the exception indicates that it is a Terminal exception
// 'Terminal' exception are specified in generator configuration
func (rm *resourceManager) terminalAWSError(err error) bool {
	if err == nil {
		return false
	}

	var terminalErr smithy.APIError
	if !errors.As(err, &terminalErr) {
		return false
	}
	switch terminalErr.ErrorCode() {
	case "BadRequestException",
		"ConflictException",
		"NotFoundException",
		"InvalidParameter":
		return true
	default:
		return false
	}
}
