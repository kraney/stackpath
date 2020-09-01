# StorageCreateBucketResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | Pointer to [**StorageBucket**](storageBucket.md) |  | [optional] 

## Methods

### NewStorageCreateBucketResponse

`func NewStorageCreateBucketResponse() *StorageCreateBucketResponse`

NewStorageCreateBucketResponse instantiates a new StorageCreateBucketResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageCreateBucketResponseWithDefaults

`func NewStorageCreateBucketResponseWithDefaults() *StorageCreateBucketResponse`

NewStorageCreateBucketResponseWithDefaults instantiates a new StorageCreateBucketResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *StorageCreateBucketResponse) GetBucket() StorageBucket`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *StorageCreateBucketResponse) GetBucketOk() (*StorageBucket, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *StorageCreateBucketResponse) SetBucket(v StorageBucket)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *StorageCreateBucketResponse) HasBucket() bool`

HasBucket returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


