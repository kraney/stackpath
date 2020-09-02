# \SSLApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RecheckCertificateVerification**](SSLApi.md#RecheckCertificateVerification) | **Post** /delivery/v1/stacks/{stack_id}/certificates/{certificate_id}/recheck | Re-check certificate verification



## RecheckCertificateVerification

> RecheckCertificateVerification(ctx, stackId, certificateId)

Re-check certificate verification

Re-check a certificate's verification details while it's being issued. This applies to Stackpath's free certificates.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**certificateId** | **string**| A certificate ID | 

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

