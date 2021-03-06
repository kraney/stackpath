/*
 * SSL
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ssl

import (
	"encoding/json"
)

// CertificateDnsVerificationDetails struct for CertificateDnsVerificationDetails
type CertificateDnsVerificationDetails struct {
	// A list of bind formatted dns records required for SSL verification
	DnsRecords *[]string `json:"dnsRecords,omitempty"`
	// A list of parsed dns records required for SSL verification
	Records *[]CertificateDnsRecord `json:"records,omitempty"`
}

// NewCertificateDnsVerificationDetails instantiates a new CertificateDnsVerificationDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateDnsVerificationDetails() *CertificateDnsVerificationDetails {
	this := CertificateDnsVerificationDetails{}
	return &this
}

// NewCertificateDnsVerificationDetailsWithDefaults instantiates a new CertificateDnsVerificationDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateDnsVerificationDetailsWithDefaults() *CertificateDnsVerificationDetails {
	this := CertificateDnsVerificationDetails{}
	return &this
}

// GetDnsRecords returns the DnsRecords field value if set, zero value otherwise.
func (o *CertificateDnsVerificationDetails) GetDnsRecords() []string {
	if o == nil || o.DnsRecords == nil {
		var ret []string
		return ret
	}
	return *o.DnsRecords
}

// GetDnsRecordsOk returns a tuple with the DnsRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateDnsVerificationDetails) GetDnsRecordsOk() (*[]string, bool) {
	if o == nil || o.DnsRecords == nil {
		return nil, false
	}
	return o.DnsRecords, true
}

// HasDnsRecords returns a boolean if a field has been set.
func (o *CertificateDnsVerificationDetails) HasDnsRecords() bool {
	if o != nil && o.DnsRecords != nil {
		return true
	}

	return false
}

// SetDnsRecords gets a reference to the given []string and assigns it to the DnsRecords field.
func (o *CertificateDnsVerificationDetails) SetDnsRecords(v []string) {
	o.DnsRecords = &v
}

// GetRecords returns the Records field value if set, zero value otherwise.
func (o *CertificateDnsVerificationDetails) GetRecords() []CertificateDnsRecord {
	if o == nil || o.Records == nil {
		var ret []CertificateDnsRecord
		return ret
	}
	return *o.Records
}

// GetRecordsOk returns a tuple with the Records field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateDnsVerificationDetails) GetRecordsOk() (*[]CertificateDnsRecord, bool) {
	if o == nil || o.Records == nil {
		return nil, false
	}
	return o.Records, true
}

// HasRecords returns a boolean if a field has been set.
func (o *CertificateDnsVerificationDetails) HasRecords() bool {
	if o != nil && o.Records != nil {
		return true
	}

	return false
}

// SetRecords gets a reference to the given []CertificateDnsRecord and assigns it to the Records field.
func (o *CertificateDnsVerificationDetails) SetRecords(v []CertificateDnsRecord) {
	o.Records = &v
}

func (o CertificateDnsVerificationDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DnsRecords != nil {
		toSerialize["dnsRecords"] = o.DnsRecords
	}
	if o.Records != nil {
		toSerialize["records"] = o.Records
	}
	return json.Marshal(toSerialize)
}

type NullableCertificateDnsVerificationDetails struct {
	value *CertificateDnsVerificationDetails
	isSet bool
}

func (v NullableCertificateDnsVerificationDetails) Get() *CertificateDnsVerificationDetails {
	return v.value
}

func (v *NullableCertificateDnsVerificationDetails) Set(val *CertificateDnsVerificationDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateDnsVerificationDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateDnsVerificationDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateDnsVerificationDetails(val *CertificateDnsVerificationDetails) *NullableCertificateDnsVerificationDetails {
	return &NullableCertificateDnsVerificationDetails{value: val, isSet: true}
}

func (v NullableCertificateDnsVerificationDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateDnsVerificationDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
