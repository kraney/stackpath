# TrafficPassedToOrigin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LegitimateRequest** | **string** | The number of requests that passed all policies and rules and were sent to the origin | [optional] 
**Whitelisted** | **string** | The number of requests allowed by whitelist | [optional] 
**AllowedSearchEngine** | **string** | The number of requests allowed by allow search bot policies | [optional] 
**StaticContent** | **string** | The number of requests for static content that was not analyzed by the WAF | [optional] 
**Uncategorized** | **string** | All uncategorized requests that the WAF passed to the origin | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


