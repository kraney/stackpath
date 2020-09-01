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

// V2ReplaceMonitorRequest struct for V2ReplaceMonitorRequest
type V2ReplaceMonitorRequest struct {
	// The name of the monitor.
	Name *string `json:"name,omitempty"`
	// The list of locations a monitor is checked from.
	Locations *[]string `json:"locations,omitempty"`
	// How often a monitor is checked.
	Interval *string `json:"interval,omitempty"`
	// The amount of time to wait before a monitor check times out.
	Timeout *string `json:"timeout,omitempty"`
	IpVersion *V2IpVersion `json:"ipVersion,omitempty"`
	Http *V2HttpConfiguration `json:"http,omitempty"`
	Tcp *V2TcpConfiguration `json:"tcp,omitempty"`
	// Whether a monitor is enabled or not.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewV2ReplaceMonitorRequest instantiates a new V2ReplaceMonitorRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2ReplaceMonitorRequest() *V2ReplaceMonitorRequest {
	this := V2ReplaceMonitorRequest{}
	var ipVersion V2IpVersion = "IPV4"
	this.IpVersion = &ipVersion
	return &this
}

// NewV2ReplaceMonitorRequestWithDefaults instantiates a new V2ReplaceMonitorRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2ReplaceMonitorRequestWithDefaults() *V2ReplaceMonitorRequest {
	this := V2ReplaceMonitorRequest{}
	var ipVersion V2IpVersion = "IPV4"
	this.IpVersion = &ipVersion
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V2ReplaceMonitorRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ReplaceMonitorRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V2ReplaceMonitorRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V2ReplaceMonitorRequest) SetName(v string) {
	o.Name = &v
}

// GetLocations returns the Locations field value if set, zero value otherwise.
func (o *V2ReplaceMonitorRequest) GetLocations() []string {
	if o == nil || o.Locations == nil {
		var ret []string
		return ret
	}
	return *o.Locations
}

// GetLocationsOk returns a tuple with the Locations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ReplaceMonitorRequest) GetLocationsOk() (*[]string, bool) {
	if o == nil || o.Locations == nil {
		return nil, false
	}
	return o.Locations, true
}

// HasLocations returns a boolean if a field has been set.
func (o *V2ReplaceMonitorRequest) HasLocations() bool {
	if o != nil && o.Locations != nil {
		return true
	}

	return false
}

// SetLocations gets a reference to the given []string and assigns it to the Locations field.
func (o *V2ReplaceMonitorRequest) SetLocations(v []string) {
	o.Locations = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *V2ReplaceMonitorRequest) GetInterval() string {
	if o == nil || o.Interval == nil {
		var ret string
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ReplaceMonitorRequest) GetIntervalOk() (*string, bool) {
	if o == nil || o.Interval == nil {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *V2ReplaceMonitorRequest) HasInterval() bool {
	if o != nil && o.Interval != nil {
		return true
	}

	return false
}

// SetInterval gets a reference to the given string and assigns it to the Interval field.
func (o *V2ReplaceMonitorRequest) SetInterval(v string) {
	o.Interval = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *V2ReplaceMonitorRequest) GetTimeout() string {
	if o == nil || o.Timeout == nil {
		var ret string
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ReplaceMonitorRequest) GetTimeoutOk() (*string, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *V2ReplaceMonitorRequest) HasTimeout() bool {
	if o != nil && o.Timeout != nil {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given string and assigns it to the Timeout field.
func (o *V2ReplaceMonitorRequest) SetTimeout(v string) {
	o.Timeout = &v
}

// GetIpVersion returns the IpVersion field value if set, zero value otherwise.
func (o *V2ReplaceMonitorRequest) GetIpVersion() V2IpVersion {
	if o == nil || o.IpVersion == nil {
		var ret V2IpVersion
		return ret
	}
	return *o.IpVersion
}

// GetIpVersionOk returns a tuple with the IpVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ReplaceMonitorRequest) GetIpVersionOk() (*V2IpVersion, bool) {
	if o == nil || o.IpVersion == nil {
		return nil, false
	}
	return o.IpVersion, true
}

// HasIpVersion returns a boolean if a field has been set.
func (o *V2ReplaceMonitorRequest) HasIpVersion() bool {
	if o != nil && o.IpVersion != nil {
		return true
	}

	return false
}

// SetIpVersion gets a reference to the given V2IpVersion and assigns it to the IpVersion field.
func (o *V2ReplaceMonitorRequest) SetIpVersion(v V2IpVersion) {
	o.IpVersion = &v
}

// GetHttp returns the Http field value if set, zero value otherwise.
func (o *V2ReplaceMonitorRequest) GetHttp() V2HttpConfiguration {
	if o == nil || o.Http == nil {
		var ret V2HttpConfiguration
		return ret
	}
	return *o.Http
}

// GetHttpOk returns a tuple with the Http field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ReplaceMonitorRequest) GetHttpOk() (*V2HttpConfiguration, bool) {
	if o == nil || o.Http == nil {
		return nil, false
	}
	return o.Http, true
}

// HasHttp returns a boolean if a field has been set.
func (o *V2ReplaceMonitorRequest) HasHttp() bool {
	if o != nil && o.Http != nil {
		return true
	}

	return false
}

// SetHttp gets a reference to the given V2HttpConfiguration and assigns it to the Http field.
func (o *V2ReplaceMonitorRequest) SetHttp(v V2HttpConfiguration) {
	o.Http = &v
}

// GetTcp returns the Tcp field value if set, zero value otherwise.
func (o *V2ReplaceMonitorRequest) GetTcp() V2TcpConfiguration {
	if o == nil || o.Tcp == nil {
		var ret V2TcpConfiguration
		return ret
	}
	return *o.Tcp
}

// GetTcpOk returns a tuple with the Tcp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ReplaceMonitorRequest) GetTcpOk() (*V2TcpConfiguration, bool) {
	if o == nil || o.Tcp == nil {
		return nil, false
	}
	return o.Tcp, true
}

// HasTcp returns a boolean if a field has been set.
func (o *V2ReplaceMonitorRequest) HasTcp() bool {
	if o != nil && o.Tcp != nil {
		return true
	}

	return false
}

// SetTcp gets a reference to the given V2TcpConfiguration and assigns it to the Tcp field.
func (o *V2ReplaceMonitorRequest) SetTcp(v V2TcpConfiguration) {
	o.Tcp = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *V2ReplaceMonitorRequest) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ReplaceMonitorRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *V2ReplaceMonitorRequest) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *V2ReplaceMonitorRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o V2ReplaceMonitorRequest) MarshalJSON() ([]byte, error) {
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

type NullableV2ReplaceMonitorRequest struct {
	value *V2ReplaceMonitorRequest
	isSet bool
}

func (v NullableV2ReplaceMonitorRequest) Get() *V2ReplaceMonitorRequest {
	return v.value
}

func (v *NullableV2ReplaceMonitorRequest) Set(val *V2ReplaceMonitorRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV2ReplaceMonitorRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV2ReplaceMonitorRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2ReplaceMonitorRequest(val *V2ReplaceMonitorRequest) *NullableV2ReplaceMonitorRequest {
	return &NullableV2ReplaceMonitorRequest{value: val, isSet: true}
}

func (v NullableV2ReplaceMonitorRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2ReplaceMonitorRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}