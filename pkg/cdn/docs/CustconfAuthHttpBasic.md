# CustconfAuthHttpBasic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | This is used by the API to perform conflict checking | [optional] 
**BindingPoint** | Pointer to **string** | This is a URL to a resource on the authentication server responsible for authentication of users. | [optional] 
**Realm** | Pointer to **string** | This is the authentication realm given back to the user on requests which do not contain the authentication credentials. Browsers typically display this value to the user when the login credentials are requested. | [optional] 
**Ttl** | Pointer to **float32** | This is the number of seconds the authentication session will be cached by the browsers. After this time, browsers will be asked to present the user credentials again for re-authentication. | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewCustconfAuthHttpBasic

`func NewCustconfAuthHttpBasic() *CustconfAuthHttpBasic`

NewCustconfAuthHttpBasic instantiates a new CustconfAuthHttpBasic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustconfAuthHttpBasicWithDefaults

`func NewCustconfAuthHttpBasicWithDefaults() *CustconfAuthHttpBasic`

NewCustconfAuthHttpBasicWithDefaults instantiates a new CustconfAuthHttpBasic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustconfAuthHttpBasic) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustconfAuthHttpBasic) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustconfAuthHttpBasic) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustconfAuthHttpBasic) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBindingPoint

`func (o *CustconfAuthHttpBasic) GetBindingPoint() string`

GetBindingPoint returns the BindingPoint field if non-nil, zero value otherwise.

### GetBindingPointOk

`func (o *CustconfAuthHttpBasic) GetBindingPointOk() (*string, bool)`

GetBindingPointOk returns a tuple with the BindingPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingPoint

`func (o *CustconfAuthHttpBasic) SetBindingPoint(v string)`

SetBindingPoint sets BindingPoint field to given value.

### HasBindingPoint

`func (o *CustconfAuthHttpBasic) HasBindingPoint() bool`

HasBindingPoint returns a boolean if a field has been set.

### GetRealm

`func (o *CustconfAuthHttpBasic) GetRealm() string`

GetRealm returns the Realm field if non-nil, zero value otherwise.

### GetRealmOk

`func (o *CustconfAuthHttpBasic) GetRealmOk() (*string, bool)`

GetRealmOk returns a tuple with the Realm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealm

`func (o *CustconfAuthHttpBasic) SetRealm(v string)`

SetRealm sets Realm field to given value.

### HasRealm

`func (o *CustconfAuthHttpBasic) HasRealm() bool`

HasRealm returns a boolean if a field has been set.

### GetTtl

`func (o *CustconfAuthHttpBasic) GetTtl() float32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *CustconfAuthHttpBasic) GetTtlOk() (*float32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *CustconfAuthHttpBasic) SetTtl(v float32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *CustconfAuthHttpBasic) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetEnabled

`func (o *CustconfAuthHttpBasic) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustconfAuthHttpBasic) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustconfAuthHttpBasic) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CustconfAuthHttpBasic) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


