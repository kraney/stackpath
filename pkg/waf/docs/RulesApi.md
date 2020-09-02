# \RulesApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkDeleteRules**](RulesApi.md#BulkDeleteRules) | **Post** /waf/v1/stacks/{stack_id}/sites/{site_id}/rules/bulk_delete | Delete multiple rules
[**CreateRule**](RulesApi.md#CreateRule) | **Post** /waf/v1/stacks/{stack_id}/sites/{site_id}/rules | Create a rule
[**DeleteRule**](RulesApi.md#DeleteRule) | **Delete** /waf/v1/stacks/{stack_id}/sites/{site_id}/rules/{rule_id} | Delete a rule
[**DisableRule**](RulesApi.md#DisableRule) | **Post** /waf/v1/stacks/{stack_id}/sites/{site_id}/rules/{rule_id}/disable | Disable a rule
[**EnableRule**](RulesApi.md#EnableRule) | **Post** /waf/v1/stacks/{stack_id}/sites/{site_id}/rules/{rule_id}/enable | Enable a rule
[**GetRule**](RulesApi.md#GetRule) | **Get** /waf/v1/stacks/{stack_id}/sites/{site_id}/rules/{rule_id} | Get a rule
[**GetRules**](RulesApi.md#GetRules) | **Get** /waf/v1/stacks/{stack_id}/sites/{site_id}/rules | Get all rules
[**UpdateRule**](RulesApi.md#UpdateRule) | **Patch** /waf/v1/stacks/{stack_id}/sites/{site_id}/rules/{rule_id} | Update a rule



## BulkDeleteRules

> BulkDeleteRules(ctx, stackId, siteId, wafBulkDeleteRulesRequest)

Delete multiple rules

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
**wafBulkDeleteRulesRequest** | [**WafBulkDeleteRulesRequest**](WafBulkDeleteRulesRequest.md)|  | 

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


## CreateRule

> WafCreateRuleResponse CreateRule(ctx, stackId, siteId, wafCreateRuleRequest)

Create a rule

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
**wafCreateRuleRequest** | [**WafCreateRuleRequest**](WafCreateRuleRequest.md)|  | 

### Return type

[**WafCreateRuleResponse**](wafCreateRuleResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRule

> DeleteRule(ctx, stackId, siteId, ruleId)

Delete a rule

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
**ruleId** | **string**| A WAF rule ID | 

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


## DisableRule

> DisableRule(ctx, stackId, siteId, ruleId)

Disable a rule

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
**ruleId** | **string**| A WAF rule ID | 

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


## EnableRule

> EnableRule(ctx, stackId, siteId, ruleId)

Enable a rule

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
**ruleId** | **string**| A WAF rule ID | 

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


## GetRule

> WafGetRuleResponse GetRule(ctx, stackId, siteId, ruleId)

Get a rule

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
**ruleId** | **string**| A WAF rule ID | 

### Return type

[**WafGetRuleResponse**](wafGetRuleResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRules

> WafGetRulesResponse GetRules(ctx, stackId, siteId, optional)

Get all rules

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
 **optional** | ***GetRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetRulesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageRequestFirst** | **optional.String**| The number of items desired. | 
 **pageRequestAfter** | **optional.String**| The cursor value after which data will be returned. | 
 **pageRequestFilter** | **optional.String**| SQL-style constraint filters. | 
 **pageRequestSortBy** | **optional.String**| Sort the response by the given field. | 

### Return type

[**WafGetRulesResponse**](wafGetRulesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRule

> WafUpdateRuleResponse UpdateRule(ctx, stackId, siteId, ruleId, wafUpdateRuleRequest)

Update a rule

Only properties present in the request will be updated.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
**ruleId** | **string**| A WAF rule ID | 
**wafUpdateRuleRequest** | [**WafUpdateRuleRequest**](WafUpdateRuleRequest.md)|  | 

### Return type

[**WafUpdateRuleResponse**](wafUpdateRuleResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

