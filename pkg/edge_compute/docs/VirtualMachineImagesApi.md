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

> V1CreateImageResponse CreateImage(ctx, stackId, imageFamily, imageTag).V1CreateImageRequest(v1CreateImageRequest).Execute()

Create an image

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
    imageFamily := "imageFamily_example" // string | The image's family
    imageTag := "imageTag_example" // string | The image's tag
    v1CreateImageRequest := openapiclient.v1CreateImageRequest{Image: openapiclient.v1Image{StackId: "StackId_example", Id: "Id_example", Family: "Family_example", Tag: "Tag_example", Metadata: openapiclient.v1ImageMetadata{Annotations: map[string]string{ "Key" = "Value" }, Labels: map[string]string{ "Key" = "Value" }, CreatedAt: "TODO", UpdatedAt: "TODO"}, Description: "Description_example", Status: openapiclient.v1ImageStatus{}, ImageSize: "ImageSize_example", VolumeSize: "VolumeSize_example", Deprecation: openapiclient.v1ImageDeprecation{DeprecationDate: "TODO", ObsoletionDate: "TODO", DeletionDate: "TODO"}, Conditions: []V1ImageCondition{openapiclient.v1ImageCondition{Type: "Type_example", Status: openapiclient.v1ImageConditionStatus{}, CheckedAt: "TODO", TransitionedAt: "TODO", Reason: "Reason_example", Message: "Message_example"})}, InstanceVolumeSource: openapiclient.v1ImageSourceInstanceVolume{WorkloadId: "WorkloadId_example", InstanceName: "InstanceName_example", VolumeClaimSlug: "VolumeClaimSlug_example"}} // V1CreateImageRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualMachineImagesApi.CreateImage(context.Background(), stackId, imageFamily, imageTag, v1CreateImageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineImagesApi.CreateImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateImage`: V1CreateImageResponse
    fmt.Fprintf(os.Stdout, "Response from `VirtualMachineImagesApi.CreateImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**imageFamily** | **string** | The image&#39;s family | 
**imageTag** | **string** | The image&#39;s tag | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **v1CreateImageRequest** | [**V1CreateImageRequest**](V1CreateImageRequest.md) |  | 

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

> DeleteImage(ctx, stackId, imageFamily, imageTag).Execute()

Delete an image

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
    imageFamily := "imageFamily_example" // string | An image's family
    imageTag := "imageTag_example" // string | An image's tag

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualMachineImagesApi.DeleteImage(context.Background(), stackId, imageFamily, imageTag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineImagesApi.DeleteImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**imageFamily** | **string** | An image&#39;s family | 
**imageTag** | **string** | An image&#39;s tag | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteImageRequest struct via the builder pattern


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


## DeleteImagesForFamily

> DeleteImagesForFamily(ctx, stackId, imageFamily).Execute()

Delete a family's images

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
    imageFamily := "imageFamily_example" // string | An image family

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualMachineImagesApi.DeleteImagesForFamily(context.Background(), stackId, imageFamily).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineImagesApi.DeleteImagesForFamily``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**imageFamily** | **string** | An image family | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteImagesForFamilyRequest struct via the builder pattern


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


## GetImage

> V1GetImageResponse GetImage(ctx, stackId, imageFamily, imageTag).Execute()

Get an image

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
    imageFamily := "imageFamily_example" // string | An image family
    imageTag := "imageTag_example" // string | An image tag (default to "default")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualMachineImagesApi.GetImage(context.Background(), stackId, imageFamily, imageTag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineImagesApi.GetImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetImage`: V1GetImageResponse
    fmt.Fprintf(os.Stdout, "Response from `VirtualMachineImagesApi.GetImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**imageFamily** | **string** | An image family | 
**imageTag** | **string** | An image tag | [default to &quot;default&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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

> V1GetImagesResponse GetImages(ctx, stackId).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).Deprecated(deprecated).Execute()

Get all images



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
    deprecated := true // bool | If present and true, include deprecated images in the result. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualMachineImagesApi.GetImages(context.Background(), stackId).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).Deprecated(deprecated).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineImagesApi.GetImages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetImages`: V1GetImagesResponse
    fmt.Fprintf(os.Stdout, "Response from `VirtualMachineImagesApi.GetImages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageRequestFirst** | **string** | The number of items desired. | 
 **pageRequestAfter** | **string** | The cursor value after which data will be returned. | 
 **pageRequestFilter** | **string** | SQL-style constraint filters. | 
 **pageRequestSortBy** | **string** | Sort the response by the given field. | 
 **deprecated** | **bool** | If present and true, include deprecated images in the result. | 

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

> V1GetImagesForFamilyResponse GetImagesForFamily(ctx, stackId, imageFamily).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).Deprecated(deprecated).Execute()

Get a family's images



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
    imageFamily := "imageFamily_example" // string | An image family
    pageRequestFirst := "pageRequestFirst_example" // string | The number of items desired. (optional)
    pageRequestAfter := "pageRequestAfter_example" // string | The cursor value after which data will be returned. (optional)
    pageRequestFilter := "pageRequestFilter_example" // string | SQL-style constraint filters. (optional)
    pageRequestSortBy := "pageRequestSortBy_example" // string | Sort the response by the given field. (optional)
    deprecated := true // bool | If present and true, include deprecated images in the result. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualMachineImagesApi.GetImagesForFamily(context.Background(), stackId, imageFamily).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).Deprecated(deprecated).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineImagesApi.GetImagesForFamily``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetImagesForFamily`: V1GetImagesForFamilyResponse
    fmt.Fprintf(os.Stdout, "Response from `VirtualMachineImagesApi.GetImagesForFamily`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**imageFamily** | **string** | An image family | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImagesForFamilyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageRequestFirst** | **string** | The number of items desired. | 
 **pageRequestAfter** | **string** | The cursor value after which data will be returned. | 
 **pageRequestFilter** | **string** | SQL-style constraint filters. | 
 **pageRequestSortBy** | **string** | Sort the response by the given field. | 
 **deprecated** | **bool** | If present and true, include deprecated images in the result. | 

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

> V1UpdateImageResponse UpdateImage(ctx, stackId, imageFamily, imageTag).V1UpdateImageRequest(v1UpdateImageRequest).Execute()

Update an image



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
    imageFamily := "imageFamily_example" // string | An image's family!
    imageTag := "imageTag_example" // string | An image's tag!
    v1UpdateImageRequest := openapiclient.v1UpdateImageRequest{Image: openapiclient.v1Image{StackId: "StackId_example", Id: "Id_example", Family: "Family_example", Tag: "Tag_example", Metadata: openapiclient.v1ImageMetadata{Annotations: map[string]string{ "Key" = "Value" }, Labels: map[string]string{ "Key" = "Value" }, CreatedAt: "TODO", UpdatedAt: "TODO"}, Description: "Description_example", Status: openapiclient.v1ImageStatus{}, ImageSize: "ImageSize_example", VolumeSize: "VolumeSize_example", Deprecation: openapiclient.v1ImageDeprecation{DeprecationDate: "TODO", ObsoletionDate: "TODO", DeletionDate: "TODO"}, Conditions: []V1ImageCondition{openapiclient.v1ImageCondition{Type: "Type_example", Status: openapiclient.v1ImageConditionStatus{}, CheckedAt: "TODO", TransitionedAt: "TODO", Reason: "Reason_example", Message: "Message_example"})}} // V1UpdateImageRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualMachineImagesApi.UpdateImage(context.Background(), stackId, imageFamily, imageTag, v1UpdateImageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineImagesApi.UpdateImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateImage`: V1UpdateImageResponse
    fmt.Fprintf(os.Stdout, "Response from `VirtualMachineImagesApi.UpdateImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**imageFamily** | **string** | An image&#39;s family! | 
**imageTag** | **string** | An image&#39;s tag! | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **v1UpdateImageRequest** | [**V1UpdateImageRequest**](V1UpdateImageRequest.md) |  | 

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

> V1UpdateImageDeprecationResponse UpdateImageDeprecation(ctx, stackId, imageFamily, imageTag).V1ImageDeprecation(v1ImageDeprecation).Execute()

Update deprecation settings



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
    imageFamily := "imageFamily_example" // string | An image family
    imageTag := "imageTag_example" // string | An image tag
    v1ImageDeprecation :=  // V1ImageDeprecation | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualMachineImagesApi.UpdateImageDeprecation(context.Background(), stackId, imageFamily, imageTag, v1ImageDeprecation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineImagesApi.UpdateImageDeprecation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateImageDeprecation`: V1UpdateImageDeprecationResponse
    fmt.Fprintf(os.Stdout, "Response from `VirtualMachineImagesApi.UpdateImageDeprecation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**imageFamily** | **string** | An image family | 
**imageTag** | **string** | An image tag | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateImageDeprecationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **v1ImageDeprecation** | [**V1ImageDeprecation**](V1ImageDeprecation.md) |  | 

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

