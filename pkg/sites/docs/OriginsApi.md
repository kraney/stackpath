# \OriginsApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConnectScopeToOrigin**](OriginsApi.md#ConnectScopeToOrigin) | **Post** /delivery/v1/stacks/{stack_id}/sites/{site_id}/scopes/{scope_id}/origins | Connect an origin to a scope
[**GetOrigin**](OriginsApi.md#GetOrigin) | **Get** /delivery/v1/stacks/{stack_id}/origins/{origin_id} | Get an origin
[**GetOrigins**](OriginsApi.md#GetOrigins) | **Get** /delivery/v1/stacks/{stack_id}/origins | Get all origins
[**GetScopeOrigins**](OriginsApi.md#GetScopeOrigins) | **Get** /delivery/v1/stacks/{stack_id}/sites/{site_id}/scopes/{scope_id}/origins | Get a scope&#39;s origins
[**UpdateOrigin**](OriginsApi.md#UpdateOrigin) | **Patch** /delivery/v1/stacks/{stack_id}/origins/{origin_id} | Update an origin



## ConnectScopeToOrigin

> DeliveryConnectScopeToOriginResponse ConnectScopeToOrigin(ctx, stackId, siteId, scopeId).DeliveryConnectScopeToOriginRequest(deliveryConnectScopeToOriginRequest).Execute()

Connect an origin to a scope



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
    scopeId := "scopeId_example" // string | A scope ID
    deliveryConnectScopeToOriginRequest := openapiclient.deliveryConnectScopeToOriginRequest{Origin: openapiclient.deliveryConnectScopeToOriginRequestOrigin{Http: openapiclient.deliveryHTTPOrigin{Path: "Path_example", Hostname: "Hostname_example", Port: 123, SecurePort: 123, Authentication: openapiclient.deliveryOriginAuthentication{Basic: openapiclient.deliveryBasicAuthentication{Username: "Username_example", Password: "Password_example"}, S3: openapiclient.deliveryS3Authentication{AccessKey: "AccessKey_example", SecretKey: "SecretKey_example"}}, VerifyCertificate: false, CertificateCommonName: "CertificateCommonName_example"}, StackpathStorage: openapiclient.deliveryStackPathStorageOrigin{BucketName: "BucketName_example", BucketRegion: "BucketRegion_example", Authentication: openapiclient.deliveryOriginAuthentication{Basic: openapiclient.deliveryBasicAuthentication{Username: "Username_example", Password: "Password_example"}, S3: openapiclient.deliveryS3Authentication{AccessKey: "AccessKey_example", SecretKey: "SecretKey_example"}}}, S3Storage: openapiclient.deliveryAWSS3Origin{BucketName: "BucketName_example", BucketRegion: "BucketRegion_example", Authentication: }, GoogleCloudStorage: openapiclient.deliveryGoogleStorageOrigin{BucketName: "BucketName_example"}}, Priority: 123, OriginId: "OriginId_example"} // DeliveryConnectScopeToOriginRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OriginsApi.ConnectScopeToOrigin(context.Background(), stackId, siteId, scopeId, deliveryConnectScopeToOriginRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OriginsApi.ConnectScopeToOrigin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectScopeToOrigin`: DeliveryConnectScopeToOriginResponse
    fmt.Fprintf(os.Stdout, "Response from `OriginsApi.ConnectScopeToOrigin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**siteId** | **string** | A site ID | 
**scopeId** | **string** | A scope ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectScopeToOriginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **deliveryConnectScopeToOriginRequest** | [**DeliveryConnectScopeToOriginRequest**](DeliveryConnectScopeToOriginRequest.md) |  | 

### Return type

[**DeliveryConnectScopeToOriginResponse**](deliveryConnectScopeToOriginResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrigin

> DeliveryGetOriginResponse GetOrigin(ctx, stackId, originId).Execute()

Get an origin

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
    originId := "originId_example" // string | An origin ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OriginsApi.GetOrigin(context.Background(), stackId, originId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OriginsApi.GetOrigin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrigin`: DeliveryGetOriginResponse
    fmt.Fprintf(os.Stdout, "Response from `OriginsApi.GetOrigin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**originId** | **string** | An origin ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOriginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeliveryGetOriginResponse**](deliveryGetOriginResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrigins

> DeliveryGetOriginsResponse GetOrigins(ctx, stackId).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).Execute()

Get all origins



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
    pageRequestFirst := "pageRequestFirst_example" // string | The number of items desired. (optional)
    pageRequestAfter := "pageRequestAfter_example" // string | The cursor value after which data will be returned. (optional)
    pageRequestFilter := "pageRequestFilter_example" // string | SQL-style constraint filters. (optional)
    pageRequestSortBy := "pageRequestSortBy_example" // string | Sort the response by the given field. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OriginsApi.GetOrigins(context.Background(), stackId).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OriginsApi.GetOrigins``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrigins`: DeliveryGetOriginsResponse
    fmt.Fprintf(os.Stdout, "Response from `OriginsApi.GetOrigins`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOriginsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageRequestFirst** | **string** | The number of items desired. | 
 **pageRequestAfter** | **string** | The cursor value after which data will be returned. | 
 **pageRequestFilter** | **string** | SQL-style constraint filters. | 
 **pageRequestSortBy** | **string** | Sort the response by the given field. | 

### Return type

[**DeliveryGetOriginsResponse**](deliveryGetOriginsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetScopeOrigins

> DeliveryGetScopeOriginsResponse GetScopeOrigins(ctx, stackId, siteId, scopeId).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).Execute()

Get a scope's origins

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
    scopeId := "scopeId_example" // string | A scope ID
    pageRequestFirst := "pageRequestFirst_example" // string | The number of items desired. (optional)
    pageRequestAfter := "pageRequestAfter_example" // string | The cursor value after which data will be returned. (optional)
    pageRequestFilter := "pageRequestFilter_example" // string | SQL-style constraint filters. (optional)
    pageRequestSortBy := "pageRequestSortBy_example" // string | Sort the response by the given field. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OriginsApi.GetScopeOrigins(context.Background(), stackId, siteId, scopeId).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OriginsApi.GetScopeOrigins``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetScopeOrigins`: DeliveryGetScopeOriginsResponse
    fmt.Fprintf(os.Stdout, "Response from `OriginsApi.GetScopeOrigins`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**siteId** | **string** | A site ID | 
**scopeId** | **string** | A scope ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetScopeOriginsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pageRequestFirst** | **string** | The number of items desired. | 
 **pageRequestAfter** | **string** | The cursor value after which data will be returned. | 
 **pageRequestFilter** | **string** | SQL-style constraint filters. | 
 **pageRequestSortBy** | **string** | Sort the response by the given field. | 

### Return type

[**DeliveryGetScopeOriginsResponse**](deliveryGetScopeOriginsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrigin

> DeliveryUpdateOriginResponse UpdateOrigin(ctx, stackId, originId).DeliveryUpdateOriginRequest(deliveryUpdateOriginRequest).Execute()

Update an origin

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
    originId := "originId_example" // string | An origin ID
    deliveryUpdateOriginRequest := openapiclient.deliveryUpdateOriginRequest{Http: openapiclient.deliveryHTTPOrigin{Path: "Path_example", Hostname: "Hostname_example", Port: 123, SecurePort: 123, Authentication: , VerifyCertificate: false, CertificateCommonName: "CertificateCommonName_example"}, StackpathStorage: openapiclient.deliveryStackPathStorageOrigin{BucketName: "BucketName_example", BucketRegion: "BucketRegion_example", Authentication: }, S3Storage: openapiclient.deliveryAWSS3Origin{BucketName: "BucketName_example", BucketRegion: "BucketRegion_example", Authentication: }, GoogleCloudStorage: openapiclient.deliveryGoogleStorageOrigin{BucketName: "BucketName_example"}} // DeliveryUpdateOriginRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OriginsApi.UpdateOrigin(context.Background(), stackId, originId, deliveryUpdateOriginRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OriginsApi.UpdateOrigin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrigin`: DeliveryUpdateOriginResponse
    fmt.Fprintf(os.Stdout, "Response from `OriginsApi.UpdateOrigin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**originId** | **string** | An origin ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOriginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deliveryUpdateOriginRequest** | [**DeliveryUpdateOriginRequest**](DeliveryUpdateOriginRequest.md) |  | 

### Return type

[**DeliveryUpdateOriginResponse**](deliveryUpdateOriginResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

