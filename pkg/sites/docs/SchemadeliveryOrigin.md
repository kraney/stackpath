# SchemadeliveryOrigin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | An origin&#39;s unique identifier | [optional] 
**StackId** | Pointer to **string** | The stack for which the origin belongs to | [optional] 
**Dedicated** | Pointer to **bool** | Whether or not an origin is dedicated to a CDN site  Dedicated origins cannot be used by any site other than that to which it is dedicated. | [optional] 
**SiteId** | Pointer to **string** | The ID of the CDN site to which an origin is dedicated | [optional] 
**Http** | Pointer to [**DeliveryHTTPOrigin**](deliveryHTTPOrigin.md) |  | [optional] 
**StackpathStorage** | Pointer to [**DeliveryStackPathStorageOrigin**](deliveryStackPathStorageOrigin.md) |  | [optional] 
**S3Storage** | Pointer to [**DeliveryAWSS3Origin**](deliveryAWSS3Origin.md) |  | [optional] 
**GoogleCloudStorage** | Pointer to [**DeliveryGoogleStorageOrigin**](deliveryGoogleStorageOrigin.md) |  | [optional] 

## Methods

### NewSchemadeliveryOrigin

`func NewSchemadeliveryOrigin() *SchemadeliveryOrigin`

NewSchemadeliveryOrigin instantiates a new SchemadeliveryOrigin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemadeliveryOriginWithDefaults

`func NewSchemadeliveryOriginWithDefaults() *SchemadeliveryOrigin`

NewSchemadeliveryOriginWithDefaults instantiates a new SchemadeliveryOrigin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SchemadeliveryOrigin) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SchemadeliveryOrigin) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SchemadeliveryOrigin) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SchemadeliveryOrigin) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStackId

`func (o *SchemadeliveryOrigin) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *SchemadeliveryOrigin) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *SchemadeliveryOrigin) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *SchemadeliveryOrigin) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetDedicated

`func (o *SchemadeliveryOrigin) GetDedicated() bool`

GetDedicated returns the Dedicated field if non-nil, zero value otherwise.

### GetDedicatedOk

`func (o *SchemadeliveryOrigin) GetDedicatedOk() (*bool, bool)`

GetDedicatedOk returns a tuple with the Dedicated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicated

`func (o *SchemadeliveryOrigin) SetDedicated(v bool)`

SetDedicated sets Dedicated field to given value.

### HasDedicated

`func (o *SchemadeliveryOrigin) HasDedicated() bool`

HasDedicated returns a boolean if a field has been set.

### GetSiteId

`func (o *SchemadeliveryOrigin) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *SchemadeliveryOrigin) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *SchemadeliveryOrigin) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *SchemadeliveryOrigin) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetHttp

`func (o *SchemadeliveryOrigin) GetHttp() DeliveryHTTPOrigin`

GetHttp returns the Http field if non-nil, zero value otherwise.

### GetHttpOk

`func (o *SchemadeliveryOrigin) GetHttpOk() (*DeliveryHTTPOrigin, bool)`

GetHttpOk returns a tuple with the Http field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp

`func (o *SchemadeliveryOrigin) SetHttp(v DeliveryHTTPOrigin)`

SetHttp sets Http field to given value.

### HasHttp

`func (o *SchemadeliveryOrigin) HasHttp() bool`

HasHttp returns a boolean if a field has been set.

### GetStackpathStorage

`func (o *SchemadeliveryOrigin) GetStackpathStorage() DeliveryStackPathStorageOrigin`

GetStackpathStorage returns the StackpathStorage field if non-nil, zero value otherwise.

### GetStackpathStorageOk

`func (o *SchemadeliveryOrigin) GetStackpathStorageOk() (*DeliveryStackPathStorageOrigin, bool)`

GetStackpathStorageOk returns a tuple with the StackpathStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackpathStorage

`func (o *SchemadeliveryOrigin) SetStackpathStorage(v DeliveryStackPathStorageOrigin)`

SetStackpathStorage sets StackpathStorage field to given value.

### HasStackpathStorage

`func (o *SchemadeliveryOrigin) HasStackpathStorage() bool`

HasStackpathStorage returns a boolean if a field has been set.

### GetS3Storage

`func (o *SchemadeliveryOrigin) GetS3Storage() DeliveryAWSS3Origin`

GetS3Storage returns the S3Storage field if non-nil, zero value otherwise.

### GetS3StorageOk

`func (o *SchemadeliveryOrigin) GetS3StorageOk() (*DeliveryAWSS3Origin, bool)`

GetS3StorageOk returns a tuple with the S3Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Storage

`func (o *SchemadeliveryOrigin) SetS3Storage(v DeliveryAWSS3Origin)`

SetS3Storage sets S3Storage field to given value.

### HasS3Storage

`func (o *SchemadeliveryOrigin) HasS3Storage() bool`

HasS3Storage returns a boolean if a field has been set.

### GetGoogleCloudStorage

`func (o *SchemadeliveryOrigin) GetGoogleCloudStorage() DeliveryGoogleStorageOrigin`

GetGoogleCloudStorage returns the GoogleCloudStorage field if non-nil, zero value otherwise.

### GetGoogleCloudStorageOk

`func (o *SchemadeliveryOrigin) GetGoogleCloudStorageOk() (*DeliveryGoogleStorageOrigin, bool)`

GetGoogleCloudStorageOk returns a tuple with the GoogleCloudStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleCloudStorage

`func (o *SchemadeliveryOrigin) SetGoogleCloudStorage(v DeliveryGoogleStorageOrigin)`

SetGoogleCloudStorage sets GoogleCloudStorage field to given value.

### HasGoogleCloudStorage

`func (o *SchemadeliveryOrigin) HasGoogleCloudStorage() bool`

HasGoogleCloudStorage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


