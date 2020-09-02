# \WorkloadsApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWorkload**](WorkloadsApi.md#CreateWorkload) | **Post** /workload/v1/stacks/{stack_id}/workloads | Create a workload
[**DeleteWorkload**](WorkloadsApi.md#DeleteWorkload) | **Delete** /workload/v1/stacks/{stack_id}/workloads/{workload_id} | Delete a workload
[**GetWorkload**](WorkloadsApi.md#GetWorkload) | **Get** /workload/v1/stacks/{stack_id}/workloads/{workload_id} | Get a workload
[**GetWorkloads**](WorkloadsApi.md#GetWorkloads) | **Get** /workload/v1/stacks/{stack_id}/workloads | Get all workloads
[**UpdateWorkload**](WorkloadsApi.md#UpdateWorkload) | **Patch** /workload/v1/stacks/{stack_id}/workloads/{workload_id} | Update a workload



## CreateWorkload

> V1CreateWorkloadResponse CreateWorkload(ctx, stackId, v1CreateWorkloadRequest)

Create a workload

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**v1CreateWorkloadRequest** | [**V1CreateWorkloadRequest**](V1CreateWorkloadRequest.md)|  | 

### Return type

[**V1CreateWorkloadResponse**](v1CreateWorkloadResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWorkload

> DeleteWorkload(ctx, stackId, workloadId)

Delete a workload

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**workloadId** | **string**| An EdgeCompute workload ID | 

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


## GetWorkload

> V1GetWorkloadResponse GetWorkload(ctx, stackId, workloadId)

Get a workload

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**workloadId** | **string**| An EdgeCompute workload ID | 

### Return type

[**V1GetWorkloadResponse**](v1GetWorkloadResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkloads

> V1GetWorkloadsResponse GetWorkloads(ctx, stackId, optional)

Get all workloads

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
 **optional** | ***GetWorkloadsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetWorkloadsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageRequestFirst** | **optional.String**| The number of items desired. | 
 **pageRequestAfter** | **optional.String**| The cursor value after which data will be returned. | 
 **pageRequestFilter** | **optional.String**| SQL-style constraint filters. | 
 **pageRequestSortBy** | **optional.String**| Sort the response by the given field. | 

### Return type

[**V1GetWorkloadsResponse**](v1GetWorkloadsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkload

> V1UpdateWorkloadResponse UpdateWorkload(ctx, stackId, workloadId, v1UpdateWorkloadRequest)

Update a workload

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**workloadId** | **string**| An EdgeCompute workload ID | 
**v1UpdateWorkloadRequest** | [**V1UpdateWorkloadRequest**](V1UpdateWorkloadRequest.md)|  | 

### Return type

[**V1UpdateWorkloadResponse**](v1UpdateWorkloadResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

