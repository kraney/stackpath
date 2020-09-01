# V1NetworkPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StackId** | Pointer to **string** | The ID of the stack that a network policy belongs to | [optional] [readonly] 
**Id** | Pointer to **string** | A network policy&#39;s unique identifier | [optional] [readonly] 
**Name** | Pointer to **string** | A network policy&#39;s name as it appears in the StackPath portal | [optional] 
**Slug** | Pointer to **string** | A network policy&#39;s programmatic name  Network policy slugs are used to build its instances names | [optional] 
**Description** | Pointer to **string** | Detailed summary of what the policy does | [optional] 
**Metadata** | Pointer to [**NetworkMetadata**](networkMetadata.md) |  | [optional] 
**Spec** | Pointer to [**V1NetworkPolicySpec**](v1NetworkPolicySpec.md) |  | [optional] 

## Methods

### NewV1NetworkPolicy

`func NewV1NetworkPolicy() *V1NetworkPolicy`

NewV1NetworkPolicy instantiates a new V1NetworkPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1NetworkPolicyWithDefaults

`func NewV1NetworkPolicyWithDefaults() *V1NetworkPolicy`

NewV1NetworkPolicyWithDefaults instantiates a new V1NetworkPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStackId

`func (o *V1NetworkPolicy) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *V1NetworkPolicy) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *V1NetworkPolicy) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *V1NetworkPolicy) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetId

`func (o *V1NetworkPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1NetworkPolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1NetworkPolicy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *V1NetworkPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *V1NetworkPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1NetworkPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1NetworkPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1NetworkPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *V1NetworkPolicy) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *V1NetworkPolicy) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *V1NetworkPolicy) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *V1NetworkPolicy) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetDescription

`func (o *V1NetworkPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V1NetworkPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V1NetworkPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V1NetworkPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *V1NetworkPolicy) GetMetadata() NetworkMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V1NetworkPolicy) GetMetadataOk() (*NetworkMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V1NetworkPolicy) SetMetadata(v NetworkMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V1NetworkPolicy) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *V1NetworkPolicy) GetSpec() V1NetworkPolicySpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *V1NetworkPolicy) GetSpecOk() (*V1NetworkPolicySpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *V1NetworkPolicy) SetSpec(v V1NetworkPolicySpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *V1NetworkPolicy) HasSpec() bool`

HasSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


