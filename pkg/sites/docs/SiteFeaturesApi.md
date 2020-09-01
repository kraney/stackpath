# \SiteFeaturesApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisableSiteCDN**](SiteFeaturesApi.md#DisableSiteCDN) | **Delete** /delivery/v1/stacks/{stack_id}/sites/{site_id}/cdn | Disable CDN
[**DisableSiteEdgeEngine2**](SiteFeaturesApi.md#DisableSiteEdgeEngine2) | **Delete** /delivery/v1/stacks/{stack_id}/sites/{site_id}/scripting | Disable serverless scripting
[**DisableSiteWAF**](SiteFeaturesApi.md#DisableSiteWAF) | **Delete** /delivery/v1/stacks/{stack_id}/sites/{site_id}/waf | Disable WAF
[**EnableSiteCDN**](SiteFeaturesApi.md#EnableSiteCDN) | **Post** /delivery/v1/stacks/{stack_id}/sites/{site_id}/cdn | Enable CDN
[**EnableSiteEdgeEngine2**](SiteFeaturesApi.md#EnableSiteEdgeEngine2) | **Post** /delivery/v1/stacks/{stack_id}/sites/{site_id}/scripting | Enable serverless scripting
[**EnableSiteWAF**](SiteFeaturesApi.md#EnableSiteWAF) | **Post** /delivery/v1/stacks/{stack_id}/sites/{site_id}/waf | Enable WAF



## DisableSiteCDN

> DisableSiteCDN(ctx, stackId, siteId).Execute()

Disable CDN

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
    resp, r, err := api_client.SiteFeaturesApi.DisableSiteCDN(context.Background(), stackId, siteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SiteFeaturesApi.DisableSiteCDN``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDisableSiteCDNRequest struct via the builder pattern


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


## DisableSiteEdgeEngine2

> DisableSiteEdgeEngine2(ctx, stackId, siteId).Execute()

Disable serverless scripting

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
    resp, r, err := api_client.SiteFeaturesApi.DisableSiteEdgeEngine2(context.Background(), stackId, siteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SiteFeaturesApi.DisableSiteEdgeEngine2``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDisableSiteEdgeEngine2Request struct via the builder pattern


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


## DisableSiteWAF

> DisableSiteWAF(ctx, stackId, siteId).Execute()

Disable WAF

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
    resp, r, err := api_client.SiteFeaturesApi.DisableSiteWAF(context.Background(), stackId, siteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SiteFeaturesApi.DisableSiteWAF``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDisableSiteWAFRequest struct via the builder pattern


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


## EnableSiteCDN

> EnableSiteCDN(ctx, stackId, siteId).Execute()

Enable CDN

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
    resp, r, err := api_client.SiteFeaturesApi.EnableSiteCDN(context.Background(), stackId, siteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SiteFeaturesApi.EnableSiteCDN``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnableSiteCDNRequest struct via the builder pattern


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


## EnableSiteEdgeEngine2

> EnableSiteEdgeEngine2(ctx, stackId, siteId).Execute()

Enable serverless scripting

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
    resp, r, err := api_client.SiteFeaturesApi.EnableSiteEdgeEngine2(context.Background(), stackId, siteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SiteFeaturesApi.EnableSiteEdgeEngine2``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnableSiteEdgeEngine2Request struct via the builder pattern


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


## EnableSiteWAF

> EnableSiteWAF(ctx, stackId, siteId).DeliveryEnableSiteWAFRequest(deliveryEnableSiteWAFRequest).Execute()

Enable WAF

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
    deliveryEnableSiteWAFRequest := openapiclient.deliveryEnableSiteWAFRequest{ApiUrls: []string{"ApiUrls_example")} // DeliveryEnableSiteWAFRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SiteFeaturesApi.EnableSiteWAF(context.Background(), stackId, siteId, deliveryEnableSiteWAFRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SiteFeaturesApi.EnableSiteWAF``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnableSiteWAFRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deliveryEnableSiteWAFRequest** | [**DeliveryEnableSiteWAFRequest**](DeliveryEnableSiteWAFRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

