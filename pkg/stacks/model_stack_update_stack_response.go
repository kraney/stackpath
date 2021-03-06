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

// StackUpdateStackResponse struct for StackUpdateStackResponse
type StackUpdateStackResponse struct {
	Stack *StackStack `json:"stack,omitempty"`
}

// NewStackUpdateStackResponse instantiates a new StackUpdateStackResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStackUpdateStackResponse() *StackUpdateStackResponse {
	this := StackUpdateStackResponse{}
	return &this
}

// NewStackUpdateStackResponseWithDefaults instantiates a new StackUpdateStackResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStackUpdateStackResponseWithDefaults() *StackUpdateStackResponse {
	this := StackUpdateStackResponse{}
	return &this
}

// GetStack returns the Stack field value if set, zero value otherwise.
func (o *StackUpdateStackResponse) GetStack() StackStack {
	if o == nil || o.Stack == nil {
		var ret StackStack
		return ret
	}
	return *o.Stack
}

// GetStackOk returns a tuple with the Stack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StackUpdateStackResponse) GetStackOk() (*StackStack, bool) {
	if o == nil || o.Stack == nil {
		return nil, false
	}
	return o.Stack, true
}

// HasStack returns a boolean if a field has been set.
func (o *StackUpdateStackResponse) HasStack() bool {
	if o != nil && o.Stack != nil {
		return true
	}

	return false
}

// SetStack gets a reference to the given StackStack and assigns it to the Stack field.
func (o *StackUpdateStackResponse) SetStack(v StackStack) {
	o.Stack = &v
}

func (o StackUpdateStackResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Stack != nil {
		toSerialize["stack"] = o.Stack
	}
	return json.Marshal(toSerialize)
}

type NullableStackUpdateStackResponse struct {
	value *StackUpdateStackResponse
	isSet bool
}

func (v NullableStackUpdateStackResponse) Get() *StackUpdateStackResponse {
	return v.value
}

func (v *NullableStackUpdateStackResponse) Set(val *StackUpdateStackResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStackUpdateStackResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStackUpdateStackResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStackUpdateStackResponse(val *StackUpdateStackResponse) *NullableStackUpdateStackResponse {
	return &NullableStackUpdateStackResponse{value: val, isSet: true}
}

func (v NullableStackUpdateStackResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStackUpdateStackResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
