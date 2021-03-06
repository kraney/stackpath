/*
 * Content Delivery Network
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package cdn

import (
	"encoding/json"
)

// CdnUpdateScopeRuleConfigurationRequest struct for CdnUpdateScopeRuleConfigurationRequest
type CdnUpdateScopeRuleConfigurationRequest struct {
	Configuration *CustconfConfiguration `json:"configuration,omitempty"`
}

// NewCdnUpdateScopeRuleConfigurationRequest instantiates a new CdnUpdateScopeRuleConfigurationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCdnUpdateScopeRuleConfigurationRequest() *CdnUpdateScopeRuleConfigurationRequest {
	this := CdnUpdateScopeRuleConfigurationRequest{}
	return &this
}

// NewCdnUpdateScopeRuleConfigurationRequestWithDefaults instantiates a new CdnUpdateScopeRuleConfigurationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCdnUpdateScopeRuleConfigurationRequestWithDefaults() *CdnUpdateScopeRuleConfigurationRequest {
	this := CdnUpdateScopeRuleConfigurationRequest{}
	return &this
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *CdnUpdateScopeRuleConfigurationRequest) GetConfiguration() CustconfConfiguration {
	if o == nil || o.Configuration == nil {
		var ret CustconfConfiguration
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdnUpdateScopeRuleConfigurationRequest) GetConfigurationOk() (*CustconfConfiguration, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *CdnUpdateScopeRuleConfigurationRequest) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given CustconfConfiguration and assigns it to the Configuration field.
func (o *CdnUpdateScopeRuleConfigurationRequest) SetConfiguration(v CustconfConfiguration) {
	o.Configuration = &v
}

func (o CdnUpdateScopeRuleConfigurationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Configuration != nil {
		toSerialize["configuration"] = o.Configuration
	}
	return json.Marshal(toSerialize)
}

type NullableCdnUpdateScopeRuleConfigurationRequest struct {
	value *CdnUpdateScopeRuleConfigurationRequest
	isSet bool
}

func (v NullableCdnUpdateScopeRuleConfigurationRequest) Get() *CdnUpdateScopeRuleConfigurationRequest {
	return v.value
}

func (v *NullableCdnUpdateScopeRuleConfigurationRequest) Set(val *CdnUpdateScopeRuleConfigurationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCdnUpdateScopeRuleConfigurationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCdnUpdateScopeRuleConfigurationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCdnUpdateScopeRuleConfigurationRequest(val *CdnUpdateScopeRuleConfigurationRequest) *NullableCdnUpdateScopeRuleConfigurationRequest {
	return &NullableCdnUpdateScopeRuleConfigurationRequest{value: val, isSet: true}
}

func (v NullableCdnUpdateScopeRuleConfigurationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCdnUpdateScopeRuleConfigurationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
