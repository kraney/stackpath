# \AccountPoliciesApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIAMPolicy**](AccountPoliciesApi.md#GetIAMPolicy) | **Get** /identity/v1/accounts/{account_id}/iam/policy | Get all account policies
[**SetIAMPolicy**](AccountPoliciesApi.md#SetIAMPolicy) | **Put** /identity/v1/accounts/{account_id}/iam/policy | Update account policies
[**TestIAMPermissions**](AccountPoliciesApi.md#TestIAMPermissions) | **Post** /identity/v1/accounts/{account_id}/iam/test | Test account policies



## GetIAMPolicy

> IdentityGetIAMPolicyResponse GetIAMPolicy(ctx, accountId).Execute()

Get all account policies

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := "accountId_example" // string | An account ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountPoliciesApi.GetIAMPolicy(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountPoliciesApi.GetIAMPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIAMPolicy`: IdentityGetIAMPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountPoliciesApi.GetIAMPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | An account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIAMPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IdentityGetIAMPolicyResponse**](identityGetIAMPolicyResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetIAMPolicy

> IdentitySetIAMPolicyResponse SetIAMPolicy(ctx, accountId).IdentitySetIAMPolicyRequest(identitySetIAMPolicyRequest).Execute()

Update account policies

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := "accountId_example" // string | An account ID
    identitySetIAMPolicyRequest := openapiclient.identitySetIAMPolicyRequest{Policy: openapiclient.iamPolicy{Bindings: []PolicyBinding{openapiclient.PolicyBinding{Role: "Role_example", Members: []string{"Members_example")}), Version: int64(123), CreatedAt: "TODO", UpdatedAt: "TODO"}} // IdentitySetIAMPolicyRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountPoliciesApi.SetIAMPolicy(context.Background(), accountId, identitySetIAMPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountPoliciesApi.SetIAMPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetIAMPolicy`: IdentitySetIAMPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountPoliciesApi.SetIAMPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | An account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetIAMPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **identitySetIAMPolicyRequest** | [**IdentitySetIAMPolicyRequest**](IdentitySetIAMPolicyRequest.md) |  | 

### Return type

[**IdentitySetIAMPolicyResponse**](identitySetIAMPolicyResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestIAMPermissions

> IdentityTestIAMPermissionsResponse TestIAMPermissions(ctx, accountId).IdentityTestIAMPermissionsRequest(identityTestIAMPermissionsRequest).Execute()

Test account policies

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := "accountId_example" // string | An account ID
    identityTestIAMPermissionsRequest := openapiclient.identityTestIAMPermissionsRequest{Permissions: []string{"Permissions_example")} // IdentityTestIAMPermissionsRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountPoliciesApi.TestIAMPermissions(context.Background(), accountId, identityTestIAMPermissionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountPoliciesApi.TestIAMPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestIAMPermissions`: IdentityTestIAMPermissionsResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountPoliciesApi.TestIAMPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | An account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestIAMPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **identityTestIAMPermissionsRequest** | [**IdentityTestIAMPermissionsRequest**](IdentityTestIAMPermissionsRequest.md) |  | 

### Return type

[**IdentityTestIAMPermissionsResponse**](identityTestIAMPermissionsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

