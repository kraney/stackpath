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

// MetricsRules The \"top threats\" metric, the number of events encountered per rule
type MetricsRules struct {
	// The name of the rule
	RuleName *string `json:"ruleName,omitempty"`
	// The number of requests the rule analyzed
	Count *string `json:"count,omitempty"`
}

// NewMetricsRules instantiates a new MetricsRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricsRules() *MetricsRules {
	this := MetricsRules{}
	return &this
}

// NewMetricsRulesWithDefaults instantiates a new MetricsRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsRulesWithDefaults() *MetricsRules {
	this := MetricsRules{}
	return &this
}

// GetRuleName returns the RuleName field value if set, zero value otherwise.
func (o *MetricsRules) GetRuleName() string {
	if o == nil || o.RuleName == nil {
		var ret string
		return ret
	}
	return *o.RuleName
}

// GetRuleNameOk returns a tuple with the RuleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsRules) GetRuleNameOk() (*string, bool) {
	if o == nil || o.RuleName == nil {
		return nil, false
	}
	return o.RuleName, true
}

// HasRuleName returns a boolean if a field has been set.
func (o *MetricsRules) HasRuleName() bool {
	if o != nil && o.RuleName != nil {
		return true
	}

	return false
}

// SetRuleName gets a reference to the given string and assigns it to the RuleName field.
func (o *MetricsRules) SetRuleName(v string) {
	o.RuleName = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *MetricsRules) GetCount() string {
	if o == nil || o.Count == nil {
		var ret string
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsRules) GetCountOk() (*string, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *MetricsRules) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given string and assigns it to the Count field.
func (o *MetricsRules) SetCount(v string) {
	o.Count = &v
}

func (o MetricsRules) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RuleName != nil {
		toSerialize["ruleName"] = o.RuleName
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableMetricsRules struct {
	value *MetricsRules
	isSet bool
}

func (v NullableMetricsRules) Get() *MetricsRules {
	return v.value
}

func (v *NullableMetricsRules) Set(val *MetricsRules) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricsRules) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricsRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricsRules(val *MetricsRules) *NullableMetricsRules {
	return &NullableMetricsRules{value: val, isSet: true}
}

func (v NullableMetricsRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricsRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
