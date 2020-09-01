# V1Workload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StackId** | Pointer to **string** | The ID of the stack that a workload belongs to | [optional] [readonly] 
**Id** | Pointer to **string** | A workload&#39;s unique identifier | [optional] [readonly] 
**Name** | Pointer to **string** | A workload&#39;s name as it appears in the StackPath portal | [optional] 
**Slug** | Pointer to **string** | A workload&#39;s programmatic name  Workload slugs are used to build its instances names | [optional] 
**Metadata** | Pointer to [**V1Metadata**](v1Metadata.md) |  | [optional] 
**Spec** | Pointer to [**V1WorkloadSpec**](v1WorkloadSpec.md) |  | [optional] 
**Targets** | Pointer to [**map[string]V1Target**](v1Target.md) | A string to deployment target key/value pair | [optional] 
**Status** | Pointer to [**V1WorkloadStatus**](v1WorkloadStatus.md) |  | [optional] [default to "ACTIVE"]

## Methods

### NewV1Workload

`func NewV1Workload() *V1Workload`

NewV1Workload instantiates a new V1Workload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1WorkloadWithDefaults

`func NewV1WorkloadWithDefaults() *V1Workload`

NewV1WorkloadWithDefaults instantiates a new V1Workload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStackId

`func (o *V1Workload) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *V1Workload) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *V1Workload) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *V1Workload) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetId

`func (o *V1Workload) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1Workload) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1Workload) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *V1Workload) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *V1Workload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1Workload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1Workload) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1Workload) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *V1Workload) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *V1Workload) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *V1Workload) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *V1Workload) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetMetadata

`func (o *V1Workload) GetMetadata() V1Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V1Workload) GetMetadataOk() (*V1Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V1Workload) SetMetadata(v V1Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V1Workload) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *V1Workload) GetSpec() V1WorkloadSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *V1Workload) GetSpecOk() (*V1WorkloadSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *V1Workload) SetSpec(v V1WorkloadSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *V1Workload) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetTargets

`func (o *V1Workload) GetTargets() map[string]V1Target`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *V1Workload) GetTargetsOk() (*map[string]V1Target, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *V1Workload) SetTargets(v map[string]V1Target)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *V1Workload) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetStatus

`func (o *V1Workload) GetStatus() V1WorkloadStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1Workload) GetStatusOk() (*V1WorkloadStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1Workload) SetStatus(v V1WorkloadStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V1Workload) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


