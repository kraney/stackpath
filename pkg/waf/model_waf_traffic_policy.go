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

// WafTrafficPolicy Aspects of WAF traffic related to policies
type WafTrafficPolicy struct {
	// The number of requests either handled or blocked by anti_automation and behavioral templates
	AntiautomationAndBehavioral *string `json:"antiautomationAndBehavioral,omitempty"`
	// The number of requests either handled or blocked by the CSRF template
	Csrf *string `json:"csrf,omitempty"`
	// The number of requests either handled or blocked by the IP reputation template
	IpReputation *string `json:"ipReputation,omitempty"`
	// The number of requests either handled or blocked by the spam and abuse template
	SpamAndAbuse *string `json:"spamAndAbuse,omitempty"`
	// The number of requests either handled or blocked by REM templates
	Waf *string `json:"waf,omitempty"`
}

// NewWafTrafficPolicy instantiates a new WafTrafficPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafTrafficPolicy() *WafTrafficPolicy {
	this := WafTrafficPolicy{}
	return &this
}

// NewWafTrafficPolicyWithDefaults instantiates a new WafTrafficPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafTrafficPolicyWithDefaults() *WafTrafficPolicy {
	this := WafTrafficPolicy{}
	return &this
}

// GetAntiautomationAndBehavioral returns the AntiautomationAndBehavioral field value if set, zero value otherwise.
func (o *WafTrafficPolicy) GetAntiautomationAndBehavioral() string {
	if o == nil || o.AntiautomationAndBehavioral == nil {
		var ret string
		return ret
	}
	return *o.AntiautomationAndBehavioral
}

// GetAntiautomationAndBehavioralOk returns a tuple with the AntiautomationAndBehavioral field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafTrafficPolicy) GetAntiautomationAndBehavioralOk() (*string, bool) {
	if o == nil || o.AntiautomationAndBehavioral == nil {
		return nil, false
	}
	return o.AntiautomationAndBehavioral, true
}

// HasAntiautomationAndBehavioral returns a boolean if a field has been set.
func (o *WafTrafficPolicy) HasAntiautomationAndBehavioral() bool {
	if o != nil && o.AntiautomationAndBehavioral != nil {
		return true
	}

	return false
}

// SetAntiautomationAndBehavioral gets a reference to the given string and assigns it to the AntiautomationAndBehavioral field.
func (o *WafTrafficPolicy) SetAntiautomationAndBehavioral(v string) {
	o.AntiautomationAndBehavioral = &v
}

// GetCsrf returns the Csrf field value if set, zero value otherwise.
func (o *WafTrafficPolicy) GetCsrf() string {
	if o == nil || o.Csrf == nil {
		var ret string
		return ret
	}
	return *o.Csrf
}

// GetCsrfOk returns a tuple with the Csrf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafTrafficPolicy) GetCsrfOk() (*string, bool) {
	if o == nil || o.Csrf == nil {
		return nil, false
	}
	return o.Csrf, true
}

// HasCsrf returns a boolean if a field has been set.
func (o *WafTrafficPolicy) HasCsrf() bool {
	if o != nil && o.Csrf != nil {
		return true
	}

	return false
}

// SetCsrf gets a reference to the given string and assigns it to the Csrf field.
func (o *WafTrafficPolicy) SetCsrf(v string) {
	o.Csrf = &v
}

// GetIpReputation returns the IpReputation field value if set, zero value otherwise.
func (o *WafTrafficPolicy) GetIpReputation() string {
	if o == nil || o.IpReputation == nil {
		var ret string
		return ret
	}
	return *o.IpReputation
}

// GetIpReputationOk returns a tuple with the IpReputation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafTrafficPolicy) GetIpReputationOk() (*string, bool) {
	if o == nil || o.IpReputation == nil {
		return nil, false
	}
	return o.IpReputation, true
}

// HasIpReputation returns a boolean if a field has been set.
func (o *WafTrafficPolicy) HasIpReputation() bool {
	if o != nil && o.IpReputation != nil {
		return true
	}

	return false
}

// SetIpReputation gets a reference to the given string and assigns it to the IpReputation field.
func (o *WafTrafficPolicy) SetIpReputation(v string) {
	o.IpReputation = &v
}

// GetSpamAndAbuse returns the SpamAndAbuse field value if set, zero value otherwise.
func (o *WafTrafficPolicy) GetSpamAndAbuse() string {
	if o == nil || o.SpamAndAbuse == nil {
		var ret string
		return ret
	}
	return *o.SpamAndAbuse
}

// GetSpamAndAbuseOk returns a tuple with the SpamAndAbuse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafTrafficPolicy) GetSpamAndAbuseOk() (*string, bool) {
	if o == nil || o.SpamAndAbuse == nil {
		return nil, false
	}
	return o.SpamAndAbuse, true
}

// HasSpamAndAbuse returns a boolean if a field has been set.
func (o *WafTrafficPolicy) HasSpamAndAbuse() bool {
	if o != nil && o.SpamAndAbuse != nil {
		return true
	}

	return false
}

// SetSpamAndAbuse gets a reference to the given string and assigns it to the SpamAndAbuse field.
func (o *WafTrafficPolicy) SetSpamAndAbuse(v string) {
	o.SpamAndAbuse = &v
}

// GetWaf returns the Waf field value if set, zero value otherwise.
func (o *WafTrafficPolicy) GetWaf() string {
	if o == nil || o.Waf == nil {
		var ret string
		return ret
	}
	return *o.Waf
}

// GetWafOk returns a tuple with the Waf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafTrafficPolicy) GetWafOk() (*string, bool) {
	if o == nil || o.Waf == nil {
		return nil, false
	}
	return o.Waf, true
}

// HasWaf returns a boolean if a field has been set.
func (o *WafTrafficPolicy) HasWaf() bool {
	if o != nil && o.Waf != nil {
		return true
	}

	return false
}

// SetWaf gets a reference to the given string and assigns it to the Waf field.
func (o *WafTrafficPolicy) SetWaf(v string) {
	o.Waf = &v
}

func (o WafTrafficPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AntiautomationAndBehavioral != nil {
		toSerialize["antiautomationAndBehavioral"] = o.AntiautomationAndBehavioral
	}
	if o.Csrf != nil {
		toSerialize["csrf"] = o.Csrf
	}
	if o.IpReputation != nil {
		toSerialize["ipReputation"] = o.IpReputation
	}
	if o.SpamAndAbuse != nil {
		toSerialize["spamAndAbuse"] = o.SpamAndAbuse
	}
	if o.Waf != nil {
		toSerialize["waf"] = o.Waf
	}
	return json.Marshal(toSerialize)
}

type NullableWafTrafficPolicy struct {
	value *WafTrafficPolicy
	isSet bool
}

func (v NullableWafTrafficPolicy) Get() *WafTrafficPolicy {
	return v.value
}

func (v *NullableWafTrafficPolicy) Set(val *WafTrafficPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableWafTrafficPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableWafTrafficPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWafTrafficPolicy(val *WafTrafficPolicy) *NullableWafTrafficPolicy {
	return &NullableWafTrafficPolicy{value: val, isSet: true}
}

func (v NullableWafTrafficPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWafTrafficPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
