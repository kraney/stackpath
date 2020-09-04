# WafPolicyGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A policy group&#39;s ID | [optional] 
**Name** | Pointer to **string** | A policy group&#39;s name | [optional] [readonly] 
**Policies** | Pointer to [**[]SchemawafPolicy**](schemawafPolicy.md) | A list of the policies in a policy group | [optional] 

## Methods

### NewWafPolicyGroup

`func NewWafPolicyGroup() *WafPolicyGroup`

NewWafPolicyGroup instantiates a new WafPolicyGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafPolicyGroupWithDefaults

`func NewWafPolicyGroupWithDefaults() *WafPolicyGroup`

NewWafPolicyGroupWithDefaults instantiates a new WafPolicyGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WafPolicyGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WafPolicyGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WafPolicyGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WafPolicyGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *WafPolicyGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WafPolicyGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WafPolicyGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WafPolicyGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPolicies

`func (o *WafPolicyGroup) GetPolicies() []SchemawafPolicy`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *WafPolicyGroup) GetPoliciesOk() (*[]SchemawafPolicy, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *WafPolicyGroup) SetPolicies(v []SchemawafPolicy)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *WafPolicyGroup) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


