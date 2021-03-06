/*
 * Edge Compute
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package edge_compute

import (
	"encoding/json"
)

// V1MetricSpec struct for V1MetricSpec
type V1MetricSpec struct {
	// The metric to measure against
	Metric *string `json:"metric,omitempty"`
	// The target average value for the metric measured across all instances within a deployment, expressed as a quantity.
	AverageValue *string `json:"averageValue,omitempty"`
	// The target average value for the metric measured across all instances within a deployment, expressed as a percentage of requested resources.
	AverageUtilization *int32 `json:"averageUtilization,omitempty"`
}

// NewV1MetricSpec instantiates a new V1MetricSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1MetricSpec() *V1MetricSpec {
	this := V1MetricSpec{}
	return &this
}

// NewV1MetricSpecWithDefaults instantiates a new V1MetricSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1MetricSpecWithDefaults() *V1MetricSpec {
	this := V1MetricSpec{}
	return &this
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *V1MetricSpec) GetMetric() string {
	if o == nil || o.Metric == nil {
		var ret string
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1MetricSpec) GetMetricOk() (*string, bool) {
	if o == nil || o.Metric == nil {
		return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *V1MetricSpec) HasMetric() bool {
	if o != nil && o.Metric != nil {
		return true
	}

	return false
}

// SetMetric gets a reference to the given string and assigns it to the Metric field.
func (o *V1MetricSpec) SetMetric(v string) {
	o.Metric = &v
}

// GetAverageValue returns the AverageValue field value if set, zero value otherwise.
func (o *V1MetricSpec) GetAverageValue() string {
	if o == nil || o.AverageValue == nil {
		var ret string
		return ret
	}
	return *o.AverageValue
}

// GetAverageValueOk returns a tuple with the AverageValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1MetricSpec) GetAverageValueOk() (*string, bool) {
	if o == nil || o.AverageValue == nil {
		return nil, false
	}
	return o.AverageValue, true
}

// HasAverageValue returns a boolean if a field has been set.
func (o *V1MetricSpec) HasAverageValue() bool {
	if o != nil && o.AverageValue != nil {
		return true
	}

	return false
}

// SetAverageValue gets a reference to the given string and assigns it to the AverageValue field.
func (o *V1MetricSpec) SetAverageValue(v string) {
	o.AverageValue = &v
}

// GetAverageUtilization returns the AverageUtilization field value if set, zero value otherwise.
func (o *V1MetricSpec) GetAverageUtilization() int32 {
	if o == nil || o.AverageUtilization == nil {
		var ret int32
		return ret
	}
	return *o.AverageUtilization
}

// GetAverageUtilizationOk returns a tuple with the AverageUtilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1MetricSpec) GetAverageUtilizationOk() (*int32, bool) {
	if o == nil || o.AverageUtilization == nil {
		return nil, false
	}
	return o.AverageUtilization, true
}

// HasAverageUtilization returns a boolean if a field has been set.
func (o *V1MetricSpec) HasAverageUtilization() bool {
	if o != nil && o.AverageUtilization != nil {
		return true
	}

	return false
}

// SetAverageUtilization gets a reference to the given int32 and assigns it to the AverageUtilization field.
func (o *V1MetricSpec) SetAverageUtilization(v int32) {
	o.AverageUtilization = &v
}

func (o V1MetricSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Metric != nil {
		toSerialize["metric"] = o.Metric
	}
	if o.AverageValue != nil {
		toSerialize["averageValue"] = o.AverageValue
	}
	if o.AverageUtilization != nil {
		toSerialize["averageUtilization"] = o.AverageUtilization
	}
	return json.Marshal(toSerialize)
}

type NullableV1MetricSpec struct {
	value *V1MetricSpec
	isSet bool
}

func (v NullableV1MetricSpec) Get() *V1MetricSpec {
	return v.value
}

func (v *NullableV1MetricSpec) Set(val *V1MetricSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableV1MetricSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableV1MetricSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1MetricSpec(val *V1MetricSpec) *NullableV1MetricSpec {
	return &NullableV1MetricSpec{value: val, isSet: true}
}

func (v NullableV1MetricSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1MetricSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
