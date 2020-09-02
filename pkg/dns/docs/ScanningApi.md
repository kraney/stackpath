# \ScanningApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDiscoveryProviderDetails**](ScanningApi.md#GetDiscoveryProviderDetails) | **Get** /dns/v1/discovery/{domain}/provider_details | Get provider details
[**ScanDomainForRecords**](ScanningApi.md#ScanDomainForRecords) | **Post** /dns/v1/discovery/{domain}/records | Get resource records



## GetDiscoveryProviderDetails

> ZoneGetDiscoveryProviderDetailsResponse GetDiscoveryProviderDetails(ctx, domain)

Get provider details

Scan a domain for DNS provider information

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string**| A hostname to scan for provider information | 

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

> ZoneScanDomainForRecordsResponse ScanDomainForRecords(ctx, domain, scanDomainForRecordsRequestProviderConfig)

Get resource records

Scan a domain for resource records. This call returns the records that StackPath is able to scan at the time of execution. It performs a best effort, but cannot guarantee all resource records were found.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string**| A hostname to scan for resource records | 
**scanDomainForRecordsRequestProviderConfig** | [**ScanDomainForRecordsRequestProviderConfig**](ScanDomainForRecordsRequestProviderConfig.md)|  | 

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

