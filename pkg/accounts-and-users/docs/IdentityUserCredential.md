# IdentityUserCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | An API credential&#39;s unique identifier | [optional] 
**Name** | Pointer to **string** | An API credential&#39;s name  API credential names are typically associated with the user&#39;s application and operating environment | [optional] 
**ClientId** | Pointer to **string** | An API credential&#39;s OAuth2 client ID | [optional] 
**CreatedAt** | Pointer to [**time.Time**](time.Time.md) | The date an API credential was created | [optional] 
**UpdatedAt** | Pointer to [**time.Time**](time.Time.md) | The date an API credential was last updated | [optional] 

## Methods

### NewIdentityUserCredential

`func NewIdentityUserCredential() *IdentityUserCredential`

NewIdentityUserCredential instantiates a new IdentityUserCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityUserCredentialWithDefaults

`func NewIdentityUserCredentialWithDefaults() *IdentityUserCredential`

NewIdentityUserCredentialWithDefaults instantiates a new IdentityUserCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IdentityUserCredential) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityUserCredential) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityUserCredential) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityUserCredential) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IdentityUserCredential) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityUserCredential) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityUserCredential) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdentityUserCredential) HasName() bool`

HasName returns a boolean if a field has been set.

### GetClientId

`func (o *IdentityUserCredential) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *IdentityUserCredential) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *IdentityUserCredential) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *IdentityUserCredential) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IdentityUserCredential) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IdentityUserCredential) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IdentityUserCredential) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IdentityUserCredential) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IdentityUserCredential) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IdentityUserCredential) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IdentityUserCredential) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IdentityUserCredential) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


