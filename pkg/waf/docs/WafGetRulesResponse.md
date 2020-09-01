# WafGetRulesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageInfo** | Pointer to [**PaginationPageInfo**](paginationPageInfo.md) |  | [optional] 
**Rules** | Pointer to [**[]WafRule**](wafRule.md) | The requested WAF rules | [optional] 

## Methods

### NewWafGetRulesResponse

`func NewWafGetRulesResponse() *WafGetRulesResponse`

NewWafGetRulesResponse instantiates a new WafGetRulesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafGetRulesResponseWithDefaults

`func NewWafGetRulesResponseWithDefaults() *WafGetRulesResponse`

NewWafGetRulesResponseWithDefaults instantiates a new WafGetRulesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageInfo

`func (o *WafGetRulesResponse) GetPageInfo() PaginationPageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *WafGetRulesResponse) GetPageInfoOk() (*PaginationPageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *WafGetRulesResponse) SetPageInfo(v PaginationPageInfo)`

SetPageInfo sets PageInfo field to given value.

### HasPageInfo

`func (o *WafGetRulesResponse) HasPageInfo() bool`

HasPageInfo returns a boolean if a field has been set.

### GetRules

`func (o *WafGetRulesResponse) GetRules() []WafRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *WafGetRulesResponse) GetRulesOk() (*[]WafRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *WafGetRulesResponse) SetRules(v []WafRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *WafGetRulesResponse) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


