# V1MetricSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | Pointer to **string** | The metric to measure against | [optional] 
**AverageValue** | Pointer to **string** | The target average value for the metric measured across all instances within a deployment, expressed as a quantity. | [optional] 
**AverageUtilization** | Pointer to **int32** | The target average value for the metric measured across all instances within a deployment, expressed as a percentage of requested resources. | [optional] 

## Methods

### NewV1MetricSpec

`func NewV1MetricSpec() *V1MetricSpec`

NewV1MetricSpec instantiates a new V1MetricSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1MetricSpecWithDefaults

`func NewV1MetricSpecWithDefaults() *V1MetricSpec`

NewV1MetricSpecWithDefaults instantiates a new V1MetricSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetric

`func (o *V1MetricSpec) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *V1MetricSpec) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *V1MetricSpec) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *V1MetricSpec) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetAverageValue

`func (o *V1MetricSpec) GetAverageValue() string`

GetAverageValue returns the AverageValue field if non-nil, zero value otherwise.

### GetAverageValueOk

`func (o *V1MetricSpec) GetAverageValueOk() (*string, bool)`

GetAverageValueOk returns a tuple with the AverageValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageValue

`func (o *V1MetricSpec) SetAverageValue(v string)`

SetAverageValue sets AverageValue field to given value.

### HasAverageValue

`func (o *V1MetricSpec) HasAverageValue() bool`

HasAverageValue returns a boolean if a field has been set.

### GetAverageUtilization

`func (o *V1MetricSpec) GetAverageUtilization() int32`

GetAverageUtilization returns the AverageUtilization field if non-nil, zero value otherwise.

### GetAverageUtilizationOk

`func (o *V1MetricSpec) GetAverageUtilizationOk() (*int32, bool)`

GetAverageUtilizationOk returns a tuple with the AverageUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageUtilization

`func (o *V1MetricSpec) SetAverageUtilization(v int32)`

SetAverageUtilization sets AverageUtilization field to given value.

### HasAverageUtilization

`func (o *V1MetricSpec) HasAverageUtilization() bool`

HasAverageUtilization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


