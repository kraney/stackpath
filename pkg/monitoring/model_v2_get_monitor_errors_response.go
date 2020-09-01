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

// V2GetMonitorErrorsResponse A response from a request to retrieve errors for a monitor.
type V2GetMonitorErrorsResponse struct {
	// The requested monitor errors.
	Results *[]V2MonitorError `json:"results,omitempty"`
}

// NewV2GetMonitorErrorsResponse instantiates a new V2GetMonitorErrorsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2GetMonitorErrorsResponse() *V2GetMonitorErrorsResponse {
	this := V2GetMonitorErrorsResponse{}
	return &this
}

// NewV2GetMonitorErrorsResponseWithDefaults instantiates a new V2GetMonitorErrorsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2GetMonitorErrorsResponseWithDefaults() *V2GetMonitorErrorsResponse {
	this := V2GetMonitorErrorsResponse{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *V2GetMonitorErrorsResponse) GetResults() []V2MonitorError {
	if o == nil || o.Results == nil {
		var ret []V2MonitorError
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2GetMonitorErrorsResponse) GetResultsOk() (*[]V2MonitorError, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *V2GetMonitorErrorsResponse) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []V2MonitorError and assigns it to the Results field.
func (o *V2GetMonitorErrorsResponse) SetResults(v []V2MonitorError) {
	o.Results = &v
}

func (o V2GetMonitorErrorsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableV2GetMonitorErrorsResponse struct {
	value *V2GetMonitorErrorsResponse
	isSet bool
}

func (v NullableV2GetMonitorErrorsResponse) Get() *V2GetMonitorErrorsResponse {
	return v.value
}

func (v *NullableV2GetMonitorErrorsResponse) Set(val *V2GetMonitorErrorsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableV2GetMonitorErrorsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableV2GetMonitorErrorsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2GetMonitorErrorsResponse(val *V2GetMonitorErrorsResponse) *NullableV2GetMonitorErrorsResponse {
	return &NullableV2GetMonitorErrorsResponse{value: val, isSet: true}
}

func (v NullableV2GetMonitorErrorsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2GetMonitorErrorsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}