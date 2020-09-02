# EventRuleAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleName** | **string** | The name of the rule that triggered the event action | [optional] 
**RuleId** | **string** | The unique ID of the rule that triggered the event action | [optional] 
**ActionTaken** | [**WafRuleAction**](wafRuleAction.md) |  | [optional] 
**Blocked** | **bool** | Whether the requesting client was blocked or not | [optional] 
**Engine** | **string** | The name of the internal WAF engine powering the rule | [optional] 
**RequestType** | [**EventWafRequestType**](EventWafRequestType.md) |  | [optional] 
**Result** | [**RuleActionResultType**](RuleActionResultType.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


