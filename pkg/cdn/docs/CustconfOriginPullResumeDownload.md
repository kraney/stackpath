# CustconfOriginPullResumeDownload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | This is used by the API to perform conflict checking | [optional] 
**Enabled** | **bool** |  | [optional] 
**OriginalStatusCodeMatch** | **string** | String of values delimited by a &#39;,&#39; character. Comma separated list of glob patterns that indicate which origin pulls this policy applies to based on the status code of the original origin response. This feature is limited to 2xx responses, but this policy provides fine control, such as permitting resume of 200 responses by not 206, or for adding a 2xx response code other than 200 or 206. | [optional] 
**MinimumBodyBytesConsumed** | **string** | Minimum number of body bytes that CDN must have consumed during an Origin Pull before encountering an error before it is permitted to attempt resuming the download. This value does not include the headers bytes. | [optional] 
**MaximumAttempts** | **float32** | Maximum number of times beyond the initial request that the CDN is permitted to attempt resuming an Origin Pull download. | [optional] 
**RequireOriginRangeSupport** | **bool** | The CDN resumes an Origin Pull even if the origin does not support range requests. If the origin does not support range requests and/or returns a 200 response for the same version given in the original response, the CDN fast-forwards (reads and discards bytes) until it reaches its previous position in the asset. | [optional] 
**EtagValidation** | [**OriginPullResumeDownloadEtagValidationEnumWrapperValue**](OriginPullResumeDownloadEtagValidationEnumWrapperValue.md) |  | [optional] 
**HeaderFilter** | **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**PathFilter** | **string** | String of values delimited by a &#39;,&#39; character. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


