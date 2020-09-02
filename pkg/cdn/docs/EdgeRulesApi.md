# \EdgeRulesApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateScopeRule**](EdgeRulesApi.md#CreateScopeRule) | **Post** /cdn/v1/stacks/{stack_id}/sites/{site_id}/scopes/{scope_id}/rules | Create an EdgeRule
[**DeleteScopeRule**](EdgeRulesApi.md#DeleteScopeRule) | **Delete** /cdn/v1/stacks/{stack_id}/sites/{site_id}/scopes/{scope_id}/rules/{rule_id} | Delete an EdgeRule
[**GetScopeRule**](EdgeRulesApi.md#GetScopeRule) | **Get** /cdn/v1/stacks/{stack_id}/sites/{site_id}/scopes/{scope_id}/rules/{rule_id} | Get an EdgeRule
[**GetScopeRuleConfiguration**](EdgeRulesApi.md#GetScopeRuleConfiguration) | **Get** /cdn/v1/stacks/{stack_id}/sites/{site_id}/scopes/{scope_id}/rules/{rule_id}/configuration | Get an EdgeRule&#39;s configuration
[**GetScopeRules**](EdgeRulesApi.md#GetScopeRules) | **Get** /cdn/v1/stacks/{stack_id}/sites/{site_id}/scopes/{scope_id}/rules | Get all EdgeRules
[**UpdateScopeRuleConfiguration**](EdgeRulesApi.md#UpdateScopeRuleConfiguration) | **Patch** /cdn/v1/stacks/{stack_id}/sites/{site_id}/scopes/{scope_id}/rules/{rule_id}/configuration | Update an EdgeRule&#39;s configuration



## CreateScopeRule

> CdnCreateScopeRuleResponse CreateScopeRule(ctx, stackId, siteId, scopeId, cdnCreateScopeRuleRequest)

Create an EdgeRule

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
**scopeId** | **string**| A scope ID | 
**cdnCreateScopeRuleRequest** | [**CdnCreateScopeRuleRequest**](CdnCreateScopeRuleRequest.md)|  | 

### Return type

[**CdnCreateScopeRuleResponse**](cdnCreateScopeRuleResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteScopeRule

> DeleteScopeRule(ctx, stackId, siteId, scopeId, ruleId)

Delete an EdgeRule

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
**scopeId** | **string**| A scope ID | 
**ruleId** | **string**| An EdgeRule ID | 

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


## GetScopeRule

> CdnGetScopeRuleResponse GetScopeRule(ctx, stackId, siteId, scopeId, ruleId)

Get an EdgeRule

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
**scopeId** | **string**| A scope ID | 
**ruleId** | **string**| An EdgeRule ID | 

### Return type

[**CdnGetScopeRuleResponse**](cdnGetScopeRuleResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetScopeRuleConfiguration

> CdnGetScopeRuleConfigurationResponse GetScopeRuleConfiguration(ctx, stackId, siteId, scopeId, ruleId)

Get an EdgeRule's configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
**scopeId** | **string**| A scope ID | 
**ruleId** | **string**| An EdgeRule ID | 

### Return type

[**CdnGetScopeRuleConfigurationResponse**](cdnGetScopeRuleConfigurationResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetScopeRules

> CdnGetScopeRulesResponse GetScopeRules(ctx, stackId, siteId, scopeId, optional)

Get all EdgeRules

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
**scopeId** | **string**| A scope ID | 
 **optional** | ***GetScopeRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetScopeRulesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pageRequestFirst** | **optional.String**| The number of items desired. | 
 **pageRequestAfter** | **optional.String**| The cursor value after which data will be returned. | 
 **pageRequestFilter** | **optional.String**| SQL-style constraint filters. | 
 **pageRequestSortBy** | **optional.String**| Sort the response by the given field. | 

### Return type

[**CdnGetScopeRulesResponse**](cdnGetScopeRulesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateScopeRuleConfiguration

> CdnUpdateScopeRuleConfigurationResponse UpdateScopeRuleConfiguration(ctx, stackId, siteId, scopeId, ruleId, cdnUpdateScopeRuleConfigurationRequest)

Update an EdgeRule's configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
**scopeId** | **string**| A scope ID | 
**ruleId** | **string**| An EdgeRule ID | 
**cdnUpdateScopeRuleConfigurationRequest** | [**CdnUpdateScopeRuleConfigurationRequest**](CdnUpdateScopeRuleConfigurationRequest.md)|  | 

### Return type

[**CdnUpdateScopeRuleConfigurationResponse**](cdnUpdateScopeRuleConfigurationResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

