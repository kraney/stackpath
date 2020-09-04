# DeliveryCreateSiteRequestOrigin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** | The path the site should request from the origin | [optional] 
**Hostname** | Pointer to **string** | The origin&#39;s fully-qualified domain name | [optional] 
**Port** | Pointer to **int32** | The TCP port the site should connect to for HTTP requests | [optional] 
**SecurePort** | Pointer to **int32** | The TCP port the site should connect to for HTTPS requests | [optional] 
**Http** | Pointer to [**DeliveryHTTPOrigin**](deliveryHTTPOrigin.md) |  | [optional] 
**StackpathStorage** | Pointer to [**DeliveryStackPathStorageOrigin**](deliveryStackPathStorageOrigin.md) |  | [optional] 
**S3Storage** | Pointer to [**DeliveryAWSS3Origin**](deliveryAWSS3Origin.md) |  | [optional] 
**GoogleCloudStorage** | Pointer to [**DeliveryGoogleStorageOrigin**](deliveryGoogleStorageOrigin.md) |  | [optional] 

## Methods

### NewDeliveryCreateSiteRequestOrigin

`func NewDeliveryCreateSiteRequestOrigin() *DeliveryCreateSiteRequestOrigin`

NewDeliveryCreateSiteRequestOrigin instantiates a new DeliveryCreateSiteRequestOrigin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryCreateSiteRequestOriginWithDefaults

`func NewDeliveryCreateSiteRequestOriginWithDefaults() *DeliveryCreateSiteRequestOrigin`

NewDeliveryCreateSiteRequestOriginWithDefaults instantiates a new DeliveryCreateSiteRequestOrigin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *DeliveryCreateSiteRequestOrigin) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DeliveryCreateSiteRequestOrigin) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DeliveryCreateSiteRequestOrigin) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *DeliveryCreateSiteRequestOrigin) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetHostname

`func (o *DeliveryCreateSiteRequestOrigin) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *DeliveryCreateSiteRequestOrigin) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *DeliveryCreateSiteRequestOrigin) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *DeliveryCreateSiteRequestOrigin) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetPort

`func (o *DeliveryCreateSiteRequestOrigin) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DeliveryCreateSiteRequestOrigin) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DeliveryCreateSiteRequestOrigin) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *DeliveryCreateSiteRequestOrigin) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSecurePort

`func (o *DeliveryCreateSiteRequestOrigin) GetSecurePort() int32`

GetSecurePort returns the SecurePort field if non-nil, zero value otherwise.

### GetSecurePortOk

`func (o *DeliveryCreateSiteRequestOrigin) GetSecurePortOk() (*int32, bool)`

GetSecurePortOk returns a tuple with the SecurePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurePort

`func (o *DeliveryCreateSiteRequestOrigin) SetSecurePort(v int32)`

SetSecurePort sets SecurePort field to given value.

### HasSecurePort

`func (o *DeliveryCreateSiteRequestOrigin) HasSecurePort() bool`

HasSecurePort returns a boolean if a field has been set.

### GetHttp

`func (o *DeliveryCreateSiteRequestOrigin) GetHttp() DeliveryHTTPOrigin`

GetHttp returns the Http field if non-nil, zero value otherwise.

### GetHttpOk

`func (o *DeliveryCreateSiteRequestOrigin) GetHttpOk() (*DeliveryHTTPOrigin, bool)`

GetHttpOk returns a tuple with the Http field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp

`func (o *DeliveryCreateSiteRequestOrigin) SetHttp(v DeliveryHTTPOrigin)`

SetHttp sets Http field to given value.

### HasHttp

`func (o *DeliveryCreateSiteRequestOrigin) HasHttp() bool`

HasHttp returns a boolean if a field has been set.

### GetStackpathStorage

`func (o *DeliveryCreateSiteRequestOrigin) GetStackpathStorage() DeliveryStackPathStorageOrigin`

GetStackpathStorage returns the StackpathStorage field if non-nil, zero value otherwise.

### GetStackpathStorageOk

`func (o *DeliveryCreateSiteRequestOrigin) GetStackpathStorageOk() (*DeliveryStackPathStorageOrigin, bool)`

GetStackpathStorageOk returns a tuple with the StackpathStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackpathStorage

`func (o *DeliveryCreateSiteRequestOrigin) SetStackpathStorage(v DeliveryStackPathStorageOrigin)`

SetStackpathStorage sets StackpathStorage field to given value.

### HasStackpathStorage

`func (o *DeliveryCreateSiteRequestOrigin) HasStackpathStorage() bool`

HasStackpathStorage returns a boolean if a field has been set.

### GetS3Storage

`func (o *DeliveryCreateSiteRequestOrigin) GetS3Storage() DeliveryAWSS3Origin`

GetS3Storage returns the S3Storage field if non-nil, zero value otherwise.

### GetS3StorageOk

`func (o *DeliveryCreateSiteRequestOrigin) GetS3StorageOk() (*DeliveryAWSS3Origin, bool)`

GetS3StorageOk returns a tuple with the S3Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Storage

`func (o *DeliveryCreateSiteRequestOrigin) SetS3Storage(v DeliveryAWSS3Origin)`

SetS3Storage sets S3Storage field to given value.

### HasS3Storage

`func (o *DeliveryCreateSiteRequestOrigin) HasS3Storage() bool`

HasS3Storage returns a boolean if a field has been set.

### GetGoogleCloudStorage

`func (o *DeliveryCreateSiteRequestOrigin) GetGoogleCloudStorage() DeliveryGoogleStorageOrigin`

GetGoogleCloudStorage returns the GoogleCloudStorage field if non-nil, zero value otherwise.

### GetGoogleCloudStorageOk

`func (o *DeliveryCreateSiteRequestOrigin) GetGoogleCloudStorageOk() (*DeliveryGoogleStorageOrigin, bool)`

GetGoogleCloudStorageOk returns a tuple with the GoogleCloudStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleCloudStorage

`func (o *DeliveryCreateSiteRequestOrigin) SetGoogleCloudStorage(v DeliveryGoogleStorageOrigin)`

SetGoogleCloudStorage sets GoogleCloudStorage field to given value.

### HasGoogleCloudStorage

`func (o *DeliveryCreateSiteRequestOrigin) HasGoogleCloudStorage() bool`

HasGoogleCloudStorage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


