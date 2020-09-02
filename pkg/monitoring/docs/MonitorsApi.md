# \MonitorsApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchDeleteMonitors**](MonitorsApi.md#BatchDeleteMonitors) | **Post** /monitoring/v2/stacks/{stack_id}/monitors/batch_delete | Delete multiple monitors
[**CreateMonitor**](MonitorsApi.md#CreateMonitor) | **Post** /monitoring/v2/stacks/{stack_id}/monitors | Create a monitor
[**DeleteMonitor**](MonitorsApi.md#DeleteMonitor) | **Delete** /monitoring/v2/stacks/{stack_id}/monitors/{monitor_id} | Delete a monitor
[**Disable**](MonitorsApi.md#Disable) | **Post** /monitoring/v2/stacks/{stack_id}/disable | Disable all monitors
[**Enable**](MonitorsApi.md#Enable) | **Post** /monitoring/v2/stacks/{stack_id}/enable | Enable all monitors
[**GetMonitor**](MonitorsApi.md#GetMonitor) | **Get** /monitoring/v2/stacks/{stack_id}/monitors/{monitor_id} | Get a monitor
[**GetMonitorErrors**](MonitorsApi.md#GetMonitorErrors) | **Get** /monitoring/v2/stacks/{stack_id}/monitors/{monitor_id}/errors | Get monitoring errors
[**GetMonitorLocations**](MonitorsApi.md#GetMonitorLocations) | **Get** /monitoring/v2/stacks/{stack_id}/monitors/{monitor_id}/locations | Get a monitor&#39;s locations
[**GetMonitors**](MonitorsApi.md#GetMonitors) | **Get** /monitoring/v2/stacks/{stack_id}/monitors | Get all monitors
[**ReplaceMonitor**](MonitorsApi.md#ReplaceMonitor) | **Put** /monitoring/v2/stacks/{stack_id}/monitors/{monitor_id} | Replace a monitor
[**UpdateMonitor**](MonitorsApi.md#UpdateMonitor) | **Patch** /monitoring/v2/stacks/{stack_id}/monitors/{monitor_id} | Update a monitor



## BatchDeleteMonitors

> BatchDeleteMonitors(ctx, stackId, v2BatchDeleteMonitorsRequest)

Delete multiple monitors

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**v2BatchDeleteMonitorsRequest** | [**V2BatchDeleteMonitorsRequest**](V2BatchDeleteMonitorsRequest.md)|  | 

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


## CreateMonitor

> V2CreateMonitorResponse CreateMonitor(ctx, stackId, v2CreateMonitorRequest)

Create a monitor

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**v2CreateMonitorRequest** | [**V2CreateMonitorRequest**](V2CreateMonitorRequest.md)|  | 

### Return type

[**V2CreateMonitorResponse**](v2CreateMonitorResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMonitor

> DeleteMonitor(ctx, stackId, monitorId)

Delete a monitor

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**monitorId** | **string**| A monitor ID | 

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


## Disable

> Disable(ctx, stackId)

Disable all monitors

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 

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


## Enable

> Enable(ctx, stackId)

Enable all monitors

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 

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


## GetMonitor

> V2GetMonitorResponse GetMonitor(ctx, stackId, monitorId)

Get a monitor

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**monitorId** | **string**| A monitor ID | 

### Return type

[**V2GetMonitorResponse**](v2GetMonitorResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMonitorErrors

> V2GetMonitorErrorsResponse GetMonitorErrors(ctx, stackId, monitorId)

Get monitoring errors

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**monitorId** | **string**| A monitor ID | 

### Return type

[**V2GetMonitorErrorsResponse**](v2GetMonitorErrorsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMonitorLocations

> V2GetMonitorLocationsResponse GetMonitorLocations(ctx, stackId, monitorId)

Get a monitor's locations

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**monitorId** | **string**| A monitor ID | 

### Return type

[**V2GetMonitorLocationsResponse**](v2GetMonitorLocationsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMonitors

> V2GetMonitorsResponse GetMonitors(ctx, stackId, optional)

Get all monitors

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
 **optional** | ***GetMonitorsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetMonitorsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageRequestFirst** | **optional.String**| The number of items desired. | 
 **pageRequestAfter** | **optional.String**| The cursor value after which data will be returned. | 
 **pageRequestFilter** | **optional.String**| SQL-style constraint filters. | 
 **pageRequestSortBy** | **optional.String**| Sort the response by the given field. | 

### Return type

[**V2GetMonitorsResponse**](v2GetMonitorsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceMonitor

> V2ReplaceMonitorResponse ReplaceMonitor(ctx, stackId, monitorId, v2ReplaceMonitorRequest)

Replace a monitor

When replacing a monitor you must provide all fields or they will be overwritten with empty or default values.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**monitorId** | **string**| A monitor ID | 
**v2ReplaceMonitorRequest** | [**V2ReplaceMonitorRequest**](V2ReplaceMonitorRequest.md)|  | 

### Return type

[**V2ReplaceMonitorResponse**](v2ReplaceMonitorResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMonitor

> V2UpdateMonitorResponse UpdateMonitor(ctx, stackId, monitorId, v2UpdateMonitorRequest)

Update a monitor

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**monitorId** | **string**| A monitor ID | 
**v2UpdateMonitorRequest** | [**V2UpdateMonitorRequest**](V2UpdateMonitorRequest.md)|  | 

### Return type

[**V2UpdateMonitorResponse**](v2UpdateMonitorResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

