# StackCreateStackRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | The ID of the account to create the stack for | [optional] 
**Slug** | Pointer to **string** | The StackPath stack&#39;s friendly, text-based unique identifier  This can be used in place of a stack&#39;s ID in most situations to identify a stack. It can be no larger than 30 characters, each being a lowercase letter, number, or dash. It also cannot start with a dash, end with a dash, or have consecutive dashes.If this value is not present, it is derived from the name. This value cannot be updated. | [optional] 
**Name** | Pointer to **string** | The stack&#39;s name | [optional] 

## Methods

### NewStackCreateStackRequest

`func NewStackCreateStackRequest() *StackCreateStackRequest`

NewStackCreateStackRequest instantiates a new StackCreateStackRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackCreateStackRequestWithDefaults

`func NewStackCreateStackRequestWithDefaults() *StackCreateStackRequest`

NewStackCreateStackRequestWithDefaults instantiates a new StackCreateStackRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *StackCreateStackRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *StackCreateStackRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *StackCreateStackRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *StackCreateStackRequest) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetSlug

`func (o *StackCreateStackRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *StackCreateStackRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *StackCreateStackRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *StackCreateStackRequest) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetName

`func (o *StackCreateStackRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StackCreateStackRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StackCreateStackRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StackCreateStackRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


