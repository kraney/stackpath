# \SSLApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConnectSiteToCertificate**](SSLApi.md#ConnectSiteToCertificate) | **Put** /waf/v1/stacks/{stack_id}/sites/{site_id}/delivery/certificates/{certificate_id} | Update SAN hosts
[**CreateCertificate**](SSLApi.md#CreateCertificate) | **Post** /waf/v1/stacks/{stack_id}/delivery/certificates | Add a certificate
[**DeleteCertificate**](SSLApi.md#DeleteCertificate) | **Delete** /waf/v1/stacks/{stack_id}/delivery/certificates/{certificate_id} | Delete a certificate
[**GetCertificate**](SSLApi.md#GetCertificate) | **Get** /waf/v1/stacks/{stack_id}/delivery/certificates/{certificate_id} | Get a certificate
[**GetCertificateVerificationDetails**](SSLApi.md#GetCertificateVerificationDetails) | **Get** /waf/v1/stacks/{stack_id}/delivery/certificates/{certificate_id}/verification_details | Get sites associated with a certificate
[**GetCertificates**](SSLApi.md#GetCertificates) | **Get** /waf/v1/stacks/{stack_id}/delivery/certificates | Get all certificates
[**GetSiteCertificates**](SSLApi.md#GetSiteCertificates) | **Get** /waf/v1/stacks/{stack_id}/sites/{site_id}/delivery/certificates | Get all site certificates
[**RenewCertificate**](SSLApi.md#RenewCertificate) | **Post** /waf/v1/stacks/{stack_id}/delivery/certificates/{certificate_id}/renew | Renew a certificate
[**RequestCertificate**](SSLApi.md#RequestCertificate) | **Post** /waf/v1/stacks/{stack_id}/sites/{site_id}/delivery/certificates/request | Request a free certificate
[**UpdateCertificate**](SSLApi.md#UpdateCertificate) | **Put** /waf/v1/stacks/{stack_id}/delivery/certificates/{certificate_id} | Update a certificate



## ConnectSiteToCertificate

> WafConnectSiteToCertificateResponse ConnectSiteToCertificate(ctx, stackId, siteId, certificateId)

Update SAN hosts

Association is performed without validating if the site has a hostname covered by the certificate. This is useful for preparation work required for getting a site ready for traffic.  If a certificate is uploaded which contains hostnames for sites, it will automatically be connected to those sites. If a hostname is added to a site which is covered by an SSL certificate, it will automatically be connected to the certificate.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
**certificateId** | **string**| A certificate ID | 

### Return type

[**WafConnectSiteToCertificateResponse**](wafConnectSiteToCertificateResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCertificate

> WafCreateCertificateResponse CreateCertificate(ctx, stackId, wafCreateCertificateRequest)

Add a certificate

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**wafCreateCertificateRequest** | [**WafCreateCertificateRequest**](WafCreateCertificateRequest.md)|  | 

### Return type

[**WafCreateCertificateResponse**](wafCreateCertificateResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCertificate

> DeleteCertificate(ctx, stackId, certificateId)

Delete a certificate

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**certificateId** | **string**| A certificate ID | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertificate

> WafGetCertificateResponse GetCertificate(ctx, stackId, certificateId)

Get a certificate

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**certificateId** | **string**| A certificate ID | 

### Return type

[**WafGetCertificateResponse**](wafGetCertificateResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertificateVerificationDetails

> WafGetCertificateVerificationDetailsResponse GetCertificateVerificationDetails(ctx, stackId, certificateId)

Get sites associated with a certificate

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**certificateId** | **string**| A certificate ID | 

### Return type

[**WafGetCertificateVerificationDetailsResponse**](wafGetCertificateVerificationDetailsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertificates

> WafGetCertificatesResponse GetCertificates(ctx, stackId, optional)

Get all certificates

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
 **optional** | ***GetCertificatesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCertificatesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageRequestFirst** | **optional.String**| The number of items desired. | 
 **pageRequestAfter** | **optional.String**| The cursor value after which data will be returned. | 
 **pageRequestFilter** | **optional.String**| SQL-style constraint filters. | 
 **pageRequestSortBy** | **optional.String**| Sort the response by the given field. | 

### Return type

[**WafGetCertificatesResponse**](wafGetCertificatesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteCertificates

> WafGetSiteCertificatesResponse GetSiteCertificates(ctx, stackId, siteId, optional)

Get all site certificates

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
 **optional** | ***GetSiteCertificatesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSiteCertificatesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageRequestFirst** | **optional.String**| The number of items desired. | 
 **pageRequestAfter** | **optional.String**| The cursor value after which data will be returned. | 
 **pageRequestFilter** | **optional.String**| SQL-style constraint filters. | 
 **pageRequestSortBy** | **optional.String**| Sort the response by the given field. | 

### Return type

[**WafGetSiteCertificatesResponse**](wafGetSiteCertificatesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenewCertificate

> RenewCertificate(ctx, stackId, certificateId)

Renew a certificate

StackPath automatically renews certificates that are 30 days from expiration. Call this to retry a renewal that previously failed.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**certificateId** | **string**| A certificate ID | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestCertificate

> WafRequestCertificateResponse RequestCertificate(ctx, stackId, siteId, wafRequestCertificateRequest)

Request a free certificate

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
**wafRequestCertificateRequest** | [**WafRequestCertificateRequest**](WafRequestCertificateRequest.md)|  | 

### Return type

[**WafRequestCertificateResponse**](wafRequestCertificateResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCertificate

> WafUpdateCertificateResponse UpdateCertificate(ctx, stackId, certificateId, wafUpdateCertificateRequest)

Update a certificate

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**certificateId** | **string**| A certificate ID | 
**wafUpdateCertificateRequest** | [**WafUpdateCertificateRequest**](WafUpdateCertificateRequest.md)|  | 

### Return type

[**WafUpdateCertificateResponse**](wafUpdateCertificateResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

