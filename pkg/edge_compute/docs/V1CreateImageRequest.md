# V1CreateImageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | Pointer to [**V1Image**](v1Image.md) |  | [optional] 
**InstanceVolumeSource** | Pointer to [**V1ImageSourceInstanceVolume**](v1ImageSourceInstanceVolume.md) |  | [optional] 

## Methods

### NewV1CreateImageRequest

`func NewV1CreateImageRequest() *V1CreateImageRequest`

NewV1CreateImageRequest instantiates a new V1CreateImageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1CreateImageRequestWithDefaults

`func NewV1CreateImageRequestWithDefaults() *V1CreateImageRequest`

NewV1CreateImageRequestWithDefaults instantiates a new V1CreateImageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImage

`func (o *V1CreateImageRequest) GetImage() V1Image`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *V1CreateImageRequest) GetImageOk() (*V1Image, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *V1CreateImageRequest) SetImage(v V1Image)`

SetImage sets Image field to given value.

### HasImage

`func (o *V1CreateImageRequest) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetInstanceVolumeSource

`func (o *V1CreateImageRequest) GetInstanceVolumeSource() V1ImageSourceInstanceVolume`

GetInstanceVolumeSource returns the InstanceVolumeSource field if non-nil, zero value otherwise.

### GetInstanceVolumeSourceOk

`func (o *V1CreateImageRequest) GetInstanceVolumeSourceOk() (*V1ImageSourceInstanceVolume, bool)`

GetInstanceVolumeSourceOk returns a tuple with the InstanceVolumeSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceVolumeSource

`func (o *V1CreateImageRequest) SetInstanceVolumeSource(v V1ImageSourceInstanceVolume)`

SetInstanceVolumeSource sets InstanceVolumeSource field to given value.

### HasInstanceVolumeSource

`func (o *V1CreateImageRequest) HasInstanceVolumeSource() bool`

HasInstanceVolumeSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


