# CdnRequestCertificateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | Pointer to [**CdnCertificate**](cdnCertificate.md) |  | [optional] 
**VerificationRequirements** | Pointer to [**[]CdnVerificationRequirements**](cdnVerificationRequirements.md) | The certificate&#39;s verification requirements | [optional] 

## Methods

### NewCdnRequestCertificateResponse

`func NewCdnRequestCertificateResponse() *CdnRequestCertificateResponse`

NewCdnRequestCertificateResponse instantiates a new CdnRequestCertificateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdnRequestCertificateResponseWithDefaults

`func NewCdnRequestCertificateResponseWithDefaults() *CdnRequestCertificateResponse`

NewCdnRequestCertificateResponseWithDefaults instantiates a new CdnRequestCertificateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *CdnRequestCertificateResponse) GetCertificate() CdnCertificate`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *CdnRequestCertificateResponse) GetCertificateOk() (*CdnCertificate, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *CdnRequestCertificateResponse) SetCertificate(v CdnCertificate)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *CdnRequestCertificateResponse) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetVerificationRequirements

`func (o *CdnRequestCertificateResponse) GetVerificationRequirements() []CdnVerificationRequirements`

GetVerificationRequirements returns the VerificationRequirements field if non-nil, zero value otherwise.

### GetVerificationRequirementsOk

`func (o *CdnRequestCertificateResponse) GetVerificationRequirementsOk() (*[]CdnVerificationRequirements, bool)`

GetVerificationRequirementsOk returns a tuple with the VerificationRequirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationRequirements

`func (o *CdnRequestCertificateResponse) SetVerificationRequirements(v []CdnVerificationRequirements)`

SetVerificationRequirements sets VerificationRequirements field to given value.

### HasVerificationRequirements

`func (o *CdnRequestCertificateResponse) HasVerificationRequirements() bool`

HasVerificationRequirements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


