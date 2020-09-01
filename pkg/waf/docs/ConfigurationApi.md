# \ConfigurationApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSiteDnsTargets**](ConfigurationApi.md#GetSiteDnsTargets) | **Get** /waf/v1/stacks/{stack_id}/sites/{site_id}/delivery/dns/targets | Get CNAME targets



## GetSiteDnsTargets

> WafGetSiteDnsTargetsResponse GetSiteDnsTargets(ctx, stackId, siteId).Execute()

Get CNAME targets



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
    resp, r, err := api_client.ConfigurationApi.GetSiteDnsTargets(context.Background(), stackId, siteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.GetSiteDnsTargets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSiteDnsTargets`: WafGetSiteDnsTargetsResponse
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationApi.GetSiteDnsTargets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**siteId** | **string** | A site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteDnsTargetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WafGetSiteDnsTargetsResponse**](wafGetSiteDnsTargetsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

