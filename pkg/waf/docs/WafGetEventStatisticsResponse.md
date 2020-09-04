# WafGetEventStatisticsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statistics** | Pointer to [**EventStatistics**](EventStatistics.md) |  | [optional] 

## Methods

### NewWafGetEventStatisticsResponse

`func NewWafGetEventStatisticsResponse() *WafGetEventStatisticsResponse`

NewWafGetEventStatisticsResponse instantiates a new WafGetEventStatisticsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafGetEventStatisticsResponseWithDefaults

`func NewWafGetEventStatisticsResponseWithDefaults() *WafGetEventStatisticsResponse`

NewWafGetEventStatisticsResponseWithDefaults instantiates a new WafGetEventStatisticsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatistics

`func (o *WafGetEventStatisticsResponse) GetStatistics() EventStatistics`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *WafGetEventStatisticsResponse) GetStatisticsOk() (*EventStatistics, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *WafGetEventStatisticsResponse) SetStatistics(v EventStatistics)`

SetStatistics sets Statistics field to given value.

### HasStatistics

`func (o *WafGetEventStatisticsResponse) HasStatistics() bool`

HasStatistics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


