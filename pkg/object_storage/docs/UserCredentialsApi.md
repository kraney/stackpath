# \UserCredentialsApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCredential**](UserCredentialsApi.md#DeleteCredential) | **Delete** /storage/v1/stacks/{stack_id}/users/{user_id}/credentials/{access_key} | Delete credentials
[**GenerateCredentials**](UserCredentialsApi.md#GenerateCredentials) | **Post** /storage/v1/stacks/{stack_id}/users/{user_id}/credentials/generate | Create credentials
[**GetCredentials**](UserCredentialsApi.md#GetCredentials) | **Get** /storage/v1/stacks/{stack_id}/users/{user_id}/credentials | Get credentials



## DeleteCredential

> DeleteCredential(ctx, stackId, userId, accessKey).Execute()

Delete credentials

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
    userId := "userId_example" // string | A user ID
    accessKey := "accessKey_example" // string | A user's access key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserCredentialsApi.DeleteCredential(context.Background(), stackId, userId, accessKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserCredentialsApi.DeleteCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**userId** | **string** | A user ID | 
**accessKey** | **string** | A user&#39;s access key | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCredentialRequest struct via the builder pattern


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


## GenerateCredentials

> StorageGenerateCredentialsResponse GenerateCredentials(ctx, stackId, userId).Execute()

Create credentials



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
    userId := "userId_example" // string | A user ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserCredentialsApi.GenerateCredentials(context.Background(), stackId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserCredentialsApi.GenerateCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateCredentials`: StorageGenerateCredentialsResponse
    fmt.Fprintf(os.Stdout, "Response from `UserCredentialsApi.GenerateCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**userId** | **string** | A user ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**StorageGenerateCredentialsResponse**](storageGenerateCredentialsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCredentials

> StorageGetCredentialsResponse GetCredentials(ctx, stackId, userId).Execute()

Get credentials

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
    userId := "userId_example" // string | A user ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserCredentialsApi.GetCredentials(context.Background(), stackId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserCredentialsApi.GetCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCredentials`: StorageGetCredentialsResponse
    fmt.Fprintf(os.Stdout, "Response from `UserCredentialsApi.GetCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**userId** | **string** | A user ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**StorageGetCredentialsResponse**](storageGetCredentialsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

