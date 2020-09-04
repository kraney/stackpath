# WafDnsVerificationDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsRecords** | Pointer to **[]string** | A list of DNS records that will validate domain ownership | [optional] 
**Records** | Pointer to [**[]WafDnsRecord**](wafDnsRecord.md) | A list of parsed dns records required for SSL verification | [optional] 

## Methods

### NewWafDnsVerificationDetails

`func NewWafDnsVerificationDetails() *WafDnsVerificationDetails`

NewWafDnsVerificationDetails instantiates a new WafDnsVerificationDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafDnsVerificationDetailsWithDefaults

`func NewWafDnsVerificationDetailsWithDefaults() *WafDnsVerificationDetails`

NewWafDnsVerificationDetailsWithDefaults instantiates a new WafDnsVerificationDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsRecords

`func (o *WafDnsVerificationDetails) GetDnsRecords() []string`

GetDnsRecords returns the DnsRecords field if non-nil, zero value otherwise.

### GetDnsRecordsOk

`func (o *WafDnsVerificationDetails) GetDnsRecordsOk() (*[]string, bool)`

GetDnsRecordsOk returns a tuple with the DnsRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsRecords

`func (o *WafDnsVerificationDetails) SetDnsRecords(v []string)`

SetDnsRecords sets DnsRecords field to given value.

### HasDnsRecords

`func (o *WafDnsVerificationDetails) HasDnsRecords() bool`

HasDnsRecords returns a boolean if a field has been set.

### GetRecords

`func (o *WafDnsVerificationDetails) GetRecords() []WafDnsRecord`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *WafDnsVerificationDetails) GetRecordsOk() (*[]WafDnsRecord, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *WafDnsVerificationDetails) SetRecords(v []WafDnsRecord)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *WafDnsVerificationDetails) HasRecords() bool`

HasRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


