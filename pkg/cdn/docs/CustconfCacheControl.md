# CustconfCacheControl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | This is used by the API to perform conflict checking | [optional] 
**StatusCodeMatch** | **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**MaxAge** | **int32** | The client TTL controls the lifetime of the asset in the browser&#39;s cache. The value entered here will be sent to the browser in the Cache-Control max-age directive for HTTP 1.1 clients and the Expires header for HTTP 1.0 clients. | [optional] 
**MustRevalidate** | **bool** | Selecting this option instructs the CDN caching servers to insert the must-revalidate directive on all HTTP responses sent to clients. | [optional] 
**SynchronizeMaxAge** | **bool** | Selecting this option allows the CDN to synchronize the Max-Age header it sends to clients with the remaining TTL of the asset in the cache. This allows assets to expire from the browser cache at the same time they expire from the CDN. | [optional] 
**Override** | **string** | This allows you to specify a custom Cache-Control header to be used by the CDN on all HTTP responses targeted by this policy. Note: Do not include the header name (Cache-Control) in this field. Only the value of the header should be specified. | [optional] 
**Enabled** | **bool** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


