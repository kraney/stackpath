# V1Image

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StackId** | **string** | The ID of the stack that an image belongs to | [optional] [readonly] 
**Id** | **string** | An image&#39;s unique identifier | [optional] [readonly] 
**Family** | **string** | An image&#39;s family  Families are logical groupings of images | [optional] 
**Tag** | **string** | The image&#39;s tag  Image tags are akin to versions | [optional] 
**Metadata** | [**V1ImageMetadata**](v1ImageMetadata.md) |  | [optional] 
**Description** | **string** | An image&#39;s description  This is optional and may not be more than 1,000 characters | [optional] 
**Status** | [**V1ImageStatus**](v1ImageStatus.md) |  | [optional] 
**ImageSize** | **string** | An image&#39;s archive size in bytes  This is only available on ready images | [optional] [readonly] 
**VolumeSize** | **string** | An image&#39;s volume size in bytes  This is only available on ready images | [optional] [readonly] 
**Deprecation** | [**V1ImageDeprecation**](v1ImageDeprecation.md) |  | [optional] 
**Conditions** | [**[]V1ImageCondition**](v1ImageCondition.md) | Details about an image&#39;s status | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


