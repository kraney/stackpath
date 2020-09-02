# \MetricsApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMetrics**](MetricsApi.md#GetMetrics) | **Get** /cdn/v1/stacks/{stack_id}/metrics | Get metrics



## GetMetrics

> CdnGetMetricsResponse GetMetrics(ctx, stackId, optional)

Get metrics

The last 24 hours of metrics are retrieved if the start and end dates are not provided

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
 **optional** | ***GetMetricsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetMetricsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **optional.Time**| The starting date to retrieve metrics for. | 
 **endDate** | **optional.Time**| The ending date to retrieve metrics for. | 
 **granularity** | **optional.String**| An ISO 8601 duration to roll up metrics by | [default to AUTO]
 **platforms** | **optional.String**| A comma-separated list of billing platforms to filter metrics for. | 
 **pops** | **optional.String**| A comma-separated list of StackPath point of presence location codes to filter metrics for. | 
 **billingRegions** | **optional.String**| A comma-separated list of billing regions to filter metrics for. | 
 **sites** | **optional.String**| A comma-separated list of site IDs to filter metrics for. | 
 **groupBy** | **optional.String**| The field to group metrics by | [default to NONE]
 **siteTypeFilter** | **optional.String**| The type of sites to retrieve metrics for | [default to ALL]

### Return type

[**CdnGetMetricsResponse**](cdnGetMetricsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

