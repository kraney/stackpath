# ConditionUrlCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | The URL to match | [optional] 
**ExactMatch** | Pointer to **bool** | Whether to perform an exact match or a string contains match of the URL | [optional] 

## Methods

### NewConditionUrlCondition

`func NewConditionUrlCondition() *ConditionUrlCondition`

NewConditionUrlCondition instantiates a new ConditionUrlCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionUrlConditionWithDefaults

`func NewConditionUrlConditionWithDefaults() *ConditionUrlCondition`

NewConditionUrlConditionWithDefaults instantiates a new ConditionUrlCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *ConditionUrlCondition) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ConditionUrlCondition) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ConditionUrlCondition) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ConditionUrlCondition) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetExactMatch

`func (o *ConditionUrlCondition) GetExactMatch() bool`

GetExactMatch returns the ExactMatch field if non-nil, zero value otherwise.

### GetExactMatchOk

`func (o *ConditionUrlCondition) GetExactMatchOk() (*bool, bool)`

GetExactMatchOk returns a tuple with the ExactMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExactMatch

`func (o *ConditionUrlCondition) SetExactMatch(v bool)`

SetExactMatch sets ExactMatch field to given value.

### HasExactMatch

`func (o *ConditionUrlCondition) HasExactMatch() bool`

HasExactMatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


