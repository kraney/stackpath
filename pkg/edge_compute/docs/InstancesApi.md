# \InstancesApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWorkloadInstance**](InstancesApi.md#GetWorkloadInstance) | **Get** /workload/v1/stacks/{stack_id}/workloads/{workload_id}/instances/{instance_name} | Get a workload instance
[**GetWorkloadInstances**](InstancesApi.md#GetWorkloadInstances) | **Get** /workload/v1/stacks/{stack_id}/workloads/{workload_id}/instances | Get all workload instances
[**RestartInstance**](InstancesApi.md#RestartInstance) | **Post** /workload/v1/stacks/{stack_id}/workloads/{workload_id}/instances/{instance_name}/power/restart | Restart a workload instance



## GetWorkloadInstance

> V1GetWorkloadInstanceResponse GetWorkloadInstance(ctx, stackId, workloadId, instanceName)

Get a workload instance

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**workloadId** | **string**| An EdgeCompute workload ID | 
**instanceName** | **string**| An EdgeCompute workload instance name | 

### Return type

[**V1GetWorkloadInstanceResponse**](v1GetWorkloadInstanceResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkloadInstances

> V1GetWorkloadInstancesResponse GetWorkloadInstances(ctx, stackId, workloadId, optional)

Get all workload instances

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**workloadId** | **string**| An EdgeCompute workload ID | 
 **optional** | ***GetWorkloadInstancesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetWorkloadInstancesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageRequestFirst** | **optional.String**| The number of items desired. | 
 **pageRequestAfter** | **optional.String**| The cursor value after which data will be returned. | 
 **pageRequestFilter** | **optional.String**| SQL-style constraint filters. | 
 **pageRequestSortBy** | **optional.String**| Sort the response by the given field. | 

### Return type

[**V1GetWorkloadInstancesResponse**](v1GetWorkloadInstancesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestartInstance

> RestartInstance(ctx, stackId, workloadId, instanceName)

Restart a workload instance

The action is performed asynchronously and a successful response does not mean the instance has restarted yet.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**workloadId** | **string**| An EdgeCompute workload ID | 
**instanceName** | **string**| An EdgeCompute workload instance name | 

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

