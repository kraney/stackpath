/*
 * Monitoring
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package monitoring

import (
	"encoding/json"
)

// V2GetLocationsResponse A response from a request to retrieve locations.
type V2GetLocationsResponse struct {
	// The list of monitoring locations for a stack.
	Results *[]Monitoringv2Location `json:"results,omitempty"`
}

// NewV2GetLocationsResponse instantiates a new V2GetLocationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2GetLocationsResponse() *V2GetLocationsResponse {
	this := V2GetLocationsResponse{}
	return &this
}

// NewV2GetLocationsResponseWithDefaults instantiates a new V2GetLocationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2GetLocationsResponseWithDefaults() *V2GetLocationsResponse {
	this := V2GetLocationsResponse{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *V2GetLocationsResponse) GetResults() []Monitoringv2Location {
	if o == nil || o.Results == nil {
		var ret []Monitoringv2Location
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2GetLocationsResponse) GetResultsOk() (*[]Monitoringv2Location, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *V2GetLocationsResponse) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []Monitoringv2Location and assigns it to the Results field.
func (o *V2GetLocationsResponse) SetResults(v []Monitoringv2Location) {
	o.Results = &v
}

func (o V2GetLocationsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableV2GetLocationsResponse struct {
	value *V2GetLocationsResponse
	isSet bool
}

func (v NullableV2GetLocationsResponse) Get() *V2GetLocationsResponse {
	return v.value
}

func (v *NullableV2GetLocationsResponse) Set(val *V2GetLocationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableV2GetLocationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableV2GetLocationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2GetLocationsResponse(val *V2GetLocationsResponse) *NullableV2GetLocationsResponse {
	return &NullableV2GetLocationsResponse{value: val, isSet: true}
}

func (v NullableV2GetLocationsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2GetLocationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
