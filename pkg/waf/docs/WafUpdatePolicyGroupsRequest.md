# WafUpdatePolicyGroupsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyGroups** | Pointer to [**[]WafPolicyGroup**](wafPolicyGroup.md) | A site&#39;s WAF policy groups | [optional] 

## Methods

### NewWafUpdatePolicyGroupsRequest

`func NewWafUpdatePolicyGroupsRequest() *WafUpdatePolicyGroupsRequest`

NewWafUpdatePolicyGroupsRequest instantiates a new WafUpdatePolicyGroupsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafUpdatePolicyGroupsRequestWithDefaults

`func NewWafUpdatePolicyGroupsRequestWithDefaults() *WafUpdatePolicyGroupsRequest`

NewWafUpdatePolicyGroupsRequestWithDefaults instantiates a new WafUpdatePolicyGroupsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyGroups

`func (o *WafUpdatePolicyGroupsRequest) GetPolicyGroups() []WafPolicyGroup`

GetPolicyGroups returns the PolicyGroups field if non-nil, zero value otherwise.

### GetPolicyGroupsOk

`func (o *WafUpdatePolicyGroupsRequest) GetPolicyGroupsOk() (*[]WafPolicyGroup, bool)`

GetPolicyGroupsOk returns a tuple with the PolicyGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyGroups

`func (o *WafUpdatePolicyGroupsRequest) SetPolicyGroups(v []WafPolicyGroup)`

SetPolicyGroups sets PolicyGroups field to given value.

### HasPolicyGroups

`func (o *WafUpdatePolicyGroupsRequest) HasPolicyGroups() bool`

HasPolicyGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


