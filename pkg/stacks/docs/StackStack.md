# StackStack

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A stack&#39;s unique identifier | [optional] 
**AccountId** | Pointer to **string** | The ID of the account that a stack belongs to | [optional] 
**Slug** | Pointer to **string** | A stack&#39;s human-friendly identifier  This can be used in place of a stack&#39;s ID in most situations to identify a stack. It can be no longer than 30 characters, each being a lowercase letter, number, or dash. It also cannot start with a dash, end with a dash, or have consecutive dashes. This value can be provided on creation or is derived from the stack&#39;s name if a slug isn&#39;t provided. This value cannot be updated once it&#39;s created. | [optional] 
**Name** | Pointer to **string** | A stack&#39;s name | [optional] 
**CreatedAt** | Pointer to [**time.Time**](time.Time.md) | The date a stack was created | [optional] [readonly] 
**UpdatedAt** | Pointer to [**time.Time**](time.Time.md) | The date a stack was last updated | [optional] [readonly] 
**Status** | Pointer to [**StackStackStatus**](stackStackStatus.md) |  | [optional] [default to "PENDING"]

## Methods

### NewStackStack

`func NewStackStack() *StackStack`

NewStackStack instantiates a new StackStack object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackStackWithDefaults

`func NewStackStackWithDefaults() *StackStack`

NewStackStackWithDefaults instantiates a new StackStack object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StackStack) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StackStack) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StackStack) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StackStack) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountId

`func (o *StackStack) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *StackStack) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *StackStack) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *StackStack) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetSlug

`func (o *StackStack) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *StackStack) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *StackStack) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *StackStack) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetName

`func (o *StackStack) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StackStack) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StackStack) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StackStack) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *StackStack) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StackStack) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StackStack) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *StackStack) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *StackStack) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *StackStack) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *StackStack) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *StackStack) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetStatus

`func (o *StackStack) GetStatus() StackStackStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StackStack) GetStatusOk() (*StackStackStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StackStack) SetStatus(v StackStackStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StackStack) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


