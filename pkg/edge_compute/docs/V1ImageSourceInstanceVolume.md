# V1ImageSourceInstanceVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkloadId** | Pointer to **string** | The ID or slug of the workload containing the instance to reference | [optional] 
**InstanceName** | Pointer to **string** | An instance&#39;s name to reference a volume on | [optional] 
**VolumeClaimSlug** | Pointer to **string** | A reference to the volume to create an image from or unset for the root volume | [optional] 

## Methods

### NewV1ImageSourceInstanceVolume

`func NewV1ImageSourceInstanceVolume() *V1ImageSourceInstanceVolume`

NewV1ImageSourceInstanceVolume instantiates a new V1ImageSourceInstanceVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ImageSourceInstanceVolumeWithDefaults

`func NewV1ImageSourceInstanceVolumeWithDefaults() *V1ImageSourceInstanceVolume`

NewV1ImageSourceInstanceVolumeWithDefaults instantiates a new V1ImageSourceInstanceVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkloadId

`func (o *V1ImageSourceInstanceVolume) GetWorkloadId() string`

GetWorkloadId returns the WorkloadId field if non-nil, zero value otherwise.

### GetWorkloadIdOk

`func (o *V1ImageSourceInstanceVolume) GetWorkloadIdOk() (*string, bool)`

GetWorkloadIdOk returns a tuple with the WorkloadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadId

`func (o *V1ImageSourceInstanceVolume) SetWorkloadId(v string)`

SetWorkloadId sets WorkloadId field to given value.

### HasWorkloadId

`func (o *V1ImageSourceInstanceVolume) HasWorkloadId() bool`

HasWorkloadId returns a boolean if a field has been set.

### GetInstanceName

`func (o *V1ImageSourceInstanceVolume) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *V1ImageSourceInstanceVolume) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *V1ImageSourceInstanceVolume) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.

### HasInstanceName

`func (o *V1ImageSourceInstanceVolume) HasInstanceName() bool`

HasInstanceName returns a boolean if a field has been set.

### GetVolumeClaimSlug

`func (o *V1ImageSourceInstanceVolume) GetVolumeClaimSlug() string`

GetVolumeClaimSlug returns the VolumeClaimSlug field if non-nil, zero value otherwise.

### GetVolumeClaimSlugOk

`func (o *V1ImageSourceInstanceVolume) GetVolumeClaimSlugOk() (*string, bool)`

GetVolumeClaimSlugOk returns a tuple with the VolumeClaimSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeClaimSlug

`func (o *V1ImageSourceInstanceVolume) SetVolumeClaimSlug(v string)`

SetVolumeClaimSlug sets VolumeClaimSlug field to given value.

### HasVolumeClaimSlug

`func (o *V1ImageSourceInstanceVolume) HasVolumeClaimSlug() bool`

HasVolumeClaimSlug returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


