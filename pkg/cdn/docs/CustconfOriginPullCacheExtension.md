# CustconfOriginPullCacheExtension

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | This is used by the API to perform conflict checking | [optional] 
**ExpiredCacheExtension** | Pointer to **int32** | This is the number of seconds to extend an asset&#39;s TTL when the origin is unavailable. The CDN will continue to retry the origin up to the Origin Unavailable Max TTL. | [optional] 
**OriginUnreachableCacheExtension** | Pointer to **int32** | The origin unavailable max TTL value is used by the caching server when your origin is unresponsive or the CDN cannot establish a connection to your origin. Under these conditions, the CDN can continue to serve expired assets from the cache. The value specified in this field establishes a maximum allowable TTL for your expired assets. If your origin connectivity or responsiveness is not corrected within your maximum allowable TTL, the CDN no longer serves your expired assets. | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewCustconfOriginPullCacheExtension

`func NewCustconfOriginPullCacheExtension() *CustconfOriginPullCacheExtension`

NewCustconfOriginPullCacheExtension instantiates a new CustconfOriginPullCacheExtension object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustconfOriginPullCacheExtensionWithDefaults

`func NewCustconfOriginPullCacheExtensionWithDefaults() *CustconfOriginPullCacheExtension`

NewCustconfOriginPullCacheExtensionWithDefaults instantiates a new CustconfOriginPullCacheExtension object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustconfOriginPullCacheExtension) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustconfOriginPullCacheExtension) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustconfOriginPullCacheExtension) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustconfOriginPullCacheExtension) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExpiredCacheExtension

`func (o *CustconfOriginPullCacheExtension) GetExpiredCacheExtension() int32`

GetExpiredCacheExtension returns the ExpiredCacheExtension field if non-nil, zero value otherwise.

### GetExpiredCacheExtensionOk

`func (o *CustconfOriginPullCacheExtension) GetExpiredCacheExtensionOk() (*int32, bool)`

GetExpiredCacheExtensionOk returns a tuple with the ExpiredCacheExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredCacheExtension

`func (o *CustconfOriginPullCacheExtension) SetExpiredCacheExtension(v int32)`

SetExpiredCacheExtension sets ExpiredCacheExtension field to given value.

### HasExpiredCacheExtension

`func (o *CustconfOriginPullCacheExtension) HasExpiredCacheExtension() bool`

HasExpiredCacheExtension returns a boolean if a field has been set.

### GetOriginUnreachableCacheExtension

`func (o *CustconfOriginPullCacheExtension) GetOriginUnreachableCacheExtension() int32`

GetOriginUnreachableCacheExtension returns the OriginUnreachableCacheExtension field if non-nil, zero value otherwise.

### GetOriginUnreachableCacheExtensionOk

`func (o *CustconfOriginPullCacheExtension) GetOriginUnreachableCacheExtensionOk() (*int32, bool)`

GetOriginUnreachableCacheExtensionOk returns a tuple with the OriginUnreachableCacheExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginUnreachableCacheExtension

`func (o *CustconfOriginPullCacheExtension) SetOriginUnreachableCacheExtension(v int32)`

SetOriginUnreachableCacheExtension sets OriginUnreachableCacheExtension field to given value.

### HasOriginUnreachableCacheExtension

`func (o *CustconfOriginPullCacheExtension) HasOriginUnreachableCacheExtension() bool`

HasOriginUnreachableCacheExtension returns a boolean if a field has been set.

### GetEnabled

`func (o *CustconfOriginPullCacheExtension) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustconfOriginPullCacheExtension) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustconfOriginPullCacheExtension) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CustconfOriginPullCacheExtension) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


