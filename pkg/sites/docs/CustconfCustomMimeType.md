# CustconfCustomMimeType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | This is used by the API to perform conflict checking | [optional] 
**Code** | **string** | String of values delimited by a &#39;,&#39; character. A comma separated list of status codes that apply to this policy | [optional] 
**ExtensionMap** | **string** | String of values delimited by a &#39;,&#39; character. This is a comma separated list of file extension and mime type pairs that describe the mime type mapping for the CDN caching servers. The file extension and mime type values should be delimited by a colon (:). For example, to map files ending with jpg to the image/jpeg mime type and files ending with png to image/png, set this value to: jpg:image/jpeg,png:image/png. | [optional] 
**Enabled** | **bool** |  | [optional] 
**MethodFilter** | **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**PathFilter** | **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**HeaderFilter** | **string** | String of values delimited by a &#39;,&#39; character. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


