# \StacksApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStack**](StacksApi.md#CreateStack) | **Post** /stack/v1/stacks | Create a stack
[**GetStack**](StacksApi.md#GetStack) | **Get** /stack/v1/stacks/{stack_id} | Get a stack
[**GetStacks**](StacksApi.md#GetStacks) | **Get** /stack/v1/stacks | Get all stacks
[**UpdateStack2**](StacksApi.md#UpdateStack2) | **Patch** /stack/v1/stacks/{stack_id} | Update a stack



## CreateStack

> StackCreateStackResponse CreateStack(ctx).StackCreateStackRequest(stackCreateStackRequest).Execute()

Create a stack

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
    stackCreateStackRequest := openapiclient.stackCreateStackRequest{AccountId: "AccountId_example", Slug: "Slug_example", Name: "Name_example"} // StackCreateStackRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StacksApi.CreateStack(context.Background(), stackCreateStackRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StacksApi.CreateStack``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStack`: StackCreateStackResponse
    fmt.Fprintf(os.Stdout, "Response from `StacksApi.CreateStack`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateStackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stackCreateStackRequest** | [**StackCreateStackRequest**](StackCreateStackRequest.md) |  | 

### Return type

[**StackCreateStackResponse**](stackCreateStackResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStack

> StackStack GetStack(ctx, stackId).Execute()

Get a stack

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StacksApi.GetStack(context.Background(), stackId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StacksApi.GetStack``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStack`: StackStack
    fmt.Fprintf(os.Stdout, "Response from `StacksApi.GetStack`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StackStack**](stackStack.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStacks

> StackGetStacksResponse GetStacks(ctx).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).AccountId(accountId).Execute()

Get all stacks

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
    accountId := "accountId_example" // string | An account ID (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StacksApi.GetStacks(context.Background(), ).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).AccountId(accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StacksApi.GetStacks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStacks`: StackGetStacksResponse
    fmt.Fprintf(os.Stdout, "Response from `StacksApi.GetStacks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStacksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageRequestFirst** | **string** | The number of items desired. | 
 **pageRequestAfter** | **string** | The cursor value after which data will be returned. | 
 **pageRequestFilter** | **string** | SQL-style constraint filters. | 
 **pageRequestSortBy** | **string** | Sort the response by the given field. | 
 **accountId** | **string** | An account ID | 

### Return type

[**StackGetStacksResponse**](stackGetStacksResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStack2

> StackUpdateStackResponse UpdateStack2(ctx, stackId).StackUpdateStackRequest2(stackUpdateStackRequest2).Execute()

Update a stack

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
    stackUpdateStackRequest2 := openapiclient.stackUpdateStackRequest2{Name: "Name_example"} // StackUpdateStackRequest2 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StacksApi.UpdateStack2(context.Background(), stackId, stackUpdateStackRequest2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StacksApi.UpdateStack2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStack2`: StackUpdateStackResponse
    fmt.Fprintf(os.Stdout, "Response from `StacksApi.UpdateStack2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStack2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **stackUpdateStackRequest2** | [**StackUpdateStackRequest2**](StackUpdateStackRequest2.md) |  | 

### Return type

[**StackUpdateStackResponse**](stackUpdateStackResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

