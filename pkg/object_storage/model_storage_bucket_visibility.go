/*
 * Object Storage
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package object_storage

import (
	"encoding/json"
)

// StorageBucketVisibility - PRIVATE: The bucket is private and only accessibly with credentials  - PUBLIC: The bucket is public and accessible over the internet
type StorageBucketVisibility string

// List of storageBucketVisibility
const (
	STORAGEBUCKETVISIBILITY_PRIVATE StorageBucketVisibility = "PRIVATE"
	STORAGEBUCKETVISIBILITY_PUBLIC StorageBucketVisibility = "PUBLIC"
)

// Ptr returns reference to storageBucketVisibility value
func (v StorageBucketVisibility) Ptr() *StorageBucketVisibility {
	return &v
}


type NullableStorageBucketVisibility struct {
	value *StorageBucketVisibility
	isSet bool
}

func (v NullableStorageBucketVisibility) Get() *StorageBucketVisibility {
	return v.value
}

func (v *NullableStorageBucketVisibility) Set(val *StorageBucketVisibility) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageBucketVisibility) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageBucketVisibility) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageBucketVisibility(val *StorageBucketVisibility) *NullableStorageBucketVisibility {
	return &NullableStorageBucketVisibility{value: val, isSet: true}
}

func (v NullableStorageBucketVisibility) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageBucketVisibility) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
