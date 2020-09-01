# IdentityAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A StackPath account&#39;s unique identifier | [optional] 
**ParentId** | Pointer to **string** | The ID of an account that a StackPath account may belong to | [optional] 
**Name** | Pointer to **string** | A StackPath account&#39;s name | [optional] 
**RootUserId** | Pointer to **string** | The ID of a StackPath account&#39;s primary user | [optional] 
**CreatedAt** | Pointer to [**time.Time**](time.Time.md) | The date a StackPath account was created | [optional] 
**UpdatedAt** | Pointer to [**time.Time**](time.Time.md) | The date a StackPath account was last updated | [optional] 

## Methods

### NewIdentityAccount

`func NewIdentityAccount() *IdentityAccount`

NewIdentityAccount instantiates a new IdentityAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityAccountWithDefaults

`func NewIdentityAccountWithDefaults() *IdentityAccount`

NewIdentityAccountWithDefaults instantiates a new IdentityAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IdentityAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityAccount) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetParentId

`func (o *IdentityAccount) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *IdentityAccount) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *IdentityAccount) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *IdentityAccount) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetName

`func (o *IdentityAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityAccount) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdentityAccount) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRootUserId

`func (o *IdentityAccount) GetRootUserId() string`

GetRootUserId returns the RootUserId field if non-nil, zero value otherwise.

### GetRootUserIdOk

`func (o *IdentityAccount) GetRootUserIdOk() (*string, bool)`

GetRootUserIdOk returns a tuple with the RootUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootUserId

`func (o *IdentityAccount) SetRootUserId(v string)`

SetRootUserId sets RootUserId field to given value.

### HasRootUserId

`func (o *IdentityAccount) HasRootUserId() bool`

HasRootUserId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IdentityAccount) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IdentityAccount) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IdentityAccount) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IdentityAccount) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IdentityAccount) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IdentityAccount) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IdentityAccount) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IdentityAccount) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


