# \VirtualMachineImagesApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateImage**](VirtualMachineImagesApi.md#CreateImage) | **Post** /workload/v1/stacks/{stack_id}/images/{image_family}/{image_tag} | Create an image
[**DeleteImage**](VirtualMachineImagesApi.md#DeleteImage) | **Delete** /workload/v1/stacks/{stack_id}/images/{image_family}/{image_tag} | Delete an image
[**DeleteImagesForFamily**](VirtualMachineImagesApi.md#DeleteImagesForFamily) | **Delete** /workload/v1/stacks/{stack_id}/images/{image_family} | Delete a family&#39;s images
[**GetImage**](VirtualMachineImagesApi.md#GetImage) | **Get** /workload/v1/stacks/{stack_id}/images/{image_family}/{image_tag} | Get an image
[**GetImages**](VirtualMachineImagesApi.md#GetImages) | **Get** /workload/v1/stacks/{stack_id}/images | Get all images
[**GetImagesForFamily**](VirtualMachineImagesApi.md#GetImagesForFamily) | **Get** /workload/v1/stacks/{stack_id}/images/{image_family} | Get a family&#39;s images
[**UpdateImage**](VirtualMachineImagesApi.md#UpdateImage) | **Patch** /workload/v1/stacks/{stack_id}/images/{image_family}/{image_tag} | Update an image
[**UpdateImageDeprecation**](VirtualMachineImagesApi.md#UpdateImageDeprecation) | **Put** /workload/v1/stacks/{stack_id}/images/{image_family}/{image_tag}/deprecation | Update deprecation settings



## CreateImage

> V1CreateImageResponse CreateImage(ctx, stackId, imageFamily, imageTag, v1CreateImageRequest)

Create an image

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**imageFamily** | **string**| The image&#39;s family | 
**imageTag** | **string**| The image&#39;s tag | 
**v1CreateImageRequest** | [**V1CreateImageRequest**](V1CreateImageRequest.md)|  | 

### Return type

[**V1CreateImageResponse**](v1CreateImageResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteImage

> DeleteImage(ctx, stackId, imageFamily, imageTag)

Delete an image

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**imageFamily** | **string**| An image&#39;s family | 
**imageTag** | **string**| An image&#39;s tag | 

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


## DeleteImagesForFamily

> DeleteImagesForFamily(ctx, stackId, imageFamily)

Delete a family's images

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**imageFamily** | **string**| An image family | 

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


## GetImage

> V1GetImageResponse GetImage(ctx, stackId, imageFamily, imageTag)

Get an image

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**imageFamily** | **string**| An image family | 
**imageTag** | **string**| An image tag | [default to default]

### Return type

[**V1GetImageResponse**](v1GetImageResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImages

> V1GetImagesResponse GetImages(ctx, stackId, optional)

Get all images

Only non-deprecated images are returned by default

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
 **optional** | ***GetImagesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetImagesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageRequestFirst** | **optional.String**| The number of items desired. | 
 **pageRequestAfter** | **optional.String**| The cursor value after which data will be returned. | 
 **pageRequestFilter** | **optional.String**| SQL-style constraint filters. | 
 **pageRequestSortBy** | **optional.String**| Sort the response by the given field. | 
 **deprecated** | **optional.Bool**| If present and true, include deprecated images in the result. | 

### Return type

[**V1GetImagesResponse**](v1GetImagesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImagesForFamily

> V1GetImagesForFamilyResponse GetImagesForFamily(ctx, stackId, imageFamily, optional)

Get a family's images

Only non-deprecated images are returned by default. This will not error but instead return an empty list if the family is not found.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**imageFamily** | **string**| An image family | 
 **optional** | ***GetImagesForFamilyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetImagesForFamilyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageRequestFirst** | **optional.String**| The number of items desired. | 
 **pageRequestAfter** | **optional.String**| The cursor value after which data will be returned. | 
 **pageRequestFilter** | **optional.String**| SQL-style constraint filters. | 
 **pageRequestSortBy** | **optional.String**| Sort the response by the given field. | 
 **deprecated** | **optional.Bool**| If present and true, include deprecated images in the result. | 

### Return type

[**V1GetImagesForFamilyResponse**](v1GetImagesForFamilyResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateImage

> V1UpdateImageResponse UpdateImage(ctx, stackId, imageFamily, imageTag, v1UpdateImageRequest)

Update an image

Only metadata and description can be updated. The metadata, if set, replaces the entire existing metadata set. The tag cannot be \"default\".

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**imageFamily** | **string**| An image&#39;s family! | 
**imageTag** | **string**| An image&#39;s tag! | 
**v1UpdateImageRequest** | [**V1UpdateImageRequest**](V1UpdateImageRequest.md)|  | 

### Return type

[**V1UpdateImageResponse**](v1UpdateImageResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateImageDeprecation

> V1UpdateImageDeprecationResponse UpdateImageDeprecation(ctx, stackId, imageFamily, imageTag, v1ImageDeprecation)

Update deprecation settings

This replaces an image's deprecation settings, so it can also undeprecate an image.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**imageFamily** | **string**| An image family | 
**imageTag** | **string**| An image tag | 
**v1ImageDeprecation** | [**V1ImageDeprecation**](V1ImageDeprecation.md)|  | 

### Return type

[**V1UpdateImageDeprecationResponse**](v1UpdateImageDeprecationResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

