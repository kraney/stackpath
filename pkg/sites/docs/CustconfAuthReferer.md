# CustconfAuthReferer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | This is used by the API to perform conflict checking | [optional] 
**Referer** | Pointer to **string** | String of values delimited by a &#39;|&#39; character. This is a list of domains authorized to access content from this path in the container. You may use wildcards to specify multiple websites hosted on the same domain. | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewCustconfAuthReferer

`func NewCustconfAuthReferer() *CustconfAuthReferer`

NewCustconfAuthReferer instantiates a new CustconfAuthReferer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustconfAuthRefererWithDefaults

`func NewCustconfAuthRefererWithDefaults() *CustconfAuthReferer`

NewCustconfAuthRefererWithDefaults instantiates a new CustconfAuthReferer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustconfAuthReferer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustconfAuthReferer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustconfAuthReferer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustconfAuthReferer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReferer

`func (o *CustconfAuthReferer) GetReferer() string`

GetReferer returns the Referer field if non-nil, zero value otherwise.

### GetRefererOk

`func (o *CustconfAuthReferer) GetRefererOk() (*string, bool)`

GetRefererOk returns a tuple with the Referer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferer

`func (o *CustconfAuthReferer) SetReferer(v string)`

SetReferer sets Referer field to given value.

### HasReferer

`func (o *CustconfAuthReferer) HasReferer() bool`

HasReferer returns a boolean if a field has been set.

### GetEnabled

`func (o *CustconfAuthReferer) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustconfAuthReferer) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustconfAuthReferer) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CustconfAuthReferer) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


