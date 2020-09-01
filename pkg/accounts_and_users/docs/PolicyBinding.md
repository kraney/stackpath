# PolicyBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to **string** | The name of the role for this binding.  This is either roles/&lt;name&gt; or accounts/&lt;id&gt;/roles/&lt;name&gt;. | [optional] 
**Members** | Pointer to **[]string** | The members assigned to the roles in this binding.  This is user:&lt;name&gt;. | [optional] 

## Methods

### NewPolicyBinding

`func NewPolicyBinding() *PolicyBinding`

NewPolicyBinding instantiates a new PolicyBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyBindingWithDefaults

`func NewPolicyBindingWithDefaults() *PolicyBinding`

NewPolicyBindingWithDefaults instantiates a new PolicyBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *PolicyBinding) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PolicyBinding) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PolicyBinding) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *PolicyBinding) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetMembers

`func (o *PolicyBinding) GetMembers() []string`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *PolicyBinding) GetMembersOk() (*[]string, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *PolicyBinding) SetMembers(v []string)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *PolicyBinding) HasMembers() bool`

HasMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


