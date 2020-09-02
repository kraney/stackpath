# UpdateMonitorRequestPatchHttpConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | The URL of the service that will be monitored.  The URL must begin with an HTTP scheme of &#39;http&#39; or &#39;https&#39;. | [optional] 
**Method** | **string** | The HTTP method used when the monitor makes a request to the service. | [optional] 
**Body** | **string** | A base64 encoded HTTP request body to use when a monitor checks the service. | [optional] 
**Headers** | [**PatchHttpConfigurationHeaderValue**](PatchHttpConfigurationHeaderValue.md) |  | [optional] 
**Basic** | [**UpdateMonitorRequestPatchHttpConfigurationBasicAuth**](UpdateMonitorRequestPatchHttpConfigurationBasicAuth.md) |  | [optional] 
**Digest** | [**UpdateMonitorRequestPatchHttpConfigurationDigestAuth**](UpdateMonitorRequestPatchHttpConfigurationDigestAuth.md) |  | [optional] 
**ClientCertificate** | [**UpdateMonitorRequestPatchHttpConfigurationClientCertificate**](UpdateMonitorRequestPatchHttpConfigurationClientCertificate.md) |  | [optional] 
**ValidateCertificate** | **bool** | A flag for validating a service&#39;s certificate. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


