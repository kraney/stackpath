# DeliveryConnectScopeToOriginRequestOrigin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Http** | Pointer to [**DeliveryHTTPOrigin**](deliveryHTTPOrigin.md) |  | [optional] 
**StackpathStorage** | Pointer to [**DeliveryStackPathStorageOrigin**](deliveryStackPathStorageOrigin.md) |  | [optional] 
**S3Storage** | Pointer to [**DeliveryAWSS3Origin**](deliveryAWSS3Origin.md) |  | [optional] 
**GoogleCloudStorage** | Pointer to [**DeliveryGoogleStorageOrigin**](deliveryGoogleStorageOrigin.md) |  | [optional] 

## Methods

### NewDeliveryConnectScopeToOriginRequestOrigin

`func NewDeliveryConnectScopeToOriginRequestOrigin() *DeliveryConnectScopeToOriginRequestOrigin`

NewDeliveryConnectScopeToOriginRequestOrigin instantiates a new DeliveryConnectScopeToOriginRequestOrigin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryConnectScopeToOriginRequestOriginWithDefaults

`func NewDeliveryConnectScopeToOriginRequestOriginWithDefaults() *DeliveryConnectScopeToOriginRequestOrigin`

NewDeliveryConnectScopeToOriginRequestOriginWithDefaults instantiates a new DeliveryConnectScopeToOriginRequestOrigin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttp

`func (o *DeliveryConnectScopeToOriginRequestOrigin) GetHttp() DeliveryHTTPOrigin`

GetHttp returns the Http field if non-nil, zero value otherwise.

### GetHttpOk

`func (o *DeliveryConnectScopeToOriginRequestOrigin) GetHttpOk() (*DeliveryHTTPOrigin, bool)`

GetHttpOk returns a tuple with the Http field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp

`func (o *DeliveryConnectScopeToOriginRequestOrigin) SetHttp(v DeliveryHTTPOrigin)`

SetHttp sets Http field to given value.

### HasHttp

`func (o *DeliveryConnectScopeToOriginRequestOrigin) HasHttp() bool`

HasHttp returns a boolean if a field has been set.

### GetStackpathStorage

`func (o *DeliveryConnectScopeToOriginRequestOrigin) GetStackpathStorage() DeliveryStackPathStorageOrigin`

GetStackpathStorage returns the StackpathStorage field if non-nil, zero value otherwise.

### GetStackpathStorageOk

`func (o *DeliveryConnectScopeToOriginRequestOrigin) GetStackpathStorageOk() (*DeliveryStackPathStorageOrigin, bool)`

GetStackpathStorageOk returns a tuple with the StackpathStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackpathStorage

`func (o *DeliveryConnectScopeToOriginRequestOrigin) SetStackpathStorage(v DeliveryStackPathStorageOrigin)`

SetStackpathStorage sets StackpathStorage field to given value.

### HasStackpathStorage

`func (o *DeliveryConnectScopeToOriginRequestOrigin) HasStackpathStorage() bool`

HasStackpathStorage returns a boolean if a field has been set.

### GetS3Storage

`func (o *DeliveryConnectScopeToOriginRequestOrigin) GetS3Storage() DeliveryAWSS3Origin`

GetS3Storage returns the S3Storage field if non-nil, zero value otherwise.

### GetS3StorageOk

`func (o *DeliveryConnectScopeToOriginRequestOrigin) GetS3StorageOk() (*DeliveryAWSS3Origin, bool)`

GetS3StorageOk returns a tuple with the S3Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Storage

`func (o *DeliveryConnectScopeToOriginRequestOrigin) SetS3Storage(v DeliveryAWSS3Origin)`

SetS3Storage sets S3Storage field to given value.

### HasS3Storage

`func (o *DeliveryConnectScopeToOriginRequestOrigin) HasS3Storage() bool`

HasS3Storage returns a boolean if a field has been set.

### GetGoogleCloudStorage

`func (o *DeliveryConnectScopeToOriginRequestOrigin) GetGoogleCloudStorage() DeliveryGoogleStorageOrigin`

GetGoogleCloudStorage returns the GoogleCloudStorage field if non-nil, zero value otherwise.

### GetGoogleCloudStorageOk

`func (o *DeliveryConnectScopeToOriginRequestOrigin) GetGoogleCloudStorageOk() (*DeliveryGoogleStorageOrigin, bool)`

GetGoogleCloudStorageOk returns a tuple with the GoogleCloudStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleCloudStorage

`func (o *DeliveryConnectScopeToOriginRequestOrigin) SetGoogleCloudStorage(v DeliveryGoogleStorageOrigin)`

SetGoogleCloudStorage sets GoogleCloudStorage field to given value.

### HasGoogleCloudStorage

`func (o *DeliveryConnectScopeToOriginRequestOrigin) HasGoogleCloudStorage() bool`

HasGoogleCloudStorage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


