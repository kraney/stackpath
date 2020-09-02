# \AccountPoliciesApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIAMPolicy**](AccountPoliciesApi.md#GetIAMPolicy) | **Get** /identity/v1/accounts/{account_id}/iam/policy | Get all account policies
[**SetIAMPolicy**](AccountPoliciesApi.md#SetIAMPolicy) | **Put** /identity/v1/accounts/{account_id}/iam/policy | Update account policies
[**TestIAMPermissions**](AccountPoliciesApi.md#TestIAMPermissions) | **Post** /identity/v1/accounts/{account_id}/iam/test | Test account policies



## GetIAMPolicy

> IdentityGetIamPolicyResponse GetIAMPolicy(ctx, accountId)

Get all account policies

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| An account ID | 

### Return type

[**IdentityGetIamPolicyResponse**](identityGetIAMPolicyResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetIAMPolicy

> IdentitySetIamPolicyResponse SetIAMPolicy(ctx, accountId, identitySetIamPolicyRequest)

Update account policies

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| An account ID | 
**identitySetIamPolicyRequest** | [**IdentitySetIamPolicyRequest**](IdentitySetIamPolicyRequest.md)|  | 

### Return type

[**IdentitySetIamPolicyResponse**](identitySetIAMPolicyResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestIAMPermissions

> IdentityTestIamPermissionsResponse TestIAMPermissions(ctx, accountId, identityTestIamPermissionsRequest)

Test account policies

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| An account ID | 
**identityTestIamPermissionsRequest** | [**IdentityTestIamPermissionsRequest**](IdentityTestIamPermissionsRequest.md)|  | 

### Return type

[**IdentityTestIamPermissionsResponse**](identityTestIAMPermissionsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

