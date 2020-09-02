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

> IdentityCreateUserCredentialResponse CreateUserCredential(ctx, userId, identityCreateUserCredentialRequest)

Create API credentials

The client secret is returned only once and is not stored by StackPath. Please take care to save this response.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| A user ID | 
**identityCreateUserCredentialRequest** | [**IdentityCreateUserCredentialRequest**](IdentityCreateUserCredentialRequest.md)|  | 

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

> DeleteUserCredential(ctx, userId, credentialId)

Delete API credentials

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| A user ID | 
**credentialId** | **string**| The ID of the API credentials to remove | 

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

> IdentityGetAccessTokenResponse GetAccessToken(ctx, identityGetAccessTokenRequest)

Generate a bearer token

Authenticate to the StackPath API. Use the resulting bearer token to authorize other StackPath API calls. API credentials can be generated in the [StackPath customer portal](https://control.stackpath.com/).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityGetAccessTokenRequest** | [**IdentityGetAccessTokenRequest**](IdentityGetAccessTokenRequest.md)|  | 

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

> IdentityGetUserCredentialsResponse GetUserCredentials(ctx, userId, optional)

Get all API credentials

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| A user ID | 
 **optional** | ***GetUserCredentialsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetUserCredentialsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageRequestFirst** | **optional.String**| The number of items desired. | 
 **pageRequestAfter** | **optional.String**| The cursor value after which data will be returned. | 
 **pageRequestFilter** | **optional.String**| SQL-style constraint filters. | 
 **pageRequestSortBy** | **optional.String**| Sort the response by the given field. | 

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

> IdentityRotateUserCredentialSecretResponse RotateUserCredentialSecret(ctx, userId, credentialId)

Create a new API secret

The client secret is returned only once and is not stored by StackPath. Please take care to save this response.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| A user ID | 
**credentialId** | **string**| The ID of the API client credential to rotate secrets for | 

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

