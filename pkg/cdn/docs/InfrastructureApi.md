# \InfrastructureApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCDNIPs**](InfrastructureApi.md#GetCDNIPs) | **Get** /cdn/v1/ips | Get IP addresses
[**GetClosestPops**](InfrastructureApi.md#GetClosestPops) | **Get** /cdn/v1/pops/closest | Get POP performance
[**GetPops**](InfrastructureApi.md#GetPops) | **Get** /cdn/v1/pops | Get points of presence
[**ScanOrigin**](InfrastructureApi.md#ScanOrigin) | **Post** /cdn/v1/origins/scan | Scan an origin



## GetCDNIPs

> CdnGetCdniPsResponse GetCDNIPs(ctx, optional)

Get IP addresses

Retrieve all IP addresses used by the StackPath CDN

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetCDNIPsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCDNIPsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| Whether to search for IPv4, IPv6, or all IP addresses | [default to ALL]
 **responseType** | **optional.String**| The format to return the result in | [default to JSON]

### Return type

[**CdnGetCdniPsResponse**](cdnGetCDNIPsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClosestPops

> CdnGetClosestPopsResponse GetClosestPops(ctx, optional)

Get POP performance

Scan a URL from the StackPath edge network and return a performance report. Results are ordered with the fastest POP response first.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetClosestPopsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetClosestPopsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **url** | **optional.String**| The URL to scan. | 

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

> CdnGetPopsResponse GetPops(ctx, )

Get points of presence

Get the StackPath CDN's points of presence

### Required Parameters

This endpoint does not need any parameter.

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

> CdnScanOriginResponse ScanOrigin(ctx, cdnScanOriginRequest)

Scan an origin

Scan an origin from StackPath's CDN. Retrieve information regarding the origin, such as its IP address and whether or not it supports SSL.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cdnScanOriginRequest** | [**CdnScanOriginRequest**](CdnScanOriginRequest.md)|  | 

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

