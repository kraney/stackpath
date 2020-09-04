# GetMetricsResponseMetricSeries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of data in the analytics set | [optional] 
**Key** | Pointer to **string** | A graphable key for the type of data in the analytics set | [optional] 
**Metrics** | Pointer to **[]string** | Descriptions of the CDN metrics&#39; samples | [optional] 
**Samples** | Pointer to [**[]GetMetricsResponseMetricSample**](GetMetricsResponseMetricSample.md) | The CDN metrics&#39; data points | [optional] 

## Methods

### NewGetMetricsResponseMetricSeries

`func NewGetMetricsResponseMetricSeries() *GetMetricsResponseMetricSeries`

NewGetMetricsResponseMetricSeries instantiates a new GetMetricsResponseMetricSeries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMetricsResponseMetricSeriesWithDefaults

`func NewGetMetricsResponseMetricSeriesWithDefaults() *GetMetricsResponseMetricSeries`

NewGetMetricsResponseMetricSeriesWithDefaults instantiates a new GetMetricsResponseMetricSeries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GetMetricsResponseMetricSeries) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetMetricsResponseMetricSeries) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetMetricsResponseMetricSeries) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetMetricsResponseMetricSeries) HasType() bool`

HasType returns a boolean if a field has been set.

### GetKey

`func (o *GetMetricsResponseMetricSeries) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *GetMetricsResponseMetricSeries) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *GetMetricsResponseMetricSeries) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *GetMetricsResponseMetricSeries) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMetrics

`func (o *GetMetricsResponseMetricSeries) GetMetrics() []string`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *GetMetricsResponseMetricSeries) GetMetricsOk() (*[]string, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *GetMetricsResponseMetricSeries) SetMetrics(v []string)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *GetMetricsResponseMetricSeries) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetSamples

`func (o *GetMetricsResponseMetricSeries) GetSamples() []GetMetricsResponseMetricSample`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *GetMetricsResponseMetricSeries) GetSamplesOk() (*[]GetMetricsResponseMetricSample, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *GetMetricsResponseMetricSeries) SetSamples(v []GetMetricsResponseMetricSample)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *GetMetricsResponseMetricSeries) HasSamples() bool`

HasSamples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


