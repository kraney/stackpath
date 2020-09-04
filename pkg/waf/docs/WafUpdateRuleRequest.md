# WafUpdateRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The WAF rule&#39;s name | [optional] 
**Description** | Pointer to **string** | A rule&#39;s description | [optional] 
**Conditions** | Pointer to [**[]RuleCondition**](RuleCondition.md) | The conditions required for a WAF rule to trigger | [optional] 
**ActionValue** | Pointer to [**WafRuleAction**](wafRuleAction.md) |  | [optional] [default to "BLOCK"]
**Enabled** | Pointer to **bool** | Whether or not the rule should be enabled | [optional] 
**StatusCode** | Pointer to [**RuleStatusCode**](RuleStatusCode.md) |  | [optional] [default to "FORBIDDEN_403"]
**ActionDuration** | Pointer to **string** | How long a rule&#39;s block action will apply to subsequent requests  Durations only apply to rules with block actions. | [optional] 

## Methods

### NewWafUpdateRuleRequest

`func NewWafUpdateRuleRequest() *WafUpdateRuleRequest`

NewWafUpdateRuleRequest instantiates a new WafUpdateRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafUpdateRuleRequestWithDefaults

`func NewWafUpdateRuleRequestWithDefaults() *WafUpdateRuleRequest`

NewWafUpdateRuleRequestWithDefaults instantiates a new WafUpdateRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WafUpdateRuleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WafUpdateRuleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WafUpdateRuleRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WafUpdateRuleRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *WafUpdateRuleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WafUpdateRuleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WafUpdateRuleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WafUpdateRuleRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetConditions

`func (o *WafUpdateRuleRequest) GetConditions() []RuleCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *WafUpdateRuleRequest) GetConditionsOk() (*[]RuleCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *WafUpdateRuleRequest) SetConditions(v []RuleCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *WafUpdateRuleRequest) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetActionValue

`func (o *WafUpdateRuleRequest) GetActionValue() WafRuleAction`

GetActionValue returns the ActionValue field if non-nil, zero value otherwise.

### GetActionValueOk

`func (o *WafUpdateRuleRequest) GetActionValueOk() (*WafRuleAction, bool)`

GetActionValueOk returns a tuple with the ActionValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionValue

`func (o *WafUpdateRuleRequest) SetActionValue(v WafRuleAction)`

SetActionValue sets ActionValue field to given value.

### HasActionValue

`func (o *WafUpdateRuleRequest) HasActionValue() bool`

HasActionValue returns a boolean if a field has been set.

### GetEnabled

`func (o *WafUpdateRuleRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WafUpdateRuleRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WafUpdateRuleRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WafUpdateRuleRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetStatusCode

`func (o *WafUpdateRuleRequest) GetStatusCode() RuleStatusCode`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *WafUpdateRuleRequest) GetStatusCodeOk() (*RuleStatusCode, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *WafUpdateRuleRequest) SetStatusCode(v RuleStatusCode)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *WafUpdateRuleRequest) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetActionDuration

`func (o *WafUpdateRuleRequest) GetActionDuration() string`

GetActionDuration returns the ActionDuration field if non-nil, zero value otherwise.

### GetActionDurationOk

`func (o *WafUpdateRuleRequest) GetActionDurationOk() (*string, bool)`

GetActionDurationOk returns a tuple with the ActionDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionDuration

`func (o *WafUpdateRuleRequest) SetActionDuration(v string)`

SetActionDuration sets ActionDuration field to given value.

### HasActionDuration

`func (o *WafUpdateRuleRequest) HasActionDuration() bool`

HasActionDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


