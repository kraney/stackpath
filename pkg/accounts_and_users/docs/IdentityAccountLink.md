# IdentityAccountLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | An account link&#39;s unique identifier | [optional] 
**Provider** | Pointer to **string** | The name of the identity provider powering an account link | [optional] 
**RemoteId** | Pointer to **string** | A StackPath account&#39;s foreign key at an identity provider | [optional] 
**AccountId** | Pointer to **string** | The ID of the StackPath account linked to an identify provider | [optional] 

## Methods

### NewIdentityAccountLink

`func NewIdentityAccountLink() *IdentityAccountLink`

NewIdentityAccountLink instantiates a new IdentityAccountLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityAccountLinkWithDefaults

`func NewIdentityAccountLinkWithDefaults() *IdentityAccountLink`

NewIdentityAccountLinkWithDefaults instantiates a new IdentityAccountLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IdentityAccountLink) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityAccountLink) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityAccountLink) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityAccountLink) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProvider

`func (o *IdentityAccountLink) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *IdentityAccountLink) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *IdentityAccountLink) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *IdentityAccountLink) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRemoteId

`func (o *IdentityAccountLink) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *IdentityAccountLink) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *IdentityAccountLink) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *IdentityAccountLink) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### GetAccountId

`func (o *IdentityAccountLink) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *IdentityAccountLink) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *IdentityAccountLink) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *IdentityAccountLink) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


