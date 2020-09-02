# Workloadv1Instance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StackId** | **string** | The ID of the stack that an instance belongs to | [optional] 
**WorkloadId** | **string** | The ID of the workload that an instance belongs to | [optional] 
**Id** | **string** | An instance&#39;s unique identifier | [optional] 
**Name** | **string** | An instance&#39;s name  Instance names are generated from their corresponsing workload&#39;s slug, followed by a unique hash | [optional] 
**Metadata** | [**V1Metadata**](v1Metadata.md) |  | [optional] 
**Location** | [**Workloadv1Location**](workloadv1Location.md) |  | [optional] 
**Phase** | [**Workloadv1InstanceInstancePhase**](workloadv1InstanceInstancePhase.md) |  | [optional] 
**IpAddress** | **string** | An instance&#39;s IP address | [optional] 
**ExternalIpAddress** | **string** | An instance&#39;s publicly accessible IP address | [optional] 
**CreatedAt** | [**time.Time**](time.Time.md) | The date an instance was created | [optional] 
**StartedAt** | [**time.Time**](time.Time.md) | The date an instance was started | [optional] 
**DeletedAt** | [**time.Time**](time.Time.md) | The date an instance was deleted | [optional] 
**NetworkInterfaces** | [**[]Workloadv1NetworkInterfaceStatus**](workloadv1NetworkInterfaceStatus.md) | An instance&#39;s network interfaces | [optional] 
**Resources** | [**V1ResourceRequirements**](v1ResourceRequirements.md) |  | [optional] 
**Containers** | [**map[string]V1ContainerSpec**](v1ContainerSpec.md) | A string to container configuration key/value pair | [optional] 
**ContainerStatuses** | [**[]V1ContainerStatus**](v1ContainerStatus.md) | Status of the containers running within the workload instance | [optional] 
**VirtualMachines** | [**map[string]V1VirtualMachineSpec**](v1VirtualMachineSpec.md) | A string to virtual machine configuration key/value pair | [optional] 
**VirtualMachineStatuses** | [**[]V1VirtualMachineStatus**](v1VirtualMachineStatus.md) | The status of the virtual machines running within the workload instance | [optional] 
**Reason** | **string** | A short reason that explains why an instance is in a phase | [optional] 
**Message** | **string** | A longer message that provides more detail on why an instance is in a phase | [optional] 
**ScheduledAt** | [**time.Time**](time.Time.md) | The date an instance was scheduled | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


