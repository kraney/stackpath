# SchemawafRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A request&#39;s unique identifier | [optional] 
**Path** | **string** | The path of the requested URL | [optional] 
**ClientIp** | **string** | The originating IP address | [optional] 
**Method** | [**WafHttpMethod**](wafHttpMethod.md) |  | [optional] 
**RuleName** | **string** | The name of the WAF rule triggered by the request | [optional] 
**Country** | **string** | The ISO 3166-1 alpha-2 code of the country where the request originated from | [optional] 
**Action** | [**WafRequestAction**](wafRequestAction.md) |  | [optional] 
**RuleId** | **string** | The unique ID of the WAF rule that triggered the action against the request | [optional] 
**UserAgent** | **string** | The full user agent string for the request | [optional] 
**UserAgentClient** | **string** | The name of the requesting client, typically the name of the requesting web browser | [optional] 
**Organization** | **string** | The organization that owns the request&#39;s originating IP address according to a WHOIS lookup | [optional] 
**RequestTime** | [**time.Time**](time.Time.md) | When the request occurred | [optional] 
**ReferenceId** | **string** | The request&#39;s user-facing identifier  Reference IDs are displayed to the end user when the WAF blocks a request to a site. Please note that a request&#39;s ID and reference ID are different. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


