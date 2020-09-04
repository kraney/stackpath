# V1VolumeClaim

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StackId** | Pointer to **string** | The ID of the stack that a volume claim belongs to | [optional] [readonly] 
**Id** | Pointer to **string** | A volume claim&#39;s unique identifier | [optional] [readonly] 
**Name** | Pointer to **string** | A volume claim&#39;s name as it appears in the StackPath portal | [optional] 
**Slug** | Pointer to **string** | A volume claim&#39;s programmatic name  Volume claim slugs are used to programatically label a claim | [optional] 
**Metadata** | Pointer to [**V1Metadata**](v1Metadata.md) |  | [optional] 
**Spec** | Pointer to [**V1VolumeClaimSpec**](v1VolumeClaimSpec.md) |  | [optional] 
**Phase** | Pointer to [**VolumeClaimVolumeClaimPhase**](VolumeClaimVolumeClaimPhase.md) |  | [optional] [default to "VOLUME_CLAIM_PHASE_UNSPECIFIED"]

## Methods

### NewV1VolumeClaim

`func NewV1VolumeClaim() *V1VolumeClaim`

NewV1VolumeClaim instantiates a new V1VolumeClaim object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1VolumeClaimWithDefaults

`func NewV1VolumeClaimWithDefaults() *V1VolumeClaim`

NewV1VolumeClaimWithDefaults instantiates a new V1VolumeClaim object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStackId

`func (o *V1VolumeClaim) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *V1VolumeClaim) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *V1VolumeClaim) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *V1VolumeClaim) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetId

`func (o *V1VolumeClaim) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1VolumeClaim) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1VolumeClaim) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *V1VolumeClaim) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *V1VolumeClaim) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1VolumeClaim) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1VolumeClaim) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1VolumeClaim) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *V1VolumeClaim) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *V1VolumeClaim) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *V1VolumeClaim) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *V1VolumeClaim) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetMetadata

`func (o *V1VolumeClaim) GetMetadata() V1Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V1VolumeClaim) GetMetadataOk() (*V1Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V1VolumeClaim) SetMetadata(v V1Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V1VolumeClaim) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *V1VolumeClaim) GetSpec() V1VolumeClaimSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *V1VolumeClaim) GetSpecOk() (*V1VolumeClaimSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *V1VolumeClaim) SetSpec(v V1VolumeClaimSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *V1VolumeClaim) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetPhase

`func (o *V1VolumeClaim) GetPhase() VolumeClaimVolumeClaimPhase`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *V1VolumeClaim) GetPhaseOk() (*VolumeClaimVolumeClaimPhase, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *V1VolumeClaim) SetPhase(v VolumeClaimVolumeClaimPhase)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *V1VolumeClaim) HasPhase() bool`

HasPhase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


