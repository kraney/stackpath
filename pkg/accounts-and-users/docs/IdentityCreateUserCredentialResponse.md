# IdentityCreateUserCredentialResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The new API credential&#39;s unique identifier | [optional] 
**Name** | Pointer to **string** | The new API credential&#39;s name | [optional] 
**ClientId** | Pointer to **string** | The new API credential&#39;s OAuth2 client ID | [optional] 
**ClientSecret** | Pointer to **string** | The new API credential&#39;s OAuth2 client secret  The client secret is returned only once and is not stored by StackPath. Please take care to save this response. | [optional] 

## Methods

### NewIdentityCreateUserCredentialResponse

`func NewIdentityCreateUserCredentialResponse() *IdentityCreateUserCredentialResponse`

NewIdentityCreateUserCredentialResponse instantiates a new IdentityCreateUserCredentialResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityCreateUserCredentialResponseWithDefaults

`func NewIdentityCreateUserCredentialResponseWithDefaults() *IdentityCreateUserCredentialResponse`

NewIdentityCreateUserCredentialResponseWithDefaults instantiates a new IdentityCreateUserCredentialResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IdentityCreateUserCredentialResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityCreateUserCredentialResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityCreateUserCredentialResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityCreateUserCredentialResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IdentityCreateUserCredentialResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityCreateUserCredentialResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityCreateUserCredentialResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdentityCreateUserCredentialResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetClientId

`func (o *IdentityCreateUserCredentialResponse) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *IdentityCreateUserCredentialResponse) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *IdentityCreateUserCredentialResponse) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *IdentityCreateUserCredentialResponse) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *IdentityCreateUserCredentialResponse) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *IdentityCreateUserCredentialResponse) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *IdentityCreateUserCredentialResponse) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *IdentityCreateUserCredentialResponse) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


