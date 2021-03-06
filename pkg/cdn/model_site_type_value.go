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

// SiteTypeValue A CDN site's type  A site's type determines how StackPath delivers content to incoming HTTP(S) requests.   - UNKNOWN: StackPath is unable to determine a site's type  - CDN: A site is CDN only site  - WAF: A site is either a standalone WAF site or a WAF site with attached CDN service  - API: A site is an API delivery site. API delivery sites are powered by both the WAF and CDN and have custom rulesets for each.
type SiteTypeValue string

// List of SiteTypeValue
const (
	SITETYPEVALUE_UNKNOWN SiteTypeValue = "UNKNOWN"
	SITETYPEVALUE_CDN SiteTypeValue = "CDN"
	SITETYPEVALUE_WAF SiteTypeValue = "WAF"
	SITETYPEVALUE_API SiteTypeValue = "API"
)

// Ptr returns reference to SiteTypeValue value
func (v SiteTypeValue) Ptr() *SiteTypeValue {
	return &v
}


type NullableSiteTypeValue struct {
	value *SiteTypeValue
	isSet bool
}

func (v NullableSiteTypeValue) Get() *SiteTypeValue {
	return v.value
}

func (v *NullableSiteTypeValue) Set(val *SiteTypeValue) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteTypeValue) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteTypeValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteTypeValue(val *SiteTypeValue) *NullableSiteTypeValue {
	return &NullableSiteTypeValue{value: val, isSet: true}
}

func (v NullableSiteTypeValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteTypeValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
