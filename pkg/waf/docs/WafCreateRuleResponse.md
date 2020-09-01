# WafCreateRuleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rule** | Pointer to [**WafRule**](wafRule.md) |  | [optional] 

## Methods

### NewWafCreateRuleResponse

`func NewWafCreateRuleResponse() *WafCreateRuleResponse`

NewWafCreateRuleResponse instantiates a new WafCreateRuleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafCreateRuleResponseWithDefaults

`func NewWafCreateRuleResponseWithDefaults() *WafCreateRuleResponse`

NewWafCreateRuleResponseWithDefaults instantiates a new WafCreateRuleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRule

`func (o *WafCreateRuleResponse) GetRule() WafRule`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *WafCreateRuleResponse) GetRuleOk() (*WafRule, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *WafCreateRuleResponse) SetRule(v WafRule)`

SetRule sets Rule field to given value.

### HasRule

`func (o *WafCreateRuleResponse) HasRule() bool`

HasRule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


