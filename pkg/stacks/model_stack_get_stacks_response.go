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

// StackGetStacksResponse struct for StackGetStacksResponse
type StackGetStacksResponse struct {
	PageInfo *PaginationPageInfo `json:"pageInfo,omitempty"`
	Results *[]StackStack `json:"results,omitempty"`
}

// NewStackGetStacksResponse instantiates a new StackGetStacksResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStackGetStacksResponse() *StackGetStacksResponse {
	this := StackGetStacksResponse{}
	return &this
}

// NewStackGetStacksResponseWithDefaults instantiates a new StackGetStacksResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStackGetStacksResponseWithDefaults() *StackGetStacksResponse {
	this := StackGetStacksResponse{}
	return &this
}

// GetPageInfo returns the PageInfo field value if set, zero value otherwise.
func (o *StackGetStacksResponse) GetPageInfo() PaginationPageInfo {
	if o == nil || o.PageInfo == nil {
		var ret PaginationPageInfo
		return ret
	}
	return *o.PageInfo
}

// GetPageInfoOk returns a tuple with the PageInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StackGetStacksResponse) GetPageInfoOk() (*PaginationPageInfo, bool) {
	if o == nil || o.PageInfo == nil {
		return nil, false
	}
	return o.PageInfo, true
}

// HasPageInfo returns a boolean if a field has been set.
func (o *StackGetStacksResponse) HasPageInfo() bool {
	if o != nil && o.PageInfo != nil {
		return true
	}

	return false
}

// SetPageInfo gets a reference to the given PaginationPageInfo and assigns it to the PageInfo field.
func (o *StackGetStacksResponse) SetPageInfo(v PaginationPageInfo) {
	o.PageInfo = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *StackGetStacksResponse) GetResults() []StackStack {
	if o == nil || o.Results == nil {
		var ret []StackStack
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StackGetStacksResponse) GetResultsOk() (*[]StackStack, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *StackGetStacksResponse) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []StackStack and assigns it to the Results field.
func (o *StackGetStacksResponse) SetResults(v []StackStack) {
	o.Results = &v
}

func (o StackGetStacksResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PageInfo != nil {
		toSerialize["pageInfo"] = o.PageInfo
	}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableStackGetStacksResponse struct {
	value *StackGetStacksResponse
	isSet bool
}

func (v NullableStackGetStacksResponse) Get() *StackGetStacksResponse {
	return v.value
}

func (v *NullableStackGetStacksResponse) Set(val *StackGetStacksResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStackGetStacksResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStackGetStacksResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStackGetStacksResponse(val *StackGetStacksResponse) *NullableStackGetStacksResponse {
	return &NullableStackGetStacksResponse{value: val, isSet: true}
}

func (v NullableStackGetStacksResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStackGetStacksResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}