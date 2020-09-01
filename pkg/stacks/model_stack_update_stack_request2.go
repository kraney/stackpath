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

// StackUpdateStackRequest2 struct for StackUpdateStackRequest2
type StackUpdateStackRequest2 struct {
	// The stack's new name
	Name *string `json:"name,omitempty"`
}

// NewStackUpdateStackRequest2 instantiates a new StackUpdateStackRequest2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStackUpdateStackRequest2() *StackUpdateStackRequest2 {
	this := StackUpdateStackRequest2{}
	return &this
}

// NewStackUpdateStackRequest2WithDefaults instantiates a new StackUpdateStackRequest2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStackUpdateStackRequest2WithDefaults() *StackUpdateStackRequest2 {
	this := StackUpdateStackRequest2{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StackUpdateStackRequest2) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StackUpdateStackRequest2) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StackUpdateStackRequest2) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StackUpdateStackRequest2) SetName(v string) {
	o.Name = &v
}

func (o StackUpdateStackRequest2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableStackUpdateStackRequest2 struct {
	value *StackUpdateStackRequest2
	isSet bool
}

func (v NullableStackUpdateStackRequest2) Get() *StackUpdateStackRequest2 {
	return v.value
}

func (v *NullableStackUpdateStackRequest2) Set(val *StackUpdateStackRequest2) {
	v.value = val
	v.isSet = true
}

func (v NullableStackUpdateStackRequest2) IsSet() bool {
	return v.isSet
}

func (v *NullableStackUpdateStackRequest2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStackUpdateStackRequest2(val *StackUpdateStackRequest2) *NullableStackUpdateStackRequest2 {
	return &NullableStackUpdateStackRequest2{value: val, isSet: true}
}

func (v NullableStackUpdateStackRequest2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStackUpdateStackRequest2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}