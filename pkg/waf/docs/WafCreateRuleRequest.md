# WafCreateRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The WAF rule&#39;s name | [optional] 
**Description** | Pointer to **string** | A rule&#39;s description | [optional] 
**Conditions** | Pointer to [**[]RuleCondition**](RuleCondition.md) | The conditions required for a WAF rule to trigger | [optional] 
**Action** | Pointer to [**WafRuleAction**](wafRuleAction.md) |  | [optional] [default to "BLOCK"]
**Enabled** | Pointer to **bool** | Whether or not the rule should be enabled on creation | [optional] 
**StatusCode** | Pointer to [**RuleStatusCode**](RuleStatusCode.md) |  | [optional] [default to "FORBIDDEN_403"]
**ActionDuration** | Pointer to **string** | How long a rule&#39;s block action will apply to subsequent requests  Durations only apply to rules with block actions. | [optional] 

## Methods

### NewWafCreateRuleRequest

`func NewWafCreateRuleRequest() *WafCreateRuleRequest`

NewWafCreateRuleRequest instantiates a new WafCreateRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafCreateRuleRequestWithDefaults

`func NewWafCreateRuleRequestWithDefaults() *WafCreateRuleRequest`

NewWafCreateRuleRequestWithDefaults instantiates a new WafCreateRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WafCreateRuleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WafCreateRuleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WafCreateRuleRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WafCreateRuleRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *WafCreateRuleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WafCreateRuleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WafCreateRuleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WafCreateRuleRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetConditions

`func (o *WafCreateRuleRequest) GetConditions() []RuleCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *WafCreateRuleRequest) GetConditionsOk() (*[]RuleCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *WafCreateRuleRequest) SetConditions(v []RuleCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *WafCreateRuleRequest) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetAction

`func (o *WafCreateRuleRequest) GetAction() WafRuleAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *WafCreateRuleRequest) GetActionOk() (*WafRuleAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *WafCreateRuleRequest) SetAction(v WafRuleAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *WafCreateRuleRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetEnabled

`func (o *WafCreateRuleRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WafCreateRuleRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WafCreateRuleRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WafCreateRuleRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetStatusCode

`func (o *WafCreateRuleRequest) GetStatusCode() RuleStatusCode`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *WafCreateRuleRequest) GetStatusCodeOk() (*RuleStatusCode, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *WafCreateRuleRequest) SetStatusCode(v RuleStatusCode)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *WafCreateRuleRequest) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetActionDuration

`func (o *WafCreateRuleRequest) GetActionDuration() string`

GetActionDuration returns the ActionDuration field if non-nil, zero value otherwise.

### GetActionDurationOk

`func (o *WafCreateRuleRequest) GetActionDurationOk() (*string, bool)`

GetActionDurationOk returns a tuple with the ActionDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionDuration

`func (o *WafCreateRuleRequest) SetActionDuration(v string)`

SetActionDuration sets ActionDuration field to given value.

### HasActionDuration

`func (o *WafCreateRuleRequest) HasActionDuration() bool`

HasActionDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


