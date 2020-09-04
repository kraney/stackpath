# \MonitorsApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchDeleteMonitors**](MonitorsApi.md#BatchDeleteMonitors) | **Post** /monitoring/v2/stacks/{stack_id}/monitors/batch_delete | Delete multiple monitors
[**CreateMonitor**](MonitorsApi.md#CreateMonitor) | **Post** /monitoring/v2/stacks/{stack_id}/monitors | Create a monitor
[**DeleteMonitor**](MonitorsApi.md#DeleteMonitor) | **Delete** /monitoring/v2/stacks/{stack_id}/monitors/{monitor_id} | Delete a monitor
[**Disable**](MonitorsApi.md#Disable) | **Post** /monitoring/v2/stacks/{stack_id}/disable | Disable all monitors
[**Enable**](MonitorsApi.md#Enable) | **Post** /monitoring/v2/stacks/{stack_id}/enable | Enable all monitors
[**GetMonitor**](MonitorsApi.md#GetMonitor) | **Get** /monitoring/v2/stacks/{stack_id}/monitors/{monitor_id} | Get a monitor
[**GetMonitorErrors**](MonitorsApi.md#GetMonitorErrors) | **Get** /monitoring/v2/stacks/{stack_id}/monitors/{monitor_id}/errors | Get monitoring errors
[**GetMonitorLocations**](MonitorsApi.md#GetMonitorLocations) | **Get** /monitoring/v2/stacks/{stack_id}/monitors/{monitor_id}/locations | Get a monitor&#39;s locations
[**GetMonitors**](MonitorsApi.md#GetMonitors) | **Get** /monitoring/v2/stacks/{stack_id}/monitors | Get all monitors
[**ReplaceMonitor**](MonitorsApi.md#ReplaceMonitor) | **Put** /monitoring/v2/stacks/{stack_id}/monitors/{monitor_id} | Replace a monitor
[**UpdateMonitor**](MonitorsApi.md#UpdateMonitor) | **Patch** /monitoring/v2/stacks/{stack_id}/monitors/{monitor_id} | Update a monitor



## BatchDeleteMonitors

> BatchDeleteMonitors(ctx, stackId).V2BatchDeleteMonitorsRequest(v2BatchDeleteMonitorsRequest).Execute()

Delete multiple monitors

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
    v2BatchDeleteMonitorsRequest := openapiclient.v2BatchDeleteMonitorsRequest{MonitorIds: []string{"MonitorIds_example")} // V2BatchDeleteMonitorsRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorsApi.BatchDeleteMonitors(context.Background(), stackId, v2BatchDeleteMonitorsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.BatchDeleteMonitors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchDeleteMonitorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v2BatchDeleteMonitorsRequest** | [**V2BatchDeleteMonitorsRequest**](V2BatchDeleteMonitorsRequest.md) |  | 

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


## CreateMonitor

> V2CreateMonitorResponse CreateMonitor(ctx, stackId).V2CreateMonitorRequest(v2CreateMonitorRequest).Execute()

Create a monitor

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
    v2CreateMonitorRequest := openapiclient.v2CreateMonitorRequest{Name: "Name_example", Locations: []string{"Locations_example"), Interval: "Interval_example", Timeout: "Timeout_example", IpVersion: openapiclient.v2IpVersion{}, Http: openapiclient.v2HttpConfiguration{Url: "Url_example", Method: "Method_example", Body: 123, Headers: []V2Header{openapiclient.v2Header{Header: "Header_example", Value: "Value_example"}), Basic: openapiclient.v2HttpConfigurationBasicAuth{Username: "Username_example", Password: "Password_example"}, Digest: openapiclient.v2HttpConfigurationDigestAuth{Username: "Username_example", Password: "Password_example"}, ClientCertificate: openapiclient.v2HttpConfigurationClientCertificate{PublicCertificate: "PublicCertificate_example", PrivateKey: "PrivateKey_example", CaBundle: "CaBundle_example"}, ValidateCertificate: false}, Tcp: openapiclient.v2TcpConfiguration{Host: "Host_example", Port: 123, Data: 123}, Enabled: false} // V2CreateMonitorRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorsApi.CreateMonitor(context.Background(), stackId, v2CreateMonitorRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.CreateMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMonitor`: V2CreateMonitorResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitorsApi.CreateMonitor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v2CreateMonitorRequest** | [**V2CreateMonitorRequest**](V2CreateMonitorRequest.md) |  | 

### Return type

[**V2CreateMonitorResponse**](v2CreateMonitorResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMonitor

> DeleteMonitor(ctx, stackId, monitorId).Execute()

Delete a monitor

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorsApi.DeleteMonitor(context.Background(), stackId, monitorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.DeleteMonitor``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteMonitorRequest struct via the builder pattern


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


## Disable

> Disable(ctx, stackId).Execute()

Disable all monitors

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorsApi.Disable(context.Background(), stackId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.Disable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableRequest struct via the builder pattern


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


## Enable

> Enable(ctx, stackId).Execute()

Enable all monitors

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorsApi.Enable(context.Background(), stackId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.Enable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableRequest struct via the builder pattern


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


## GetMonitor

> V2GetMonitorResponse GetMonitor(ctx, stackId, monitorId).Execute()

Get a monitor

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorsApi.GetMonitor(context.Background(), stackId, monitorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.GetMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMonitor`: V2GetMonitorResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitorsApi.GetMonitor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**monitorId** | **string** | A monitor ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**V2GetMonitorResponse**](v2GetMonitorResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMonitorErrors

> V2GetMonitorErrorsResponse GetMonitorErrors(ctx, stackId, monitorId).Execute()

Get monitoring errors

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorsApi.GetMonitorErrors(context.Background(), stackId, monitorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.GetMonitorErrors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMonitorErrors`: V2GetMonitorErrorsResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitorsApi.GetMonitorErrors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**monitorId** | **string** | A monitor ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMonitorErrorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**V2GetMonitorErrorsResponse**](v2GetMonitorErrorsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMonitorLocations

> V2GetMonitorLocationsResponse GetMonitorLocations(ctx, stackId, monitorId).Execute()

Get a monitor's locations

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorsApi.GetMonitorLocations(context.Background(), stackId, monitorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.GetMonitorLocations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMonitorLocations`: V2GetMonitorLocationsResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitorsApi.GetMonitorLocations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**monitorId** | **string** | A monitor ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMonitorLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**V2GetMonitorLocationsResponse**](v2GetMonitorLocationsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMonitors

> V2GetMonitorsResponse GetMonitors(ctx, stackId).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).Execute()

Get all monitors

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
    pageRequestFirst := "pageRequestFirst_example" // string | The number of items desired. (optional)
    pageRequestAfter := "pageRequestAfter_example" // string | The cursor value after which data will be returned. (optional)
    pageRequestFilter := "pageRequestFilter_example" // string | SQL-style constraint filters. (optional)
    pageRequestSortBy := "pageRequestSortBy_example" // string | Sort the response by the given field. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorsApi.GetMonitors(context.Background(), stackId).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.GetMonitors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMonitors`: V2GetMonitorsResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitorsApi.GetMonitors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMonitorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageRequestFirst** | **string** | The number of items desired. | 
 **pageRequestAfter** | **string** | The cursor value after which data will be returned. | 
 **pageRequestFilter** | **string** | SQL-style constraint filters. | 
 **pageRequestSortBy** | **string** | Sort the response by the given field. | 

### Return type

[**V2GetMonitorsResponse**](v2GetMonitorsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceMonitor

> V2ReplaceMonitorResponse ReplaceMonitor(ctx, stackId, monitorId).V2ReplaceMonitorRequest(v2ReplaceMonitorRequest).Execute()

Replace a monitor



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
    v2ReplaceMonitorRequest := openapiclient.v2ReplaceMonitorRequest{Name: "Name_example", Locations: []string{"Locations_example"), Interval: "Interval_example", Timeout: "Timeout_example", IpVersion: openapiclient.v2IpVersion{}, Http: openapiclient.v2HttpConfiguration{Url: "Url_example", Method: "Method_example", Body: 123, Headers: []V2Header{openapiclient.v2Header{Header: "Header_example", Value: "Value_example"}), Basic: openapiclient.v2HttpConfigurationBasicAuth{Username: "Username_example", Password: "Password_example"}, Digest: openapiclient.v2HttpConfigurationDigestAuth{Username: "Username_example", Password: "Password_example"}, ClientCertificate: openapiclient.v2HttpConfigurationClientCertificate{PublicCertificate: "PublicCertificate_example", PrivateKey: "PrivateKey_example", CaBundle: "CaBundle_example"}, ValidateCertificate: false}, Tcp: openapiclient.v2TcpConfiguration{Host: "Host_example", Port: 123, Data: 123}, Enabled: false} // V2ReplaceMonitorRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorsApi.ReplaceMonitor(context.Background(), stackId, monitorId, v2ReplaceMonitorRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.ReplaceMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceMonitor`: V2ReplaceMonitorResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitorsApi.ReplaceMonitor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**monitorId** | **string** | A monitor ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **v2ReplaceMonitorRequest** | [**V2ReplaceMonitorRequest**](V2ReplaceMonitorRequest.md) |  | 

### Return type

[**V2ReplaceMonitorResponse**](v2ReplaceMonitorResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMonitor

> V2UpdateMonitorResponse UpdateMonitor(ctx, stackId, monitorId).V2UpdateMonitorRequest(v2UpdateMonitorRequest).Execute()

Update a monitor

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
    v2UpdateMonitorRequest := openapiclient.v2UpdateMonitorRequest{Name: "Name_example", Locations: openapiclient.UpdateMonitorRequestLocationsValue{Value: []string{"Value_example")}, Interval: "Interval_example", Timeout: "Timeout_example", IpVersion: openapiclient.UpdateMonitorRequestIpVersionValue{Value: }, Http: openapiclient.UpdateMonitorRequestPatchHttpConfiguration{Url: "Url_example", Method: "Method_example", Body: "Body_example", Headers: openapiclient.PatchHttpConfigurationHeaderValue{Value: []V2Header{)}, Basic: openapiclient.UpdateMonitorRequestPatchHttpConfigurationBasicAuth{Username: "Username_example", Password: "Password_example"}, Digest: openapiclient.UpdateMonitorRequestPatchHttpConfigurationDigestAuth{Username: "Username_example", Password: "Password_example"}, ClientCertificate: openapiclient.UpdateMonitorRequestPatchHttpConfigurationClientCertificate{PublicCertificate: "PublicCertificate_example", PrivateKey: "PrivateKey_example", CaBundle: "CaBundle_example"}, ValidateCertificate: false}, Tcp: openapiclient.UpdateMonitorRequestPatchTcpConfiguration{Host: "Host_example", Port: 123, Data: "Data_example"}, Enabled: false} // V2UpdateMonitorRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorsApi.UpdateMonitor(context.Background(), stackId, monitorId, v2UpdateMonitorRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.UpdateMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMonitor`: V2UpdateMonitorResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitorsApi.UpdateMonitor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**monitorId** | **string** | A monitor ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **v2UpdateMonitorRequest** | [**V2UpdateMonitorRequest**](V2UpdateMonitorRequest.md) |  | 

### Return type

[**V2UpdateMonitorResponse**](v2UpdateMonitorResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

