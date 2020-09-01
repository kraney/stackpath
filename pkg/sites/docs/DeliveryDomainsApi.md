# \DeliveryDomainsApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSiteDeliveryDomain**](DeliveryDomainsApi.md#CreateSiteDeliveryDomain) | **Post** /delivery/v1/stacks/{stack_id}/sites/{site_id}/delivery_domains | Add a delivery domain to a site
[**DeleteSiteDeliveryDomain**](DeliveryDomainsApi.md#DeleteSiteDeliveryDomain) | **Delete** /delivery/v1/stacks/{stack_id}/sites/{site_id}/delivery_domains/{domain} | Remove a delivery domain from a site
[**GetSiteDeliveryDomains2**](DeliveryDomainsApi.md#GetSiteDeliveryDomains2) | **Get** /delivery/v1/stacks/{stack_id}/sites/{site_id}/delivery_domains | Retrieve the delivery domains configured on a site



## CreateSiteDeliveryDomain

> DeliveryCreateSiteDeliveryDomainResponse CreateSiteDeliveryDomain(ctx, stackId, siteId).DeliveryCreateSiteDeliveryDomainRequest(deliveryCreateSiteDeliveryDomainRequest).Execute()

Add a delivery domain to a site

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
    deliveryCreateSiteDeliveryDomainRequest := openapiclient.deliveryCreateSiteDeliveryDomainRequest{Domain: "Domain_example"} // DeliveryCreateSiteDeliveryDomainRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeliveryDomainsApi.CreateSiteDeliveryDomain(context.Background(), stackId, siteId, deliveryCreateSiteDeliveryDomainRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryDomainsApi.CreateSiteDeliveryDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSiteDeliveryDomain`: DeliveryCreateSiteDeliveryDomainResponse
    fmt.Fprintf(os.Stdout, "Response from `DeliveryDomainsApi.CreateSiteDeliveryDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**siteId** | **string** | A site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSiteDeliveryDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deliveryCreateSiteDeliveryDomainRequest** | [**DeliveryCreateSiteDeliveryDomainRequest**](DeliveryCreateSiteDeliveryDomainRequest.md) |  | 

### Return type

[**DeliveryCreateSiteDeliveryDomainResponse**](deliveryCreateSiteDeliveryDomainResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSiteDeliveryDomain

> DeleteSiteDeliveryDomain(ctx, stackId, siteId, domain).Execute()

Remove a delivery domain from a site

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
    domain := "domain_example" // string | The delivery domain to remove from a site

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeliveryDomainsApi.DeleteSiteDeliveryDomain(context.Background(), stackId, siteId, domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryDomainsApi.DeleteSiteDeliveryDomain``: %v\n", err)
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
**domain** | **string** | The delivery domain to remove from a site | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSiteDeliveryDomainRequest struct via the builder pattern


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


## GetSiteDeliveryDomains2

> DeliveryGetSiteDeliveryDomainsResponse GetSiteDeliveryDomains2(ctx, stackId, siteId).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).Execute()

Retrieve the delivery domains configured on a site



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
    pageRequestFirst := "pageRequestFirst_example" // string | The number of items desired. (optional)
    pageRequestAfter := "pageRequestAfter_example" // string | The cursor value after which data will be returned. (optional)
    pageRequestFilter := "pageRequestFilter_example" // string | SQL-style constraint filters. (optional)
    pageRequestSortBy := "pageRequestSortBy_example" // string | Sort the response by the given field. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeliveryDomainsApi.GetSiteDeliveryDomains2(context.Background(), stackId, siteId).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryDomainsApi.GetSiteDeliveryDomains2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSiteDeliveryDomains2`: DeliveryGetSiteDeliveryDomainsResponse
    fmt.Fprintf(os.Stdout, "Response from `DeliveryDomainsApi.GetSiteDeliveryDomains2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**siteId** | **string** | A site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteDeliveryDomains2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageRequestFirst** | **string** | The number of items desired. | 
 **pageRequestAfter** | **string** | The cursor value after which data will be returned. | 
 **pageRequestFilter** | **string** | SQL-style constraint filters. | 
 **pageRequestSortBy** | **string** | Sort the response by the given field. | 

### Return type

[**DeliveryGetSiteDeliveryDomainsResponse**](deliveryGetSiteDeliveryDomainsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

