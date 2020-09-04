# \WAFFeaturesApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDdosSettings**](WAFFeaturesApi.md#GetDdosSettings) | **Get** /waf/v1/stacks/{stack_id}/sites/{site_id}/ddos | Get DDOS settings
[**GetTags**](WAFFeaturesApi.md#GetTags) | **Get** /waf/v1/tags | Get all tags
[**MonitorSite**](WAFFeaturesApi.md#MonitorSite) | **Post** /waf/v1/stacks/{stack_id}/sites/{site_id}/monitoring | Enable monitoring mode
[**UnMonitorSite**](WAFFeaturesApi.md#UnMonitorSite) | **Delete** /waf/v1/stacks/{stack_id}/sites/{site_id}/monitoring | Disable monitoring mode
[**UpdateDdosSettings**](WAFFeaturesApi.md#UpdateDdosSettings) | **Patch** /waf/v1/stacks/{stack_id}/sites/{site_id}/ddos | Update DDOS settings
[**UpdateSiteApiUrls**](WAFFeaturesApi.md#UpdateSiteApiUrls) | **Put** /waf/v1/stacks/{stack_id}/sites/{site_id}/api_urls | Update API URLs



## GetDdosSettings

> WafGetDdosSettingsResponse GetDdosSettings(ctx, stackId, siteId).Execute()

Get DDOS settings

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
    siteId := "siteId_example" // string | A site ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WAFFeaturesApi.GetDdosSettings(context.Background(), stackId, siteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WAFFeaturesApi.GetDdosSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDdosSettings`: WafGetDdosSettingsResponse
    fmt.Fprintf(os.Stdout, "Response from `WAFFeaturesApi.GetDdosSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**siteId** | **string** | A site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDdosSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WafGetDdosSettingsResponse**](wafGetDdosSettingsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTags

> WafGetTagsResponse GetTags(ctx).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).Execute()

Get all tags



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
    pageRequestFirst := "pageRequestFirst_example" // string | The number of items desired. (optional)
    pageRequestAfter := "pageRequestAfter_example" // string | The cursor value after which data will be returned. (optional)
    pageRequestFilter := "pageRequestFilter_example" // string | SQL-style constraint filters. (optional)
    pageRequestSortBy := "pageRequestSortBy_example" // string | Sort the response by the given field. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WAFFeaturesApi.GetTags(context.Background(), ).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WAFFeaturesApi.GetTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTags`: WafGetTagsResponse
    fmt.Fprintf(os.Stdout, "Response from `WAFFeaturesApi.GetTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageRequestFirst** | **string** | The number of items desired. | 
 **pageRequestAfter** | **string** | The cursor value after which data will be returned. | 
 **pageRequestFilter** | **string** | SQL-style constraint filters. | 
 **pageRequestSortBy** | **string** | Sort the response by the given field. | 

### Return type

[**WafGetTagsResponse**](wafGetTagsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorSite

> MonitorSite(ctx, stackId, siteId).Execute()

Enable monitoring mode



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
    siteId := "siteId_example" // string | A site ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WAFFeaturesApi.MonitorSite(context.Background(), stackId, siteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WAFFeaturesApi.MonitorSite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**siteId** | **string** | A site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnMonitorSite

> UnMonitorSite(ctx, stackId, siteId).Execute()

Disable monitoring mode



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
    siteId := "siteId_example" // string | A site ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WAFFeaturesApi.UnMonitorSite(context.Background(), stackId, siteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WAFFeaturesApi.UnMonitorSite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**siteId** | **string** | A site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnMonitorSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDdosSettings

> WafUpdateDdosSettingsResponse UpdateDdosSettings(ctx, stackId, siteId).WafUpdateDdosSettingsRequest(wafUpdateDdosSettingsRequest).Execute()

Update DDOS settings

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
    siteId := "siteId_example" // string | A site ID
    wafUpdateDdosSettingsRequest := openapiclient.wafUpdateDdosSettingsRequest{GlobalThreshold: "GlobalThreshold_example", BurstThreshold: "BurstThreshold_example", SubSecondBurstThreshold: "SubSecondBurstThreshold_example"} // WafUpdateDdosSettingsRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WAFFeaturesApi.UpdateDdosSettings(context.Background(), stackId, siteId, wafUpdateDdosSettingsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WAFFeaturesApi.UpdateDdosSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDdosSettings`: WafUpdateDdosSettingsResponse
    fmt.Fprintf(os.Stdout, "Response from `WAFFeaturesApi.UpdateDdosSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**siteId** | **string** | A site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDdosSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **wafUpdateDdosSettingsRequest** | [**WafUpdateDdosSettingsRequest**](WafUpdateDdosSettingsRequest.md) |  | 

### Return type

[**WafUpdateDdosSettingsResponse**](wafUpdateDdosSettingsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSiteApiUrls

> WafUpdateSiteApiUrlsResponse UpdateSiteApiUrls(ctx, stackId, siteId).WafUpdateSiteApiUrlsRequest(wafUpdateSiteApiUrlsRequest).Execute()

Update API URLs

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
    siteId := "siteId_example" // string | A site ID
    wafUpdateSiteApiUrlsRequest := openapiclient.wafUpdateSiteApiUrlsRequest{ApiUrls: []string{"ApiUrls_example")} // WafUpdateSiteApiUrlsRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WAFFeaturesApi.UpdateSiteApiUrls(context.Background(), stackId, siteId, wafUpdateSiteApiUrlsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WAFFeaturesApi.UpdateSiteApiUrls``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSiteApiUrls`: WafUpdateSiteApiUrlsResponse
    fmt.Fprintf(os.Stdout, "Response from `WAFFeaturesApi.UpdateSiteApiUrls`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**siteId** | **string** | A site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSiteApiUrlsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **wafUpdateSiteApiUrlsRequest** | [**WafUpdateSiteApiUrlsRequest**](WafUpdateSiteApiUrlsRequest.md) |  | 

### Return type

[**WafUpdateSiteApiUrlsResponse**](wafUpdateSiteApiUrlsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

