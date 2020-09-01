# CdnCreateCertificateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | Pointer to **string** | A PEM PKCS #7 formatted SSL certificate | [optional] 
**Key** | Pointer to **string** | A PEM PKCS #7 formatted private key  Private keys are sent directly to the edge nodes and are not stored elsewhere on StackPath&#39;s systems. | [optional] 
**CaBundle** | Pointer to **string** | A PEM PKCS #7 formatted certificate authority bundle | [optional] 

## Methods

### NewCdnCreateCertificateRequest

`func NewCdnCreateCertificateRequest() *CdnCreateCertificateRequest`

NewCdnCreateCertificateRequest instantiates a new CdnCreateCertificateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdnCreateCertificateRequestWithDefaults

`func NewCdnCreateCertificateRequestWithDefaults() *CdnCreateCertificateRequest`

NewCdnCreateCertificateRequestWithDefaults instantiates a new CdnCreateCertificateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *CdnCreateCertificateRequest) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *CdnCreateCertificateRequest) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *CdnCreateCertificateRequest) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *CdnCreateCertificateRequest) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetKey

`func (o *CdnCreateCertificateRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CdnCreateCertificateRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CdnCreateCertificateRequest) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CdnCreateCertificateRequest) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetCaBundle

`func (o *CdnCreateCertificateRequest) GetCaBundle() string`

GetCaBundle returns the CaBundle field if non-nil, zero value otherwise.

### GetCaBundleOk

`func (o *CdnCreateCertificateRequest) GetCaBundleOk() (*string, bool)`

GetCaBundleOk returns a tuple with the CaBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaBundle

`func (o *CdnCreateCertificateRequest) SetCaBundle(v string)`

SetCaBundle sets CaBundle field to given value.

### HasCaBundle

`func (o *CdnCreateCertificateRequest) HasCaBundle() bool`

HasCaBundle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


