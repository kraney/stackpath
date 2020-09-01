# \SSLApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConnectSiteToCertificate**](SSLApi.md#ConnectSiteToCertificate) | **Put** /waf/v1/stacks/{stack_id}/sites/{site_id}/delivery/certificates/{certificate_id} | Update SAN hosts
[**CreateCertificate**](SSLApi.md#CreateCertificate) | **Post** /waf/v1/stacks/{stack_id}/delivery/certificates | Add a certificate
[**DeleteCertificate**](SSLApi.md#DeleteCertificate) | **Delete** /waf/v1/stacks/{stack_id}/delivery/certificates/{certificate_id} | Delete a certificate
[**GetCertificate**](SSLApi.md#GetCertificate) | **Get** /waf/v1/stacks/{stack_id}/delivery/certificates/{certificate_id} | Get a certificate
[**GetCertificateVerificationDetails**](SSLApi.md#GetCertificateVerificationDetails) | **Get** /waf/v1/stacks/{stack_id}/delivery/certificates/{certificate_id}/verification_details | Get sites associated with a certificate
[**GetCertificates**](SSLApi.md#GetCertificates) | **Get** /waf/v1/stacks/{stack_id}/delivery/certificates | Get all certificates
[**GetSiteCertificates**](SSLApi.md#GetSiteCertificates) | **Get** /waf/v1/stacks/{stack_id}/sites/{site_id}/delivery/certificates | Get all site certificates
[**RenewCertificate**](SSLApi.md#RenewCertificate) | **Post** /waf/v1/stacks/{stack_id}/delivery/certificates/{certificate_id}/renew | Renew a certificate
[**RequestCertificate**](SSLApi.md#RequestCertificate) | **Post** /waf/v1/stacks/{stack_id}/sites/{site_id}/delivery/certificates/request | Request a free certificate
[**UpdateCertificate**](SSLApi.md#UpdateCertificate) | **Put** /waf/v1/stacks/{stack_id}/delivery/certificates/{certificate_id} | Update a certificate



## ConnectSiteToCertificate

> WafConnectSiteToCertificateResponse ConnectSiteToCertificate(ctx, stackId, siteId, certificateId).Execute()

Update SAN hosts



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
    certificateId := "certificateId_example" // string | A certificate ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SSLApi.ConnectSiteToCertificate(context.Background(), stackId, siteId, certificateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSLApi.ConnectSiteToCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectSiteToCertificate`: WafConnectSiteToCertificateResponse
    fmt.Fprintf(os.Stdout, "Response from `SSLApi.ConnectSiteToCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**siteId** | **string** | A site ID | 
**certificateId** | **string** | A certificate ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectSiteToCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**WafConnectSiteToCertificateResponse**](wafConnectSiteToCertificateResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCertificate

> WafCreateCertificateResponse CreateCertificate(ctx, stackId).WafCreateCertificateRequest(wafCreateCertificateRequest).Execute()

Add a certificate

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
    wafCreateCertificateRequest := openapiclient.wafCreateCertificateRequest{Certificate: "Certificate_example", Key: "Key_example", CaBundle: "CaBundle_example"} // WafCreateCertificateRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SSLApi.CreateCertificate(context.Background(), stackId, wafCreateCertificateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSLApi.CreateCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCertificate`: WafCreateCertificateResponse
    fmt.Fprintf(os.Stdout, "Response from `SSLApi.CreateCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **wafCreateCertificateRequest** | [**WafCreateCertificateRequest**](WafCreateCertificateRequest.md) |  | 

### Return type

[**WafCreateCertificateResponse**](wafCreateCertificateResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCertificate

> DeleteCertificate(ctx, stackId, certificateId).Execute()

Delete a certificate

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
    certificateId := "certificateId_example" // string | A certificate ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SSLApi.DeleteCertificate(context.Background(), stackId, certificateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSLApi.DeleteCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**certificateId** | **string** | A certificate ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCertificateRequest struct via the builder pattern


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


## GetCertificate

> WafGetCertificateResponse GetCertificate(ctx, stackId, certificateId).Execute()

Get a certificate

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
    certificateId := "certificateId_example" // string | A certificate ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SSLApi.GetCertificate(context.Background(), stackId, certificateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSLApi.GetCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertificate`: WafGetCertificateResponse
    fmt.Fprintf(os.Stdout, "Response from `SSLApi.GetCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**certificateId** | **string** | A certificate ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WafGetCertificateResponse**](wafGetCertificateResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertificateVerificationDetails

> WafGetCertificateVerificationDetailsResponse GetCertificateVerificationDetails(ctx, stackId, certificateId).Execute()

Get sites associated with a certificate

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
    certificateId := "certificateId_example" // string | A certificate ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SSLApi.GetCertificateVerificationDetails(context.Background(), stackId, certificateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSLApi.GetCertificateVerificationDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertificateVerificationDetails`: WafGetCertificateVerificationDetailsResponse
    fmt.Fprintf(os.Stdout, "Response from `SSLApi.GetCertificateVerificationDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**certificateId** | **string** | A certificate ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificateVerificationDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WafGetCertificateVerificationDetailsResponse**](wafGetCertificateVerificationDetailsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertificates

> WafGetCertificatesResponse GetCertificates(ctx, stackId).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).Execute()

Get all certificates

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
    resp, r, err := api_client.SSLApi.GetCertificates(context.Background(), stackId).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSLApi.GetCertificates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertificates`: WafGetCertificatesResponse
    fmt.Fprintf(os.Stdout, "Response from `SSLApi.GetCertificates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageRequestFirst** | **string** | The number of items desired. | 
 **pageRequestAfter** | **string** | The cursor value after which data will be returned. | 
 **pageRequestFilter** | **string** | SQL-style constraint filters. | 
 **pageRequestSortBy** | **string** | Sort the response by the given field. | 

### Return type

[**WafGetCertificatesResponse**](wafGetCertificatesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteCertificates

> WafGetSiteCertificatesResponse GetSiteCertificates(ctx, stackId, siteId).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).Execute()

Get all site certificates

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
    resp, r, err := api_client.SSLApi.GetSiteCertificates(context.Background(), stackId, siteId).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSLApi.GetSiteCertificates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSiteCertificates`: WafGetSiteCertificatesResponse
    fmt.Fprintf(os.Stdout, "Response from `SSLApi.GetSiteCertificates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**siteId** | **string** | A site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteCertificatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageRequestFirst** | **string** | The number of items desired. | 
 **pageRequestAfter** | **string** | The cursor value after which data will be returned. | 
 **pageRequestFilter** | **string** | SQL-style constraint filters. | 
 **pageRequestSortBy** | **string** | Sort the response by the given field. | 

### Return type

[**WafGetSiteCertificatesResponse**](wafGetSiteCertificatesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenewCertificate

> RenewCertificate(ctx, stackId, certificateId).Execute()

Renew a certificate



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
    certificateId := "certificateId_example" // string | A certificate ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SSLApi.RenewCertificate(context.Background(), stackId, certificateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSLApi.RenewCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**certificateId** | **string** | A certificate ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRenewCertificateRequest struct via the builder pattern


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


## RequestCertificate

> WafRequestCertificateResponse RequestCertificate(ctx, stackId, siteId).WafRequestCertificateRequest(wafRequestCertificateRequest).Execute()

Request a free certificate

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
    wafRequestCertificateRequest := openapiclient.wafRequestCertificateRequest{Hosts: []string{"Hosts_example"), VerificationMethod: openapiclient.wafCertificateVerificationMethod{}} // WafRequestCertificateRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SSLApi.RequestCertificate(context.Background(), stackId, siteId, wafRequestCertificateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSLApi.RequestCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RequestCertificate`: WafRequestCertificateResponse
    fmt.Fprintf(os.Stdout, "Response from `SSLApi.RequestCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**siteId** | **string** | A site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **wafRequestCertificateRequest** | [**WafRequestCertificateRequest**](WafRequestCertificateRequest.md) |  | 

### Return type

[**WafRequestCertificateResponse**](wafRequestCertificateResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCertificate

> WafUpdateCertificateResponse UpdateCertificate(ctx, stackId, certificateId).WafUpdateCertificateRequest(wafUpdateCertificateRequest).Execute()

Update a certificate

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
    certificateId := "certificateId_example" // string | A certificate ID
    wafUpdateCertificateRequest := openapiclient.wafUpdateCertificateRequest{Certificate: "Certificate_example", Key: "Key_example", CaBundle: "CaBundle_example"} // WafUpdateCertificateRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SSLApi.UpdateCertificate(context.Background(), stackId, certificateId, wafUpdateCertificateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSLApi.UpdateCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCertificate`: WafUpdateCertificateResponse
    fmt.Fprintf(os.Stdout, "Response from `SSLApi.UpdateCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**certificateId** | **string** | A certificate ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **wafUpdateCertificateRequest** | [**WafUpdateCertificateRequest**](WafUpdateCertificateRequest.md) |  | 

### Return type

[**WafUpdateCertificateResponse**](wafUpdateCertificateResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

