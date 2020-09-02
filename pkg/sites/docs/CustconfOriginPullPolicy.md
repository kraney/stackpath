# CustconfOriginPullPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | This is used by the API to perform conflict checking | [optional] 
**StatusCodeMatch** | **string** | String of values delimited by a &#39;,&#39; character. This is a pattern match expression for each status code this policy applies to. For example, 2*, 3* applies this policy to all 200 and 300 level HTTP responses from your origin. | [optional] 
**ExpirePolicy** | [**OriginPullPolicyExpirePolicyEnumWrapperValue**](OriginPullPolicyExpirePolicyEnumWrapperValue.md) |  | [optional] 
**ExpireSeconds** | **int32** | This is the expiration time used for assets pulled from your origin. When using Cache-Control headers expiration methods, this value is used if your origin doesn&#39;t return a max-age directive in the Cache-Control HTTP header. Please note that a value of 0 in this fields instructs the caching server to retain assets for as long as possible. | [optional] 
**HonorNoStore** | **bool** | This enables the processing of no-store HTTP Cache-Control directives on your container. By enabling this option, responses from your origin containing the no-store directive are not cached. Be aware that requests for non-cacheable assets are always forwarded to your origin and may impose a high request and bandwidth load on your origin. | [optional] 
**HonorNoCache** | **bool** | This enables the processing of no-cache HTTP Cache-Control directives on your container. By enabling this option, responses from your origin containing the no-cache directive force the CDN to submit every subsequent request to your origin for validation before serving the asset stored in the cache. | [optional] 
**HonorMustRevalidate** | **bool** |  | [optional] 
**NoCacheBehavior** | [**OriginPullPolicyNoCacheBehaviorEnumWrapperValue**](OriginPullPolicyNoCacheBehaviorEnumWrapperValue.md) |  | [optional] 
**MaxAgeZeroToNoCache** | **bool** | This enables the CDN to apply the no-cache behavior for assets delivered by your origin containing a max-age directive equal to zero. | [optional] 
**MustRevalidateToNoCache** | **bool** | This enables the CDN to apply the no-cache behavior for assets delivered by your origin containing the must-revalidate directive. | [optional] 
**BypassCacheIdentifier** | **string** | This allows you to define a custom directive that, when used by your origin in the Cache-Control response headers, forces the CDN to proxy the request to the end user without caching the result. | [optional] 
**ForceBypassCache** | **bool** | This forces the CDN to not cache any asset pulled from your origin that would otherwise be stored at this location in the cache. Typically this policy is used to prevent 4XX and 5XX response codes from overwriting a file in the cache when used with corresponding Origin Status Code Match setting. If bypass cache behavior is desired for all assets at a scope, Origin Pull Queue Behavior in the Origin Pull Settings also needs to be set to NOCACHE for that scope. | [optional] 
**HttpHeaders** | **string** | String of values delimited by a &#39;,&#39; character. This is the list of your origin&#39;s HTTP headers that you want the CDN to cache and deliver to end users. | [optional] 
**HonorPrivate** | **bool** | This enables the processing of private HTTP Cache-Control directives on your container. By enabling this option, responses from your origin containing the private directive are not cached. Be aware that requests for non-cacheable assets are always forwarded to your origin and may impose a high request and bandwidth load on your origin. | [optional] 
**HonorSMaxAge** | **bool** | This enables the processing of s-maxage HTTP Cache-Control directives on your container. By enabling this option, the s-maxage HTTP Cache-Control directive in the responses from your origin takes precedence over the max-age directive. If both max-age and s-maxage need to be preserved in the client response, the Cache-Control header must be added to the \&quot;Http Header Caching\&quot; setting. | [optional] 
**UpdateHttpHeadersOn304Response** | **bool** |  | [optional] 
**DefaultCacheBehavior** | [**OriginPullPolicyDefaultCacheBehaviorEnumWrapperValue**](OriginPullPolicyDefaultCacheBehaviorEnumWrapperValue.md) |  | [optional] 
**Enabled** | **bool** |  | [optional] 
**MethodFilter** | **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**PathFilter** | **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**HeaderFilter** | **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**ContentTypeFilter** | **string** | String of values delimited by a &#39;,&#39; character. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


