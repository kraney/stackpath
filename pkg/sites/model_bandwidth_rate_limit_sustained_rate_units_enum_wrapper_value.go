/*
 * Sites
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package sites

import (
	"encoding/json"
)

// BandwidthRateLimitSustainedRateUnitsEnumWrapperValue the model 'BandwidthRateLimitSustainedRateUnitsEnumWrapperValue'
type BandwidthRateLimitSustainedRateUnitsEnumWrapperValue string

// List of BandwidthRateLimitSustainedRateUnitsEnumWrapperValue
const (
	UNKNOWN BandwidthRateLimitSustainedRateUnitsEnumWrapperValue = "UNKNOWN"
	KILOBIT BandwidthRateLimitSustainedRateUnitsEnumWrapperValue = "kilobit"
	KILOBYTE BandwidthRateLimitSustainedRateUnitsEnumWrapperValue = "kilobyte"
)

// Ptr returns reference to BandwidthRateLimitSustainedRateUnitsEnumWrapperValue value
func (v BandwidthRateLimitSustainedRateUnitsEnumWrapperValue) Ptr() *BandwidthRateLimitSustainedRateUnitsEnumWrapperValue {
	return &v
}


type NullableBandwidthRateLimitSustainedRateUnitsEnumWrapperValue struct {
	value *BandwidthRateLimitSustainedRateUnitsEnumWrapperValue
	isSet bool
}

func (v NullableBandwidthRateLimitSustainedRateUnitsEnumWrapperValue) Get() *BandwidthRateLimitSustainedRateUnitsEnumWrapperValue {
	return v.value
}

func (v *NullableBandwidthRateLimitSustainedRateUnitsEnumWrapperValue) Set(val *BandwidthRateLimitSustainedRateUnitsEnumWrapperValue) {
	v.value = val
	v.isSet = true
}

func (v NullableBandwidthRateLimitSustainedRateUnitsEnumWrapperValue) IsSet() bool {
	return v.isSet
}

func (v *NullableBandwidthRateLimitSustainedRateUnitsEnumWrapperValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBandwidthRateLimitSustainedRateUnitsEnumWrapperValue(val *BandwidthRateLimitSustainedRateUnitsEnumWrapperValue) *NullableBandwidthRateLimitSustainedRateUnitsEnumWrapperValue {
	return &NullableBandwidthRateLimitSustainedRateUnitsEnumWrapperValue{value: val, isSet: true}
}

func (v NullableBandwidthRateLimitSustainedRateUnitsEnumWrapperValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBandwidthRateLimitSustainedRateUnitsEnumWrapperValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}