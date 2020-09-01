# CdnGetCertificateVerificationDetailsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ManualVerificationRequired** | Pointer to **bool** | Whether or not the end user must provide their own certificate verification | [optional] 
**VerificationRequirements** | Pointer to [**[]CdnVerificationRequirements**](cdnVerificationRequirements.md) | An SSL certificate&#39;s verification requirements | [optional] 

## Methods

### NewCdnGetCertificateVerificationDetailsResponse

`func NewCdnGetCertificateVerificationDetailsResponse() *CdnGetCertificateVerificationDetailsResponse`

NewCdnGetCertificateVerificationDetailsResponse instantiates a new CdnGetCertificateVerificationDetailsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdnGetCertificateVerificationDetailsResponseWithDefaults

`func NewCdnGetCertificateVerificationDetailsResponseWithDefaults() *CdnGetCertificateVerificationDetailsResponse`

NewCdnGetCertificateVerificationDetailsResponseWithDefaults instantiates a new CdnGetCertificateVerificationDetailsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManualVerificationRequired

`func (o *CdnGetCertificateVerificationDetailsResponse) GetManualVerificationRequired() bool`

GetManualVerificationRequired returns the ManualVerificationRequired field if non-nil, zero value otherwise.

### GetManualVerificationRequiredOk

`func (o *CdnGetCertificateVerificationDetailsResponse) GetManualVerificationRequiredOk() (*bool, bool)`

GetManualVerificationRequiredOk returns a tuple with the ManualVerificationRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualVerificationRequired

`func (o *CdnGetCertificateVerificationDetailsResponse) SetManualVerificationRequired(v bool)`

SetManualVerificationRequired sets ManualVerificationRequired field to given value.

### HasManualVerificationRequired

`func (o *CdnGetCertificateVerificationDetailsResponse) HasManualVerificationRequired() bool`

HasManualVerificationRequired returns a boolean if a field has been set.

### GetVerificationRequirements

`func (o *CdnGetCertificateVerificationDetailsResponse) GetVerificationRequirements() []CdnVerificationRequirements`

GetVerificationRequirements returns the VerificationRequirements field if non-nil, zero value otherwise.

### GetVerificationRequirementsOk

`func (o *CdnGetCertificateVerificationDetailsResponse) GetVerificationRequirementsOk() (*[]CdnVerificationRequirements, bool)`

GetVerificationRequirementsOk returns a tuple with the VerificationRequirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationRequirements

`func (o *CdnGetCertificateVerificationDetailsResponse) SetVerificationRequirements(v []CdnVerificationRequirements)`

SetVerificationRequirements sets VerificationRequirements field to given value.

### HasVerificationRequirements

`func (o *CdnGetCertificateVerificationDetailsResponse) HasVerificationRequirements() bool`

HasVerificationRequirements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


