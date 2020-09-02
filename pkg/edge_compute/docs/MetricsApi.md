# \MetricsApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMetrics**](MetricsApi.md#GetMetrics) | **Get** /workload/v1/stacks/{stack_id}/metrics | Get metrics



## GetMetrics

> PrometheusMetrics GetMetrics(ctx, stackId, optional)

Get metrics

Retrieve usage metrics for all workloads in a stack, a specific workload, or a specific instance in a workload

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

 **workloadId** | **optional.String**| An EdgeCompute workload ID | 
 **startDate** | **optional.Time**| A lower bound date to search metrics for. | 
 **endDate** | **optional.Time**| An upper bound date to search metrics for. | 
 **type_** | **optional.String**|  | [default to BANDWIDTH]
 **granularity** | **optional.String**|  | [default to DEFAULT]
 **instanceName** | **optional.String**| An instance name within a workload to filter metrics for | 
 **pop** | **optional.String**| A StackPath POP to filter traffic metrics for. This field does not apply when retrieving INSTANCE type metrics | 
 **region** | **optional.String**|  | [default to ALL]
 **groupBy** | **optional.String**|  | [default to NONE]
 **grouping** | [**optional.Interface of []string**](string.md)| List of fields to group by. | 

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

