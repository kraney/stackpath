# IdentityCreateUserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to [**IdentityUser**](identityUser.md) |  | [optional] 

## Methods

### NewIdentityCreateUserResponse

`func NewIdentityCreateUserResponse() *IdentityCreateUserResponse`

NewIdentityCreateUserResponse instantiates a new IdentityCreateUserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityCreateUserResponseWithDefaults

`func NewIdentityCreateUserResponseWithDefaults() *IdentityCreateUserResponse`

NewIdentityCreateUserResponseWithDefaults instantiates a new IdentityCreateUserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *IdentityCreateUserResponse) GetUser() IdentityUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *IdentityCreateUserResponse) GetUserOk() (*IdentityUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *IdentityCreateUserResponse) SetUser(v IdentityUser)`

SetUser sets User field to given value.

### HasUser

`func (o *IdentityCreateUserResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


