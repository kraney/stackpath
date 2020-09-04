# CdncustconfOrigin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | This is used by the API to perform conflict checking | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**OriginPullHeaders** | Pointer to **string** | String of values delimited by a &#39;|&#39; character. | [optional] 
**OriginCacheHeaders** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 

## Methods

### NewCdncustconfOrigin

`func NewCdncustconfOrigin() *CdncustconfOrigin`

NewCdncustconfOrigin instantiates a new CdncustconfOrigin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdncustconfOriginWithDefaults

`func NewCdncustconfOriginWithDefaults() *CdncustconfOrigin`

NewCdncustconfOriginWithDefaults instantiates a new CdncustconfOrigin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CdncustconfOrigin) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CdncustconfOrigin) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CdncustconfOrigin) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CdncustconfOrigin) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHost

`func (o *CdncustconfOrigin) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *CdncustconfOrigin) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *CdncustconfOrigin) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *CdncustconfOrigin) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetOriginPullHeaders

`func (o *CdncustconfOrigin) GetOriginPullHeaders() string`

GetOriginPullHeaders returns the OriginPullHeaders field if non-nil, zero value otherwise.

### GetOriginPullHeadersOk

`func (o *CdncustconfOrigin) GetOriginPullHeadersOk() (*string, bool)`

GetOriginPullHeadersOk returns a tuple with the OriginPullHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginPullHeaders

`func (o *CdncustconfOrigin) SetOriginPullHeaders(v string)`

SetOriginPullHeaders sets OriginPullHeaders field to given value.

### HasOriginPullHeaders

`func (o *CdncustconfOrigin) HasOriginPullHeaders() bool`

HasOriginPullHeaders returns a boolean if a field has been set.

### GetOriginCacheHeaders

`func (o *CdncustconfOrigin) GetOriginCacheHeaders() string`

GetOriginCacheHeaders returns the OriginCacheHeaders field if non-nil, zero value otherwise.

### GetOriginCacheHeadersOk

`func (o *CdncustconfOrigin) GetOriginCacheHeadersOk() (*string, bool)`

GetOriginCacheHeadersOk returns a tuple with the OriginCacheHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginCacheHeaders

`func (o *CdncustconfOrigin) SetOriginCacheHeaders(v string)`

SetOriginCacheHeaders sets OriginCacheHeaders field to given value.

### HasOriginCacheHeaders

`func (o *CdncustconfOrigin) HasOriginCacheHeaders() bool`

HasOriginCacheHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


