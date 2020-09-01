# V1WorkloadSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkInterfaces** | Pointer to [**[]V1NetworkInterface**](v1NetworkInterface.md) | Network interfaces to bind to the workload&#39;s instances | [optional] 
**Containers** | Pointer to [**map[string]V1ContainerSpec**](v1ContainerSpec.md) | A string to container configuration key/value pair | [optional] 
**VirtualMachines** | Pointer to [**map[string]V1VirtualMachineSpec**](v1VirtualMachineSpec.md) | A string to virtual machine configuration key/value pair | [optional] 
**VolumeClaimTemplates** | Pointer to [**[]V1VolumeClaim**](v1VolumeClaim.md) | A list of claims that instances may reference  The slug of the claim will be used in combination with the name of the instance to create a stable identifier. The slug should be used in the volume mount specifications for containers and VMs. | [optional] 
**ImagePullCredentials** | Pointer to [**[]V1ImagePullCredential**](v1ImagePullCredential.md) | The credentials needed to pull a container image | [optional] 

## Methods

### NewV1WorkloadSpec

`func NewV1WorkloadSpec() *V1WorkloadSpec`

NewV1WorkloadSpec instantiates a new V1WorkloadSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1WorkloadSpecWithDefaults

`func NewV1WorkloadSpecWithDefaults() *V1WorkloadSpec`

NewV1WorkloadSpecWithDefaults instantiates a new V1WorkloadSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkInterfaces

`func (o *V1WorkloadSpec) GetNetworkInterfaces() []V1NetworkInterface`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *V1WorkloadSpec) GetNetworkInterfacesOk() (*[]V1NetworkInterface, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *V1WorkloadSpec) SetNetworkInterfaces(v []V1NetworkInterface)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *V1WorkloadSpec) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.

### GetContainers

`func (o *V1WorkloadSpec) GetContainers() map[string]V1ContainerSpec`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *V1WorkloadSpec) GetContainersOk() (*map[string]V1ContainerSpec, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *V1WorkloadSpec) SetContainers(v map[string]V1ContainerSpec)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *V1WorkloadSpec) HasContainers() bool`

HasContainers returns a boolean if a field has been set.

### GetVirtualMachines

`func (o *V1WorkloadSpec) GetVirtualMachines() map[string]V1VirtualMachineSpec`

GetVirtualMachines returns the VirtualMachines field if non-nil, zero value otherwise.

### GetVirtualMachinesOk

`func (o *V1WorkloadSpec) GetVirtualMachinesOk() (*map[string]V1VirtualMachineSpec, bool)`

GetVirtualMachinesOk returns a tuple with the VirtualMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachines

`func (o *V1WorkloadSpec) SetVirtualMachines(v map[string]V1VirtualMachineSpec)`

SetVirtualMachines sets VirtualMachines field to given value.

### HasVirtualMachines

`func (o *V1WorkloadSpec) HasVirtualMachines() bool`

HasVirtualMachines returns a boolean if a field has been set.

### GetVolumeClaimTemplates

`func (o *V1WorkloadSpec) GetVolumeClaimTemplates() []V1VolumeClaim`

GetVolumeClaimTemplates returns the VolumeClaimTemplates field if non-nil, zero value otherwise.

### GetVolumeClaimTemplatesOk

`func (o *V1WorkloadSpec) GetVolumeClaimTemplatesOk() (*[]V1VolumeClaim, bool)`

GetVolumeClaimTemplatesOk returns a tuple with the VolumeClaimTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeClaimTemplates

`func (o *V1WorkloadSpec) SetVolumeClaimTemplates(v []V1VolumeClaim)`

SetVolumeClaimTemplates sets VolumeClaimTemplates field to given value.

### HasVolumeClaimTemplates

`func (o *V1WorkloadSpec) HasVolumeClaimTemplates() bool`

HasVolumeClaimTemplates returns a boolean if a field has been set.

### GetImagePullCredentials

`func (o *V1WorkloadSpec) GetImagePullCredentials() []V1ImagePullCredential`

GetImagePullCredentials returns the ImagePullCredentials field if non-nil, zero value otherwise.

### GetImagePullCredentialsOk

`func (o *V1WorkloadSpec) GetImagePullCredentialsOk() (*[]V1ImagePullCredential, bool)`

GetImagePullCredentialsOk returns a tuple with the ImagePullCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagePullCredentials

`func (o *V1WorkloadSpec) SetImagePullCredentials(v []V1ImagePullCredential)`

SetImagePullCredentials sets ImagePullCredentials field to given value.

### HasImagePullCredentials

`func (o *V1WorkloadSpec) HasImagePullCredentials() bool`

HasImagePullCredentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


