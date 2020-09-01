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

> DisablePolicy(ctx, stackId, siteId, policyGroupId, policyId).Execute()

Disable a policy

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
    siteId := "siteId_example" // string | A site ID
    policyGroupId := "policyGroupId_example" // string | A WAF policy group ID
    policyId := "policyId_example" // string | A WAF policy ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.DisablePolicy(context.Background(), stackId, siteId, policyGroupId, policyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.DisablePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**siteId** | **string** | A site ID | 
**policyGroupId** | **string** | A WAF policy group ID | 
**policyId** | **string** | A WAF policy ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisablePolicyRequest struct via the builder pattern


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


## DisablePolicyGroup

> DisablePolicyGroup(ctx, stackId, siteId, policyGroupId).Execute()

Disable all policies in a group

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
    siteId := "siteId_example" // string | A site ID
    policyGroupId := "policyGroupId_example" // string | A WAF policy group ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.DisablePolicyGroup(context.Background(), stackId, siteId, policyGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.DisablePolicyGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**siteId** | **string** | A site ID | 
**policyGroupId** | **string** | A WAF policy group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisablePolicyGroupRequest struct via the builder pattern


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


## EnablePolicy

> EnablePolicy(ctx, stackId, siteId, policyGroupId, policyId).Execute()

Enable a policy

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
    siteId := "siteId_example" // string | A site ID
    policyGroupId := "policyGroupId_example" // string | A WAF policy group ID
    policyId := "policyId_example" // string | A WAF policy ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.EnablePolicy(context.Background(), stackId, siteId, policyGroupId, policyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.EnablePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**siteId** | **string** | A site ID | 
**policyGroupId** | **string** | A WAF policy group ID | 
**policyId** | **string** | A WAF policy ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnablePolicyRequest struct via the builder pattern


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


## EnablePolicyGroup

> EnablePolicyGroup(ctx, stackId, siteId, policyGroupId).Execute()

Enable all policies in a group

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
    siteId := "siteId_example" // string | A site ID
    policyGroupId := "policyGroupId_example" // string | A WAF policy group ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.EnablePolicyGroup(context.Background(), stackId, siteId, policyGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.EnablePolicyGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**siteId** | **string** | A site ID | 
**policyGroupId** | **string** | A WAF policy group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnablePolicyGroupRequest struct via the builder pattern


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


## GetPolicies

> WafGetPoliciesResponse GetPolicies(ctx, stackId, siteId, policyGroupId).Execute()

Get all policies in a group

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
    siteId := "siteId_example" // string | A site ID
    policyGroupId := "policyGroupId_example" // string | A WAF policy group ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.GetPolicies(context.Background(), stackId, siteId, policyGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.GetPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPolicies`: WafGetPoliciesResponse
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.GetPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**siteId** | **string** | A site ID | 
**policyGroupId** | **string** | A WAF policy group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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

> WafGetPolicyResponse GetPolicy(ctx, stackId, siteId, policyGroupId, policyId).Execute()

Get a policy

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
    siteId := "siteId_example" // string | A site ID
    policyGroupId := "policyGroupId_example" // string | A WAF policy group ID
    policyId := "policyId_example" // string | A WAF policy ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.GetPolicy(context.Background(), stackId, siteId, policyGroupId, policyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.GetPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPolicy`: WafGetPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.GetPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**siteId** | **string** | A site ID | 
**policyGroupId** | **string** | A WAF policy group ID | 
**policyId** | **string** | A WAF policy ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





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

> WafGetPolicyGroupResponse GetPolicyGroup(ctx, stackId, siteId, policyGroupId).Execute()

Get a policy group

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
    siteId := "siteId_example" // string | A site ID
    policyGroupId := "policyGroupId_example" // string | A WAF policy group ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.GetPolicyGroup(context.Background(), stackId, siteId, policyGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.GetPolicyGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPolicyGroup`: WafGetPolicyGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.GetPolicyGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**siteId** | **string** | A site ID | 
**policyGroupId** | **string** | A WAF policy group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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

> WafGetPolicyGroupsResponse GetPolicyGroups(ctx, stackId, siteId).Execute()

Get all policy groups

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
    siteId := "siteId_example" // string | A site ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.GetPolicyGroups(context.Background(), stackId, siteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.GetPolicyGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPolicyGroups`: WafGetPolicyGroupsResponse
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.GetPolicyGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**siteId** | **string** | A site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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

> WafGetPolicyGroupsResponse UpdatePolicyGroups(ctx, stackId, siteId).WafUpdatePolicyGroupsRequest(wafUpdatePolicyGroupsRequest).Execute()

Update policy groups



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
    siteId := "siteId_example" // string | A site ID
    wafUpdatePolicyGroupsRequest := openapiclient.wafUpdatePolicyGroupsRequest{PolicyGroups: []WafPolicyGroup{openapiclient.wafPolicyGroup{Id: "Id_example", Name: "Name_example", Policies: []SchemawafPolicy{openapiclient.schemawafPolicy{Id: "Id_example", Name: "Name_example", Description: "Description_example", Action: openapiclient.wafPolicyAction{}, Enabled: false})})} // WafUpdatePolicyGroupsRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.UpdatePolicyGroups(context.Background(), stackId, siteId, wafUpdatePolicyGroupsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.UpdatePolicyGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePolicyGroups`: WafGetPolicyGroupsResponse
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.UpdatePolicyGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**siteId** | **string** | A site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePolicyGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **wafUpdatePolicyGroupsRequest** | [**WafUpdatePolicyGroupsRequest**](WafUpdatePolicyGroupsRequest.md) |  | 

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

