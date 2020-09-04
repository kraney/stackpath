# \MetricsApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMetrics**](MetricsApi.md#GetMetrics) | **Get** /monitoring/v2/stacks/{stack_id}/monitors/{monitor_id}/metrics | Get metrics



## GetMetrics

> PrometheusMetrics GetMetrics(ctx, stackId, monitorId).StartDate(startDate).EndDate(endDate).Pops(pops).Metrics(metrics).Aggregation(aggregation).GroupBy(groupBy).Granularity(granularity).GranularityFunction(granularityFunction).Execute()

Get metrics

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    stackId := "stackId_example" // string | A stack ID or slug
    monitorId := "monitorId_example" // string | A monitor ID
    startDate := Get-Date // time.Time | The start date of the time range to look for metrics. (optional)
    endDate := Get-Date // time.Time | The end date of the time range to look for metrics. (optional)
    pops := []string{"Inner_example"} // []string | A list of point of presence to retrieve metrics for. (optional)
    metrics := []string{"Inner_example"} // []string | A list of metrics to retrieve. (optional)
    aggregation := "aggregation_example" // string |  (optional) (default to "DEFAULT")
    groupBy := "groupBy_example" // string |  (optional) (default to "NONE")
    granularity := "granularity_example" // string |  (optional) (default to "DEFAULT_GRANULARITY")
    granularityFunction := "granularityFunction_example" // string |  (optional) (default to "DEFAULT")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MetricsApi.GetMetrics(context.Background(), stackId, monitorId).StartDate(startDate).EndDate(endDate).Pops(pops).Metrics(metrics).Aggregation(aggregation).GroupBy(groupBy).Granularity(granularity).GranularityFunction(granularityFunction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.GetMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMetrics`: PrometheusMetrics
    fmt.Fprintf(os.Stdout, "Response from `MetricsApi.GetMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**monitorId** | **string** | A monitor ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startDate** | **time.Time** | The start date of the time range to look for metrics. | 
 **endDate** | **time.Time** | The end date of the time range to look for metrics. | 
 **pops** | [**[]string**](string.md) | A list of point of presence to retrieve metrics for. | 
 **metrics** | [**[]string**](string.md) | A list of metrics to retrieve. | 
 **aggregation** | **string** |  | [default to &quot;DEFAULT&quot;]
 **groupBy** | **string** |  | [default to &quot;NONE&quot;]
 **granularity** | **string** |  | [default to &quot;DEFAULT_GRANULARITY&quot;]
 **granularityFunction** | **string** |  | [default to &quot;DEFAULT&quot;]

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

