# CustconfBypassCache

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | This is used by the API to perform conflict checking | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**MethodFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**PathFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**HeaderFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**CookieFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 

## Methods

### NewCustconfBypassCache

`func NewCustconfBypassCache() *CustconfBypassCache`

NewCustconfBypassCache instantiates a new CustconfBypassCache object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustconfBypassCacheWithDefaults

`func NewCustconfBypassCacheWithDefaults() *CustconfBypassCache`

NewCustconfBypassCacheWithDefaults instantiates a new CustconfBypassCache object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustconfBypassCache) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustconfBypassCache) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustconfBypassCache) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustconfBypassCache) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEnabled

`func (o *CustconfBypassCache) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustconfBypassCache) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustconfBypassCache) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CustconfBypassCache) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMethodFilter

`func (o *CustconfBypassCache) GetMethodFilter() string`

GetMethodFilter returns the MethodFilter field if non-nil, zero value otherwise.

### GetMethodFilterOk

`func (o *CustconfBypassCache) GetMethodFilterOk() (*string, bool)`

GetMethodFilterOk returns a tuple with the MethodFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodFilter

`func (o *CustconfBypassCache) SetMethodFilter(v string)`

SetMethodFilter sets MethodFilter field to given value.

### HasMethodFilter

`func (o *CustconfBypassCache) HasMethodFilter() bool`

HasMethodFilter returns a boolean if a field has been set.

### GetPathFilter

`func (o *CustconfBypassCache) GetPathFilter() string`

GetPathFilter returns the PathFilter field if non-nil, zero value otherwise.

### GetPathFilterOk

`func (o *CustconfBypassCache) GetPathFilterOk() (*string, bool)`

GetPathFilterOk returns a tuple with the PathFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathFilter

`func (o *CustconfBypassCache) SetPathFilter(v string)`

SetPathFilter sets PathFilter field to given value.

### HasPathFilter

`func (o *CustconfBypassCache) HasPathFilter() bool`

HasPathFilter returns a boolean if a field has been set.

### GetHeaderFilter

`func (o *CustconfBypassCache) GetHeaderFilter() string`

GetHeaderFilter returns the HeaderFilter field if non-nil, zero value otherwise.

### GetHeaderFilterOk

`func (o *CustconfBypassCache) GetHeaderFilterOk() (*string, bool)`

GetHeaderFilterOk returns a tuple with the HeaderFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderFilter

`func (o *CustconfBypassCache) SetHeaderFilter(v string)`

SetHeaderFilter sets HeaderFilter field to given value.

### HasHeaderFilter

`func (o *CustconfBypassCache) HasHeaderFilter() bool`

HasHeaderFilter returns a boolean if a field has been set.

### GetCookieFilter

`func (o *CustconfBypassCache) GetCookieFilter() string`

GetCookieFilter returns the CookieFilter field if non-nil, zero value otherwise.

### GetCookieFilterOk

`func (o *CustconfBypassCache) GetCookieFilterOk() (*string, bool)`

GetCookieFilterOk returns a tuple with the CookieFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookieFilter

`func (o *CustconfBypassCache) SetCookieFilter(v string)`

SetCookieFilter sets CookieFilter field to given value.

### HasCookieFilter

`func (o *CustconfBypassCache) HasCookieFilter() bool`

HasCookieFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


