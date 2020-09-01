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

// V2LocationWithAnalytics A location with additional metrics  The additional metrics provide status and response times about recent monitor checks for a location.
type V2LocationWithAnalytics struct {
	Location *Monitoringv2Location `json:"location,omitempty"`
	// The HTTP status code of the most recent monitoring check from the location.
	StatusCode *int32 `json:"statusCode,omitempty"`
	// The response time of the most recent monitoring check from the location.
	ResponseTime *string `json:"responseTime,omitempty"`
}

// NewV2LocationWithAnalytics instantiates a new V2LocationWithAnalytics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2LocationWithAnalytics() *V2LocationWithAnalytics {
	this := V2LocationWithAnalytics{}
	return &this
}

// NewV2LocationWithAnalyticsWithDefaults instantiates a new V2LocationWithAnalytics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2LocationWithAnalyticsWithDefaults() *V2LocationWithAnalytics {
	this := V2LocationWithAnalytics{}
	return &this
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *V2LocationWithAnalytics) GetLocation() Monitoringv2Location {
	if o == nil || o.Location == nil {
		var ret Monitoringv2Location
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2LocationWithAnalytics) GetLocationOk() (*Monitoringv2Location, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *V2LocationWithAnalytics) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given Monitoringv2Location and assigns it to the Location field.
func (o *V2LocationWithAnalytics) SetLocation(v Monitoringv2Location) {
	o.Location = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *V2LocationWithAnalytics) GetStatusCode() int32 {
	if o == nil || o.StatusCode == nil {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2LocationWithAnalytics) GetStatusCodeOk() (*int32, bool) {
	if o == nil || o.StatusCode == nil {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *V2LocationWithAnalytics) HasStatusCode() bool {
	if o != nil && o.StatusCode != nil {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *V2LocationWithAnalytics) SetStatusCode(v int32) {
	o.StatusCode = &v
}

// GetResponseTime returns the ResponseTime field value if set, zero value otherwise.
func (o *V2LocationWithAnalytics) GetResponseTime() string {
	if o == nil || o.ResponseTime == nil {
		var ret string
		return ret
	}
	return *o.ResponseTime
}

// GetResponseTimeOk returns a tuple with the ResponseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2LocationWithAnalytics) GetResponseTimeOk() (*string, bool) {
	if o == nil || o.ResponseTime == nil {
		return nil, false
	}
	return o.ResponseTime, true
}

// HasResponseTime returns a boolean if a field has been set.
func (o *V2LocationWithAnalytics) HasResponseTime() bool {
	if o != nil && o.ResponseTime != nil {
		return true
	}

	return false
}

// SetResponseTime gets a reference to the given string and assigns it to the ResponseTime field.
func (o *V2LocationWithAnalytics) SetResponseTime(v string) {
	o.ResponseTime = &v
}

func (o V2LocationWithAnalytics) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if o.StatusCode != nil {
		toSerialize["statusCode"] = o.StatusCode
	}
	if o.ResponseTime != nil {
		toSerialize["responseTime"] = o.ResponseTime
	}
	return json.Marshal(toSerialize)
}

type NullableV2LocationWithAnalytics struct {
	value *V2LocationWithAnalytics
	isSet bool
}

func (v NullableV2LocationWithAnalytics) Get() *V2LocationWithAnalytics {
	return v.value
}

func (v *NullableV2LocationWithAnalytics) Set(val *V2LocationWithAnalytics) {
	v.value = val
	v.isSet = true
}

func (v NullableV2LocationWithAnalytics) IsSet() bool {
	return v.isSet
}

func (v *NullableV2LocationWithAnalytics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2LocationWithAnalytics(val *V2LocationWithAnalytics) *NullableV2LocationWithAnalytics {
	return &NullableV2LocationWithAnalytics{value: val, isSet: true}
}

func (v NullableV2LocationWithAnalytics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2LocationWithAnalytics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}