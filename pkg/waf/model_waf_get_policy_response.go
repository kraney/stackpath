/*
 * Web Application Firewall
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package waf

import (
	"encoding/json"
)

// WafGetPolicyResponse A response from a request to retrieve an individual WAF policy
type WafGetPolicyResponse struct {
	Policy *SchemawafPolicy `json:"policy,omitempty"`
}

// NewWafGetPolicyResponse instantiates a new WafGetPolicyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafGetPolicyResponse() *WafGetPolicyResponse {
	this := WafGetPolicyResponse{}
	return &this
}

// NewWafGetPolicyResponseWithDefaults instantiates a new WafGetPolicyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafGetPolicyResponseWithDefaults() *WafGetPolicyResponse {
	this := WafGetPolicyResponse{}
	return &this
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *WafGetPolicyResponse) GetPolicy() SchemawafPolicy {
	if o == nil || o.Policy == nil {
		var ret SchemawafPolicy
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafGetPolicyResponse) GetPolicyOk() (*SchemawafPolicy, bool) {
	if o == nil || o.Policy == nil {
		return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *WafGetPolicyResponse) HasPolicy() bool {
	if o != nil && o.Policy != nil {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given SchemawafPolicy and assigns it to the Policy field.
func (o *WafGetPolicyResponse) SetPolicy(v SchemawafPolicy) {
	o.Policy = &v
}

func (o WafGetPolicyResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Policy != nil {
		toSerialize["policy"] = o.Policy
	}
	return json.Marshal(toSerialize)
}

type NullableWafGetPolicyResponse struct {
	value *WafGetPolicyResponse
	isSet bool
}

func (v NullableWafGetPolicyResponse) Get() *WafGetPolicyResponse {
	return v.value
}

func (v *NullableWafGetPolicyResponse) Set(val *WafGetPolicyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWafGetPolicyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWafGetPolicyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWafGetPolicyResponse(val *WafGetPolicyResponse) *NullableWafGetPolicyResponse {
	return &NullableWafGetPolicyResponse{value: val, isSet: true}
}

func (v NullableWafGetPolicyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWafGetPolicyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
