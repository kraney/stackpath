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

// StackCreateStackResponse struct for StackCreateStackResponse
type StackCreateStackResponse struct {
	Stack *StackStack `json:"stack,omitempty"`
}

// NewStackCreateStackResponse instantiates a new StackCreateStackResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStackCreateStackResponse() *StackCreateStackResponse {
	this := StackCreateStackResponse{}
	return &this
}

// NewStackCreateStackResponseWithDefaults instantiates a new StackCreateStackResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStackCreateStackResponseWithDefaults() *StackCreateStackResponse {
	this := StackCreateStackResponse{}
	return &this
}

// GetStack returns the Stack field value if set, zero value otherwise.
func (o *StackCreateStackResponse) GetStack() StackStack {
	if o == nil || o.Stack == nil {
		var ret StackStack
		return ret
	}
	return *o.Stack
}

// GetStackOk returns a tuple with the Stack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StackCreateStackResponse) GetStackOk() (*StackStack, bool) {
	if o == nil || o.Stack == nil {
		return nil, false
	}
	return o.Stack, true
}

// HasStack returns a boolean if a field has been set.
func (o *StackCreateStackResponse) HasStack() bool {
	if o != nil && o.Stack != nil {
		return true
	}

	return false
}

// SetStack gets a reference to the given StackStack and assigns it to the Stack field.
func (o *StackCreateStackResponse) SetStack(v StackStack) {
	o.Stack = &v
}

func (o StackCreateStackResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Stack != nil {
		toSerialize["stack"] = o.Stack
	}
	return json.Marshal(toSerialize)
}

type NullableStackCreateStackResponse struct {
	value *StackCreateStackResponse
	isSet bool
}

func (v NullableStackCreateStackResponse) Get() *StackCreateStackResponse {
	return v.value
}

func (v *NullableStackCreateStackResponse) Set(val *StackCreateStackResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStackCreateStackResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStackCreateStackResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStackCreateStackResponse(val *StackCreateStackResponse) *NullableStackCreateStackResponse {
	return &NullableStackCreateStackResponse{value: val, isSet: true}
}

func (v NullableStackCreateStackResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStackCreateStackResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}