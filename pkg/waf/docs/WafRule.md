# WafRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A rule&#39;s ID | [optional] 
**Name** | **string** | A rule&#39;s name | [optional] 
**Description** | **string** | A rule&#39;s description | [optional] 
**Conditions** | [**[]RuleCondition**](RuleCondition.md) | The conditions required for the WAF engine to trigger the rule  Rules may have between 1 and 5 conditions. All conditions must pass for the rule to trigger. | [optional] 
**Action** | [**WafRuleAction**](wafRuleAction.md) |  | [optional] 
**Enabled** | **bool** | Whether or not a WAF rule is enabled | [optional] 
**ActionDuration** | **string** | How long a rule&#39;s block action will apply to subsequent requests  Durations only apply to rules with block actions. | [optional] 
**StatusCode** | [**RuleStatusCode**](RuleStatusCode.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


