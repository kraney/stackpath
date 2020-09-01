# \AlertConditionsApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchDeleteAlertConditions**](AlertConditionsApi.md#BatchDeleteAlertConditions) | **Post** /monitoring/v2/stacks/{stack_id}/monitors/{monitor_id}/alerts/conditions/batch_delete | Delete multiple alert conditions
[**CreateAlertCondition**](AlertConditionsApi.md#CreateAlertCondition) | **Post** /monitoring/v2/stacks/{stack_id}/monitors/{monitor_id}/alerts/conditions | Create an alert condition
[**DeleteAlertCondition**](AlertConditionsApi.md#DeleteAlertCondition) | **Delete** /monitoring/v2/stacks/{stack_id}/monitors/{monitor_id}/alerts/conditions/{condition_id} | Delete an alert condition
[**DisableAlertCondition**](AlertConditionsApi.md#DisableAlertCondition) | **Post** /monitoring/v2/stacks/{stack_id}/monitors/{monitor_id}/alerts/conditions/{condition_id}/disable | Disable an alert condition
[**EnableAlertCondition**](AlertConditionsApi.md#EnableAlertCondition) | **Post** /monitoring/v2/stacks/{stack_id}/monitors/{monitor_id}/alerts/conditions/{condition_id}/enable | Enable an alert condition
[**GetAlertCondition**](AlertConditionsApi.md#GetAlertCondition) | **Get** /monitoring/v2/stacks/{stack_id}/monitors/{monitor_id}/alerts/conditions/{condition_id} | Get an alert condition
[**GetAlertConditions**](AlertConditionsApi.md#GetAlertConditions) | **Get** /monitoring/v2/stacks/{stack_id}/monitors/{monitor_id}/alerts/conditions | Get all alert conditions
[**UpdateAlertCondition**](AlertConditionsApi.md#UpdateAlertCondition) | **Patch** /monitoring/v2/stacks/{stack_id}/monitors/{monitor_id}/alerts/conditions/{condition_id} | Update an alert condition



## BatchDeleteAlertConditions

> BatchDeleteAlertConditions(ctx, stackId, monitorId).V2BatchDeleteAlertConditionsRequest(v2BatchDeleteAlertConditionsRequest).Execute()

Delete multiple alert conditions

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
    monitorId := "monitorId_example" // string | A monitor ID
    v2BatchDeleteAlertConditionsRequest := openapiclient.v2BatchDeleteAlertConditionsRequest{ConditionIds: []string{"ConditionIds_example")} // V2BatchDeleteAlertConditionsRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertConditionsApi.BatchDeleteAlertConditions(context.Background(), stackId, monitorId, v2BatchDeleteAlertConditionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertConditionsApi.BatchDeleteAlertConditions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**monitorId** | **string** | A monitor ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchDeleteAlertConditionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **v2BatchDeleteAlertConditionsRequest** | [**V2BatchDeleteAlertConditionsRequest**](V2BatchDeleteAlertConditionsRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAlertCondition

> V2CreateAlertConditionResponse CreateAlertCondition(ctx, stackId, monitorId).V2CreateAlertConditionRequest(v2CreateAlertConditionRequest).Execute()

Create an alert condition



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
    monitorId := "monitorId_example" // string | A monitor ID
    v2CreateAlertConditionRequest := openapiclient.v2CreateAlertConditionRequest{Metric: openapiclient.v2AlertConditionMetric{}, Comparator: openapiclient.AlertConditionComparator{}, Value: "Value_example", AlarmDelay: 123, AlarmEvery: 123, AlarmUntilSilenced: false, Enabled: false, MinimumLocations: 123} // V2CreateAlertConditionRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertConditionsApi.CreateAlertCondition(context.Background(), stackId, monitorId, v2CreateAlertConditionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertConditionsApi.CreateAlertCondition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAlertCondition`: V2CreateAlertConditionResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertConditionsApi.CreateAlertCondition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**monitorId** | **string** | A monitor ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAlertConditionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **v2CreateAlertConditionRequest** | [**V2CreateAlertConditionRequest**](V2CreateAlertConditionRequest.md) |  | 

### Return type

[**V2CreateAlertConditionResponse**](v2CreateAlertConditionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAlertCondition

> DeleteAlertCondition(ctx, stackId, monitorId, conditionId).Execute()

Delete an alert condition

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
    monitorId := "monitorId_example" // string | A monitor ID
    conditionId := "conditionId_example" // string | A monitoring alert condition ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertConditionsApi.DeleteAlertCondition(context.Background(), stackId, monitorId, conditionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertConditionsApi.DeleteAlertCondition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**monitorId** | **string** | A monitor ID | 
**conditionId** | **string** | A monitoring alert condition ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAlertConditionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisableAlertCondition

> DisableAlertCondition(ctx, stackId, monitorId, conditionId).Execute()

Disable an alert condition

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
    monitorId := "monitorId_example" // string | A monitor ID
    conditionId := "conditionId_example" // string | A monitoring alert condition ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertConditionsApi.DisableAlertCondition(context.Background(), stackId, monitorId, conditionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertConditionsApi.DisableAlertCondition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**monitorId** | **string** | A monitor ID | 
**conditionId** | **string** | A monitoring alert condition ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableAlertConditionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableAlertCondition

> EnableAlertCondition(ctx, stackId, monitorId, conditionId).Execute()

Enable an alert condition

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
    monitorId := "monitorId_example" // string | A monitor ID
    conditionId := "conditionId_example" // string | A monitoring alert condition ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertConditionsApi.EnableAlertCondition(context.Background(), stackId, monitorId, conditionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertConditionsApi.EnableAlertCondition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**monitorId** | **string** | A monitor ID | 
**conditionId** | **string** | A monitoring alert condition ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableAlertConditionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlertCondition

> V2GetAlertConditionResponse GetAlertCondition(ctx, stackId, monitorId, conditionId).Execute()

Get an alert condition

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
    monitorId := "monitorId_example" // string | A monitor ID
    conditionId := "conditionId_example" // string | A monitoring alert condition ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertConditionsApi.GetAlertCondition(context.Background(), stackId, monitorId, conditionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertConditionsApi.GetAlertCondition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlertCondition`: V2GetAlertConditionResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertConditionsApi.GetAlertCondition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**monitorId** | **string** | A monitor ID | 
**conditionId** | **string** | A monitoring alert condition ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertConditionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**V2GetAlertConditionResponse**](v2GetAlertConditionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlertConditions

> V2GetAlertConditionsResponse GetAlertConditions(ctx, stackId, monitorId).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).Execute()

Get all alert conditions

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
    monitorId := "monitorId_example" // string | A monitor ID
    pageRequestFirst := "pageRequestFirst_example" // string | The number of items desired. (optional)
    pageRequestAfter := "pageRequestAfter_example" // string | The cursor value after which data will be returned. (optional)
    pageRequestFilter := "pageRequestFilter_example" // string | SQL-style constraint filters. (optional)
    pageRequestSortBy := "pageRequestSortBy_example" // string | Sort the response by the given field. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertConditionsApi.GetAlertConditions(context.Background(), stackId, monitorId).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertConditionsApi.GetAlertConditions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlertConditions`: V2GetAlertConditionsResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertConditionsApi.GetAlertConditions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**monitorId** | **string** | A monitor ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertConditionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageRequestFirst** | **string** | The number of items desired. | 
 **pageRequestAfter** | **string** | The cursor value after which data will be returned. | 
 **pageRequestFilter** | **string** | SQL-style constraint filters. | 
 **pageRequestSortBy** | **string** | Sort the response by the given field. | 

### Return type

[**V2GetAlertConditionsResponse**](v2GetAlertConditionsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAlertCondition

> V2UpdateAlertConditionResponse UpdateAlertCondition(ctx, stackId, monitorId, conditionId).V2UpdateAlertConditionRequest(v2UpdateAlertConditionRequest).Execute()

Update an alert condition

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
    monitorId := "monitorId_example" // string | A monitor ID
    conditionId := "conditionId_example" // string | A monitoring alert condition ID
    v2UpdateAlertConditionRequest := openapiclient.v2UpdateAlertConditionRequest{Metric: openapiclient.v2AlertConditionMetric{}, Comparator: openapiclient.AlertConditionComparator{}, Value: "Value_example", AlarmDelay: 123, AlarmEvery: 123, AlarmUntilSilenced: false, Enabled: false, MinimumLocations: 123} // V2UpdateAlertConditionRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertConditionsApi.UpdateAlertCondition(context.Background(), stackId, monitorId, conditionId, v2UpdateAlertConditionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertConditionsApi.UpdateAlertCondition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAlertCondition`: V2UpdateAlertConditionResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertConditionsApi.UpdateAlertCondition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**monitorId** | **string** | A monitor ID | 
**conditionId** | **string** | A monitoring alert condition ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAlertConditionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **v2UpdateAlertConditionRequest** | [**V2UpdateAlertConditionRequest**](V2UpdateAlertConditionRequest.md) |  | 

### Return type

[**V2UpdateAlertConditionResponse**](v2UpdateAlertConditionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

