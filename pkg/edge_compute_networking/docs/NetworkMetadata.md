# NetworkMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotations** | Pointer to **map[string]string** | A string to string key/value pair | [optional] 
**Labels** | Pointer to **map[string]string** | A string to string key/value pair | [optional] 
**CreatedAt** | Pointer to [**NullableTime**](time.Time.md) | The date that a metadata entry was created | [optional] [readonly] 
**UpdatedAt** | Pointer to [**NullableTime**](time.Time.md) | The date that a metadata entry was last updated | [optional] [readonly] 
**DeleteRequestedAt** | Pointer to [**NullableTime**](time.Time.md) | The date that a network policy was requested for deletion | [optional] [readonly] 
**Version** | Pointer to **string** | An entity&#39;s version number  Versions start at 1 when they are created and increment by 1 every time they are updated. | [optional] 

## Methods

### NewNetworkMetadata

`func NewNetworkMetadata() *NetworkMetadata`

NewNetworkMetadata instantiates a new NetworkMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkMetadataWithDefaults

`func NewNetworkMetadataWithDefaults() *NetworkMetadata`

NewNetworkMetadataWithDefaults instantiates a new NetworkMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotations

`func (o *NetworkMetadata) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *NetworkMetadata) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *NetworkMetadata) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *NetworkMetadata) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetLabels

`func (o *NetworkMetadata) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *NetworkMetadata) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *NetworkMetadata) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *NetworkMetadata) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetCreatedAt

`func (o *NetworkMetadata) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NetworkMetadata) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NetworkMetadata) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NetworkMetadata) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *NetworkMetadata) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *NetworkMetadata) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *NetworkMetadata) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NetworkMetadata) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NetworkMetadata) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *NetworkMetadata) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *NetworkMetadata) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *NetworkMetadata) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetDeleteRequestedAt

`func (o *NetworkMetadata) GetDeleteRequestedAt() time.Time`

GetDeleteRequestedAt returns the DeleteRequestedAt field if non-nil, zero value otherwise.

### GetDeleteRequestedAtOk

`func (o *NetworkMetadata) GetDeleteRequestedAtOk() (*time.Time, bool)`

GetDeleteRequestedAtOk returns a tuple with the DeleteRequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteRequestedAt

`func (o *NetworkMetadata) SetDeleteRequestedAt(v time.Time)`

SetDeleteRequestedAt sets DeleteRequestedAt field to given value.

### HasDeleteRequestedAt

`func (o *NetworkMetadata) HasDeleteRequestedAt() bool`

HasDeleteRequestedAt returns a boolean if a field has been set.

### SetDeleteRequestedAtNil

`func (o *NetworkMetadata) SetDeleteRequestedAtNil(b bool)`

 SetDeleteRequestedAtNil sets the value for DeleteRequestedAt to be an explicit nil

### UnsetDeleteRequestedAt
`func (o *NetworkMetadata) UnsetDeleteRequestedAt()`

UnsetDeleteRequestedAt ensures that no value is present for DeleteRequestedAt, not even an explicit nil
### GetVersion

`func (o *NetworkMetadata) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NetworkMetadata) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NetworkMetadata) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *NetworkMetadata) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


