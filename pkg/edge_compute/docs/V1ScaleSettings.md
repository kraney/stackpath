# V1ScaleSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metrics** | Pointer to [**[]V1MetricSpec**](v1MetricSpec.md) | Metrics to observe for invoking scaling events. | [optional] 

## Methods

### NewV1ScaleSettings

`func NewV1ScaleSettings() *V1ScaleSettings`

NewV1ScaleSettings instantiates a new V1ScaleSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ScaleSettingsWithDefaults

`func NewV1ScaleSettingsWithDefaults() *V1ScaleSettings`

NewV1ScaleSettingsWithDefaults instantiates a new V1ScaleSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetrics

`func (o *V1ScaleSettings) GetMetrics() []V1MetricSpec`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *V1ScaleSettings) GetMetricsOk() (*[]V1MetricSpec, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *V1ScaleSettings) SetMetrics(v []V1MetricSpec)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *V1ScaleSettings) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


