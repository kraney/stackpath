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

// V2GetMonitorsResponse A response from a request to retrieve a stack's monitors.
type V2GetMonitorsResponse struct {
	PageInfo *PaginationPageInfo `json:"pageInfo,omitempty"`
	// The requested monitors.
	Results *[]V2Monitor `json:"results,omitempty"`
}

// NewV2GetMonitorsResponse instantiates a new V2GetMonitorsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2GetMonitorsResponse() *V2GetMonitorsResponse {
	this := V2GetMonitorsResponse{}
	return &this
}

// NewV2GetMonitorsResponseWithDefaults instantiates a new V2GetMonitorsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2GetMonitorsResponseWithDefaults() *V2GetMonitorsResponse {
	this := V2GetMonitorsResponse{}
	return &this
}

// GetPageInfo returns the PageInfo field value if set, zero value otherwise.
func (o *V2GetMonitorsResponse) GetPageInfo() PaginationPageInfo {
	if o == nil || o.PageInfo == nil {
		var ret PaginationPageInfo
		return ret
	}
	return *o.PageInfo
}

// GetPageInfoOk returns a tuple with the PageInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2GetMonitorsResponse) GetPageInfoOk() (*PaginationPageInfo, bool) {
	if o == nil || o.PageInfo == nil {
		return nil, false
	}
	return o.PageInfo, true
}

// HasPageInfo returns a boolean if a field has been set.
func (o *V2GetMonitorsResponse) HasPageInfo() bool {
	if o != nil && o.PageInfo != nil {
		return true
	}

	return false
}

// SetPageInfo gets a reference to the given PaginationPageInfo and assigns it to the PageInfo field.
func (o *V2GetMonitorsResponse) SetPageInfo(v PaginationPageInfo) {
	o.PageInfo = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *V2GetMonitorsResponse) GetResults() []V2Monitor {
	if o == nil || o.Results == nil {
		var ret []V2Monitor
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2GetMonitorsResponse) GetResultsOk() (*[]V2Monitor, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *V2GetMonitorsResponse) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []V2Monitor and assigns it to the Results field.
func (o *V2GetMonitorsResponse) SetResults(v []V2Monitor) {
	o.Results = &v
}

func (o V2GetMonitorsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PageInfo != nil {
		toSerialize["pageInfo"] = o.PageInfo
	}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableV2GetMonitorsResponse struct {
	value *V2GetMonitorsResponse
	isSet bool
}

func (v NullableV2GetMonitorsResponse) Get() *V2GetMonitorsResponse {
	return v.value
}

func (v *NullableV2GetMonitorsResponse) Set(val *V2GetMonitorsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableV2GetMonitorsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableV2GetMonitorsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2GetMonitorsResponse(val *V2GetMonitorsResponse) *NullableV2GetMonitorsResponse {
	return &NullableV2GetMonitorsResponse{value: val, isSet: true}
}

func (v NullableV2GetMonitorsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2GetMonitorsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}