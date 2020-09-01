# IdentityUpdateUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The user&#39;s new name | [optional] 
**PhoneNumber** | Pointer to **string** | The user&#39;s new phone number | [optional] 

## Methods

### NewIdentityUpdateUserRequest

`func NewIdentityUpdateUserRequest() *IdentityUpdateUserRequest`

NewIdentityUpdateUserRequest instantiates a new IdentityUpdateUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityUpdateUserRequestWithDefaults

`func NewIdentityUpdateUserRequestWithDefaults() *IdentityUpdateUserRequest`

NewIdentityUpdateUserRequestWithDefaults instantiates a new IdentityUpdateUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IdentityUpdateUserRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityUpdateUserRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityUpdateUserRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdentityUpdateUserRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *IdentityUpdateUserRequest) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *IdentityUpdateUserRequest) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *IdentityUpdateUserRequest) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *IdentityUpdateUserRequest) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


