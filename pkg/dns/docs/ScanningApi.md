# \ScanningApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDiscoveryProviderDetails**](ScanningApi.md#GetDiscoveryProviderDetails) | **Get** /dns/v1/discovery/{domain}/provider_details | Get provider details
[**ScanDomainForRecords**](ScanningApi.md#ScanDomainForRecords) | **Post** /dns/v1/discovery/{domain}/records | Get resource records



## GetDiscoveryProviderDetails

> ZoneGetDiscoveryProviderDetailsResponse GetDiscoveryProviderDetails(ctx, domain).Execute()

Get provider details



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
    domain := "domain_example" // string | A hostname to scan for provider information

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScanningApi.GetDiscoveryProviderDetails(context.Background(), domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScanningApi.GetDiscoveryProviderDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDiscoveryProviderDetails`: ZoneGetDiscoveryProviderDetailsResponse
    fmt.Fprintf(os.Stdout, "Response from `ScanningApi.GetDiscoveryProviderDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | A hostname to scan for provider information | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDiscoveryProviderDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ZoneGetDiscoveryProviderDetailsResponse**](zoneGetDiscoveryProviderDetailsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScanDomainForRecords

> ZoneScanDomainForRecordsResponse ScanDomainForRecords(ctx, domain).ScanDomainForRecordsRequestProviderConfig(scanDomainForRecordsRequestProviderConfig).Execute()

Get resource records



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
    domain := "domain_example" // string | A hostname to scan for resource records
    scanDomainForRecordsRequestProviderConfig := openapiclient.ScanDomainForRecordsRequestProviderConfig{DnsProvider: openapiclient.zoneDnsProvider{}, AuthenticationUser: "AuthenticationUser_example", ApiKey: "ApiKey_example"} // ScanDomainForRecordsRequestProviderConfig | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScanningApi.ScanDomainForRecords(context.Background(), domain, scanDomainForRecordsRequestProviderConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScanningApi.ScanDomainForRecords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScanDomainForRecords`: ZoneScanDomainForRecordsResponse
    fmt.Fprintf(os.Stdout, "Response from `ScanningApi.ScanDomainForRecords`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | A hostname to scan for resource records | 

### Other Parameters

Other parameters are passed through a pointer to a apiScanDomainForRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scanDomainForRecordsRequestProviderConfig** | [**ScanDomainForRecordsRequestProviderConfig**](ScanDomainForRecordsRequestProviderConfig.md) |  | 

### Return type

[**ZoneScanDomainForRecordsResponse**](zoneScanDomainForRecordsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

