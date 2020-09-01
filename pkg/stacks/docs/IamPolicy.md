# IamPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bindings** | Pointer to [**[]PolicyBinding**](PolicyBinding.md) | Bindings to assign members to roles | [optional] 
**Version** | Pointer to **int64** | The current update number used to ensure updates happen to the expected version.  On first update this must be 0 and on each successive update this must be the last known version. When getting or as the result of a set, this will be the current version. | [optional] 
**CreatedAt** | Pointer to [**time.Time**](time.Time.md) | When this policy was created.  Always present on get, ignored on set. | [optional] [readonly] 
**UpdatedAt** | Pointer to [**time.Time**](time.Time.md) | When this policy was last updated.  Could be absent on get if not updated since initial creation. Ignored on set. | [optional] [readonly] 

## Methods

### NewIamPolicy

`func NewIamPolicy() *IamPolicy`

NewIamPolicy instantiates a new IamPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamPolicyWithDefaults

`func NewIamPolicyWithDefaults() *IamPolicy`

NewIamPolicyWithDefaults instantiates a new IamPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBindings

`func (o *IamPolicy) GetBindings() []PolicyBinding`

GetBindings returns the Bindings field if non-nil, zero value otherwise.

### GetBindingsOk

`func (o *IamPolicy) GetBindingsOk() (*[]PolicyBinding, bool)`

GetBindingsOk returns a tuple with the Bindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindings

`func (o *IamPolicy) SetBindings(v []PolicyBinding)`

SetBindings sets Bindings field to given value.

### HasBindings

`func (o *IamPolicy) HasBindings() bool`

HasBindings returns a boolean if a field has been set.

### GetVersion

`func (o *IamPolicy) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *IamPolicy) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *IamPolicy) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *IamPolicy) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IamPolicy) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IamPolicy) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IamPolicy) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IamPolicy) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IamPolicy) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IamPolicy) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IamPolicy) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IamPolicy) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


