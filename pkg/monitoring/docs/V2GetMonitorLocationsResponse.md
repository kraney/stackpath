# V2GetMonitorLocationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Monitor** | Pointer to [**V2Monitor**](v2Monitor.md) |  | [optional] 
**Locations** | Pointer to [**[]V2LocationWithAnalytics**](v2LocationWithAnalytics.md) | The requested locations of the monitor. | [optional] 

## Methods

### NewV2GetMonitorLocationsResponse

`func NewV2GetMonitorLocationsResponse() *V2GetMonitorLocationsResponse`

NewV2GetMonitorLocationsResponse instantiates a new V2GetMonitorLocationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2GetMonitorLocationsResponseWithDefaults

`func NewV2GetMonitorLocationsResponseWithDefaults() *V2GetMonitorLocationsResponse`

NewV2GetMonitorLocationsResponseWithDefaults instantiates a new V2GetMonitorLocationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitor

`func (o *V2GetMonitorLocationsResponse) GetMonitor() V2Monitor`

GetMonitor returns the Monitor field if non-nil, zero value otherwise.

### GetMonitorOk

`func (o *V2GetMonitorLocationsResponse) GetMonitorOk() (*V2Monitor, bool)`

GetMonitorOk returns a tuple with the Monitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitor

`func (o *V2GetMonitorLocationsResponse) SetMonitor(v V2Monitor)`

SetMonitor sets Monitor field to given value.

### HasMonitor

`func (o *V2GetMonitorLocationsResponse) HasMonitor() bool`

HasMonitor returns a boolean if a field has been set.

### GetLocations

`func (o *V2GetMonitorLocationsResponse) GetLocations() []V2LocationWithAnalytics`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *V2GetMonitorLocationsResponse) GetLocationsOk() (*[]V2LocationWithAnalytics, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *V2GetMonitorLocationsResponse) SetLocations(v []V2LocationWithAnalytics)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *V2GetMonitorLocationsResponse) HasLocations() bool`

HasLocations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


