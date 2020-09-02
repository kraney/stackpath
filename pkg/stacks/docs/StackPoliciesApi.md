# \StackPoliciesApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIAMPolicy**](StackPoliciesApi.md#DeleteIAMPolicy) | **Delete** /stack/v1/stacks/{stack_id}/iam/policy | Delete a stack&#39;s IAM policy
[**GetIAMPolicy**](StackPoliciesApi.md#GetIAMPolicy) | **Get** /stack/v1/stacks/{stack_id}/iam/policy | Get a stack&#39;s IAM policy
[**SetIAMPolicy**](StackPoliciesApi.md#SetIAMPolicy) | **Put** /stack/v1/stacks/{stack_id}/iam/policy | Set a stack&#39;s IAM policy
[**TestIAMPermissions**](StackPoliciesApi.md#TestIAMPermissions) | **Post** /stack/v1/stacks/{stack_id}/iam/test | Test stack policies



## DeleteIAMPolicy

> DeleteIAMPolicy(ctx, stackId)

Delete a stack's IAM policy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 

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


## GetIAMPolicy

> StackGetIamPolicyResponse GetIAMPolicy(ctx, stackId)

Get a stack's IAM policy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 

### Return type

[**StackGetIamPolicyResponse**](stackGetIAMPolicyResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetIAMPolicy

> StackSetIamPolicyResponse SetIAMPolicy(ctx, stackId, stackSetIamPolicyRequest)

Set a stack's IAM policy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**stackSetIamPolicyRequest** | [**StackSetIamPolicyRequest**](StackSetIamPolicyRequest.md)|  | 

### Return type

[**StackSetIamPolicyResponse**](stackSetIAMPolicyResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestIAMPermissions

> StackTestIamPermissionsResponse TestIAMPermissions(ctx, stackId, stackTestIamPermissionsRequest)

Test stack policies

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**stackTestIamPermissionsRequest** | [**StackTestIamPermissionsRequest**](StackTestIamPermissionsRequest.md)|  | 

### Return type

[**StackTestIamPermissionsResponse**](stackTestIAMPermissionsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

