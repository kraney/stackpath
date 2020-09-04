# WafTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A tag&#39;s name | [optional] 
**Description** | Pointer to **string** | A tag&#39;s human readable description | [optional] 

## Methods

### NewWafTag

`func NewWafTag() *WafTag`

NewWafTag instantiates a new WafTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafTagWithDefaults

`func NewWafTagWithDefaults() *WafTag`

NewWafTagWithDefaults instantiates a new WafTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WafTag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WafTag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WafTag) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WafTag) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *WafTag) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WafTag) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WafTag) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WafTag) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


