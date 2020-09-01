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

// WafGetPolicyGroupsResponse A response from a request to retrieve a WAF site's policy groups
type WafGetPolicyGroupsResponse struct {
	// The requested WAF policy groups
	PolicyGroups *[]WafPolicyGroup `json:"policyGroups,omitempty"`
}

// NewWafGetPolicyGroupsResponse instantiates a new WafGetPolicyGroupsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafGetPolicyGroupsResponse() *WafGetPolicyGroupsResponse {
	this := WafGetPolicyGroupsResponse{}
	return &this
}

// NewWafGetPolicyGroupsResponseWithDefaults instantiates a new WafGetPolicyGroupsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafGetPolicyGroupsResponseWithDefaults() *WafGetPolicyGroupsResponse {
	this := WafGetPolicyGroupsResponse{}
	return &this
}

// GetPolicyGroups returns the PolicyGroups field value if set, zero value otherwise.
func (o *WafGetPolicyGroupsResponse) GetPolicyGroups() []WafPolicyGroup {
	if o == nil || o.PolicyGroups == nil {
		var ret []WafPolicyGroup
		return ret
	}
	return *o.PolicyGroups
}

// GetPolicyGroupsOk returns a tuple with the PolicyGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafGetPolicyGroupsResponse) GetPolicyGroupsOk() (*[]WafPolicyGroup, bool) {
	if o == nil || o.PolicyGroups == nil {
		return nil, false
	}
	return o.PolicyGroups, true
}

// HasPolicyGroups returns a boolean if a field has been set.
func (o *WafGetPolicyGroupsResponse) HasPolicyGroups() bool {
	if o != nil && o.PolicyGroups != nil {
		return true
	}

	return false
}

// SetPolicyGroups gets a reference to the given []WafPolicyGroup and assigns it to the PolicyGroups field.
func (o *WafGetPolicyGroupsResponse) SetPolicyGroups(v []WafPolicyGroup) {
	o.PolicyGroups = &v
}

func (o WafGetPolicyGroupsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PolicyGroups != nil {
		toSerialize["policyGroups"] = o.PolicyGroups
	}
	return json.Marshal(toSerialize)
}

type NullableWafGetPolicyGroupsResponse struct {
	value *WafGetPolicyGroupsResponse
	isSet bool
}

func (v NullableWafGetPolicyGroupsResponse) Get() *WafGetPolicyGroupsResponse {
	return v.value
}

func (v *NullableWafGetPolicyGroupsResponse) Set(val *WafGetPolicyGroupsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWafGetPolicyGroupsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWafGetPolicyGroupsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWafGetPolicyGroupsResponse(val *WafGetPolicyGroupsResponse) *NullableWafGetPolicyGroupsResponse {
	return &NullableWafGetPolicyGroupsResponse{value: val, isSet: true}
}

func (v NullableWafGetPolicyGroupsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWafGetPolicyGroupsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}