# \DeliveryDomainsApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSiteDeliveryDomain**](DeliveryDomainsApi.md#CreateSiteDeliveryDomain) | **Post** /delivery/v1/stacks/{stack_id}/sites/{site_id}/delivery_domains | Add a delivery domain to a site
[**DeleteSiteDeliveryDomain**](DeliveryDomainsApi.md#DeleteSiteDeliveryDomain) | **Delete** /delivery/v1/stacks/{stack_id}/sites/{site_id}/delivery_domains/{domain} | Remove a delivery domain from a site
[**GetSiteDeliveryDomains2**](DeliveryDomainsApi.md#GetSiteDeliveryDomains2) | **Get** /delivery/v1/stacks/{stack_id}/sites/{site_id}/delivery_domains | Retrieve the delivery domains configured on a site



## CreateSiteDeliveryDomain

> DeliveryCreateSiteDeliveryDomainResponse CreateSiteDeliveryDomain(ctx, stackId, siteId, deliveryCreateSiteDeliveryDomainRequest)

Add a delivery domain to a site

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
**deliveryCreateSiteDeliveryDomainRequest** | [**DeliveryCreateSiteDeliveryDomainRequest**](DeliveryCreateSiteDeliveryDomainRequest.md)|  | 

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

> DeleteSiteDeliveryDomain(ctx, stackId, siteId, domain)

Remove a delivery domain from a site

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
**domain** | **string**| The delivery domain to remove from a site | 

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

> DeliveryGetSiteDeliveryDomainsResponse GetSiteDeliveryDomains2(ctx, stackId, siteId, optional)

Retrieve the delivery domains configured on a site

Delivery domains allow the CDN to recognize an HTTP request and associate it with a site.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
 **optional** | ***GetSiteDeliveryDomains2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSiteDeliveryDomains2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageRequestFirst** | **optional.String**| The number of items desired. | 
 **pageRequestAfter** | **optional.String**| The cursor value after which data will be returned. | 
 **pageRequestFilter** | **optional.String**| SQL-style constraint filters. | 
 **pageRequestSortBy** | **optional.String**| Sort the response by the given field. | 

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

