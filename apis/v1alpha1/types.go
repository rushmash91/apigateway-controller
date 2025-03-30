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
	"github.com/aws/aws-sdk-go/aws"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = &aws.JSONValue{}
	_ = ackv1alpha1.AWSAccountID("")
)

// API stage name of the associated API stage in a usage plan.
type APIStage struct {
	APIID *string `json:"apiID,omitempty"`
	Stage *string `json:"stage,omitempty"`
}

// Access log settings, including the access log format and access log destination
// ARN.
type AccessLogSettings struct {
	DestinationARN *string `json:"destinationARN,omitempty"`
	Format         *string `json:"format,omitempty"`
}

// Configuration settings of a canary deployment.
type CanarySettings struct {
	DeploymentID           *string            `json:"deploymentID,omitempty"`
	PercentTraffic         *float64           `json:"percentTraffic,omitempty"`
	StageVariableOverrides map[string]*string `json:"stageVariableOverrides,omitempty"`
	UseStageCache          *bool              `json:"useStageCache,omitempty"`
}

// The input configuration for a canary deployment.
type DeploymentCanarySettings struct {
	PercentTraffic         *float64           `json:"percentTraffic,omitempty"`
	StageVariableOverrides map[string]*string `json:"stageVariableOverrides,omitempty"`
	UseStageCache          *bool              `json:"useStageCache,omitempty"`
}

// Specifies the target API entity to which the documentation applies.
type DocumentationPartLocation struct {
	Method *string `json:"method,omitempty"`
	Name   *string `json:"name,omitempty"`
	Path   *string `json:"path,omitempty"`
}

// Represents a domain name access association between an access association
// source and a private custom domain name. With a domain name access association,
// an access association source can invoke a private custom domain name while
// isolated from the public internet.
type DomainNameAccessAssociation struct {
	AccessAssociationSource        *string            `json:"accessAssociationSource,omitempty"`
	DomainNameAccessAssociationARN *string            `json:"domainNameAccessAssociationARN,omitempty"`
	DomainNameARN                  *string            `json:"domainNameARN,omitempty"`
	Tags                           map[string]*string `json:"tags,omitempty"`
}

// Represents a custom domain name as a user-friendly host name of an API (RestApi).
type DomainName_SDK struct {
	CertificateARN           *string      `json:"certificateARN,omitempty"`
	CertificateName          *string      `json:"certificateName,omitempty"`
	CertificateUploadDate    *metav1.Time `json:"certificateUploadDate,omitempty"`
	DistributionDomainName   *string      `json:"distributionDomainName,omitempty"`
	DistributionHostedZoneID *string      `json:"distributionHostedZoneID,omitempty"`
	DomainName               *string      `json:"domainName,omitempty"`
	DomainNameARN            *string      `json:"domainNameARN,omitempty"`
	DomainNameID             *string      `json:"domainNameID,omitempty"`
	DomainNameStatus         *string      `json:"domainNameStatus,omitempty"`
	DomainNameStatusMessage  *string      `json:"domainNameStatusMessage,omitempty"`
	// The endpoint configuration to indicate the types of endpoints an API (RestApi)
	// or its custom domain name (DomainName) has.
	EndpointConfiguration *EndpointConfiguration `json:"endpointConfiguration,omitempty"`
	ManagementPolicy      *string                `json:"managementPolicy,omitempty"`
	// The mutual TLS authentication configuration for a custom domain name. If
	// specified, API Gateway performs two-way authentication between the client
	// and the server. Clients must present a trusted certificate to access your
	// API.
	MutualTLSAuthentication             *MutualTLSAuthentication `json:"mutualTLSAuthentication,omitempty"`
	OwnershipVerificationCertificateARN *string                  `json:"ownershipVerificationCertificateARN,omitempty"`
	Policy                              *string                  `json:"policy,omitempty"`
	RegionalCertificateARN              *string                  `json:"regionalCertificateARN,omitempty"`
	RegionalCertificateName             *string                  `json:"regionalCertificateName,omitempty"`
	RegionalDomainName                  *string                  `json:"regionalDomainName,omitempty"`
	RegionalHostedZoneID                *string                  `json:"regionalHostedZoneID,omitempty"`
	SecurityPolicy                      *string                  `json:"securityPolicy,omitempty"`
	Tags                                map[string]*string       `json:"tags,omitempty"`
}

// The endpoint configuration to indicate the types of endpoints an API (RestApi)
// or its custom domain name (DomainName) has.
type EndpointConfiguration struct {
	Types          []*string `json:"types,omitempty"`
	VPCEndpointIDs []*string `json:"vpcEndpointIDs,omitempty"`
	// Reference field for VPCEndpointIDs
	VPCEndpointRefs []*ackv1alpha1.AWSResourceReferenceWrapper `json:"vpcEndpointRefs,omitempty"`
}

// Represents an integration response. The status code must map to an existing
// MethodResponse, and parameters and templates can be used to transform the
// back-end response.
type IntegrationResponse struct {
	ContentHandling    *string            `json:"contentHandling,omitempty"`
	ResponseParameters map[string]*string `json:"responseParameters,omitempty"`
	ResponseTemplates  map[string]*string `json:"responseTemplates,omitempty"`
	SelectionPattern   *string            `json:"selectionPattern,omitempty"`
	// The status code.
	StatusCode *string `json:"statusCode,omitempty"`
}

// Represents an HTTP, HTTP_PROXY, AWS, AWS_PROXY, or Mock integration.
type Integration_SDK struct {
	CacheKeyParameters   []*string                       `json:"cacheKeyParameters,omitempty"`
	CacheNamespace       *string                         `json:"cacheNamespace,omitempty"`
	ConnectionID         *string                         `json:"connectionID,omitempty"`
	ConnectionType       *string                         `json:"connectionType,omitempty"`
	ContentHandling      *string                         `json:"contentHandling,omitempty"`
	Credentials          *string                         `json:"credentials,omitempty"`
	HTTPMethod           *string                         `json:"httpMethod,omitempty"`
	IntegrationResponses map[string]*IntegrationResponse `json:"integrationResponses,omitempty"`
	PassthroughBehavior  *string                         `json:"passthroughBehavior,omitempty"`
	RequestParameters    map[string]*string              `json:"requestParameters,omitempty"`
	RequestTemplates     map[string]*string              `json:"requestTemplates,omitempty"`
	TimeoutInMillis      *int64                          `json:"timeoutInMillis,omitempty"`
	// Specifies the TLS configuration for an integration.
	TLSConfig *TLSConfig `json:"tlsConfig,omitempty"`
	// The integration type. The valid value is HTTP for integrating an API method
	// with an HTTP backend; AWS with any Amazon Web Services service endpoints;
	// MOCK for testing without actually invoking the backend; HTTP_PROXY for integrating
	// with the HTTP proxy integration; AWS_PROXY for integrating with the Lambda
	// proxy integration.
	Type *string `json:"type_,omitempty"`
	URI  *string `json:"uri,omitempty"`
}

// Represents a client-facing interface by which the client calls the API to
// access back-end resources. A Method resource is integrated with an Integration
// resource. Both consist of a request and one or more responses. The method
// request takes the client input that is passed to the back end through the
// integration request. A method response returns the output from the back end
// to the client through an integration response. A method request is embodied
// in a Method resource, whereas an integration request is embodied in an Integration
// resource. On the other hand, a method response is represented by a MethodResponse
// resource, whereas an integration response is represented by an IntegrationResponse
// resource.
type Method struct {
	APIKeyRequired      *bool     `json:"apiKeyRequired,omitempty"`
	AuthorizationScopes []*string `json:"authorizationScopes,omitempty"`
	AuthorizationType   *string   `json:"authorizationType,omitempty"`
	AuthorizerID        *string   `json:"authorizerID,omitempty"`
	HTTPMethod          *string   `json:"httpMethod,omitempty"`
	// Represents an HTTP, HTTP_PROXY, AWS, AWS_PROXY, or Mock integration.
	MethodIntegration  *Integration_SDK           `json:"methodIntegration,omitempty"`
	MethodResponses    map[string]*MethodResponse `json:"methodResponses,omitempty"`
	OperationName      *string                    `json:"operationName,omitempty"`
	RequestModels      map[string]*string         `json:"requestModels,omitempty"`
	RequestParameters  map[string]*bool           `json:"requestParameters,omitempty"`
	RequestValidatorID *string                    `json:"requestValidatorID,omitempty"`
}

// Represents a method response of a given HTTP status code returned to the
// client. The method response is passed from the back end through the associated
// integration response that can be transformed using a mapping template.
type MethodResponse struct {
	ResponseModels     map[string]*string `json:"responseModels,omitempty"`
	ResponseParameters map[string]*bool   `json:"responseParameters,omitempty"`
	// The status code.
	StatusCode *string `json:"statusCode,omitempty"`
}

// Specifies the method setting properties.
type MethodSetting struct {
	CacheDataEncrypted                     *bool    `json:"cacheDataEncrypted,omitempty"`
	CacheTTLInSeconds                      *int64   `json:"cacheTTLInSeconds,omitempty"`
	CachingEnabled                         *bool    `json:"cachingEnabled,omitempty"`
	DataTraceEnabled                       *bool    `json:"dataTraceEnabled,omitempty"`
	LoggingLevel                           *string  `json:"loggingLevel,omitempty"`
	MetricsEnabled                         *bool    `json:"metricsEnabled,omitempty"`
	RequireAuthorizationForCacheControl    *bool    `json:"requireAuthorizationForCacheControl,omitempty"`
	ThrottlingBurstLimit                   *int64   `json:"throttlingBurstLimit,omitempty"`
	ThrottlingRateLimit                    *float64 `json:"throttlingRateLimit,omitempty"`
	UnauthorizedCacheControlHeaderStrategy *string  `json:"unauthorizedCacheControlHeaderStrategy,omitempty"`
}

// Represents a summary of a Method resource, given a particular date and time.
type MethodSnapshot struct {
	APIKeyRequired    *bool   `json:"apiKeyRequired,omitempty"`
	AuthorizationType *string `json:"authorizationType,omitempty"`
}

// The mutual TLS authentication configuration for a custom domain name. If
// specified, API Gateway performs two-way authentication between the client
// and the server. Clients must present a trusted certificate to access your
// API.
type MutualTLSAuthentication struct {
	TruststoreURI      *string   `json:"truststoreURI,omitempty"`
	TruststoreVersion  *string   `json:"truststoreVersion,omitempty"`
	TruststoreWarnings []*string `json:"truststoreWarnings,omitempty"`
}

// The mutual TLS authentication configuration for a custom domain name. If
// specified, API Gateway performs two-way authentication between the client
// and the server. Clients must present a trusted certificate to access your
// API.
type MutualTLSAuthenticationInput struct {
	TruststoreURI     *string `json:"truststoreURI,omitempty"`
	TruststoreVersion *string `json:"truststoreVersion,omitempty"`
}

// For more information about supported patch operations, see Patch Operations
// (https://docs.aws.amazon.com/apigateway/latest/api/patch-operations.html).
type PatchOperation struct {
	From  *string `json:"from,omitempty"`
	Op    *string `json:"op,omitempty"`
	Path  *string `json:"path,omitempty"`
	Value *string `json:"value,omitempty"`
}

// Quotas configured for a usage plan.
type QuotaSettings struct {
	Limit  *int64 `json:"limit,omitempty"`
	Offset *int64 `json:"offset,omitempty"`
}

// Represents an API resource.
type Resource_SDK struct {
	ID       *string `json:"id,omitempty"`
	ParentID *string `json:"parentID,omitempty"`
	Path     *string `json:"path,omitempty"`
	PathPart *string `json:"pathPart,omitempty"`
}

// Represents a REST API.
type RestAPI_SDK struct {
	APIKeySource              *string      `json:"apiKeySource,omitempty"`
	BinaryMediaTypes          []*string    `json:"binaryMediaTypes,omitempty"`
	CreatedDate               *metav1.Time `json:"createdDate,omitempty"`
	Description               *string      `json:"description,omitempty"`
	DisableExecuteAPIEndpoint *bool        `json:"disableExecuteAPIEndpoint,omitempty"`
	// The endpoint configuration to indicate the types of endpoints an API (RestApi)
	// or its custom domain name (DomainName) has.
	EndpointConfiguration  *EndpointConfiguration `json:"endpointConfiguration,omitempty"`
	ID                     *string                `json:"id,omitempty"`
	MinimumCompressionSize *int64                 `json:"minimumCompressionSize,omitempty"`
	Name                   *string                `json:"name,omitempty"`
	Policy                 *string                `json:"policy,omitempty"`
	RootResourceID         *string                `json:"rootResourceID,omitempty"`
	Tags                   map[string]*string     `json:"tags,omitempty"`
	Version                *string                `json:"version,omitempty"`
	Warnings               []*string              `json:"warnings,omitempty"`
}

// A configuration property of an SDK type.
type SDKConfigurationProperty struct {
	DefaultValue *string `json:"defaultValue,omitempty"`
	Description  *string `json:"description,omitempty"`
	FriendlyName *string `json:"friendlyName,omitempty"`
	Name         *string `json:"name,omitempty"`
	Required     *bool   `json:"required,omitempty"`
}

// A reference to a unique stage identified in the format {restApiId}/{stage}.
type StageKey struct {
	RestAPIID *string `json:"restAPIID,omitempty"`
	StageName *string `json:"stageName,omitempty"`
}

// Represents a unique identifier for a version of a deployed RestApi that is
// callable by users.
type Stage_SDK struct {
	// Access log settings, including the access log format and access log destination
	// ARN.
	AccessLogSettings   *AccessLogSettings `json:"accessLogSettings,omitempty"`
	CacheClusterEnabled *bool              `json:"cacheClusterEnabled,omitempty"`
	// Returns the size of the CacheCluster.
	CacheClusterSize *string `json:"cacheClusterSize,omitempty"`
	// Returns the status of the CacheCluster.
	CacheClusterStatus *string `json:"cacheClusterStatus,omitempty"`
	// Configuration settings of a canary deployment.
	CanarySettings       *CanarySettings           `json:"canarySettings,omitempty"`
	ClientCertificateID  *string                   `json:"clientCertificateID,omitempty"`
	CreatedDate          *metav1.Time              `json:"createdDate,omitempty"`
	DeploymentID         *string                   `json:"deploymentID,omitempty"`
	Description          *string                   `json:"description,omitempty"`
	DocumentationVersion *string                   `json:"documentationVersion,omitempty"`
	LastUpdatedDate      *metav1.Time              `json:"lastUpdatedDate,omitempty"`
	MethodSettings       map[string]*MethodSetting `json:"methodSettings,omitempty"`
	StageName            *string                   `json:"stageName,omitempty"`
	Tags                 map[string]*string        `json:"tags,omitempty"`
	TracingEnabled       *bool                     `json:"tracingEnabled,omitempty"`
	Variables            map[string]*string        `json:"variables,omitempty"`
	WebACLARN            *string                   `json:"webACLARN,omitempty"`
}

// Specifies the TLS configuration for an integration.
type TLSConfig struct {
	InsecureSkipVerification *bool `json:"insecureSkipVerification,omitempty"`
}

// The API request rate limits.
type ThrottleSettings struct {
	BurstLimit *int64   `json:"burstLimit,omitempty"`
	RateLimit  *float64 `json:"rateLimit,omitempty"`
}

// An API Gateway VPC link for a RestApi to access resources in an Amazon Virtual
// Private Cloud (VPC).
type VPCLink_SDK struct {
	Description   *string            `json:"description,omitempty"`
	ID            *string            `json:"id,omitempty"`
	Name          *string            `json:"name,omitempty"`
	Status        *string            `json:"status,omitempty"`
	StatusMessage *string            `json:"statusMessage,omitempty"`
	Tags          map[string]*string `json:"tags,omitempty"`
	TargetARNs    []*string          `json:"targetARNs,omitempty"`
}
