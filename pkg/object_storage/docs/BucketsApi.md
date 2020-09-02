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

> StorageCreateBucketResponse CreateBucket(ctx, stackId, storageCreateBucketRequest)

Create a bucket

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**storageCreateBucketRequest** | [**StorageCreateBucketRequest**](StorageCreateBucketRequest.md)|  | 

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

> DeleteBucket(ctx, stackId, bucketId, optional)

Delete a bucket

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**bucketId** | **string**| A storage bucket ID | 
 **optional** | ***DeleteBucketOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteBucketOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **forceDelete** | **optional.Bool**| Force bucket deletion even if there are contents inside it | 

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

> StorageGetBucketResponse GetBucket(ctx, stackId, bucketId)

Get a bucket

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**bucketId** | **string**| A storage bucket ID | 

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

> StorageGetBucketsResponse GetBuckets(ctx, stackId, optional)

Get all buckets

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
 **optional** | ***GetBucketsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetBucketsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageRequestFirst** | **optional.String**| The number of items desired. | 
 **pageRequestAfter** | **optional.String**| The cursor value after which data will be returned. | 
 **pageRequestFilter** | **optional.String**| SQL-style constraint filters. | 
 **pageRequestSortBy** | **optional.String**| Sort the response by the given field. | 

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

> StorageUpdateBucketResponse UpdateBucket(ctx, stackId, bucketId, storageUpdateBucketRequest)

Update a bucket

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**bucketId** | **string**| A storage bucket ID | 
**storageUpdateBucketRequest** | [**StorageUpdateBucketRequest**](StorageUpdateBucketRequest.md)|  | 

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

