# WafUpdateRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The WAF rule&#39;s name | [optional] 
**Description** | **string** | A rule&#39;s description | [optional] 
**Conditions** | [**[]RuleCondition**](RuleCondition.md) | The conditions required for a WAF rule to trigger | [optional] 
**ActionValue** | [**WafRuleAction**](wafRuleAction.md) |  | [optional] 
**Enabled** | **bool** | Whether or not the rule should be enabled | [optional] 
**StatusCode** | [**RuleStatusCode**](RuleStatusCode.md) |  | [optional] 
**ActionDuration** | **string** | How long a rule&#39;s block action will apply to subsequent requests  Durations only apply to rules with block actions. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


