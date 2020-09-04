# IdentityUserIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**IdentityId** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to [**IdentityProvider**](identityProvider.md) |  | [optional] 

## Methods

### NewIdentityUserIdentity

`func NewIdentityUserIdentity() *IdentityUserIdentity`

NewIdentityUserIdentity instantiates a new IdentityUserIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityUserIdentityWithDefaults

`func NewIdentityUserIdentityWithDefaults() *IdentityUserIdentity`

NewIdentityUserIdentityWithDefaults instantiates a new IdentityUserIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IdentityUserIdentity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityUserIdentity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityUserIdentity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityUserIdentity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentityId

`func (o *IdentityUserIdentity) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *IdentityUserIdentity) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *IdentityUserIdentity) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.

### HasIdentityId

`func (o *IdentityUserIdentity) HasIdentityId() bool`

HasIdentityId returns a boolean if a field has been set.

### GetEmail

`func (o *IdentityUserIdentity) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *IdentityUserIdentity) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *IdentityUserIdentity) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *IdentityUserIdentity) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetProvider

`func (o *IdentityUserIdentity) GetProvider() IdentityProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *IdentityUserIdentity) GetProviderOk() (*IdentityProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *IdentityUserIdentity) SetProvider(v IdentityProvider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *IdentityUserIdentity) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


