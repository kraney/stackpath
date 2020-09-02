# V2HttpConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | The URL of the service to monitor  The URL must begin with a case insensitive HTTP scheme of &#39;http&#39; or &#39;https&#39;. | [optional] 
**Method** | **string** | The HTTP method used when the monitor makes a request to the service. | [optional] 
**Body** | **string** | A base64 encoded HTTP request body to use when a monitor checks the service. | [optional] 
**Headers** | [**[]V2Header**](v2Header.md) | A list of HTTP headers to add to the request when a monitor checks a service. | [optional] 
**Basic** | [**V2HttpConfigurationBasicAuth**](v2HttpConfigurationBasicAuth.md) |  | [optional] 
**Digest** | [**V2HttpConfigurationDigestAuth**](v2HttpConfigurationDigestAuth.md) |  | [optional] 
**ClientCertificate** | [**V2HttpConfigurationClientCertificate**](v2HttpConfigurationClientCertificate.md) |  | [optional] 
**ValidateCertificate** | **bool** | Whether or not to validate a service&#39;s certificate. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


