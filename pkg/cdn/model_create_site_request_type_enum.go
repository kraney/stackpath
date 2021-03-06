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

// CreateSiteRequestTypeEnum The CDN site's type  A site's type determines how StackPath delivers content to incoming HTTP(S) requests.   - CDN: The site is a CDN only site  - WAF: The site is either a standalone WAF site or a WAF site with attached CDN service  - API: The site is an API delivery site. API delivery sites are powered by both the WAF and CDN and have custom rulesets for each.
type CreateSiteRequestTypeEnum string

// List of CreateSiteRequestTypeEnum
const (
	CREATESITEREQUESTTYPEENUM_CDN CreateSiteRequestTypeEnum = "CDN"
	CREATESITEREQUESTTYPEENUM_WAF CreateSiteRequestTypeEnum = "WAF"
	CREATESITEREQUESTTYPEENUM_API CreateSiteRequestTypeEnum = "API"
)

// Ptr returns reference to CreateSiteRequestTypeEnum value
func (v CreateSiteRequestTypeEnum) Ptr() *CreateSiteRequestTypeEnum {
	return &v
}


type NullableCreateSiteRequestTypeEnum struct {
	value *CreateSiteRequestTypeEnum
	isSet bool
}

func (v NullableCreateSiteRequestTypeEnum) Get() *CreateSiteRequestTypeEnum {
	return v.value
}

func (v *NullableCreateSiteRequestTypeEnum) Set(val *CreateSiteRequestTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSiteRequestTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSiteRequestTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSiteRequestTypeEnum(val *CreateSiteRequestTypeEnum) *NullableCreateSiteRequestTypeEnum {
	return &NullableCreateSiteRequestTypeEnum{value: val, isSet: true}
}

func (v NullableCreateSiteRequestTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSiteRequestTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
