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

// V1ImageStatus Which capture status an image is currently in  - IMAGE_STATUS_UNKNOWN: The image status is unknown  - PENDING: The image is pending creation  - PROCESSING: The image is processing  - READY: The image is ready  - FAILED: The image failed to be created
type V1ImageStatus string

// List of v1ImageStatus
const (
	IMAGE_STATUS_UNKNOWN V1ImageStatus = "IMAGE_STATUS_UNKNOWN"
	PENDING V1ImageStatus = "PENDING"
	PROCESSING V1ImageStatus = "PROCESSING"
	READY V1ImageStatus = "READY"
	FAILED V1ImageStatus = "FAILED"
)

// Ptr returns reference to v1ImageStatus value
func (v V1ImageStatus) Ptr() *V1ImageStatus {
	return &v
}


type NullableV1ImageStatus struct {
	value *V1ImageStatus
	isSet bool
}

func (v NullableV1ImageStatus) Get() *V1ImageStatus {
	return v.value
}

func (v *NullableV1ImageStatus) Set(val *V1ImageStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ImageStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ImageStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ImageStatus(val *V1ImageStatus) *NullableV1ImageStatus {
	return &NullableV1ImageStatus{value: val, isSet: true}
}

func (v NullableV1ImageStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ImageStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}