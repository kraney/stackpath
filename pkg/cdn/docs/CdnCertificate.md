# CdnCertificate

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
**Status** | Pointer to [**CdnCertificateStatus**](cdnCertificateStatus.md) |  | [optional] [default to "UNKNOWN"]
**ProviderManaged** | Pointer to **bool** | Whether a certificate is managed by StackPath or the end user | [optional] 

## Methods

### NewCdnCertificate

`func NewCdnCertificate() *CdnCertificate`

NewCdnCertificate instantiates a new CdnCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdnCertificateWithDefaults

`func NewCdnCertificateWithDefaults() *CdnCertificate`

NewCdnCertificateWithDefaults instantiates a new CdnCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CdnCertificate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CdnCertificate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CdnCertificate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CdnCertificate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFingerprint

`func (o *CdnCertificate) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *CdnCertificate) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *CdnCertificate) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *CdnCertificate) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetCommonName

`func (o *CdnCertificate) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *CdnCertificate) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *CdnCertificate) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *CdnCertificate) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### GetIssuer

`func (o *CdnCertificate) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *CdnCertificate) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *CdnCertificate) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *CdnCertificate) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetCaBundle

`func (o *CdnCertificate) GetCaBundle() string`

GetCaBundle returns the CaBundle field if non-nil, zero value otherwise.

### GetCaBundleOk

`func (o *CdnCertificate) GetCaBundleOk() (*string, bool)`

GetCaBundleOk returns a tuple with the CaBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaBundle

`func (o *CdnCertificate) SetCaBundle(v string)`

SetCaBundle sets CaBundle field to given value.

### HasCaBundle

`func (o *CdnCertificate) HasCaBundle() bool`

HasCaBundle returns a boolean if a field has been set.

### GetTrusted

`func (o *CdnCertificate) GetTrusted() bool`

GetTrusted returns the Trusted field if non-nil, zero value otherwise.

### GetTrustedOk

`func (o *CdnCertificate) GetTrustedOk() (*bool, bool)`

GetTrustedOk returns a tuple with the Trusted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrusted

`func (o *CdnCertificate) SetTrusted(v bool)`

SetTrusted sets Trusted field to given value.

### HasTrusted

`func (o *CdnCertificate) HasTrusted() bool`

HasTrusted returns a boolean if a field has been set.

### GetExpirationDate

`func (o *CdnCertificate) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *CdnCertificate) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *CdnCertificate) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *CdnCertificate) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetCreateDate

`func (o *CdnCertificate) GetCreateDate() time.Time`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *CdnCertificate) GetCreateDateOk() (*time.Time, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *CdnCertificate) SetCreateDate(v time.Time)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *CdnCertificate) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### GetUpdateDate

`func (o *CdnCertificate) GetUpdateDate() time.Time`

GetUpdateDate returns the UpdateDate field if non-nil, zero value otherwise.

### GetUpdateDateOk

`func (o *CdnCertificate) GetUpdateDateOk() (*time.Time, bool)`

GetUpdateDateOk returns a tuple with the UpdateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDate

`func (o *CdnCertificate) SetUpdateDate(v time.Time)`

SetUpdateDate sets UpdateDate field to given value.

### HasUpdateDate

`func (o *CdnCertificate) HasUpdateDate() bool`

HasUpdateDate returns a boolean if a field has been set.

### GetSubjectAlternativeNames

`func (o *CdnCertificate) GetSubjectAlternativeNames() []string`

GetSubjectAlternativeNames returns the SubjectAlternativeNames field if non-nil, zero value otherwise.

### GetSubjectAlternativeNamesOk

`func (o *CdnCertificate) GetSubjectAlternativeNamesOk() (*[]string, bool)`

GetSubjectAlternativeNamesOk returns a tuple with the SubjectAlternativeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAlternativeNames

`func (o *CdnCertificate) SetSubjectAlternativeNames(v []string)`

SetSubjectAlternativeNames sets SubjectAlternativeNames field to given value.

### HasSubjectAlternativeNames

`func (o *CdnCertificate) HasSubjectAlternativeNames() bool`

HasSubjectAlternativeNames returns a boolean if a field has been set.

### GetStatus

`func (o *CdnCertificate) GetStatus() CdnCertificateStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CdnCertificate) GetStatusOk() (*CdnCertificateStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CdnCertificate) SetStatus(v CdnCertificateStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CdnCertificate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetProviderManaged

`func (o *CdnCertificate) GetProviderManaged() bool`

GetProviderManaged returns the ProviderManaged field if non-nil, zero value otherwise.

### GetProviderManagedOk

`func (o *CdnCertificate) GetProviderManagedOk() (*bool, bool)`

GetProviderManagedOk returns a tuple with the ProviderManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderManaged

`func (o *CdnCertificate) SetProviderManaged(v bool)`

SetProviderManaged sets ProviderManaged field to given value.

### HasProviderManaged

`func (o *CdnCertificate) HasProviderManaged() bool`

HasProviderManaged returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


