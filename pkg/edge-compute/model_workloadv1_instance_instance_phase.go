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
)

// Workloadv1InstanceInstancePhase An instance's startup state  - INSTANCE_PHASE_UNSPECIFIED: StackPath is unable to determine the instance's startup state  - STARTING: The instance is still initializing  - RUNNING: The instance is running  - FAILED: The instance failed to start  - COMPLETED: The instance finished running  - SCHEDULING: The instance is being scheduled  - STOPPED: The instance is stopped
type Workloadv1InstanceInstancePhase string

// List of workloadv1InstanceInstancePhase
const (
	INSTANCE_PHASE_UNSPECIFIED Workloadv1InstanceInstancePhase = "INSTANCE_PHASE_UNSPECIFIED"
	STARTING Workloadv1InstanceInstancePhase = "STARTING"
	RUNNING Workloadv1InstanceInstancePhase = "RUNNING"
	FAILED Workloadv1InstanceInstancePhase = "FAILED"
	COMPLETED Workloadv1InstanceInstancePhase = "COMPLETED"
	SCHEDULING Workloadv1InstanceInstancePhase = "SCHEDULING"
	STOPPED Workloadv1InstanceInstancePhase = "STOPPED"
)

// Ptr returns reference to workloadv1InstanceInstancePhase value
func (v Workloadv1InstanceInstancePhase) Ptr() *Workloadv1InstanceInstancePhase {
	return &v
}


type NullableWorkloadv1InstanceInstancePhase struct {
	value *Workloadv1InstanceInstancePhase
	isSet bool
}

func (v NullableWorkloadv1InstanceInstancePhase) Get() *Workloadv1InstanceInstancePhase {
	return v.value
}

func (v *NullableWorkloadv1InstanceInstancePhase) Set(val *Workloadv1InstanceInstancePhase) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkloadv1InstanceInstancePhase) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkloadv1InstanceInstancePhase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkloadv1InstanceInstancePhase(val *Workloadv1InstanceInstancePhase) *NullableWorkloadv1InstanceInstancePhase {
	return &NullableWorkloadv1InstanceInstancePhase{value: val, isSet: true}
}

func (v NullableWorkloadv1InstanceInstancePhase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkloadv1InstanceInstancePhase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}