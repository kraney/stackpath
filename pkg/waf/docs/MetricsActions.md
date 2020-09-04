# MetricsActions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to [**WafRuleAction**](wafRuleAction.md) |  | [optional] [default to "BLOCK"]
**Count** | Pointer to **string** | The number of requests that resulted in that action | [optional] 

## Methods

### NewMetricsActions

`func NewMetricsActions() *MetricsActions`

NewMetricsActions instantiates a new MetricsActions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsActionsWithDefaults

`func NewMetricsActionsWithDefaults() *MetricsActions`

NewMetricsActionsWithDefaults instantiates a new MetricsActions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *MetricsActions) GetAction() WafRuleAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *MetricsActions) GetActionOk() (*WafRuleAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *MetricsActions) SetAction(v WafRuleAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *MetricsActions) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCount

`func (o *MetricsActions) GetCount() string`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *MetricsActions) GetCountOk() (*string, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *MetricsActions) SetCount(v string)`

SetCount sets Count field to given value.

### HasCount

`func (o *MetricsActions) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


