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

// GetMetricsResponseMetricSample The data points associated with a series of metrics
type GetMetricsResponseMetricSample struct {
	// An individual data point
	Values *[]float64 `json:"values,omitempty"`
}

// NewGetMetricsResponseMetricSample instantiates a new GetMetricsResponseMetricSample object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMetricsResponseMetricSample() *GetMetricsResponseMetricSample {
	this := GetMetricsResponseMetricSample{}
	return &this
}

// NewGetMetricsResponseMetricSampleWithDefaults instantiates a new GetMetricsResponseMetricSample object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMetricsResponseMetricSampleWithDefaults() *GetMetricsResponseMetricSample {
	this := GetMetricsResponseMetricSample{}
	return &this
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *GetMetricsResponseMetricSample) GetValues() []float64 {
	if o == nil || o.Values == nil {
		var ret []float64
		return ret
	}
	return *o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMetricsResponseMetricSample) GetValuesOk() (*[]float64, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *GetMetricsResponseMetricSample) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given []float64 and assigns it to the Values field.
func (o *GetMetricsResponseMetricSample) SetValues(v []float64) {
	o.Values = &v
}

func (o GetMetricsResponseMetricSample) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}
	return json.Marshal(toSerialize)
}

type NullableGetMetricsResponseMetricSample struct {
	value *GetMetricsResponseMetricSample
	isSet bool
}

func (v NullableGetMetricsResponseMetricSample) Get() *GetMetricsResponseMetricSample {
	return v.value
}

func (v *NullableGetMetricsResponseMetricSample) Set(val *GetMetricsResponseMetricSample) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMetricsResponseMetricSample) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMetricsResponseMetricSample) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMetricsResponseMetricSample(val *GetMetricsResponseMetricSample) *NullableGetMetricsResponseMetricSample {
	return &NullableGetMetricsResponseMetricSample{value: val, isSet: true}
}

func (v NullableGetMetricsResponseMetricSample) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMetricsResponseMetricSample) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}