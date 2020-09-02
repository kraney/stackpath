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

> ZoneBulkCreateOrUpdateZoneRecordsResponse BulkCreateOrUpdateZoneRecords(ctx, stackId, zoneId, zoneBulkCreateOrUpdateZoneRecordsRequest)

Create or update multiple records

Existing resource records are updated while new records are created

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**zoneId** | **string**| A DNS zone ID | 
**zoneBulkCreateOrUpdateZoneRecordsRequest** | [**ZoneBulkCreateOrUpdateZoneRecordsRequest**](ZoneBulkCreateOrUpdateZoneRecordsRequest.md)|  | 

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

> BulkDeleteZoneRecords(ctx, stackId, zoneId, zoneBulkDeleteZoneRecordsMessage)

Delete multiple records

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**zoneId** | **string**| A DNS zone ID | 
**zoneBulkDeleteZoneRecordsMessage** | [**ZoneBulkDeleteZoneRecordsMessage**](ZoneBulkDeleteZoneRecordsMessage.md)|  | 

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

> ZoneCreateZoneRecordResponse CreateZoneRecord(ctx, stackId, zoneId, zoneUpdateZoneRecordMessage)

Create a record

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**zoneId** | **string**| A DNS zone ID | 
**zoneUpdateZoneRecordMessage** | [**ZoneUpdateZoneRecordMessage**](ZoneUpdateZoneRecordMessage.md)|  | 

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

> DeleteZoneRecord(ctx, stackId, zoneId, zoneRecordId)

Delete a record

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**zoneId** | **string**| A DNS zone ID | 
**zoneRecordId** | **string**| A DNS resource record ID | 

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

> ZoneGetZoneRecordResponse GetZoneRecord(ctx, stackId, zoneId, zoneRecordId)

Get a record

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**zoneId** | **string**| A DNS zone ID | 
**zoneRecordId** | **string**| A DNS resource record ID | 

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

> ZoneGetZoneRecordsResponse GetZoneRecords(ctx, stackId, zoneId, optional)

Get all records

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**zoneId** | **string**| A DNS zone ID | 
 **optional** | ***GetZoneRecordsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetZoneRecordsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageRequestFirst** | **optional.String**| The number of items desired. | 
 **pageRequestAfter** | **optional.String**| The cursor value after which data will be returned. | 
 **pageRequestFilter** | **optional.String**| SQL-style constraint filters. | 
 **pageRequestSortBy** | **optional.String**| Sort the response by the given field. | 

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

> ZoneUpdateZoneRecordResponse PatchZoneRecord(ctx, stackId, zoneId, zoneRecordId, zonePatchZoneRecordMessage)

Replace a record

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**zoneId** | **string**| A DNS zone ID | 
**zoneRecordId** | **string**| A DNS resource record ID | 
**zonePatchZoneRecordMessage** | [**ZonePatchZoneRecordMessage**](ZonePatchZoneRecordMessage.md)|  | 

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

> ZoneUpdateZoneRecordResponse UpdateZoneRecord(ctx, stackId, zoneId, zoneRecordId, zoneUpdateZoneRecordMessage)

Update a record

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**zoneId** | **string**| A DNS zone ID | 
**zoneRecordId** | **string**| A DNS resource record ID | 
**zoneUpdateZoneRecordMessage** | [**ZoneUpdateZoneRecordMessage**](ZoneUpdateZoneRecordMessage.md)|  | 

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

