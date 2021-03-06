/*
 * Edge Compute
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package edge_compute

import (
	"encoding/json"
	"time"
)

// V1ImageCondition Further information about an image's status
type V1ImageCondition struct {
	// Type of the condition
	Type *string `json:"type,omitempty"`
	Status *V1ImageConditionStatus `json:"status,omitempty"`
	// The last time the condition was checked
	CheckedAt *time.Time `json:"checkedAt,omitempty"`
	// The last time the condition transitioned status
	TransitionedAt *time.Time `json:"transitionedAt,omitempty"`
	// A stable identifier for the reason the condition is in its current state
	Reason *string `json:"reason,omitempty"`
	// A human readable message with details regarding the condition
	Message *string `json:"message,omitempty"`
}

// NewV1ImageCondition instantiates a new V1ImageCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ImageCondition() *V1ImageCondition {
	this := V1ImageCondition{}
	var status V1ImageConditionStatus = "IMAGE_CONDITION_STATUS_UNKNOWN"
	this.Status = &status
	return &this
}

// NewV1ImageConditionWithDefaults instantiates a new V1ImageCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ImageConditionWithDefaults() *V1ImageCondition {
	this := V1ImageCondition{}
	var status V1ImageConditionStatus = "IMAGE_CONDITION_STATUS_UNKNOWN"
	this.Status = &status
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *V1ImageCondition) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ImageCondition) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *V1ImageCondition) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *V1ImageCondition) SetType(v string) {
	o.Type = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *V1ImageCondition) GetStatus() V1ImageConditionStatus {
	if o == nil || o.Status == nil {
		var ret V1ImageConditionStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ImageCondition) GetStatusOk() (*V1ImageConditionStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *V1ImageCondition) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given V1ImageConditionStatus and assigns it to the Status field.
func (o *V1ImageCondition) SetStatus(v V1ImageConditionStatus) {
	o.Status = &v
}

// GetCheckedAt returns the CheckedAt field value if set, zero value otherwise.
func (o *V1ImageCondition) GetCheckedAt() time.Time {
	if o == nil || o.CheckedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CheckedAt
}

// GetCheckedAtOk returns a tuple with the CheckedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ImageCondition) GetCheckedAtOk() (*time.Time, bool) {
	if o == nil || o.CheckedAt == nil {
		return nil, false
	}
	return o.CheckedAt, true
}

// HasCheckedAt returns a boolean if a field has been set.
func (o *V1ImageCondition) HasCheckedAt() bool {
	if o != nil && o.CheckedAt != nil {
		return true
	}

	return false
}

// SetCheckedAt gets a reference to the given time.Time and assigns it to the CheckedAt field.
func (o *V1ImageCondition) SetCheckedAt(v time.Time) {
	o.CheckedAt = &v
}

// GetTransitionedAt returns the TransitionedAt field value if set, zero value otherwise.
func (o *V1ImageCondition) GetTransitionedAt() time.Time {
	if o == nil || o.TransitionedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.TransitionedAt
}

// GetTransitionedAtOk returns a tuple with the TransitionedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ImageCondition) GetTransitionedAtOk() (*time.Time, bool) {
	if o == nil || o.TransitionedAt == nil {
		return nil, false
	}
	return o.TransitionedAt, true
}

// HasTransitionedAt returns a boolean if a field has been set.
func (o *V1ImageCondition) HasTransitionedAt() bool {
	if o != nil && o.TransitionedAt != nil {
		return true
	}

	return false
}

// SetTransitionedAt gets a reference to the given time.Time and assigns it to the TransitionedAt field.
func (o *V1ImageCondition) SetTransitionedAt(v time.Time) {
	o.TransitionedAt = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *V1ImageCondition) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ImageCondition) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *V1ImageCondition) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *V1ImageCondition) SetReason(v string) {
	o.Reason = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *V1ImageCondition) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ImageCondition) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *V1ImageCondition) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *V1ImageCondition) SetMessage(v string) {
	o.Message = &v
}

func (o V1ImageCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.CheckedAt != nil {
		toSerialize["checkedAt"] = o.CheckedAt
	}
	if o.TransitionedAt != nil {
		toSerialize["transitionedAt"] = o.TransitionedAt
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableV1ImageCondition struct {
	value *V1ImageCondition
	isSet bool
}

func (v NullableV1ImageCondition) Get() *V1ImageCondition {
	return v.value
}

func (v *NullableV1ImageCondition) Set(val *V1ImageCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ImageCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ImageCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ImageCondition(val *V1ImageCondition) *NullableV1ImageCondition {
	return &NullableV1ImageCondition{value: val, isSet: true}
}

func (v NullableV1ImageCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ImageCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
