# V2CreateAlertConditionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | Pointer to [**V2AlertConditionMetric**](v2AlertConditionMetric.md) |  | [optional] [default to "STATUS"]
**Comparator** | Pointer to [**AlertConditionComparator**](AlertConditionComparator.md) |  | [optional] [default to "EQUAL"]
**Value** | Pointer to **string** | The value to check to determine if an alert should be generated. | [optional] 
**AlarmDelay** | Pointer to **int32** | The amount of time to wait before sending an alert.  This is useful to prevent alerts due to a flapping service. | [optional] 
**AlarmEvery** | Pointer to **int32** | How often an alert should be generated if an alert condition is still being met. | [optional] 
**AlarmUntilSilenced** | Pointer to **bool** | A flag that determines if an alert should continue to trigger until it is resolved. | [optional] 
**Enabled** | Pointer to **bool** | Whether an alert condition is enabled or not. | [optional] 
**MinimumLocations** | Pointer to **int32** | The minimum number of locations that the alert condition must be triggered for before an alert is generated.  This should always be less than or equal to the number of locations a monitor has. | [optional] 

## Methods

### NewV2CreateAlertConditionRequest

`func NewV2CreateAlertConditionRequest() *V2CreateAlertConditionRequest`

NewV2CreateAlertConditionRequest instantiates a new V2CreateAlertConditionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2CreateAlertConditionRequestWithDefaults

`func NewV2CreateAlertConditionRequestWithDefaults() *V2CreateAlertConditionRequest`

NewV2CreateAlertConditionRequestWithDefaults instantiates a new V2CreateAlertConditionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetric

`func (o *V2CreateAlertConditionRequest) GetMetric() V2AlertConditionMetric`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *V2CreateAlertConditionRequest) GetMetricOk() (*V2AlertConditionMetric, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *V2CreateAlertConditionRequest) SetMetric(v V2AlertConditionMetric)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *V2CreateAlertConditionRequest) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetComparator

`func (o *V2CreateAlertConditionRequest) GetComparator() AlertConditionComparator`

GetComparator returns the Comparator field if non-nil, zero value otherwise.

### GetComparatorOk

`func (o *V2CreateAlertConditionRequest) GetComparatorOk() (*AlertConditionComparator, bool)`

GetComparatorOk returns a tuple with the Comparator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparator

`func (o *V2CreateAlertConditionRequest) SetComparator(v AlertConditionComparator)`

SetComparator sets Comparator field to given value.

### HasComparator

`func (o *V2CreateAlertConditionRequest) HasComparator() bool`

HasComparator returns a boolean if a field has been set.

### GetValue

`func (o *V2CreateAlertConditionRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *V2CreateAlertConditionRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *V2CreateAlertConditionRequest) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *V2CreateAlertConditionRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetAlarmDelay

`func (o *V2CreateAlertConditionRequest) GetAlarmDelay() int32`

GetAlarmDelay returns the AlarmDelay field if non-nil, zero value otherwise.

### GetAlarmDelayOk

`func (o *V2CreateAlertConditionRequest) GetAlarmDelayOk() (*int32, bool)`

GetAlarmDelayOk returns a tuple with the AlarmDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmDelay

`func (o *V2CreateAlertConditionRequest) SetAlarmDelay(v int32)`

SetAlarmDelay sets AlarmDelay field to given value.

### HasAlarmDelay

`func (o *V2CreateAlertConditionRequest) HasAlarmDelay() bool`

HasAlarmDelay returns a boolean if a field has been set.

### GetAlarmEvery

`func (o *V2CreateAlertConditionRequest) GetAlarmEvery() int32`

GetAlarmEvery returns the AlarmEvery field if non-nil, zero value otherwise.

### GetAlarmEveryOk

`func (o *V2CreateAlertConditionRequest) GetAlarmEveryOk() (*int32, bool)`

GetAlarmEveryOk returns a tuple with the AlarmEvery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmEvery

`func (o *V2CreateAlertConditionRequest) SetAlarmEvery(v int32)`

SetAlarmEvery sets AlarmEvery field to given value.

### HasAlarmEvery

`func (o *V2CreateAlertConditionRequest) HasAlarmEvery() bool`

HasAlarmEvery returns a boolean if a field has been set.

### GetAlarmUntilSilenced

`func (o *V2CreateAlertConditionRequest) GetAlarmUntilSilenced() bool`

GetAlarmUntilSilenced returns the AlarmUntilSilenced field if non-nil, zero value otherwise.

### GetAlarmUntilSilencedOk

`func (o *V2CreateAlertConditionRequest) GetAlarmUntilSilencedOk() (*bool, bool)`

GetAlarmUntilSilencedOk returns a tuple with the AlarmUntilSilenced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmUntilSilenced

`func (o *V2CreateAlertConditionRequest) SetAlarmUntilSilenced(v bool)`

SetAlarmUntilSilenced sets AlarmUntilSilenced field to given value.

### HasAlarmUntilSilenced

`func (o *V2CreateAlertConditionRequest) HasAlarmUntilSilenced() bool`

HasAlarmUntilSilenced returns a boolean if a field has been set.

### GetEnabled

`func (o *V2CreateAlertConditionRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *V2CreateAlertConditionRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *V2CreateAlertConditionRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *V2CreateAlertConditionRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMinimumLocations

`func (o *V2CreateAlertConditionRequest) GetMinimumLocations() int32`

GetMinimumLocations returns the MinimumLocations field if non-nil, zero value otherwise.

### GetMinimumLocationsOk

`func (o *V2CreateAlertConditionRequest) GetMinimumLocationsOk() (*int32, bool)`

GetMinimumLocationsOk returns a tuple with the MinimumLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumLocations

`func (o *V2CreateAlertConditionRequest) SetMinimumLocations(v int32)`

SetMinimumLocations sets MinimumLocations field to given value.

### HasMinimumLocations

`func (o *V2CreateAlertConditionRequest) HasMinimumLocations() bool`

HasMinimumLocations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


