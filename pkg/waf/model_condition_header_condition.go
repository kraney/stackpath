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

// ConditionHeaderCondition Match an incoming request header
type ConditionHeaderCondition struct {
	// The request header name
	Header *string `json:"header,omitempty"`
	// The request header value
	Value *string `json:"value,omitempty"`
	// Whether or not to perform an exact match of the request header and value
	ExactMatch *bool `json:"exactMatch,omitempty"`
}

// NewConditionHeaderCondition instantiates a new ConditionHeaderCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConditionHeaderCondition() *ConditionHeaderCondition {
	this := ConditionHeaderCondition{}
	return &this
}

// NewConditionHeaderConditionWithDefaults instantiates a new ConditionHeaderCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConditionHeaderConditionWithDefaults() *ConditionHeaderCondition {
	this := ConditionHeaderCondition{}
	return &this
}

// GetHeader returns the Header field value if set, zero value otherwise.
func (o *ConditionHeaderCondition) GetHeader() string {
	if o == nil || o.Header == nil {
		var ret string
		return ret
	}
	return *o.Header
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionHeaderCondition) GetHeaderOk() (*string, bool) {
	if o == nil || o.Header == nil {
		return nil, false
	}
	return o.Header, true
}

// HasHeader returns a boolean if a field has been set.
func (o *ConditionHeaderCondition) HasHeader() bool {
	if o != nil && o.Header != nil {
		return true
	}

	return false
}

// SetHeader gets a reference to the given string and assigns it to the Header field.
func (o *ConditionHeaderCondition) SetHeader(v string) {
	o.Header = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ConditionHeaderCondition) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionHeaderCondition) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ConditionHeaderCondition) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ConditionHeaderCondition) SetValue(v string) {
	o.Value = &v
}

// GetExactMatch returns the ExactMatch field value if set, zero value otherwise.
func (o *ConditionHeaderCondition) GetExactMatch() bool {
	if o == nil || o.ExactMatch == nil {
		var ret bool
		return ret
	}
	return *o.ExactMatch
}

// GetExactMatchOk returns a tuple with the ExactMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionHeaderCondition) GetExactMatchOk() (*bool, bool) {
	if o == nil || o.ExactMatch == nil {
		return nil, false
	}
	return o.ExactMatch, true
}

// HasExactMatch returns a boolean if a field has been set.
func (o *ConditionHeaderCondition) HasExactMatch() bool {
	if o != nil && o.ExactMatch != nil {
		return true
	}

	return false
}

// SetExactMatch gets a reference to the given bool and assigns it to the ExactMatch field.
func (o *ConditionHeaderCondition) SetExactMatch(v bool) {
	o.ExactMatch = &v
}

func (o ConditionHeaderCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Header != nil {
		toSerialize["header"] = o.Header
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.ExactMatch != nil {
		toSerialize["exactMatch"] = o.ExactMatch
	}
	return json.Marshal(toSerialize)
}

type NullableConditionHeaderCondition struct {
	value *ConditionHeaderCondition
	isSet bool
}

func (v NullableConditionHeaderCondition) Get() *ConditionHeaderCondition {
	return v.value
}

func (v *NullableConditionHeaderCondition) Set(val *ConditionHeaderCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableConditionHeaderCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableConditionHeaderCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConditionHeaderCondition(val *ConditionHeaderCondition) *NullableConditionHeaderCondition {
	return &NullableConditionHeaderCondition{value: val, isSet: true}
}

func (v NullableConditionHeaderCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConditionHeaderCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
