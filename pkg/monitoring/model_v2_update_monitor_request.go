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

// V2UpdateMonitorRequest struct for V2UpdateMonitorRequest
type V2UpdateMonitorRequest struct {
	// The monitor's name
	Name *string `json:"name,omitempty"`
	Locations *UpdateMonitorRequestLocationsValue `json:"locations,omitempty"`
	// How often a monitor should be checked.
	Interval *string `json:"interval,omitempty"`
	// The length of time a monitor check should wait before timing out.
	Timeout *string `json:"timeout,omitempty"`
	IpVersion *UpdateMonitorRequestIpVersionValue `json:"ipVersion,omitempty"`
	Http *UpdateMonitorRequestPatchHttpConfiguration `json:"http,omitempty"`
	Tcp *UpdateMonitorRequestPatchTcpConfiguration `json:"tcp,omitempty"`
	// Whether a monitor is enabled or not.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewV2UpdateMonitorRequest instantiates a new V2UpdateMonitorRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2UpdateMonitorRequest() *V2UpdateMonitorRequest {
	this := V2UpdateMonitorRequest{}
	return &this
}

// NewV2UpdateMonitorRequestWithDefaults instantiates a new V2UpdateMonitorRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2UpdateMonitorRequestWithDefaults() *V2UpdateMonitorRequest {
	this := V2UpdateMonitorRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V2UpdateMonitorRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2UpdateMonitorRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V2UpdateMonitorRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V2UpdateMonitorRequest) SetName(v string) {
	o.Name = &v
}

// GetLocations returns the Locations field value if set, zero value otherwise.
func (o *V2UpdateMonitorRequest) GetLocations() UpdateMonitorRequestLocationsValue {
	if o == nil || o.Locations == nil {
		var ret UpdateMonitorRequestLocationsValue
		return ret
	}
	return *o.Locations
}

// GetLocationsOk returns a tuple with the Locations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2UpdateMonitorRequest) GetLocationsOk() (*UpdateMonitorRequestLocationsValue, bool) {
	if o == nil || o.Locations == nil {
		return nil, false
	}
	return o.Locations, true
}

// HasLocations returns a boolean if a field has been set.
func (o *V2UpdateMonitorRequest) HasLocations() bool {
	if o != nil && o.Locations != nil {
		return true
	}

	return false
}

// SetLocations gets a reference to the given UpdateMonitorRequestLocationsValue and assigns it to the Locations field.
func (o *V2UpdateMonitorRequest) SetLocations(v UpdateMonitorRequestLocationsValue) {
	o.Locations = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *V2UpdateMonitorRequest) GetInterval() string {
	if o == nil || o.Interval == nil {
		var ret string
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2UpdateMonitorRequest) GetIntervalOk() (*string, bool) {
	if o == nil || o.Interval == nil {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *V2UpdateMonitorRequest) HasInterval() bool {
	if o != nil && o.Interval != nil {
		return true
	}

	return false
}

// SetInterval gets a reference to the given string and assigns it to the Interval field.
func (o *V2UpdateMonitorRequest) SetInterval(v string) {
	o.Interval = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *V2UpdateMonitorRequest) GetTimeout() string {
	if o == nil || o.Timeout == nil {
		var ret string
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2UpdateMonitorRequest) GetTimeoutOk() (*string, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *V2UpdateMonitorRequest) HasTimeout() bool {
	if o != nil && o.Timeout != nil {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given string and assigns it to the Timeout field.
func (o *V2UpdateMonitorRequest) SetTimeout(v string) {
	o.Timeout = &v
}

// GetIpVersion returns the IpVersion field value if set, zero value otherwise.
func (o *V2UpdateMonitorRequest) GetIpVersion() UpdateMonitorRequestIpVersionValue {
	if o == nil || o.IpVersion == nil {
		var ret UpdateMonitorRequestIpVersionValue
		return ret
	}
	return *o.IpVersion
}

// GetIpVersionOk returns a tuple with the IpVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2UpdateMonitorRequest) GetIpVersionOk() (*UpdateMonitorRequestIpVersionValue, bool) {
	if o == nil || o.IpVersion == nil {
		return nil, false
	}
	return o.IpVersion, true
}

// HasIpVersion returns a boolean if a field has been set.
func (o *V2UpdateMonitorRequest) HasIpVersion() bool {
	if o != nil && o.IpVersion != nil {
		return true
	}

	return false
}

// SetIpVersion gets a reference to the given UpdateMonitorRequestIpVersionValue and assigns it to the IpVersion field.
func (o *V2UpdateMonitorRequest) SetIpVersion(v UpdateMonitorRequestIpVersionValue) {
	o.IpVersion = &v
}

// GetHttp returns the Http field value if set, zero value otherwise.
func (o *V2UpdateMonitorRequest) GetHttp() UpdateMonitorRequestPatchHttpConfiguration {
	if o == nil || o.Http == nil {
		var ret UpdateMonitorRequestPatchHttpConfiguration
		return ret
	}
	return *o.Http
}

// GetHttpOk returns a tuple with the Http field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2UpdateMonitorRequest) GetHttpOk() (*UpdateMonitorRequestPatchHttpConfiguration, bool) {
	if o == nil || o.Http == nil {
		return nil, false
	}
	return o.Http, true
}

// HasHttp returns a boolean if a field has been set.
func (o *V2UpdateMonitorRequest) HasHttp() bool {
	if o != nil && o.Http != nil {
		return true
	}

	return false
}

// SetHttp gets a reference to the given UpdateMonitorRequestPatchHttpConfiguration and assigns it to the Http field.
func (o *V2UpdateMonitorRequest) SetHttp(v UpdateMonitorRequestPatchHttpConfiguration) {
	o.Http = &v
}

// GetTcp returns the Tcp field value if set, zero value otherwise.
func (o *V2UpdateMonitorRequest) GetTcp() UpdateMonitorRequestPatchTcpConfiguration {
	if o == nil || o.Tcp == nil {
		var ret UpdateMonitorRequestPatchTcpConfiguration
		return ret
	}
	return *o.Tcp
}

// GetTcpOk returns a tuple with the Tcp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2UpdateMonitorRequest) GetTcpOk() (*UpdateMonitorRequestPatchTcpConfiguration, bool) {
	if o == nil || o.Tcp == nil {
		return nil, false
	}
	return o.Tcp, true
}

// HasTcp returns a boolean if a field has been set.
func (o *V2UpdateMonitorRequest) HasTcp() bool {
	if o != nil && o.Tcp != nil {
		return true
	}

	return false
}

// SetTcp gets a reference to the given UpdateMonitorRequestPatchTcpConfiguration and assigns it to the Tcp field.
func (o *V2UpdateMonitorRequest) SetTcp(v UpdateMonitorRequestPatchTcpConfiguration) {
	o.Tcp = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *V2UpdateMonitorRequest) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2UpdateMonitorRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *V2UpdateMonitorRequest) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *V2UpdateMonitorRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o V2UpdateMonitorRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Locations != nil {
		toSerialize["locations"] = o.Locations
	}
	if o.Interval != nil {
		toSerialize["interval"] = o.Interval
	}
	if o.Timeout != nil {
		toSerialize["timeout"] = o.Timeout
	}
	if o.IpVersion != nil {
		toSerialize["ipVersion"] = o.IpVersion
	}
	if o.Http != nil {
		toSerialize["http"] = o.Http
	}
	if o.Tcp != nil {
		toSerialize["tcp"] = o.Tcp
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableV2UpdateMonitorRequest struct {
	value *V2UpdateMonitorRequest
	isSet bool
}

func (v NullableV2UpdateMonitorRequest) Get() *V2UpdateMonitorRequest {
	return v.value
}

func (v *NullableV2UpdateMonitorRequest) Set(val *V2UpdateMonitorRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV2UpdateMonitorRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV2UpdateMonitorRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2UpdateMonitorRequest(val *V2UpdateMonitorRequest) *NullableV2UpdateMonitorRequest {
	return &NullableV2UpdateMonitorRequest{value: val, isSet: true}
}

func (v NullableV2UpdateMonitorRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2UpdateMonitorRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}