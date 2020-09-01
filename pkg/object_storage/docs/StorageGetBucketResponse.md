# StorageGetBucketResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | Pointer to [**StorageBucket**](storageBucket.md) |  | [optional] 

## Methods

### NewStorageGetBucketResponse

`func NewStorageGetBucketResponse() *StorageGetBucketResponse`

NewStorageGetBucketResponse instantiates a new StorageGetBucketResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageGetBucketResponseWithDefaults

`func NewStorageGetBucketResponseWithDefaults() *StorageGetBucketResponse`

NewStorageGetBucketResponseWithDefaults instantiates a new StorageGetBucketResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *StorageGetBucketResponse) GetBucket() StorageBucket`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *StorageGetBucketResponse) GetBucketOk() (*StorageBucket, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *StorageGetBucketResponse) SetBucket(v StorageBucket)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *StorageGetBucketResponse) HasBucket() bool`

HasBucket returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


