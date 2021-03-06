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

// ConditionSessionRequestCountCondition Match the number of dynamic page requests made in a WAF session  The WAF injects a session cookie into protected pages to track multi-request security issues. This condition matches against the number of dynamic page requests made during that session's lifetime.
type ConditionSessionRequestCountCondition struct {
	// The number of dynamic requests in the session
	RequestCount *string `json:"requestCount,omitempty"`
}

// NewConditionSessionRequestCountCondition instantiates a new ConditionSessionRequestCountCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConditionSessionRequestCountCondition() *ConditionSessionRequestCountCondition {
	this := ConditionSessionRequestCountCondition{}
	return &this
}

// NewConditionSessionRequestCountConditionWithDefaults instantiates a new ConditionSessionRequestCountCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConditionSessionRequestCountConditionWithDefaults() *ConditionSessionRequestCountCondition {
	this := ConditionSessionRequestCountCondition{}
	return &this
}

// GetRequestCount returns the RequestCount field value if set, zero value otherwise.
func (o *ConditionSessionRequestCountCondition) GetRequestCount() string {
	if o == nil || o.RequestCount == nil {
		var ret string
		return ret
	}
	return *o.RequestCount
}

// GetRequestCountOk returns a tuple with the RequestCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionSessionRequestCountCondition) GetRequestCountOk() (*string, bool) {
	if o == nil || o.RequestCount == nil {
		return nil, false
	}
	return o.RequestCount, true
}

// HasRequestCount returns a boolean if a field has been set.
func (o *ConditionSessionRequestCountCondition) HasRequestCount() bool {
	if o != nil && o.RequestCount != nil {
		return true
	}

	return false
}

// SetRequestCount gets a reference to the given string and assigns it to the RequestCount field.
func (o *ConditionSessionRequestCountCondition) SetRequestCount(v string) {
	o.RequestCount = &v
}

func (o ConditionSessionRequestCountCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RequestCount != nil {
		toSerialize["requestCount"] = o.RequestCount
	}
	return json.Marshal(toSerialize)
}

type NullableConditionSessionRequestCountCondition struct {
	value *ConditionSessionRequestCountCondition
	isSet bool
}

func (v NullableConditionSessionRequestCountCondition) Get() *ConditionSessionRequestCountCondition {
	return v.value
}

func (v *NullableConditionSessionRequestCountCondition) Set(val *ConditionSessionRequestCountCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableConditionSessionRequestCountCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableConditionSessionRequestCountCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConditionSessionRequestCountCondition(val *ConditionSessionRequestCountCondition) *NullableConditionSessionRequestCountCondition {
	return &NullableConditionSessionRequestCountCondition{value: val, isSet: true}
}

func (v NullableConditionSessionRequestCountCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConditionSessionRequestCountCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
