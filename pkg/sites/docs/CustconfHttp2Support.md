# CustconfHttp2Support

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | This is used by the API to perform conflict checking | [optional] 
**Enabled** | Pointer to **bool** | Enable support of HTTP2 | [optional] 

## Methods

### NewCustconfHttp2Support

`func NewCustconfHttp2Support() *CustconfHttp2Support`

NewCustconfHttp2Support instantiates a new CustconfHttp2Support object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustconfHttp2SupportWithDefaults

`func NewCustconfHttp2SupportWithDefaults() *CustconfHttp2Support`

NewCustconfHttp2SupportWithDefaults instantiates a new CustconfHttp2Support object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustconfHttp2Support) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustconfHttp2Support) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustconfHttp2Support) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustconfHttp2Support) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEnabled

`func (o *CustconfHttp2Support) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustconfHttp2Support) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustconfHttp2Support) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CustconfHttp2Support) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


