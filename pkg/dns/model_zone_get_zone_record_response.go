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

// ZoneGetZoneRecordResponse A response from a request to retrieve a DNS zone resource record
type ZoneGetZoneRecordResponse struct {
	Record *ZoneZoneRecord `json:"record,omitempty"`
}

// NewZoneGetZoneRecordResponse instantiates a new ZoneGetZoneRecordResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZoneGetZoneRecordResponse() *ZoneGetZoneRecordResponse {
	this := ZoneGetZoneRecordResponse{}
	return &this
}

// NewZoneGetZoneRecordResponseWithDefaults instantiates a new ZoneGetZoneRecordResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZoneGetZoneRecordResponseWithDefaults() *ZoneGetZoneRecordResponse {
	this := ZoneGetZoneRecordResponse{}
	return &this
}

// GetRecord returns the Record field value if set, zero value otherwise.
func (o *ZoneGetZoneRecordResponse) GetRecord() ZoneZoneRecord {
	if o == nil || o.Record == nil {
		var ret ZoneZoneRecord
		return ret
	}
	return *o.Record
}

// GetRecordOk returns a tuple with the Record field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneGetZoneRecordResponse) GetRecordOk() (*ZoneZoneRecord, bool) {
	if o == nil || o.Record == nil {
		return nil, false
	}
	return o.Record, true
}

// HasRecord returns a boolean if a field has been set.
func (o *ZoneGetZoneRecordResponse) HasRecord() bool {
	if o != nil && o.Record != nil {
		return true
	}

	return false
}

// SetRecord gets a reference to the given ZoneZoneRecord and assigns it to the Record field.
func (o *ZoneGetZoneRecordResponse) SetRecord(v ZoneZoneRecord) {
	o.Record = &v
}

func (o ZoneGetZoneRecordResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Record != nil {
		toSerialize["record"] = o.Record
	}
	return json.Marshal(toSerialize)
}

type NullableZoneGetZoneRecordResponse struct {
	value *ZoneGetZoneRecordResponse
	isSet bool
}

func (v NullableZoneGetZoneRecordResponse) Get() *ZoneGetZoneRecordResponse {
	return v.value
}

func (v *NullableZoneGetZoneRecordResponse) Set(val *ZoneGetZoneRecordResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableZoneGetZoneRecordResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableZoneGetZoneRecordResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZoneGetZoneRecordResponse(val *ZoneGetZoneRecordResponse) *NullableZoneGetZoneRecordResponse {
	return &NullableZoneGetZoneRecordResponse{value: val, isSet: true}
}

func (v NullableZoneGetZoneRecordResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZoneGetZoneRecordResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
