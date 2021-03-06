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

// WafRule A custom WAF rule applied to a site  Unlike pre-defined policies, rules are provided by the user for custom WAF functionality.
type WafRule struct {
	// A rule's ID
	Id *string `json:"id,omitempty"`
	// A rule's name
	Name *string `json:"name,omitempty"`
	// A rule's description
	Description *string `json:"description,omitempty"`
	// The conditions required for the WAF engine to trigger the rule  Rules may have between 1 and 5 conditions. All conditions must pass for the rule to trigger.
	Conditions *[]RuleCondition `json:"conditions,omitempty"`
	Action *WafRuleAction `json:"action,omitempty"`
	// Whether or not a WAF rule is enabled
	Enabled *bool `json:"enabled,omitempty"`
	// How long a rule's block action will apply to subsequent requests  Durations only apply to rules with block actions.
	ActionDuration *string `json:"actionDuration,omitempty"`
	StatusCode *RuleStatusCode `json:"statusCode,omitempty"`
}

// NewWafRule instantiates a new WafRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafRule() *WafRule {
	this := WafRule{}
	var action WafRuleAction = "BLOCK"
	this.Action = &action
	var statusCode RuleStatusCode = "FORBIDDEN_403"
	this.StatusCode = &statusCode
	return &this
}

// NewWafRuleWithDefaults instantiates a new WafRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafRuleWithDefaults() *WafRule {
	this := WafRule{}
	var action WafRuleAction = "BLOCK"
	this.Action = &action
	var statusCode RuleStatusCode = "FORBIDDEN_403"
	this.StatusCode = &statusCode
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WafRule) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafRule) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WafRule) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WafRule) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WafRule) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafRule) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WafRule) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WafRule) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WafRule) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafRule) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WafRule) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WafRule) SetDescription(v string) {
	o.Description = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *WafRule) GetConditions() []RuleCondition {
	if o == nil || o.Conditions == nil {
		var ret []RuleCondition
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafRule) GetConditionsOk() (*[]RuleCondition, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *WafRule) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []RuleCondition and assigns it to the Conditions field.
func (o *WafRule) SetConditions(v []RuleCondition) {
	o.Conditions = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *WafRule) GetAction() WafRuleAction {
	if o == nil || o.Action == nil {
		var ret WafRuleAction
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafRule) GetActionOk() (*WafRuleAction, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *WafRule) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given WafRuleAction and assigns it to the Action field.
func (o *WafRule) SetAction(v WafRuleAction) {
	o.Action = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *WafRule) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafRule) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *WafRule) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *WafRule) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetActionDuration returns the ActionDuration field value if set, zero value otherwise.
func (o *WafRule) GetActionDuration() string {
	if o == nil || o.ActionDuration == nil {
		var ret string
		return ret
	}
	return *o.ActionDuration
}

// GetActionDurationOk returns a tuple with the ActionDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafRule) GetActionDurationOk() (*string, bool) {
	if o == nil || o.ActionDuration == nil {
		return nil, false
	}
	return o.ActionDuration, true
}

// HasActionDuration returns a boolean if a field has been set.
func (o *WafRule) HasActionDuration() bool {
	if o != nil && o.ActionDuration != nil {
		return true
	}

	return false
}

// SetActionDuration gets a reference to the given string and assigns it to the ActionDuration field.
func (o *WafRule) SetActionDuration(v string) {
	o.ActionDuration = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *WafRule) GetStatusCode() RuleStatusCode {
	if o == nil || o.StatusCode == nil {
		var ret RuleStatusCode
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafRule) GetStatusCodeOk() (*RuleStatusCode, bool) {
	if o == nil || o.StatusCode == nil {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *WafRule) HasStatusCode() bool {
	if o != nil && o.StatusCode != nil {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given RuleStatusCode and assigns it to the StatusCode field.
func (o *WafRule) SetStatusCode(v RuleStatusCode) {
	o.StatusCode = &v
}

func (o WafRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.ActionDuration != nil {
		toSerialize["actionDuration"] = o.ActionDuration
	}
	if o.StatusCode != nil {
		toSerialize["statusCode"] = o.StatusCode
	}
	return json.Marshal(toSerialize)
}

type NullableWafRule struct {
	value *WafRule
	isSet bool
}

func (v NullableWafRule) Get() *WafRule {
	return v.value
}

func (v *NullableWafRule) Set(val *WafRule) {
	v.value = val
	v.isSet = true
}

func (v NullableWafRule) IsSet() bool {
	return v.isSet
}

func (v *NullableWafRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWafRule(val *WafRule) *NullableWafRule {
	return &NullableWafRule{value: val, isSet: true}
}

func (v NullableWafRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWafRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
