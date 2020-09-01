# \StackPoliciesApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIAMPolicy**](StackPoliciesApi.md#DeleteIAMPolicy) | **Delete** /stack/v1/stacks/{stack_id}/iam/policy | Delete a stack&#39;s IAM policy
[**GetIAMPolicy**](StackPoliciesApi.md#GetIAMPolicy) | **Get** /stack/v1/stacks/{stack_id}/iam/policy | Get a stack&#39;s IAM policy
[**SetIAMPolicy**](StackPoliciesApi.md#SetIAMPolicy) | **Put** /stack/v1/stacks/{stack_id}/iam/policy | Set a stack&#39;s IAM policy
[**TestIAMPermissions**](StackPoliciesApi.md#TestIAMPermissions) | **Post** /stack/v1/stacks/{stack_id}/iam/test | Test stack policies



## DeleteIAMPolicy

> DeleteIAMPolicy(ctx, stackId).Execute()

Delete a stack's IAM policy

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
    stackId := "stackId_example" // string | A stack ID or slug

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StackPoliciesApi.DeleteIAMPolicy(context.Background(), stackId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StackPoliciesApi.DeleteIAMPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIAMPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> StackGetIAMPolicyResponse GetIAMPolicy(ctx, stackId).Execute()

Get a stack's IAM policy

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
    stackId := "stackId_example" // string | A stack ID or slug

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StackPoliciesApi.GetIAMPolicy(context.Background(), stackId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StackPoliciesApi.GetIAMPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIAMPolicy`: StackGetIAMPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `StackPoliciesApi.GetIAMPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIAMPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StackGetIAMPolicyResponse**](stackGetIAMPolicyResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetIAMPolicy

> StackSetIAMPolicyResponse SetIAMPolicy(ctx, stackId).StackSetIAMPolicyRequest(stackSetIAMPolicyRequest).Execute()

Set a stack's IAM policy

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
    stackId := "stackId_example" // string | A stack ID or slug
    stackSetIAMPolicyRequest := openapiclient.stackSetIAMPolicyRequest{Policy: openapiclient.iamPolicy{Bindings: []PolicyBinding{openapiclient.PolicyBinding{Role: "Role_example", Members: []string{"Members_example")}), Version: int64(123), CreatedAt: "TODO", UpdatedAt: "TODO"}} // StackSetIAMPolicyRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StackPoliciesApi.SetIAMPolicy(context.Background(), stackId, stackSetIAMPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StackPoliciesApi.SetIAMPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetIAMPolicy`: StackSetIAMPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `StackPoliciesApi.SetIAMPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetIAMPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **stackSetIAMPolicyRequest** | [**StackSetIAMPolicyRequest**](StackSetIAMPolicyRequest.md) |  | 

### Return type

[**StackSetIAMPolicyResponse**](stackSetIAMPolicyResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestIAMPermissions

> StackTestIAMPermissionsResponse TestIAMPermissions(ctx, stackId).StackTestIAMPermissionsRequest(stackTestIAMPermissionsRequest).Execute()

Test stack policies

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
    stackId := "stackId_example" // string | A stack ID or slug
    stackTestIAMPermissionsRequest := openapiclient.stackTestIAMPermissionsRequest{Permissions: []string{"Permissions_example")} // StackTestIAMPermissionsRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StackPoliciesApi.TestIAMPermissions(context.Background(), stackId, stackTestIAMPermissionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StackPoliciesApi.TestIAMPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestIAMPermissions`: StackTestIAMPermissionsResponse
    fmt.Fprintf(os.Stdout, "Response from `StackPoliciesApi.TestIAMPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestIAMPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **stackTestIAMPermissionsRequest** | [**StackTestIAMPermissionsRequest**](StackTestIAMPermissionsRequest.md) |  | 

### Return type

[**StackTestIAMPermissionsResponse**](stackTestIAMPermissionsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

