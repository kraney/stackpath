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

// SearchFilterActionFilter Potential actions the WAF took on an event
type SearchFilterActionFilter string

// List of SearchFilterActionFilter
const (
	ANY_ACTION SearchFilterActionFilter = "ANY_ACTION"
	BLOCK_ACTION SearchFilterActionFilter = "BLOCK_ACTION"
	ALLOW_ACTION SearchFilterActionFilter = "ALLOW_ACTION"
	CAPTCHA_ACTION SearchFilterActionFilter = "CAPTCHA_ACTION"
	HANDSHAKE_ACTION SearchFilterActionFilter = "HANDSHAKE_ACTION"
	MONITOR_ACTION SearchFilterActionFilter = "MONITOR_ACTION"
)

// Ptr returns reference to SearchFilterActionFilter value
func (v SearchFilterActionFilter) Ptr() *SearchFilterActionFilter {
	return &v
}


type NullableSearchFilterActionFilter struct {
	value *SearchFilterActionFilter
	isSet bool
}

func (v NullableSearchFilterActionFilter) Get() *SearchFilterActionFilter {
	return v.value
}

func (v *NullableSearchFilterActionFilter) Set(val *SearchFilterActionFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchFilterActionFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchFilterActionFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchFilterActionFilter(val *SearchFilterActionFilter) *NullableSearchFilterActionFilter {
	return &NullableSearchFilterActionFilter{value: val, isSet: true}
}

func (v NullableSearchFilterActionFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchFilterActionFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}