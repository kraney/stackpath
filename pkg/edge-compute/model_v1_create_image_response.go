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

// V1CreateImageResponse A response from a request to create an image
type V1CreateImageResponse struct {
	Image *V1Image `json:"image,omitempty"`
}

// NewV1CreateImageResponse instantiates a new V1CreateImageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1CreateImageResponse() *V1CreateImageResponse {
	this := V1CreateImageResponse{}
	return &this
}

// NewV1CreateImageResponseWithDefaults instantiates a new V1CreateImageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1CreateImageResponseWithDefaults() *V1CreateImageResponse {
	this := V1CreateImageResponse{}
	return &this
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *V1CreateImageResponse) GetImage() V1Image {
	if o == nil || o.Image == nil {
		var ret V1Image
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1CreateImageResponse) GetImageOk() (*V1Image, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *V1CreateImageResponse) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given V1Image and assigns it to the Image field.
func (o *V1CreateImageResponse) SetImage(v V1Image) {
	o.Image = &v
}

func (o V1CreateImageResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Image != nil {
		toSerialize["image"] = o.Image
	}
	return json.Marshal(toSerialize)
}

type NullableV1CreateImageResponse struct {
	value *V1CreateImageResponse
	isSet bool
}

func (v NullableV1CreateImageResponse) Get() *V1CreateImageResponse {
	return v.value
}

func (v *NullableV1CreateImageResponse) Set(val *V1CreateImageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableV1CreateImageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableV1CreateImageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1CreateImageResponse(val *V1CreateImageResponse) *NullableV1CreateImageResponse {
	return &NullableV1CreateImageResponse{value: val, isSet: true}
}

func (v NullableV1CreateImageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1CreateImageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}