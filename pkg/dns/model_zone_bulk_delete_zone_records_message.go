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

// ZoneBulkDeleteZoneRecordsMessage The fields needed to delete multiple DNS zone resource records
type ZoneBulkDeleteZoneRecordsMessage struct {
	// The IDs of the resource records to delete
	ZoneRecordIds *[]string `json:"zoneRecordIds,omitempty"`
}

// NewZoneBulkDeleteZoneRecordsMessage instantiates a new ZoneBulkDeleteZoneRecordsMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZoneBulkDeleteZoneRecordsMessage() *ZoneBulkDeleteZoneRecordsMessage {
	this := ZoneBulkDeleteZoneRecordsMessage{}
	return &this
}

// NewZoneBulkDeleteZoneRecordsMessageWithDefaults instantiates a new ZoneBulkDeleteZoneRecordsMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZoneBulkDeleteZoneRecordsMessageWithDefaults() *ZoneBulkDeleteZoneRecordsMessage {
	this := ZoneBulkDeleteZoneRecordsMessage{}
	return &this
}

// GetZoneRecordIds returns the ZoneRecordIds field value if set, zero value otherwise.
func (o *ZoneBulkDeleteZoneRecordsMessage) GetZoneRecordIds() []string {
	if o == nil || o.ZoneRecordIds == nil {
		var ret []string
		return ret
	}
	return *o.ZoneRecordIds
}

// GetZoneRecordIdsOk returns a tuple with the ZoneRecordIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneBulkDeleteZoneRecordsMessage) GetZoneRecordIdsOk() (*[]string, bool) {
	if o == nil || o.ZoneRecordIds == nil {
		return nil, false
	}
	return o.ZoneRecordIds, true
}

// HasZoneRecordIds returns a boolean if a field has been set.
func (o *ZoneBulkDeleteZoneRecordsMessage) HasZoneRecordIds() bool {
	if o != nil && o.ZoneRecordIds != nil {
		return true
	}

	return false
}

// SetZoneRecordIds gets a reference to the given []string and assigns it to the ZoneRecordIds field.
func (o *ZoneBulkDeleteZoneRecordsMessage) SetZoneRecordIds(v []string) {
	o.ZoneRecordIds = &v
}

func (o ZoneBulkDeleteZoneRecordsMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ZoneRecordIds != nil {
		toSerialize["zoneRecordIds"] = o.ZoneRecordIds
	}
	return json.Marshal(toSerialize)
}

type NullableZoneBulkDeleteZoneRecordsMessage struct {
	value *ZoneBulkDeleteZoneRecordsMessage
	isSet bool
}

func (v NullableZoneBulkDeleteZoneRecordsMessage) Get() *ZoneBulkDeleteZoneRecordsMessage {
	return v.value
}

func (v *NullableZoneBulkDeleteZoneRecordsMessage) Set(val *ZoneBulkDeleteZoneRecordsMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableZoneBulkDeleteZoneRecordsMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableZoneBulkDeleteZoneRecordsMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZoneBulkDeleteZoneRecordsMessage(val *ZoneBulkDeleteZoneRecordsMessage) *NullableZoneBulkDeleteZoneRecordsMessage {
	return &NullableZoneBulkDeleteZoneRecordsMessage{value: val, isSet: true}
}

func (v NullableZoneBulkDeleteZoneRecordsMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZoneBulkDeleteZoneRecordsMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
