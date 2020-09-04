# V1Image

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StackId** | Pointer to **string** | The ID of the stack that an image belongs to | [optional] [readonly] 
**Id** | Pointer to **string** | An image&#39;s unique identifier | [optional] [readonly] 
**Family** | Pointer to **string** | An image&#39;s family  Families are logical groupings of images | [optional] 
**Tag** | Pointer to **string** | The image&#39;s tag  Image tags are akin to versions | [optional] 
**Metadata** | Pointer to [**V1ImageMetadata**](v1ImageMetadata.md) |  | [optional] 
**Description** | Pointer to **string** | An image&#39;s description  This is optional and may not be more than 1,000 characters | [optional] 
**Status** | Pointer to [**V1ImageStatus**](v1ImageStatus.md) |  | [optional] [default to "IMAGE_STATUS_UNKNOWN"]
**ImageSize** | Pointer to **string** | An image&#39;s archive size in bytes  This is only available on ready images | [optional] [readonly] 
**VolumeSize** | Pointer to **string** | An image&#39;s volume size in bytes  This is only available on ready images | [optional] [readonly] 
**Deprecation** | Pointer to [**V1ImageDeprecation**](v1ImageDeprecation.md) |  | [optional] 
**Conditions** | Pointer to [**[]V1ImageCondition**](v1ImageCondition.md) | Details about an image&#39;s status | [optional] [readonly] 

## Methods

### NewV1Image

`func NewV1Image() *V1Image`

NewV1Image instantiates a new V1Image object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ImageWithDefaults

`func NewV1ImageWithDefaults() *V1Image`

NewV1ImageWithDefaults instantiates a new V1Image object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStackId

`func (o *V1Image) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *V1Image) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *V1Image) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *V1Image) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetId

`func (o *V1Image) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1Image) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1Image) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *V1Image) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFamily

`func (o *V1Image) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *V1Image) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *V1Image) SetFamily(v string)`

SetFamily sets Family field to given value.

### HasFamily

`func (o *V1Image) HasFamily() bool`

HasFamily returns a boolean if a field has been set.

### GetTag

`func (o *V1Image) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *V1Image) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *V1Image) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *V1Image) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetMetadata

`func (o *V1Image) GetMetadata() V1ImageMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V1Image) GetMetadataOk() (*V1ImageMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V1Image) SetMetadata(v V1ImageMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V1Image) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDescription

`func (o *V1Image) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V1Image) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V1Image) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V1Image) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *V1Image) GetStatus() V1ImageStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1Image) GetStatusOk() (*V1ImageStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1Image) SetStatus(v V1ImageStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V1Image) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetImageSize

`func (o *V1Image) GetImageSize() string`

GetImageSize returns the ImageSize field if non-nil, zero value otherwise.

### GetImageSizeOk

`func (o *V1Image) GetImageSizeOk() (*string, bool)`

GetImageSizeOk returns a tuple with the ImageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageSize

`func (o *V1Image) SetImageSize(v string)`

SetImageSize sets ImageSize field to given value.

### HasImageSize

`func (o *V1Image) HasImageSize() bool`

HasImageSize returns a boolean if a field has been set.

### GetVolumeSize

`func (o *V1Image) GetVolumeSize() string`

GetVolumeSize returns the VolumeSize field if non-nil, zero value otherwise.

### GetVolumeSizeOk

`func (o *V1Image) GetVolumeSizeOk() (*string, bool)`

GetVolumeSizeOk returns a tuple with the VolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeSize

`func (o *V1Image) SetVolumeSize(v string)`

SetVolumeSize sets VolumeSize field to given value.

### HasVolumeSize

`func (o *V1Image) HasVolumeSize() bool`

HasVolumeSize returns a boolean if a field has been set.

### GetDeprecation

`func (o *V1Image) GetDeprecation() V1ImageDeprecation`

GetDeprecation returns the Deprecation field if non-nil, zero value otherwise.

### GetDeprecationOk

`func (o *V1Image) GetDeprecationOk() (*V1ImageDeprecation, bool)`

GetDeprecationOk returns a tuple with the Deprecation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecation

`func (o *V1Image) SetDeprecation(v V1ImageDeprecation)`

SetDeprecation sets Deprecation field to given value.

### HasDeprecation

`func (o *V1Image) HasDeprecation() bool`

HasDeprecation returns a boolean if a field has been set.

### GetConditions

`func (o *V1Image) GetConditions() []V1ImageCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *V1Image) GetConditionsOk() (*[]V1ImageCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *V1Image) SetConditions(v []V1ImageCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *V1Image) HasConditions() bool`

HasConditions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


