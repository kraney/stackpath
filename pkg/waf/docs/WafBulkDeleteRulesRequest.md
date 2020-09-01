# WafBulkDeleteRulesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleIds** | Pointer to **[]string** | The IDs of the rules to delete | [optional] 

## Methods

### NewWafBulkDeleteRulesRequest

`func NewWafBulkDeleteRulesRequest() *WafBulkDeleteRulesRequest`

NewWafBulkDeleteRulesRequest instantiates a new WafBulkDeleteRulesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafBulkDeleteRulesRequestWithDefaults

`func NewWafBulkDeleteRulesRequestWithDefaults() *WafBulkDeleteRulesRequest`

NewWafBulkDeleteRulesRequestWithDefaults instantiates a new WafBulkDeleteRulesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleIds

`func (o *WafBulkDeleteRulesRequest) GetRuleIds() []string`

GetRuleIds returns the RuleIds field if non-nil, zero value otherwise.

### GetRuleIdsOk

`func (o *WafBulkDeleteRulesRequest) GetRuleIdsOk() (*[]string, bool)`

GetRuleIdsOk returns a tuple with the RuleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleIds

`func (o *WafBulkDeleteRulesRequest) SetRuleIds(v []string)`

SetRuleIds sets RuleIds field to given value.

### HasRuleIds

`func (o *WafBulkDeleteRulesRequest) HasRuleIds() bool`

HasRuleIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


