# V2LocationWithAnalytics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | Pointer to [**Monitoringv2Location**](monitoringv2Location.md) |  | [optional] 
**StatusCode** | Pointer to **int32** | The HTTP status code of the most recent monitoring check from the location. | [optional] 
**ResponseTime** | Pointer to **string** | The response time of the most recent monitoring check from the location. | [optional] 

## Methods

### NewV2LocationWithAnalytics

`func NewV2LocationWithAnalytics() *V2LocationWithAnalytics`

NewV2LocationWithAnalytics instantiates a new V2LocationWithAnalytics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2LocationWithAnalyticsWithDefaults

`func NewV2LocationWithAnalyticsWithDefaults() *V2LocationWithAnalytics`

NewV2LocationWithAnalyticsWithDefaults instantiates a new V2LocationWithAnalytics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *V2LocationWithAnalytics) GetLocation() Monitoringv2Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *V2LocationWithAnalytics) GetLocationOk() (*Monitoringv2Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *V2LocationWithAnalytics) SetLocation(v Monitoringv2Location)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *V2LocationWithAnalytics) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetStatusCode

`func (o *V2LocationWithAnalytics) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *V2LocationWithAnalytics) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *V2LocationWithAnalytics) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *V2LocationWithAnalytics) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetResponseTime

`func (o *V2LocationWithAnalytics) GetResponseTime() string`

GetResponseTime returns the ResponseTime field if non-nil, zero value otherwise.

### GetResponseTimeOk

`func (o *V2LocationWithAnalytics) GetResponseTimeOk() (*string, bool)`

GetResponseTimeOk returns a tuple with the ResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTime

`func (o *V2LocationWithAnalytics) SetResponseTime(v string)`

SetResponseTime sets ResponseTime field to given value.

### HasResponseTime

`func (o *V2LocationWithAnalytics) HasResponseTime() bool`

HasResponseTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


