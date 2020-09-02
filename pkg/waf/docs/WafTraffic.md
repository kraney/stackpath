# WafTraffic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | [**time.Time**](time.Time.md) | The date and time traffic metrics were recorded for | [optional] 
**PolicyCount** | [**WafTrafficPolicy**](wafTrafficPolicy.md) |  | [optional] 
**PolicyBlocked** | [**WafTrafficPolicy**](wafTrafficPolicy.md) |  | [optional] 
**RuleCount** | **string** | The number of requests analyzed by rules | [optional] 
**RuleBlocked** | **string** | The number of requests blocked by rules | [optional] 
**PassedToOrigin** | [**TrafficPassedToOrigin**](TrafficPassedToOrigin.md) |  | [optional] 
**Ddos** | **string** | The number of requests handled by the layer 7 DDOS engine | [optional] 
**OriginError4xx** | **string** | The number of HTTP 4xx errors received from the origin | [optional] 
**OriginError5xx** | **string** | The number of HTTP 5xx errors received from the origin | [optional] 
**OriginTimeout** | **string** | The number of request timeouts from the WAF to the origin | [optional] 
**ResponseTime** | **string** | Average response time to the origin in milliseconds | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


