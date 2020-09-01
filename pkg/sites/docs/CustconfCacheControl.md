# CustconfCacheControl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | This is used by the API to perform conflict checking | [optional] 
**StatusCodeMatch** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**MaxAge** | Pointer to **int32** | The client TTL controls the lifetime of the asset in the browser&#39;s cache. The value entered here will be sent to the browser in the Cache-Control max-age directive for HTTP 1.1 clients and the Expires header for HTTP 1.0 clients. | [optional] 
**MustRevalidate** | Pointer to **bool** | Selecting this option instructs the CDN caching servers to insert the must-revalidate directive on all HTTP responses sent to clients. | [optional] 
**SynchronizeMaxAge** | Pointer to **bool** | Selecting this option allows the CDN to synchronize the Max-Age header it sends to clients with the remaining TTL of the asset in the cache. This allows assets to expire from the browser cache at the same time they expire from the CDN. | [optional] 
**Override** | Pointer to **string** | This allows you to specify a custom Cache-Control header to be used by the CDN on all HTTP responses targeted by this policy. Note: Do not include the header name (Cache-Control) in this field. Only the value of the header should be specified. | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewCustconfCacheControl

`func NewCustconfCacheControl() *CustconfCacheControl`

NewCustconfCacheControl instantiates a new CustconfCacheControl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustconfCacheControlWithDefaults

`func NewCustconfCacheControlWithDefaults() *CustconfCacheControl`

NewCustconfCacheControlWithDefaults instantiates a new CustconfCacheControl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustconfCacheControl) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustconfCacheControl) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustconfCacheControl) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustconfCacheControl) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatusCodeMatch

`func (o *CustconfCacheControl) GetStatusCodeMatch() string`

GetStatusCodeMatch returns the StatusCodeMatch field if non-nil, zero value otherwise.

### GetStatusCodeMatchOk

`func (o *CustconfCacheControl) GetStatusCodeMatchOk() (*string, bool)`

GetStatusCodeMatchOk returns a tuple with the StatusCodeMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCodeMatch

`func (o *CustconfCacheControl) SetStatusCodeMatch(v string)`

SetStatusCodeMatch sets StatusCodeMatch field to given value.

### HasStatusCodeMatch

`func (o *CustconfCacheControl) HasStatusCodeMatch() bool`

HasStatusCodeMatch returns a boolean if a field has been set.

### GetMaxAge

`func (o *CustconfCacheControl) GetMaxAge() int32`

GetMaxAge returns the MaxAge field if non-nil, zero value otherwise.

### GetMaxAgeOk

`func (o *CustconfCacheControl) GetMaxAgeOk() (*int32, bool)`

GetMaxAgeOk returns a tuple with the MaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAge

`func (o *CustconfCacheControl) SetMaxAge(v int32)`

SetMaxAge sets MaxAge field to given value.

### HasMaxAge

`func (o *CustconfCacheControl) HasMaxAge() bool`

HasMaxAge returns a boolean if a field has been set.

### GetMustRevalidate

`func (o *CustconfCacheControl) GetMustRevalidate() bool`

GetMustRevalidate returns the MustRevalidate field if non-nil, zero value otherwise.

### GetMustRevalidateOk

`func (o *CustconfCacheControl) GetMustRevalidateOk() (*bool, bool)`

GetMustRevalidateOk returns a tuple with the MustRevalidate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMustRevalidate

`func (o *CustconfCacheControl) SetMustRevalidate(v bool)`

SetMustRevalidate sets MustRevalidate field to given value.

### HasMustRevalidate

`func (o *CustconfCacheControl) HasMustRevalidate() bool`

HasMustRevalidate returns a boolean if a field has been set.

### GetSynchronizeMaxAge

`func (o *CustconfCacheControl) GetSynchronizeMaxAge() bool`

GetSynchronizeMaxAge returns the SynchronizeMaxAge field if non-nil, zero value otherwise.

### GetSynchronizeMaxAgeOk

`func (o *CustconfCacheControl) GetSynchronizeMaxAgeOk() (*bool, bool)`

GetSynchronizeMaxAgeOk returns a tuple with the SynchronizeMaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynchronizeMaxAge

`func (o *CustconfCacheControl) SetSynchronizeMaxAge(v bool)`

SetSynchronizeMaxAge sets SynchronizeMaxAge field to given value.

### HasSynchronizeMaxAge

`func (o *CustconfCacheControl) HasSynchronizeMaxAge() bool`

HasSynchronizeMaxAge returns a boolean if a field has been set.

### GetOverride

`func (o *CustconfCacheControl) GetOverride() string`

GetOverride returns the Override field if non-nil, zero value otherwise.

### GetOverrideOk

`func (o *CustconfCacheControl) GetOverrideOk() (*string, bool)`

GetOverrideOk returns a tuple with the Override field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverride

`func (o *CustconfCacheControl) SetOverride(v string)`

SetOverride sets Override field to given value.

### HasOverride

`func (o *CustconfCacheControl) HasOverride() bool`

HasOverride returns a boolean if a field has been set.

### GetEnabled

`func (o *CustconfCacheControl) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustconfCacheControl) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustconfCacheControl) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CustconfCacheControl) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


