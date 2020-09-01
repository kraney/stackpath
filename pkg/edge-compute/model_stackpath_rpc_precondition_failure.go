/*
 * Edge Compute
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package edge-compute

import (
	"encoding/json"
)

// StackpathRpcPreconditionFailure struct for StackpathRpcPreconditionFailure
type StackpathRpcPreconditionFailure struct {
	ApiStatusDetail
	Violations *[]StackpathRpcPreconditionFailureViolation `json:"violations,omitempty"`
}

// NewStackpathRpcPreconditionFailure instantiates a new StackpathRpcPreconditionFailure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStackpathRpcPreconditionFailure() *StackpathRpcPreconditionFailure {
	this := StackpathRpcPreconditionFailure{}
	return &this
}

// NewStackpathRpcPreconditionFailureWithDefaults instantiates a new StackpathRpcPreconditionFailure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStackpathRpcPreconditionFailureWithDefaults() *StackpathRpcPreconditionFailure {
	this := StackpathRpcPreconditionFailure{}
	return &this
}

// GetViolations returns the Violations field value if set, zero value otherwise.
func (o *StackpathRpcPreconditionFailure) GetViolations() []StackpathRpcPreconditionFailureViolation {
	if o == nil || o.Violations == nil {
		var ret []StackpathRpcPreconditionFailureViolation
		return ret
	}
	return *o.Violations
}

// GetViolationsOk returns a tuple with the Violations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StackpathRpcPreconditionFailure) GetViolationsOk() (*[]StackpathRpcPreconditionFailureViolation, bool) {
	if o == nil || o.Violations == nil {
		return nil, false
	}
	return o.Violations, true
}

// HasViolations returns a boolean if a field has been set.
func (o *StackpathRpcPreconditionFailure) HasViolations() bool {
	if o != nil && o.Violations != nil {
		return true
	}

	return false
}

// SetViolations gets a reference to the given []StackpathRpcPreconditionFailureViolation and assigns it to the Violations field.
func (o *StackpathRpcPreconditionFailure) SetViolations(v []StackpathRpcPreconditionFailureViolation) {
	o.Violations = &v
}

func (o StackpathRpcPreconditionFailure) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedApiStatusDetail, errApiStatusDetail := json.Marshal(o.ApiStatusDetail)
	if errApiStatusDetail != nil {
		return []byte{}, errApiStatusDetail
	}
	errApiStatusDetail = json.Unmarshal([]byte(serializedApiStatusDetail), &toSerialize)
	if errApiStatusDetail != nil {
		return []byte{}, errApiStatusDetail
	}
	if o.Violations != nil {
		toSerialize["violations"] = o.Violations
	}
	return json.Marshal(toSerialize)
}

type NullableStackpathRpcPreconditionFailure struct {
	value *StackpathRpcPreconditionFailure
	isSet bool
}

func (v NullableStackpathRpcPreconditionFailure) Get() *StackpathRpcPreconditionFailure {
	return v.value
}

func (v *NullableStackpathRpcPreconditionFailure) Set(val *StackpathRpcPreconditionFailure) {
	v.value = val
	v.isSet = true
}

func (v NullableStackpathRpcPreconditionFailure) IsSet() bool {
	return v.isSet
}

func (v *NullableStackpathRpcPreconditionFailure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStackpathRpcPreconditionFailure(val *StackpathRpcPreconditionFailure) *NullableStackpathRpcPreconditionFailure {
	return &NullableStackpathRpcPreconditionFailure{value: val, isSet: true}
}

func (v NullableStackpathRpcPreconditionFailure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStackpathRpcPreconditionFailure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}