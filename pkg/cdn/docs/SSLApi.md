# \SSLApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConnectSiteToCertificate**](SSLApi.md#ConnectSiteToCertificate) | **Put** /cdn/v1/stacks/{stack_id}/sites/{site_id}/certificates/{certificate_id} | Get a site&#39;s certificate
[**CreateCertificate**](SSLApi.md#CreateCertificate) | **Post** /cdn/v1/stacks/{stack_id}/certificates | Add a certificate
[**DeleteCertificate**](SSLApi.md#DeleteCertificate) | **Delete** /cdn/v1/stacks/{stack_id}/certificates/{certificate_id} | Delete a certificate
[**GetCertificate**](SSLApi.md#GetCertificate) | **Get** /cdn/v1/stacks/{stack_id}/certificates/{certificate_id} | Get a certificate
[**GetCertificateSites**](SSLApi.md#GetCertificateSites) | **Get** /cdn/v1/stacks/{stack_id}/certificates/{certificate_id}/sites | Get sites associated with a certificate
[**GetCertificateVerificationDetails**](SSLApi.md#GetCertificateVerificationDetails) | **Get** /cdn/v1/stacks/{stack_id}/certificates/{certificate_id}/verification_details | Get verification details
[**GetCertificates**](SSLApi.md#GetCertificates) | **Get** /cdn/v1/stacks/{stack_id}/certificates | Get all certificates
[**GetSiteCertificates**](SSLApi.md#GetSiteCertificates) | **Get** /cdn/v1/stacks/{stack_id}/sites/{site_id}/certificates | Get all site certificates
[**RenewCertificate**](SSLApi.md#RenewCertificate) | **Post** /cdn/v1/stacks/{stack_id}/certificates/{certificate_id}/renew | Renew a certificate
[**RequestCertificate**](SSLApi.md#RequestCertificate) | **Post** /cdn/v1/stacks/{stack_id}/sites/{site_id}/certificates/request | Request a free certificate
[**UpdateCertificate**](SSLApi.md#UpdateCertificate) | **Put** /cdn/v1/stacks/{stack_id}/certificates/{certificate_id} | Update a certificate
[**UpdateSiteCertificateHosts**](SSLApi.md#UpdateSiteCertificateHosts) | **Put** /cdn/v1/stacks/{stack_id}/sites/{site_id}/certificates/{certificate_id}/hosts | Update SAN hosts



## ConnectSiteToCertificate

> CdnConnectSiteToCertificateResponse ConnectSiteToCertificate(ctx, stackId, siteId, certificateId).Execute()

Get a site's certificate



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
    // response from `ConnectSiteToCertificate`: CdnConnectSiteToCertificateResponse
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

[**CdnConnectSiteToCertificateResponse**](cdnConnectSiteToCertificateResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCertificate

> CdnCreateCertificateResponse CreateCertificate(ctx, stackId).CdnCreateCertificateRequest(cdnCreateCertificateRequest).Execute()

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
    cdnCreateCertificateRequest := openapiclient.cdnCreateCertificateRequest{Certificate: "Certificate_example", Key: "Key_example", CaBundle: "CaBundle_example"} // CdnCreateCertificateRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SSLApi.CreateCertificate(context.Background(), stackId, cdnCreateCertificateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSLApi.CreateCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCertificate`: CdnCreateCertificateResponse
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

 **cdnCreateCertificateRequest** | [**CdnCreateCertificateRequest**](CdnCreateCertificateRequest.md) |  | 

### Return type

[**CdnCreateCertificateResponse**](cdnCreateCertificateResponse.md)

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

> CdnGetCertificateResponse GetCertificate(ctx, stackId, certificateId).Execute()

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
    // response from `GetCertificate`: CdnGetCertificateResponse
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

[**CdnGetCertificateResponse**](cdnGetCertificateResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertificateSites

> CdnGetCertificateSitesResponse GetCertificateSites(ctx, stackId, certificateId).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).Execute()

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
    pageRequestFirst := "pageRequestFirst_example" // string | The number of items desired. (optional)
    pageRequestAfter := "pageRequestAfter_example" // string | The cursor value after which data will be returned. (optional)
    pageRequestFilter := "pageRequestFilter_example" // string | SQL-style constraint filters. (optional)
    pageRequestSortBy := "pageRequestSortBy_example" // string | Sort the response by the given field. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SSLApi.GetCertificateSites(context.Background(), stackId, certificateId).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSLApi.GetCertificateSites``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertificateSites`: CdnGetCertificateSitesResponse
    fmt.Fprintf(os.Stdout, "Response from `SSLApi.GetCertificateSites`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**certificateId** | **string** | A certificate ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificateSitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageRequestFirst** | **string** | The number of items desired. | 
 **pageRequestAfter** | **string** | The cursor value after which data will be returned. | 
 **pageRequestFilter** | **string** | SQL-style constraint filters. | 
 **pageRequestSortBy** | **string** | Sort the response by the given field. | 

### Return type

[**CdnGetCertificateSitesResponse**](cdnGetCertificateSitesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertificateVerificationDetails

> CdnGetCertificateVerificationDetailsResponse GetCertificateVerificationDetails(ctx, stackId, certificateId).Execute()

Get verification details

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
    // response from `GetCertificateVerificationDetails`: CdnGetCertificateVerificationDetailsResponse
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

[**CdnGetCertificateVerificationDetailsResponse**](cdnGetCertificateVerificationDetailsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertificates

> CdnGetCertificatesResponse GetCertificates(ctx, stackId).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).Execute()

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
    // response from `GetCertificates`: CdnGetCertificatesResponse
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

[**CdnGetCertificatesResponse**](cdnGetCertificatesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteCertificates

> CdnGetSiteCertificatesResponse GetSiteCertificates(ctx, stackId, siteId).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).Execute()

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
    // response from `GetSiteCertificates`: CdnGetSiteCertificatesResponse
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

[**CdnGetSiteCertificatesResponse**](cdnGetSiteCertificatesResponse.md)

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

> CdnRequestCertificateResponse RequestCertificate(ctx, stackId, siteId).CdnRequestCertificateRequest(cdnRequestCertificateRequest).Execute()

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
    cdnRequestCertificateRequest := openapiclient.cdnRequestCertificateRequest{Hosts: []string{"Hosts_example"), VerificationMethod: openapiclient.cdnCertificateVerificationMethod{}} // CdnRequestCertificateRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SSLApi.RequestCertificate(context.Background(), stackId, siteId, cdnRequestCertificateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSLApi.RequestCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RequestCertificate`: CdnRequestCertificateResponse
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


 **cdnRequestCertificateRequest** | [**CdnRequestCertificateRequest**](CdnRequestCertificateRequest.md) |  | 

### Return type

[**CdnRequestCertificateResponse**](cdnRequestCertificateResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCertificate

> CdnUpdateCertificateResponse UpdateCertificate(ctx, stackId, certificateId).CdnUpdateCertificateRequest(cdnUpdateCertificateRequest).Execute()

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
    cdnUpdateCertificateRequest := openapiclient.cdnUpdateCertificateRequest{Certificate: "Certificate_example", Key: "Key_example", CaBundle: "CaBundle_example"} // CdnUpdateCertificateRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SSLApi.UpdateCertificate(context.Background(), stackId, certificateId, cdnUpdateCertificateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSLApi.UpdateCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCertificate`: CdnUpdateCertificateResponse
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


 **cdnUpdateCertificateRequest** | [**CdnUpdateCertificateRequest**](CdnUpdateCertificateRequest.md) |  | 

### Return type

[**CdnUpdateCertificateResponse**](cdnUpdateCertificateResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSiteCertificateHosts

> UpdateSiteCertificateHosts(ctx, stackId, siteId, certificateId).CdnUpdateSiteCertificateHostsRequest(cdnUpdateSiteCertificateHostsRequest).Execute()

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
    cdnUpdateSiteCertificateHostsRequest := openapiclient.cdnUpdateSiteCertificateHostsRequest{Hosts: []string{"Hosts_example")} // CdnUpdateSiteCertificateHostsRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SSLApi.UpdateSiteCertificateHosts(context.Background(), stackId, siteId, certificateId, cdnUpdateSiteCertificateHostsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSLApi.UpdateSiteCertificateHosts``: %v\n", err)
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
**certificateId** | **string** | A certificate ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSiteCertificateHostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **cdnUpdateSiteCertificateHostsRequest** | [**CdnUpdateSiteCertificateHostsRequest**](CdnUpdateSiteCertificateHostsRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

