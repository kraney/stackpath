# \SSLApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RecheckCertificateVerification**](SSLApi.md#RecheckCertificateVerification) | **Post** /delivery/v1/stacks/{stack_id}/certificates/{certificate_id}/recheck | Re-check certificate verification



## RecheckCertificateVerification

> RecheckCertificateVerification(ctx, stackId, certificateId).Execute()

Re-check certificate verification



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
    resp, r, err := api_client.SSLApi.RecheckCertificateVerification(context.Background(), stackId, certificateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSLApi.RecheckCertificateVerification``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRecheckCertificateVerificationRequest struct via the builder pattern


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

