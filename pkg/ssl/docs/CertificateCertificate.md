# CertificateCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StackId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Hosts** | Pointer to **[]string** |  | [optional] 
**Issuer** | Pointer to **string** |  | [optional] 
**Certificate** | Pointer to **string** |  | [optional] 
**CaBundle** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**CertificateCertificateStatus**](certificateCertificateStatus.md) |  | [optional] [default to "UNKNOWN"]
**ExpiresAt** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**CreatedAt** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**UpdatedAt** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 

## Methods

### NewCertificateCertificate

`func NewCertificateCertificate() *CertificateCertificate`

NewCertificateCertificate instantiates a new CertificateCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateCertificateWithDefaults

`func NewCertificateCertificateWithDefaults() *CertificateCertificate`

NewCertificateCertificateWithDefaults instantiates a new CertificateCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStackId

`func (o *CertificateCertificate) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *CertificateCertificate) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *CertificateCertificate) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *CertificateCertificate) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetId

`func (o *CertificateCertificate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertificateCertificate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertificateCertificate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CertificateCertificate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHosts

`func (o *CertificateCertificate) GetHosts() []string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *CertificateCertificate) GetHostsOk() (*[]string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *CertificateCertificate) SetHosts(v []string)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *CertificateCertificate) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetIssuer

`func (o *CertificateCertificate) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *CertificateCertificate) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *CertificateCertificate) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *CertificateCertificate) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetCertificate

`func (o *CertificateCertificate) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *CertificateCertificate) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *CertificateCertificate) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *CertificateCertificate) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetCaBundle

`func (o *CertificateCertificate) GetCaBundle() string`

GetCaBundle returns the CaBundle field if non-nil, zero value otherwise.

### GetCaBundleOk

`func (o *CertificateCertificate) GetCaBundleOk() (*string, bool)`

GetCaBundleOk returns a tuple with the CaBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaBundle

`func (o *CertificateCertificate) SetCaBundle(v string)`

SetCaBundle sets CaBundle field to given value.

### HasCaBundle

`func (o *CertificateCertificate) HasCaBundle() bool`

HasCaBundle returns a boolean if a field has been set.

### GetStatus

`func (o *CertificateCertificate) GetStatus() CertificateCertificateStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CertificateCertificate) GetStatusOk() (*CertificateCertificateStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CertificateCertificate) SetStatus(v CertificateCertificateStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CertificateCertificate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetExpiresAt

`func (o *CertificateCertificate) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CertificateCertificate) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CertificateCertificate) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CertificateCertificate) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CertificateCertificate) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CertificateCertificate) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CertificateCertificate) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CertificateCertificate) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CertificateCertificate) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CertificateCertificate) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CertificateCertificate) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CertificateCertificate) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


