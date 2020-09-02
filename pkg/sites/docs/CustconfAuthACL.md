# CustconfAuthAcl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | This is used by the API to perform conflict checking | [optional] 
**AccessCode** | [**AuthAclAccessCodeEnumWrapperValue**](AuthACLAccessCodeEnumWrapperValue.md) |  | [optional] 
**IpList** | **string** | String of values delimited by a &#39;,&#39; character. This is a list of IP addresses considered for this policy. | [optional] 
**Protocol** | [**CustconfAuthAclProtocolEnumWrapperValue**](custconfAuthACLProtocolEnumWrapperValue.md) |  | [optional] 
**ClientIPSrc** | [**AuthAclClientIpSrcEnumWrapperValue**](AuthACLClientIPSrcEnumWrapperValue.md) |  | [optional] 
**Header** | **string** | This allows you to specify the name of a HTTP request header which will contain the client IP address to use for this policy. | [optional] 
**Enabled** | **bool** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


