# WafCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A certificate&#39;s unique ID | [optional] 
**Fingerprint** | Pointer to **string** | A unique hash of a certificate&#39;s contents | [optional] 
**CommonName** | Pointer to **string** | A certificate&#39;s common name, or the primary domain name the certificate is used for | [optional] 
**Issuer** | Pointer to **string** | The name of the certificate&#39;s issuing certificate authority | [optional] 
**CaBundle** | Pointer to **string** | A PEM PKCS #7 formatted certificate authority bundle | [optional] 
**Trusted** | Pointer to **bool** | Whether or not the certificate&#39;s authority is trusted by a web browser | [optional] 
**ExpirationDate** | Pointer to [**time.Time**](time.Time.md) | The date an SSL certificate will expire | [optional] 
**CreateDate** | Pointer to [**time.Time**](time.Time.md) | The date an SSL certificate was created | [optional] 
**UpdateDate** | Pointer to [**time.Time**](time.Time.md) | The date an SSL certificate was last updated | [optional] 
**SubjectAlternativeNames** | Pointer to **[]string** | A list of Subject Alternative Names in the certificate  Certificates for multiple domains define their domains in certificate&#39;s SAN list. | [optional] 
**Status** | Pointer to [**WafCertificateStatus**](wafCertificateStatus.md) |  | [optional] [default to "UNKNOWN"]
**ProviderManaged** | Pointer to **bool** | Whether a certificate is managed by StackPath or the end user | [optional] 

## Methods

### NewWafCertificate

`func NewWafCertificate() *WafCertificate`

NewWafCertificate instantiates a new WafCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafCertificateWithDefaults

`func NewWafCertificateWithDefaults() *WafCertificate`

NewWafCertificateWithDefaults instantiates a new WafCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WafCertificate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WafCertificate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WafCertificate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WafCertificate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFingerprint

`func (o *WafCertificate) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *WafCertificate) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *WafCertificate) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *WafCertificate) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetCommonName

`func (o *WafCertificate) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *WafCertificate) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *WafCertificate) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *WafCertificate) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### GetIssuer

`func (o *WafCertificate) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *WafCertificate) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *WafCertificate) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *WafCertificate) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetCaBundle

`func (o *WafCertificate) GetCaBundle() string`

GetCaBundle returns the CaBundle field if non-nil, zero value otherwise.

### GetCaBundleOk

`func (o *WafCertificate) GetCaBundleOk() (*string, bool)`

GetCaBundleOk returns a tuple with the CaBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaBundle

`func (o *WafCertificate) SetCaBundle(v string)`

SetCaBundle sets CaBundle field to given value.

### HasCaBundle

`func (o *WafCertificate) HasCaBundle() bool`

HasCaBundle returns a boolean if a field has been set.

### GetTrusted

`func (o *WafCertificate) GetTrusted() bool`

GetTrusted returns the Trusted field if non-nil, zero value otherwise.

### GetTrustedOk

`func (o *WafCertificate) GetTrustedOk() (*bool, bool)`

GetTrustedOk returns a tuple with the Trusted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrusted

`func (o *WafCertificate) SetTrusted(v bool)`

SetTrusted sets Trusted field to given value.

### HasTrusted

`func (o *WafCertificate) HasTrusted() bool`

HasTrusted returns a boolean if a field has been set.

### GetExpirationDate

`func (o *WafCertificate) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *WafCertificate) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *WafCertificate) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *WafCertificate) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetCreateDate

`func (o *WafCertificate) GetCreateDate() time.Time`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *WafCertificate) GetCreateDateOk() (*time.Time, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *WafCertificate) SetCreateDate(v time.Time)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *WafCertificate) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### GetUpdateDate

`func (o *WafCertificate) GetUpdateDate() time.Time`

GetUpdateDate returns the UpdateDate field if non-nil, zero value otherwise.

### GetUpdateDateOk

`func (o *WafCertificate) GetUpdateDateOk() (*time.Time, bool)`

GetUpdateDateOk returns a tuple with the UpdateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDate

`func (o *WafCertificate) SetUpdateDate(v time.Time)`

SetUpdateDate sets UpdateDate field to given value.

### HasUpdateDate

`func (o *WafCertificate) HasUpdateDate() bool`

HasUpdateDate returns a boolean if a field has been set.

### GetSubjectAlternativeNames

`func (o *WafCertificate) GetSubjectAlternativeNames() []string`

GetSubjectAlternativeNames returns the SubjectAlternativeNames field if non-nil, zero value otherwise.

### GetSubjectAlternativeNamesOk

`func (o *WafCertificate) GetSubjectAlternativeNamesOk() (*[]string, bool)`

GetSubjectAlternativeNamesOk returns a tuple with the SubjectAlternativeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAlternativeNames

`func (o *WafCertificate) SetSubjectAlternativeNames(v []string)`

SetSubjectAlternativeNames sets SubjectAlternativeNames field to given value.

### HasSubjectAlternativeNames

`func (o *WafCertificate) HasSubjectAlternativeNames() bool`

HasSubjectAlternativeNames returns a boolean if a field has been set.

### GetStatus

`func (o *WafCertificate) GetStatus() WafCertificateStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WafCertificate) GetStatusOk() (*WafCertificateStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WafCertificate) SetStatus(v WafCertificateStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WafCertificate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetProviderManaged

`func (o *WafCertificate) GetProviderManaged() bool`

GetProviderManaged returns the ProviderManaged field if non-nil, zero value otherwise.

### GetProviderManagedOk

`func (o *WafCertificate) GetProviderManagedOk() (*bool, bool)`

GetProviderManagedOk returns a tuple with the ProviderManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderManaged

`func (o *WafCertificate) SetProviderManaged(v bool)`

SetProviderManaged sets ProviderManaged field to given value.

### HasProviderManaged

`func (o *WafCertificate) HasProviderManaged() bool`

HasProviderManaged returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


