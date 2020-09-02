# WafEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A WAF event&#39;s unique identifier | [optional] 
**ReferenceId** | **string** | An event&#39;s user-facing identifier  Reference IDs are displayed to the end user when the WAF blocks a request to a site. Please note that an event&#39;s ID and reference ID are different. | [optional] 
**EventDate** | [**time.Time**](time.Time.md) | When a WAF event occurred | [optional] 
**Request** | [**WafEventRequest**](wafEventRequest.md) |  | [optional] 
**Action** | [**EventRuleAction**](EventRuleAction.md) |  | [optional] 
**Client** | [**WafEventNetwork**](wafEventNetwork.md) |  | [optional] 
**Count** | **string** | Number of events which matched this | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


