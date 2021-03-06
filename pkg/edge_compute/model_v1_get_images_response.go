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

// V1GetImagesResponse A response from a request to retrieve images for a stack
type V1GetImagesResponse struct {
	PageInfo *PaginationPageInfo `json:"pageInfo,omitempty"`
	// The requested images
	Results *[]V1Image `json:"results,omitempty"`
}

// NewV1GetImagesResponse instantiates a new V1GetImagesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1GetImagesResponse() *V1GetImagesResponse {
	this := V1GetImagesResponse{}
	return &this
}

// NewV1GetImagesResponseWithDefaults instantiates a new V1GetImagesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1GetImagesResponseWithDefaults() *V1GetImagesResponse {
	this := V1GetImagesResponse{}
	return &this
}

// GetPageInfo returns the PageInfo field value if set, zero value otherwise.
func (o *V1GetImagesResponse) GetPageInfo() PaginationPageInfo {
	if o == nil || o.PageInfo == nil {
		var ret PaginationPageInfo
		return ret
	}
	return *o.PageInfo
}

// GetPageInfoOk returns a tuple with the PageInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GetImagesResponse) GetPageInfoOk() (*PaginationPageInfo, bool) {
	if o == nil || o.PageInfo == nil {
		return nil, false
	}
	return o.PageInfo, true
}

// HasPageInfo returns a boolean if a field has been set.
func (o *V1GetImagesResponse) HasPageInfo() bool {
	if o != nil && o.PageInfo != nil {
		return true
	}

	return false
}

// SetPageInfo gets a reference to the given PaginationPageInfo and assigns it to the PageInfo field.
func (o *V1GetImagesResponse) SetPageInfo(v PaginationPageInfo) {
	o.PageInfo = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *V1GetImagesResponse) GetResults() []V1Image {
	if o == nil || o.Results == nil {
		var ret []V1Image
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GetImagesResponse) GetResultsOk() (*[]V1Image, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *V1GetImagesResponse) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []V1Image and assigns it to the Results field.
func (o *V1GetImagesResponse) SetResults(v []V1Image) {
	o.Results = &v
}

func (o V1GetImagesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PageInfo != nil {
		toSerialize["pageInfo"] = o.PageInfo
	}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableV1GetImagesResponse struct {
	value *V1GetImagesResponse
	isSet bool
}

func (v NullableV1GetImagesResponse) Get() *V1GetImagesResponse {
	return v.value
}

func (v *NullableV1GetImagesResponse) Set(val *V1GetImagesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableV1GetImagesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableV1GetImagesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1GetImagesResponse(val *V1GetImagesResponse) *NullableV1GetImagesResponse {
	return &NullableV1GetImagesResponse{value: val, isSet: true}
}

func (v NullableV1GetImagesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1GetImagesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
