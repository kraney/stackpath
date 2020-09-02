# \PurgeContentApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPurgeStatus**](PurgeContentApi.md#GetPurgeStatus) | **Get** /cdn/v1/stacks/{stack_id}/purge/{purge_id} | Get purge status
[**PurgeContent**](PurgeContentApi.md#PurgeContent) | **Post** /cdn/v1/stacks/{stack_id}/purge | Purge content



## GetPurgeStatus

> CdnGetPurgeStatusResponse GetPurgeStatus(ctx, stackId, purgeId)

Get purge status

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**purgeId** | **string**| A CDN purge ID | 

### Return type

[**CdnGetPurgeStatusResponse**](cdnGetPurgeStatusResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PurgeContent

> CdnPurgeContentResponse PurgeContent(ctx, stackId, cdnPurgeContentRequest)

Purge content

Purge cached content for all CDN sites on a stack. Content is re-cached on the CDN the next time it is requested. Use the returned purge ID to see the status of a purge request.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**cdnPurgeContentRequest** | [**CdnPurgeContentRequest**](CdnPurgeContentRequest.md)|  | 

### Return type

[**CdnPurgeContentResponse**](cdnPurgeContentResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

