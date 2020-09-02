# V1WorkloadSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkInterfaces** | [**[]V1NetworkInterface**](v1NetworkInterface.md) | Network interfaces to bind to the workload&#39;s instances | [optional] 
**Containers** | [**map[string]V1ContainerSpec**](v1ContainerSpec.md) | A string to container configuration key/value pair | [optional] 
**VirtualMachines** | [**map[string]V1VirtualMachineSpec**](v1VirtualMachineSpec.md) | A string to virtual machine configuration key/value pair | [optional] 
**VolumeClaimTemplates** | [**[]V1VolumeClaim**](v1VolumeClaim.md) | A list of claims that instances may reference  The slug of the claim will be used in combination with the name of the instance to create a stable identifier. The slug should be used in the volume mount specifications for containers and VMs. | [optional] 
**ImagePullCredentials** | [**[]V1ImagePullCredential**](v1ImagePullCredential.md) | The credentials needed to pull a container image | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


