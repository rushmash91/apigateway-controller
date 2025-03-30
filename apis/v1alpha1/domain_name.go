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

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DomainNameSpec defines the desired state of DomainName.
//
// Represents a custom domain name as a user-friendly host name of an API (RestApi).
type DomainNameSpec struct {

	// The reference to an Amazon Web Services-managed certificate that will be
	// used by edge-optimized endpoint or private endpoint for this domain name.
	// Certificate Manager is the only supported source.
	CertificateARN *string `json:"certificateARN,omitempty"`
	// [Deprecated] The body of the server certificate that will be used by edge-optimized
	// endpoint or private endpoint for this domain name provided by your certificate
	// authority.
	CertificateBody *string `json:"certificateBody,omitempty"`
	// [Deprecated] The intermediate certificates and optionally the root certificate,
	// one after the other without any blank lines, used by an edge-optimized endpoint
	// for this domain name. If you include the root certificate, your certificate
	// chain must start with intermediate certificates and end with the root certificate.
	// Use the intermediate certificates that were provided by your certificate
	// authority. Do not include any intermediaries that are not in the chain of
	// trust path.
	CertificateChain *string `json:"certificateChain,omitempty"`
	// The user-friendly name of the certificate that will be used by edge-optimized
	// endpoint or private endpoint for this domain name.
	CertificateName *string `json:"certificateName,omitempty"`
	// [Deprecated] Your edge-optimized endpoint's domain name certificate's private
	// key.
	CertificatePrivateKey *string `json:"certificatePrivateKey,omitempty"`
	// The name of the DomainName resource.
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Value is immutable once set"
	// +kubebuilder:validation:Required
	DomainName *string `json:"domainName"`
	// The endpoint configuration of this DomainName showing the endpoint types
	// of the domain name.
	EndpointConfiguration   *EndpointConfiguration        `json:"endpointConfiguration,omitempty"`
	MutualTLSAuthentication *MutualTLSAuthenticationInput `json:"mutualTLSAuthentication,omitempty"`
	// The ARN of the public certificate issued by ACM to validate ownership of
	// your custom domain. Only required when configuring mutual TLS and using an
	// ACM imported or private CA certificate ARN as the regionalCertificateArn.
	OwnershipVerificationCertificateARN *string `json:"ownershipVerificationCertificateARN,omitempty"`
	// A stringified JSON policy document that applies to the execute-api service
	// for this DomainName regardless of the caller and Method configuration. Supported
	// only for private custom domain names.
	Policy *string `json:"policy,omitempty"`
	// The reference to an Amazon Web Services-managed certificate that will be
	// used by regional endpoint for this domain name. Certificate Manager is the
	// only supported source.
	RegionalCertificateARN *string `json:"regionalCertificateARN,omitempty"`
	// The user-friendly name of the certificate that will be used by regional endpoint
	// for this domain name.
	RegionalCertificateName *string `json:"regionalCertificateName,omitempty"`
	// The Transport Layer Security (TLS) version + cipher suite for this DomainName.
	// The valid values are TLS_1_0 and TLS_1_2.
	SecurityPolicy *string `json:"securityPolicy,omitempty"`
	// The key-value map of strings. The valid character set is [a-zA-Z+-=._:/].
	// The tag key can be up to 128 characters and must not start with aws:. The
	// tag value can be up to 256 characters.
	Tags map[string]*string `json:"tags,omitempty"`
}

// DomainNameStatus defines the observed state of DomainName
type DomainNameStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRs managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// The timestamp when the certificate that was used by edge-optimized endpoint
	// or private endpoint for this domain name was uploaded.
	// +kubebuilder:validation:Optional
	CertificateUploadDate *metav1.Time `json:"certificateUploadDate,omitempty"`
	// The domain name of the Amazon CloudFront distribution associated with this
	// custom domain name for an edge-optimized endpoint. You set up this association
	// when adding a DNS record pointing the custom domain name to this distribution
	// name. For more information about CloudFront distributions, see the Amazon
	// CloudFront documentation.
	// +kubebuilder:validation:Optional
	DistributionDomainName *string `json:"distributionDomainName,omitempty"`
	// The region-agnostic Amazon Route 53 Hosted Zone ID of the edge-optimized
	// endpoint. The valid value is Z2FDTNDATAQYW2 for all the regions. For more
	// information, see Set up a Regional Custom Domain Name and AWS Regions and
	// Endpoints for API Gateway.
	// +kubebuilder:validation:Optional
	DistributionHostedZoneID *string `json:"distributionHostedZoneID,omitempty"`
	// The identifier for the domain name resource. Supported only for private custom
	// domain names.
	// +kubebuilder:validation:Optional
	DomainNameID *string `json:"domainNameID,omitempty"`
	// The status of the DomainName migration. The valid values are AVAILABLE and
	// UPDATING. If the status is UPDATING, the domain cannot be modified further
	// until the existing operation is complete. If it is AVAILABLE, the domain
	// can be updated.
	// +kubebuilder:validation:Optional
	DomainNameStatus *string `json:"domainNameStatus,omitempty"`
	// An optional text message containing detailed information about status of
	// the DomainName migration.
	// +kubebuilder:validation:Optional
	DomainNameStatusMessage *string `json:"domainNameStatusMessage,omitempty"`
	// A stringified JSON policy document that applies to the API Gateway Management
	// service for this DomainName. This policy document controls access for access
	// association sources to create domain name access associations with this DomainName.
	// Supported only for private custom domain names.
	// +kubebuilder:validation:Optional
	ManagementPolicy *string `json:"managementPolicy,omitempty"`
	// The domain name associated with the regional endpoint for this custom domain
	// name. You set up this association by adding a DNS record that points the
	// custom domain name to this regional domain name. The regional domain name
	// is returned by API Gateway when you create a regional endpoint.
	// +kubebuilder:validation:Optional
	RegionalDomainName *string `json:"regionalDomainName,omitempty"`
	// The region-specific Amazon Route 53 Hosted Zone ID of the regional endpoint.
	// For more information, see Set up a Regional Custom Domain Name and AWS Regions
	// and Endpoints for API Gateway.
	// +kubebuilder:validation:Optional
	RegionalHostedZoneID *string `json:"regionalHostedZoneID,omitempty"`
}

// DomainName is the Schema for the DomainNames API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type DomainName struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DomainNameSpec   `json:"spec,omitempty"`
	Status            DomainNameStatus `json:"status,omitempty"`
}

// DomainNameList contains a list of DomainName
// +kubebuilder:object:root=true
type DomainNameList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DomainName `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DomainName{}, &DomainNameList{})
}
