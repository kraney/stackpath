# V1DockerRegistryCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Server** | Pointer to **string** | The server that the credentials should be used with  The server that the credentials should be used with. This value will default to the docker hub registry when not set. | [optional] 
**Username** | Pointer to **string** | The username that should be used for authenticate the image pull | [optional] 
**Password** | Pointer to **string** | The password that should be used to authenticate the image pull | [optional] 
**Email** | Pointer to **string** | The email address to use for the docker registry account | [optional] 

## Methods

### NewV1DockerRegistryCredentials

`func NewV1DockerRegistryCredentials() *V1DockerRegistryCredentials`

NewV1DockerRegistryCredentials instantiates a new V1DockerRegistryCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DockerRegistryCredentialsWithDefaults

`func NewV1DockerRegistryCredentialsWithDefaults() *V1DockerRegistryCredentials`

NewV1DockerRegistryCredentialsWithDefaults instantiates a new V1DockerRegistryCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServer

`func (o *V1DockerRegistryCredentials) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *V1DockerRegistryCredentials) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *V1DockerRegistryCredentials) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *V1DockerRegistryCredentials) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetUsername

`func (o *V1DockerRegistryCredentials) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *V1DockerRegistryCredentials) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *V1DockerRegistryCredentials) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *V1DockerRegistryCredentials) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *V1DockerRegistryCredentials) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *V1DockerRegistryCredentials) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *V1DockerRegistryCredentials) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *V1DockerRegistryCredentials) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetEmail

`func (o *V1DockerRegistryCredentials) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *V1DockerRegistryCredentials) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *V1DockerRegistryCredentials) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *V1DockerRegistryCredentials) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


