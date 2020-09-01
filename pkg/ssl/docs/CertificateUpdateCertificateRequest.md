# CertificateUpdateCertificateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommonName** | Pointer to **string** |  | [optional] 
**Hosts** | Pointer to **[]string** |  | [optional] 
**Organization** | Pointer to **string** |  | [optional] 
**OrganizationalUnit** | Pointer to **string** |  | [optional] 
**Locality** | Pointer to **string** |  | [optional] 
**Province** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 

## Methods

### NewCertificateUpdateCertificateRequest

`func NewCertificateUpdateCertificateRequest() *CertificateUpdateCertificateRequest`

NewCertificateUpdateCertificateRequest instantiates a new CertificateUpdateCertificateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateUpdateCertificateRequestWithDefaults

`func NewCertificateUpdateCertificateRequestWithDefaults() *CertificateUpdateCertificateRequest`

NewCertificateUpdateCertificateRequestWithDefaults instantiates a new CertificateUpdateCertificateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommonName

`func (o *CertificateUpdateCertificateRequest) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *CertificateUpdateCertificateRequest) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *CertificateUpdateCertificateRequest) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *CertificateUpdateCertificateRequest) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### GetHosts

`func (o *CertificateUpdateCertificateRequest) GetHosts() []string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *CertificateUpdateCertificateRequest) GetHostsOk() (*[]string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *CertificateUpdateCertificateRequest) SetHosts(v []string)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *CertificateUpdateCertificateRequest) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetOrganization

`func (o *CertificateUpdateCertificateRequest) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *CertificateUpdateCertificateRequest) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *CertificateUpdateCertificateRequest) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *CertificateUpdateCertificateRequest) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetOrganizationalUnit

`func (o *CertificateUpdateCertificateRequest) GetOrganizationalUnit() string`

GetOrganizationalUnit returns the OrganizationalUnit field if non-nil, zero value otherwise.

### GetOrganizationalUnitOk

`func (o *CertificateUpdateCertificateRequest) GetOrganizationalUnitOk() (*string, bool)`

GetOrganizationalUnitOk returns a tuple with the OrganizationalUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationalUnit

`func (o *CertificateUpdateCertificateRequest) SetOrganizationalUnit(v string)`

SetOrganizationalUnit sets OrganizationalUnit field to given value.

### HasOrganizationalUnit

`func (o *CertificateUpdateCertificateRequest) HasOrganizationalUnit() bool`

HasOrganizationalUnit returns a boolean if a field has been set.

### GetLocality

`func (o *CertificateUpdateCertificateRequest) GetLocality() string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *CertificateUpdateCertificateRequest) GetLocalityOk() (*string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *CertificateUpdateCertificateRequest) SetLocality(v string)`

SetLocality sets Locality field to given value.

### HasLocality

`func (o *CertificateUpdateCertificateRequest) HasLocality() bool`

HasLocality returns a boolean if a field has been set.

### GetProvince

`func (o *CertificateUpdateCertificateRequest) GetProvince() string`

GetProvince returns the Province field if non-nil, zero value otherwise.

### GetProvinceOk

`func (o *CertificateUpdateCertificateRequest) GetProvinceOk() (*string, bool)`

GetProvinceOk returns a tuple with the Province field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvince

`func (o *CertificateUpdateCertificateRequest) SetProvince(v string)`

SetProvince sets Province field to given value.

### HasProvince

`func (o *CertificateUpdateCertificateRequest) HasProvince() bool`

HasProvince returns a boolean if a field has been set.

### GetCountry

`func (o *CertificateUpdateCertificateRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CertificateUpdateCertificateRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CertificateUpdateCertificateRequest) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CertificateUpdateCertificateRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


