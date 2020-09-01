# ConditionUserAgentCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserAgent** | Pointer to **string** | The user agent value to match | [optional] 
**ExactMatch** | Pointer to **bool** | Whether to perform an exact match or a string contains match of the user agent | [optional] 

## Methods

### NewConditionUserAgentCondition

`func NewConditionUserAgentCondition() *ConditionUserAgentCondition`

NewConditionUserAgentCondition instantiates a new ConditionUserAgentCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionUserAgentConditionWithDefaults

`func NewConditionUserAgentConditionWithDefaults() *ConditionUserAgentCondition`

NewConditionUserAgentConditionWithDefaults instantiates a new ConditionUserAgentCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserAgent

`func (o *ConditionUserAgentCondition) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *ConditionUserAgentCondition) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *ConditionUserAgentCondition) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *ConditionUserAgentCondition) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetExactMatch

`func (o *ConditionUserAgentCondition) GetExactMatch() bool`

GetExactMatch returns the ExactMatch field if non-nil, zero value otherwise.

### GetExactMatchOk

`func (o *ConditionUserAgentCondition) GetExactMatchOk() (*bool, bool)`

GetExactMatchOk returns a tuple with the ExactMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExactMatch

`func (o *ConditionUserAgentCondition) SetExactMatch(v bool)`

SetExactMatch sets ExactMatch field to given value.

### HasExactMatch

`func (o *ConditionUserAgentCondition) HasExactMatch() bool`

HasExactMatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


