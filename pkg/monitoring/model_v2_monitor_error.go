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
	"time"
)

// V2MonitorError An error that ocurred during a monitoring check.
type V2MonitorError struct {
	// The monitor identifier associated with the error.
	MonitorId *string `json:"monitorId,omitempty"`
	// A list of locations that have the error.
	Locations *[]Monitoringv2Location `json:"locations,omitempty"`
	// The error text for a monitor error.
	Error *string `json:"error,omitempty"`
	// The date a monitor error was created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
}

// NewV2MonitorError instantiates a new V2MonitorError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2MonitorError() *V2MonitorError {
	this := V2MonitorError{}
	return &this
}

// NewV2MonitorErrorWithDefaults instantiates a new V2MonitorError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2MonitorErrorWithDefaults() *V2MonitorError {
	this := V2MonitorError{}
	return &this
}

// GetMonitorId returns the MonitorId field value if set, zero value otherwise.
func (o *V2MonitorError) GetMonitorId() string {
	if o == nil || o.MonitorId == nil {
		var ret string
		return ret
	}
	return *o.MonitorId
}

// GetMonitorIdOk returns a tuple with the MonitorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2MonitorError) GetMonitorIdOk() (*string, bool) {
	if o == nil || o.MonitorId == nil {
		return nil, false
	}
	return o.MonitorId, true
}

// HasMonitorId returns a boolean if a field has been set.
func (o *V2MonitorError) HasMonitorId() bool {
	if o != nil && o.MonitorId != nil {
		return true
	}

	return false
}

// SetMonitorId gets a reference to the given string and assigns it to the MonitorId field.
func (o *V2MonitorError) SetMonitorId(v string) {
	o.MonitorId = &v
}

// GetLocations returns the Locations field value if set, zero value otherwise.
func (o *V2MonitorError) GetLocations() []Monitoringv2Location {
	if o == nil || o.Locations == nil {
		var ret []Monitoringv2Location
		return ret
	}
	return *o.Locations
}

// GetLocationsOk returns a tuple with the Locations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2MonitorError) GetLocationsOk() (*[]Monitoringv2Location, bool) {
	if o == nil || o.Locations == nil {
		return nil, false
	}
	return o.Locations, true
}

// HasLocations returns a boolean if a field has been set.
func (o *V2MonitorError) HasLocations() bool {
	if o != nil && o.Locations != nil {
		return true
	}

	return false
}

// SetLocations gets a reference to the given []Monitoringv2Location and assigns it to the Locations field.
func (o *V2MonitorError) SetLocations(v []Monitoringv2Location) {
	o.Locations = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *V2MonitorError) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2MonitorError) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *V2MonitorError) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *V2MonitorError) SetError(v string) {
	o.Error = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *V2MonitorError) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2MonitorError) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *V2MonitorError) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *V2MonitorError) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

func (o V2MonitorError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MonitorId != nil {
		toSerialize["monitorId"] = o.MonitorId
	}
	if o.Locations != nil {
		toSerialize["locations"] = o.Locations
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableV2MonitorError struct {
	value *V2MonitorError
	isSet bool
}

func (v NullableV2MonitorError) Get() *V2MonitorError {
	return v.value
}

func (v *NullableV2MonitorError) Set(val *V2MonitorError) {
	v.value = val
	v.isSet = true
}

func (v NullableV2MonitorError) IsSet() bool {
	return v.isSet
}

func (v *NullableV2MonitorError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2MonitorError(val *V2MonitorError) *NullableV2MonitorError {
	return &NullableV2MonitorError{value: val, isSet: true}
}

func (v NullableV2MonitorError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2MonitorError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}