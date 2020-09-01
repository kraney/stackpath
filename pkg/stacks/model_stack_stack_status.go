/*
 * Stacks
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package stacks

import (
	"encoding/json"
)

// StackStackStatus A stack's status
type StackStackStatus string

// List of stackStackStatus
const (
	PENDING StackStackStatus = "PENDING"
	ACTIVE StackStackStatus = "ACTIVE"
	DISABLED StackStackStatus = "DISABLED"
	SUSPENDED StackStackStatus = "SUSPENDED"
	BILLING_SUSPENDED StackStackStatus = "BILLING_SUSPENDED"
	CANCELLED StackStackStatus = "CANCELLED"
	DELETED StackStackStatus = "DELETED"
)

// Ptr returns reference to stackStackStatus value
func (v StackStackStatus) Ptr() *StackStackStatus {
	return &v
}


type NullableStackStackStatus struct {
	value *StackStackStatus
	isSet bool
}

func (v NullableStackStackStatus) Get() *StackStackStatus {
	return v.value
}

func (v *NullableStackStackStatus) Set(val *StackStackStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableStackStackStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableStackStackStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStackStackStatus(val *StackStackStatus) *NullableStackStackStatus {
	return &NullableStackStackStatus{value: val, isSet: true}
}

func (v NullableStackStackStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStackStackStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}