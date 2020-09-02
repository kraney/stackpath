# ConditionRequestRateCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ips** | **[]string** | A list of source IPs that can trigger a request rate condition | [optional] 
**HttpMethods** | [**[]WafHttpMethod**](wafHttpMethod.md) | Possible HTTP request methods that can trigger a request rate condition | [optional] 
**PathPattern** | **string** | A regular expression matching the URL path of the incoming request | [optional] 
**Requests** | **string** | The number of incoming requests over the given time that can trigger a request rate condition | [optional] 
**Time** | **string** | The number of seconds that the WAF measures incoming requests over before triggering a request rate condition | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


