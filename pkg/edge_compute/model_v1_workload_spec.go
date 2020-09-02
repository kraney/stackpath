/*
 * Edge Compute
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package edge_compute
// V1WorkloadSpec The specification for the desired state of a workload
type V1WorkloadSpec struct {
	// Network interfaces to bind to the workload's instances
	NetworkInterfaces []V1NetworkInterface `json:"networkInterfaces,omitempty"`
	// A string to container configuration key/value pair
	Containers map[string]V1ContainerSpec `json:"containers,omitempty"`
	// A string to virtual machine configuration key/value pair
	VirtualMachines map[string]V1VirtualMachineSpec `json:"virtualMachines,omitempty"`
	// A list of claims that instances may reference  The slug of the claim will be used in combination with the name of the instance to create a stable identifier. The slug should be used in the volume mount specifications for containers and VMs.
	VolumeClaimTemplates []V1VolumeClaim `json:"volumeClaimTemplates,omitempty"`
	// The credentials needed to pull a container image
	ImagePullCredentials []V1ImagePullCredential `json:"imagePullCredentials,omitempty"`
}