# CustconfDynamicContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | This is used by the API to perform conflict checking | [optional] 
**QueryParams** | **string** | String of values delimited by a &#39;,&#39; character. A comma separated list of query string parameters you want to include in the cache key generation. NOTE: This list is ignored when the Key Location is set to header. | [optional] 
**HeaderFields** | **string** | String of values delimited by a &#39;,&#39; character. A comma-separated list of glob patterns that represent HTTP request headers you want included in the cache key generation. Via the use of a colon &#39;:&#39;, users can define each glob pattern to match a header name only, or the pattern can be used to match both the header name and value. A pattern without a colon &#39;:&#39; is treated as a header name ONLY match. If the pattern matches the header, then all values are used as a part of the cache key. If the pattern contains a colon, the CDN uses the pattern on the left of the colon to match the header. The pattern to the right of the colon is used to match the value. The CDN only uses the header/value as a part of the cache key if both patterns result in a match. Note, no pattern after a colon indicates an empty header (no value). See the fnmatch manpage for acceptable glob patterns. | [optional] 
**Enabled** | **bool** |  | [optional] 
**MethodFilter** | **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**PathFilter** | **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**HeaderFilter** | **string** | String of values delimited by a &#39;,&#39; character. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


