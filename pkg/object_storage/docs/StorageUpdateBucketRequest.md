# StorageUpdateBucketRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Visibility** | Pointer to [**StorageBucketVisibility**](storageBucketVisibility.md) |  | [optional] [default to "PRIVATE"]

## Methods

### NewStorageUpdateBucketRequest

`func NewStorageUpdateBucketRequest() *StorageUpdateBucketRequest`

NewStorageUpdateBucketRequest instantiates a new StorageUpdateBucketRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageUpdateBucketRequestWithDefaults

`func NewStorageUpdateBucketRequestWithDefaults() *StorageUpdateBucketRequest`

NewStorageUpdateBucketRequestWithDefaults instantiates a new StorageUpdateBucketRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVisibility

`func (o *StorageUpdateBucketRequest) GetVisibility() StorageBucketVisibility`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *StorageUpdateBucketRequest) GetVisibilityOk() (*StorageBucketVisibility, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *StorageUpdateBucketRequest) SetVisibility(v StorageBucketVisibility)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *StorageUpdateBucketRequest) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


