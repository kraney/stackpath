# \InfrastructureApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLocations**](InfrastructureApi.md#GetLocations) | **Get** /workload/v1/locations | Get compute locations



## GetLocations

> V1GetLocationsResponse GetLocations(ctx).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).Execute()

Get compute locations



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
    pageRequestFirst := "pageRequestFirst_example" // string | The number of items desired. (optional)
    pageRequestAfter := "pageRequestAfter_example" // string | The cursor value after which data will be returned. (optional)
    pageRequestFilter := "pageRequestFilter_example" // string | SQL-style constraint filters. (optional)
    pageRequestSortBy := "pageRequestSortBy_example" // string | Sort the response by the given field. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InfrastructureApi.GetLocations(context.Background(), ).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfrastructureApi.GetLocations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLocations`: V1GetLocationsResponse
    fmt.Fprintf(os.Stdout, "Response from `InfrastructureApi.GetLocations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageRequestFirst** | **string** | The number of items desired. | 
 **pageRequestAfter** | **string** | The cursor value after which data will be returned. | 
 **pageRequestFilter** | **string** | SQL-style constraint filters. | 
 **pageRequestSortBy** | **string** | Sort the response by the given field. | 

### Return type

[**V1GetLocationsResponse**](v1GetLocationsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

