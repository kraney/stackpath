# SchemawafPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A WAF policy&#39;s unique identifier | [optional] 
**Name** | Pointer to **string** | A WAF policy&#39;s name | [optional] [readonly] 
**Description** | Pointer to **string** | A WAF policy&#39;s description | [optional] [readonly] 
**Action** | Pointer to [**WafPolicyAction**](wafPolicyAction.md) |  | [optional] [default to "BLOCK"]
**Enabled** | Pointer to **bool** | Whether or not a WAF policy is enabled | [optional] 

## Methods

### NewSchemawafPolicy

`func NewSchemawafPolicy() *SchemawafPolicy`

NewSchemawafPolicy instantiates a new SchemawafPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemawafPolicyWithDefaults

`func NewSchemawafPolicyWithDefaults() *SchemawafPolicy`

NewSchemawafPolicyWithDefaults instantiates a new SchemawafPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SchemawafPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SchemawafPolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SchemawafPolicy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SchemawafPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SchemawafPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SchemawafPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SchemawafPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SchemawafPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *SchemawafPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SchemawafPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SchemawafPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SchemawafPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAction

`func (o *SchemawafPolicy) GetAction() WafPolicyAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SchemawafPolicy) GetActionOk() (*WafPolicyAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SchemawafPolicy) SetAction(v WafPolicyAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *SchemawafPolicy) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetEnabled

`func (o *SchemawafPolicy) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SchemawafPolicy) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SchemawafPolicy) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SchemawafPolicy) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


