/*
 * Edge Compute
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package edge-compute

import (
	"encoding/json"
	"time"
)

// V1ImageDeprecation An image's deprecation settings
type V1ImageDeprecation struct {
	// The date when an image is considered deprecated  Deprecated images may be used in new virtual machine based workloads, but are no longer considered for the \"default\" tag, nor are they returned in image lists by default.
	DeprecationDate *time.Time `json:"deprecationDate,omitempty"`
	// The date when an image is considered obsolete  Obsolete images may not be used in new virtual machine based workloads. If an obsoletion date is present then the deprecation date must also be present with a same or earlier date.
	ObsoletionDate *time.Time `json:"obsoletionDate,omitempty"`
	// The date when an image is deleted  Deletion dates cannot be before an image's deprecation or obsoletion dates.
	DeletionDate *time.Time `json:"deletionDate,omitempty"`
}

// NewV1ImageDeprecation instantiates a new V1ImageDeprecation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ImageDeprecation() *V1ImageDeprecation {
	this := V1ImageDeprecation{}
	return &this
}

// NewV1ImageDeprecationWithDefaults instantiates a new V1ImageDeprecation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ImageDeprecationWithDefaults() *V1ImageDeprecation {
	this := V1ImageDeprecation{}
	return &this
}

// GetDeprecationDate returns the DeprecationDate field value if set, zero value otherwise.
func (o *V1ImageDeprecation) GetDeprecationDate() time.Time {
	if o == nil || o.DeprecationDate == nil {
		var ret time.Time
		return ret
	}
	return *o.DeprecationDate
}

// GetDeprecationDateOk returns a tuple with the DeprecationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ImageDeprecation) GetDeprecationDateOk() (*time.Time, bool) {
	if o == nil || o.DeprecationDate == nil {
		return nil, false
	}
	return o.DeprecationDate, true
}

// HasDeprecationDate returns a boolean if a field has been set.
func (o *V1ImageDeprecation) HasDeprecationDate() bool {
	if o != nil && o.DeprecationDate != nil {
		return true
	}

	return false
}

// SetDeprecationDate gets a reference to the given time.Time and assigns it to the DeprecationDate field.
func (o *V1ImageDeprecation) SetDeprecationDate(v time.Time) {
	o.DeprecationDate = &v
}

// GetObsoletionDate returns the ObsoletionDate field value if set, zero value otherwise.
func (o *V1ImageDeprecation) GetObsoletionDate() time.Time {
	if o == nil || o.ObsoletionDate == nil {
		var ret time.Time
		return ret
	}
	return *o.ObsoletionDate
}

// GetObsoletionDateOk returns a tuple with the ObsoletionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ImageDeprecation) GetObsoletionDateOk() (*time.Time, bool) {
	if o == nil || o.ObsoletionDate == nil {
		return nil, false
	}
	return o.ObsoletionDate, true
}

// HasObsoletionDate returns a boolean if a field has been set.
func (o *V1ImageDeprecation) HasObsoletionDate() bool {
	if o != nil && o.ObsoletionDate != nil {
		return true
	}

	return false
}

// SetObsoletionDate gets a reference to the given time.Time and assigns it to the ObsoletionDate field.
func (o *V1ImageDeprecation) SetObsoletionDate(v time.Time) {
	o.ObsoletionDate = &v
}

// GetDeletionDate returns the DeletionDate field value if set, zero value otherwise.
func (o *V1ImageDeprecation) GetDeletionDate() time.Time {
	if o == nil || o.DeletionDate == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletionDate
}

// GetDeletionDateOk returns a tuple with the DeletionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ImageDeprecation) GetDeletionDateOk() (*time.Time, bool) {
	if o == nil || o.DeletionDate == nil {
		return nil, false
	}
	return o.DeletionDate, true
}

// HasDeletionDate returns a boolean if a field has been set.
func (o *V1ImageDeprecation) HasDeletionDate() bool {
	if o != nil && o.DeletionDate != nil {
		return true
	}

	return false
}

// SetDeletionDate gets a reference to the given time.Time and assigns it to the DeletionDate field.
func (o *V1ImageDeprecation) SetDeletionDate(v time.Time) {
	o.DeletionDate = &v
}

func (o V1ImageDeprecation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeprecationDate != nil {
		toSerialize["deprecationDate"] = o.DeprecationDate
	}
	if o.ObsoletionDate != nil {
		toSerialize["obsoletionDate"] = o.ObsoletionDate
	}
	if o.DeletionDate != nil {
		toSerialize["deletionDate"] = o.DeletionDate
	}
	return json.Marshal(toSerialize)
}

type NullableV1ImageDeprecation struct {
	value *V1ImageDeprecation
	isSet bool
}

func (v NullableV1ImageDeprecation) Get() *V1ImageDeprecation {
	return v.value
}

func (v *NullableV1ImageDeprecation) Set(val *V1ImageDeprecation) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ImageDeprecation) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ImageDeprecation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ImageDeprecation(val *V1ImageDeprecation) *NullableV1ImageDeprecation {
	return &NullableV1ImageDeprecation{value: val, isSet: true}
}

func (v NullableV1ImageDeprecation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ImageDeprecation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}