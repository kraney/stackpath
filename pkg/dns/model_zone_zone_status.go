/*
 * DNS
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dns

import (
	"encoding/json"
)

// ZoneZoneStatus the model 'ZoneZoneStatus'
type ZoneZoneStatus string

// List of zoneZoneStatus
const (
	ZONEZONESTATUS_ACTIVE ZoneZoneStatus = "ACTIVE"
	ZONEZONESTATUS_SUSPENDED ZoneZoneStatus = "SUSPENDED"
	ZONEZONESTATUS_BILLING_SUSPENDED ZoneZoneStatus = "BILLING_SUSPENDED"
	ZONEZONESTATUS_INACTIVE ZoneZoneStatus = "INACTIVE"
)

// Ptr returns reference to zoneZoneStatus value
func (v ZoneZoneStatus) Ptr() *ZoneZoneStatus {
	return &v
}


type NullableZoneZoneStatus struct {
	value *ZoneZoneStatus
	isSet bool
}

func (v NullableZoneZoneStatus) Get() *ZoneZoneStatus {
	return v.value
}

func (v *NullableZoneZoneStatus) Set(val *ZoneZoneStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableZoneZoneStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableZoneZoneStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZoneZoneStatus(val *ZoneZoneStatus) *NullableZoneZoneStatus {
	return &NullableZoneZoneStatus{value: val, isSet: true}
}

func (v NullableZoneZoneStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZoneZoneStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
