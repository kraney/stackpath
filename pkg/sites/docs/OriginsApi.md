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

> DeliveryConnectScopeToOriginResponse ConnectScopeToOrigin(ctx, stackId, siteId, scopeId, deliveryConnectScopeToOriginRequest)

Connect an origin to a scope

The origin is automatically created if necessary. When the request contains a priority which an origin already associated with the scope has set, the existing origin is disconnected. The priority of an origin already associated with a scope can be modified via this endpoint.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
**scopeId** | **string**| A scope ID | 
**deliveryConnectScopeToOriginRequest** | [**DeliveryConnectScopeToOriginRequest**](DeliveryConnectScopeToOriginRequest.md)|  | 

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

> DeliveryGetOriginResponse GetOrigin(ctx, stackId, originId)

Get an origin

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**originId** | **string**| An origin ID | 

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

> DeliveryGetOriginsResponse GetOrigins(ctx, stackId, optional)

Get all origins

Retrieve all origins associated with a site's stack

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
 **optional** | ***GetOriginsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetOriginsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageRequestFirst** | **optional.String**| The number of items desired. | 
 **pageRequestAfter** | **optional.String**| The cursor value after which data will be returned. | 
 **pageRequestFilter** | **optional.String**| SQL-style constraint filters. | 
 **pageRequestSortBy** | **optional.String**| Sort the response by the given field. | 

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

> DeliveryGetScopeOriginsResponse GetScopeOrigins(ctx, stackId, siteId, scopeId, optional)

Get a scope's origins

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
**scopeId** | **string**| A scope ID | 
 **optional** | ***GetScopeOriginsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetScopeOriginsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pageRequestFirst** | **optional.String**| The number of items desired. | 
 **pageRequestAfter** | **optional.String**| The cursor value after which data will be returned. | 
 **pageRequestFilter** | **optional.String**| SQL-style constraint filters. | 
 **pageRequestSortBy** | **optional.String**| Sort the response by the given field. | 

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

> DeliveryUpdateOriginResponse UpdateOrigin(ctx, stackId, originId, deliveryUpdateOriginRequest)

Update an origin

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**originId** | **string**| An origin ID | 
**deliveryUpdateOriginRequest** | [**DeliveryUpdateOriginRequest**](DeliveryUpdateOriginRequest.md)|  | 

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

