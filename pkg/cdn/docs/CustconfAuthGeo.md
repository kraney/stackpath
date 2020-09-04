# CustconfAuthGeo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | This is used by the API to perform conflict checking | [optional] 
**Code** | Pointer to [**AuthGeoCodeEnumWrapperValue**](AuthGeoCodeEnumWrapperValue.md) |  | [optional] [default to "UNKNOWN"]
**Values** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. These are the region codes you are targeting for this policy. The values that can be supplied within this field are those that are supported by the MaxMind® GeoIP database. | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewCustconfAuthGeo

`func NewCustconfAuthGeo() *CustconfAuthGeo`

NewCustconfAuthGeo instantiates a new CustconfAuthGeo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustconfAuthGeoWithDefaults

`func NewCustconfAuthGeoWithDefaults() *CustconfAuthGeo`

NewCustconfAuthGeoWithDefaults instantiates a new CustconfAuthGeo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustconfAuthGeo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustconfAuthGeo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustconfAuthGeo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustconfAuthGeo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *CustconfAuthGeo) GetCode() AuthGeoCodeEnumWrapperValue`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CustconfAuthGeo) GetCodeOk() (*AuthGeoCodeEnumWrapperValue, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CustconfAuthGeo) SetCode(v AuthGeoCodeEnumWrapperValue)`

SetCode sets Code field to given value.

### HasCode

`func (o *CustconfAuthGeo) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetValues

`func (o *CustconfAuthGeo) GetValues() string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *CustconfAuthGeo) GetValuesOk() (*string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *CustconfAuthGeo) SetValues(v string)`

SetValues sets Values field to given value.

### HasValues

`func (o *CustconfAuthGeo) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetEnabled

`func (o *CustconfAuthGeo) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustconfAuthGeo) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustconfAuthGeo) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CustconfAuthGeo) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


