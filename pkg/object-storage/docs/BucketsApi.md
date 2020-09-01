# \BucketsApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBucket**](BucketsApi.md#CreateBucket) | **Post** /storage/v1/stacks/{stack_id}/buckets | Create a bucket
[**DeleteBucket**](BucketsApi.md#DeleteBucket) | **Delete** /storage/v1/stacks/{stack_id}/buckets/{bucket_id} | Delete a bucket
[**GetBucket**](BucketsApi.md#GetBucket) | **Get** /storage/v1/stacks/{stack_id}/buckets/{bucket_id} | Get a bucket
[**GetBuckets**](BucketsApi.md#GetBuckets) | **Get** /storage/v1/stacks/{stack_id}/buckets | Get all buckets
[**UpdateBucket**](BucketsApi.md#UpdateBucket) | **Put** /storage/v1/stacks/{stack_id}/buckets/{bucket_id} | Update a bucket



## CreateBucket

> StorageCreateBucketResponse CreateBucket(ctx, stackId).StorageCreateBucketRequest(storageCreateBucketRequest).Execute()

Create a bucket

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
    storageCreateBucketRequest := openapiclient.storageCreateBucketRequest{Label: "Label_example", Region: "Region_example"} // StorageCreateBucketRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BucketsApi.CreateBucket(context.Background(), stackId, storageCreateBucketRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketsApi.CreateBucket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBucket`: StorageCreateBucketResponse
    fmt.Fprintf(os.Stdout, "Response from `BucketsApi.CreateBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **storageCreateBucketRequest** | [**StorageCreateBucketRequest**](StorageCreateBucketRequest.md) |  | 

### Return type

[**StorageCreateBucketResponse**](storageCreateBucketResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBucket

> DeleteBucket(ctx, stackId, bucketId).ForceDelete(forceDelete).Execute()

Delete a bucket

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
    bucketId := "bucketId_example" // string | A storage bucket ID
    forceDelete := true // bool | Force bucket deletion even if there are contents inside it (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BucketsApi.DeleteBucket(context.Background(), stackId, bucketId).ForceDelete(forceDelete).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketsApi.DeleteBucket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**bucketId** | **string** | A storage bucket ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **forceDelete** | **bool** | Force bucket deletion even if there are contents inside it | 

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


## GetBucket

> StorageGetBucketResponse GetBucket(ctx, stackId, bucketId).Execute()

Get a bucket

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
    bucketId := "bucketId_example" // string | A storage bucket ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BucketsApi.GetBucket(context.Background(), stackId, bucketId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketsApi.GetBucket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBucket`: StorageGetBucketResponse
    fmt.Fprintf(os.Stdout, "Response from `BucketsApi.GetBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**bucketId** | **string** | A storage bucket ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**StorageGetBucketResponse**](storageGetBucketResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBuckets

> StorageGetBucketsResponse GetBuckets(ctx, stackId).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).Execute()

Get all buckets

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
    resp, r, err := api_client.BucketsApi.GetBuckets(context.Background(), stackId).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketsApi.GetBuckets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuckets`: StorageGetBucketsResponse
    fmt.Fprintf(os.Stdout, "Response from `BucketsApi.GetBuckets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBucketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageRequestFirst** | **string** | The number of items desired. | 
 **pageRequestAfter** | **string** | The cursor value after which data will be returned. | 
 **pageRequestFilter** | **string** | SQL-style constraint filters. | 
 **pageRequestSortBy** | **string** | Sort the response by the given field. | 

### Return type

[**StorageGetBucketsResponse**](storageGetBucketsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBucket

> StorageUpdateBucketResponse UpdateBucket(ctx, stackId, bucketId).StorageUpdateBucketRequest(storageUpdateBucketRequest).Execute()

Update a bucket

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
    bucketId := "bucketId_example" // string | A storage bucket ID
    storageUpdateBucketRequest := openapiclient.storageUpdateBucketRequest{Visibility: openapiclient.storageBucketVisibility{}} // StorageUpdateBucketRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BucketsApi.UpdateBucket(context.Background(), stackId, bucketId, storageUpdateBucketRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketsApi.UpdateBucket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBucket`: StorageUpdateBucketResponse
    fmt.Fprintf(os.Stdout, "Response from `BucketsApi.UpdateBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**bucketId** | **string** | A storage bucket ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **storageUpdateBucketRequest** | [**StorageUpdateBucketRequest**](StorageUpdateBucketRequest.md) |  | 

### Return type

[**StorageUpdateBucketResponse**](storageUpdateBucketResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

