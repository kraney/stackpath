/*
 * Content Delivery Network
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package cdn

import (
	"encoding/json"
)

// CdnGetCertificateVerificationDetailsResponse The response from a request to retrieve an SSL certificate's manual verification details
type CdnGetCertificateVerificationDetailsResponse struct {
	// Whether or not the end user must provide their own certificate verification
	ManualVerificationRequired *bool `json:"manualVerificationRequired,omitempty"`
	// An SSL certificate's verification requirements
	VerificationRequirements *[]CdnVerificationRequirements `json:"verificationRequirements,omitempty"`
}

// NewCdnGetCertificateVerificationDetailsResponse instantiates a new CdnGetCertificateVerificationDetailsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCdnGetCertificateVerificationDetailsResponse() *CdnGetCertificateVerificationDetailsResponse {
	this := CdnGetCertificateVerificationDetailsResponse{}
	return &this
}

// NewCdnGetCertificateVerificationDetailsResponseWithDefaults instantiates a new CdnGetCertificateVerificationDetailsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCdnGetCertificateVerificationDetailsResponseWithDefaults() *CdnGetCertificateVerificationDetailsResponse {
	this := CdnGetCertificateVerificationDetailsResponse{}
	return &this
}

// GetManualVerificationRequired returns the ManualVerificationRequired field value if set, zero value otherwise.
func (o *CdnGetCertificateVerificationDetailsResponse) GetManualVerificationRequired() bool {
	if o == nil || o.ManualVerificationRequired == nil {
		var ret bool
		return ret
	}
	return *o.ManualVerificationRequired
}

// GetManualVerificationRequiredOk returns a tuple with the ManualVerificationRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdnGetCertificateVerificationDetailsResponse) GetManualVerificationRequiredOk() (*bool, bool) {
	if o == nil || o.ManualVerificationRequired == nil {
		return nil, false
	}
	return o.ManualVerificationRequired, true
}

// HasManualVerificationRequired returns a boolean if a field has been set.
func (o *CdnGetCertificateVerificationDetailsResponse) HasManualVerificationRequired() bool {
	if o != nil && o.ManualVerificationRequired != nil {
		return true
	}

	return false
}

// SetManualVerificationRequired gets a reference to the given bool and assigns it to the ManualVerificationRequired field.
func (o *CdnGetCertificateVerificationDetailsResponse) SetManualVerificationRequired(v bool) {
	o.ManualVerificationRequired = &v
}

// GetVerificationRequirements returns the VerificationRequirements field value if set, zero value otherwise.
func (o *CdnGetCertificateVerificationDetailsResponse) GetVerificationRequirements() []CdnVerificationRequirements {
	if o == nil || o.VerificationRequirements == nil {
		var ret []CdnVerificationRequirements
		return ret
	}
	return *o.VerificationRequirements
}

// GetVerificationRequirementsOk returns a tuple with the VerificationRequirements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdnGetCertificateVerificationDetailsResponse) GetVerificationRequirementsOk() (*[]CdnVerificationRequirements, bool) {
	if o == nil || o.VerificationRequirements == nil {
		return nil, false
	}
	return o.VerificationRequirements, true
}

// HasVerificationRequirements returns a boolean if a field has been set.
func (o *CdnGetCertificateVerificationDetailsResponse) HasVerificationRequirements() bool {
	if o != nil && o.VerificationRequirements != nil {
		return true
	}

	return false
}

// SetVerificationRequirements gets a reference to the given []CdnVerificationRequirements and assigns it to the VerificationRequirements field.
func (o *CdnGetCertificateVerificationDetailsResponse) SetVerificationRequirements(v []CdnVerificationRequirements) {
	o.VerificationRequirements = &v
}

func (o CdnGetCertificateVerificationDetailsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ManualVerificationRequired != nil {
		toSerialize["manualVerificationRequired"] = o.ManualVerificationRequired
	}
	if o.VerificationRequirements != nil {
		toSerialize["verificationRequirements"] = o.VerificationRequirements
	}
	return json.Marshal(toSerialize)
}

type NullableCdnGetCertificateVerificationDetailsResponse struct {
	value *CdnGetCertificateVerificationDetailsResponse
	isSet bool
}

func (v NullableCdnGetCertificateVerificationDetailsResponse) Get() *CdnGetCertificateVerificationDetailsResponse {
	return v.value
}

func (v *NullableCdnGetCertificateVerificationDetailsResponse) Set(val *CdnGetCertificateVerificationDetailsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCdnGetCertificateVerificationDetailsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCdnGetCertificateVerificationDetailsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCdnGetCertificateVerificationDetailsResponse(val *CdnGetCertificateVerificationDetailsResponse) *NullableCdnGetCertificateVerificationDetailsResponse {
	return &NullableCdnGetCertificateVerificationDetailsResponse{value: val, isSet: true}
}

func (v NullableCdnGetCertificateVerificationDetailsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCdnGetCertificateVerificationDetailsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}