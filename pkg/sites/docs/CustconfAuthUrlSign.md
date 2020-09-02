# CustconfAuthUrlSign

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | This is used by the API to perform conflict checking | [optional] 
**TokenField** | **string** | This is the name of the query string parameter that will be used by the publisher to specify the signature for the URL. | [optional] 
**IgnoreFieldsAfterToken** | **bool** | Select this option if you want StackPath to exclude query string parameters specified after the passphrase in the validation process. | [optional] 
**PassPhraseField** | **string** | This is the name of the query string parameter that contains the value of the secret. This query string parameter is only used during the generation and validation of a URL signature and should not be present in the published URL. | [optional] 
**PassPhrase** | **string** | The shared secret used during the signing process. This value should only be known by StackPath and systems authorized to sign your content. | [optional] 
**ExpiresField** | **string** | This is the name of the query string parameter that contains the Epoch time after which the URL is considered invalid. | [optional] 
**IpAddressField** | **string** | This is a query string parameter that contains an IPv4 address to which the published URL will be restricted. | [optional] 
**UriLengthField** | **string** | This is a query string parameter used to limit the number of characters in the path that should be considered when validating the URL signature. This feature allows you to reuse the same signature on all assets stored in the same cache path. Setting this value to 0 will strip off the filename in the URL (leaving the trailing slash) when calculating the checksum. | [optional] 
**UserAgentField** | **string** | This is a query string parameter used to restrict the published URL to a specific user agent. Publishers can use this feature during the signing process to ensure that only a specific user agent can access the published content. You do not need to specify the user agent on the published URL. StackPath will use the HTTP User-Agent header value during signature validation. | [optional] 
**Enabled** | **bool** |  | [optional] 
**MethodFilter** | **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**PathFilter** | **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**HeaderFilter** | **string** | String of values delimited by a &#39;,&#39; character. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


