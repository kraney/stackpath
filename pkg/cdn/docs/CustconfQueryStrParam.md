# CustconfQueryStrParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | This is used by the API to perform conflict checking | [optional] 
**DispositionName** | **string** | This is the name of the query string parameter which contains the name of the file to specify in the Content-Disposition HTTP response header. This setting is typically used by customers to configure a \&quot;friendly name\&quot; for URLs that have obfuscated file names. This setting controls the \&quot;filename\&quot; directive that is part of the Content-Disposition HTTP header. | [optional] 
**DispositionType** | **string** | This is the name of the query string parameter which contains the disposition type to use in the Content-Disposition HTTP header. Typically, this value is set to \&quot;attachment,\&quot; but you may supply a custom string using this setting. | [optional] 
**DispositionOverride** | **string** | This setting allows you to completely override the Content-Disposition HTTP header that the CDN caching servers use on a response. | [optional] 
**JumpToTimeStart** | **string** | This is the name of the query string parameter that indicates to the CDN the start time offset of the video returned. This parameter is part of the jump-to-time feature that is initiated on a per request basis. | [optional] 
**JumpToTimeEnd** | **string** | This is the name of the query string parameter that indicates to the CDN the end time offset of the video that should be returned. This parameter is part of the jump-to-time feature that is initiated on a per request basis. | [optional] 
**JumpToByteInitialBytes** | **string** | This is the  name of the query string parameter that indicates to the CDN the initial bytes of a video that should be returned before sending the requested byte offset. This parameter is part of the jump-to-byte feature that is initiated on a per request basis. | [optional] 
**JumpToByteStartOffset** | **string** | This is the name of the query string parameter that indicates to the CDN the specific offset into the video that is being requested. This parameter is part of the jump-to-byte feature that is initiated on a per request basis. | [optional] 
**RateLimitInitial** | **string** | This is the name of the query string parameter that indicates to the CDN an initial burst rate to use when delivering a file. This parameter is part of the bandwidth limiting feature that is initiated on a per request basis. | [optional] 
**RateLimitSustained** | **string** | This is the name of the query string parameter that indicates to the CDN the sustained rate being requested for the delivery of a file. This parameter is part of the bandwidth throttling feature that is initiated on a per request basis. | [optional] 
**Enabled** | **bool** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


