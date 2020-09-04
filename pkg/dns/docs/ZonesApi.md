# \ZonesApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateZone**](ZonesApi.md#CreateZone) | **Post** /dns/v1/stacks/{stack_id}/zones | Create a zone
[**DeleteZone**](ZonesApi.md#DeleteZone) | **Delete** /dns/v1/stacks/{stack_id}/zones/{zone_id} | Delete a zone
[**DisableZone**](ZonesApi.md#DisableZone) | **Post** /dns/v1/stacks/{stack_id}/zones/{zone_id}/disable | Disable a zone
[**EnableZone**](ZonesApi.md#EnableZone) | **Post** /dns/v1/stacks/{stack_id}/zones/{zone_id}/enable | Enable a zone
[**GetNameserversForZone**](ZonesApi.md#GetNameserversForZone) | **Get** /dns/v1/stacks/{stack_id}/zones/{zone_id}/discover_nameservers | Get a zone&#39;s nameservers
[**GetZone**](ZonesApi.md#GetZone) | **Get** /dns/v1/stacks/{stack_id}/zones/{zone_id} | Get a zone
[**GetZones**](ZonesApi.md#GetZones) | **Get** /dns/v1/stacks/{stack_id}/zones | Get all zones
[**ParseRecordsFromZoneFile**](ZonesApi.md#ParseRecordsFromZoneFile) | **Post** /dns/v1/stacks/{stack_id}/zones/{zone_id}/parse | Parse a zone file
[**PushFullZone**](ZonesApi.md#PushFullZone) | **Post** /dns/v1/stacks/{stack_id}/zones/{zone_id}/repush | Publish a zone
[**UpdateZone**](ZonesApi.md#UpdateZone) | **Put** /dns/v1/stacks/{stack_id}/zones/{zone_id} | Update a zone



## CreateZone

> ZoneCreateZoneResponse CreateZone(ctx, stackId).ZoneCreateZoneMessage(zoneCreateZoneMessage).Execute()

Create a zone

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
    zoneCreateZoneMessage := openapiclient.zoneCreateZoneMessage{Domain: "Domain_example", Labels: map[string]string{ "Key" = "Value" }, UseApexDomain: false} // ZoneCreateZoneMessage | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ZonesApi.CreateZone(context.Background(), stackId, zoneCreateZoneMessage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZonesApi.CreateZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateZone`: ZoneCreateZoneResponse
    fmt.Fprintf(os.Stdout, "Response from `ZonesApi.CreateZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **zoneCreateZoneMessage** | [**ZoneCreateZoneMessage**](ZoneCreateZoneMessage.md) |  | 

### Return type

[**ZoneCreateZoneResponse**](zoneCreateZoneResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteZone

> DeleteZone(ctx, stackId, zoneId).Execute()

Delete a zone

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
    zoneId := "zoneId_example" // string | A DNS zone ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ZonesApi.DeleteZone(context.Background(), stackId, zoneId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZonesApi.DeleteZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**zoneId** | **string** | A DNS zone ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteZoneRequest struct via the builder pattern


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


## DisableZone

> DisableZone(ctx, stackId, zoneId).Execute()

Disable a zone

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
    zoneId := "zoneId_example" // string | A DNS zone ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ZonesApi.DisableZone(context.Background(), stackId, zoneId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZonesApi.DisableZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**zoneId** | **string** | A DNS zone ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableZoneRequest struct via the builder pattern


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


## EnableZone

> EnableZone(ctx, stackId, zoneId).Execute()

Enable a zone

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
    zoneId := "zoneId_example" // string | A DNS zone ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ZonesApi.EnableZone(context.Background(), stackId, zoneId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZonesApi.EnableZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**zoneId** | **string** | A DNS zone ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableZoneRequest struct via the builder pattern


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


## GetNameserversForZone

> ZoneGetNameserversForZoneResponse GetNameserversForZone(ctx, stackId, zoneId).Execute()

Get a zone's nameservers

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
    zoneId := "zoneId_example" // string | A DNS zone ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ZonesApi.GetNameserversForZone(context.Background(), stackId, zoneId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZonesApi.GetNameserversForZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNameserversForZone`: ZoneGetNameserversForZoneResponse
    fmt.Fprintf(os.Stdout, "Response from `ZonesApi.GetNameserversForZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**zoneId** | **string** | A DNS zone ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNameserversForZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ZoneGetNameserversForZoneResponse**](zoneGetNameserversForZoneResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetZone

> ZoneGetZoneResponse GetZone(ctx, stackId, zoneId).Execute()

Get a zone

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
    zoneId := "zoneId_example" // string | A DNS zone ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ZonesApi.GetZone(context.Background(), stackId, zoneId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZonesApi.GetZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetZone`: ZoneGetZoneResponse
    fmt.Fprintf(os.Stdout, "Response from `ZonesApi.GetZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**zoneId** | **string** | A DNS zone ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ZoneGetZoneResponse**](zoneGetZoneResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetZones

> ZoneGetZonesResponse GetZones(ctx, stackId).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).Execute()

Get all zones

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
    resp, r, err := api_client.ZonesApi.GetZones(context.Background(), stackId).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZonesApi.GetZones``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetZones`: ZoneGetZonesResponse
    fmt.Fprintf(os.Stdout, "Response from `ZonesApi.GetZones`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetZonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageRequestFirst** | **string** | The number of items desired. | 
 **pageRequestAfter** | **string** | The cursor value after which data will be returned. | 
 **pageRequestFilter** | **string** | SQL-style constraint filters. | 
 **pageRequestSortBy** | **string** | Sort the response by the given field. | 

### Return type

[**ZoneGetZonesResponse**](zoneGetZonesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ParseRecordsFromZoneFile

> ZoneParseRecordsFromZoneFileResponse ParseRecordsFromZoneFile(ctx, stackId, zoneId).ZoneParseRecordsFromZoneFileRequest(zoneParseRecordsFromZoneFileRequest).Execute()

Parse a zone file



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
    zoneId := "zoneId_example" // string | A DNS zone ID
    zoneParseRecordsFromZoneFileRequest := openapiclient.zoneParseRecordsFromZoneFileRequest{ZoneFile: 123} // ZoneParseRecordsFromZoneFileRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ZonesApi.ParseRecordsFromZoneFile(context.Background(), stackId, zoneId, zoneParseRecordsFromZoneFileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZonesApi.ParseRecordsFromZoneFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ParseRecordsFromZoneFile`: ZoneParseRecordsFromZoneFileResponse
    fmt.Fprintf(os.Stdout, "Response from `ZonesApi.ParseRecordsFromZoneFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**zoneId** | **string** | A DNS zone ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiParseRecordsFromZoneFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **zoneParseRecordsFromZoneFileRequest** | [**ZoneParseRecordsFromZoneFileRequest**](ZoneParseRecordsFromZoneFileRequest.md) |  | 

### Return type

[**ZoneParseRecordsFromZoneFileResponse**](zoneParseRecordsFromZoneFileResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PushFullZone

> PushFullZone(ctx, stackId, zoneId).Execute()

Publish a zone



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
    zoneId := "zoneId_example" // string | A DNS zone ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ZonesApi.PushFullZone(context.Background(), stackId, zoneId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZonesApi.PushFullZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**zoneId** | **string** | A DNS zone ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPushFullZoneRequest struct via the builder pattern


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


## UpdateZone

> ZoneUpdateZoneResponse UpdateZone(ctx, stackId, zoneId).ZoneUpdateZoneMessage(zoneUpdateZoneMessage).Execute()

Update a zone

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
    zoneId := "zoneId_example" // string | A DNS zone ID
    zoneUpdateZoneMessage := openapiclient.zoneUpdateZoneMessage{Labels: map[string]string{ "Key" = "Value" }} // ZoneUpdateZoneMessage | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ZonesApi.UpdateZone(context.Background(), stackId, zoneId, zoneUpdateZoneMessage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZonesApi.UpdateZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateZone`: ZoneUpdateZoneResponse
    fmt.Fprintf(os.Stdout, "Response from `ZonesApi.UpdateZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**zoneId** | **string** | A DNS zone ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **zoneUpdateZoneMessage** | [**ZoneUpdateZoneMessage**](ZoneUpdateZoneMessage.md) |  | 

### Return type

[**ZoneUpdateZoneResponse**](zoneUpdateZoneResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

