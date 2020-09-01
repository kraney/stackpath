/*
 * Edge Compute
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package edge-compute

import (
	"encoding/json"
)

// V1CreateWorkloadRequest struct for V1CreateWorkloadRequest
type V1CreateWorkloadRequest struct {
	Workload *V1Workload `json:"workload,omitempty"`
}

// NewV1CreateWorkloadRequest instantiates a new V1CreateWorkloadRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1CreateWorkloadRequest() *V1CreateWorkloadRequest {
	this := V1CreateWorkloadRequest{}
	return &this
}

// NewV1CreateWorkloadRequestWithDefaults instantiates a new V1CreateWorkloadRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1CreateWorkloadRequestWithDefaults() *V1CreateWorkloadRequest {
	this := V1CreateWorkloadRequest{}
	return &this
}

// GetWorkload returns the Workload field value if set, zero value otherwise.
func (o *V1CreateWorkloadRequest) GetWorkload() V1Workload {
	if o == nil || o.Workload == nil {
		var ret V1Workload
		return ret
	}
	return *o.Workload
}

// GetWorkloadOk returns a tuple with the Workload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1CreateWorkloadRequest) GetWorkloadOk() (*V1Workload, bool) {
	if o == nil || o.Workload == nil {
		return nil, false
	}
	return o.Workload, true
}

// HasWorkload returns a boolean if a field has been set.
func (o *V1CreateWorkloadRequest) HasWorkload() bool {
	if o != nil && o.Workload != nil {
		return true
	}

	return false
}

// SetWorkload gets a reference to the given V1Workload and assigns it to the Workload field.
func (o *V1CreateWorkloadRequest) SetWorkload(v V1Workload) {
	o.Workload = &v
}

func (o V1CreateWorkloadRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Workload != nil {
		toSerialize["workload"] = o.Workload
	}
	return json.Marshal(toSerialize)
}

type NullableV1CreateWorkloadRequest struct {
	value *V1CreateWorkloadRequest
	isSet bool
}

func (v NullableV1CreateWorkloadRequest) Get() *V1CreateWorkloadRequest {
	return v.value
}

func (v *NullableV1CreateWorkloadRequest) Set(val *V1CreateWorkloadRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV1CreateWorkloadRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV1CreateWorkloadRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1CreateWorkloadRequest(val *V1CreateWorkloadRequest) *NullableV1CreateWorkloadRequest {
	return &NullableV1CreateWorkloadRequest{value: val, isSet: true}
}

func (v NullableV1CreateWorkloadRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1CreateWorkloadRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}