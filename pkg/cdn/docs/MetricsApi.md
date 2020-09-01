# \MetricsApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMetrics**](MetricsApi.md#GetMetrics) | **Get** /cdn/v1/stacks/{stack_id}/metrics | Get metrics



## GetMetrics

> CdnGetMetricsResponse GetMetrics(ctx, stackId).StartDate(startDate).EndDate(endDate).Granularity(granularity).Platforms(platforms).Pops(pops).BillingRegions(billingRegions).Sites(sites).GroupBy(groupBy).SiteTypeFilter(siteTypeFilter).Execute()

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
    granularity := "granularity_example" // string | An ISO 8601 duration to roll up metrics by (optional) (default to "AUTO")
    platforms := "platforms_example" // string | A comma-separated list of billing platforms to filter metrics for. (optional)
    pops := "pops_example" // string | A comma-separated list of StackPath point of presence location codes to filter metrics for. (optional)
    billingRegions := "billingRegions_example" // string | A comma-separated list of billing regions to filter metrics for. (optional)
    sites := "sites_example" // string | A comma-separated list of site IDs to filter metrics for. (optional)
    groupBy := "groupBy_example" // string | The field to group metrics by (optional) (default to "NONE")
    siteTypeFilter := "siteTypeFilter_example" // string | The type of sites to retrieve metrics for (optional) (default to "ALL")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MetricsApi.GetMetrics(context.Background(), stackId).StartDate(startDate).EndDate(endDate).Granularity(granularity).Platforms(platforms).Pops(pops).BillingRegions(billingRegions).Sites(sites).GroupBy(groupBy).SiteTypeFilter(siteTypeFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.GetMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMetrics`: CdnGetMetricsResponse
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
 **granularity** | **string** | An ISO 8601 duration to roll up metrics by | [default to &quot;AUTO&quot;]
 **platforms** | **string** | A comma-separated list of billing platforms to filter metrics for. | 
 **pops** | **string** | A comma-separated list of StackPath point of presence location codes to filter metrics for. | 
 **billingRegions** | **string** | A comma-separated list of billing regions to filter metrics for. | 
 **sites** | **string** | A comma-separated list of site IDs to filter metrics for. | 
 **groupBy** | **string** | The field to group metrics by | [default to &quot;NONE&quot;]
 **siteTypeFilter** | **string** | The type of sites to retrieve metrics for | [default to &quot;ALL&quot;]

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

