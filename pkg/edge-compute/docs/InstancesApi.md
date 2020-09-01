# \InstancesApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWorkloadInstance**](InstancesApi.md#GetWorkloadInstance) | **Get** /workload/v1/stacks/{stack_id}/workloads/{workload_id}/instances/{instance_name} | Get a workload instance
[**GetWorkloadInstances**](InstancesApi.md#GetWorkloadInstances) | **Get** /workload/v1/stacks/{stack_id}/workloads/{workload_id}/instances | Get all workload instances
[**RestartInstance**](InstancesApi.md#RestartInstance) | **Post** /workload/v1/stacks/{stack_id}/workloads/{workload_id}/instances/{instance_name}/power/restart | Restart a workload instance



## GetWorkloadInstance

> V1GetWorkloadInstanceResponse GetWorkloadInstance(ctx, stackId, workloadId, instanceName).Execute()

Get a workload instance

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.GetWorkloadInstance(context.Background(), stackId, workloadId, instanceName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.GetWorkloadInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkloadInstance`: V1GetWorkloadInstanceResponse
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.GetWorkloadInstance`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetWorkloadInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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

> V1GetWorkloadInstancesResponse GetWorkloadInstances(ctx, stackId, workloadId).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).Execute()

Get all workload instances

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
    pageRequestFirst := "pageRequestFirst_example" // string | The number of items desired. (optional)
    pageRequestAfter := "pageRequestAfter_example" // string | The cursor value after which data will be returned. (optional)
    pageRequestFilter := "pageRequestFilter_example" // string | SQL-style constraint filters. (optional)
    pageRequestSortBy := "pageRequestSortBy_example" // string | Sort the response by the given field. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.GetWorkloadInstances(context.Background(), stackId, workloadId).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.GetWorkloadInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkloadInstances`: V1GetWorkloadInstancesResponse
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.GetWorkloadInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**workloadId** | **string** | An EdgeCompute workload ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkloadInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageRequestFirst** | **string** | The number of items desired. | 
 **pageRequestAfter** | **string** | The cursor value after which data will be returned. | 
 **pageRequestFilter** | **string** | SQL-style constraint filters. | 
 **pageRequestSortBy** | **string** | Sort the response by the given field. | 

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

> RestartInstance(ctx, stackId, workloadId, instanceName).Execute()

Restart a workload instance



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.RestartInstance(context.Background(), stackId, workloadId, instanceName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.RestartInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
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

Other parameters are passed through a pointer to a apiRestartInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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

