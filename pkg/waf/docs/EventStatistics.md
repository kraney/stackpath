# EventStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to [**StatisticsMetrics**](StatisticsMetrics.md) |  | [optional] 
**Blocked** | Pointer to [**StatisticsMetrics**](StatisticsMetrics.md) |  | [optional] 

## Methods

### NewEventStatistics

`func NewEventStatistics() *EventStatistics`

NewEventStatistics instantiates a new EventStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventStatisticsWithDefaults

`func NewEventStatisticsWithDefaults() *EventStatistics`

NewEventStatisticsWithDefaults instantiates a new EventStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *EventStatistics) GetCount() StatisticsMetrics`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *EventStatistics) GetCountOk() (*StatisticsMetrics, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *EventStatistics) SetCount(v StatisticsMetrics)`

SetCount sets Count field to given value.

### HasCount

`func (o *EventStatistics) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetBlocked

`func (o *EventStatistics) GetBlocked() StatisticsMetrics`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *EventStatistics) GetBlockedOk() (*StatisticsMetrics, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *EventStatistics) SetBlocked(v StatisticsMetrics)`

SetBlocked sets Blocked field to given value.

### HasBlocked

`func (o *EventStatistics) HasBlocked() bool`

HasBlocked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


