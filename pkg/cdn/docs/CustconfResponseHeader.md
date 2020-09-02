# CustconfResponseHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | This is used by the API to perform conflict checking | [optional] 
**Http** | **string** | A pipe delimited list of rules that instructs the CDN caching servers to include a content-disposition header. The rules included in this setting must be entered in the following format: Content-Disposition:&lt;User Agent&gt;:&lt;file extension 1&gt;, &lt;file extension 2&gt;. For example, to send the Content-Disposition header for all Mozilla browsers on the delivery of mp3, exe, tar, zip, gz and rar files, you would enter the following in the field: Content-Disposition:Mozilla*:mp3,exe,tar,zip,gz,rar | [optional] 
**EnableETag** | **bool** | This gives the ability to disable the ETag header on the response. | [optional] 
**Enabled** | **bool** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


