# \SSLApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConnectSiteToCertificate**](SSLApi.md#ConnectSiteToCertificate) | **Put** /cdn/v1/stacks/{stack_id}/sites/{site_id}/certificates/{certificate_id} | Get a site&#39;s certificate
[**CreateCertificate**](SSLApi.md#CreateCertificate) | **Post** /cdn/v1/stacks/{stack_id}/certificates | Add a certificate
[**DeleteCertificate**](SSLApi.md#DeleteCertificate) | **Delete** /cdn/v1/stacks/{stack_id}/certificates/{certificate_id} | Delete a certificate
[**GetCertificate**](SSLApi.md#GetCertificate) | **Get** /cdn/v1/stacks/{stack_id}/certificates/{certificate_id} | Get a certificate
[**GetCertificateSites**](SSLApi.md#GetCertificateSites) | **Get** /cdn/v1/stacks/{stack_id}/certificates/{certificate_id}/sites | Get sites associated with a certificate
[**GetCertificateVerificationDetails**](SSLApi.md#GetCertificateVerificationDetails) | **Get** /cdn/v1/stacks/{stack_id}/certificates/{certificate_id}/verification_details | Get verification details
[**GetCertificates**](SSLApi.md#GetCertificates) | **Get** /cdn/v1/stacks/{stack_id}/certificates | Get all certificates
[**GetSiteCertificates**](SSLApi.md#GetSiteCertificates) | **Get** /cdn/v1/stacks/{stack_id}/sites/{site_id}/certificates | Get all site certificates
[**RenewCertificate**](SSLApi.md#RenewCertificate) | **Post** /cdn/v1/stacks/{stack_id}/certificates/{certificate_id}/renew | Renew a certificate
[**RequestCertificate**](SSLApi.md#RequestCertificate) | **Post** /cdn/v1/stacks/{stack_id}/sites/{site_id}/certificates/request | Request a free certificate
[**UpdateCertificate**](SSLApi.md#UpdateCertificate) | **Put** /cdn/v1/stacks/{stack_id}/certificates/{certificate_id} | Update a certificate
[**UpdateSiteCertificateHosts**](SSLApi.md#UpdateSiteCertificateHosts) | **Put** /cdn/v1/stacks/{stack_id}/sites/{site_id}/certificates/{certificate_id}/hosts | Update SAN hosts



## ConnectSiteToCertificate

> CdnConnectSiteToCertificateResponse ConnectSiteToCertificate(ctx, stackId, siteId, certificateId)

Get a site's certificate

Association is performed without validating if the site has a hostname covered by the certificate. This is useful for preparation work required for getting a site ready for traffic.  If a certificate is uploaded which contains hostnames for sites, it will automatically be connected to those sites. If a hostname is added to a site which is covered by an SSL certificate, it will automatically be connected to the certificate.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
**certificateId** | **string**| A certificate ID | 

### Return type

[**CdnConnectSiteToCertificateResponse**](cdnConnectSiteToCertificateResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCertificate

> CdnCreateCertificateResponse CreateCertificate(ctx, stackId, cdnCreateCertificateRequest)

Add a certificate

The certificate is automatically associated with CDN site scope hostnames that match either the certificate's subject or its alternative names.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**cdnCreateCertificateRequest** | [**CdnCreateCertificateRequest**](CdnCreateCertificateRequest.md)|  | 

### Return type

[**CdnCreateCertificateResponse**](cdnCreateCertificateResponse.md)

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

> CdnGetCertificateResponse GetCertificate(ctx, stackId, certificateId)

Get a certificate

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**certificateId** | **string**| A certificate ID | 

### Return type

[**CdnGetCertificateResponse**](cdnGetCertificateResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertificateSites

> CdnGetCertificateSitesResponse GetCertificateSites(ctx, stackId, certificateId, optional)

Get sites associated with a certificate

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**certificateId** | **string**| A certificate ID | 
 **optional** | ***GetCertificateSitesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCertificateSitesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageRequestFirst** | **optional.String**| The number of items desired. | 
 **pageRequestAfter** | **optional.String**| The cursor value after which data will be returned. | 
 **pageRequestFilter** | **optional.String**| SQL-style constraint filters. | 
 **pageRequestSortBy** | **optional.String**| Sort the response by the given field. | 

### Return type

[**CdnGetCertificateSitesResponse**](cdnGetCertificateSitesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertificateVerificationDetails

> CdnGetCertificateVerificationDetailsResponse GetCertificateVerificationDetails(ctx, stackId, certificateId)

Get verification details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**certificateId** | **string**| A certificate ID | 

### Return type

[**CdnGetCertificateVerificationDetailsResponse**](cdnGetCertificateVerificationDetailsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertificates

> CdnGetCertificatesResponse GetCertificates(ctx, stackId, optional)

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

[**CdnGetCertificatesResponse**](cdnGetCertificatesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteCertificates

> CdnGetSiteCertificatesResponse GetSiteCertificates(ctx, stackId, siteId, optional)

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

[**CdnGetSiteCertificatesResponse**](cdnGetSiteCertificatesResponse.md)

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

> CdnRequestCertificateResponse RequestCertificate(ctx, stackId, siteId, cdnRequestCertificateRequest)

Request a free certificate

The optional list of hosts should be delivery domains for the site. If no hosts parameter is provided, all delivery domains for a site will be included in the SAN field. If the hosts parameter is provided, then the first entry in the list will be used as the certificate's common name.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
**cdnRequestCertificateRequest** | [**CdnRequestCertificateRequest**](CdnRequestCertificateRequest.md)|  | 

### Return type

[**CdnRequestCertificateResponse**](cdnRequestCertificateResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCertificate

> CdnUpdateCertificateResponse UpdateCertificate(ctx, stackId, certificateId, cdnUpdateCertificateRequest)

Update a certificate

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**certificateId** | **string**| A certificate ID | 
**cdnUpdateCertificateRequest** | [**CdnUpdateCertificateRequest**](CdnUpdateCertificateRequest.md)|  | 

### Return type

[**CdnUpdateCertificateResponse**](cdnUpdateCertificateResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSiteCertificateHosts

> UpdateSiteCertificateHosts(ctx, stackId, siteId, certificateId, cdnUpdateSiteCertificateHostsRequest)

Update SAN hosts

Updating hosts issues a new certificate.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
**certificateId** | **string**| A certificate ID | 
**cdnUpdateSiteCertificateHostsRequest** | [**CdnUpdateSiteCertificateHostsRequest**](CdnUpdateSiteCertificateHostsRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

