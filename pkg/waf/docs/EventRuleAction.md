# EventRuleAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleName** | Pointer to **string** | The name of the rule that triggered the event action | [optional] 
**RuleId** | Pointer to **string** | The unique ID of the rule that triggered the event action | [optional] 
**ActionTaken** | Pointer to [**WafRuleAction**](wafRuleAction.md) |  | [optional] [default to "BLOCK"]
**Blocked** | Pointer to **bool** | Whether the requesting client was blocked or not | [optional] 
**Engine** | Pointer to **string** | The name of the internal WAF engine powering the rule | [optional] 
**RequestType** | Pointer to [**EventWafRequestType**](EventWafRequestType.md) |  | [optional] [default to "CHALLENGE"]
**Result** | Pointer to [**RuleActionResultType**](RuleActionResultType.md) |  | [optional] [default to "RESULT_TYPE_UNSPECIFIED"]

## Methods

### NewEventRuleAction

`func NewEventRuleAction() *EventRuleAction`

NewEventRuleAction instantiates a new EventRuleAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventRuleActionWithDefaults

`func NewEventRuleActionWithDefaults() *EventRuleAction`

NewEventRuleActionWithDefaults instantiates a new EventRuleAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleName

`func (o *EventRuleAction) GetRuleName() string`

GetRuleName returns the RuleName field if non-nil, zero value otherwise.

### GetRuleNameOk

`func (o *EventRuleAction) GetRuleNameOk() (*string, bool)`

GetRuleNameOk returns a tuple with the RuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleName

`func (o *EventRuleAction) SetRuleName(v string)`

SetRuleName sets RuleName field to given value.

### HasRuleName

`func (o *EventRuleAction) HasRuleName() bool`

HasRuleName returns a boolean if a field has been set.

### GetRuleId

`func (o *EventRuleAction) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *EventRuleAction) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *EventRuleAction) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *EventRuleAction) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetActionTaken

`func (o *EventRuleAction) GetActionTaken() WafRuleAction`

GetActionTaken returns the ActionTaken field if non-nil, zero value otherwise.

### GetActionTakenOk

`func (o *EventRuleAction) GetActionTakenOk() (*WafRuleAction, bool)`

GetActionTakenOk returns a tuple with the ActionTaken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionTaken

`func (o *EventRuleAction) SetActionTaken(v WafRuleAction)`

SetActionTaken sets ActionTaken field to given value.

### HasActionTaken

`func (o *EventRuleAction) HasActionTaken() bool`

HasActionTaken returns a boolean if a field has been set.

### GetBlocked

`func (o *EventRuleAction) GetBlocked() bool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *EventRuleAction) GetBlockedOk() (*bool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *EventRuleAction) SetBlocked(v bool)`

SetBlocked sets Blocked field to given value.

### HasBlocked

`func (o *EventRuleAction) HasBlocked() bool`

HasBlocked returns a boolean if a field has been set.

### GetEngine

`func (o *EventRuleAction) GetEngine() string`

GetEngine returns the Engine field if non-nil, zero value otherwise.

### GetEngineOk

`func (o *EventRuleAction) GetEngineOk() (*string, bool)`

GetEngineOk returns a tuple with the Engine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngine

`func (o *EventRuleAction) SetEngine(v string)`

SetEngine sets Engine field to given value.

### HasEngine

`func (o *EventRuleAction) HasEngine() bool`

HasEngine returns a boolean if a field has been set.

### GetRequestType

`func (o *EventRuleAction) GetRequestType() EventWafRequestType`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *EventRuleAction) GetRequestTypeOk() (*EventWafRequestType, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *EventRuleAction) SetRequestType(v EventWafRequestType)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *EventRuleAction) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetResult

`func (o *EventRuleAction) GetResult() RuleActionResultType`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *EventRuleAction) GetResultOk() (*RuleActionResultType, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *EventRuleAction) SetResult(v RuleActionResultType)`

SetResult sets Result field to given value.

### HasResult

`func (o *EventRuleAction) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


