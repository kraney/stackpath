# CertificateDnsVerificationDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsRecords** | Pointer to **[]string** | A list of bind formatted dns records required for SSL verification | [optional] 
**Records** | Pointer to [**[]CertificateDnsRecord**](certificateDnsRecord.md) | A list of parsed dns records required for SSL verification | [optional] 

## Methods

### NewCertificateDnsVerificationDetails

`func NewCertificateDnsVerificationDetails() *CertificateDnsVerificationDetails`

NewCertificateDnsVerificationDetails instantiates a new CertificateDnsVerificationDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateDnsVerificationDetailsWithDefaults

`func NewCertificateDnsVerificationDetailsWithDefaults() *CertificateDnsVerificationDetails`

NewCertificateDnsVerificationDetailsWithDefaults instantiates a new CertificateDnsVerificationDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsRecords

`func (o *CertificateDnsVerificationDetails) GetDnsRecords() []string`

GetDnsRecords returns the DnsRecords field if non-nil, zero value otherwise.

### GetDnsRecordsOk

`func (o *CertificateDnsVerificationDetails) GetDnsRecordsOk() (*[]string, bool)`

GetDnsRecordsOk returns a tuple with the DnsRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsRecords

`func (o *CertificateDnsVerificationDetails) SetDnsRecords(v []string)`

SetDnsRecords sets DnsRecords field to given value.

### HasDnsRecords

`func (o *CertificateDnsVerificationDetails) HasDnsRecords() bool`

HasDnsRecords returns a boolean if a field has been set.

### GetRecords

`func (o *CertificateDnsVerificationDetails) GetRecords() []CertificateDnsRecord`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *CertificateDnsVerificationDetails) GetRecordsOk() (*[]CertificateDnsRecord, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *CertificateDnsVerificationDetails) SetRecords(v []CertificateDnsRecord)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *CertificateDnsVerificationDetails) HasRecords() bool`

HasRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


