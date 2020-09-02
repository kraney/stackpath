# CustconfRedirectMappings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | This is used by the API to perform conflict checking | [optional] 
**Code** | **float32** | The HTTP cache response code that applies to this policy. | [optional] 
**RedirectURL** | **string** | The URL that clients would be redirected to when applying this policy. | [optional] 
**ReplacementToken** | **string** | A token that will be replaced by the caching server with the URL of the request that triggered the policy. This token can be positioned anywhere in the redirect URL. | [optional] 
**Enabled** | **bool** |  | [optional] 
**MethodFilter** | **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**PathFilter** | **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**HeaderFilter** | **string** | String of values delimited by a &#39;,&#39; character. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


