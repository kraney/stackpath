# PurgeContentRequestItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | The URL at which to delete content | [optional] 
**Recursive** | **bool** | Whether or not to recursively delete content from the CDN | [optional] 
**InvalidateOnly** | **bool** | Whether or not to mark the asset as expired and re-validate instead of deleting | [optional] 
**PurgeAllDynamic** | **bool** | Whether or not to purge dynamic versions of assets  This is ignored if recursive is true. | [optional] 
**Headers** | **[]string** | A list of HTTP request headers used to construct a cache key to purge content by. These headers must be configured in the site configuration&#39;s _DynamicContent.headerFields_ property. | [optional] 
**PurgeSelector** | [**[]PurgeContentRequestPurgeSelector**](PurgeContentRequestPurgeSelector.md) | A key/value pair definition of content to purge from the CDN | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


