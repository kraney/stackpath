# \PoliciesApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisablePolicy**](PoliciesApi.md#DisablePolicy) | **Post** /waf/v1/stacks/{stack_id}/sites/{site_id}/policy_groups/{policy_group_id}/policies/{policy_id}/disable | Disable a policy
[**DisablePolicyGroup**](PoliciesApi.md#DisablePolicyGroup) | **Post** /waf/v1/stacks/{stack_id}/sites/{site_id}/policy_groups/{policy_group_id}/disable | Disable all policies in a group
[**EnablePolicy**](PoliciesApi.md#EnablePolicy) | **Post** /waf/v1/stacks/{stack_id}/sites/{site_id}/policy_groups/{policy_group_id}/policies/{policy_id}/enable | Enable a policy
[**EnablePolicyGroup**](PoliciesApi.md#EnablePolicyGroup) | **Post** /waf/v1/stacks/{stack_id}/sites/{site_id}/policy_groups/{policy_group_id}/enable | Enable all policies in a group
[**GetPolicies**](PoliciesApi.md#GetPolicies) | **Get** /waf/v1/stacks/{stack_id}/sites/{site_id}/policy_groups/{policy_group_id}/policies | Get all policies in a group
[**GetPolicy**](PoliciesApi.md#GetPolicy) | **Get** /waf/v1/stacks/{stack_id}/sites/{site_id}/policy_groups/{policy_group_id}/policies/{policy_id} | Get a policy
[**GetPolicyGroup**](PoliciesApi.md#GetPolicyGroup) | **Get** /waf/v1/stacks/{stack_id}/sites/{site_id}/policy_groups/{policy_group_id} | Get a policy group
[**GetPolicyGroups**](PoliciesApi.md#GetPolicyGroups) | **Get** /waf/v1/stacks/{stack_id}/sites/{site_id}/policy_groups | Get all policy groups
[**UpdatePolicyGroups**](PoliciesApi.md#UpdatePolicyGroups) | **Patch** /waf/v1/stacks/{stack_id}/sites/{site_id}/policy_groups | Update policy groups



## DisablePolicy

> DisablePolicy(ctx, stackId, siteId, policyGroupId, policyId)

Disable a policy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
**policyGroupId** | **string**| A WAF policy group ID | 
**policyId** | **string**| A WAF policy ID | 

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


## DisablePolicyGroup

> DisablePolicyGroup(ctx, stackId, siteId, policyGroupId)

Disable all policies in a group

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
**policyGroupId** | **string**| A WAF policy group ID | 

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


## EnablePolicy

> EnablePolicy(ctx, stackId, siteId, policyGroupId, policyId)

Enable a policy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
**policyGroupId** | **string**| A WAF policy group ID | 
**policyId** | **string**| A WAF policy ID | 

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


## EnablePolicyGroup

> EnablePolicyGroup(ctx, stackId, siteId, policyGroupId)

Enable all policies in a group

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
**policyGroupId** | **string**| A WAF policy group ID | 

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


## GetPolicies

> WafGetPoliciesResponse GetPolicies(ctx, stackId, siteId, policyGroupId)

Get all policies in a group

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
**policyGroupId** | **string**| A WAF policy group ID | 

### Return type

[**WafGetPoliciesResponse**](wafGetPoliciesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicy

> WafGetPolicyResponse GetPolicy(ctx, stackId, siteId, policyGroupId, policyId)

Get a policy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
**policyGroupId** | **string**| A WAF policy group ID | 
**policyId** | **string**| A WAF policy ID | 

### Return type

[**WafGetPolicyResponse**](wafGetPolicyResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicyGroup

> WafGetPolicyGroupResponse GetPolicyGroup(ctx, stackId, siteId, policyGroupId)

Get a policy group

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
**policyGroupId** | **string**| A WAF policy group ID | 

### Return type

[**WafGetPolicyGroupResponse**](wafGetPolicyGroupResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicyGroups

> WafGetPolicyGroupsResponse GetPolicyGroups(ctx, stackId, siteId)

Get all policy groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 

### Return type

[**WafGetPolicyGroupsResponse**](wafGetPolicyGroupsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePolicyGroups

> WafGetPolicyGroupsResponse UpdatePolicyGroups(ctx, stackId, siteId, wafUpdatePolicyGroupsRequest)

Update policy groups

Provide the IDs and enabled states of the policies in their policy groups that should change

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
**wafUpdatePolicyGroupsRequest** | [**WafUpdatePolicyGroupsRequest**](WafUpdatePolicyGroupsRequest.md)|  | 

### Return type

[**WafGetPolicyGroupsResponse**](wafGetPolicyGroupsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

