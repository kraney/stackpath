# WafRequestCertificateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | Pointer to [**WafCertificate**](wafCertificate.md) |  | [optional] 
**VerificationRequirements** | Pointer to [**[]WafVerificationRequirements**](wafVerificationRequirements.md) | The certificate&#39;s verification requirements | [optional] 

## Methods

### NewWafRequestCertificateResponse

`func NewWafRequestCertificateResponse() *WafRequestCertificateResponse`

NewWafRequestCertificateResponse instantiates a new WafRequestCertificateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafRequestCertificateResponseWithDefaults

`func NewWafRequestCertificateResponseWithDefaults() *WafRequestCertificateResponse`

NewWafRequestCertificateResponseWithDefaults instantiates a new WafRequestCertificateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *WafRequestCertificateResponse) GetCertificate() WafCertificate`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *WafRequestCertificateResponse) GetCertificateOk() (*WafCertificate, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *WafRequestCertificateResponse) SetCertificate(v WafCertificate)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *WafRequestCertificateResponse) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetVerificationRequirements

`func (o *WafRequestCertificateResponse) GetVerificationRequirements() []WafVerificationRequirements`

GetVerificationRequirements returns the VerificationRequirements field if non-nil, zero value otherwise.

### GetVerificationRequirementsOk

`func (o *WafRequestCertificateResponse) GetVerificationRequirementsOk() (*[]WafVerificationRequirements, bool)`

GetVerificationRequirementsOk returns a tuple with the VerificationRequirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationRequirements

`func (o *WafRequestCertificateResponse) SetVerificationRequirements(v []WafVerificationRequirements)`

SetVerificationRequirements sets VerificationRequirements field to given value.

### HasVerificationRequirements

`func (o *WafRequestCertificateResponse) HasVerificationRequirements() bool`

HasVerificationRequirements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


