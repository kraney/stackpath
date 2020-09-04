# \InstanceLogsApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLogs**](InstanceLogsApi.md#GetLogs) | **Get** /workload/v1/stacks/{stack_id}/workloads/{workload_id}/instances/{instance_name}/logs | Get log stream



## GetLogs

> V1LogChunk GetLogs(ctx, stackId, workloadId, instanceName).ContainerName(containerName).Follow(follow).Previous(previous).SinceSeconds(sinceSeconds).SinceTime(sinceTime).Timestamps(timestamps).TailLines(tailLines).LimitBytes(limitBytes).Execute()

Get log stream



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    stackId := "stackId_example" // string | A stack ID or slug
    workloadId := "workloadId_example" // string | An EdgeCompute workload ID
    instanceName := "instanceName_example" // string | An EdgeCompute workload instance name
    containerName := "containerName_example" // string | The name of the container to obtain logs for. This defaults to first container in the instance. (optional)
    follow := true // bool | Whether or not to follow the instance's log stream. This defaults to false. (optional)
    previous := true // bool | Whether or not to return log entries made by previously terminated containers. This defaults to false. (optional)
    sinceSeconds := "sinceSeconds_example" // string | A relative time in seconds before the current time from which to show logs. If this value precedes the time an instance was started, only logs since the instance's start time will be returned. If this value is in the future, no logs will be returned.  Only one of since_seconds or since_time may be specified (optional)
    sinceTime := Get-Date // time.Time | An RFC3339 timestamp from which to show logs. If this value precedes the time an instance was started, only logs since the instance's start time will be returned. If this value is in the future, no logs will be returned.  Only one of since_seconds or since_time may be specified (optional)
    timestamps := true // bool | Whether or not to add an RFC3339 or RFC3339Nano timestamp at the beginning of every line of log output. This defaults to false. (optional)
    tailLines := "tailLines_example" // string | The number of lines from the end of the logs to show. If not specified, logs are shown from the creation of the container or since_seconds or since_time. (optional)
    limitBytes := "limitBytes_example" // string | The number of bytes to read from the server before terminating log output. This may not display a complete final line of logging, and may return slightly more or slightly less than the specified limit. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstanceLogsApi.GetLogs(context.Background(), stackId, workloadId, instanceName).ContainerName(containerName).Follow(follow).Previous(previous).SinceSeconds(sinceSeconds).SinceTime(sinceTime).Timestamps(timestamps).TailLines(tailLines).LimitBytes(limitBytes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstanceLogsApi.GetLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogs`: V1LogChunk
    fmt.Fprintf(os.Stdout, "Response from `InstanceLogsApi.GetLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**workloadId** | **string** | An EdgeCompute workload ID | 
**instanceName** | **string** | An EdgeCompute workload instance name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **containerName** | **string** | The name of the container to obtain logs for. This defaults to first container in the instance. | 
 **follow** | **bool** | Whether or not to follow the instance&#39;s log stream. This defaults to false. | 
 **previous** | **bool** | Whether or not to return log entries made by previously terminated containers. This defaults to false. | 
 **sinceSeconds** | **string** | A relative time in seconds before the current time from which to show logs. If this value precedes the time an instance was started, only logs since the instance&#39;s start time will be returned. If this value is in the future, no logs will be returned.  Only one of since_seconds or since_time may be specified | 
 **sinceTime** | **time.Time** | An RFC3339 timestamp from which to show logs. If this value precedes the time an instance was started, only logs since the instance&#39;s start time will be returned. If this value is in the future, no logs will be returned.  Only one of since_seconds or since_time may be specified | 
 **timestamps** | **bool** | Whether or not to add an RFC3339 or RFC3339Nano timestamp at the beginning of every line of log output. This defaults to false. | 
 **tailLines** | **string** | The number of lines from the end of the logs to show. If not specified, logs are shown from the creation of the container or since_seconds or since_time. | 
 **limitBytes** | **string** | The number of bytes to read from the server before terminating log output. This may not display a complete final line of logging, and may return slightly more or slightly less than the specified limit. | 

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

