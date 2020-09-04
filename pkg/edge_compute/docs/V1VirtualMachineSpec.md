# V1VirtualMachineSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | Pointer to **string** | The image to use for the virtual machine  This is in the format of &lt;stack-slug&gt;/&lt;image-family&gt;[:&lt;image-tag&gt;]. If the image tag portion is omitted, &#39;default&#39; is assumed which is the most recently created, ready, and non-deprecated image of that slug. A set of common images is present on the &#39;stackpath-edge&#39; stack. | [optional] 
**UserData** | Pointer to **string** | Base64 encoded cloud-init compatible user-data | [optional] 
**Ports** | Pointer to [**map[string]V1InstancePort**](v1InstancePort.md) | A string to network port key/value pair | [optional] 
**LivenessProbe** | Pointer to [**V1Probe**](v1Probe.md) |  | [optional] 
**ReadinessProbe** | Pointer to [**V1Probe**](v1Probe.md) |  | [optional] 
**Resources** | Pointer to [**V1ResourceRequirements**](v1ResourceRequirements.md) |  | [optional] 
**VolumeMounts** | Pointer to [**[]V1InstanceVolumeMount**](v1InstanceVolumeMount.md) | Volumes to mount in the virtual machine | [optional] 

## Methods

### NewV1VirtualMachineSpec

`func NewV1VirtualMachineSpec() *V1VirtualMachineSpec`

NewV1VirtualMachineSpec instantiates a new V1VirtualMachineSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1VirtualMachineSpecWithDefaults

`func NewV1VirtualMachineSpecWithDefaults() *V1VirtualMachineSpec`

NewV1VirtualMachineSpecWithDefaults instantiates a new V1VirtualMachineSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImage

`func (o *V1VirtualMachineSpec) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *V1VirtualMachineSpec) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *V1VirtualMachineSpec) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *V1VirtualMachineSpec) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetUserData

`func (o *V1VirtualMachineSpec) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *V1VirtualMachineSpec) GetUserDataOk() (*string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *V1VirtualMachineSpec) SetUserData(v string)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *V1VirtualMachineSpec) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### GetPorts

`func (o *V1VirtualMachineSpec) GetPorts() map[string]V1InstancePort`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *V1VirtualMachineSpec) GetPortsOk() (*map[string]V1InstancePort, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *V1VirtualMachineSpec) SetPorts(v map[string]V1InstancePort)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *V1VirtualMachineSpec) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetLivenessProbe

`func (o *V1VirtualMachineSpec) GetLivenessProbe() V1Probe`

GetLivenessProbe returns the LivenessProbe field if non-nil, zero value otherwise.

### GetLivenessProbeOk

`func (o *V1VirtualMachineSpec) GetLivenessProbeOk() (*V1Probe, bool)`

GetLivenessProbeOk returns a tuple with the LivenessProbe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivenessProbe

`func (o *V1VirtualMachineSpec) SetLivenessProbe(v V1Probe)`

SetLivenessProbe sets LivenessProbe field to given value.

### HasLivenessProbe

`func (o *V1VirtualMachineSpec) HasLivenessProbe() bool`

HasLivenessProbe returns a boolean if a field has been set.

### GetReadinessProbe

`func (o *V1VirtualMachineSpec) GetReadinessProbe() V1Probe`

GetReadinessProbe returns the ReadinessProbe field if non-nil, zero value otherwise.

### GetReadinessProbeOk

`func (o *V1VirtualMachineSpec) GetReadinessProbeOk() (*V1Probe, bool)`

GetReadinessProbeOk returns a tuple with the ReadinessProbe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessProbe

`func (o *V1VirtualMachineSpec) SetReadinessProbe(v V1Probe)`

SetReadinessProbe sets ReadinessProbe field to given value.

### HasReadinessProbe

`func (o *V1VirtualMachineSpec) HasReadinessProbe() bool`

HasReadinessProbe returns a boolean if a field has been set.

### GetResources

`func (o *V1VirtualMachineSpec) GetResources() V1ResourceRequirements`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *V1VirtualMachineSpec) GetResourcesOk() (*V1ResourceRequirements, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *V1VirtualMachineSpec) SetResources(v V1ResourceRequirements)`

SetResources sets Resources field to given value.

### HasResources

`func (o *V1VirtualMachineSpec) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetVolumeMounts

`func (o *V1VirtualMachineSpec) GetVolumeMounts() []V1InstanceVolumeMount`

GetVolumeMounts returns the VolumeMounts field if non-nil, zero value otherwise.

### GetVolumeMountsOk

`func (o *V1VirtualMachineSpec) GetVolumeMountsOk() (*[]V1InstanceVolumeMount, bool)`

GetVolumeMountsOk returns a tuple with the VolumeMounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeMounts

`func (o *V1VirtualMachineSpec) SetVolumeMounts(v []V1InstanceVolumeMount)`

SetVolumeMounts sets VolumeMounts field to given value.

### HasVolumeMounts

`func (o *V1VirtualMachineSpec) HasVolumeMounts() bool`

HasVolumeMounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


