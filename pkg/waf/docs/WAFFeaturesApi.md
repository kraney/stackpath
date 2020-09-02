# \WAFFeaturesApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDdosSettings**](WAFFeaturesApi.md#GetDdosSettings) | **Get** /waf/v1/stacks/{stack_id}/sites/{site_id}/ddos | Get DDOS settings
[**GetTags**](WAFFeaturesApi.md#GetTags) | **Get** /waf/v1/tags | Get all tags
[**MonitorSite**](WAFFeaturesApi.md#MonitorSite) | **Post** /waf/v1/stacks/{stack_id}/sites/{site_id}/monitoring | Enable monitoring mode
[**UnMonitorSite**](WAFFeaturesApi.md#UnMonitorSite) | **Delete** /waf/v1/stacks/{stack_id}/sites/{site_id}/monitoring | Disable monitoring mode
[**UpdateDdosSettings**](WAFFeaturesApi.md#UpdateDdosSettings) | **Patch** /waf/v1/stacks/{stack_id}/sites/{site_id}/ddos | Update DDOS settings
[**UpdateSiteApiUrls**](WAFFeaturesApi.md#UpdateSiteApiUrls) | **Put** /waf/v1/stacks/{stack_id}/sites/{site_id}/api_urls | Update API URLs



## GetDdosSettings

> WafGetDdosSettingsResponse GetDdosSettings(ctx, stackId, siteId)

Get DDOS settings

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 

### Return type

[**WafGetDdosSettingsResponse**](wafGetDdosSettingsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTags

> WafGetTagsResponse GetTags(ctx, optional)

Get all tags

Get tags available for use in WAF rule conditons. Tags describe aspects of an incoming web request and acn be used to create complex application-level custom rules.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetTagsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTagsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageRequestFirst** | **optional.String**| The number of items desired. | 
 **pageRequestAfter** | **optional.String**| The cursor value after which data will be returned. | 
 **pageRequestFilter** | **optional.String**| SQL-style constraint filters. | 
 **pageRequestSortBy** | **optional.String**| Sort the response by the given field. | 

### Return type

[**WafGetTagsResponse**](wafGetTagsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorSite

> MonitorSite(ctx, stackId, siteId)

Enable monitoring mode

Sites in monitoring mode accept incoming requests and log but take no action on policy and rule violations. A sitge's WAF feature must be enabled for monitoring mode to apply. Otherwise, monitoring mode will take effect the next time the WAF feature is enabled.

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


## UnMonitorSite

> UnMonitorSite(ctx, stackId, siteId)

Disable monitoring mode

Restore the WAF feature's original enabled or disabled state.

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


## UpdateDdosSettings

> WafUpdateDdosSettingsResponse UpdateDdosSettings(ctx, stackId, siteId, wafUpdateDdosSettingsRequest)

Update DDOS settings

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
**wafUpdateDdosSettingsRequest** | [**WafUpdateDdosSettingsRequest**](WafUpdateDdosSettingsRequest.md)|  | 

### Return type

[**WafUpdateDdosSettingsResponse**](wafUpdateDdosSettingsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSiteApiUrls

> WafUpdateSiteApiUrlsResponse UpdateSiteApiUrls(ctx, stackId, siteId, wafUpdateSiteApiUrlsRequest)

Update API URLs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
**wafUpdateSiteApiUrlsRequest** | [**WafUpdateSiteApiUrlsRequest**](WafUpdateSiteApiUrlsRequest.md)|  | 

### Return type

[**WafUpdateSiteApiUrlsResponse**](wafUpdateSiteApiUrlsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

