# \StacksApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStack**](StacksApi.md#CreateStack) | **Post** /stack/v1/stacks | Create a stack
[**GetStack**](StacksApi.md#GetStack) | **Get** /stack/v1/stacks/{stack_id} | Get a stack
[**GetStacks**](StacksApi.md#GetStacks) | **Get** /stack/v1/stacks | Get all stacks
[**UpdateStack2**](StacksApi.md#UpdateStack2) | **Patch** /stack/v1/stacks/{stack_id} | Update a stack



## CreateStack

> StackCreateStackResponse CreateStack(ctx, stackCreateStackRequest)

Create a stack

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackCreateStackRequest** | [**StackCreateStackRequest**](StackCreateStackRequest.md)|  | 

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

> StackStack GetStack(ctx, stackId)

Get a stack

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 

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

> StackGetStacksResponse GetStacks(ctx, optional)

Get all stacks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetStacksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetStacksOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageRequestFirst** | **optional.String**| The number of items desired. | 
 **pageRequestAfter** | **optional.String**| The cursor value after which data will be returned. | 
 **pageRequestFilter** | **optional.String**| SQL-style constraint filters. | 
 **pageRequestSortBy** | **optional.String**| Sort the response by the given field. | 
 **accountId** | **optional.String**| An account ID | 

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

> StackUpdateStackResponse UpdateStack2(ctx, stackId, stackUpdateStackRequest2)

Update a stack

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**stackUpdateStackRequest2** | [**StackUpdateStackRequest2**](StackUpdateStackRequest2.md)|  | 

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

