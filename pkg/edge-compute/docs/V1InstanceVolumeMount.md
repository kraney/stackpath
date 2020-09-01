# V1InstanceVolumeMount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slug** | Pointer to **string** | The slug of the volume claim to mount | [optional] 
**MountPath** | Pointer to **string** | The path in an instance to mount a volume | [optional] 

## Methods

### NewV1InstanceVolumeMount

`func NewV1InstanceVolumeMount() *V1InstanceVolumeMount`

NewV1InstanceVolumeMount instantiates a new V1InstanceVolumeMount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1InstanceVolumeMountWithDefaults

`func NewV1InstanceVolumeMountWithDefaults() *V1InstanceVolumeMount`

NewV1InstanceVolumeMountWithDefaults instantiates a new V1InstanceVolumeMount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlug

`func (o *V1InstanceVolumeMount) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *V1InstanceVolumeMount) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *V1InstanceVolumeMount) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *V1InstanceVolumeMount) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetMountPath

`func (o *V1InstanceVolumeMount) GetMountPath() string`

GetMountPath returns the MountPath field if non-nil, zero value otherwise.

### GetMountPathOk

`func (o *V1InstanceVolumeMount) GetMountPathOk() (*string, bool)`

GetMountPathOk returns a tuple with the MountPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPath

`func (o *V1InstanceVolumeMount) SetMountPath(v string)`

SetMountPath sets MountPath field to given value.

### HasMountPath

`func (o *V1InstanceVolumeMount) HasMountPath() bool`

HasMountPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


