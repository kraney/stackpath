# CertificateVerificationRequirements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsVerificationDetails** | Pointer to [**CertificateDnsVerificationDetails**](certificateDnsVerificationDetails.md) |  | [optional] 
**HttpVerificationDetails** | Pointer to [**CertificateHttpVerificationDetails**](certificateHttpVerificationDetails.md) |  | [optional] 
**VerificationMethod** | Pointer to [**CertificateCertificateVerificationMethod**](certificateCertificateVerificationMethod.md) |  | [optional] [default to "DNS"]

## Methods

### NewCertificateVerificationRequirements

`func NewCertificateVerificationRequirements() *CertificateVerificationRequirements`

NewCertificateVerificationRequirements instantiates a new CertificateVerificationRequirements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateVerificationRequirementsWithDefaults

`func NewCertificateVerificationRequirementsWithDefaults() *CertificateVerificationRequirements`

NewCertificateVerificationRequirementsWithDefaults instantiates a new CertificateVerificationRequirements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsVerificationDetails

`func (o *CertificateVerificationRequirements) GetDnsVerificationDetails() CertificateDnsVerificationDetails`

GetDnsVerificationDetails returns the DnsVerificationDetails field if non-nil, zero value otherwise.

### GetDnsVerificationDetailsOk

`func (o *CertificateVerificationRequirements) GetDnsVerificationDetailsOk() (*CertificateDnsVerificationDetails, bool)`

GetDnsVerificationDetailsOk returns a tuple with the DnsVerificationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsVerificationDetails

`func (o *CertificateVerificationRequirements) SetDnsVerificationDetails(v CertificateDnsVerificationDetails)`

SetDnsVerificationDetails sets DnsVerificationDetails field to given value.

### HasDnsVerificationDetails

`func (o *CertificateVerificationRequirements) HasDnsVerificationDetails() bool`

HasDnsVerificationDetails returns a boolean if a field has been set.

### GetHttpVerificationDetails

`func (o *CertificateVerificationRequirements) GetHttpVerificationDetails() CertificateHttpVerificationDetails`

GetHttpVerificationDetails returns the HttpVerificationDetails field if non-nil, zero value otherwise.

### GetHttpVerificationDetailsOk

`func (o *CertificateVerificationRequirements) GetHttpVerificationDetailsOk() (*CertificateHttpVerificationDetails, bool)`

GetHttpVerificationDetailsOk returns a tuple with the HttpVerificationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpVerificationDetails

`func (o *CertificateVerificationRequirements) SetHttpVerificationDetails(v CertificateHttpVerificationDetails)`

SetHttpVerificationDetails sets HttpVerificationDetails field to given value.

### HasHttpVerificationDetails

`func (o *CertificateVerificationRequirements) HasHttpVerificationDetails() bool`

HasHttpVerificationDetails returns a boolean if a field has been set.

### GetVerificationMethod

`func (o *CertificateVerificationRequirements) GetVerificationMethod() CertificateCertificateVerificationMethod`

GetVerificationMethod returns the VerificationMethod field if non-nil, zero value otherwise.

### GetVerificationMethodOk

`func (o *CertificateVerificationRequirements) GetVerificationMethodOk() (*CertificateCertificateVerificationMethod, bool)`

GetVerificationMethodOk returns a tuple with the VerificationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationMethod

`func (o *CertificateVerificationRequirements) SetVerificationMethod(v CertificateCertificateVerificationMethod)`

SetVerificationMethod sets VerificationMethod field to given value.

### HasVerificationMethod

`func (o *CertificateVerificationRequirements) HasVerificationMethod() bool`

HasVerificationMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


