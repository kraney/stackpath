# \ServerlessScriptingApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSiteScript**](ServerlessScriptingApi.md#CreateSiteScript) | **Post** /cdn/v1/stacks/{stack_id}/sites/{site_id}/scripts | Create a serverless script
[**DeleteSiteScript**](ServerlessScriptingApi.md#DeleteSiteScript) | **Delete** /cdn/v1/stacks/{stack_id}/sites/{site_id}/scripts/{script_id} | Delete a serverless script
[**GetSiteScript**](ServerlessScriptingApi.md#GetSiteScript) | **Get** /cdn/v1/stacks/{stack_id}/sites/{site_id}/scripts/{script_id} | Get a serverless script
[**GetSiteScript2**](ServerlessScriptingApi.md#GetSiteScript2) | **Get** /cdn/v1/stacks/{stack_id}/sites/{site_id}/scripts/{script_id}/{script_version} | Get a serverless script version
[**GetSiteScripts**](ServerlessScriptingApi.md#GetSiteScripts) | **Get** /cdn/v1/stacks/{stack_id}/sites/{site_id}/scripts | Get all serverless scripts
[**UpdateSiteScript**](ServerlessScriptingApi.md#UpdateSiteScript) | **Patch** /cdn/v1/stacks/{stack_id}/sites/{site_id}/scripts/{script_id} | Update a serverless script



## CreateSiteScript

> CdnCreateSiteScriptResponse CreateSiteScript(ctx, stackId, siteId, cdnCreateSiteScriptRequest)

Create a serverless script

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
**cdnCreateSiteScriptRequest** | [**CdnCreateSiteScriptRequest**](CdnCreateSiteScriptRequest.md)|  | 

### Return type

[**CdnCreateSiteScriptResponse**](cdnCreateSiteScriptResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSiteScript

> DeleteSiteScript(ctx, stackId, siteId, scriptId)

Delete a serverless script

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
**scriptId** | **string**| A serverless script ID | 

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


## GetSiteScript

> CdnGetSiteScriptResponse GetSiteScript(ctx, stackId, siteId, scriptId, optional)

Get a serverless script

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
**scriptId** | **string**| A serverless script ID | 
 **optional** | ***GetSiteScriptOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSiteScriptOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **scriptVersion** | **optional.String**| The version of the serverless script to get | 

### Return type

[**CdnGetSiteScriptResponse**](cdnGetSiteScriptResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteScript2

> CdnGetSiteScriptResponse GetSiteScript2(ctx, stackId, siteId, scriptId, scriptVersion)

Get a serverless script version

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
**scriptId** | **string**| A serverless script ID | 
**scriptVersion** | **string**| The version of the serverless script to get | 

### Return type

[**CdnGetSiteScriptResponse**](cdnGetSiteScriptResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteScripts

> CdnGetSiteScriptsResponse GetSiteScripts(ctx, stackId, siteId, optional)

Get all serverless scripts

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
 **optional** | ***GetSiteScriptsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSiteScriptsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageRequestFirst** | **optional.String**| The number of items desired. | 
 **pageRequestAfter** | **optional.String**| The cursor value after which data will be returned. | 
 **pageRequestFilter** | **optional.String**| SQL-style constraint filters. | 
 **pageRequestSortBy** | **optional.String**| Sort the response by the given field. | 

### Return type

[**CdnGetSiteScriptsResponse**](cdnGetSiteScriptsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSiteScript

> CdnUpdateSiteScriptResponse UpdateSiteScript(ctx, stackId, siteId, scriptId, cdnUpdateSiteScriptRequest)

Update a serverless script

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
**scriptId** | **string**| A serverless script ID | 
**cdnUpdateSiteScriptRequest** | [**CdnUpdateSiteScriptRequest**](CdnUpdateSiteScriptRequest.md)|  | 

### Return type

[**CdnUpdateSiteScriptResponse**](cdnUpdateSiteScriptResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

