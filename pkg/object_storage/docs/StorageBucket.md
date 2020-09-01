# StorageBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID for the bucket | [optional] 
**Label** | Pointer to **string** | The name of the bucket | [optional] 
**EndpointUrl** | Pointer to **string** | The URL used to access the bucket | [optional] 
**Visibility** | Pointer to [**StorageBucketVisibility**](storageBucketVisibility.md) |  | [optional] [default to "PRIVATE"]
**CreatedAt** | Pointer to [**time.Time**](time.Time.md) | The date when the bucket was created | [optional] 
**UpdatedAt** | Pointer to [**time.Time**](time.Time.md) | The date when the bucket was last updated | [optional] 
**Region** | Pointer to **string** | The region in which the bucket is created. Available regions are: us-east-1, us-west-1, eu-central-1 | [optional] 

## Methods

### NewStorageBucket

`func NewStorageBucket() *StorageBucket`

NewStorageBucket instantiates a new StorageBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageBucketWithDefaults

`func NewStorageBucketWithDefaults() *StorageBucket`

NewStorageBucketWithDefaults instantiates a new StorageBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StorageBucket) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StorageBucket) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StorageBucket) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StorageBucket) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *StorageBucket) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *StorageBucket) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *StorageBucket) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *StorageBucket) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetEndpointUrl

`func (o *StorageBucket) GetEndpointUrl() string`

GetEndpointUrl returns the EndpointUrl field if non-nil, zero value otherwise.

### GetEndpointUrlOk

`func (o *StorageBucket) GetEndpointUrlOk() (*string, bool)`

GetEndpointUrlOk returns a tuple with the EndpointUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointUrl

`func (o *StorageBucket) SetEndpointUrl(v string)`

SetEndpointUrl sets EndpointUrl field to given value.

### HasEndpointUrl

`func (o *StorageBucket) HasEndpointUrl() bool`

HasEndpointUrl returns a boolean if a field has been set.

### GetVisibility

`func (o *StorageBucket) GetVisibility() StorageBucketVisibility`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *StorageBucket) GetVisibilityOk() (*StorageBucketVisibility, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *StorageBucket) SetVisibility(v StorageBucketVisibility)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *StorageBucket) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetCreatedAt

`func (o *StorageBucket) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StorageBucket) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StorageBucket) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *StorageBucket) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *StorageBucket) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *StorageBucket) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *StorageBucket) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *StorageBucket) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetRegion

`func (o *StorageBucket) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *StorageBucket) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *StorageBucket) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *StorageBucket) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


