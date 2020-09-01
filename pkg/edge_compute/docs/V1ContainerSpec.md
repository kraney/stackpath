# V1ContainerSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | Pointer to **string** | The location of a Docker image to run as a container | [optional] 
**Command** | Pointer to **[]string** | The commands that start a container | [optional] 
**Env** | Pointer to [**map[string]V1EnvironmentVariable**](v1EnvironmentVariable.md) | A string to environment variable key/value pair | [optional] 
**Ports** | Pointer to [**map[string]V1InstancePort**](v1InstancePort.md) | A string to network port key/value pair | [optional] 
**LivenessProbe** | Pointer to [**V1Probe**](v1Probe.md) |  | [optional] 
**ReadinessProbe** | Pointer to [**V1Probe**](v1Probe.md) |  | [optional] 
**Resources** | Pointer to [**V1ResourceRequirements**](v1ResourceRequirements.md) |  | [optional] 
**VolumeMounts** | Pointer to [**[]V1InstanceVolumeMount**](v1InstanceVolumeMount.md) | Volumes to mount in the container | [optional] 

## Methods

### NewV1ContainerSpec

`func NewV1ContainerSpec() *V1ContainerSpec`

NewV1ContainerSpec instantiates a new V1ContainerSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ContainerSpecWithDefaults

`func NewV1ContainerSpecWithDefaults() *V1ContainerSpec`

NewV1ContainerSpecWithDefaults instantiates a new V1ContainerSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImage

`func (o *V1ContainerSpec) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *V1ContainerSpec) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *V1ContainerSpec) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *V1ContainerSpec) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetCommand

`func (o *V1ContainerSpec) GetCommand() []string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *V1ContainerSpec) GetCommandOk() (*[]string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *V1ContainerSpec) SetCommand(v []string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *V1ContainerSpec) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetEnv

`func (o *V1ContainerSpec) GetEnv() map[string]V1EnvironmentVariable`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *V1ContainerSpec) GetEnvOk() (*map[string]V1EnvironmentVariable, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *V1ContainerSpec) SetEnv(v map[string]V1EnvironmentVariable)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *V1ContainerSpec) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetPorts

`func (o *V1ContainerSpec) GetPorts() map[string]V1InstancePort`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *V1ContainerSpec) GetPortsOk() (*map[string]V1InstancePort, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *V1ContainerSpec) SetPorts(v map[string]V1InstancePort)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *V1ContainerSpec) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetLivenessProbe

`func (o *V1ContainerSpec) GetLivenessProbe() V1Probe`

GetLivenessProbe returns the LivenessProbe field if non-nil, zero value otherwise.

### GetLivenessProbeOk

`func (o *V1ContainerSpec) GetLivenessProbeOk() (*V1Probe, bool)`

GetLivenessProbeOk returns a tuple with the LivenessProbe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivenessProbe

`func (o *V1ContainerSpec) SetLivenessProbe(v V1Probe)`

SetLivenessProbe sets LivenessProbe field to given value.

### HasLivenessProbe

`func (o *V1ContainerSpec) HasLivenessProbe() bool`

HasLivenessProbe returns a boolean if a field has been set.

### GetReadinessProbe

`func (o *V1ContainerSpec) GetReadinessProbe() V1Probe`

GetReadinessProbe returns the ReadinessProbe field if non-nil, zero value otherwise.

### GetReadinessProbeOk

`func (o *V1ContainerSpec) GetReadinessProbeOk() (*V1Probe, bool)`

GetReadinessProbeOk returns a tuple with the ReadinessProbe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessProbe

`func (o *V1ContainerSpec) SetReadinessProbe(v V1Probe)`

SetReadinessProbe sets ReadinessProbe field to given value.

### HasReadinessProbe

`func (o *V1ContainerSpec) HasReadinessProbe() bool`

HasReadinessProbe returns a boolean if a field has been set.

### GetResources

`func (o *V1ContainerSpec) GetResources() V1ResourceRequirements`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *V1ContainerSpec) GetResourcesOk() (*V1ResourceRequirements, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *V1ContainerSpec) SetResources(v V1ResourceRequirements)`

SetResources sets Resources field to given value.

### HasResources

`func (o *V1ContainerSpec) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetVolumeMounts

`func (o *V1ContainerSpec) GetVolumeMounts() []V1InstanceVolumeMount`

GetVolumeMounts returns the VolumeMounts field if non-nil, zero value otherwise.

### GetVolumeMountsOk

`func (o *V1ContainerSpec) GetVolumeMountsOk() (*[]V1InstanceVolumeMount, bool)`

GetVolumeMountsOk returns a tuple with the VolumeMounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeMounts

`func (o *V1ContainerSpec) SetVolumeMounts(v []V1InstanceVolumeMount)`

SetVolumeMounts sets VolumeMounts field to given value.

### HasVolumeMounts

`func (o *V1ContainerSpec) HasVolumeMounts() bool`

HasVolumeMounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


