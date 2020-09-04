# V1ImageDeprecation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeprecationDate** | Pointer to [**time.Time**](time.Time.md) | The date when an image is considered deprecated  Deprecated images may be used in new virtual machine based workloads, but are no longer considered for the \&quot;default\&quot; tag, nor are they returned in image lists by default. | [optional] 
**ObsoletionDate** | Pointer to [**time.Time**](time.Time.md) | The date when an image is considered obsolete  Obsolete images may not be used in new virtual machine based workloads. If an obsoletion date is present then the deprecation date must also be present with a same or earlier date. | [optional] 
**DeletionDate** | Pointer to [**time.Time**](time.Time.md) | The date when an image is deleted  Deletion dates cannot be before an image&#39;s deprecation or obsoletion dates. | [optional] 

## Methods

### NewV1ImageDeprecation

`func NewV1ImageDeprecation() *V1ImageDeprecation`

NewV1ImageDeprecation instantiates a new V1ImageDeprecation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ImageDeprecationWithDefaults

`func NewV1ImageDeprecationWithDefaults() *V1ImageDeprecation`

NewV1ImageDeprecationWithDefaults instantiates a new V1ImageDeprecation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeprecationDate

`func (o *V1ImageDeprecation) GetDeprecationDate() time.Time`

GetDeprecationDate returns the DeprecationDate field if non-nil, zero value otherwise.

### GetDeprecationDateOk

`func (o *V1ImageDeprecation) GetDeprecationDateOk() (*time.Time, bool)`

GetDeprecationDateOk returns a tuple with the DeprecationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecationDate

`func (o *V1ImageDeprecation) SetDeprecationDate(v time.Time)`

SetDeprecationDate sets DeprecationDate field to given value.

### HasDeprecationDate

`func (o *V1ImageDeprecation) HasDeprecationDate() bool`

HasDeprecationDate returns a boolean if a field has been set.

### GetObsoletionDate

`func (o *V1ImageDeprecation) GetObsoletionDate() time.Time`

GetObsoletionDate returns the ObsoletionDate field if non-nil, zero value otherwise.

### GetObsoletionDateOk

`func (o *V1ImageDeprecation) GetObsoletionDateOk() (*time.Time, bool)`

GetObsoletionDateOk returns a tuple with the ObsoletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObsoletionDate

`func (o *V1ImageDeprecation) SetObsoletionDate(v time.Time)`

SetObsoletionDate sets ObsoletionDate field to given value.

### HasObsoletionDate

`func (o *V1ImageDeprecation) HasObsoletionDate() bool`

HasObsoletionDate returns a boolean if a field has been set.

### GetDeletionDate

`func (o *V1ImageDeprecation) GetDeletionDate() time.Time`

GetDeletionDate returns the DeletionDate field if non-nil, zero value otherwise.

### GetDeletionDateOk

`func (o *V1ImageDeprecation) GetDeletionDateOk() (*time.Time, bool)`

GetDeletionDateOk returns a tuple with the DeletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionDate

`func (o *V1ImageDeprecation) SetDeletionDate(v time.Time)`

SetDeletionDate sets DeletionDate field to given value.

### HasDeletionDate

`func (o *V1ImageDeprecation) HasDeletionDate() bool`

HasDeletionDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


