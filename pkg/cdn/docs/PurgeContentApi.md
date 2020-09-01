# \PurgeContentApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPurgeStatus**](PurgeContentApi.md#GetPurgeStatus) | **Get** /cdn/v1/stacks/{stack_id}/purge/{purge_id} | Get purge status
[**PurgeContent**](PurgeContentApi.md#PurgeContent) | **Post** /cdn/v1/stacks/{stack_id}/purge | Purge content



## GetPurgeStatus

> CdnGetPurgeStatusResponse GetPurgeStatus(ctx, stackId, purgeId).Execute()

Get purge status

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
    purgeId := "purgeId_example" // string | A CDN purge ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PurgeContentApi.GetPurgeStatus(context.Background(), stackId, purgeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PurgeContentApi.GetPurgeStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPurgeStatus`: CdnGetPurgeStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `PurgeContentApi.GetPurgeStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**purgeId** | **string** | A CDN purge ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPurgeStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CdnGetPurgeStatusResponse**](cdnGetPurgeStatusResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PurgeContent

> CdnPurgeContentResponse PurgeContent(ctx, stackId).CdnPurgeContentRequest(cdnPurgeContentRequest).Execute()

Purge content



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
    cdnPurgeContentRequest := openapiclient.cdnPurgeContentRequest{Items: []PurgeContentRequestItem{openapiclient.PurgeContentRequestItem{Url: "Url_example", Recursive: false, InvalidateOnly: false, PurgeAllDynamic: false, Headers: []string{"Headers_example"), PurgeSelector: []PurgeContentRequestPurgeSelector{openapiclient.PurgeContentRequestPurgeSelector{SelectorType: openapiclient.PurgeContentRequestPurgeSelectorType{}, SelectorName: "SelectorName_example", SelectorValue: "SelectorValue_example", SelectorValueDelimiter: "SelectorValueDelimiter_example"})})} // CdnPurgeContentRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PurgeContentApi.PurgeContent(context.Background(), stackId, cdnPurgeContentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PurgeContentApi.PurgeContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PurgeContent`: CdnPurgeContentResponse
    fmt.Fprintf(os.Stdout, "Response from `PurgeContentApi.PurgeContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiPurgeContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cdnPurgeContentRequest** | [**CdnPurgeContentRequest**](CdnPurgeContentRequest.md) |  | 

### Return type

[**CdnPurgeContentResponse**](cdnPurgeContentResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

