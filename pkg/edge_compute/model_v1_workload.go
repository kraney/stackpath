/*
 * Edge Compute
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package edge_compute
// V1Workload A computing workload  Workloads define a computing application to deploy to StackPath's edge network.
type V1Workload struct {
	// The ID of the stack that a workload belongs to
	StackId string `json:"stackId,omitempty"`
	// A workload's unique identifier
	Id string `json:"id,omitempty"`
	// A workload's name as it appears in the StackPath portal
	Name string `json:"name,omitempty"`
	// A workload's programmatic name  Workload slugs are used to build its instances names
	Slug string `json:"slug,omitempty"`
	Metadata V1Metadata `json:"metadata,omitempty"`
	Spec V1WorkloadSpec `json:"spec,omitempty"`
	// A string to deployment target key/value pair
	Targets map[string]V1Target `json:"targets,omitempty"`
	Status V1WorkloadStatus `json:"status,omitempty"`
}