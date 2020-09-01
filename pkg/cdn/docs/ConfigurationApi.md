# \ConfigurationApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConnectScopeToOrigin**](ConfigurationApi.md#ConnectScopeToOrigin) | **Post** /cdn/v1/stacks/{stack_id}/sites/{site_id}/scopes/{scope_id}/origins | Connect an origin to a scope
[**CreateScope**](ConfigurationApi.md#CreateScope) | **Post** /cdn/v1/stacks/{stack_id}/sites/{site_id}/scopes | Create a scope
[**CreateScopeHostname**](ConfigurationApi.md#CreateScopeHostname) | **Post** /cdn/v1/stacks/{stack_id}/sites/{site_id}/scopes/{scope_id}/hostnames | Add a scope hostname
[**DeleteScope**](ConfigurationApi.md#DeleteScope) | **Delete** /cdn/v1/stacks/{stack_id}/sites/{site_id}/scopes/{scope_id} | Delete a scope
[**DeleteScopeHostname**](ConfigurationApi.md#DeleteScopeHostname) | **Delete** /cdn/v1/stacks/{stack_id}/sites/{site_id}/scopes/{scope_id}/hostnames/{domain} | Delete a scope hostname
[**DisconnectScopeFromOrigin**](ConfigurationApi.md#DisconnectScopeFromOrigin) | **Delete** /cdn/v1/stacks/{stack_id}/sites/{site_id}/scopes/{scope_id}/origins/{origin_id} | Delete a scope origin
[**GetScopeConfiguration**](ConfigurationApi.md#GetScopeConfiguration) | **Get** /cdn/v1/stacks/{stack_id}/sites/{site_id}/scopes/{scope_id}/configuration | Get a scope&#39;s configuraiton
[**GetScopeHostnames**](ConfigurationApi.md#GetScopeHostnames) | **Get** /cdn/v1/stacks/{stack_id}/sites/{site_id}/scopes/{scope_id}/hostnames | Get all scope hostnames
[**GetScopeOrigins**](ConfigurationApi.md#GetScopeOrigins) | **Get** /cdn/v1/stacks/{stack_id}/sites/{site_id}/scopes/{scope_id}/origins | Get all scope origins
[**GetSiteDnsTargets**](ConfigurationApi.md#GetSiteDnsTargets) | **Get** /cdn/v1/stacks/{stack_id}/sites/{site_id}/dns/targets | Get CNAME targets
[**GetSiteScopes**](ConfigurationApi.md#GetSiteScopes) | **Get** /cdn/v1/stacks/{stack_id}/sites/{site_id}/scopes | Get all scopes
[**UpdateScopeConfiguration**](ConfigurationApi.md#UpdateScopeConfiguration) | **Patch** /cdn/v1/stacks/{stack_id}/sites/{site_id}/scopes/{scope_id}/configuration | Update a scope&#39;s configuration



## ConnectScopeToOrigin

> CdnConnectScopeToOriginResponse ConnectScopeToOrigin(ctx, stackId, siteId, scopeId).CdnConnectScopeToOriginRequest(cdnConnectScopeToOriginRequest).Execute()

Connect an origin to a scope



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
    scopeId := "scopeId_example" // string | A scope ID
    cdnConnectScopeToOriginRequest := openapiclient.cdnConnectScopeToOriginRequest{Origin: openapiclient.cdnConnectScopeToOriginRequestOrigin{Path: "Path_example", Hostname: "Hostname_example", Port: 123, SecurePort: 123}, Priority: 123, OriginId: "OriginId_example"} // CdnConnectScopeToOriginRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigurationApi.ConnectScopeToOrigin(context.Background(), stackId, siteId, scopeId, cdnConnectScopeToOriginRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.ConnectScopeToOrigin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectScopeToOrigin`: CdnConnectScopeToOriginResponse
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationApi.ConnectScopeToOrigin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**siteId** | **string** | A site ID | 
**scopeId** | **string** | A scope ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectScopeToOriginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **cdnConnectScopeToOriginRequest** | [**CdnConnectScopeToOriginRequest**](CdnConnectScopeToOriginRequest.md) |  | 

### Return type

[**CdnConnectScopeToOriginResponse**](cdnConnectScopeToOriginResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateScope

> CdnCreateScopeResponse CreateScope(ctx, stackId, siteId).CdnCreateScopeRequest(cdnCreateScopeRequest).Execute()

Create a scope

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
    cdnCreateScopeRequest := openapiclient.cdnCreateScopeRequest{Path: "Path_example", Platform: "Platform_example"} // CdnCreateScopeRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigurationApi.CreateScope(context.Background(), stackId, siteId, cdnCreateScopeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.CreateScope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateScope`: CdnCreateScopeResponse
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationApi.CreateScope`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**siteId** | **string** | A site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cdnCreateScopeRequest** | [**CdnCreateScopeRequest**](CdnCreateScopeRequest.md) |  | 

### Return type

[**CdnCreateScopeResponse**](cdnCreateScopeResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateScopeHostname

> CdnCreateScopeHostnameResponse CreateScopeHostname(ctx, stackId, siteId, scopeId).CdnCreateScopeHostnameRequest(cdnCreateScopeHostnameRequest).Execute()

Add a scope hostname

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
    scopeId := "scopeId_example" // string | A scope ID
    cdnCreateScopeHostnameRequest := openapiclient.cdnCreateScopeHostnameRequest{Domain: "Domain_example", DisableTransparentMode: false} // CdnCreateScopeHostnameRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigurationApi.CreateScopeHostname(context.Background(), stackId, siteId, scopeId, cdnCreateScopeHostnameRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.CreateScopeHostname``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateScopeHostname`: CdnCreateScopeHostnameResponse
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationApi.CreateScopeHostname`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**siteId** | **string** | A site ID | 
**scopeId** | **string** | A scope ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateScopeHostnameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **cdnCreateScopeHostnameRequest** | [**CdnCreateScopeHostnameRequest**](CdnCreateScopeHostnameRequest.md) |  | 

### Return type

[**CdnCreateScopeHostnameResponse**](cdnCreateScopeHostnameResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteScope

> DeleteScope(ctx, stackId, siteId, scopeId).Execute()

Delete a scope

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
    scopeId := "scopeId_example" // string | A scope ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigurationApi.DeleteScope(context.Background(), stackId, siteId, scopeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.DeleteScope``: %v\n", err)
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
**scopeId** | **string** | A scope ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteScopeRequest struct via the builder pattern


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


## DeleteScopeHostname

> DeleteScopeHostname(ctx, stackId, siteId, scopeId, domain).DisableTransparentMode(disableTransparentMode).Execute()

Delete a scope hostname

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
    scopeId := "scopeId_example" // string | A scope ID
    domain := "domain_example" // string | The hostname to remove from a scope
    disableTransparentMode := true // bool | Whether or not to remove the hostname from a CDN site's CDN scope or its WAF scope. When true, this call removes the hostname from a CDN site's scope instead of loading from a CDN site's WAF scope, if the site has WAF service. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigurationApi.DeleteScopeHostname(context.Background(), stackId, siteId, scopeId, domain).DisableTransparentMode(disableTransparentMode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.DeleteScopeHostname``: %v\n", err)
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
**scopeId** | **string** | A scope ID | 
**domain** | **string** | The hostname to remove from a scope | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteScopeHostnameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **disableTransparentMode** | **bool** | Whether or not to remove the hostname from a CDN site&#39;s CDN scope or its WAF scope. When true, this call removes the hostname from a CDN site&#39;s scope instead of loading from a CDN site&#39;s WAF scope, if the site has WAF service. | 

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


## DisconnectScopeFromOrigin

> DisconnectScopeFromOrigin(ctx, stackId, siteId, scopeId, originId).Execute()

Delete a scope origin

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
    scopeId := "scopeId_example" // string | A scope ID
    originId := "originId_example" // string | An origin ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigurationApi.DisconnectScopeFromOrigin(context.Background(), stackId, siteId, scopeId, originId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.DisconnectScopeFromOrigin``: %v\n", err)
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
**scopeId** | **string** | A scope ID | 
**originId** | **string** | An origin ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisconnectScopeFromOriginRequest struct via the builder pattern


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


## GetScopeConfiguration

> CdnGetScopeConfigurationResponse GetScopeConfiguration(ctx, stackId, siteId, scopeId).Execute()

Get a scope's configuraiton

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
    scopeId := "scopeId_example" // string | A scope ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigurationApi.GetScopeConfiguration(context.Background(), stackId, siteId, scopeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.GetScopeConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetScopeConfiguration`: CdnGetScopeConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationApi.GetScopeConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**siteId** | **string** | A site ID | 
**scopeId** | **string** | A scope ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetScopeConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**CdnGetScopeConfigurationResponse**](cdnGetScopeConfigurationResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetScopeHostnames

> CdnGetScopeHostnamesResponse GetScopeHostnames(ctx, stackId, siteId, scopeId).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).DisableTransparentMode(disableTransparentMode).Execute()

Get all scope hostnames



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
    scopeId := "scopeId_example" // string | A scope ID
    pageRequestFirst := "pageRequestFirst_example" // string | The number of items desired. (optional)
    pageRequestAfter := "pageRequestAfter_example" // string | The cursor value after which data will be returned. (optional)
    pageRequestFilter := "pageRequestFilter_example" // string | SQL-style constraint filters. (optional)
    pageRequestSortBy := "pageRequestSortBy_example" // string | Sort the response by the given field. (optional)
    disableTransparentMode := true // bool | Whether or not to load hostnames from a CDN site's CDN scope or its WAF scope. When true, this call loads scope hostnames from a CDN site's scope instead of loading from a CDN site's WAF scope, if the site has WAF service. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigurationApi.GetScopeHostnames(context.Background(), stackId, siteId, scopeId).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).DisableTransparentMode(disableTransparentMode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.GetScopeHostnames``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetScopeHostnames`: CdnGetScopeHostnamesResponse
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationApi.GetScopeHostnames`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**siteId** | **string** | A site ID | 
**scopeId** | **string** | A scope ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetScopeHostnamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pageRequestFirst** | **string** | The number of items desired. | 
 **pageRequestAfter** | **string** | The cursor value after which data will be returned. | 
 **pageRequestFilter** | **string** | SQL-style constraint filters. | 
 **pageRequestSortBy** | **string** | Sort the response by the given field. | 
 **disableTransparentMode** | **bool** | Whether or not to load hostnames from a CDN site&#39;s CDN scope or its WAF scope. When true, this call loads scope hostnames from a CDN site&#39;s scope instead of loading from a CDN site&#39;s WAF scope, if the site has WAF service. | 

### Return type

[**CdnGetScopeHostnamesResponse**](cdnGetScopeHostnamesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetScopeOrigins

> CdnGetScopeOriginsResponse GetScopeOrigins(ctx, stackId, siteId, scopeId).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).Execute()

Get all scope origins

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
    scopeId := "scopeId_example" // string | A scope ID
    pageRequestFirst := "pageRequestFirst_example" // string | The number of items desired. (optional)
    pageRequestAfter := "pageRequestAfter_example" // string | The cursor value after which data will be returned. (optional)
    pageRequestFilter := "pageRequestFilter_example" // string | SQL-style constraint filters. (optional)
    pageRequestSortBy := "pageRequestSortBy_example" // string | Sort the response by the given field. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigurationApi.GetScopeOrigins(context.Background(), stackId, siteId, scopeId).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.GetScopeOrigins``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetScopeOrigins`: CdnGetScopeOriginsResponse
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationApi.GetScopeOrigins`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**siteId** | **string** | A site ID | 
**scopeId** | **string** | A scope ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetScopeOriginsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pageRequestFirst** | **string** | The number of items desired. | 
 **pageRequestAfter** | **string** | The cursor value after which data will be returned. | 
 **pageRequestFilter** | **string** | SQL-style constraint filters. | 
 **pageRequestSortBy** | **string** | Sort the response by the given field. | 

### Return type

[**CdnGetScopeOriginsResponse**](cdnGetScopeOriginsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteDnsTargets

> CdnGetSiteDnsTargetsResponse GetSiteDnsTargets(ctx, stackId, siteId).Execute()

Get CNAME targets



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
    resp, r, err := api_client.ConfigurationApi.GetSiteDnsTargets(context.Background(), stackId, siteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.GetSiteDnsTargets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSiteDnsTargets`: CdnGetSiteDnsTargetsResponse
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationApi.GetSiteDnsTargets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**siteId** | **string** | A site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteDnsTargetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CdnGetSiteDnsTargetsResponse**](cdnGetSiteDnsTargetsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteScopes

> CdnGetSiteScopesResponse GetSiteScopes(ctx, stackId, siteId).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).DisableTransparentMode(disableTransparentMode).Execute()

Get all scopes

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
    disableTransparentMode := true // bool | Whether or not to retrieve the site's CDN scope or its WAF scope. When true, this call removes the hostname from a CDN site's scope instead of loading from a CDN site's WAF scope, if the site has WAF service. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigurationApi.GetSiteScopes(context.Background(), stackId, siteId).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).DisableTransparentMode(disableTransparentMode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.GetSiteScopes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSiteScopes`: CdnGetSiteScopesResponse
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationApi.GetSiteScopes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**siteId** | **string** | A site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteScopesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageRequestFirst** | **string** | The number of items desired. | 
 **pageRequestAfter** | **string** | The cursor value after which data will be returned. | 
 **pageRequestFilter** | **string** | SQL-style constraint filters. | 
 **pageRequestSortBy** | **string** | Sort the response by the given field. | 
 **disableTransparentMode** | **bool** | Whether or not to retrieve the site&#39;s CDN scope or its WAF scope. When true, this call removes the hostname from a CDN site&#39;s scope instead of loading from a CDN site&#39;s WAF scope, if the site has WAF service. | 

### Return type

[**CdnGetSiteScopesResponse**](cdnGetSiteScopesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateScopeConfiguration

> CdnUpdateScopeConfigurationResponse UpdateScopeConfiguration(ctx, stackId, siteId, scopeId).CdnUpdateScopeConfigurationRequest(cdnUpdateScopeConfigurationRequest).Execute()

Update a scope's configuration

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
    scopeId := "scopeId_example" // string | A scope ID
    cdnUpdateScopeConfigurationRequest := openapiclient.cdnUpdateScopeConfigurationRequest{Configuration: openapiclient.custconfConfiguration{AccessLogs: openapiclient.custconfAccessLogs{Id: "Id_example", Enabled: false}, AccessLogsConfig: openapiclient.custconfAccessLogsConfig{Id: "Id_example", ExtraLogFields: "ExtraLogFields_example", Enabled: false}, AuthACL: []CustconfAuthACL{openapiclient.custconfAuthACL{Id: "Id_example", AccessCode: openapiclient.AuthACLAccessCodeEnumWrapperValue{}, IpList: "IpList_example", Protocol: openapiclient.custconfAuthACLProtocolEnumWrapperValue{}, ClientIPSrc: openapiclient.AuthACLClientIPSrcEnumWrapperValue{}, Header: "Header_example", Enabled: false}), AuthGeo: []CustconfAuthGeo{openapiclient.custconfAuthGeo{Id: "Id_example", Code: openapiclient.AuthGeoCodeEnumWrapperValue{}, Values: "Values_example", Enabled: false}), AuthHttpBasic: openapiclient.custconfAuthHttpBasic{Id: "Id_example", BindingPoint: "BindingPoint_example", Realm: "Realm_example", Ttl: 123, Enabled: false}, AuthReferer: openapiclient.custconfAuthReferer{Id: "Id_example", Referer: "Referer_example", Enabled: false}, AuthUrlSign: []CustconfAuthUrlSign{openapiclient.custconfAuthUrlSign{Id: "Id_example", TokenField: "TokenField_example", IgnoreFieldsAfterToken: false, PassPhraseField: "PassPhraseField_example", PassPhrase: "PassPhrase_example", ExpiresField: "ExpiresField_example", IpAddressField: "IpAddressField_example", UriLengthField: "UriLengthField_example", UserAgentField: "UserAgentField_example", Enabled: false, MethodFilter: "MethodFilter_example", PathFilter: "PathFilter_example", HeaderFilter: "HeaderFilter_example"}), AuthUrlSignAliCloudA: []CustconfAuthUrlSignAliCloudA{openapiclient.custconfAuthUrlSignAliCloudA{Id: "Id_example", PassPhrase: "PassPhrase_example", TokenField: "TokenField_example", IncludeParamsBeforeToken: false, ExpirationExtension: 123, Enabled: false, MethodFilter: "MethodFilter_example", PathFilter: "PathFilter_example", HeaderFilter: "HeaderFilter_example"}), AuthUrlSignAliCloudB: []CustconfAuthUrlSignAliCloudB{openapiclient.custconfAuthUrlSignAliCloudB{Id: "Id_example", PassPhrase: "PassPhrase_example", ExpirationExtension: 123, Enabled: false, MethodFilter: "MethodFilter_example", PathFilter: "PathFilter_example", HeaderFilter: "HeaderFilter_example"}), AuthUrlSignAliCloudC: []CustconfAuthUrlSignAliCloudC{openapiclient.custconfAuthUrlSignAliCloudC{Id: "Id_example", PassPhrase: "PassPhrase_example", ExpirationExtension: 123, TokenField: "TokenField_example", ExpireField: "ExpireField_example", Enabled: false, MethodFilter: "MethodFilter_example", PathFilter: "PathFilter_example", HeaderFilter: "HeaderFilter_example"}), AuthUrlSignHmacTlu: []CustconfAuthUrlSignHmacTlu{openapiclient.custconfAuthUrlSignHmacTlu{Id: "Id_example", ExpireParameterName: "ExpireParameterName_example", KeyIdParameterName: "KeyIdParameterName_example", AlgorithmIdParameterName: "AlgorithmIdParameterName_example", DigestParameterName: "DigestParameterName_example", SymmetricKeyIdMap: map[string]string{ "Key" = "Value" }, AlgorithmIdMap: map[string]string{ "Key" = "Value" }, Enabled: false, MethodFilter: "MethodFilter_example", PathFilter: "PathFilter_example", HeaderFilter: "HeaderFilter_example"}), AuthUrlSignIq: []CustconfAuthUrlSignIq{openapiclient.custconfAuthUrlSignIq{Id: "Id_example", SecretKey: "SecretKey_example", Enabled: false, MethodFilter: "MethodFilter_example", PathFilter: "PathFilter_example", HeaderFilter: "HeaderFilter_example"}), AuthUrlAsymmetricSignTlu: []CustconfAuthUrlAsymmetricSignTlu{openapiclient.custconfAuthUrlAsymmetricSignTlu{Id: "Id_example", ExpireParameterName: "ExpireParameterName_example", KeyIdParameterName: "KeyIdParameterName_example", AlgorithmIdParameterName: "AlgorithmIdParameterName_example", DigestParameterName: "DigestParameterName_example", PublicKeyIdMap: map[string]string{ "Key" = "Value" }, AlgorithmIdMap: map[string]string{ "Key" = "Value" }, Enabled: false, MethodFilter: "MethodFilter_example", PathFilter: "PathFilter_example", HeaderFilter: "HeaderFilter_example"}), BandWidthLimit: openapiclient.custconfBandWidthLimit{Id: "Id_example", Rule: "Rule_example", Values: "Values_example", Enabled: false}, BandwidthRateLimit: openapiclient.custconfBandwidthRateLimit{Id: "Id_example", InitialBurstName: "InitialBurstName_example", SustainedRateName: "SustainedRateName_example", InitialBurstUnits: openapiclient.BandwidthRateLimitInitialBurstUnitsEnumWrapperValue{}, SustainedRateUnits: openapiclient.BandwidthRateLimitSustainedRateUnitsEnumWrapperValue{}, Enabled: false, MethodFilter: "MethodFilter_example", PathFilter: "PathFilter_example", HeaderFilter: "HeaderFilter_example"}, BypassCache: []CustconfBypassCache{openapiclient.custconfBypassCache{Id: "Id_example", Enabled: false, MethodFilter: "MethodFilter_example", PathFilter: "PathFilter_example", HeaderFilter: "HeaderFilter_example", CookieFilter: "CookieFilter_example"}), CacheControl: []CustconfCacheControl{openapiclient.custconfCacheControl{Id: "Id_example", StatusCodeMatch: "StatusCodeMatch_example", MaxAge: 123, MustRevalidate: false, SynchronizeMaxAge: false, Override: "Override_example", Enabled: false}), CacheKeyModification: openapiclient.custconfCacheKeyModification{Id: "Id_example", NormalizeKeyPathToLowerCase: false, Enabled: false}, ClientRequestModification: []CustconfClientRequestModification{openapiclient.custconfClientRequestModification{Id: "Id_example", UrlPattern: "UrlPattern_example", UrlRewrite: "UrlRewrite_example", HeaderPattern: "HeaderPattern_example", HeaderRewrite: "HeaderRewrite_example", AddHeaders: "AddHeaders_example", FlowControl: openapiclient.custconfClientRequestModificationFlowControlEnumWrapperValue{}, Enabled: false, MethodFilter: "MethodFilter_example", PathFilter: "PathFilter_example", HeaderFilter: "HeaderFilter_example", CookieFilter: "CookieFilter_example"}), ClientResponseModification: []CustconfClientResponseModification{openapiclient.custconfClientResponseModification{Id: "Id_example", StatusCodeRewrite: 123, HeaderPattern: "HeaderPattern_example", HeaderRewrite: "HeaderRewrite_example", AddHeaders: "AddHeaders_example", FlowControl: openapiclient.custconfClientResponseModificationFlowControlEnumWrapperValue{}, Enabled: false, MethodFilter: "MethodFilter_example", PathFilter: "PathFilter_example", HeaderFilter: "HeaderFilter_example"}), Compression: openapiclient.custconfCompression{Id: "Id_example", Gzip: "Gzip_example", Mime: "Mime_example", Level: "Level_example", Enabled: false}, ContentDispositionByURL: openapiclient.custconfContentDispositionByURL{Id: "Id_example", DispositionNameQSParam: "DispositionNameQSParam_example", DispositionTypeQSParam: "DispositionTypeQSParam_example", DispositionOverrideQSParam: "DispositionOverrideQSParam_example", OverrideOriginHeader: false, Enabled: false}, ContentDispositionByHeader: []CustconfContentDispositionByHeader{openapiclient.custconfContentDispositionByHeader{Id: "Id_example", HeaderFieldName: "HeaderFieldName_example", HeaderValueMatch: "HeaderValueMatch_example", DefaultType: openapiclient.ContentDispositionByHeaderDefaultTypeEnumWrapperValue{}, OverrideOriginHeader: false, Enabled: false, MethodFilter: "MethodFilter_example", PathFilter: "PathFilter_example", HeaderFilter: "HeaderFilter_example"}), Customer: openapiclient.custconfCustomer{Id: "Id_example", AccessLogFields: "AccessLogFields_example", OpLogFields: "OpLogFields_example", ReceiptLogFields: "ReceiptLogFields_example"}, CustomHeader: openapiclient.custconfCustomHeader{Id: "Id_example", XForwardedForAuth: "XForwardedForAuth_example", XForwardedForOrigin: "XForwardedForOrigin_example", Enabled: false}, CustomMimeType: []CustconfCustomMimeType{openapiclient.custconfCustomMimeType{Id: "Id_example", Code: "Code_example", ExtensionMap: "ExtensionMap_example", Enabled: false, MethodFilter: "MethodFilter_example", PathFilter: "PathFilter_example", HeaderFilter: "HeaderFilter_example"}), DynamicCacheRule: []CustconfDynamicCacheRule{openapiclient.custconfDynamicCacheRule{Id: "Id_example", StatusCode: 123, Headers: "Headers_example", Enabled: false, MethodFilter: "MethodFilter_example", PathFilter: "PathFilter_example", HeaderFilter: "HeaderFilter_example"}), DynamicContent: []CustconfDynamicContent{openapiclient.custconfDynamicContent{Id: "Id_example", QueryParams: "QueryParams_example", HeaderFields: "HeaderFields_example", Enabled: false, MethodFilter: "MethodFilter_example", PathFilter: "PathFilter_example", HeaderFilter: "HeaderFilter_example"}), FailSafeOriginPull: openapiclient.custconfFailSafeOriginPull{Id: "Id_example", Enabled: false, StatusCodeMatch: "StatusCodeMatch_example", PathFilter: "PathFilter_example"}, FarAheadRangeProxy: openapiclient.custconfFarAheadRangeProxy{Id: "Id_example", Enabled: false, ThresholdBytes: 123}, FileSegmentation: openapiclient.custconfFileSegmentation{Id: "Id_example", Enabled: false}, FlvPseudoStreaming: openapiclient.custconfFlvPseudoStreaming{Id: "Id_example", JumpToByteInitialBytesParam: "JumpToByteInitialBytesParam_example", JumpToByteStartOffsetParam: "JumpToByteStartOffsetParam_example", Enabled: false}, GzipOriginPull: openapiclient.custconfGzipOriginPull{Id: "Id_example", Enabled: false}, HttpMethods: openapiclient.custconfHttpMethods{Id: "Id_example", PassThru: "PassThru_example", Enabled: false}, Origin: []CdncustconfOrigin{openapiclient.cdncustconfOrigin{Id: "Id_example", Host: "Host_example", OriginPullHeaders: "OriginPullHeaders_example", OriginCacheHeaders: "OriginCacheHeaders_example"}), OriginPersistentConnections: openapiclient.custconfOriginPersistentConnections{Id: "Id_example", Enabled: false}, OriginPull: openapiclient.custconfOriginPull{Id: "Id_example", RedirectAction: openapiclient.OriginPullRedirectActionEnumWrapperValue{}, NoQSParams: false, RetryMethods: "RetryMethods_example", Enabled: false}, OriginPullCacheExtension: openapiclient.custconfOriginPullCacheExtension{Id: "Id_example", ExpiredCacheExtension: 123, OriginUnreachableCacheExtension: 123, Enabled: false}, OriginPullHost: []CustconfOriginPullHost{openapiclient.custconfOriginPullHost{Id: "Id_example", OriginUrl: "OriginUrl_example", UserName: "UserName_example", Password: "Password_example", MethodFilter: "MethodFilter_example", PathFilter: "PathFilter_example", HeaderFilter: "HeaderFilter_example"}), OriginPullProtocol: openapiclient.custconfOriginPullProtocol{Id: "Id_example", Protocol: openapiclient.custconfOriginPullProtocolProtocolEnumWrapperValue{}, EnableSNI: false}, OriginPullLogs: openapiclient.custconfOriginPullLogs{Id: "Id_example", Enabled: false}, OriginPullLogsConfig: openapiclient.custconfOriginPullLogsConfig{Id: "Id_example", ExtraLogFields: "ExtraLogFields_example", Enabled: false}, OriginPullPolicy: []CustconfOriginPullPolicy{openapiclient.custconfOriginPullPolicy{Id: "Id_example", StatusCodeMatch: "StatusCodeMatch_example", ExpirePolicy: openapiclient.OriginPullPolicyExpirePolicyEnumWrapperValue{}, ExpireSeconds: 123, HonorNoStore: false, HonorNoCache: false, HonorMustRevalidate: false, NoCacheBehavior: openapiclient.OriginPullPolicyNoCacheBehaviorEnumWrapperValue{}, MaxAgeZeroToNoCache: false, MustRevalidateToNoCache: false, BypassCacheIdentifier: "BypassCacheIdentifier_example", ForceBypassCache: false, HttpHeaders: "HttpHeaders_example", HonorPrivate: false, HonorSMaxAge: false, UpdateHttpHeadersOn304Response: false, DefaultCacheBehavior: openapiclient.OriginPullPolicyDefaultCacheBehaviorEnumWrapperValue{}, Enabled: false, MethodFilter: "MethodFilter_example", PathFilter: "PathFilter_example", HeaderFilter: "HeaderFilter_example", ContentTypeFilter: "ContentTypeFilter_example"}), OriginRequestModification: []CustconfOriginRequestModification{openapiclient.custconfOriginRequestModification{Id: "Id_example", UrlPattern: "UrlPattern_example", UrlRewrite: "UrlRewrite_example", HeaderPattern: "HeaderPattern_example", HeaderRewrite: "HeaderRewrite_example", AddHeaders: "AddHeaders_example", FlowControl: "FlowControl_example", Enabled: false, MethodFilter: "MethodFilter_example", PathFilter: "PathFilter_example", HeaderFilter: "HeaderFilter_example", CookieFilter: "CookieFilter_example"}), OriginResponseModification: []CustconfOriginResponseModification{openapiclient.custconfOriginResponseModification{Id: "Id_example", StatusCodeRewrite: 123, HeaderPattern: "HeaderPattern_example", HeaderRewrite: "HeaderRewrite_example", AddHeaders: "AddHeaders_example", FlowControl: openapiclient.custconfOriginResponseModificationFlowControlEnumWrapperValue{}, Enabled: false, MethodFilter: "MethodFilter_example", PathFilter: "PathFilter_example", HeaderFilter: "HeaderFilter_example"}), QueryStrParam: openapiclient.custconfQueryStrParam{Id: "Id_example", DispositionName: "DispositionName_example", DispositionType: "DispositionType_example", DispositionOverride: "DispositionOverride_example", JumpToTimeStart: "JumpToTimeStart_example", JumpToTimeEnd: "JumpToTimeEnd_example", JumpToByteInitialBytes: "JumpToByteInitialBytes_example", JumpToByteStartOffset: "JumpToByteStartOffset_example", RateLimitInitial: "RateLimitInitial_example", RateLimitSustained: "RateLimitSustained_example", Enabled: false}, ReceiptLogsConfig: openapiclient.custconfReceiptLogsConfig{Id: "Id_example", ExtraLogFields: "ExtraLogFields_example", Enabled: false}, RedirectExceptions: openapiclient.custconfRedirectExceptions{Id: "Id_example", RedirectAgentCode: "RedirectAgentCode_example", Enabled: false}, RedirectMappings: []CustconfRedirectMappings{openapiclient.custconfRedirectMappings{Id: "Id_example", Code: 123, RedirectURL: "RedirectURL_example", ReplacementToken: "ReplacementToken_example", Enabled: false, MethodFilter: "MethodFilter_example", PathFilter: "PathFilter_example", HeaderFilter: "HeaderFilter_example"}), ResponseHeader: openapiclient.custconfResponseHeader{Id: "Id_example", Http: "Http_example", EnableETag: false, Enabled: false}, OriginPullResumeDownload: []CustconfOriginPullResumeDownload{openapiclient.custconfOriginPullResumeDownload{Id: "Id_example", Enabled: false, OriginalStatusCodeMatch: "OriginalStatusCodeMatch_example", MinimumBodyBytesConsumed: "MinimumBodyBytesConsumed_example", MaximumAttempts: 123, RequireOriginRangeSupport: false, EtagValidation: openapiclient.OriginPullResumeDownloadEtagValidationEnumWrapperValue{}, HeaderFilter: "HeaderFilter_example", PathFilter: "PathFilter_example"}), RobotsTxt: []CustconfRobotsTxt{openapiclient.custconfRobotsTxt{Id: "Id_example", File: "File_example", CacheControlHeader: "CacheControlHeader_example", Enabled: false, MethodFilter: "MethodFilter_example", PathFilter: "PathFilter_example", HeaderFilter: "HeaderFilter_example"}), AwsSignedOriginPullV4: []CustconfAwsSignedOriginPullV4{openapiclient.custconfAwsSignedOriginPullV4{Id: "Id_example", Enabled: false, BucketIdentifier: "BucketIdentifier_example", AccessKeyId: "AccessKeyId_example", SecretAccessKey: "SecretAccessKey_example", AwsRegion: "AwsRegion_example", AwsService: "AwsService_example", ExpireTimeSeconds: 123, SignedHeaders: "SignedHeaders_example", AuthenticationType: openapiclient.custconfAwsSignedOriginPullV4AuthenticationTypeEnumWrapperValue{}, HeaderFilter: "HeaderFilter_example", MethodFilter: "MethodFilter_example", PathFilter: "PathFilter_example"}), AwsSignedS3PostV4: []CustconfAwsSignedS3PostV4{openapiclient.custconfAwsSignedS3PostV4{Id: "Id_example", Enabled: false, BucketIdentifier: "BucketIdentifier_example", AccessKeyId: "AccessKeyId_example", SecretAccessKey: "SecretAccessKey_example", AwsRegion: "AwsRegion_example", AwsService: "AwsService_example", ExpireTimeSeconds: 123, SignedHeaders: "SignedHeaders_example", AuthenticationType: openapiclient.custconfAwsSignedS3PostV4AuthenticationTypeEnumWrapperValue{}, HeaderFilter: "HeaderFilter_example", MethodFilter: "MethodFilter_example", PathFilter: "PathFilter_example"}), StaticHeader: []CustconfStaticHeader{openapiclient.custconfStaticHeader{Id: "Id_example", ClientRequest: "ClientRequest_example", Http: "Http_example", OriginPull: "OriginPull_example", Enabled: false, MethodFilter: "MethodFilter_example", PathFilter: "PathFilter_example", HeaderFilter: "HeaderFilter_example"}), TimePseudoStreaming: openapiclient.custconfTimePseudoStreaming{Id: "Id_example", JumpToTimeStartParam: "JumpToTimeStartParam_example", JumpToTimeEndParam: "JumpToTimeEndParam_example", Enabled: false}, Http2Support: openapiclient.custconfHttp2Support{Id: "Id_example", Enabled: false}, VHost: []CustconfVHost{openapiclient.custconfVHost{Id: "Id_example", Domain: "Domain_example", Enabled: false})}} // CdnUpdateScopeConfigurationRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigurationApi.UpdateScopeConfiguration(context.Background(), stackId, siteId, scopeId, cdnUpdateScopeConfigurationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.UpdateScopeConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateScopeConfiguration`: CdnUpdateScopeConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationApi.UpdateScopeConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**siteId** | **string** | A site ID | 
**scopeId** | **string** | A scope ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateScopeConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **cdnUpdateScopeConfigurationRequest** | [**CdnUpdateScopeConfigurationRequest**](CdnUpdateScopeConfigurationRequest.md) |  | 

### Return type

[**CdnUpdateScopeConfigurationResponse**](cdnUpdateScopeConfigurationResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

