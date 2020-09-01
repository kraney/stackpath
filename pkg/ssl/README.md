# Go API client for ssl

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientExperimentalCodegen
For more information, please visit [https://support.stackpath.com/](https://support.stackpath.com/)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./ssl"
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identifield by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://gateway.stackpath.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*CertificatesApi* | [**DeleteCertificate**](docs/CertificatesApi.md#deletecertificate) | **Delete** /ssl/v1/stacks/{stack_id}/certificates/{certificate_id} | Delete a certificate
*CertificatesApi* | [**GetCertificate**](docs/CertificatesApi.md#getcertificate) | **Get** /ssl/v1/stacks/{stack_id}/certificates/{certificate_id} | Get a certificate
*CertificatesApi* | [**GetCertificates**](docs/CertificatesApi.md#getcertificates) | **Get** /ssl/v1/stacks/{stack_id}/certificates | Get all certificates
*CertificatesApi* | [**GetLatestCertificate**](docs/CertificatesApi.md#getlatestcertificate) | **Get** /ssl/v1/stacks/{stack_id}/certificates/{certificate_id}/latest | Get a certificate&#39;s latest version
*CertificatesApi* | [**RenewCertificate**](docs/CertificatesApi.md#renewcertificate) | **Post** /ssl/v1/stacks/{stack_id}/certificates/{certificate_id}/renew | Renew a certificate
*CertificatesApi* | [**RevokeCertificate**](docs/CertificatesApi.md#revokecertificate) | **Post** /ssl/v1/stacks/{stack_id}/certificates/{certificate_id}/revoke | Revoke a certificate
*CertificatesApi* | [**UpdateCertificate**](docs/CertificatesApi.md#updatecertificate) | **Patch** /ssl/v1/stacks/{stack_id}/certificates/{certificate_id} | Update a certificate


## Documentation For Models

 - [ApiStatus](docs/ApiStatus.md)
 - [ApiStatusDetail](docs/ApiStatusDetail.md)
 - [CertificateCertificate](docs/CertificateCertificate.md)
 - [CertificateCertificateStatus](docs/CertificateCertificateStatus.md)
 - [CertificateCertificateVerificationMethod](docs/CertificateCertificateVerificationMethod.md)
 - [CertificateDnsRecord](docs/CertificateDnsRecord.md)
 - [CertificateDnsVerificationDetails](docs/CertificateDnsVerificationDetails.md)
 - [CertificateGetCertificateResponse](docs/CertificateGetCertificateResponse.md)
 - [CertificateGetCertificatesResponse](docs/CertificateGetCertificatesResponse.md)
 - [CertificateGetLatestCertificateResponse](docs/CertificateGetLatestCertificateResponse.md)
 - [CertificateHttpVerificationDetails](docs/CertificateHttpVerificationDetails.md)
 - [CertificateRenewCertificateResponse](docs/CertificateRenewCertificateResponse.md)
 - [CertificateUpdateCertificateRequest](docs/CertificateUpdateCertificateRequest.md)
 - [CertificateUpdateCertificateResponse](docs/CertificateUpdateCertificateResponse.md)
 - [CertificateVerificationRequirements](docs/CertificateVerificationRequirements.md)
 - [PaginationPageInfo](docs/PaginationPageInfo.md)
 - [StackpathRpcBadRequest](docs/StackpathRpcBadRequest.md)
 - [StackpathRpcBadRequestAllOf](docs/StackpathRpcBadRequestAllOf.md)
 - [StackpathRpcBadRequestFieldViolation](docs/StackpathRpcBadRequestFieldViolation.md)
 - [StackpathRpcHelp](docs/StackpathRpcHelp.md)
 - [StackpathRpcHelpAllOf](docs/StackpathRpcHelpAllOf.md)
 - [StackpathRpcHelpLink](docs/StackpathRpcHelpLink.md)
 - [StackpathRpcLocalizedMessage](docs/StackpathRpcLocalizedMessage.md)
 - [StackpathRpcLocalizedMessageAllOf](docs/StackpathRpcLocalizedMessageAllOf.md)
 - [StackpathRpcPreconditionFailure](docs/StackpathRpcPreconditionFailure.md)
 - [StackpathRpcPreconditionFailureAllOf](docs/StackpathRpcPreconditionFailureAllOf.md)
 - [StackpathRpcPreconditionFailureViolation](docs/StackpathRpcPreconditionFailureViolation.md)
 - [StackpathRpcQuotaFailure](docs/StackpathRpcQuotaFailure.md)
 - [StackpathRpcQuotaFailureAllOf](docs/StackpathRpcQuotaFailureAllOf.md)
 - [StackpathRpcQuotaFailureViolation](docs/StackpathRpcQuotaFailureViolation.md)
 - [StackpathRpcRequestInfo](docs/StackpathRpcRequestInfo.md)
 - [StackpathRpcRequestInfoAllOf](docs/StackpathRpcRequestInfoAllOf.md)
 - [StackpathRpcResourceInfo](docs/StackpathRpcResourceInfo.md)
 - [StackpathRpcResourceInfoAllOf](docs/StackpathRpcResourceInfoAllOf.md)
 - [StackpathRpcRetryInfo](docs/StackpathRpcRetryInfo.md)
 - [StackpathRpcRetryInfoAllOf](docs/StackpathRpcRetryInfoAllOf.md)


## Documentation For Authorization



### oauth2


- **Type**: OAuth
- **Flow**: application
- **Authorization URL**: 
- **Scopes**: N/A

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author


