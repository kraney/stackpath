# StorageCreateBucketRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | The name of the bucket to be created | [optional] 
**Region** | Pointer to **string** | The region where to create the bucket, defaults to us-east-1 | [optional] 

## Methods

### NewStorageCreateBucketRequest

`func NewStorageCreateBucketRequest() *StorageCreateBucketRequest`

NewStorageCreateBucketRequest instantiates a new StorageCreateBucketRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageCreateBucketRequestWithDefaults

`func NewStorageCreateBucketRequestWithDefaults() *StorageCreateBucketRequest`

NewStorageCreateBucketRequestWithDefaults instantiates a new StorageCreateBucketRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *StorageCreateBucketRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *StorageCreateBucketRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *StorageCreateBucketRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *StorageCreateBucketRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetRegion

`func (o *StorageCreateBucketRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *StorageCreateBucketRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *StorageCreateBucketRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *StorageCreateBucketRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


