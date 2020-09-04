# \AuthenticationApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUserCredential**](AuthenticationApi.md#CreateUserCredential) | **Post** /identity/v1/users/{user_id}/credentials | Create API credentials
[**DeleteUserCredential**](AuthenticationApi.md#DeleteUserCredential) | **Delete** /identity/v1/users/{user_id}/credentials/{credential_id} | Delete API credentials
[**GetAccessToken**](AuthenticationApi.md#GetAccessToken) | **Post** /identity/v1/oauth2/token | Generate a bearer token
[**GetUserCredentials**](AuthenticationApi.md#GetUserCredentials) | **Get** /identity/v1/users/{user_id}/credentials | Get all API credentials
[**RotateUserCredentialSecret**](AuthenticationApi.md#RotateUserCredentialSecret) | **Post** /identity/v1/users/{user_id}/credentials/{credential_id}/rotate | Create a new API secret



## CreateUserCredential

> IdentityCreateUserCredentialResponse CreateUserCredential(ctx, userId).IdentityCreateUserCredentialRequest(identityCreateUserCredentialRequest).Execute()

Create API credentials



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
    userId := "userId_example" // string | A user ID
    identityCreateUserCredentialRequest := openapiclient.identityCreateUserCredentialRequest{Name: "Name_example"} // IdentityCreateUserCredentialRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationApi.CreateUserCredential(context.Background(), userId, identityCreateUserCredentialRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationApi.CreateUserCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUserCredential`: IdentityCreateUserCredentialResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationApi.CreateUserCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | A user ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **identityCreateUserCredentialRequest** | [**IdentityCreateUserCredentialRequest**](IdentityCreateUserCredentialRequest.md) |  | 

### Return type

[**IdentityCreateUserCredentialResponse**](identityCreateUserCredentialResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserCredential

> DeleteUserCredential(ctx, userId, credentialId).Execute()

Delete API credentials

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
    userId := "userId_example" // string | A user ID
    credentialId := "credentialId_example" // string | The ID of the API credentials to remove

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationApi.DeleteUserCredential(context.Background(), userId, credentialId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationApi.DeleteUserCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | A user ID | 
**credentialId** | **string** | The ID of the API credentials to remove | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserCredentialRequest struct via the builder pattern


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


## GetAccessToken

> IdentityGetAccessTokenResponse GetAccessToken(ctx).IdentityGetAccessTokenRequest(identityGetAccessTokenRequest).Execute()

Generate a bearer token



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
    identityGetAccessTokenRequest := openapiclient.identityGetAccessTokenRequest{GrantType: "GrantType_example", ClientId: "ClientId_example", ClientSecret: "ClientSecret_example"} // IdentityGetAccessTokenRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationApi.GetAccessToken(context.Background(), identityGetAccessTokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationApi.GetAccessToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccessToken`: IdentityGetAccessTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationApi.GetAccessToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identityGetAccessTokenRequest** | [**IdentityGetAccessTokenRequest**](IdentityGetAccessTokenRequest.md) |  | 

### Return type

[**IdentityGetAccessTokenResponse**](identityGetAccessTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserCredentials

> IdentityGetUserCredentialsResponse GetUserCredentials(ctx, userId).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).Execute()

Get all API credentials

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
    userId := "userId_example" // string | A user ID
    pageRequestFirst := "pageRequestFirst_example" // string | The number of items desired. (optional)
    pageRequestAfter := "pageRequestAfter_example" // string | The cursor value after which data will be returned. (optional)
    pageRequestFilter := "pageRequestFilter_example" // string | SQL-style constraint filters. (optional)
    pageRequestSortBy := "pageRequestSortBy_example" // string | Sort the response by the given field. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationApi.GetUserCredentials(context.Background(), userId).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationApi.GetUserCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserCredentials`: IdentityGetUserCredentialsResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationApi.GetUserCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | A user ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageRequestFirst** | **string** | The number of items desired. | 
 **pageRequestAfter** | **string** | The cursor value after which data will be returned. | 
 **pageRequestFilter** | **string** | SQL-style constraint filters. | 
 **pageRequestSortBy** | **string** | Sort the response by the given field. | 

### Return type

[**IdentityGetUserCredentialsResponse**](identityGetUserCredentialsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RotateUserCredentialSecret

> IdentityRotateUserCredentialSecretResponse RotateUserCredentialSecret(ctx, userId, credentialId).Execute()

Create a new API secret



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
    userId := "userId_example" // string | A user ID
    credentialId := "credentialId_example" // string | The ID of the API client credential to rotate secrets for

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationApi.RotateUserCredentialSecret(context.Background(), userId, credentialId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationApi.RotateUserCredentialSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RotateUserCredentialSecret`: IdentityRotateUserCredentialSecretResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationApi.RotateUserCredentialSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | A user ID | 
**credentialId** | **string** | The ID of the API client credential to rotate secrets for | 

### Other Parameters

Other parameters are passed through a pointer to a apiRotateUserCredentialSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IdentityRotateUserCredentialSecretResponse**](identityRotateUserCredentialSecretResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

