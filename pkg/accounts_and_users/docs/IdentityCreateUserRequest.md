# IdentityCreateUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | The user&#39;s email address  The user will use this as their login name at the StackPath customer portal | [optional] 
**Name** | Pointer to **string** | The user&#39;s name | [optional] 
**PhoneNumber** | Pointer to **string** | The user&#39;s phone number | [optional] 
**AuthorizedApplications** | Pointer to **[]string** | A set of application slugs the user should have access to  If no application slugs are applied then the user will have access to the StackPath customer portal and API by default. | [optional] 

## Methods

### NewIdentityCreateUserRequest

`func NewIdentityCreateUserRequest() *IdentityCreateUserRequest`

NewIdentityCreateUserRequest instantiates a new IdentityCreateUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityCreateUserRequestWithDefaults

`func NewIdentityCreateUserRequestWithDefaults() *IdentityCreateUserRequest`

NewIdentityCreateUserRequestWithDefaults instantiates a new IdentityCreateUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *IdentityCreateUserRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *IdentityCreateUserRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *IdentityCreateUserRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *IdentityCreateUserRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *IdentityCreateUserRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityCreateUserRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityCreateUserRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdentityCreateUserRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *IdentityCreateUserRequest) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *IdentityCreateUserRequest) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *IdentityCreateUserRequest) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *IdentityCreateUserRequest) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetAuthorizedApplications

`func (o *IdentityCreateUserRequest) GetAuthorizedApplications() []string`

GetAuthorizedApplications returns the AuthorizedApplications field if non-nil, zero value otherwise.

### GetAuthorizedApplicationsOk

`func (o *IdentityCreateUserRequest) GetAuthorizedApplicationsOk() (*[]string, bool)`

GetAuthorizedApplicationsOk returns a tuple with the AuthorizedApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedApplications

`func (o *IdentityCreateUserRequest) SetAuthorizedApplications(v []string)`

SetAuthorizedApplications sets AuthorizedApplications field to given value.

### HasAuthorizedApplications

`func (o *IdentityCreateUserRequest) HasAuthorizedApplications() bool`

HasAuthorizedApplications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


