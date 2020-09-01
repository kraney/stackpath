# ConditionResponseHeaderCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Header** | Pointer to **string** | The response header name | [optional] 
**Value** | Pointer to **string** | The response header value | [optional] 
**ExactMatch** | Pointer to **bool** | Whether to perform an exact match of the response header and value | [optional] 

## Methods

### NewConditionResponseHeaderCondition

`func NewConditionResponseHeaderCondition() *ConditionResponseHeaderCondition`

NewConditionResponseHeaderCondition instantiates a new ConditionResponseHeaderCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionResponseHeaderConditionWithDefaults

`func NewConditionResponseHeaderConditionWithDefaults() *ConditionResponseHeaderCondition`

NewConditionResponseHeaderConditionWithDefaults instantiates a new ConditionResponseHeaderCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeader

`func (o *ConditionResponseHeaderCondition) GetHeader() string`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *ConditionResponseHeaderCondition) GetHeaderOk() (*string, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *ConditionResponseHeaderCondition) SetHeader(v string)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *ConditionResponseHeaderCondition) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetValue

`func (o *ConditionResponseHeaderCondition) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ConditionResponseHeaderCondition) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ConditionResponseHeaderCondition) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ConditionResponseHeaderCondition) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetExactMatch

`func (o *ConditionResponseHeaderCondition) GetExactMatch() bool`

GetExactMatch returns the ExactMatch field if non-nil, zero value otherwise.

### GetExactMatchOk

`func (o *ConditionResponseHeaderCondition) GetExactMatchOk() (*bool, bool)`

GetExactMatchOk returns a tuple with the ExactMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExactMatch

`func (o *ConditionResponseHeaderCondition) SetExactMatch(v bool)`

SetExactMatch sets ExactMatch field to given value.

### HasExactMatch

`func (o *ConditionResponseHeaderCondition) HasExactMatch() bool`

HasExactMatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


