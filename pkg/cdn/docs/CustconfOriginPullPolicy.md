# CustconfOriginPullPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | This is used by the API to perform conflict checking | [optional] 
**StatusCodeMatch** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. This is a pattern match expression for each status code this policy applies to. For example, 2*, 3* applies this policy to all 200 and 300 level HTTP responses from your origin. | [optional] 
**ExpirePolicy** | Pointer to [**OriginPullPolicyExpirePolicyEnumWrapperValue**](OriginPullPolicyExpirePolicyEnumWrapperValue.md) |  | [optional] [default to "UNKNOWN"]
**ExpireSeconds** | Pointer to **int32** | This is the expiration time used for assets pulled from your origin. When using Cache-Control headers expiration methods, this value is used if your origin doesn&#39;t return a max-age directive in the Cache-Control HTTP header. Please note that a value of 0 in this fields instructs the caching server to retain assets for as long as possible. | [optional] 
**HonorNoStore** | Pointer to **bool** | This enables the processing of no-store HTTP Cache-Control directives on your container. By enabling this option, responses from your origin containing the no-store directive are not cached. Be aware that requests for non-cacheable assets are always forwarded to your origin and may impose a high request and bandwidth load on your origin. | [optional] 
**HonorNoCache** | Pointer to **bool** | This enables the processing of no-cache HTTP Cache-Control directives on your container. By enabling this option, responses from your origin containing the no-cache directive force the CDN to submit every subsequent request to your origin for validation before serving the asset stored in the cache. | [optional] 
**HonorMustRevalidate** | Pointer to **bool** |  | [optional] 
**NoCacheBehavior** | Pointer to [**OriginPullPolicyNoCacheBehaviorEnumWrapperValue**](OriginPullPolicyNoCacheBehaviorEnumWrapperValue.md) |  | [optional] [default to "UNKNOWN"]
**MaxAgeZeroToNoCache** | Pointer to **bool** | This enables the CDN to apply the no-cache behavior for assets delivered by your origin containing a max-age directive equal to zero. | [optional] 
**MustRevalidateToNoCache** | Pointer to **bool** | This enables the CDN to apply the no-cache behavior for assets delivered by your origin containing the must-revalidate directive. | [optional] 
**BypassCacheIdentifier** | Pointer to **string** | This allows you to define a custom directive that, when used by your origin in the Cache-Control response headers, forces the CDN to proxy the request to the end user without caching the result. | [optional] 
**ForceBypassCache** | Pointer to **bool** | This forces the CDN to not cache any asset pulled from your origin that would otherwise be stored at this location in the cache. Typically this policy is used to prevent 4XX and 5XX response codes from overwriting a file in the cache when used with corresponding Origin Status Code Match setting. If bypass cache behavior is desired for all assets at a scope, Origin Pull Queue Behavior in the Origin Pull Settings also needs to be set to NOCACHE for that scope. | [optional] 
**HttpHeaders** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. This is the list of your origin&#39;s HTTP headers that you want the CDN to cache and deliver to end users. | [optional] 
**HonorPrivate** | Pointer to **bool** | This enables the processing of private HTTP Cache-Control directives on your container. By enabling this option, responses from your origin containing the private directive are not cached. Be aware that requests for non-cacheable assets are always forwarded to your origin and may impose a high request and bandwidth load on your origin. | [optional] 
**HonorSMaxAge** | Pointer to **bool** | This enables the processing of s-maxage HTTP Cache-Control directives on your container. By enabling this option, the s-maxage HTTP Cache-Control directive in the responses from your origin takes precedence over the max-age directive. If both max-age and s-maxage need to be preserved in the client response, the Cache-Control header must be added to the \&quot;Http Header Caching\&quot; setting. | [optional] 
**UpdateHttpHeadersOn304Response** | Pointer to **bool** |  | [optional] 
**DefaultCacheBehavior** | Pointer to [**OriginPullPolicyDefaultCacheBehaviorEnumWrapperValue**](OriginPullPolicyDefaultCacheBehaviorEnumWrapperValue.md) |  | [optional] [default to "UNKNOWN"]
**Enabled** | Pointer to **bool** |  | [optional] 
**MethodFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**PathFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**HeaderFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**ContentTypeFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 

## Methods

### NewCustconfOriginPullPolicy

`func NewCustconfOriginPullPolicy() *CustconfOriginPullPolicy`

NewCustconfOriginPullPolicy instantiates a new CustconfOriginPullPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustconfOriginPullPolicyWithDefaults

`func NewCustconfOriginPullPolicyWithDefaults() *CustconfOriginPullPolicy`

NewCustconfOriginPullPolicyWithDefaults instantiates a new CustconfOriginPullPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustconfOriginPullPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustconfOriginPullPolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustconfOriginPullPolicy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustconfOriginPullPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatusCodeMatch

`func (o *CustconfOriginPullPolicy) GetStatusCodeMatch() string`

GetStatusCodeMatch returns the StatusCodeMatch field if non-nil, zero value otherwise.

### GetStatusCodeMatchOk

`func (o *CustconfOriginPullPolicy) GetStatusCodeMatchOk() (*string, bool)`

GetStatusCodeMatchOk returns a tuple with the StatusCodeMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCodeMatch

`func (o *CustconfOriginPullPolicy) SetStatusCodeMatch(v string)`

SetStatusCodeMatch sets StatusCodeMatch field to given value.

### HasStatusCodeMatch

`func (o *CustconfOriginPullPolicy) HasStatusCodeMatch() bool`

HasStatusCodeMatch returns a boolean if a field has been set.

### GetExpirePolicy

`func (o *CustconfOriginPullPolicy) GetExpirePolicy() OriginPullPolicyExpirePolicyEnumWrapperValue`

GetExpirePolicy returns the ExpirePolicy field if non-nil, zero value otherwise.

### GetExpirePolicyOk

`func (o *CustconfOriginPullPolicy) GetExpirePolicyOk() (*OriginPullPolicyExpirePolicyEnumWrapperValue, bool)`

GetExpirePolicyOk returns a tuple with the ExpirePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirePolicy

`func (o *CustconfOriginPullPolicy) SetExpirePolicy(v OriginPullPolicyExpirePolicyEnumWrapperValue)`

SetExpirePolicy sets ExpirePolicy field to given value.

### HasExpirePolicy

`func (o *CustconfOriginPullPolicy) HasExpirePolicy() bool`

HasExpirePolicy returns a boolean if a field has been set.

### GetExpireSeconds

`func (o *CustconfOriginPullPolicy) GetExpireSeconds() int32`

GetExpireSeconds returns the ExpireSeconds field if non-nil, zero value otherwise.

### GetExpireSecondsOk

`func (o *CustconfOriginPullPolicy) GetExpireSecondsOk() (*int32, bool)`

GetExpireSecondsOk returns a tuple with the ExpireSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireSeconds

`func (o *CustconfOriginPullPolicy) SetExpireSeconds(v int32)`

SetExpireSeconds sets ExpireSeconds field to given value.

### HasExpireSeconds

`func (o *CustconfOriginPullPolicy) HasExpireSeconds() bool`

HasExpireSeconds returns a boolean if a field has been set.

### GetHonorNoStore

`func (o *CustconfOriginPullPolicy) GetHonorNoStore() bool`

GetHonorNoStore returns the HonorNoStore field if non-nil, zero value otherwise.

### GetHonorNoStoreOk

`func (o *CustconfOriginPullPolicy) GetHonorNoStoreOk() (*bool, bool)`

GetHonorNoStoreOk returns a tuple with the HonorNoStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHonorNoStore

`func (o *CustconfOriginPullPolicy) SetHonorNoStore(v bool)`

SetHonorNoStore sets HonorNoStore field to given value.

### HasHonorNoStore

`func (o *CustconfOriginPullPolicy) HasHonorNoStore() bool`

HasHonorNoStore returns a boolean if a field has been set.

### GetHonorNoCache

`func (o *CustconfOriginPullPolicy) GetHonorNoCache() bool`

GetHonorNoCache returns the HonorNoCache field if non-nil, zero value otherwise.

### GetHonorNoCacheOk

`func (o *CustconfOriginPullPolicy) GetHonorNoCacheOk() (*bool, bool)`

GetHonorNoCacheOk returns a tuple with the HonorNoCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHonorNoCache

`func (o *CustconfOriginPullPolicy) SetHonorNoCache(v bool)`

SetHonorNoCache sets HonorNoCache field to given value.

### HasHonorNoCache

`func (o *CustconfOriginPullPolicy) HasHonorNoCache() bool`

HasHonorNoCache returns a boolean if a field has been set.

### GetHonorMustRevalidate

`func (o *CustconfOriginPullPolicy) GetHonorMustRevalidate() bool`

GetHonorMustRevalidate returns the HonorMustRevalidate field if non-nil, zero value otherwise.

### GetHonorMustRevalidateOk

`func (o *CustconfOriginPullPolicy) GetHonorMustRevalidateOk() (*bool, bool)`

GetHonorMustRevalidateOk returns a tuple with the HonorMustRevalidate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHonorMustRevalidate

`func (o *CustconfOriginPullPolicy) SetHonorMustRevalidate(v bool)`

SetHonorMustRevalidate sets HonorMustRevalidate field to given value.

### HasHonorMustRevalidate

`func (o *CustconfOriginPullPolicy) HasHonorMustRevalidate() bool`

HasHonorMustRevalidate returns a boolean if a field has been set.

### GetNoCacheBehavior

`func (o *CustconfOriginPullPolicy) GetNoCacheBehavior() OriginPullPolicyNoCacheBehaviorEnumWrapperValue`

GetNoCacheBehavior returns the NoCacheBehavior field if non-nil, zero value otherwise.

### GetNoCacheBehaviorOk

`func (o *CustconfOriginPullPolicy) GetNoCacheBehaviorOk() (*OriginPullPolicyNoCacheBehaviorEnumWrapperValue, bool)`

GetNoCacheBehaviorOk returns a tuple with the NoCacheBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoCacheBehavior

`func (o *CustconfOriginPullPolicy) SetNoCacheBehavior(v OriginPullPolicyNoCacheBehaviorEnumWrapperValue)`

SetNoCacheBehavior sets NoCacheBehavior field to given value.

### HasNoCacheBehavior

`func (o *CustconfOriginPullPolicy) HasNoCacheBehavior() bool`

HasNoCacheBehavior returns a boolean if a field has been set.

### GetMaxAgeZeroToNoCache

`func (o *CustconfOriginPullPolicy) GetMaxAgeZeroToNoCache() bool`

GetMaxAgeZeroToNoCache returns the MaxAgeZeroToNoCache field if non-nil, zero value otherwise.

### GetMaxAgeZeroToNoCacheOk

`func (o *CustconfOriginPullPolicy) GetMaxAgeZeroToNoCacheOk() (*bool, bool)`

GetMaxAgeZeroToNoCacheOk returns a tuple with the MaxAgeZeroToNoCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAgeZeroToNoCache

`func (o *CustconfOriginPullPolicy) SetMaxAgeZeroToNoCache(v bool)`

SetMaxAgeZeroToNoCache sets MaxAgeZeroToNoCache field to given value.

### HasMaxAgeZeroToNoCache

`func (o *CustconfOriginPullPolicy) HasMaxAgeZeroToNoCache() bool`

HasMaxAgeZeroToNoCache returns a boolean if a field has been set.

### GetMustRevalidateToNoCache

`func (o *CustconfOriginPullPolicy) GetMustRevalidateToNoCache() bool`

GetMustRevalidateToNoCache returns the MustRevalidateToNoCache field if non-nil, zero value otherwise.

### GetMustRevalidateToNoCacheOk

`func (o *CustconfOriginPullPolicy) GetMustRevalidateToNoCacheOk() (*bool, bool)`

GetMustRevalidateToNoCacheOk returns a tuple with the MustRevalidateToNoCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMustRevalidateToNoCache

`func (o *CustconfOriginPullPolicy) SetMustRevalidateToNoCache(v bool)`

SetMustRevalidateToNoCache sets MustRevalidateToNoCache field to given value.

### HasMustRevalidateToNoCache

`func (o *CustconfOriginPullPolicy) HasMustRevalidateToNoCache() bool`

HasMustRevalidateToNoCache returns a boolean if a field has been set.

### GetBypassCacheIdentifier

`func (o *CustconfOriginPullPolicy) GetBypassCacheIdentifier() string`

GetBypassCacheIdentifier returns the BypassCacheIdentifier field if non-nil, zero value otherwise.

### GetBypassCacheIdentifierOk

`func (o *CustconfOriginPullPolicy) GetBypassCacheIdentifierOk() (*string, bool)`

GetBypassCacheIdentifierOk returns a tuple with the BypassCacheIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassCacheIdentifier

`func (o *CustconfOriginPullPolicy) SetBypassCacheIdentifier(v string)`

SetBypassCacheIdentifier sets BypassCacheIdentifier field to given value.

### HasBypassCacheIdentifier

`func (o *CustconfOriginPullPolicy) HasBypassCacheIdentifier() bool`

HasBypassCacheIdentifier returns a boolean if a field has been set.

### GetForceBypassCache

`func (o *CustconfOriginPullPolicy) GetForceBypassCache() bool`

GetForceBypassCache returns the ForceBypassCache field if non-nil, zero value otherwise.

### GetForceBypassCacheOk

`func (o *CustconfOriginPullPolicy) GetForceBypassCacheOk() (*bool, bool)`

GetForceBypassCacheOk returns a tuple with the ForceBypassCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceBypassCache

`func (o *CustconfOriginPullPolicy) SetForceBypassCache(v bool)`

SetForceBypassCache sets ForceBypassCache field to given value.

### HasForceBypassCache

`func (o *CustconfOriginPullPolicy) HasForceBypassCache() bool`

HasForceBypassCache returns a boolean if a field has been set.

### GetHttpHeaders

`func (o *CustconfOriginPullPolicy) GetHttpHeaders() string`

GetHttpHeaders returns the HttpHeaders field if non-nil, zero value otherwise.

### GetHttpHeadersOk

`func (o *CustconfOriginPullPolicy) GetHttpHeadersOk() (*string, bool)`

GetHttpHeadersOk returns a tuple with the HttpHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpHeaders

`func (o *CustconfOriginPullPolicy) SetHttpHeaders(v string)`

SetHttpHeaders sets HttpHeaders field to given value.

### HasHttpHeaders

`func (o *CustconfOriginPullPolicy) HasHttpHeaders() bool`

HasHttpHeaders returns a boolean if a field has been set.

### GetHonorPrivate

`func (o *CustconfOriginPullPolicy) GetHonorPrivate() bool`

GetHonorPrivate returns the HonorPrivate field if non-nil, zero value otherwise.

### GetHonorPrivateOk

`func (o *CustconfOriginPullPolicy) GetHonorPrivateOk() (*bool, bool)`

GetHonorPrivateOk returns a tuple with the HonorPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHonorPrivate

`func (o *CustconfOriginPullPolicy) SetHonorPrivate(v bool)`

SetHonorPrivate sets HonorPrivate field to given value.

### HasHonorPrivate

`func (o *CustconfOriginPullPolicy) HasHonorPrivate() bool`

HasHonorPrivate returns a boolean if a field has been set.

### GetHonorSMaxAge

`func (o *CustconfOriginPullPolicy) GetHonorSMaxAge() bool`

GetHonorSMaxAge returns the HonorSMaxAge field if non-nil, zero value otherwise.

### GetHonorSMaxAgeOk

`func (o *CustconfOriginPullPolicy) GetHonorSMaxAgeOk() (*bool, bool)`

GetHonorSMaxAgeOk returns a tuple with the HonorSMaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHonorSMaxAge

`func (o *CustconfOriginPullPolicy) SetHonorSMaxAge(v bool)`

SetHonorSMaxAge sets HonorSMaxAge field to given value.

### HasHonorSMaxAge

`func (o *CustconfOriginPullPolicy) HasHonorSMaxAge() bool`

HasHonorSMaxAge returns a boolean if a field has been set.

### GetUpdateHttpHeadersOn304Response

`func (o *CustconfOriginPullPolicy) GetUpdateHttpHeadersOn304Response() bool`

GetUpdateHttpHeadersOn304Response returns the UpdateHttpHeadersOn304Response field if non-nil, zero value otherwise.

### GetUpdateHttpHeadersOn304ResponseOk

`func (o *CustconfOriginPullPolicy) GetUpdateHttpHeadersOn304ResponseOk() (*bool, bool)`

GetUpdateHttpHeadersOn304ResponseOk returns a tuple with the UpdateHttpHeadersOn304Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateHttpHeadersOn304Response

`func (o *CustconfOriginPullPolicy) SetUpdateHttpHeadersOn304Response(v bool)`

SetUpdateHttpHeadersOn304Response sets UpdateHttpHeadersOn304Response field to given value.

### HasUpdateHttpHeadersOn304Response

`func (o *CustconfOriginPullPolicy) HasUpdateHttpHeadersOn304Response() bool`

HasUpdateHttpHeadersOn304Response returns a boolean if a field has been set.

### GetDefaultCacheBehavior

`func (o *CustconfOriginPullPolicy) GetDefaultCacheBehavior() OriginPullPolicyDefaultCacheBehaviorEnumWrapperValue`

GetDefaultCacheBehavior returns the DefaultCacheBehavior field if non-nil, zero value otherwise.

### GetDefaultCacheBehaviorOk

`func (o *CustconfOriginPullPolicy) GetDefaultCacheBehaviorOk() (*OriginPullPolicyDefaultCacheBehaviorEnumWrapperValue, bool)`

GetDefaultCacheBehaviorOk returns a tuple with the DefaultCacheBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCacheBehavior

`func (o *CustconfOriginPullPolicy) SetDefaultCacheBehavior(v OriginPullPolicyDefaultCacheBehaviorEnumWrapperValue)`

SetDefaultCacheBehavior sets DefaultCacheBehavior field to given value.

### HasDefaultCacheBehavior

`func (o *CustconfOriginPullPolicy) HasDefaultCacheBehavior() bool`

HasDefaultCacheBehavior returns a boolean if a field has been set.

### GetEnabled

`func (o *CustconfOriginPullPolicy) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustconfOriginPullPolicy) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustconfOriginPullPolicy) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CustconfOriginPullPolicy) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMethodFilter

`func (o *CustconfOriginPullPolicy) GetMethodFilter() string`

GetMethodFilter returns the MethodFilter field if non-nil, zero value otherwise.

### GetMethodFilterOk

`func (o *CustconfOriginPullPolicy) GetMethodFilterOk() (*string, bool)`

GetMethodFilterOk returns a tuple with the MethodFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodFilter

`func (o *CustconfOriginPullPolicy) SetMethodFilter(v string)`

SetMethodFilter sets MethodFilter field to given value.

### HasMethodFilter

`func (o *CustconfOriginPullPolicy) HasMethodFilter() bool`

HasMethodFilter returns a boolean if a field has been set.

### GetPathFilter

`func (o *CustconfOriginPullPolicy) GetPathFilter() string`

GetPathFilter returns the PathFilter field if non-nil, zero value otherwise.

### GetPathFilterOk

`func (o *CustconfOriginPullPolicy) GetPathFilterOk() (*string, bool)`

GetPathFilterOk returns a tuple with the PathFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathFilter

`func (o *CustconfOriginPullPolicy) SetPathFilter(v string)`

SetPathFilter sets PathFilter field to given value.

### HasPathFilter

`func (o *CustconfOriginPullPolicy) HasPathFilter() bool`

HasPathFilter returns a boolean if a field has been set.

### GetHeaderFilter

`func (o *CustconfOriginPullPolicy) GetHeaderFilter() string`

GetHeaderFilter returns the HeaderFilter field if non-nil, zero value otherwise.

### GetHeaderFilterOk

`func (o *CustconfOriginPullPolicy) GetHeaderFilterOk() (*string, bool)`

GetHeaderFilterOk returns a tuple with the HeaderFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderFilter

`func (o *CustconfOriginPullPolicy) SetHeaderFilter(v string)`

SetHeaderFilter sets HeaderFilter field to given value.

### HasHeaderFilter

`func (o *CustconfOriginPullPolicy) HasHeaderFilter() bool`

HasHeaderFilter returns a boolean if a field has been set.

### GetContentTypeFilter

`func (o *CustconfOriginPullPolicy) GetContentTypeFilter() string`

GetContentTypeFilter returns the ContentTypeFilter field if non-nil, zero value otherwise.

### GetContentTypeFilterOk

`func (o *CustconfOriginPullPolicy) GetContentTypeFilterOk() (*string, bool)`

GetContentTypeFilterOk returns a tuple with the ContentTypeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentTypeFilter

`func (o *CustconfOriginPullPolicy) SetContentTypeFilter(v string)`

SetContentTypeFilter sets ContentTypeFilter field to given value.

### HasContentTypeFilter

`func (o *CustconfOriginPullPolicy) HasContentTypeFilter() bool`

HasContentTypeFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


