# \RequestsApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRequest**](RequestsApi.md#GetRequest) | **Get** /waf/v1/stacks/{stack_id}/sites/{site_id}/requests/{request_id} | Get a request
[**GetRequestDetails**](RequestsApi.md#GetRequestDetails) | **Get** /waf/v1/stacks/{stack_id}/sites/{site_id}/requests/{request_id}/details | Get a request&#39;s details
[**GetRequests**](RequestsApi.md#GetRequests) | **Get** /waf/v1/stacks/{stack_id}/sites/{site_id}/requests | Get all requests



## GetRequest

> WafGetRequestResponse GetRequest(ctx, stackId, siteId, requestId)

Get a request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
**requestId** | **string**| A WAF request ID | 

### Return type

[**WafGetRequestResponse**](wafGetRequestResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRequestDetails

> WafGetRequestDetailsResponse GetRequestDetails(ctx, stackId, siteId, requestId)

Get a request's details

Retrieve more detailed information about a WAF request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
**requestId** | **string**| A WAF request ID | 

### Return type

[**WafGetRequestDetailsResponse**](wafGetRequestDetailsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRequests

> WafGetRequestsResponse GetRequests(ctx, stackId, siteId, optional)

Get all requests

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
 **optional** | ***GetRequestsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetRequestsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageRequestFirst** | **optional.String**| The number of items desired. | 
 **pageRequestAfter** | **optional.String**| The cursor value after which data will be returned. | 
 **pageRequestFilter** | **optional.String**| SQL-style constraint filters. | 
 **pageRequestSortBy** | **optional.String**| Sort the response by the given field. | 
 **startDate** | **optional.Time**| A lower bound date to search requests for. | 
 **endDate** | **optional.Time**| An upper bound date to search requests for. | 

### Return type

[**WafGetRequestsResponse**](wafGetRequestsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

