# CustconfOriginPullCacheExtension

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | This is used by the API to perform conflict checking | [optional] 
**ExpiredCacheExtension** | **int32** | This is the number of seconds to extend an asset&#39;s TTL when the origin is unavailable. The CDN will continue to retry the origin up to the Origin Unavailable Max TTL. | [optional] 
**OriginUnreachableCacheExtension** | **int32** | The origin unavailable max TTL value is used by the caching server when your origin is unresponsive or the CDN cannot establish a connection to your origin. Under these conditions, the CDN can continue to serve expired assets from the cache. The value specified in this field establishes a maximum allowable TTL for your expired assets. If your origin connectivity or responsiveness is not corrected within your maximum allowable TTL, the CDN no longer serves your expired assets. | [optional] 
**Enabled** | **bool** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


