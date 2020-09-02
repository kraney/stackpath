# \SiteFeaturesApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisableSiteCDN**](SiteFeaturesApi.md#DisableSiteCDN) | **Delete** /delivery/v1/stacks/{stack_id}/sites/{site_id}/cdn | Disable CDN
[**DisableSiteEdgeEngine2**](SiteFeaturesApi.md#DisableSiteEdgeEngine2) | **Delete** /delivery/v1/stacks/{stack_id}/sites/{site_id}/scripting | Disable serverless scripting
[**DisableSiteWAF**](SiteFeaturesApi.md#DisableSiteWAF) | **Delete** /delivery/v1/stacks/{stack_id}/sites/{site_id}/waf | Disable WAF
[**EnableSiteCDN**](SiteFeaturesApi.md#EnableSiteCDN) | **Post** /delivery/v1/stacks/{stack_id}/sites/{site_id}/cdn | Enable CDN
[**EnableSiteEdgeEngine2**](SiteFeaturesApi.md#EnableSiteEdgeEngine2) | **Post** /delivery/v1/stacks/{stack_id}/sites/{site_id}/scripting | Enable serverless scripting
[**EnableSiteWAF**](SiteFeaturesApi.md#EnableSiteWAF) | **Post** /delivery/v1/stacks/{stack_id}/sites/{site_id}/waf | Enable WAF



## DisableSiteCDN

> DisableSiteCDN(ctx, stackId, siteId)

Disable CDN

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 

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


## DisableSiteEdgeEngine2

> DisableSiteEdgeEngine2(ctx, stackId, siteId)

Disable serverless scripting

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 

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


## DisableSiteWAF

> DisableSiteWAF(ctx, stackId, siteId)

Disable WAF

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 

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


## EnableSiteCDN

> EnableSiteCDN(ctx, stackId, siteId)

Enable CDN

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 

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


## EnableSiteEdgeEngine2

> EnableSiteEdgeEngine2(ctx, stackId, siteId)

Enable serverless scripting

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 

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


## EnableSiteWAF

> EnableSiteWAF(ctx, stackId, siteId, deliveryEnableSiteWafRequest)

Enable WAF

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
**deliveryEnableSiteWafRequest** | [**DeliveryEnableSiteWafRequest**](DeliveryEnableSiteWafRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

