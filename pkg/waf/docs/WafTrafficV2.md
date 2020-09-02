# WafTrafficV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | [**time.Time**](time.Time.md) | The date and time traffic metrics were recorded for | [optional] 
**Api** | **string** | The number of requests to URLs designated as API URLs | [optional] 
**Ajax** | **string** | The number of Ajax requests | [optional] 
**Static** | **string** | The number of requests to static assets | [optional] 
**CustomBlocked** | **string** | The number of requests blocked by custom rules | [optional] 
**PolicyBlocked** | **string** | The number of requests blocked by policies | [optional] 
**DdosBlocked** | **string** | The number of requests blocked by layer 7 DDOS protection | [optional] 
**Monitored** | **string** | The number of monitored requests | [optional] 
**CustomWhitelisted** | **string** | The number of requests allowed by custom rules | [optional] 
**PolicyWhitelisted** | **string** | The number of requests from allowed search engines and other allowed bots | [optional] 
**OriginError4xx** | **string** | The number of HTTP 4xx errors received from the origin | [optional] 
**OriginError5xx** | **string** | The number of HTTP 5xx errors received from the origin | [optional] 
**OriginTimeout** | **string** | The number of request timeouts from the WAF to the origin | [optional] 
**Uncategorized** | **string** | The number of uncategorized requests | [optional] 
**PassedToOrigin** | **string** | The number of requests that were passed to the origin | [optional] 
**ResponseTime** | **string** | The average response time to the origin in milliseconds | [optional] 
**Total** | **string** | The total number of requests for the given time period | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


