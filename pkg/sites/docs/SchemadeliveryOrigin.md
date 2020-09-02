# SchemadeliveryOrigin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | An origin&#39;s unique identifier | [optional] 
**StackId** | **string** | The stack for which the origin belongs to | [optional] 
**Dedicated** | **bool** | Whether or not an origin is dedicated to a CDN site  Dedicated origins cannot be used by any site other than that to which it is dedicated. | [optional] 
**SiteId** | **string** | The ID of the CDN site to which an origin is dedicated | [optional] 
**Http** | [**DeliveryHttpOrigin**](deliveryHTTPOrigin.md) |  | [optional] 
**StackpathStorage** | [**DeliveryStackPathStorageOrigin**](deliveryStackPathStorageOrigin.md) |  | [optional] 
**S3Storage** | [**DeliveryAwss3Origin**](deliveryAWSS3Origin.md) |  | [optional] 
**GoogleCloudStorage** | [**DeliveryGoogleStorageOrigin**](deliveryGoogleStorageOrigin.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


