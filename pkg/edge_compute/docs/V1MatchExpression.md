# V1MatchExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | The name of the selector to perform a match against | [optional] 
**Operator** | Pointer to **string** | The operation to perform to match a selector  Valid values are \&quot;In\&quot;, \&quot;NotIn\&quot;, \&quot;Exists\&quot;, and \&quot;DoesNotExist\&quot;. | [optional] 
**Values** | Pointer to **[]string** | The values to match in the selector | [optional] 

## Methods

### NewV1MatchExpression

`func NewV1MatchExpression() *V1MatchExpression`

NewV1MatchExpression instantiates a new V1MatchExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1MatchExpressionWithDefaults

`func NewV1MatchExpressionWithDefaults() *V1MatchExpression`

NewV1MatchExpressionWithDefaults instantiates a new V1MatchExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *V1MatchExpression) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *V1MatchExpression) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *V1MatchExpression) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *V1MatchExpression) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetOperator

`func (o *V1MatchExpression) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *V1MatchExpression) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *V1MatchExpression) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *V1MatchExpression) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValues

`func (o *V1MatchExpression) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *V1MatchExpression) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *V1MatchExpression) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *V1MatchExpression) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


