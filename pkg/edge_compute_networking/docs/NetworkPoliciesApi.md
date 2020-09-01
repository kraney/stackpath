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

> V1CreateNetworkPolicyResponse CreateNetworkPolicy(ctx, stackId).V1CreateNetworkPolicyRequest(v1CreateNetworkPolicyRequest).Execute()

Create a policy

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
    v1CreateNetworkPolicyRequest := openapiclient.v1CreateNetworkPolicyRequest{NetworkPolicy: openapiclient.v1NetworkPolicy{StackId: "StackId_example", Id: "Id_example", Name: "Name_example", Slug: "Slug_example", Description: "Description_example", Metadata: openapiclient.networkMetadata{Annotations: map[string]string{ "Key" = "Value" }, Labels: map[string]string{ "Key" = "Value" }, CreatedAt: "TODO", UpdatedAt: "TODO", DeleteRequestedAt: "TODO", Version: "Version_example"}, Spec: openapiclient.v1NetworkPolicySpec{InstanceSelectors: []V1MatchExpression{openapiclient.v1MatchExpression{Key: "Key_example", Operator: "Operator_example", Values: []string{"Values_example")}), NetworkSelectors: []V1MatchExpression{openapiclient.v1MatchExpression{Key: "Key_example", Operator: "Operator_example", Values: []string{"Values_example")}), PolicyTypes: []NetworkPolicySpecPolicyType{openapiclient.NetworkPolicySpecPolicyType{}), Priority: 123, Ingress: []V1Ingress{openapiclient.v1Ingress{Description: "Description_example", Action: openapiclient.v1Action{}, From: openapiclient.v1HostRule{IpBlock: []V1IpBlock{openapiclient.v1IpBlock{Cidr: "Cidr_example", Except: []string{"Except_example")}), InstanceSelectors: []V1MatchExpression{), NetworkSelectors: []V1MatchExpression{)}, Protocols: openapiclient.v1Protocols{Tcp: openapiclient.v1ProtocolTcp{DestinationPorts: []string{"DestinationPorts_example"), SourcePorts: []string{"SourcePorts_example")}, Udp: openapiclient.v1ProtocolUdp{DestinationPorts: []string{"DestinationPorts_example"), SourcePorts: []string{"SourcePorts_example")}, TcpUdp: openapiclient.v1ProtocolTcpUdp{DestinationPorts: []string{"DestinationPorts_example"), SourcePorts: []string{"SourcePorts_example")}, Icmp: "TODO", Ah: "TODO", Esp: "TODO", Gre: "TODO"}}), Egress: []V1Egress{openapiclient.v1Egress{Description: "Description_example", Action: openapiclient.v1Action{}, To: openapiclient.v1HostRule{IpBlock: []V1IpBlock{openapiclient.v1IpBlock{Cidr: "Cidr_example", Except: []string{"Except_example")}), InstanceSelectors: []V1MatchExpression{), NetworkSelectors: []V1MatchExpression{)}, Protocols: openapiclient.v1Protocols{Tcp: openapiclient.v1ProtocolTcp{DestinationPorts: []string{"DestinationPorts_example"), SourcePorts: []string{"SourcePorts_example")}, Udp: openapiclient.v1ProtocolUdp{DestinationPorts: []string{"DestinationPorts_example"), SourcePorts: []string{"SourcePorts_example")}, TcpUdp: openapiclient.v1ProtocolTcpUdp{DestinationPorts: []string{"DestinationPorts_example"), SourcePorts: []string{"SourcePorts_example")}, Icmp: "TODO", Ah: "TODO", Esp: "TODO", Gre: "TODO"}})}}} // V1CreateNetworkPolicyRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkPoliciesApi.CreateNetworkPolicy(context.Background(), stackId, v1CreateNetworkPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkPoliciesApi.CreateNetworkPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkPolicy`: V1CreateNetworkPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkPoliciesApi.CreateNetworkPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v1CreateNetworkPolicyRequest** | [**V1CreateNetworkPolicyRequest**](V1CreateNetworkPolicyRequest.md) |  | 

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

> DeleteNetworkPolicy(ctx, stackId, networkPolicyId).Execute()

Delete a policy

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
    networkPolicyId := "networkPolicyId_example" // string | An EdgeCompute network policy ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkPoliciesApi.DeleteNetworkPolicy(context.Background(), stackId, networkPolicyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkPoliciesApi.DeleteNetworkPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**networkPolicyId** | **string** | An EdgeCompute network policy ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkPolicyRequest struct via the builder pattern


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


## GetNetworkPolicies

> V1GetNetworkPoliciesResponse GetNetworkPolicies(ctx, stackId).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).Execute()

Get all policies

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
    pageRequestFirst := "pageRequestFirst_example" // string | The number of items desired. (optional)
    pageRequestAfter := "pageRequestAfter_example" // string | The cursor value after which data will be returned. (optional)
    pageRequestFilter := "pageRequestFilter_example" // string | SQL-style constraint filters. (optional)
    pageRequestSortBy := "pageRequestSortBy_example" // string | Sort the response by the given field. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkPoliciesApi.GetNetworkPolicies(context.Background(), stackId).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkPoliciesApi.GetNetworkPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkPolicies`: V1GetNetworkPoliciesResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkPoliciesApi.GetNetworkPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageRequestFirst** | **string** | The number of items desired. | 
 **pageRequestAfter** | **string** | The cursor value after which data will be returned. | 
 **pageRequestFilter** | **string** | SQL-style constraint filters. | 
 **pageRequestSortBy** | **string** | Sort the response by the given field. | 

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

> V1GetNetworkPolicyResponse GetNetworkPolicy(ctx, stackId, networkPolicyId).Execute()

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
    networkPolicyId := "networkPolicyId_example" // string | An EdgeCompute network policy ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkPoliciesApi.GetNetworkPolicy(context.Background(), stackId, networkPolicyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkPoliciesApi.GetNetworkPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkPolicy`: V1GetNetworkPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkPoliciesApi.GetNetworkPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**networkPolicyId** | **string** | An EdgeCompute network policy ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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

> V1UpdateNetworkPolicyResponse UpdateNetworkPolicy(ctx, stackId, networkPolicyId).V1UpdateNetworkPolicyRequest(v1UpdateNetworkPolicyRequest).Execute()

Update a policy

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
    networkPolicyId := "networkPolicyId_example" // string | An EdgeCompute network policy ID
    v1UpdateNetworkPolicyRequest := openapiclient.v1UpdateNetworkPolicyRequest{NetworkPolicy: openapiclient.v1NetworkPolicy{StackId: "StackId_example", Id: "Id_example", Name: "Name_example", Slug: "Slug_example", Description: "Description_example", Metadata: openapiclient.networkMetadata{Annotations: map[string]string{ "Key" = "Value" }, Labels: map[string]string{ "Key" = "Value" }, CreatedAt: "TODO", UpdatedAt: "TODO", DeleteRequestedAt: "TODO", Version: "Version_example"}, Spec: openapiclient.v1NetworkPolicySpec{InstanceSelectors: []V1MatchExpression{), NetworkSelectors: []V1MatchExpression{), PolicyTypes: []NetworkPolicySpecPolicyType{openapiclient.NetworkPolicySpecPolicyType{}), Priority: 123, Ingress: []V1Ingress{openapiclient.v1Ingress{Description: "Description_example", Action: , From: , Protocols: }), Egress: []V1Egress{openapiclient.v1Egress{Description: "Description_example", Action: , To: , Protocols: })}}} // V1UpdateNetworkPolicyRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkPoliciesApi.UpdateNetworkPolicy(context.Background(), stackId, networkPolicyId, v1UpdateNetworkPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkPoliciesApi.UpdateNetworkPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkPolicy`: V1UpdateNetworkPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkPoliciesApi.UpdateNetworkPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**networkPolicyId** | **string** | An EdgeCompute network policy ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **v1UpdateNetworkPolicyRequest** | [**V1UpdateNetworkPolicyRequest**](V1UpdateNetworkPolicyRequest.md) |  | 

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

