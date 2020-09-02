# \AlertConditionsApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchDeleteAlertConditions**](AlertConditionsApi.md#BatchDeleteAlertConditions) | **Post** /monitoring/v2/stacks/{stack_id}/monitors/{monitor_id}/alerts/conditions/batch_delete | Delete multiple alert conditions
[**CreateAlertCondition**](AlertConditionsApi.md#CreateAlertCondition) | **Post** /monitoring/v2/stacks/{stack_id}/monitors/{monitor_id}/alerts/conditions | Create an alert condition
[**DeleteAlertCondition**](AlertConditionsApi.md#DeleteAlertCondition) | **Delete** /monitoring/v2/stacks/{stack_id}/monitors/{monitor_id}/alerts/conditions/{condition_id} | Delete an alert condition
[**DisableAlertCondition**](AlertConditionsApi.md#DisableAlertCondition) | **Post** /monitoring/v2/stacks/{stack_id}/monitors/{monitor_id}/alerts/conditions/{condition_id}/disable | Disable an alert condition
[**EnableAlertCondition**](AlertConditionsApi.md#EnableAlertCondition) | **Post** /monitoring/v2/stacks/{stack_id}/monitors/{monitor_id}/alerts/conditions/{condition_id}/enable | Enable an alert condition
[**GetAlertCondition**](AlertConditionsApi.md#GetAlertCondition) | **Get** /monitoring/v2/stacks/{stack_id}/monitors/{monitor_id}/alerts/conditions/{condition_id} | Get an alert condition
[**GetAlertConditions**](AlertConditionsApi.md#GetAlertConditions) | **Get** /monitoring/v2/stacks/{stack_id}/monitors/{monitor_id}/alerts/conditions | Get all alert conditions
[**UpdateAlertCondition**](AlertConditionsApi.md#UpdateAlertCondition) | **Patch** /monitoring/v2/stacks/{stack_id}/monitors/{monitor_id}/alerts/conditions/{condition_id} | Update an alert condition



## BatchDeleteAlertConditions

> BatchDeleteAlertConditions(ctx, stackId, monitorId, v2BatchDeleteAlertConditionsRequest)

Delete multiple alert conditions

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**monitorId** | **string**| A monitor ID | 
**v2BatchDeleteAlertConditionsRequest** | [**V2BatchDeleteAlertConditionsRequest**](V2BatchDeleteAlertConditionsRequest.md)|  | 

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


## CreateAlertCondition

> V2CreateAlertConditionResponse CreateAlertCondition(ctx, stackId, monitorId, v2CreateAlertConditionRequest)

Create an alert condition

An alert condition defines when to be alerted by a change in the monitored service.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**monitorId** | **string**| A monitor ID | 
**v2CreateAlertConditionRequest** | [**V2CreateAlertConditionRequest**](V2CreateAlertConditionRequest.md)|  | 

### Return type

[**V2CreateAlertConditionResponse**](v2CreateAlertConditionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAlertCondition

> DeleteAlertCondition(ctx, stackId, monitorId, conditionId)

Delete an alert condition

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**monitorId** | **string**| A monitor ID | 
**conditionId** | **string**| A monitoring alert condition ID | 

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


## DisableAlertCondition

> DisableAlertCondition(ctx, stackId, monitorId, conditionId)

Disable an alert condition

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**monitorId** | **string**| A monitor ID | 
**conditionId** | **string**| A monitoring alert condition ID | 

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


## EnableAlertCondition

> EnableAlertCondition(ctx, stackId, monitorId, conditionId)

Enable an alert condition

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**monitorId** | **string**| A monitor ID | 
**conditionId** | **string**| A monitoring alert condition ID | 

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


## GetAlertCondition

> V2GetAlertConditionResponse GetAlertCondition(ctx, stackId, monitorId, conditionId)

Get an alert condition

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**monitorId** | **string**| A monitor ID | 
**conditionId** | **string**| A monitoring alert condition ID | 

### Return type

[**V2GetAlertConditionResponse**](v2GetAlertConditionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlertConditions

> V2GetAlertConditionsResponse GetAlertConditions(ctx, stackId, monitorId, optional)

Get all alert conditions

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**monitorId** | **string**| A monitor ID | 
 **optional** | ***GetAlertConditionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAlertConditionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageRequestFirst** | **optional.String**| The number of items desired. | 
 **pageRequestAfter** | **optional.String**| The cursor value after which data will be returned. | 
 **pageRequestFilter** | **optional.String**| SQL-style constraint filters. | 
 **pageRequestSortBy** | **optional.String**| Sort the response by the given field. | 

### Return type

[**V2GetAlertConditionsResponse**](v2GetAlertConditionsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAlertCondition

> V2UpdateAlertConditionResponse UpdateAlertCondition(ctx, stackId, monitorId, conditionId, v2UpdateAlertConditionRequest)

Update an alert condition

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**monitorId** | **string**| A monitor ID | 
**conditionId** | **string**| A monitoring alert condition ID | 
**v2UpdateAlertConditionRequest** | [**V2UpdateAlertConditionRequest**](V2UpdateAlertConditionRequest.md)|  | 

### Return type

[**V2UpdateAlertConditionResponse**](v2UpdateAlertConditionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

