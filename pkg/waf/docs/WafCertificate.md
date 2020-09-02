# WafCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A certificate&#39;s unique ID | [optional] 
**Fingerprint** | **string** | A unique hash of a certificate&#39;s contents | [optional] 
**CommonName** | **string** | A certificate&#39;s common name, or the primary domain name the certificate is used for | [optional] 
**Issuer** | **string** | The name of the certificate&#39;s issuing certificate authority | [optional] 
**CaBundle** | **string** | A PEM PKCS #7 formatted certificate authority bundle | [optional] 
**Trusted** | **bool** | Whether or not the certificate&#39;s authority is trusted by a web browser | [optional] 
**ExpirationDate** | [**time.Time**](time.Time.md) | The date an SSL certificate will expire | [optional] 
**CreateDate** | [**time.Time**](time.Time.md) | The date an SSL certificate was created | [optional] 
**UpdateDate** | [**time.Time**](time.Time.md) | The date an SSL certificate was last updated | [optional] 
**SubjectAlternativeNames** | **[]string** | A list of Subject Alternative Names in the certificate  Certificates for multiple domains define their domains in certificate&#39;s SAN list. | [optional] 
**Status** | [**WafCertificateStatus**](wafCertificateStatus.md) |  | [optional] 
**ProviderManaged** | **bool** | Whether a certificate is managed by StackPath or the end user | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


