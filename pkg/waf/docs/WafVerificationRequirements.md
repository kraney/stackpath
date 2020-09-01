# WafVerificationRequirements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsVerificationDetails** | Pointer to [**WafDnsVerificationDetails**](wafDnsVerificationDetails.md) |  | [optional] 
**HttpVerificationDetails** | Pointer to [**WafHttpVerificationDetails**](wafHttpVerificationDetails.md) |  | [optional] 
**VerificationMethod** | Pointer to [**WafCertificateVerificationMethod**](wafCertificateVerificationMethod.md) |  | [optional] [default to "DNS"]

## Methods

### NewWafVerificationRequirements

`func NewWafVerificationRequirements() *WafVerificationRequirements`

NewWafVerificationRequirements instantiates a new WafVerificationRequirements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafVerificationRequirementsWithDefaults

`func NewWafVerificationRequirementsWithDefaults() *WafVerificationRequirements`

NewWafVerificationRequirementsWithDefaults instantiates a new WafVerificationRequirements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsVerificationDetails

`func (o *WafVerificationRequirements) GetDnsVerificationDetails() WafDnsVerificationDetails`

GetDnsVerificationDetails returns the DnsVerificationDetails field if non-nil, zero value otherwise.

### GetDnsVerificationDetailsOk

`func (o *WafVerificationRequirements) GetDnsVerificationDetailsOk() (*WafDnsVerificationDetails, bool)`

GetDnsVerificationDetailsOk returns a tuple with the DnsVerificationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsVerificationDetails

`func (o *WafVerificationRequirements) SetDnsVerificationDetails(v WafDnsVerificationDetails)`

SetDnsVerificationDetails sets DnsVerificationDetails field to given value.

### HasDnsVerificationDetails

`func (o *WafVerificationRequirements) HasDnsVerificationDetails() bool`

HasDnsVerificationDetails returns a boolean if a field has been set.

### GetHttpVerificationDetails

`func (o *WafVerificationRequirements) GetHttpVerificationDetails() WafHttpVerificationDetails`

GetHttpVerificationDetails returns the HttpVerificationDetails field if non-nil, zero value otherwise.

### GetHttpVerificationDetailsOk

`func (o *WafVerificationRequirements) GetHttpVerificationDetailsOk() (*WafHttpVerificationDetails, bool)`

GetHttpVerificationDetailsOk returns a tuple with the HttpVerificationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpVerificationDetails

`func (o *WafVerificationRequirements) SetHttpVerificationDetails(v WafHttpVerificationDetails)`

SetHttpVerificationDetails sets HttpVerificationDetails field to given value.

### HasHttpVerificationDetails

`func (o *WafVerificationRequirements) HasHttpVerificationDetails() bool`

HasHttpVerificationDetails returns a boolean if a field has been set.

### GetVerificationMethod

`func (o *WafVerificationRequirements) GetVerificationMethod() WafCertificateVerificationMethod`

GetVerificationMethod returns the VerificationMethod field if non-nil, zero value otherwise.

### GetVerificationMethodOk

`func (o *WafVerificationRequirements) GetVerificationMethodOk() (*WafCertificateVerificationMethod, bool)`

GetVerificationMethodOk returns a tuple with the VerificationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationMethod

`func (o *WafVerificationRequirements) SetVerificationMethod(v WafCertificateVerificationMethod)`

SetVerificationMethod sets VerificationMethod field to given value.

### HasVerificationMethod

`func (o *WafVerificationRequirements) HasVerificationMethod() bool`

HasVerificationMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


