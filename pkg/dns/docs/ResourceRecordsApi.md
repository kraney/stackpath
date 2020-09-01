# \ResourceRecordsApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkCreateOrUpdateZoneRecords**](ResourceRecordsApi.md#BulkCreateOrUpdateZoneRecords) | **Post** /dns/v1/stacks/{stack_id}/zones/{zone_id}/bulk/records | Create or update multiple records
[**BulkDeleteZoneRecords**](ResourceRecordsApi.md#BulkDeleteZoneRecords) | **Post** /dns/v1/stacks/{stack_id}/zones/{zone_id}/bulk/records/delete | Delete multiple records
[**CreateZoneRecord**](ResourceRecordsApi.md#CreateZoneRecord) | **Post** /dns/v1/stacks/{stack_id}/zones/{zone_id}/records | Create a record
[**DeleteZoneRecord**](ResourceRecordsApi.md#DeleteZoneRecord) | **Delete** /dns/v1/stacks/{stack_id}/zones/{zone_id}/records/{zone_record_id} | Delete a record
[**GetZoneRecord**](ResourceRecordsApi.md#GetZoneRecord) | **Get** /dns/v1/stacks/{stack_id}/zones/{zone_id}/records/{zone_record_id} | Get a record
[**GetZoneRecords**](ResourceRecordsApi.md#GetZoneRecords) | **Get** /dns/v1/stacks/{stack_id}/zones/{zone_id}/records | Get all records
[**PatchZoneRecord**](ResourceRecordsApi.md#PatchZoneRecord) | **Patch** /dns/v1/stacks/{stack_id}/zones/{zone_id}/records/{zone_record_id} | Replace a record
[**UpdateZoneRecord**](ResourceRecordsApi.md#UpdateZoneRecord) | **Put** /dns/v1/stacks/{stack_id}/zones/{zone_id}/records/{zone_record_id} | Update a record



## BulkCreateOrUpdateZoneRecords

> ZoneBulkCreateOrUpdateZoneRecordsResponse BulkCreateOrUpdateZoneRecords(ctx, stackId, zoneId).ZoneBulkCreateOrUpdateZoneRecordsRequest(zoneBulkCreateOrUpdateZoneRecordsRequest).Execute()

Create or update multiple records



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
    zoneBulkCreateOrUpdateZoneRecordsRequest := openapiclient.zoneBulkCreateOrUpdateZoneRecordsRequest{Records: []ZoneBulkCreateOrUpdateZoneRecordMessage{openapiclient.zoneBulkCreateOrUpdateZoneRecordMessage{Name: "Name_example", Type: openapiclient.zoneRecordType{}, Ttl: 123, Data: "Data_example", Weight: 123, Labels: map[string]string{ "Key" = "Value" }, Id: "Id_example"})} // ZoneBulkCreateOrUpdateZoneRecordsRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourceRecordsApi.BulkCreateOrUpdateZoneRecords(context.Background(), stackId, zoneId, zoneBulkCreateOrUpdateZoneRecordsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceRecordsApi.BulkCreateOrUpdateZoneRecords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkCreateOrUpdateZoneRecords`: ZoneBulkCreateOrUpdateZoneRecordsResponse
    fmt.Fprintf(os.Stdout, "Response from `ResourceRecordsApi.BulkCreateOrUpdateZoneRecords`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**zoneId** | **string** | A DNS zone ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkCreateOrUpdateZoneRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **zoneBulkCreateOrUpdateZoneRecordsRequest** | [**ZoneBulkCreateOrUpdateZoneRecordsRequest**](ZoneBulkCreateOrUpdateZoneRecordsRequest.md) |  | 

### Return type

[**ZoneBulkCreateOrUpdateZoneRecordsResponse**](zoneBulkCreateOrUpdateZoneRecordsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkDeleteZoneRecords

> BulkDeleteZoneRecords(ctx, stackId, zoneId).ZoneBulkDeleteZoneRecordsMessage(zoneBulkDeleteZoneRecordsMessage).Execute()

Delete multiple records

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
    zoneBulkDeleteZoneRecordsMessage := openapiclient.zoneBulkDeleteZoneRecordsMessage{ZoneRecordIds: []string{"ZoneRecordIds_example")} // ZoneBulkDeleteZoneRecordsMessage | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourceRecordsApi.BulkDeleteZoneRecords(context.Background(), stackId, zoneId, zoneBulkDeleteZoneRecordsMessage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceRecordsApi.BulkDeleteZoneRecords``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBulkDeleteZoneRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **zoneBulkDeleteZoneRecordsMessage** | [**ZoneBulkDeleteZoneRecordsMessage**](ZoneBulkDeleteZoneRecordsMessage.md) |  | 

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


## CreateZoneRecord

> ZoneCreateZoneRecordResponse CreateZoneRecord(ctx, stackId, zoneId).ZoneUpdateZoneRecordMessage(zoneUpdateZoneRecordMessage).Execute()

Create a record

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
    zoneUpdateZoneRecordMessage := openapiclient.zoneUpdateZoneRecordMessage{Name: "Name_example", Type: openapiclient.zoneRecordType{}, Ttl: 123, Data: "Data_example", Weight: 123, Labels: map[string]string{ "Key" = "Value" }} // ZoneUpdateZoneRecordMessage | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourceRecordsApi.CreateZoneRecord(context.Background(), stackId, zoneId, zoneUpdateZoneRecordMessage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceRecordsApi.CreateZoneRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateZoneRecord`: ZoneCreateZoneRecordResponse
    fmt.Fprintf(os.Stdout, "Response from `ResourceRecordsApi.CreateZoneRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**zoneId** | **string** | A DNS zone ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateZoneRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **zoneUpdateZoneRecordMessage** | [**ZoneUpdateZoneRecordMessage**](ZoneUpdateZoneRecordMessage.md) |  | 

### Return type

[**ZoneCreateZoneRecordResponse**](zoneCreateZoneRecordResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteZoneRecord

> DeleteZoneRecord(ctx, stackId, zoneId, zoneRecordId).Execute()

Delete a record

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
    zoneRecordId := "zoneRecordId_example" // string | A DNS resource record ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourceRecordsApi.DeleteZoneRecord(context.Background(), stackId, zoneId, zoneRecordId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceRecordsApi.DeleteZoneRecord``: %v\n", err)
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
**zoneRecordId** | **string** | A DNS resource record ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteZoneRecordRequest struct via the builder pattern


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


## GetZoneRecord

> ZoneGetZoneRecordResponse GetZoneRecord(ctx, stackId, zoneId, zoneRecordId).Execute()

Get a record

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
    zoneRecordId := "zoneRecordId_example" // string | A DNS resource record ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourceRecordsApi.GetZoneRecord(context.Background(), stackId, zoneId, zoneRecordId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceRecordsApi.GetZoneRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetZoneRecord`: ZoneGetZoneRecordResponse
    fmt.Fprintf(os.Stdout, "Response from `ResourceRecordsApi.GetZoneRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**zoneId** | **string** | A DNS zone ID | 
**zoneRecordId** | **string** | A DNS resource record ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetZoneRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ZoneGetZoneRecordResponse**](zoneGetZoneRecordResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetZoneRecords

> ZoneGetZoneRecordsResponse GetZoneRecords(ctx, stackId, zoneId).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).Execute()

Get all records

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
    pageRequestFirst := "pageRequestFirst_example" // string | The number of items desired. (optional)
    pageRequestAfter := "pageRequestAfter_example" // string | The cursor value after which data will be returned. (optional)
    pageRequestFilter := "pageRequestFilter_example" // string | SQL-style constraint filters. (optional)
    pageRequestSortBy := "pageRequestSortBy_example" // string | Sort the response by the given field. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourceRecordsApi.GetZoneRecords(context.Background(), stackId, zoneId).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceRecordsApi.GetZoneRecords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetZoneRecords`: ZoneGetZoneRecordsResponse
    fmt.Fprintf(os.Stdout, "Response from `ResourceRecordsApi.GetZoneRecords`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**zoneId** | **string** | A DNS zone ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetZoneRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageRequestFirst** | **string** | The number of items desired. | 
 **pageRequestAfter** | **string** | The cursor value after which data will be returned. | 
 **pageRequestFilter** | **string** | SQL-style constraint filters. | 
 **pageRequestSortBy** | **string** | Sort the response by the given field. | 

### Return type

[**ZoneGetZoneRecordsResponse**](zoneGetZoneRecordsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchZoneRecord

> ZoneUpdateZoneRecordResponse PatchZoneRecord(ctx, stackId, zoneId, zoneRecordId).ZonePatchZoneRecordMessage(zonePatchZoneRecordMessage).Execute()

Replace a record

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
    zoneRecordId := "zoneRecordId_example" // string | A DNS resource record ID
    zonePatchZoneRecordMessage := openapiclient.zonePatchZoneRecordMessage{Name: "Name_example", Type: , Ttl: 123, Data: "Data_example", Weight: 123} // ZonePatchZoneRecordMessage | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourceRecordsApi.PatchZoneRecord(context.Background(), stackId, zoneId, zoneRecordId, zonePatchZoneRecordMessage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceRecordsApi.PatchZoneRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchZoneRecord`: ZoneUpdateZoneRecordResponse
    fmt.Fprintf(os.Stdout, "Response from `ResourceRecordsApi.PatchZoneRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**zoneId** | **string** | A DNS zone ID | 
**zoneRecordId** | **string** | A DNS resource record ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchZoneRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **zonePatchZoneRecordMessage** | [**ZonePatchZoneRecordMessage**](ZonePatchZoneRecordMessage.md) |  | 

### Return type

[**ZoneUpdateZoneRecordResponse**](zoneUpdateZoneRecordResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateZoneRecord

> ZoneUpdateZoneRecordResponse UpdateZoneRecord(ctx, stackId, zoneId, zoneRecordId).ZoneUpdateZoneRecordMessage(zoneUpdateZoneRecordMessage).Execute()

Update a record

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
    zoneRecordId := "zoneRecordId_example" // string | A DNS resource record ID
    zoneUpdateZoneRecordMessage := openapiclient.zoneUpdateZoneRecordMessage{Name: "Name_example", Type: , Ttl: 123, Data: "Data_example", Weight: 123, Labels: map[string]string{ "Key" = "Value" }} // ZoneUpdateZoneRecordMessage | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourceRecordsApi.UpdateZoneRecord(context.Background(), stackId, zoneId, zoneRecordId, zoneUpdateZoneRecordMessage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceRecordsApi.UpdateZoneRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateZoneRecord`: ZoneUpdateZoneRecordResponse
    fmt.Fprintf(os.Stdout, "Response from `ResourceRecordsApi.UpdateZoneRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**zoneId** | **string** | A DNS zone ID | 
**zoneRecordId** | **string** | A DNS resource record ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateZoneRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **zoneUpdateZoneRecordMessage** | [**ZoneUpdateZoneRecordMessage**](ZoneUpdateZoneRecordMessage.md) |  | 

### Return type

[**ZoneUpdateZoneRecordResponse**](zoneUpdateZoneRecordResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

