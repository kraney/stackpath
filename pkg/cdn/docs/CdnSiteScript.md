# CdnSiteScript

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A serverless script&#39;s unique identifier | [optional] 
**StackId** | **string** | The ID of the stack to which a serverless script&#39;s site belongs | [optional] 
**SiteId** | **string** | The ID of the site to which a serverless script belongs | [optional] 
**Name** | **string** | A serverless script&#39;s name | [optional] 
**Version** | **string** | A serverless script&#39;s version number  Version numbers start at 1 and are incremented every time the script is updated. | [optional] 
**Code** | **string** | Base64 encoded contents of a serverless script | [optional] 
**Paths** | **[]string** | The URL paths that incoming requests should answer with a serverless script  Every serverless script needs at least one path, and paths must be unique. | [optional] 
**CreatedAt** | [**time.Time**](time.Time.md) | The date a serverless script was created | [optional] 
**UpdatedAt** | [**time.Time**](time.Time.md) | The date a serverless script was last updated | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


