# \MetricsApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMetrics**](MetricsApi.md#GetMetrics) | **Get** /delivery/v1/stacks/{stack_id}/metrics | Get metrics



## GetMetrics

> PrometheusMetrics GetMetrics(ctx, stackId).StartDate(startDate).EndDate(endDate).Granularity(granularity).StatusCategories(statusCategories).StatusCodes(statusCodes).Sites(sites).GroupBy(groupBy).BillingRegions(billingRegions).Pops(pops).Platforms(platforms).SiteTypeFilter(siteTypeFilter).MetricType(metricType).Execute()

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
    startDate := Get-Date // time.Time | The starting date to retrieve metrics for. (optional)
    endDate := Get-Date // time.Time | The ending date to retrieve metrics for. (optional)
    granularity := "granularity_example" // string |  (optional) (default to "AUTO")
    statusCategories := []string{"Inner_example"} // []string | A comma-separated list of 1 digit http status codes categories to filter by. (optional)
    statusCodes := []string{"Inner_example"} // []string | A comma-separated list of 3 digit http status codes to filter by. (optional)
    sites := []string{"Inner_example"} // []string | A comma-separated list of site IDs to filter metrics for. (optional)
    groupBy := "groupBy_example" // string |  (optional) (default to "NONE")
    billingRegions := []string{"Inner_example"} // []string | A comma-separated list of billing regions to filter metrics for. (optional)
    pops := []string{"Inner_example"} // []string | A comma-separated list of StackPath point of presence location codes to filter metrics for. (optional)
    platforms := []string{"Inner_example"} // []string | A comma-separated list of billing platforms to filter metrics for. (optional)
    siteTypeFilter := []string{"SiteTypeFilter_example"} // []string | A filter to retrieve metrics by site type. (optional)
    metricType := "metricType_example" // string |  (optional) (default to "TRANSFER")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MetricsApi.GetMetrics(context.Background(), stackId).StartDate(startDate).EndDate(endDate).Granularity(granularity).StatusCategories(statusCategories).StatusCodes(statusCodes).Sites(sites).GroupBy(groupBy).BillingRegions(billingRegions).Pops(pops).Platforms(platforms).SiteTypeFilter(siteTypeFilter).MetricType(metricType).Execute()
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

 **startDate** | **time.Time** | The starting date to retrieve metrics for. | 
 **endDate** | **time.Time** | The ending date to retrieve metrics for. | 
 **granularity** | **string** |  | [default to &quot;AUTO&quot;]
 **statusCategories** | [**[]string**](string.md) | A comma-separated list of 1 digit http status codes categories to filter by. | 
 **statusCodes** | [**[]string**](string.md) | A comma-separated list of 3 digit http status codes to filter by. | 
 **sites** | [**[]string**](string.md) | A comma-separated list of site IDs to filter metrics for. | 
 **groupBy** | **string** |  | [default to &quot;NONE&quot;]
 **billingRegions** | [**[]string**](string.md) | A comma-separated list of billing regions to filter metrics for. | 
 **pops** | [**[]string**](string.md) | A comma-separated list of StackPath point of presence location codes to filter metrics for. | 
 **platforms** | [**[]string**](string.md) | A comma-separated list of billing platforms to filter metrics for. | 
 **siteTypeFilter** | [**[]string**](string.md) | A filter to retrieve metrics by site type. | 
 **metricType** | **string** |  | [default to &quot;TRANSFER&quot;]

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

