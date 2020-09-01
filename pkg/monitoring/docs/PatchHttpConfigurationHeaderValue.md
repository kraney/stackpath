# PatchHttpConfigurationHeaderValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to [**[]V2Header**](v2Header.md) | The headers value. | [optional] 

## Methods

### NewPatchHttpConfigurationHeaderValue

`func NewPatchHttpConfigurationHeaderValue() *PatchHttpConfigurationHeaderValue`

NewPatchHttpConfigurationHeaderValue instantiates a new PatchHttpConfigurationHeaderValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchHttpConfigurationHeaderValueWithDefaults

`func NewPatchHttpConfigurationHeaderValueWithDefaults() *PatchHttpConfigurationHeaderValue`

NewPatchHttpConfigurationHeaderValueWithDefaults instantiates a new PatchHttpConfigurationHeaderValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *PatchHttpConfigurationHeaderValue) GetValue() []V2Header`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PatchHttpConfigurationHeaderValue) GetValueOk() (*[]V2Header, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PatchHttpConfigurationHeaderValue) SetValue(v []V2Header)`

SetValue sets Value field to given value.

### HasValue

`func (o *PatchHttpConfigurationHeaderValue) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


