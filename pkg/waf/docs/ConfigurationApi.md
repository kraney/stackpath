# \ConfigurationApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSiteDnsTargets**](ConfigurationApi.md#GetSiteDnsTargets) | **Get** /waf/v1/stacks/{stack_id}/sites/{site_id}/delivery/dns/targets | Get CNAME targets



## GetSiteDnsTargets

> WafGetSiteDnsTargetsResponse GetSiteDnsTargets(ctx, stackId, siteId)

Get CNAME targets

A site's hostname should point to these CNAME targets in order for traffic to be sent through StackPath's edge nodes.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 

### Return type

[**WafGetSiteDnsTargetsResponse**](wafGetSiteDnsTargetsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

