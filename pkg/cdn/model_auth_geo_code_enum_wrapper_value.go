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

// AuthGeoCodeEnumWrapperValue the model 'AuthGeoCodeEnumWrapperValue'
type AuthGeoCodeEnumWrapperValue string

// List of AuthGeoCodeEnumWrapperValue
const (
	UNKNOWN AuthGeoCodeEnumWrapperValue = "UNKNOWN"
	COUNTRY_CODE AuthGeoCodeEnumWrapperValue = "countryCode"
	REGION AuthGeoCodeEnumWrapperValue = "region"
	SUBDIVISION_CODES AuthGeoCodeEnumWrapperValue = "subdivisionCodes"
	CITY AuthGeoCodeEnumWrapperValue = "city"
	POSTAL_CODE AuthGeoCodeEnumWrapperValue = "postalCode"
	CONTINENT_CODE AuthGeoCodeEnumWrapperValue = "continentCode"
	TIME_ZONE AuthGeoCodeEnumWrapperValue = "timeZone"
	DMA_CODE AuthGeoCodeEnumWrapperValue = "dmaCode"
	AREA_CODE AuthGeoCodeEnumWrapperValue = "areaCode"
)

// Ptr returns reference to AuthGeoCodeEnumWrapperValue value
func (v AuthGeoCodeEnumWrapperValue) Ptr() *AuthGeoCodeEnumWrapperValue {
	return &v
}


type NullableAuthGeoCodeEnumWrapperValue struct {
	value *AuthGeoCodeEnumWrapperValue
	isSet bool
}

func (v NullableAuthGeoCodeEnumWrapperValue) Get() *AuthGeoCodeEnumWrapperValue {
	return v.value
}

func (v *NullableAuthGeoCodeEnumWrapperValue) Set(val *AuthGeoCodeEnumWrapperValue) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthGeoCodeEnumWrapperValue) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthGeoCodeEnumWrapperValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthGeoCodeEnumWrapperValue(val *AuthGeoCodeEnumWrapperValue) *NullableAuthGeoCodeEnumWrapperValue {
	return &NullableAuthGeoCodeEnumWrapperValue{value: val, isSet: true}
}

func (v NullableAuthGeoCodeEnumWrapperValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthGeoCodeEnumWrapperValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}