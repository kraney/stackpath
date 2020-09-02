# \MetricsApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMetrics**](MetricsApi.md#GetMetrics) | **Get** /monitoring/v2/stacks/{stack_id}/monitors/{monitor_id}/metrics | Get metrics



## GetMetrics

> PrometheusMetrics GetMetrics(ctx, stackId, monitorId, optional)

Get metrics

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**monitorId** | **string**| A monitor ID | 
 **optional** | ***GetMetricsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetMetricsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startDate** | **optional.Time**| The start date of the time range to look for metrics. | 
 **endDate** | **optional.Time**| The end date of the time range to look for metrics. | 
 **pops** | [**optional.Interface of []string**](string.md)| A list of point of presence to retrieve metrics for. | 
 **metrics** | [**optional.Interface of []string**](string.md)| A list of metrics to retrieve. | 
 **aggregation** | **optional.String**|  | [default to DEFAULT]
 **groupBy** | **optional.String**|  | [default to NONE]
 **granularity** | **optional.String**|  | [default to DEFAULT_GRANULARITY]
 **granularityFunction** | **optional.String**|  | [default to DEFAULT]

### Return type

[**PrometheusMetrics**](prometheusMetrics.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

