# IdentityRotateUserCredentialSecretResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The API credential&#39;s unique identifier | [optional] 
**Name** | Pointer to **string** | The API credential&#39;s name | [optional] 
**ClientId** | Pointer to **string** | The API credential&#39;s OAuth2 client ID | [optional] 
**ClientSecret** | Pointer to **string** | The API credential&#39;s OAuth2 client secret  The client secret is returned only once and is not stored by StackPath. Please take care to save this response. | [optional] 

## Methods

### NewIdentityRotateUserCredentialSecretResponse

`func NewIdentityRotateUserCredentialSecretResponse() *IdentityRotateUserCredentialSecretResponse`

NewIdentityRotateUserCredentialSecretResponse instantiates a new IdentityRotateUserCredentialSecretResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityRotateUserCredentialSecretResponseWithDefaults

`func NewIdentityRotateUserCredentialSecretResponseWithDefaults() *IdentityRotateUserCredentialSecretResponse`

NewIdentityRotateUserCredentialSecretResponseWithDefaults instantiates a new IdentityRotateUserCredentialSecretResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IdentityRotateUserCredentialSecretResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityRotateUserCredentialSecretResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityRotateUserCredentialSecretResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityRotateUserCredentialSecretResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IdentityRotateUserCredentialSecretResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityRotateUserCredentialSecretResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityRotateUserCredentialSecretResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdentityRotateUserCredentialSecretResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetClientId

`func (o *IdentityRotateUserCredentialSecretResponse) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *IdentityRotateUserCredentialSecretResponse) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *IdentityRotateUserCredentialSecretResponse) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *IdentityRotateUserCredentialSecretResponse) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *IdentityRotateUserCredentialSecretResponse) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *IdentityRotateUserCredentialSecretResponse) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *IdentityRotateUserCredentialSecretResponse) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *IdentityRotateUserCredentialSecretResponse) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


