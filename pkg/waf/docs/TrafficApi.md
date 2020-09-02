# \TrafficApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTraffic**](TrafficApi.md#GetTraffic) | **Get** /waf/v1/stacks/{stack_id}/traffic | LEGACY: Get traffic
[**GetTrafficV2**](TrafficApi.md#GetTrafficV2) | **Get** /waf/v2/stacks/{stack_id}/traffic | Get traffic



## GetTraffic

> WafGetTrafficResponse GetTraffic(ctx, stackId, optional)

LEGACY: Get traffic

Retrieve a report of a stack or site's WAF traffic.  **Note:** This endpoint is deprecated and will be removed in the future. Please use the [v2 get traffic call](ref:gettrafficv2) to retrieve WAF site traffic.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
 **optional** | ***GetTrafficOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTrafficOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **siteId** | **optional.String**| A site ID | 
 **startDate** | **optional.Time**| A lower bound date to search traffic for. | 
 **endDate** | **optional.Time**| An upper bound date to search traffic for. | 
 **resolution** | **optional.String**|  - HOURLY: All data points represent one hour of traffic  - MINUTELY: All data points represent one minute of traffic | [default to HOURLY]

### Return type

[**WafGetTrafficResponse**](wafGetTrafficResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrafficV2

> WafGetTrafficV2Response GetTrafficV2(ctx, stackId, optional)

Get traffic

Retreive a report of a stack or site's WAF traffic

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
 **optional** | ***GetTrafficV2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTrafficV2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **siteId** | **optional.String**| A site ID | 
 **startDate** | **optional.Time**| A lower bound date to search traffic for. | 
 **endDate** | **optional.Time**| An upper bound date to search traffic for. | 
 **resolution** | **optional.String**|  - DAILY: All data points represent one day of traffic  - HOURLY: All data points represent one hour of traffic  - MINUTELY: All data points represent one minute of traffic | [default to DAILY]

### Return type

[**WafGetTrafficV2Response**](wafGetTrafficV2Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

