/*
 * Accounts and Users
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package accounts_and_users

import (
	"encoding/json"
)

// IdentitySetIAMPolicyRequest struct for IdentitySetIAMPolicyRequest
type IdentitySetIAMPolicyRequest struct {
	Policy *IamPolicy `json:"policy,omitempty"`
}

// NewIdentitySetIAMPolicyRequest instantiates a new IdentitySetIAMPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentitySetIAMPolicyRequest() *IdentitySetIAMPolicyRequest {
	this := IdentitySetIAMPolicyRequest{}
	return &this
}

// NewIdentitySetIAMPolicyRequestWithDefaults instantiates a new IdentitySetIAMPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentitySetIAMPolicyRequestWithDefaults() *IdentitySetIAMPolicyRequest {
	this := IdentitySetIAMPolicyRequest{}
	return &this
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *IdentitySetIAMPolicyRequest) GetPolicy() IamPolicy {
	if o == nil || o.Policy == nil {
		var ret IamPolicy
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySetIAMPolicyRequest) GetPolicyOk() (*IamPolicy, bool) {
	if o == nil || o.Policy == nil {
		return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *IdentitySetIAMPolicyRequest) HasPolicy() bool {
	if o != nil && o.Policy != nil {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given IamPolicy and assigns it to the Policy field.
func (o *IdentitySetIAMPolicyRequest) SetPolicy(v IamPolicy) {
	o.Policy = &v
}

func (o IdentitySetIAMPolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Policy != nil {
		toSerialize["policy"] = o.Policy
	}
	return json.Marshal(toSerialize)
}

type NullableIdentitySetIAMPolicyRequest struct {
	value *IdentitySetIAMPolicyRequest
	isSet bool
}

func (v NullableIdentitySetIAMPolicyRequest) Get() *IdentitySetIAMPolicyRequest {
	return v.value
}

func (v *NullableIdentitySetIAMPolicyRequest) Set(val *IdentitySetIAMPolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentitySetIAMPolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentitySetIAMPolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentitySetIAMPolicyRequest(val *IdentitySetIAMPolicyRequest) *NullableIdentitySetIAMPolicyRequest {
	return &NullableIdentitySetIAMPolicyRequest{value: val, isSet: true}
}

func (v NullableIdentitySetIAMPolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentitySetIAMPolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
