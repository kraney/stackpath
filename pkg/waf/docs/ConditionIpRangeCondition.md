# ConditionIpRangeCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LowerBound** | Pointer to **string** | The lower bound IPv4 address to match against | [optional] 
**UpperBound** | Pointer to **string** | The upper bound IPv4 address to match against | [optional] 

## Methods

### NewConditionIpRangeCondition

`func NewConditionIpRangeCondition() *ConditionIpRangeCondition`

NewConditionIpRangeCondition instantiates a new ConditionIpRangeCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionIpRangeConditionWithDefaults

`func NewConditionIpRangeConditionWithDefaults() *ConditionIpRangeCondition`

NewConditionIpRangeConditionWithDefaults instantiates a new ConditionIpRangeCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLowerBound

`func (o *ConditionIpRangeCondition) GetLowerBound() string`

GetLowerBound returns the LowerBound field if non-nil, zero value otherwise.

### GetLowerBoundOk

`func (o *ConditionIpRangeCondition) GetLowerBoundOk() (*string, bool)`

GetLowerBoundOk returns a tuple with the LowerBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerBound

`func (o *ConditionIpRangeCondition) SetLowerBound(v string)`

SetLowerBound sets LowerBound field to given value.

### HasLowerBound

`func (o *ConditionIpRangeCondition) HasLowerBound() bool`

HasLowerBound returns a boolean if a field has been set.

### GetUpperBound

`func (o *ConditionIpRangeCondition) GetUpperBound() string`

GetUpperBound returns the UpperBound field if non-nil, zero value otherwise.

### GetUpperBoundOk

`func (o *ConditionIpRangeCondition) GetUpperBoundOk() (*string, bool)`

GetUpperBoundOk returns a tuple with the UpperBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperBound

`func (o *ConditionIpRangeCondition) SetUpperBound(v string)`

SetUpperBound sets UpperBound field to given value.

### HasUpperBound

`func (o *ConditionIpRangeCondition) HasUpperBound() bool`

HasUpperBound returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


