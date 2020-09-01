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

// MetricsData The data points in a metrics collection
type MetricsData struct {
	Matrix *DataMatrix `json:"matrix,omitempty"`
	Vector *DataVector `json:"vector,omitempty"`
}

// NewMetricsData instantiates a new MetricsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricsData() *MetricsData {
	this := MetricsData{}
	return &this
}

// NewMetricsDataWithDefaults instantiates a new MetricsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsDataWithDefaults() *MetricsData {
	this := MetricsData{}
	return &this
}

// GetMatrix returns the Matrix field value if set, zero value otherwise.
func (o *MetricsData) GetMatrix() DataMatrix {
	if o == nil || o.Matrix == nil {
		var ret DataMatrix
		return ret
	}
	return *o.Matrix
}

// GetMatrixOk returns a tuple with the Matrix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsData) GetMatrixOk() (*DataMatrix, bool) {
	if o == nil || o.Matrix == nil {
		return nil, false
	}
	return o.Matrix, true
}

// HasMatrix returns a boolean if a field has been set.
func (o *MetricsData) HasMatrix() bool {
	if o != nil && o.Matrix != nil {
		return true
	}

	return false
}

// SetMatrix gets a reference to the given DataMatrix and assigns it to the Matrix field.
func (o *MetricsData) SetMatrix(v DataMatrix) {
	o.Matrix = &v
}

// GetVector returns the Vector field value if set, zero value otherwise.
func (o *MetricsData) GetVector() DataVector {
	if o == nil || o.Vector == nil {
		var ret DataVector
		return ret
	}
	return *o.Vector
}

// GetVectorOk returns a tuple with the Vector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsData) GetVectorOk() (*DataVector, bool) {
	if o == nil || o.Vector == nil {
		return nil, false
	}
	return o.Vector, true
}

// HasVector returns a boolean if a field has been set.
func (o *MetricsData) HasVector() bool {
	if o != nil && o.Vector != nil {
		return true
	}

	return false
}

// SetVector gets a reference to the given DataVector and assigns it to the Vector field.
func (o *MetricsData) SetVector(v DataVector) {
	o.Vector = &v
}

func (o MetricsData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Matrix != nil {
		toSerialize["matrix"] = o.Matrix
	}
	if o.Vector != nil {
		toSerialize["vector"] = o.Vector
	}
	return json.Marshal(toSerialize)
}

type NullableMetricsData struct {
	value *MetricsData
	isSet bool
}

func (v NullableMetricsData) Get() *MetricsData {
	return v.value
}

func (v *NullableMetricsData) Set(val *MetricsData) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricsData) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricsData(val *MetricsData) *NullableMetricsData {
	return &NullableMetricsData{value: val, isSet: true}
}

func (v NullableMetricsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}