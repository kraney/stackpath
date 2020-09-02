# PurgeContentRequestPurgeSelector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SelectorType** | [**PurgeContentRequestPurgeSelectorType**](PurgeContentRequestPurgeSelectorType.md) |  | [optional] 
**SelectorName** | **string** | The name of the type of content to purge  For example, the name of the HTTP response header. Names are case sensitive. | [optional] 
**SelectorValue** | **string** | The value of the content to purge  For example, the value of the HTTP response header. Values are case sensitive and may be wild-carded, but cannot match a \&quot;/\&quot;. | [optional] 
**SelectorValueDelimiter** | **string** | The delimiter to separate multiple values with  Defaults to \&quot;,\&quot;. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


