# CertificateRenewCertificateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**CertificateCertificateStatus**](certificateCertificateStatus.md) |  | [optional] [default to "UNKNOWN"]
**VerificationRequirements** | Pointer to [**[]CertificateVerificationRequirements**](certificateVerificationRequirements.md) |  | [optional] 

## Methods

### NewCertificateRenewCertificateResponse

`func NewCertificateRenewCertificateResponse() *CertificateRenewCertificateResponse`

NewCertificateRenewCertificateResponse instantiates a new CertificateRenewCertificateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateRenewCertificateResponseWithDefaults

`func NewCertificateRenewCertificateResponseWithDefaults() *CertificateRenewCertificateResponse`

NewCertificateRenewCertificateResponseWithDefaults instantiates a new CertificateRenewCertificateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateId

`func (o *CertificateRenewCertificateResponse) GetCertificateId() string`

GetCertificateId returns the CertificateId field if non-nil, zero value otherwise.

### GetCertificateIdOk

`func (o *CertificateRenewCertificateResponse) GetCertificateIdOk() (*string, bool)`

GetCertificateIdOk returns a tuple with the CertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateId

`func (o *CertificateRenewCertificateResponse) SetCertificateId(v string)`

SetCertificateId sets CertificateId field to given value.

### HasCertificateId

`func (o *CertificateRenewCertificateResponse) HasCertificateId() bool`

HasCertificateId returns a boolean if a field has been set.

### GetStatus

`func (o *CertificateRenewCertificateResponse) GetStatus() CertificateCertificateStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CertificateRenewCertificateResponse) GetStatusOk() (*CertificateCertificateStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CertificateRenewCertificateResponse) SetStatus(v CertificateCertificateStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CertificateRenewCertificateResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVerificationRequirements

`func (o *CertificateRenewCertificateResponse) GetVerificationRequirements() []CertificateVerificationRequirements`

GetVerificationRequirements returns the VerificationRequirements field if non-nil, zero value otherwise.

### GetVerificationRequirementsOk

`func (o *CertificateRenewCertificateResponse) GetVerificationRequirementsOk() (*[]CertificateVerificationRequirements, bool)`

GetVerificationRequirementsOk returns a tuple with the VerificationRequirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationRequirements

`func (o *CertificateRenewCertificateResponse) SetVerificationRequirements(v []CertificateVerificationRequirements)`

SetVerificationRequirements sets VerificationRequirements field to given value.

### HasVerificationRequirements

`func (o *CertificateRenewCertificateResponse) HasVerificationRequirements() bool`

HasVerificationRequirements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


