# SchemacdnOrigin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | An origin&#39;s unique identifier | [optional] 
**Path** | **string** | An origin&#39;s path  Paths default to \&quot;/\&quot; | [optional] 
**Hostname** | **string** | An origin&#39;s hostname or IP address | [optional] 
**Port** | **int32** | The HTTP port to connect to the origin | [optional] 
**SecurePort** | **int32** | The HTTPS port to connect to the origin | [optional] 
**Dedicated** | **bool** | Whether or not an origin is dedicated to a CDN site  Dedicated origins cannot be used by any site other than that to which it is dedicated. | [optional] 
**SiteId** | **string** | The ID of the CDN site to which an origin is dedicated | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


