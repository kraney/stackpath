# CustconfOriginPull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | This is used by the API to perform conflict checking | [optional] 
**RedirectAction** | [**OriginPullRedirectActionEnumWrapperValue**](OriginPullRedirectActionEnumWrapperValue.md) |  | [optional] 
**NoQSParams** | **bool** | GFS sends a path without any query string parameters when making external origin requests regardless if any parameters were sent by the User-Agent. | [optional] 
**RetryMethods** | **string** | String of values delimited by a &#39;,&#39; character. List of HTTP Methods that define types of origin pull requests that can be retried if a failure occurs after sending a previous request. | [optional] 
**Enabled** | **bool** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


