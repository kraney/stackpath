# CustconfClientResponseModification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | This is used by the API to perform conflict checking | [optional] 
**StatusCodeRewrite** | **float32** |  | [optional] 
**HeaderPattern** | **string** |  | [optional] 
**HeaderRewrite** | **string** |  | [optional] 
**AddHeaders** | **string** | String of values delimited by a &#39;|&#39; character. This is the static HTTP header you want inserted into the CDN response. Start value with \&quot;append:\&quot;, \&quot;replace:\&quot; or \&quot;create:\&quot; which defines if Header will be added, replaced or create if not exists. Default is append. | [optional] 
**FlowControl** | [**CustconfClientResponseModificationFlowControlEnumWrapperValue**](custconfClientResponseModificationFlowControlEnumWrapperValue.md) |  | [optional] 
**Enabled** | **bool** |  | [optional] 
**MethodFilter** | **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**PathFilter** | **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**HeaderFilter** | **string** | String of values delimited by a &#39;,&#39; character. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


