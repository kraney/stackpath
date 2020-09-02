# \EventsApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEvent**](EventsApi.md#GetEvent) | **Get** /waf/v1/stacks/{stack_id}/sites/{site_id}/events/{event_id} | LEGACY: Get an event
[**GetEventStatistics**](EventsApi.md#GetEventStatistics) | **Get** /waf/v1/stacks/{stack_id}/sites/{site_id}/event_stats | LEGACY: Get event statistics
[**SearchEvents**](EventsApi.md#SearchEvents) | **Get** /waf/v1/stacks/{stack_id}/sites/{site_id}/events | LEGACY: Get all events



## GetEvent

> WafGetEventResponse GetEvent(ctx, stackId, siteId, eventId)

LEGACY: Get an event

**Note:** This endpoint is deprecated and will be removed in the future. WAF events will be replaced with requests, which provide more functionality. Please use the [get request](ref:getrequest) and [get request details](ref:getrequestdetails) calls to retrieve WAF requests.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
**eventId** | **string**| A WAF event ID | 

### Return type

[**WafGetEventResponse**](wafGetEventResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventStatistics

> WafGetEventStatisticsResponse GetEventStatistics(ctx, stackId, siteId, optional)

LEGACY: Get event statistics

Event statistics collect the total number of and number of blocked events for a site over a given time frame. Statistics are collected per country of origin, the rules that triggered events, the requesting organization as determined by WHOIS lookup against the client IP address, and by actions taken by the WAF as a result of the event.  **Note:** This endpoint is deprecated and will be removed in the future. WAF events will be replaced with requests, which provide more functionality. A replacement for this call is in development.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
 **optional** | ***GetEventStatisticsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEventStatisticsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startDate** | **optional.Time**| A lower bound date to search events for. | 
 **endDate** | **optional.Time**| An upper bound date to search events for. | 
 **filterActionValue** | **optional.String**|  | [default to ANY_ACTION]
 **filterResultValue** | **optional.String**|  | [default to ANY_RESULT]
 **filterClientIp** | **optional.String**| Filter events by client IP address. | 
 **filterReferenceId** | **optional.String**| Filter events by reference ID. Reference IDs are displayed to the end user when the WAF blocks a request to a site. Please note that an event&#39;s ID and reference ID are different. | 

### Return type

[**WafGetEventStatisticsResponse**](wafGetEventStatisticsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchEvents

> WafSearchEventsResponse SearchEvents(ctx, stackId, siteId, optional)

LEGACY: Get all events

**Note:** This endpoint is deprecated and will be removed in the future. WAF events will be replaced with requests, which provide more functionality. Please use the [get all requests](ref:getrequests) call to search for WAF requests.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**siteId** | **string**| A site ID | 
 **optional** | ***SearchEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SearchEventsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageRequestFirst** | **optional.String**| The number of items desired. | 
 **pageRequestAfter** | **optional.String**| The cursor value after which data will be returned. | 
 **pageRequestFilter** | **optional.String**| SQL-style constraint filters. | 
 **pageRequestSortBy** | **optional.String**| Sort the response by the given field. | 
 **startDate** | **optional.Time**| A lower bound date to search events for. | 
 **endDate** | **optional.Time**| An upper bound date to search events for. | 
 **filterActionValue** | **optional.String**|  | [default to ANY_ACTION]
 **filterResultValue** | **optional.String**|  | [default to ANY_RESULT]
 **filterClientIp** | **optional.String**| Filter events by client IP address. | 
 **filterReferenceId** | **optional.String**| Filter events by reference ID. Reference IDs are displayed to the end user when the WAF blocks a request to a site. Please note that an event&#39;s ID and reference ID are different. | 
 **sortBy** | **optional.String**|  | [default to TIMESTAMP]
 **sortOrder** | **optional.String**|  | [default to ASCENDING]

### Return type

[**WafSearchEventsResponse**](wafSearchEventsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

