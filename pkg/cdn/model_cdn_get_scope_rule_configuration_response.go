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

// CdnGetScopeRuleConfigurationResponse The response from a request to retrieve an EdgeRule's configuration
type CdnGetScopeRuleConfigurationResponse struct {
	Configuration *CustconfConfiguration `json:"configuration,omitempty"`
}

// NewCdnGetScopeRuleConfigurationResponse instantiates a new CdnGetScopeRuleConfigurationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCdnGetScopeRuleConfigurationResponse() *CdnGetScopeRuleConfigurationResponse {
	this := CdnGetScopeRuleConfigurationResponse{}
	return &this
}

// NewCdnGetScopeRuleConfigurationResponseWithDefaults instantiates a new CdnGetScopeRuleConfigurationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCdnGetScopeRuleConfigurationResponseWithDefaults() *CdnGetScopeRuleConfigurationResponse {
	this := CdnGetScopeRuleConfigurationResponse{}
	return &this
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *CdnGetScopeRuleConfigurationResponse) GetConfiguration() CustconfConfiguration {
	if o == nil || o.Configuration == nil {
		var ret CustconfConfiguration
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdnGetScopeRuleConfigurationResponse) GetConfigurationOk() (*CustconfConfiguration, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *CdnGetScopeRuleConfigurationResponse) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given CustconfConfiguration and assigns it to the Configuration field.
func (o *CdnGetScopeRuleConfigurationResponse) SetConfiguration(v CustconfConfiguration) {
	o.Configuration = &v
}

func (o CdnGetScopeRuleConfigurationResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Configuration != nil {
		toSerialize["configuration"] = o.Configuration
	}
	return json.Marshal(toSerialize)
}

type NullableCdnGetScopeRuleConfigurationResponse struct {
	value *CdnGetScopeRuleConfigurationResponse
	isSet bool
}

func (v NullableCdnGetScopeRuleConfigurationResponse) Get() *CdnGetScopeRuleConfigurationResponse {
	return v.value
}

func (v *NullableCdnGetScopeRuleConfigurationResponse) Set(val *CdnGetScopeRuleConfigurationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCdnGetScopeRuleConfigurationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCdnGetScopeRuleConfigurationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCdnGetScopeRuleConfigurationResponse(val *CdnGetScopeRuleConfigurationResponse) *NullableCdnGetScopeRuleConfigurationResponse {
	return &NullableCdnGetScopeRuleConfigurationResponse{value: val, isSet: true}
}

func (v NullableCdnGetScopeRuleConfigurationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCdnGetScopeRuleConfigurationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}