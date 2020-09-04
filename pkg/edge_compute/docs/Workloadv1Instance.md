# Workloadv1Instance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StackId** | Pointer to **string** | The ID of the stack that an instance belongs to | [optional] 
**WorkloadId** | Pointer to **string** | The ID of the workload that an instance belongs to | [optional] 
**Id** | Pointer to **string** | An instance&#39;s unique identifier | [optional] 
**Name** | Pointer to **string** | An instance&#39;s name  Instance names are generated from their corresponsing workload&#39;s slug, followed by a unique hash | [optional] 
**Metadata** | Pointer to [**V1Metadata**](v1Metadata.md) |  | [optional] 
**Location** | Pointer to [**Workloadv1Location**](workloadv1Location.md) |  | [optional] 
**Phase** | Pointer to [**Workloadv1InstanceInstancePhase**](workloadv1InstanceInstancePhase.md) |  | [optional] [default to "INSTANCE_PHASE_UNSPECIFIED"]
**IpAddress** | Pointer to **string** | An instance&#39;s IP address | [optional] 
**ExternalIpAddress** | Pointer to **string** | An instance&#39;s publicly accessible IP address | [optional] 
**CreatedAt** | Pointer to [**time.Time**](time.Time.md) | The date an instance was created | [optional] 
**StartedAt** | Pointer to [**time.Time**](time.Time.md) | The date an instance was started | [optional] 
**DeletedAt** | Pointer to [**time.Time**](time.Time.md) | The date an instance was deleted | [optional] 
**NetworkInterfaces** | Pointer to [**[]Workloadv1NetworkInterfaceStatus**](workloadv1NetworkInterfaceStatus.md) | An instance&#39;s network interfaces | [optional] 
**Resources** | Pointer to [**V1ResourceRequirements**](v1ResourceRequirements.md) |  | [optional] 
**Containers** | Pointer to [**map[string]V1ContainerSpec**](v1ContainerSpec.md) | A string to container configuration key/value pair | [optional] 
**ContainerStatuses** | Pointer to [**[]V1ContainerStatus**](v1ContainerStatus.md) | Status of the containers running within the workload instance | [optional] 
**VirtualMachines** | Pointer to [**map[string]V1VirtualMachineSpec**](v1VirtualMachineSpec.md) | A string to virtual machine configuration key/value pair | [optional] 
**VirtualMachineStatuses** | Pointer to [**[]V1VirtualMachineStatus**](v1VirtualMachineStatus.md) | The status of the virtual machines running within the workload instance | [optional] 
**Reason** | Pointer to **string** | A short reason that explains why an instance is in a phase | [optional] 
**Message** | Pointer to **string** | A longer message that provides more detail on why an instance is in a phase | [optional] 
**ScheduledAt** | Pointer to [**time.Time**](time.Time.md) | The date an instance was scheduled | [optional] 

## Methods

### NewWorkloadv1Instance

`func NewWorkloadv1Instance() *Workloadv1Instance`

NewWorkloadv1Instance instantiates a new Workloadv1Instance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadv1InstanceWithDefaults

`func NewWorkloadv1InstanceWithDefaults() *Workloadv1Instance`

NewWorkloadv1InstanceWithDefaults instantiates a new Workloadv1Instance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStackId

`func (o *Workloadv1Instance) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *Workloadv1Instance) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *Workloadv1Instance) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *Workloadv1Instance) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetWorkloadId

`func (o *Workloadv1Instance) GetWorkloadId() string`

GetWorkloadId returns the WorkloadId field if non-nil, zero value otherwise.

### GetWorkloadIdOk

`func (o *Workloadv1Instance) GetWorkloadIdOk() (*string, bool)`

GetWorkloadIdOk returns a tuple with the WorkloadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadId

`func (o *Workloadv1Instance) SetWorkloadId(v string)`

SetWorkloadId sets WorkloadId field to given value.

### HasWorkloadId

`func (o *Workloadv1Instance) HasWorkloadId() bool`

HasWorkloadId returns a boolean if a field has been set.

### GetId

`func (o *Workloadv1Instance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Workloadv1Instance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Workloadv1Instance) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Workloadv1Instance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Workloadv1Instance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Workloadv1Instance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Workloadv1Instance) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Workloadv1Instance) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMetadata

`func (o *Workloadv1Instance) GetMetadata() V1Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Workloadv1Instance) GetMetadataOk() (*V1Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Workloadv1Instance) SetMetadata(v V1Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Workloadv1Instance) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetLocation

`func (o *Workloadv1Instance) GetLocation() Workloadv1Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Workloadv1Instance) GetLocationOk() (*Workloadv1Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Workloadv1Instance) SetLocation(v Workloadv1Location)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Workloadv1Instance) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetPhase

`func (o *Workloadv1Instance) GetPhase() Workloadv1InstanceInstancePhase`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *Workloadv1Instance) GetPhaseOk() (*Workloadv1InstanceInstancePhase, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *Workloadv1Instance) SetPhase(v Workloadv1InstanceInstancePhase)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *Workloadv1Instance) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetIpAddress

`func (o *Workloadv1Instance) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *Workloadv1Instance) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *Workloadv1Instance) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *Workloadv1Instance) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetExternalIpAddress

`func (o *Workloadv1Instance) GetExternalIpAddress() string`

GetExternalIpAddress returns the ExternalIpAddress field if non-nil, zero value otherwise.

### GetExternalIpAddressOk

`func (o *Workloadv1Instance) GetExternalIpAddressOk() (*string, bool)`

GetExternalIpAddressOk returns a tuple with the ExternalIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIpAddress

`func (o *Workloadv1Instance) SetExternalIpAddress(v string)`

SetExternalIpAddress sets ExternalIpAddress field to given value.

### HasExternalIpAddress

`func (o *Workloadv1Instance) HasExternalIpAddress() bool`

HasExternalIpAddress returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Workloadv1Instance) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Workloadv1Instance) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Workloadv1Instance) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Workloadv1Instance) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetStartedAt

`func (o *Workloadv1Instance) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *Workloadv1Instance) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *Workloadv1Instance) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *Workloadv1Instance) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *Workloadv1Instance) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Workloadv1Instance) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Workloadv1Instance) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Workloadv1Instance) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetNetworkInterfaces

`func (o *Workloadv1Instance) GetNetworkInterfaces() []Workloadv1NetworkInterfaceStatus`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *Workloadv1Instance) GetNetworkInterfacesOk() (*[]Workloadv1NetworkInterfaceStatus, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *Workloadv1Instance) SetNetworkInterfaces(v []Workloadv1NetworkInterfaceStatus)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *Workloadv1Instance) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.

### GetResources

`func (o *Workloadv1Instance) GetResources() V1ResourceRequirements`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *Workloadv1Instance) GetResourcesOk() (*V1ResourceRequirements, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *Workloadv1Instance) SetResources(v V1ResourceRequirements)`

SetResources sets Resources field to given value.

### HasResources

`func (o *Workloadv1Instance) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetContainers

`func (o *Workloadv1Instance) GetContainers() map[string]V1ContainerSpec`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *Workloadv1Instance) GetContainersOk() (*map[string]V1ContainerSpec, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *Workloadv1Instance) SetContainers(v map[string]V1ContainerSpec)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *Workloadv1Instance) HasContainers() bool`

HasContainers returns a boolean if a field has been set.

### GetContainerStatuses

`func (o *Workloadv1Instance) GetContainerStatuses() []V1ContainerStatus`

GetContainerStatuses returns the ContainerStatuses field if non-nil, zero value otherwise.

### GetContainerStatusesOk

`func (o *Workloadv1Instance) GetContainerStatusesOk() (*[]V1ContainerStatus, bool)`

GetContainerStatusesOk returns a tuple with the ContainerStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerStatuses

`func (o *Workloadv1Instance) SetContainerStatuses(v []V1ContainerStatus)`

SetContainerStatuses sets ContainerStatuses field to given value.

### HasContainerStatuses

`func (o *Workloadv1Instance) HasContainerStatuses() bool`

HasContainerStatuses returns a boolean if a field has been set.

### GetVirtualMachines

`func (o *Workloadv1Instance) GetVirtualMachines() map[string]V1VirtualMachineSpec`

GetVirtualMachines returns the VirtualMachines field if non-nil, zero value otherwise.

### GetVirtualMachinesOk

`func (o *Workloadv1Instance) GetVirtualMachinesOk() (*map[string]V1VirtualMachineSpec, bool)`

GetVirtualMachinesOk returns a tuple with the VirtualMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachines

`func (o *Workloadv1Instance) SetVirtualMachines(v map[string]V1VirtualMachineSpec)`

SetVirtualMachines sets VirtualMachines field to given value.

### HasVirtualMachines

`func (o *Workloadv1Instance) HasVirtualMachines() bool`

HasVirtualMachines returns a boolean if a field has been set.

### GetVirtualMachineStatuses

`func (o *Workloadv1Instance) GetVirtualMachineStatuses() []V1VirtualMachineStatus`

GetVirtualMachineStatuses returns the VirtualMachineStatuses field if non-nil, zero value otherwise.

### GetVirtualMachineStatusesOk

`func (o *Workloadv1Instance) GetVirtualMachineStatusesOk() (*[]V1VirtualMachineStatus, bool)`

GetVirtualMachineStatusesOk returns a tuple with the VirtualMachineStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachineStatuses

`func (o *Workloadv1Instance) SetVirtualMachineStatuses(v []V1VirtualMachineStatus)`

SetVirtualMachineStatuses sets VirtualMachineStatuses field to given value.

### HasVirtualMachineStatuses

`func (o *Workloadv1Instance) HasVirtualMachineStatuses() bool`

HasVirtualMachineStatuses returns a boolean if a field has been set.

### GetReason

`func (o *Workloadv1Instance) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Workloadv1Instance) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Workloadv1Instance) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *Workloadv1Instance) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetMessage

`func (o *Workloadv1Instance) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Workloadv1Instance) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Workloadv1Instance) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Workloadv1Instance) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetScheduledAt

`func (o *Workloadv1Instance) GetScheduledAt() time.Time`

GetScheduledAt returns the ScheduledAt field if non-nil, zero value otherwise.

### GetScheduledAtOk

`func (o *Workloadv1Instance) GetScheduledAtOk() (*time.Time, bool)`

GetScheduledAtOk returns a tuple with the ScheduledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAt

`func (o *Workloadv1Instance) SetScheduledAt(v time.Time)`

SetScheduledAt sets ScheduledAt field to given value.

### HasScheduledAt

`func (o *Workloadv1Instance) HasScheduledAt() bool`

HasScheduledAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


