# IdentityUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A user&#39;s unique identifier | [optional] 
**Email** | Pointer to **string** | A user&#39;s email address  Email addresses are used as login names to the StackPath customer portal | [optional] 
**Status** | Pointer to [**IdentityUserStatus**](identityUserStatus.md) |  | [optional] [default to "UNKNOWN"]
**Identities** | Pointer to [**[]IdentityUserIdentity**](identityUserIdentity.md) | A user&#39;s underlying authentication identities | [optional] 
**Accounts** | Pointer to [**[]IdentityAccount**](identityAccount.md) | The accounts that a user belongs to | [optional] 
**Name** | Pointer to **string** | A user&#39;s name | [optional] 
**PhoneNumber** | Pointer to **string** | A user&#39;s phone number | [optional] 

## Methods

### NewIdentityUser

`func NewIdentityUser() *IdentityUser`

NewIdentityUser instantiates a new IdentityUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityUserWithDefaults

`func NewIdentityUserWithDefaults() *IdentityUser`

NewIdentityUserWithDefaults instantiates a new IdentityUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IdentityUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityUser) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEmail

`func (o *IdentityUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *IdentityUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *IdentityUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *IdentityUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetStatus

`func (o *IdentityUser) GetStatus() IdentityUserStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IdentityUser) GetStatusOk() (*IdentityUserStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IdentityUser) SetStatus(v IdentityUserStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IdentityUser) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIdentities

`func (o *IdentityUser) GetIdentities() []IdentityUserIdentity`

GetIdentities returns the Identities field if non-nil, zero value otherwise.

### GetIdentitiesOk

`func (o *IdentityUser) GetIdentitiesOk() (*[]IdentityUserIdentity, bool)`

GetIdentitiesOk returns a tuple with the Identities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentities

`func (o *IdentityUser) SetIdentities(v []IdentityUserIdentity)`

SetIdentities sets Identities field to given value.

### HasIdentities

`func (o *IdentityUser) HasIdentities() bool`

HasIdentities returns a boolean if a field has been set.

### GetAccounts

`func (o *IdentityUser) GetAccounts() []IdentityAccount`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *IdentityUser) GetAccountsOk() (*[]IdentityAccount, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *IdentityUser) SetAccounts(v []IdentityAccount)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *IdentityUser) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetName

`func (o *IdentityUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityUser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdentityUser) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *IdentityUser) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *IdentityUser) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *IdentityUser) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *IdentityUser) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


