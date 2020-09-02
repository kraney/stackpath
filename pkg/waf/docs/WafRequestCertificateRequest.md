# WafRequestCertificateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hosts** | **[]string** | An optional list of delivery domains that will be included as subject alternative names on the certificate  If no hosts are provided, all delivery domains on the site will be included with the first one in the list being used as the common name.  If hosts are provided the first entry will be used as the common name.  All entries in the list are validated against the existing delivery domains for the provided site. | [optional] 
**VerificationMethod** | [**WafCertificateVerificationMethod**](wafCertificateVerificationMethod.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


