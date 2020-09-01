# V2AlertCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier for an alert condition. | [optional] 
**MonitorId** | Pointer to **string** | The monitor identifier associated with an alert condition. | [optional] 
**Metric** | Pointer to [**V2AlertConditionMetric**](v2AlertConditionMetric.md) |  | [optional] [default to "STATUS"]
**Comparator** | Pointer to [**AlertConditionComparator**](AlertConditionComparator.md) |  | [optional] [default to "EQUAL"]
**Value** | Pointer to **string** | The value to check to determine if an alert should be generated. | [optional] 
**AlarmDelay** | Pointer to **int32** | The amount of time to wait before sending an alert.  This is useful to prevent alerts due to a flapping service. | [optional] 
**AlarmEvery** | Pointer to **int32** | How often an alert should be generated if an alert condition is still being met. | [optional] 
**AlarmUntilSilenced** | Pointer to **bool** | A flag that determines if an alert should continue to trigger until it is resolved. | [optional] 
**Enabled** | Pointer to **bool** | Whether an alert condition is enabled or not. | [optional] 
**MinimumLocations** | Pointer to **int32** | The minimum number of locations that the alert condition must be triggered for before an alert is generated.  This should always be less than or equal to the number of locations a monitor has. | [optional] 
**CreatedAt** | Pointer to [**time.Time**](time.Time.md) | The date an alert condition was created. | [optional] 
**UpdatedAt** | Pointer to [**time.Time**](time.Time.md) | The date an alert condition was last updated. | [optional] 

## Methods

### NewV2AlertCondition

`func NewV2AlertCondition() *V2AlertCondition`

NewV2AlertCondition instantiates a new V2AlertCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2AlertConditionWithDefaults

`func NewV2AlertConditionWithDefaults() *V2AlertCondition`

NewV2AlertConditionWithDefaults instantiates a new V2AlertCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V2AlertCondition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V2AlertCondition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V2AlertCondition) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *V2AlertCondition) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMonitorId

`func (o *V2AlertCondition) GetMonitorId() string`

GetMonitorId returns the MonitorId field if non-nil, zero value otherwise.

### GetMonitorIdOk

`func (o *V2AlertCondition) GetMonitorIdOk() (*string, bool)`

GetMonitorIdOk returns a tuple with the MonitorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorId

`func (o *V2AlertCondition) SetMonitorId(v string)`

SetMonitorId sets MonitorId field to given value.

### HasMonitorId

`func (o *V2AlertCondition) HasMonitorId() bool`

HasMonitorId returns a boolean if a field has been set.

### GetMetric

`func (o *V2AlertCondition) GetMetric() V2AlertConditionMetric`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *V2AlertCondition) GetMetricOk() (*V2AlertConditionMetric, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *V2AlertCondition) SetMetric(v V2AlertConditionMetric)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *V2AlertCondition) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetComparator

`func (o *V2AlertCondition) GetComparator() AlertConditionComparator`

GetComparator returns the Comparator field if non-nil, zero value otherwise.

### GetComparatorOk

`func (o *V2AlertCondition) GetComparatorOk() (*AlertConditionComparator, bool)`

GetComparatorOk returns a tuple with the Comparator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparator

`func (o *V2AlertCondition) SetComparator(v AlertConditionComparator)`

SetComparator sets Comparator field to given value.

### HasComparator

`func (o *V2AlertCondition) HasComparator() bool`

HasComparator returns a boolean if a field has been set.

### GetValue

`func (o *V2AlertCondition) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *V2AlertCondition) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *V2AlertCondition) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *V2AlertCondition) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetAlarmDelay

`func (o *V2AlertCondition) GetAlarmDelay() int32`

GetAlarmDelay returns the AlarmDelay field if non-nil, zero value otherwise.

### GetAlarmDelayOk

`func (o *V2AlertCondition) GetAlarmDelayOk() (*int32, bool)`

GetAlarmDelayOk returns a tuple with the AlarmDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmDelay

`func (o *V2AlertCondition) SetAlarmDelay(v int32)`

SetAlarmDelay sets AlarmDelay field to given value.

### HasAlarmDelay

`func (o *V2AlertCondition) HasAlarmDelay() bool`

HasAlarmDelay returns a boolean if a field has been set.

### GetAlarmEvery

`func (o *V2AlertCondition) GetAlarmEvery() int32`

GetAlarmEvery returns the AlarmEvery field if non-nil, zero value otherwise.

### GetAlarmEveryOk

`func (o *V2AlertCondition) GetAlarmEveryOk() (*int32, bool)`

GetAlarmEveryOk returns a tuple with the AlarmEvery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmEvery

`func (o *V2AlertCondition) SetAlarmEvery(v int32)`

SetAlarmEvery sets AlarmEvery field to given value.

### HasAlarmEvery

`func (o *V2AlertCondition) HasAlarmEvery() bool`

HasAlarmEvery returns a boolean if a field has been set.

### GetAlarmUntilSilenced

`func (o *V2AlertCondition) GetAlarmUntilSilenced() bool`

GetAlarmUntilSilenced returns the AlarmUntilSilenced field if non-nil, zero value otherwise.

### GetAlarmUntilSilencedOk

`func (o *V2AlertCondition) GetAlarmUntilSilencedOk() (*bool, bool)`

GetAlarmUntilSilencedOk returns a tuple with the AlarmUntilSilenced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmUntilSilenced

`func (o *V2AlertCondition) SetAlarmUntilSilenced(v bool)`

SetAlarmUntilSilenced sets AlarmUntilSilenced field to given value.

### HasAlarmUntilSilenced

`func (o *V2AlertCondition) HasAlarmUntilSilenced() bool`

HasAlarmUntilSilenced returns a boolean if a field has been set.

### GetEnabled

`func (o *V2AlertCondition) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *V2AlertCondition) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *V2AlertCondition) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *V2AlertCondition) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMinimumLocations

`func (o *V2AlertCondition) GetMinimumLocations() int32`

GetMinimumLocations returns the MinimumLocations field if non-nil, zero value otherwise.

### GetMinimumLocationsOk

`func (o *V2AlertCondition) GetMinimumLocationsOk() (*int32, bool)`

GetMinimumLocationsOk returns a tuple with the MinimumLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumLocations

`func (o *V2AlertCondition) SetMinimumLocations(v int32)`

SetMinimumLocations sets MinimumLocations field to given value.

### HasMinimumLocations

`func (o *V2AlertCondition) HasMinimumLocations() bool`

HasMinimumLocations returns a boolean if a field has been set.

### GetCreatedAt

`func (o *V2AlertCondition) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *V2AlertCondition) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *V2AlertCondition) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *V2AlertCondition) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *V2AlertCondition) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *V2AlertCondition) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *V2AlertCondition) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *V2AlertCondition) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


