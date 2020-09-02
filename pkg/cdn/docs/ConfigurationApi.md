# \ConfigurationApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConnectScopeToOrigin**](ConfigurationApi.md#ConnectScopeToOrigin) | **Post** /cdn/v1/stacks/{stack_id}/sites/{site_id}/scopes/{scope_id}/origins | Connect an origin to a scope
[**CreateScope**](ConfigurationApi.md#CreateScope) | **Post** /cdn/v1/stacks/{stack_id}/sites/{site_id}/scopes | Create a scope
[**CreateScopeHostname**](ConfigurationApi.md#CreateScopeHostname) | **Post** /cdn/v1/stacks/{stack_id}/sites/{site_id}/scopes/{scope_id}/hostnames | Add a scope hostname
[**DeleteScope**](ConfigurationApi.md#DeleteScope) | **Delete** /cdn/v1/stacks/{stack_id}/sites/{site_id}/scopes/{scope_id} | Delete a scope
[**DeleteScopeHostname**](ConfigurationApi.md#DeleteScopeHostname) | **Delete** /cdn/v1/stacks/{stack_id}/sites/{site_id}/scopes/{scope_id}/hostnames/{domain} | Delete a scope hostname
[**DisconnectScopeFromOrigin**](ConfigurationApi.md#DisconnectScopeFromOrigin) | **Delete** /cdn/v1/stacks/{stack_id}/sites/{site_id}/scopes/{scope_id}/origins/{origin_id} | Delete a scope origin
[**GetScopeConfiguration**](ConfigurationApi.md#GetScopeConfiguration) | **Get** /cdn/v1/stacks/{stack_id}/sites/{site_id}/scopes/{scope_id}/configuration | Get a scope&#39;s configuraiton
[**GetScopeHostnames**](ConfigurationApi.md#GetScopeHostnames) | **Get** /cdn/v1/stacks/{stack_id}/sites/{site_id}/scopes/{scope_id}/hostnames | Get all scope hostnames
[**GetScopeOrigins**](ConfigurationApi.md#GetScopeOrigins) | **Get** /cdn/v1/stacks/{stack_id}/sites/{site_id}/scopes/{scope_id}/origins | Get all scope origins
[**GetSiteDnsTargets**](ConfigurationApi.md#GetSiteDnsTargets) | **Get** /cdn/v1/stacks/{stack_id}/sites/{site_id}/dns/targets | Get CNAME targets
[**GetSiteScopes**](ConfigurationApi.md#GetSiteScopes) | **Get** /cdn/v1/stacks/{stack_id}/sites/{site_id}/scopes | Get all scopes
[**UpdateScopeConfiguration**](ConfigurationApi.md#UpdateScopeConfiguration) | **Patch** /cdn/v1/stacks/{stack_id}/sites/{site_id}/scopes/{scope_id}/configuration | Update a scope&#39;s configuration



## ConnectScopeToOrigin

> CdnConnectScopeToOriginResponse ConnectScopeToOrigin(ctx, stackId, siteId, scopeId, cdnConnectScopeToOriginRequest)

Connect an origin to a scope

The origin is automatically created if necessary. When the request contains a priority which an origin already associated with the scope has set, the existing origin is disconnected. The priority of an origin already associated with a scope can be modified via this endpoint.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
**scopeId** | **string**| A scope ID | 
**cdnConnectScopeToOriginRequest** | [**CdnConnectScopeToOriginRequest**](CdnConnectScopeToOriginRequest.md)|  | 

### Return type

[**CdnConnectScopeToOriginResponse**](cdnConnectScopeToOriginResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateScope

> CdnCreateScopeResponse CreateScope(ctx, stackId, siteId, cdnCreateScopeRequest)

Create a scope

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
**cdnCreateScopeRequest** | [**CdnCreateScopeRequest**](CdnCreateScopeRequest.md)|  | 

### Return type

[**CdnCreateScopeResponse**](cdnCreateScopeResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateScopeHostname

> CdnCreateScopeHostnameResponse CreateScopeHostname(ctx, stackId, siteId, scopeId, cdnCreateScopeHostnameRequest)

Add a scope hostname

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
**scopeId** | **string**| A scope ID | 
**cdnCreateScopeHostnameRequest** | [**CdnCreateScopeHostnameRequest**](CdnCreateScopeHostnameRequest.md)|  | 

### Return type

[**CdnCreateScopeHostnameResponse**](cdnCreateScopeHostnameResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteScope

> DeleteScope(ctx, stackId, siteId, scopeId)

Delete a scope

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
**scopeId** | **string**| A scope ID | 

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


## DeleteScopeHostname

> DeleteScopeHostname(ctx, stackId, siteId, scopeId, domain, optional)

Delete a scope hostname

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
**scopeId** | **string**| A scope ID | 
**domain** | **string**| The hostname to remove from a scope | 
 **optional** | ***DeleteScopeHostnameOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteScopeHostnameOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **disableTransparentMode** | **optional.Bool**| Whether or not to remove the hostname from a CDN site&#39;s CDN scope or its WAF scope. When true, this call removes the hostname from a CDN site&#39;s scope instead of loading from a CDN site&#39;s WAF scope, if the site has WAF service. | 

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


## DisconnectScopeFromOrigin

> DisconnectScopeFromOrigin(ctx, stackId, siteId, scopeId, originId)

Delete a scope origin

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
**scopeId** | **string**| A scope ID | 
**originId** | **string**| An origin ID | 

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


## GetScopeConfiguration

> CdnGetScopeConfigurationResponse GetScopeConfiguration(ctx, stackId, siteId, scopeId)

Get a scope's configuraiton

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
**scopeId** | **string**| A scope ID | 

### Return type

[**CdnGetScopeConfigurationResponse**](cdnGetScopeConfigurationResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetScopeHostnames

> CdnGetScopeHostnamesResponse GetScopeHostnames(ctx, stackId, siteId, scopeId, optional)

Get all scope hostnames

Hostnames allow the CDN to recognize an HTTP request and associate it with a CDN site.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
**scopeId** | **string**| A scope ID | 
 **optional** | ***GetScopeHostnamesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetScopeHostnamesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pageRequestFirst** | **optional.String**| The number of items desired. | 
 **pageRequestAfter** | **optional.String**| The cursor value after which data will be returned. | 
 **pageRequestFilter** | **optional.String**| SQL-style constraint filters. | 
 **pageRequestSortBy** | **optional.String**| Sort the response by the given field. | 
 **disableTransparentMode** | **optional.Bool**| Whether or not to load hostnames from a CDN site&#39;s CDN scope or its WAF scope. When true, this call loads scope hostnames from a CDN site&#39;s scope instead of loading from a CDN site&#39;s WAF scope, if the site has WAF service. | 

### Return type

[**CdnGetScopeHostnamesResponse**](cdnGetScopeHostnamesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetScopeOrigins

> CdnGetScopeOriginsResponse GetScopeOrigins(ctx, stackId, siteId, scopeId, optional)

Get all scope origins

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
**scopeId** | **string**| A scope ID | 
 **optional** | ***GetScopeOriginsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetScopeOriginsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pageRequestFirst** | **optional.String**| The number of items desired. | 
 **pageRequestAfter** | **optional.String**| The cursor value after which data will be returned. | 
 **pageRequestFilter** | **optional.String**| SQL-style constraint filters. | 
 **pageRequestSortBy** | **optional.String**| Sort the response by the given field. | 

### Return type

[**CdnGetScopeOriginsResponse**](cdnGetScopeOriginsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteDnsTargets

> CdnGetSiteDnsTargetsResponse GetSiteDnsTargets(ctx, stackId, siteId)

Get CNAME targets

A site's hostname should point to these CNAME targets in order for traffic to be sent through StackPath's edge nodes.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 

### Return type

[**CdnGetSiteDnsTargetsResponse**](cdnGetSiteDnsTargetsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteScopes

> CdnGetSiteScopesResponse GetSiteScopes(ctx, stackId, siteId, optional)

Get all scopes

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
 **optional** | ***GetSiteScopesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSiteScopesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageRequestFirst** | **optional.String**| The number of items desired. | 
 **pageRequestAfter** | **optional.String**| The cursor value after which data will be returned. | 
 **pageRequestFilter** | **optional.String**| SQL-style constraint filters. | 
 **pageRequestSortBy** | **optional.String**| Sort the response by the given field. | 
 **disableTransparentMode** | **optional.Bool**| Whether or not to retrieve the site&#39;s CDN scope or its WAF scope. When true, this call removes the hostname from a CDN site&#39;s scope instead of loading from a CDN site&#39;s WAF scope, if the site has WAF service. | 

### Return type

[**CdnGetSiteScopesResponse**](cdnGetSiteScopesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateScopeConfiguration

> CdnUpdateScopeConfigurationResponse UpdateScopeConfiguration(ctx, stackId, siteId, scopeId, cdnUpdateScopeConfigurationRequest)

Update a scope's configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
**scopeId** | **string**| A scope ID | 
**cdnUpdateScopeConfigurationRequest** | [**CdnUpdateScopeConfigurationRequest**](CdnUpdateScopeConfigurationRequest.md)|  | 

### Return type

[**CdnUpdateScopeConfigurationResponse**](cdnUpdateScopeConfigurationResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

