# WafSiteCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Valid** | Pointer to **bool** | Whether or not a site&#39;s SSL certificate is valid  An SSL certificate is valid when a hostname associated with the site is covered by the certificate. | [optional] 
**Certificate** | Pointer to [**WafCertificate**](wafCertificate.md) |  | [optional] 

## Methods

### NewWafSiteCertificate

`func NewWafSiteCertificate() *WafSiteCertificate`

NewWafSiteCertificate instantiates a new WafSiteCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafSiteCertificateWithDefaults

`func NewWafSiteCertificateWithDefaults() *WafSiteCertificate`

NewWafSiteCertificateWithDefaults instantiates a new WafSiteCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValid

`func (o *WafSiteCertificate) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *WafSiteCertificate) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *WafSiteCertificate) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *WafSiteCertificate) HasValid() bool`

HasValid returns a boolean if a field has been set.

### GetCertificate

`func (o *WafSiteCertificate) GetCertificate() WafCertificate`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *WafSiteCertificate) GetCertificateOk() (*WafCertificate, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *WafSiteCertificate) SetCertificate(v WafCertificate)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *WafSiteCertificate) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


