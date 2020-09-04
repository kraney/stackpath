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

> V1CreateWorkloadResponse CreateWorkload(ctx, stackId).V1CreateWorkloadRequest(v1CreateWorkloadRequest).Execute()

Create a workload

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
    v1CreateWorkloadRequest := openapiclient.v1CreateWorkloadRequest{Workload: openapiclient.v1Workload{StackId: "StackId_example", Id: "Id_example", Name: "Name_example", Slug: "Slug_example", Metadata: openapiclient.v1Metadata{Annotations: map[string]string{ "Key" = "Value" }, Labels: map[string]string{ "Key" = "Value" }, CreatedAt: "TODO", UpdatedAt: "TODO", DeleteRequestedAt: "TODO", Version: "Version_example"}, Spec: openapiclient.v1WorkloadSpec{NetworkInterfaces: []V1NetworkInterface{openapiclient.v1NetworkInterface{Network: "Network_example"}), Containers: map[string]string{ "Key" = "Value" }, VirtualMachines: map[string]string{ "Key" = "Value" }, VolumeClaimTemplates: []V1VolumeClaim{openapiclient.v1VolumeClaim{StackId: "StackId_example", Id: "Id_example", Name: "Name_example", Slug: "Slug_example", Metadata: openapiclient.v1Metadata{Annotations: map[string]string{ "Key" = "Value" }, Labels: map[string]string{ "Key" = "Value" }, CreatedAt: "TODO", UpdatedAt: "TODO", DeleteRequestedAt: "TODO", Version: "Version_example"}, Spec: openapiclient.v1VolumeClaimSpec{Resources: openapiclient.v1ResourceRequirements{Requests: map[string]string{ "Key" = "Value" }, Limits: map[string]string{ "Key" = "Value" }}}, Phase: openapiclient.VolumeClaimVolumeClaimPhase{}}), ImagePullCredentials: []V1ImagePullCredential{openapiclient.v1ImagePullCredential{DockerRegistry: openapiclient.v1DockerRegistryCredentials{Server: "Server_example", Username: "Username_example", Password: "Password_example", Email: "Email_example"}})}, Targets: map[string]string{ "Key" = "Value" }, Status: openapiclient.v1WorkloadStatus{}}} // V1CreateWorkloadRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkloadsApi.CreateWorkload(context.Background(), stackId, v1CreateWorkloadRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadsApi.CreateWorkload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorkload`: V1CreateWorkloadResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkloadsApi.CreateWorkload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v1CreateWorkloadRequest** | [**V1CreateWorkloadRequest**](V1CreateWorkloadRequest.md) |  | 

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

> DeleteWorkload(ctx, stackId, workloadId).Execute()

Delete a workload

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkloadsApi.DeleteWorkload(context.Background(), stackId, workloadId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadsApi.DeleteWorkload``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkloadRequest struct via the builder pattern


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


## GetWorkload

> V1GetWorkloadResponse GetWorkload(ctx, stackId, workloadId).Execute()

Get a workload

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkloadsApi.GetWorkload(context.Background(), stackId, workloadId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadsApi.GetWorkload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkload`: V1GetWorkloadResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkloadsApi.GetWorkload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**workloadId** | **string** | An EdgeCompute workload ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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

> V1GetWorkloadsResponse GetWorkloads(ctx, stackId).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).Execute()

Get all workloads

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
    pageRequestFirst := "pageRequestFirst_example" // string | The number of items desired. (optional)
    pageRequestAfter := "pageRequestAfter_example" // string | The cursor value after which data will be returned. (optional)
    pageRequestFilter := "pageRequestFilter_example" // string | SQL-style constraint filters. (optional)
    pageRequestSortBy := "pageRequestSortBy_example" // string | Sort the response by the given field. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkloadsApi.GetWorkloads(context.Background(), stackId).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadsApi.GetWorkloads``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkloads`: V1GetWorkloadsResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkloadsApi.GetWorkloads`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkloadsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageRequestFirst** | **string** | The number of items desired. | 
 **pageRequestAfter** | **string** | The cursor value after which data will be returned. | 
 **pageRequestFilter** | **string** | SQL-style constraint filters. | 
 **pageRequestSortBy** | **string** | Sort the response by the given field. | 

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

> V1UpdateWorkloadResponse UpdateWorkload(ctx, stackId, workloadId).V1UpdateWorkloadRequest(v1UpdateWorkloadRequest).Execute()

Update a workload

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
    v1UpdateWorkloadRequest := openapiclient.v1UpdateWorkloadRequest{Workload: openapiclient.v1Workload{StackId: "StackId_example", Id: "Id_example", Name: "Name_example", Slug: "Slug_example", Metadata: , Spec: openapiclient.v1WorkloadSpec{NetworkInterfaces: []V1NetworkInterface{openapiclient.v1NetworkInterface{Network: "Network_example"}), Containers: map[string]string{ "Key" = "Value" }, VirtualMachines: map[string]string{ "Key" = "Value" }, VolumeClaimTemplates: []V1VolumeClaim{openapiclient.v1VolumeClaim{StackId: "StackId_example", Id: "Id_example", Name: "Name_example", Slug: "Slug_example", Metadata: , Spec: openapiclient.v1VolumeClaimSpec{Resources: openapiclient.v1ResourceRequirements{Requests: map[string]string{ "Key" = "Value" }, Limits: map[string]string{ "Key" = "Value" }}}, Phase: openapiclient.VolumeClaimVolumeClaimPhase{}}), ImagePullCredentials: []V1ImagePullCredential{openapiclient.v1ImagePullCredential{DockerRegistry: openapiclient.v1DockerRegistryCredentials{Server: "Server_example", Username: "Username_example", Password: "Password_example", Email: "Email_example"}})}, Targets: map[string]string{ "Key" = "Value" }, Status: openapiclient.v1WorkloadStatus{}}} // V1UpdateWorkloadRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkloadsApi.UpdateWorkload(context.Background(), stackId, workloadId, v1UpdateWorkloadRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadsApi.UpdateWorkload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorkload`: V1UpdateWorkloadResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkloadsApi.UpdateWorkload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**workloadId** | **string** | An EdgeCompute workload ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **v1UpdateWorkloadRequest** | [**V1UpdateWorkloadRequest**](V1UpdateWorkloadRequest.md) |  | 

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

