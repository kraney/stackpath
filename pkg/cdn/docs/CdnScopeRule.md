# CdnScopeRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | An EdgeRule&#39;s unique identifier | [optional] 
**Name** | Pointer to **string** | An EdgeRule&#39;s name | [optional] 
**Slug** | Pointer to **string** | A human and machine readable name for an EdgeRule | [optional] 
**CreatedAt** | Pointer to [**time.Time**](time.Time.md) | The date an EdgeRule was created | [optional] 
**UpdatedAt** | Pointer to [**time.Time**](time.Time.md) | The date an EdgeRule was last updated | [optional] 

## Methods

### NewCdnScopeRule

`func NewCdnScopeRule() *CdnScopeRule`

NewCdnScopeRule instantiates a new CdnScopeRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdnScopeRuleWithDefaults

`func NewCdnScopeRuleWithDefaults() *CdnScopeRule`

NewCdnScopeRuleWithDefaults instantiates a new CdnScopeRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CdnScopeRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CdnScopeRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CdnScopeRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CdnScopeRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CdnScopeRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CdnScopeRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CdnScopeRule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CdnScopeRule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *CdnScopeRule) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *CdnScopeRule) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *CdnScopeRule) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *CdnScopeRule) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CdnScopeRule) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CdnScopeRule) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CdnScopeRule) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CdnScopeRule) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CdnScopeRule) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CdnScopeRule) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CdnScopeRule) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CdnScopeRule) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


