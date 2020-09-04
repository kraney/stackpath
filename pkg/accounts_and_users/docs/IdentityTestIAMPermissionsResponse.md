# IdentityTestIAMPermissionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | The ID of the StackPath account the permissions were tested on | [optional] 
**Permissions** | Pointer to **[]string** | The set of permissions that the calling user does have on the account | [optional] 

## Methods

### NewIdentityTestIAMPermissionsResponse

`func NewIdentityTestIAMPermissionsResponse() *IdentityTestIAMPermissionsResponse`

NewIdentityTestIAMPermissionsResponse instantiates a new IdentityTestIAMPermissionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityTestIAMPermissionsResponseWithDefaults

`func NewIdentityTestIAMPermissionsResponseWithDefaults() *IdentityTestIAMPermissionsResponse`

NewIdentityTestIAMPermissionsResponseWithDefaults instantiates a new IdentityTestIAMPermissionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *IdentityTestIAMPermissionsResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *IdentityTestIAMPermissionsResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *IdentityTestIAMPermissionsResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *IdentityTestIAMPermissionsResponse) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetPermissions

`func (o *IdentityTestIAMPermissionsResponse) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *IdentityTestIAMPermissionsResponse) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *IdentityTestIAMPermissionsResponse) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *IdentityTestIAMPermissionsResponse) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


