# \InstanceLogsApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLogs**](InstanceLogsApi.md#GetLogs) | **Get** /workload/v1/stacks/{stack_id}/workloads/{workload_id}/instances/{instance_name}/logs | Get log stream



## GetLogs

> V1LogChunk GetLogs(ctx, stackId, workloadId, instanceName, optional)

Get log stream

Retrieve a stream of logs generated by a workload instance's containers. Logs are generated by the containers and are not modified by StackPath. Log contents vary by the application running in the container, though many containerized applications are configured to log to STDOUT and STDERR.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**workloadId** | **string**| An EdgeCompute workload ID | 
**instanceName** | **string**| An EdgeCompute workload instance name | 
 **optional** | ***GetLogsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetLogsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **containerName** | **optional.String**| The name of the container to obtain logs for. This defaults to first container in the instance. | 
 **follow** | **optional.Bool**| Whether or not to follow the instance&#39;s log stream. This defaults to false. | 
 **previous** | **optional.Bool**| Whether or not to return log entries made by previously terminated containers. This defaults to false. | 
 **sinceSeconds** | **optional.String**| A relative time in seconds before the current time from which to show logs. If this value precedes the time an instance was started, only logs since the instance&#39;s start time will be returned. If this value is in the future, no logs will be returned.  Only one of since_seconds or since_time may be specified | 
 **sinceTime** | **optional.Time**| An RFC3339 timestamp from which to show logs. If this value precedes the time an instance was started, only logs since the instance&#39;s start time will be returned. If this value is in the future, no logs will be returned.  Only one of since_seconds or since_time may be specified | 
 **timestamps** | **optional.Bool**| Whether or not to add an RFC3339 or RFC3339Nano timestamp at the beginning of every line of log output. This defaults to false. | 
 **tailLines** | **optional.String**| The number of lines from the end of the logs to show. If not specified, logs are shown from the creation of the container or since_seconds or since_time. | 
 **limitBytes** | **optional.String**| The number of bytes to read from the server before terminating log output. This may not display a complete final line of logging, and may return slightly more or slightly less than the specified limit. | 

### Return type

[**V1LogChunk**](v1LogChunk.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
