# \CertificatesApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCertificate**](CertificatesApi.md#DeleteCertificate) | **Delete** /ssl/v1/stacks/{stack_id}/certificates/{certificate_id} | Delete a certificate
[**GetCertificate**](CertificatesApi.md#GetCertificate) | **Get** /ssl/v1/stacks/{stack_id}/certificates/{certificate_id} | Get a certificate
[**GetCertificates**](CertificatesApi.md#GetCertificates) | **Get** /ssl/v1/stacks/{stack_id}/certificates | Get all certificates
[**GetLatestCertificate**](CertificatesApi.md#GetLatestCertificate) | **Get** /ssl/v1/stacks/{stack_id}/certificates/{certificate_id}/latest | Get a certificate&#39;s latest version
[**RenewCertificate**](CertificatesApi.md#RenewCertificate) | **Post** /ssl/v1/stacks/{stack_id}/certificates/{certificate_id}/renew | Renew a certificate
[**RevokeCertificate**](CertificatesApi.md#RevokeCertificate) | **Post** /ssl/v1/stacks/{stack_id}/certificates/{certificate_id}/revoke | Revoke a certificate
[**UpdateCertificate**](CertificatesApi.md#UpdateCertificate) | **Patch** /ssl/v1/stacks/{stack_id}/certificates/{certificate_id} | Update a certificate



## DeleteCertificate

> DeleteCertificate(ctx, stackId, certificateId, optional)

Delete a certificate

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**certificateId** | **string**| A certificate ID | 
 **optional** | ***DeleteCertificateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteCertificateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reason** | **optional.String**|  | 

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

> CertificateGetCertificateResponse GetCertificate(ctx, stackId, certificateId)

Get a certificate

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**certificateId** | **string**| A certificate ID | 

### Return type

[**CertificateGetCertificateResponse**](certificateGetCertificateResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertificates

> CertificateGetCertificatesResponse GetCertificates(ctx, stackId, optional)

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

[**CertificateGetCertificatesResponse**](certificateGetCertificatesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLatestCertificate

> CertificateGetLatestCertificateResponse GetLatestCertificate(ctx, stackId, certificateId)

Get a certificate's latest version

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**certificateId** | **string**| A certificate ID | 

### Return type

[**CertificateGetLatestCertificateResponse**](certificateGetLatestCertificateResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenewCertificate

> CertificateRenewCertificateResponse RenewCertificate(ctx, stackId, certificateId)

Renew a certificate

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**certificateId** | **string**| A certificate ID | 

### Return type

[**CertificateRenewCertificateResponse**](certificateRenewCertificateResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeCertificate

> RevokeCertificate(ctx, stackId, certificateId)

Revoke a certificate

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


## UpdateCertificate

> CertificateUpdateCertificateResponse UpdateCertificate(ctx, stackId, certificateId, certificateUpdateCertificateRequest)

Update a certificate

Updating a certificate creates a new CSR and issues a new certificate.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**certificateId** | **string**| A certificate ID | 
**certificateUpdateCertificateRequest** | [**CertificateUpdateCertificateRequest**](CertificateUpdateCertificateRequest.md)|  | 

### Return type

[**CertificateUpdateCertificateResponse**](certificateUpdateCertificateResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

