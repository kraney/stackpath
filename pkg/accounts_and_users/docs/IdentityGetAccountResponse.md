# IdentityGetAccountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**IdentityAccount**](identityAccount.md) |  | [optional] 

## Methods

### NewIdentityGetAccountResponse

`func NewIdentityGetAccountResponse() *IdentityGetAccountResponse`

NewIdentityGetAccountResponse instantiates a new IdentityGetAccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityGetAccountResponseWithDefaults

`func NewIdentityGetAccountResponseWithDefaults() *IdentityGetAccountResponse`

NewIdentityGetAccountResponseWithDefaults instantiates a new IdentityGetAccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *IdentityGetAccountResponse) GetAccount() IdentityAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IdentityGetAccountResponse) GetAccountOk() (*IdentityAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IdentityGetAccountResponse) SetAccount(v IdentityAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *IdentityGetAccountResponse) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


