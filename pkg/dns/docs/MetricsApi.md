# \MetricsApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMetrics**](MetricsApi.md#GetMetrics) | **Get** /dns/v1/stacks/{stack_id}/metrics | Get metrics



## GetMetrics

> PrometheusMetrics GetMetrics(ctx, stackId, optional)

Get metrics

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

 **zoneId** | **optional.String**| An optional zone ID to retrieve metrics for | 
 **pop** | **optional.String**| An optional StackPath POP code to retrieve metrics from | 
 **startDate** | **optional.Time**| A lower bound date to search metrics for. | 
 **endDate** | **optional.Time**| An upper bound date to search metrics for. | 
 **type_** | **optional.String**|  | [default to ZONE_QUERY_COUNT]
 **granularity** | **optional.String**|  | [default to DEFAULT]
 **dimensions** | [**optional.Interface of []string**](string.md)| The dimensions to aggregate metrics by. When empty or unset, defaults to all dimensions. If non-empty, this must always at least include the stack dimension. | 

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

