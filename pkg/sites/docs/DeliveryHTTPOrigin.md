# DeliveryHttpOrigin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | The origin&#39;s HTTP request path | [optional] 
**Hostname** | **string** | The origin&#39;s HTTP request hostname | [optional] 
**Port** | **int32** | The origin&#39;s HTTP request port  Set this to 0 to remove this value | [optional] 
**SecurePort** | **int32** | The origin&#39;s HTTPS request port  Set this to 0 to remove this value | [optional] 
**Authentication** | [**DeliveryOriginAuthentication**](deliveryOriginAuthentication.md) |  | [optional] 
**VerifyCertificate** | **bool** | Verify the origin&#39;s SSL certificate when requesting from the origin | [optional] 
**CertificateCommonName** | **string** | The CommonName to validate SSL origin requests from. This value needs to be provided when enabling &#x60;verify_certificate&#x60;. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


