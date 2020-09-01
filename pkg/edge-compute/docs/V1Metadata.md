# V1Metadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotations** | Pointer to **map[string]string** | A string to string key/value pair | [optional] 
**Labels** | Pointer to **map[string]string** | A string to string key/value pair | [optional] 
**CreatedAt** | Pointer to [**NullableTime**](time.Time.md) | The date that a metadata entry was created | [optional] [readonly] 
**UpdatedAt** | Pointer to [**NullableTime**](time.Time.md) | The date that a metadata entry was last updated | [optional] [readonly] 
**DeleteRequestedAt** | Pointer to [**NullableTime**](time.Time.md) | The date an entity was requested to be deleted | [optional] [readonly] 
**Version** | Pointer to **string** | A metadata entry&#39;s version number  Metadata versions start at 1 when they are created and increment by 1 every time they are updated. | [optional] 

## Methods

### NewV1Metadata

`func NewV1Metadata() *V1Metadata`

NewV1Metadata instantiates a new V1Metadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1MetadataWithDefaults

`func NewV1MetadataWithDefaults() *V1Metadata`

NewV1MetadataWithDefaults instantiates a new V1Metadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotations

`func (o *V1Metadata) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *V1Metadata) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *V1Metadata) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *V1Metadata) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetLabels

`func (o *V1Metadata) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *V1Metadata) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *V1Metadata) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *V1Metadata) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetCreatedAt

`func (o *V1Metadata) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *V1Metadata) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *V1Metadata) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *V1Metadata) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *V1Metadata) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *V1Metadata) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *V1Metadata) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *V1Metadata) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *V1Metadata) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *V1Metadata) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *V1Metadata) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *V1Metadata) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetDeleteRequestedAt

`func (o *V1Metadata) GetDeleteRequestedAt() time.Time`

GetDeleteRequestedAt returns the DeleteRequestedAt field if non-nil, zero value otherwise.

### GetDeleteRequestedAtOk

`func (o *V1Metadata) GetDeleteRequestedAtOk() (*time.Time, bool)`

GetDeleteRequestedAtOk returns a tuple with the DeleteRequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteRequestedAt

`func (o *V1Metadata) SetDeleteRequestedAt(v time.Time)`

SetDeleteRequestedAt sets DeleteRequestedAt field to given value.

### HasDeleteRequestedAt

`func (o *V1Metadata) HasDeleteRequestedAt() bool`

HasDeleteRequestedAt returns a boolean if a field has been set.

### SetDeleteRequestedAtNil

`func (o *V1Metadata) SetDeleteRequestedAtNil(b bool)`

 SetDeleteRequestedAtNil sets the value for DeleteRequestedAt to be an explicit nil

### UnsetDeleteRequestedAt
`func (o *V1Metadata) UnsetDeleteRequestedAt()`

UnsetDeleteRequestedAt ensures that no value is present for DeleteRequestedAt, not even an explicit nil
### GetVersion

`func (o *V1Metadata) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *V1Metadata) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *V1Metadata) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *V1Metadata) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


