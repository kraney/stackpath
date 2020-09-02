# DeliveryCreateSiteRequestOrigin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | The path the site should request from the origin | [optional] 
**Hostname** | **string** | The origin&#39;s fully-qualified domain name | [optional] 
**Port** | **int32** | The TCP port the site should connect to for HTTP requests | [optional] 
**SecurePort** | **int32** | The TCP port the site should connect to for HTTPS requests | [optional] 
**Http** | [**DeliveryHttpOrigin**](deliveryHTTPOrigin.md) |  | [optional] 
**StackpathStorage** | [**DeliveryStackPathStorageOrigin**](deliveryStackPathStorageOrigin.md) |  | [optional] 
**S3Storage** | [**DeliveryAwss3Origin**](deliveryAWSS3Origin.md) |  | [optional] 
**GoogleCloudStorage** | [**DeliveryGoogleStorageOrigin**](deliveryGoogleStorageOrigin.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


