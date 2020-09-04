# DeliveryStackPathStorageOrigin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketName** | Pointer to **string** | The origin bucket&#39;s name | [optional] 
**BucketRegion** | Pointer to **string** | The origin bucket&#39;s region | [optional] 
**Authentication** | Pointer to [**DeliveryOriginAuthentication**](deliveryOriginAuthentication.md) |  | [optional] 

## Methods

### NewDeliveryStackPathStorageOrigin

`func NewDeliveryStackPathStorageOrigin() *DeliveryStackPathStorageOrigin`

NewDeliveryStackPathStorageOrigin instantiates a new DeliveryStackPathStorageOrigin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryStackPathStorageOriginWithDefaults

`func NewDeliveryStackPathStorageOriginWithDefaults() *DeliveryStackPathStorageOrigin`

NewDeliveryStackPathStorageOriginWithDefaults instantiates a new DeliveryStackPathStorageOrigin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketName

`func (o *DeliveryStackPathStorageOrigin) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *DeliveryStackPathStorageOrigin) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *DeliveryStackPathStorageOrigin) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### HasBucketName

`func (o *DeliveryStackPathStorageOrigin) HasBucketName() bool`

HasBucketName returns a boolean if a field has been set.

### GetBucketRegion

`func (o *DeliveryStackPathStorageOrigin) GetBucketRegion() string`

GetBucketRegion returns the BucketRegion field if non-nil, zero value otherwise.

### GetBucketRegionOk

`func (o *DeliveryStackPathStorageOrigin) GetBucketRegionOk() (*string, bool)`

GetBucketRegionOk returns a tuple with the BucketRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketRegion

`func (o *DeliveryStackPathStorageOrigin) SetBucketRegion(v string)`

SetBucketRegion sets BucketRegion field to given value.

### HasBucketRegion

`func (o *DeliveryStackPathStorageOrigin) HasBucketRegion() bool`

HasBucketRegion returns a boolean if a field has been set.

### GetAuthentication

`func (o *DeliveryStackPathStorageOrigin) GetAuthentication() DeliveryOriginAuthentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *DeliveryStackPathStorageOrigin) GetAuthenticationOk() (*DeliveryOriginAuthentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *DeliveryStackPathStorageOrigin) SetAuthentication(v DeliveryOriginAuthentication)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *DeliveryStackPathStorageOrigin) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


