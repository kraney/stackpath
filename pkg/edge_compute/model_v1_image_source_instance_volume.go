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
)

// V1ImageSourceInstanceVolume A reference to the volume of an instance to capture an image from
type V1ImageSourceInstanceVolume struct {
	// The ID or slug of the workload containing the instance to reference
	WorkloadId *string `json:"workloadId,omitempty"`
	// An instance's name to reference a volume on
	InstanceName *string `json:"instanceName,omitempty"`
	// A reference to the volume to create an image from or unset for the root volume
	VolumeClaimSlug *string `json:"volumeClaimSlug,omitempty"`
}

// NewV1ImageSourceInstanceVolume instantiates a new V1ImageSourceInstanceVolume object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ImageSourceInstanceVolume() *V1ImageSourceInstanceVolume {
	this := V1ImageSourceInstanceVolume{}
	return &this
}

// NewV1ImageSourceInstanceVolumeWithDefaults instantiates a new V1ImageSourceInstanceVolume object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ImageSourceInstanceVolumeWithDefaults() *V1ImageSourceInstanceVolume {
	this := V1ImageSourceInstanceVolume{}
	return &this
}

// GetWorkloadId returns the WorkloadId field value if set, zero value otherwise.
func (o *V1ImageSourceInstanceVolume) GetWorkloadId() string {
	if o == nil || o.WorkloadId == nil {
		var ret string
		return ret
	}
	return *o.WorkloadId
}

// GetWorkloadIdOk returns a tuple with the WorkloadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ImageSourceInstanceVolume) GetWorkloadIdOk() (*string, bool) {
	if o == nil || o.WorkloadId == nil {
		return nil, false
	}
	return o.WorkloadId, true
}

// HasWorkloadId returns a boolean if a field has been set.
func (o *V1ImageSourceInstanceVolume) HasWorkloadId() bool {
	if o != nil && o.WorkloadId != nil {
		return true
	}

	return false
}

// SetWorkloadId gets a reference to the given string and assigns it to the WorkloadId field.
func (o *V1ImageSourceInstanceVolume) SetWorkloadId(v string) {
	o.WorkloadId = &v
}

// GetInstanceName returns the InstanceName field value if set, zero value otherwise.
func (o *V1ImageSourceInstanceVolume) GetInstanceName() string {
	if o == nil || o.InstanceName == nil {
		var ret string
		return ret
	}
	return *o.InstanceName
}

// GetInstanceNameOk returns a tuple with the InstanceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ImageSourceInstanceVolume) GetInstanceNameOk() (*string, bool) {
	if o == nil || o.InstanceName == nil {
		return nil, false
	}
	return o.InstanceName, true
}

// HasInstanceName returns a boolean if a field has been set.
func (o *V1ImageSourceInstanceVolume) HasInstanceName() bool {
	if o != nil && o.InstanceName != nil {
		return true
	}

	return false
}

// SetInstanceName gets a reference to the given string and assigns it to the InstanceName field.
func (o *V1ImageSourceInstanceVolume) SetInstanceName(v string) {
	o.InstanceName = &v
}

// GetVolumeClaimSlug returns the VolumeClaimSlug field value if set, zero value otherwise.
func (o *V1ImageSourceInstanceVolume) GetVolumeClaimSlug() string {
	if o == nil || o.VolumeClaimSlug == nil {
		var ret string
		return ret
	}
	return *o.VolumeClaimSlug
}

// GetVolumeClaimSlugOk returns a tuple with the VolumeClaimSlug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ImageSourceInstanceVolume) GetVolumeClaimSlugOk() (*string, bool) {
	if o == nil || o.VolumeClaimSlug == nil {
		return nil, false
	}
	return o.VolumeClaimSlug, true
}

// HasVolumeClaimSlug returns a boolean if a field has been set.
func (o *V1ImageSourceInstanceVolume) HasVolumeClaimSlug() bool {
	if o != nil && o.VolumeClaimSlug != nil {
		return true
	}

	return false
}

// SetVolumeClaimSlug gets a reference to the given string and assigns it to the VolumeClaimSlug field.
func (o *V1ImageSourceInstanceVolume) SetVolumeClaimSlug(v string) {
	o.VolumeClaimSlug = &v
}

func (o V1ImageSourceInstanceVolume) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.WorkloadId != nil {
		toSerialize["workloadId"] = o.WorkloadId
	}
	if o.InstanceName != nil {
		toSerialize["instanceName"] = o.InstanceName
	}
	if o.VolumeClaimSlug != nil {
		toSerialize["volumeClaimSlug"] = o.VolumeClaimSlug
	}
	return json.Marshal(toSerialize)
}

type NullableV1ImageSourceInstanceVolume struct {
	value *V1ImageSourceInstanceVolume
	isSet bool
}

func (v NullableV1ImageSourceInstanceVolume) Get() *V1ImageSourceInstanceVolume {
	return v.value
}

func (v *NullableV1ImageSourceInstanceVolume) Set(val *V1ImageSourceInstanceVolume) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ImageSourceInstanceVolume) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ImageSourceInstanceVolume) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ImageSourceInstanceVolume(val *V1ImageSourceInstanceVolume) *NullableV1ImageSourceInstanceVolume {
	return &NullableV1ImageSourceInstanceVolume{value: val, isSet: true}
}

func (v NullableV1ImageSourceInstanceVolume) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ImageSourceInstanceVolume) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
