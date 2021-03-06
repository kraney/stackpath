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

// UpdateMonitorRequestPatchHttpConfigurationDigestAuth Digest authentication for the HTTP configuration of a monitor update request.
type UpdateMonitorRequestPatchHttpConfigurationDigestAuth struct {
	// The username used for digest authentication by a monitor.
	Username *string `json:"username,omitempty"`
	// The password used for digest authentication by a monitor.
	Password *string `json:"password,omitempty"`
}

// NewUpdateMonitorRequestPatchHttpConfigurationDigestAuth instantiates a new UpdateMonitorRequestPatchHttpConfigurationDigestAuth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateMonitorRequestPatchHttpConfigurationDigestAuth() *UpdateMonitorRequestPatchHttpConfigurationDigestAuth {
	this := UpdateMonitorRequestPatchHttpConfigurationDigestAuth{}
	return &this
}

// NewUpdateMonitorRequestPatchHttpConfigurationDigestAuthWithDefaults instantiates a new UpdateMonitorRequestPatchHttpConfigurationDigestAuth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateMonitorRequestPatchHttpConfigurationDigestAuthWithDefaults() *UpdateMonitorRequestPatchHttpConfigurationDigestAuth {
	this := UpdateMonitorRequestPatchHttpConfigurationDigestAuth{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *UpdateMonitorRequestPatchHttpConfigurationDigestAuth) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMonitorRequestPatchHttpConfigurationDigestAuth) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *UpdateMonitorRequestPatchHttpConfigurationDigestAuth) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *UpdateMonitorRequestPatchHttpConfigurationDigestAuth) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UpdateMonitorRequestPatchHttpConfigurationDigestAuth) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMonitorRequestPatchHttpConfigurationDigestAuth) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UpdateMonitorRequestPatchHttpConfigurationDigestAuth) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UpdateMonitorRequestPatchHttpConfigurationDigestAuth) SetPassword(v string) {
	o.Password = &v
}

func (o UpdateMonitorRequestPatchHttpConfigurationDigestAuth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateMonitorRequestPatchHttpConfigurationDigestAuth struct {
	value *UpdateMonitorRequestPatchHttpConfigurationDigestAuth
	isSet bool
}

func (v NullableUpdateMonitorRequestPatchHttpConfigurationDigestAuth) Get() *UpdateMonitorRequestPatchHttpConfigurationDigestAuth {
	return v.value
}

func (v *NullableUpdateMonitorRequestPatchHttpConfigurationDigestAuth) Set(val *UpdateMonitorRequestPatchHttpConfigurationDigestAuth) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMonitorRequestPatchHttpConfigurationDigestAuth) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMonitorRequestPatchHttpConfigurationDigestAuth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMonitorRequestPatchHttpConfigurationDigestAuth(val *UpdateMonitorRequestPatchHttpConfigurationDigestAuth) *NullableUpdateMonitorRequestPatchHttpConfigurationDigestAuth {
	return &NullableUpdateMonitorRequestPatchHttpConfigurationDigestAuth{value: val, isSet: true}
}

func (v NullableUpdateMonitorRequestPatchHttpConfigurationDigestAuth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateMonitorRequestPatchHttpConfigurationDigestAuth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
