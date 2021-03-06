/*
 * Sites
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package sites

import (
	"encoding/json"
)

// MetricsDataValue An individual metric data point
type MetricsDataValue struct {
	// The time that a data point was recorded
	UnixTime *string `json:"unixTime,omitempty"`
	// A data point's value
	Value *string `json:"value,omitempty"`
}

// NewMetricsDataValue instantiates a new MetricsDataValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricsDataValue() *MetricsDataValue {
	this := MetricsDataValue{}
	return &this
}

// NewMetricsDataValueWithDefaults instantiates a new MetricsDataValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsDataValueWithDefaults() *MetricsDataValue {
	this := MetricsDataValue{}
	return &this
}

// GetUnixTime returns the UnixTime field value if set, zero value otherwise.
func (o *MetricsDataValue) GetUnixTime() string {
	if o == nil || o.UnixTime == nil {
		var ret string
		return ret
	}
	return *o.UnixTime
}

// GetUnixTimeOk returns a tuple with the UnixTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsDataValue) GetUnixTimeOk() (*string, bool) {
	if o == nil || o.UnixTime == nil {
		return nil, false
	}
	return o.UnixTime, true
}

// HasUnixTime returns a boolean if a field has been set.
func (o *MetricsDataValue) HasUnixTime() bool {
	if o != nil && o.UnixTime != nil {
		return true
	}

	return false
}

// SetUnixTime gets a reference to the given string and assigns it to the UnixTime field.
func (o *MetricsDataValue) SetUnixTime(v string) {
	o.UnixTime = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *MetricsDataValue) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsDataValue) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *MetricsDataValue) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *MetricsDataValue) SetValue(v string) {
	o.Value = &v
}

func (o MetricsDataValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnixTime != nil {
		toSerialize["unixTime"] = o.UnixTime
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableMetricsDataValue struct {
	value *MetricsDataValue
	isSet bool
}

func (v NullableMetricsDataValue) Get() *MetricsDataValue {
	return v.value
}

func (v *NullableMetricsDataValue) Set(val *MetricsDataValue) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricsDataValue) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricsDataValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricsDataValue(val *MetricsDataValue) *NullableMetricsDataValue {
	return &NullableMetricsDataValue{value: val, isSet: true}
}

func (v NullableMetricsDataValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricsDataValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
