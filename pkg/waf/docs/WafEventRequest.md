# WafEventRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | The requested domain name | [optional] 
**Method** | **string** | The HTTP method that triggered a WAF event | [optional] 
**Scheme** | **string** | The URL scheme that triggered a WAF event | [optional] 
**Uri** | **string** | The full URL that triggered a WAF event | [optional] 
**QueryString** | **string** | The query string portion of a URL that triggered a WAF event | [optional] 
**Headers** | **map[string]string** | A key/value pair of the event&#39;s request headers | [optional] 
**UserAgent** | [**EventRequestUserAgent**](EventRequestUserAgent.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


