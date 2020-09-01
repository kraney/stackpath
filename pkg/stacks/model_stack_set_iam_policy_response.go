/*
 * Stacks
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package stacks

import (
	"encoding/json"
)

// StackSetIAMPolicyResponse A response from a request to set a StackPath stack's policy
type StackSetIAMPolicyResponse struct {
	// The ID of the StackPath stack the policy is on
	StackId *string `json:"stackId,omitempty"`
	Policy *IamPolicy `json:"policy,omitempty"`
}

// NewStackSetIAMPolicyResponse instantiates a new StackSetIAMPolicyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStackSetIAMPolicyResponse() *StackSetIAMPolicyResponse {
	this := StackSetIAMPolicyResponse{}
	return &this
}

// NewStackSetIAMPolicyResponseWithDefaults instantiates a new StackSetIAMPolicyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStackSetIAMPolicyResponseWithDefaults() *StackSetIAMPolicyResponse {
	this := StackSetIAMPolicyResponse{}
	return &this
}

// GetStackId returns the StackId field value if set, zero value otherwise.
func (o *StackSetIAMPolicyResponse) GetStackId() string {
	if o == nil || o.StackId == nil {
		var ret string
		return ret
	}
	return *o.StackId
}

// GetStackIdOk returns a tuple with the StackId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StackSetIAMPolicyResponse) GetStackIdOk() (*string, bool) {
	if o == nil || o.StackId == nil {
		return nil, false
	}
	return o.StackId, true
}

// HasStackId returns a boolean if a field has been set.
func (o *StackSetIAMPolicyResponse) HasStackId() bool {
	if o != nil && o.StackId != nil {
		return true
	}

	return false
}

// SetStackId gets a reference to the given string and assigns it to the StackId field.
func (o *StackSetIAMPolicyResponse) SetStackId(v string) {
	o.StackId = &v
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *StackSetIAMPolicyResponse) GetPolicy() IamPolicy {
	if o == nil || o.Policy == nil {
		var ret IamPolicy
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StackSetIAMPolicyResponse) GetPolicyOk() (*IamPolicy, bool) {
	if o == nil || o.Policy == nil {
		return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *StackSetIAMPolicyResponse) HasPolicy() bool {
	if o != nil && o.Policy != nil {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given IamPolicy and assigns it to the Policy field.
func (o *StackSetIAMPolicyResponse) SetPolicy(v IamPolicy) {
	o.Policy = &v
}

func (o StackSetIAMPolicyResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StackId != nil {
		toSerialize["stackId"] = o.StackId
	}
	if o.Policy != nil {
		toSerialize["policy"] = o.Policy
	}
	return json.Marshal(toSerialize)
}

type NullableStackSetIAMPolicyResponse struct {
	value *StackSetIAMPolicyResponse
	isSet bool
}

func (v NullableStackSetIAMPolicyResponse) Get() *StackSetIAMPolicyResponse {
	return v.value
}

func (v *NullableStackSetIAMPolicyResponse) Set(val *StackSetIAMPolicyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStackSetIAMPolicyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStackSetIAMPolicyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStackSetIAMPolicyResponse(val *StackSetIAMPolicyResponse) *NullableStackSetIAMPolicyResponse {
	return &NullableStackSetIAMPolicyResponse{value: val, isSet: true}
}

func (v NullableStackSetIAMPolicyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStackSetIAMPolicyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}