# \UserCredentialsApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCredential**](UserCredentialsApi.md#DeleteCredential) | **Delete** /storage/v1/stacks/{stack_id}/users/{user_id}/credentials/{access_key} | Delete credentials
[**GenerateCredentials**](UserCredentialsApi.md#GenerateCredentials) | **Post** /storage/v1/stacks/{stack_id}/users/{user_id}/credentials/generate | Create credentials
[**GetCredentials**](UserCredentialsApi.md#GetCredentials) | **Get** /storage/v1/stacks/{stack_id}/users/{user_id}/credentials | Get credentials



## DeleteCredential

> DeleteCredential(ctx, stackId, userId, accessKey)

Delete credentials

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**userId** | **string**| A user ID | 
**accessKey** | **string**| A user&#39;s access key | 

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

> StorageGenerateCredentialsResponse GenerateCredentials(ctx, stackId, userId)

Create credentials

Generate storage credentials for the given user. Users can only have one set of credentials. Calling this method will generate a new set and invalidate any existing ones.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**userId** | **string**| A user ID | 

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

> StorageGetCredentialsResponse GetCredentials(ctx, stackId, userId)

Get credentials

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**userId** | **string**| A user ID | 

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

