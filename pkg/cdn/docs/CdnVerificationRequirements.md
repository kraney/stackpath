# CdnVerificationRequirements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsVerificationDetails** | Pointer to [**CdnDnsVerificationDetails**](cdnDnsVerificationDetails.md) |  | [optional] 
**HttpVerificationDetails** | Pointer to [**CdnHttpVerificationDetails**](cdnHttpVerificationDetails.md) |  | [optional] 
**VerificationMethod** | Pointer to [**CdnCertificateVerificationMethod**](cdnCertificateVerificationMethod.md) |  | [optional] [default to "DNS"]

## Methods

### NewCdnVerificationRequirements

`func NewCdnVerificationRequirements() *CdnVerificationRequirements`

NewCdnVerificationRequirements instantiates a new CdnVerificationRequirements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdnVerificationRequirementsWithDefaults

`func NewCdnVerificationRequirementsWithDefaults() *CdnVerificationRequirements`

NewCdnVerificationRequirementsWithDefaults instantiates a new CdnVerificationRequirements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsVerificationDetails

`func (o *CdnVerificationRequirements) GetDnsVerificationDetails() CdnDnsVerificationDetails`

GetDnsVerificationDetails returns the DnsVerificationDetails field if non-nil, zero value otherwise.

### GetDnsVerificationDetailsOk

`func (o *CdnVerificationRequirements) GetDnsVerificationDetailsOk() (*CdnDnsVerificationDetails, bool)`

GetDnsVerificationDetailsOk returns a tuple with the DnsVerificationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsVerificationDetails

`func (o *CdnVerificationRequirements) SetDnsVerificationDetails(v CdnDnsVerificationDetails)`

SetDnsVerificationDetails sets DnsVerificationDetails field to given value.

### HasDnsVerificationDetails

`func (o *CdnVerificationRequirements) HasDnsVerificationDetails() bool`

HasDnsVerificationDetails returns a boolean if a field has been set.

### GetHttpVerificationDetails

`func (o *CdnVerificationRequirements) GetHttpVerificationDetails() CdnHttpVerificationDetails`

GetHttpVerificationDetails returns the HttpVerificationDetails field if non-nil, zero value otherwise.

### GetHttpVerificationDetailsOk

`func (o *CdnVerificationRequirements) GetHttpVerificationDetailsOk() (*CdnHttpVerificationDetails, bool)`

GetHttpVerificationDetailsOk returns a tuple with the HttpVerificationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpVerificationDetails

`func (o *CdnVerificationRequirements) SetHttpVerificationDetails(v CdnHttpVerificationDetails)`

SetHttpVerificationDetails sets HttpVerificationDetails field to given value.

### HasHttpVerificationDetails

`func (o *CdnVerificationRequirements) HasHttpVerificationDetails() bool`

HasHttpVerificationDetails returns a boolean if a field has been set.

### GetVerificationMethod

`func (o *CdnVerificationRequirements) GetVerificationMethod() CdnCertificateVerificationMethod`

GetVerificationMethod returns the VerificationMethod field if non-nil, zero value otherwise.

### GetVerificationMethodOk

`func (o *CdnVerificationRequirements) GetVerificationMethodOk() (*CdnCertificateVerificationMethod, bool)`

GetVerificationMethodOk returns a tuple with the VerificationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationMethod

`func (o *CdnVerificationRequirements) SetVerificationMethod(v CdnCertificateVerificationMethod)`

SetVerificationMethod sets VerificationMethod field to given value.

### HasVerificationMethod

`func (o *CdnVerificationRequirements) HasVerificationMethod() bool`

HasVerificationMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


