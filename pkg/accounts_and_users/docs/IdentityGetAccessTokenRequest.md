# IdentityGetAccessTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrantType** | Pointer to **string** | The OAuth2 grant type  Currently, only the \&quot;client_credentials\&quot; grant type is supported. | [optional] [default to "client_credentials"]
**ClientId** | Pointer to **string** | The client ID from your API credential | [optional] 
**ClientSecret** | Pointer to **string** | The client secret from your API credential  For security reasons, client secrets are not stored in StackPath&#39;s internal systems after they are generated. Please record your API client secret after generating it. | [optional] 

## Methods

### NewIdentityGetAccessTokenRequest

`func NewIdentityGetAccessTokenRequest() *IdentityGetAccessTokenRequest`

NewIdentityGetAccessTokenRequest instantiates a new IdentityGetAccessTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityGetAccessTokenRequestWithDefaults

`func NewIdentityGetAccessTokenRequestWithDefaults() *IdentityGetAccessTokenRequest`

NewIdentityGetAccessTokenRequestWithDefaults instantiates a new IdentityGetAccessTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrantType

`func (o *IdentityGetAccessTokenRequest) GetGrantType() string`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *IdentityGetAccessTokenRequest) GetGrantTypeOk() (*string, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *IdentityGetAccessTokenRequest) SetGrantType(v string)`

SetGrantType sets GrantType field to given value.

### HasGrantType

`func (o *IdentityGetAccessTokenRequest) HasGrantType() bool`

HasGrantType returns a boolean if a field has been set.

### GetClientId

`func (o *IdentityGetAccessTokenRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *IdentityGetAccessTokenRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *IdentityGetAccessTokenRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *IdentityGetAccessTokenRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *IdentityGetAccessTokenRequest) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *IdentityGetAccessTokenRequest) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *IdentityGetAccessTokenRequest) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *IdentityGetAccessTokenRequest) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


