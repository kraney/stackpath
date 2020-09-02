# \NetworkPoliciesApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkPolicy**](NetworkPoliciesApi.md#CreateNetworkPolicy) | **Post** /ipam/v1/stacks/{stack_id}/network_policies | Create a policy
[**DeleteNetworkPolicy**](NetworkPoliciesApi.md#DeleteNetworkPolicy) | **Delete** /ipam/v1/stacks/{stack_id}/network_policies/{network_policy_id} | Delete a policy
[**GetNetworkPolicies**](NetworkPoliciesApi.md#GetNetworkPolicies) | **Get** /ipam/v1/stacks/{stack_id}/network_policies | Get all policies
[**GetNetworkPolicy**](NetworkPoliciesApi.md#GetNetworkPolicy) | **Get** /ipam/v1/stacks/{stack_id}/network_policies/{network_policy_id} | Get a policy
[**UpdateNetworkPolicy**](NetworkPoliciesApi.md#UpdateNetworkPolicy) | **Put** /ipam/v1/stacks/{stack_id}/network_policies/{network_policy_id} | Update a policy



## CreateNetworkPolicy

> V1CreateNetworkPolicyResponse CreateNetworkPolicy(ctx, stackId, v1CreateNetworkPolicyRequest)

Create a policy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**v1CreateNetworkPolicyRequest** | [**V1CreateNetworkPolicyRequest**](V1CreateNetworkPolicyRequest.md)|  | 

### Return type

[**V1CreateNetworkPolicyResponse**](v1CreateNetworkPolicyResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkPolicy

> DeleteNetworkPolicy(ctx, stackId, networkPolicyId)

Delete a policy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**networkPolicyId** | **string**| An EdgeCompute network policy ID | 

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


## GetNetworkPolicies

> V1GetNetworkPoliciesResponse GetNetworkPolicies(ctx, stackId, optional)

Get all policies

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
 **optional** | ***GetNetworkPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetNetworkPoliciesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageRequestFirst** | **optional.String**| The number of items desired. | 
 **pageRequestAfter** | **optional.String**| The cursor value after which data will be returned. | 
 **pageRequestFilter** | **optional.String**| SQL-style constraint filters. | 
 **pageRequestSortBy** | **optional.String**| Sort the response by the given field. | 

### Return type

[**V1GetNetworkPoliciesResponse**](v1GetNetworkPoliciesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkPolicy

> V1GetNetworkPolicyResponse GetNetworkPolicy(ctx, stackId, networkPolicyId)

Get a policy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**networkPolicyId** | **string**| An EdgeCompute network policy ID | 

### Return type

[**V1GetNetworkPolicyResponse**](v1GetNetworkPolicyResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkPolicy

> V1UpdateNetworkPolicyResponse UpdateNetworkPolicy(ctx, stackId, networkPolicyId, v1UpdateNetworkPolicyRequest)

Update a policy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**networkPolicyId** | **string**| An EdgeCompute network policy ID | 
**v1UpdateNetworkPolicyRequest** | [**V1UpdateNetworkPolicyRequest**](V1UpdateNetworkPolicyRequest.md)|  | 

### Return type

[**V1UpdateNetworkPolicyResponse**](v1UpdateNetworkPolicyResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

