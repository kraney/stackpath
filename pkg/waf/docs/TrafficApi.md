# \TrafficApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTraffic**](TrafficApi.md#GetTraffic) | **Get** /waf/v1/stacks/{stack_id}/traffic | LEGACY: Get traffic
[**GetTrafficV2**](TrafficApi.md#GetTrafficV2) | **Get** /waf/v2/stacks/{stack_id}/traffic | Get traffic



## GetTraffic

> WafGetTrafficResponse GetTraffic(ctx, stackId).SiteId(siteId).StartDate(startDate).EndDate(endDate).Resolution(resolution).Execute()

LEGACY: Get traffic



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
    siteId := "siteId_example" // string | A site ID (optional)
    startDate := Get-Date // time.Time | A lower bound date to search traffic for. (optional)
    endDate := Get-Date // time.Time | An upper bound date to search traffic for. (optional)
    resolution := "resolution_example" // string |  - HOURLY: All data points represent one hour of traffic  - MINUTELY: All data points represent one minute of traffic (optional) (default to "HOURLY")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TrafficApi.GetTraffic(context.Background(), stackId).SiteId(siteId).StartDate(startDate).EndDate(endDate).Resolution(resolution).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrafficApi.GetTraffic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTraffic`: WafGetTrafficResponse
    fmt.Fprintf(os.Stdout, "Response from `TrafficApi.GetTraffic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrafficRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **siteId** | **string** | A site ID | 
 **startDate** | **time.Time** | A lower bound date to search traffic for. | 
 **endDate** | **time.Time** | An upper bound date to search traffic for. | 
 **resolution** | **string** |  - HOURLY: All data points represent one hour of traffic  - MINUTELY: All data points represent one minute of traffic | [default to &quot;HOURLY&quot;]

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

> WafGetTrafficV2Response GetTrafficV2(ctx, stackId).SiteId(siteId).StartDate(startDate).EndDate(endDate).Resolution(resolution).Execute()

Get traffic



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
    siteId := "siteId_example" // string | A site ID (optional)
    startDate := Get-Date // time.Time | A lower bound date to search traffic for. (optional)
    endDate := Get-Date // time.Time | An upper bound date to search traffic for. (optional)
    resolution := "resolution_example" // string |  - DAILY: All data points represent one day of traffic  - HOURLY: All data points represent one hour of traffic  - MINUTELY: All data points represent one minute of traffic (optional) (default to "DAILY")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TrafficApi.GetTrafficV2(context.Background(), stackId).SiteId(siteId).StartDate(startDate).EndDate(endDate).Resolution(resolution).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrafficApi.GetTrafficV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrafficV2`: WafGetTrafficV2Response
    fmt.Fprintf(os.Stdout, "Response from `TrafficApi.GetTrafficV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrafficV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **siteId** | **string** | A site ID | 
 **startDate** | **time.Time** | A lower bound date to search traffic for. | 
 **endDate** | **time.Time** | An upper bound date to search traffic for. | 
 **resolution** | **string** |  - DAILY: All data points represent one day of traffic  - HOURLY: All data points represent one hour of traffic  - MINUTELY: All data points represent one minute of traffic | [default to &quot;DAILY&quot;]

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

