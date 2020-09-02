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

> ZoneCreateZoneResponse CreateZone(ctx, stackId, zoneCreateZoneMessage)

Create a zone

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**zoneCreateZoneMessage** | [**ZoneCreateZoneMessage**](ZoneCreateZoneMessage.md)|  | 

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

> DeleteZone(ctx, stackId, zoneId)

Delete a zone

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**zoneId** | **string**| A DNS zone ID | 

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

> DisableZone(ctx, stackId, zoneId)

Disable a zone

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**zoneId** | **string**| A DNS zone ID | 

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

> EnableZone(ctx, stackId, zoneId)

Enable a zone

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**zoneId** | **string**| A DNS zone ID | 

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

> ZoneGetNameserversForZoneResponse GetNameserversForZone(ctx, stackId, zoneId)

Get a zone's nameservers

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**zoneId** | **string**| A DNS zone ID | 

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

> ZoneGetZoneResponse GetZone(ctx, stackId, zoneId)

Get a zone

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**zoneId** | **string**| A DNS zone ID | 

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

> ZoneGetZonesResponse GetZones(ctx, stackId, optional)

Get all zones

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
 **optional** | ***GetZonesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetZonesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageRequestFirst** | **optional.String**| The number of items desired. | 
 **pageRequestAfter** | **optional.String**| The cursor value after which data will be returned. | 
 **pageRequestFilter** | **optional.String**| SQL-style constraint filters. | 
 **pageRequestSortBy** | **optional.String**| Sort the response by the given field. | 

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

> ZoneParseRecordsFromZoneFileResponse ParseRecordsFromZoneFile(ctx, stackId, zoneId, zoneParseRecordsFromZoneFileRequest)

Parse a zone file

Parse a BIND zone file. SOA records are not imported. StackPath nameserver records are automatically provided with the zone. Nameserver records are found at the root are not imported.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**zoneId** | **string**| A DNS zone ID | 
**zoneParseRecordsFromZoneFileRequest** | [**ZoneParseRecordsFromZoneFileRequest**](ZoneParseRecordsFromZoneFileRequest.md)|  | 

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

> PushFullZone(ctx, stackId, zoneId)

Publish a zone

Re-push a zone to StackPath's DNS infrastructure

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**zoneId** | **string**| A DNS zone ID | 

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

> ZoneUpdateZoneResponse UpdateZone(ctx, stackId, zoneId, zoneUpdateZoneMessage)

Update a zone

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| A stack ID or slug | 
**zoneId** | **string**| A DNS zone ID | 
**zoneUpdateZoneMessage** | [**ZoneUpdateZoneMessage**](ZoneUpdateZoneMessage.md)|  | 

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

