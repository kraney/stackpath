# \UsersApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeUserEmail**](UsersApi.md#ChangeUserEmail) | **Post** /identity/v1/users/{user_id}/change_email | Update a user&#39;s email address
[**CreateUser**](UsersApi.md#CreateUser) | **Post** /identity/v1/accounts/{account_id}/users | Create a user
[**DeleteUser**](UsersApi.md#DeleteUser) | **Delete** /identity/v1/users/{user_id} | Delete a user
[**GetAccountUsers**](UsersApi.md#GetAccountUsers) | **Get** /identity/v1/accounts/{account_id}/users | Get all users
[**GetUser**](UsersApi.md#GetUser) | **Get** /identity/v1/users/{user_id} | Get a user
[**UpdateUser**](UsersApi.md#UpdateUser) | **Patch** /identity/v1/users/{user_id} | Update a user



## ChangeUserEmail

> ChangeUserEmail(ctx, userId, identityChangeUserEmailRequest)

Update a user's email address

This immediately invalidates the user's StackPath customer portal logins and API tokens.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| A user ID | 
**identityChangeUserEmailRequest** | [**IdentityChangeUserEmailRequest**](IdentityChangeUserEmailRequest.md)|  | 

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


## CreateUser

> IdentityCreateUserResponse CreateUser(ctx, accountId, identityCreateUserRequest)

Create a user

The new user will receive an email to set their password.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| An account ID | 
**identityCreateUserRequest** | [**IdentityCreateUserRequest**](IdentityCreateUserRequest.md)|  | 

### Return type

[**IdentityCreateUserResponse**](identityCreateUserResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUser

> DeleteUser(ctx, userId)

Delete a user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| A user ID | 

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


## GetAccountUsers

> IdentityGetAccountUsersResponse GetAccountUsers(ctx, accountId, optional)

Get all users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| An account ID | 
 **optional** | ***GetAccountUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAccountUsersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageRequestFirst** | **optional.String**| The number of items desired. | 
 **pageRequestAfter** | **optional.String**| The cursor value after which data will be returned. | 
 **pageRequestFilter** | **optional.String**| SQL-style constraint filters. | 
 **pageRequestSortBy** | **optional.String**| Sort the response by the given field. | 

### Return type

[**IdentityGetAccountUsersResponse**](identityGetAccountUsersResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> IdentityGetUserResponse GetUser(ctx, userId)

Get a user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| A user ID | 

### Return type

[**IdentityGetUserResponse**](identityGetUserResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> IdentityUpdateUserResponse UpdateUser(ctx, userId, identityUpdateUserRequest)

Update a user

Update a user's non-essential properties, such as their phone number.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| A user ID | 
**identityUpdateUserRequest** | [**IdentityUpdateUserRequest**](IdentityUpdateUserRequest.md)|  | 

### Return type

[**IdentityUpdateUserResponse**](identityUpdateUserResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

