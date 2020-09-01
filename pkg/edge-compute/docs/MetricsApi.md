# \MetricsApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMetrics**](MetricsApi.md#GetMetrics) | **Get** /workload/v1/stacks/{stack_id}/metrics | Get metrics



## GetMetrics

> PrometheusMetrics GetMetrics(ctx, stackId).WorkloadId(workloadId).StartDate(startDate).EndDate(endDate).Type_(type_).Granularity(granularity).InstanceName(instanceName).Pop(pop).Region(region).GroupBy(groupBy).Grouping(grouping).Execute()

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
    workloadId := "workloadId_example" // string | An EdgeCompute workload ID (optional)
    startDate := Get-Date // time.Time | A lower bound date to search metrics for. (optional)
    endDate := Get-Date // time.Time | An upper bound date to search metrics for. (optional)
    type_ := "type__example" // string |  (optional) (default to "BANDWIDTH")
    granularity := "granularity_example" // string |  (optional) (default to "DEFAULT")
    instanceName := "instanceName_example" // string | An instance name within a workload to filter metrics for (optional)
    pop := "pop_example" // string | A StackPath POP to filter traffic metrics for. This field does not apply when retrieving INSTANCE type metrics (optional)
    region := "region_example" // string |  (optional) (default to "ALL")
    groupBy := "groupBy_example" // string |  (optional) (default to "NONE")
    grouping := []string{"Inner_example"} // []string | List of fields to group by. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MetricsApi.GetMetrics(context.Background(), stackId).WorkloadId(workloadId).StartDate(startDate).EndDate(endDate).Type_(type_).Granularity(granularity).InstanceName(instanceName).Pop(pop).Region(region).GroupBy(groupBy).Grouping(grouping).Execute()
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

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workloadId** | **string** | An EdgeCompute workload ID | 
 **startDate** | **time.Time** | A lower bound date to search metrics for. | 
 **endDate** | **time.Time** | An upper bound date to search metrics for. | 
 **type_** | **string** |  | [default to &quot;BANDWIDTH&quot;]
 **granularity** | **string** |  | [default to &quot;DEFAULT&quot;]
 **instanceName** | **string** | An instance name within a workload to filter metrics for | 
 **pop** | **string** | A StackPath POP to filter traffic metrics for. This field does not apply when retrieving INSTANCE type metrics | 
 **region** | **string** |  | [default to &quot;ALL&quot;]
 **groupBy** | **string** |  | [default to &quot;NONE&quot;]
 **grouping** | [**[]string**](string.md) | List of fields to group by. | 

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

