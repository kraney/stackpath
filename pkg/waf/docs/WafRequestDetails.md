# WafRequestDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The WAF request&#39;s unique identifier | [optional] 
**Path** | **string** | The path of the requested URL | [optional] 
**Method** | [**WafHttpMethod**](wafHttpMethod.md) |  | [optional] 
**Action** | [**WafRequestAction**](wafRequestAction.md) |  | [optional] 
**RuleId** | **string** | The unique ID of the WAF rule that triggered the action against the WAF request | [optional] 
**RuleName** | **string** | The name of the WAF rule triggered by the request | [optional] 
**UserAgent** | [**WafRequestDetailsUserAgent**](wafRequestDetailsUserAgent.md) |  | [optional] 
**Network** | [**WafRequestDetailsNetwork**](wafRequestDetailsNetwork.md) |  | [optional] 
**RequestTime** | [**time.Time**](time.Time.md) | When the WAF request occurred | [optional] 
**ReferenceId** | **string** | The WAF request&#39;s user-facing identifier  Reference IDs are displayed to the end user when the WAF blocks a request to a site. Please note that a request&#39;s ID and reference ID are different. | [optional] 
**ContentType** | **string** | The content type of the WAF request | [optional] 
**Scheme** | **string** | The HTTP scheme of the WAF request | [optional] 
**HttpStatusCode** | **string** | The HTTP status code of the WAF request | [optional] 
**HttpVersion** | **string** | The HTTP version of the WAF request | [optional] 
**Tags** | [**map[string]RequestDetailsTags**](RequestDetailsTags.md) | Aspects of a WAF request | [optional] 
**ResponseTime** | **string** | The time it took the WAF to process the WAF request | [optional] 
**RequestHeaders** | **map[string]string** | The HTTP request headers of the WAF request | [optional] 
**IncidentId** | **string** | The ID of the of challenge that was generated for the WAF request | [optional] 
**RequestType** | [**RequestDetailsRequestType**](RequestDetailsRequestType.md) |  | [optional] 
**SessionRequestCount** | **string** | The number of WAF requests in the session | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


