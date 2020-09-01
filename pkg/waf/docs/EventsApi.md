# \EventsApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEvent**](EventsApi.md#GetEvent) | **Get** /waf/v1/stacks/{stack_id}/sites/{site_id}/events/{event_id} | LEGACY: Get an event
[**GetEventStatistics**](EventsApi.md#GetEventStatistics) | **Get** /waf/v1/stacks/{stack_id}/sites/{site_id}/event_stats | LEGACY: Get event statistics
[**SearchEvents**](EventsApi.md#SearchEvents) | **Get** /waf/v1/stacks/{stack_id}/sites/{site_id}/events | LEGACY: Get all events



## GetEvent

> WafGetEventResponse GetEvent(ctx, stackId, siteId, eventId).Execute()

LEGACY: Get an event



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    stackId := "stackId_example" // string | A stack ID or slug
    siteId := "siteId_example" // string | A site ID
    eventId := "eventId_example" // string | A WAF event ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventsApi.GetEvent(context.Background(), stackId, siteId, eventId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.GetEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEvent`: WafGetEventResponse
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.GetEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**siteId** | **string** | A site ID | 
**eventId** | **string** | A WAF event ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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

> WafGetEventStatisticsResponse GetEventStatistics(ctx, stackId, siteId).StartDate(startDate).EndDate(endDate).FilterActionValue(filterActionValue).FilterResultValue(filterResultValue).FilterClientIp(filterClientIp).FilterReferenceId(filterReferenceId).Execute()

LEGACY: Get event statistics



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    stackId := "stackId_example" // string | A stack ID or slug
    siteId := "siteId_example" // string | A site ID
    startDate := Get-Date // time.Time | A lower bound date to search events for. (optional)
    endDate := Get-Date // time.Time | An upper bound date to search events for. (optional)
    filterActionValue := "filterActionValue_example" // string |  (optional) (default to "ANY_ACTION")
    filterResultValue := "filterResultValue_example" // string |  (optional) (default to "ANY_RESULT")
    filterClientIp := "filterClientIp_example" // string | Filter events by client IP address. (optional)
    filterReferenceId := "filterReferenceId_example" // string | Filter events by reference ID. Reference IDs are displayed to the end user when the WAF blocks a request to a site. Please note that an event's ID and reference ID are different. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventsApi.GetEventStatistics(context.Background(), stackId, siteId).StartDate(startDate).EndDate(endDate).FilterActionValue(filterActionValue).FilterResultValue(filterResultValue).FilterClientIp(filterClientIp).FilterReferenceId(filterReferenceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.GetEventStatistics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEventStatistics`: WafGetEventStatisticsResponse
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.GetEventStatistics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**siteId** | **string** | A site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventStatisticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startDate** | **time.Time** | A lower bound date to search events for. | 
 **endDate** | **time.Time** | An upper bound date to search events for. | 
 **filterActionValue** | **string** |  | [default to &quot;ANY_ACTION&quot;]
 **filterResultValue** | **string** |  | [default to &quot;ANY_RESULT&quot;]
 **filterClientIp** | **string** | Filter events by client IP address. | 
 **filterReferenceId** | **string** | Filter events by reference ID. Reference IDs are displayed to the end user when the WAF blocks a request to a site. Please note that an event&#39;s ID and reference ID are different. | 

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

> WafSearchEventsResponse SearchEvents(ctx, stackId, siteId).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).StartDate(startDate).EndDate(endDate).FilterActionValue(filterActionValue).FilterResultValue(filterResultValue).FilterClientIp(filterClientIp).FilterReferenceId(filterReferenceId).SortBy(sortBy).SortOrder(sortOrder).Execute()

LEGACY: Get all events



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    stackId := "stackId_example" // string | A stack ID or slug
    siteId := "siteId_example" // string | A site ID
    pageRequestFirst := "pageRequestFirst_example" // string | The number of items desired. (optional)
    pageRequestAfter := "pageRequestAfter_example" // string | The cursor value after which data will be returned. (optional)
    pageRequestFilter := "pageRequestFilter_example" // string | SQL-style constraint filters. (optional)
    pageRequestSortBy := "pageRequestSortBy_example" // string | Sort the response by the given field. (optional)
    startDate := Get-Date // time.Time | A lower bound date to search events for. (optional)
    endDate := Get-Date // time.Time | An upper bound date to search events for. (optional)
    filterActionValue := "filterActionValue_example" // string |  (optional) (default to "ANY_ACTION")
    filterResultValue := "filterResultValue_example" // string |  (optional) (default to "ANY_RESULT")
    filterClientIp := "filterClientIp_example" // string | Filter events by client IP address. (optional)
    filterReferenceId := "filterReferenceId_example" // string | Filter events by reference ID. Reference IDs are displayed to the end user when the WAF blocks a request to a site. Please note that an event's ID and reference ID are different. (optional)
    sortBy := "sortBy_example" // string |  (optional) (default to "TIMESTAMP")
    sortOrder := "sortOrder_example" // string |  (optional) (default to "ASCENDING")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventsApi.SearchEvents(context.Background(), stackId, siteId).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).StartDate(startDate).EndDate(endDate).FilterActionValue(filterActionValue).FilterResultValue(filterResultValue).FilterClientIp(filterClientIp).FilterReferenceId(filterReferenceId).SortBy(sortBy).SortOrder(sortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.SearchEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchEvents`: WafSearchEventsResponse
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.SearchEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**siteId** | **string** | A site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageRequestFirst** | **string** | The number of items desired. | 
 **pageRequestAfter** | **string** | The cursor value after which data will be returned. | 
 **pageRequestFilter** | **string** | SQL-style constraint filters. | 
 **pageRequestSortBy** | **string** | Sort the response by the given field. | 
 **startDate** | **time.Time** | A lower bound date to search events for. | 
 **endDate** | **time.Time** | An upper bound date to search events for. | 
 **filterActionValue** | **string** |  | [default to &quot;ANY_ACTION&quot;]
 **filterResultValue** | **string** |  | [default to &quot;ANY_RESULT&quot;]
 **filterClientIp** | **string** | Filter events by client IP address. | 
 **filterReferenceId** | **string** | Filter events by reference ID. Reference IDs are displayed to the end user when the WAF blocks a request to a site. Please note that an event&#39;s ID and reference ID are different. | 
 **sortBy** | **string** |  | [default to &quot;TIMESTAMP&quot;]
 **sortOrder** | **string** |  | [default to &quot;ASCENDING&quot;]

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

