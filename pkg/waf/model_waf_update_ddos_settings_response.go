/*
 * Web Application Firewall
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package waf

import (
	"encoding/json"
)

// WafUpdateDdosSettingsResponse A response from a request to update a WAF site's DDOS protection settings
type WafUpdateDdosSettingsResponse struct {
	DdosSettings *WafDdosSettings `json:"ddosSettings,omitempty"`
}

// NewWafUpdateDdosSettingsResponse instantiates a new WafUpdateDdosSettingsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafUpdateDdosSettingsResponse() *WafUpdateDdosSettingsResponse {
	this := WafUpdateDdosSettingsResponse{}
	return &this
}

// NewWafUpdateDdosSettingsResponseWithDefaults instantiates a new WafUpdateDdosSettingsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafUpdateDdosSettingsResponseWithDefaults() *WafUpdateDdosSettingsResponse {
	this := WafUpdateDdosSettingsResponse{}
	return &this
}

// GetDdosSettings returns the DdosSettings field value if set, zero value otherwise.
func (o *WafUpdateDdosSettingsResponse) GetDdosSettings() WafDdosSettings {
	if o == nil || o.DdosSettings == nil {
		var ret WafDdosSettings
		return ret
	}
	return *o.DdosSettings
}

// GetDdosSettingsOk returns a tuple with the DdosSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafUpdateDdosSettingsResponse) GetDdosSettingsOk() (*WafDdosSettings, bool) {
	if o == nil || o.DdosSettings == nil {
		return nil, false
	}
	return o.DdosSettings, true
}

// HasDdosSettings returns a boolean if a field has been set.
func (o *WafUpdateDdosSettingsResponse) HasDdosSettings() bool {
	if o != nil && o.DdosSettings != nil {
		return true
	}

	return false
}

// SetDdosSettings gets a reference to the given WafDdosSettings and assigns it to the DdosSettings field.
func (o *WafUpdateDdosSettingsResponse) SetDdosSettings(v WafDdosSettings) {
	o.DdosSettings = &v
}

func (o WafUpdateDdosSettingsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DdosSettings != nil {
		toSerialize["ddosSettings"] = o.DdosSettings
	}
	return json.Marshal(toSerialize)
}

type NullableWafUpdateDdosSettingsResponse struct {
	value *WafUpdateDdosSettingsResponse
	isSet bool
}

func (v NullableWafUpdateDdosSettingsResponse) Get() *WafUpdateDdosSettingsResponse {
	return v.value
}

func (v *NullableWafUpdateDdosSettingsResponse) Set(val *WafUpdateDdosSettingsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWafUpdateDdosSettingsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWafUpdateDdosSettingsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWafUpdateDdosSettingsResponse(val *WafUpdateDdosSettingsResponse) *NullableWafUpdateDdosSettingsResponse {
	return &NullableWafUpdateDdosSettingsResponse{value: val, isSet: true}
}

func (v NullableWafUpdateDdosSettingsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWafUpdateDdosSettingsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
