# \InfrastructureApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCDNIPs**](InfrastructureApi.md#GetCDNIPs) | **Get** /cdn/v1/ips | Get IP addresses
[**GetClosestPops**](InfrastructureApi.md#GetClosestPops) | **Get** /cdn/v1/pops/closest | Get POP performance
[**GetPops**](InfrastructureApi.md#GetPops) | **Get** /cdn/v1/pops | Get points of presence
[**ScanOrigin**](InfrastructureApi.md#ScanOrigin) | **Post** /cdn/v1/origins/scan | Scan an origin



## GetCDNIPs

> CdnGetCDNIPsResponse GetCDNIPs(ctx).Filter(filter).ResponseType(responseType).Execute()

Get IP addresses



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
    filter := "filter_example" // string | Whether to search for IPv4, IPv6, or all IP addresses (optional) (default to "ALL")
    responseType := "responseType_example" // string | The format to return the result in (optional) (default to "JSON")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InfrastructureApi.GetCDNIPs(context.Background(), ).Filter(filter).ResponseType(responseType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfrastructureApi.GetCDNIPs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCDNIPs`: CdnGetCDNIPsResponse
    fmt.Fprintf(os.Stdout, "Response from `InfrastructureApi.GetCDNIPs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCDNIPsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Whether to search for IPv4, IPv6, or all IP addresses | [default to &quot;ALL&quot;]
 **responseType** | **string** | The format to return the result in | [default to &quot;JSON&quot;]

### Return type

[**CdnGetCDNIPsResponse**](cdnGetCDNIPsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClosestPops

> CdnGetClosestPopsResponse GetClosestPops(ctx).Url(url).Execute()

Get POP performance



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
    url := "url_example" // string | The URL to scan. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InfrastructureApi.GetClosestPops(context.Background(), ).Url(url).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfrastructureApi.GetClosestPops``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClosestPops`: CdnGetClosestPopsResponse
    fmt.Fprintf(os.Stdout, "Response from `InfrastructureApi.GetClosestPops`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetClosestPopsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **url** | **string** | The URL to scan. | 

### Return type

[**CdnGetClosestPopsResponse**](cdnGetClosestPopsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPops

> CdnGetPopsResponse GetPops(ctx).Execute()

Get points of presence



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InfrastructureApi.GetPops(context.Background(), ).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfrastructureApi.GetPops``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPops`: CdnGetPopsResponse
    fmt.Fprintf(os.Stdout, "Response from `InfrastructureApi.GetPops`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPopsRequest struct via the builder pattern


### Return type

[**CdnGetPopsResponse**](cdnGetPopsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScanOrigin

> CdnScanOriginResponse ScanOrigin(ctx).CdnScanOriginRequest(cdnScanOriginRequest).Execute()

Scan an origin



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
    cdnScanOriginRequest := openapiclient.cdnScanOriginRequest{Domain: "Domain_example"} // CdnScanOriginRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InfrastructureApi.ScanOrigin(context.Background(), cdnScanOriginRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfrastructureApi.ScanOrigin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScanOrigin`: CdnScanOriginResponse
    fmt.Fprintf(os.Stdout, "Response from `InfrastructureApi.ScanOrigin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiScanOriginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cdnScanOriginRequest** | [**CdnScanOriginRequest**](CdnScanOriginRequest.md) |  | 

### Return type

[**CdnScanOriginResponse**](cdnScanOriginResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

