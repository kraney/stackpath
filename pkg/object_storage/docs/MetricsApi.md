# \MetricsApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBucketMetrics**](MetricsApi.md#GetBucketMetrics) | **Get** /storage/v1/stacks/{stack_id}/buckets/{bucket_id}/metrics | Get bucket metrics
[**GetStackMetrics**](MetricsApi.md#GetStackMetrics) | **Get** /storage/v1/stacks/{stack_id}/metrics | Get stack metrics



## GetBucketMetrics

> PrometheusMetrics GetBucketMetrics(ctx, stackId, bucketId).StartTime(startTime).EndTime(endTime).Execute()

Get bucket metrics



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
    startTime := Get-Date // time.Time | The start date for the range of metrics. (optional)
    endTime := Get-Date // time.Time | The end date for the range of metrics. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MetricsApi.GetBucketMetrics(context.Background(), stackId, bucketId).StartTime(startTime).EndTime(endTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.GetBucketMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBucketMetrics`: PrometheusMetrics
    fmt.Fprintf(os.Stdout, "Response from `MetricsApi.GetBucketMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**bucketId** | **string** | A storage bucket ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBucketMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startTime** | **time.Time** | The start date for the range of metrics. | 
 **endTime** | **time.Time** | The end date for the range of metrics. | 

### Return type

[**PrometheusMetrics**](prometheusMetrics.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStackMetrics

> PrometheusMetrics GetStackMetrics(ctx, stackId).StartTime(startTime).EndTime(endTime).Execute()

Get stack metrics



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
    startTime := Get-Date // time.Time | The start date for the range of metrics. (optional)
    endTime := Get-Date // time.Time | The end date for the range of metrics. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MetricsApi.GetStackMetrics(context.Background(), stackId).StartTime(startTime).EndTime(endTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.GetStackMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStackMetrics`: PrometheusMetrics
    fmt.Fprintf(os.Stdout, "Response from `MetricsApi.GetStackMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStackMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startTime** | **time.Time** | The start date for the range of metrics. | 
 **endTime** | **time.Time** | The end date for the range of metrics. | 

### Return type

[**PrometheusMetrics**](prometheusMetrics.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

