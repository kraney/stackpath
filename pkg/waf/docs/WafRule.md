# WafRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A rule&#39;s ID | [optional] 
**Name** | Pointer to **string** | A rule&#39;s name | [optional] 
**Description** | Pointer to **string** | A rule&#39;s description | [optional] 
**Conditions** | Pointer to [**[]RuleCondition**](RuleCondition.md) | The conditions required for the WAF engine to trigger the rule  Rules may have between 1 and 5 conditions. All conditions must pass for the rule to trigger. | [optional] 
**Action** | Pointer to [**WafRuleAction**](wafRuleAction.md) |  | [optional] [default to "BLOCK"]
**Enabled** | Pointer to **bool** | Whether or not a WAF rule is enabled | [optional] 
**ActionDuration** | Pointer to **string** | How long a rule&#39;s block action will apply to subsequent requests  Durations only apply to rules with block actions. | [optional] 
**StatusCode** | Pointer to [**RuleStatusCode**](RuleStatusCode.md) |  | [optional] [default to "FORBIDDEN_403"]

## Methods

### NewWafRule

`func NewWafRule() *WafRule`

NewWafRule instantiates a new WafRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafRuleWithDefaults

`func NewWafRuleWithDefaults() *WafRule`

NewWafRuleWithDefaults instantiates a new WafRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WafRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WafRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WafRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WafRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *WafRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WafRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WafRule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WafRule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *WafRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WafRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WafRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WafRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetConditions

`func (o *WafRule) GetConditions() []RuleCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *WafRule) GetConditionsOk() (*[]RuleCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *WafRule) SetConditions(v []RuleCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *WafRule) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetAction

`func (o *WafRule) GetAction() WafRuleAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *WafRule) GetActionOk() (*WafRuleAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *WafRule) SetAction(v WafRuleAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *WafRule) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetEnabled

`func (o *WafRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WafRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WafRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WafRule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetActionDuration

`func (o *WafRule) GetActionDuration() string`

GetActionDuration returns the ActionDuration field if non-nil, zero value otherwise.

### GetActionDurationOk

`func (o *WafRule) GetActionDurationOk() (*string, bool)`

GetActionDurationOk returns a tuple with the ActionDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionDuration

`func (o *WafRule) SetActionDuration(v string)`

SetActionDuration sets ActionDuration field to given value.

### HasActionDuration

`func (o *WafRule) HasActionDuration() bool`

HasActionDuration returns a boolean if a field has been set.

### GetStatusCode

`func (o *WafRule) GetStatusCode() RuleStatusCode`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *WafRule) GetStatusCodeOk() (*RuleStatusCode, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *WafRule) SetStatusCode(v RuleStatusCode)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *WafRule) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


