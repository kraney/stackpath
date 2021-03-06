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

// CertificateGetLatestCertificateResponse struct for CertificateGetLatestCertificateResponse
type CertificateGetLatestCertificateResponse struct {
	Certificate *CertificateCertificate `json:"certificate,omitempty"`
}

// NewCertificateGetLatestCertificateResponse instantiates a new CertificateGetLatestCertificateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateGetLatestCertificateResponse() *CertificateGetLatestCertificateResponse {
	this := CertificateGetLatestCertificateResponse{}
	return &this
}

// NewCertificateGetLatestCertificateResponseWithDefaults instantiates a new CertificateGetLatestCertificateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateGetLatestCertificateResponseWithDefaults() *CertificateGetLatestCertificateResponse {
	this := CertificateGetLatestCertificateResponse{}
	return &this
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *CertificateGetLatestCertificateResponse) GetCertificate() CertificateCertificate {
	if o == nil || o.Certificate == nil {
		var ret CertificateCertificate
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateGetLatestCertificateResponse) GetCertificateOk() (*CertificateCertificate, bool) {
	if o == nil || o.Certificate == nil {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *CertificateGetLatestCertificateResponse) HasCertificate() bool {
	if o != nil && o.Certificate != nil {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given CertificateCertificate and assigns it to the Certificate field.
func (o *CertificateGetLatestCertificateResponse) SetCertificate(v CertificateCertificate) {
	o.Certificate = &v
}

func (o CertificateGetLatestCertificateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Certificate != nil {
		toSerialize["certificate"] = o.Certificate
	}
	return json.Marshal(toSerialize)
}

type NullableCertificateGetLatestCertificateResponse struct {
	value *CertificateGetLatestCertificateResponse
	isSet bool
}

func (v NullableCertificateGetLatestCertificateResponse) Get() *CertificateGetLatestCertificateResponse {
	return v.value
}

func (v *NullableCertificateGetLatestCertificateResponse) Set(val *CertificateGetLatestCertificateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateGetLatestCertificateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateGetLatestCertificateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateGetLatestCertificateResponse(val *CertificateGetLatestCertificateResponse) *NullableCertificateGetLatestCertificateResponse {
	return &NullableCertificateGetLatestCertificateResponse{value: val, isSet: true}
}

func (v NullableCertificateGetLatestCertificateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateGetLatestCertificateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
