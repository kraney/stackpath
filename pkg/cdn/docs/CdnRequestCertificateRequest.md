# CdnRequestCertificateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hosts** | Pointer to **[]string** | An optional list of delivery domains that will be included as subject alternative names on the certificate  If no hosts are provided, all delivery domains on the site will be included with the first one in the list being used as the common name.  If hosts are provided, the first entry will be used as the common name.  All entries in the list are validated against the existing delivery domains for the provided site. | [optional] 
**VerificationMethod** | Pointer to [**CdnCertificateVerificationMethod**](cdnCertificateVerificationMethod.md) |  | [optional] [default to "DNS"]

## Methods

### NewCdnRequestCertificateRequest

`func NewCdnRequestCertificateRequest() *CdnRequestCertificateRequest`

NewCdnRequestCertificateRequest instantiates a new CdnRequestCertificateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdnRequestCertificateRequestWithDefaults

`func NewCdnRequestCertificateRequestWithDefaults() *CdnRequestCertificateRequest`

NewCdnRequestCertificateRequestWithDefaults instantiates a new CdnRequestCertificateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHosts

`func (o *CdnRequestCertificateRequest) GetHosts() []string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *CdnRequestCertificateRequest) GetHostsOk() (*[]string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *CdnRequestCertificateRequest) SetHosts(v []string)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *CdnRequestCertificateRequest) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetVerificationMethod

`func (o *CdnRequestCertificateRequest) GetVerificationMethod() CdnCertificateVerificationMethod`

GetVerificationMethod returns the VerificationMethod field if non-nil, zero value otherwise.

### GetVerificationMethodOk

`func (o *CdnRequestCertificateRequest) GetVerificationMethodOk() (*CdnCertificateVerificationMethod, bool)`

GetVerificationMethodOk returns a tuple with the VerificationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationMethod

`func (o *CdnRequestCertificateRequest) SetVerificationMethod(v CdnCertificateVerificationMethod)`

SetVerificationMethod sets VerificationMethod field to given value.

### HasVerificationMethod

`func (o *CdnRequestCertificateRequest) HasVerificationMethod() bool`

HasVerificationMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


