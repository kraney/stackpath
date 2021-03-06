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

// RuleStatusCode A custom HTTP status code that the WAF returns if a rule blocks a request
type RuleStatusCode string

// List of RuleStatusCode
const (
	RULESTATUSCODE_FORBIDDEN_403 RuleStatusCode = "FORBIDDEN_403"
	RULESTATUSCODE_TOO_MANY_REQUESTS_429 RuleStatusCode = "TOO_MANY_REQUESTS_429"
)

// Ptr returns reference to RuleStatusCode value
func (v RuleStatusCode) Ptr() *RuleStatusCode {
	return &v
}


type NullableRuleStatusCode struct {
	value *RuleStatusCode
	isSet bool
}

func (v NullableRuleStatusCode) Get() *RuleStatusCode {
	return v.value
}

func (v *NullableRuleStatusCode) Set(val *RuleStatusCode) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleStatusCode) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleStatusCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleStatusCode(val *RuleStatusCode) *NullableRuleStatusCode {
	return &NullableRuleStatusCode{value: val, isSet: true}
}

func (v NullableRuleStatusCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleStatusCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
