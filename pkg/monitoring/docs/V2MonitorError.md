# V2MonitorError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MonitorId** | Pointer to **string** | The monitor identifier associated with the error. | [optional] 
**Locations** | Pointer to [**[]Monitoringv2Location**](monitoringv2Location.md) | A list of locations that have the error. | [optional] 
**Error** | Pointer to **string** | The error text for a monitor error. | [optional] 
**CreatedAt** | Pointer to [**time.Time**](time.Time.md) | The date a monitor error was created. | [optional] 

## Methods

### NewV2MonitorError

`func NewV2MonitorError() *V2MonitorError`

NewV2MonitorError instantiates a new V2MonitorError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2MonitorErrorWithDefaults

`func NewV2MonitorErrorWithDefaults() *V2MonitorError`

NewV2MonitorErrorWithDefaults instantiates a new V2MonitorError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitorId

`func (o *V2MonitorError) GetMonitorId() string`

GetMonitorId returns the MonitorId field if non-nil, zero value otherwise.

### GetMonitorIdOk

`func (o *V2MonitorError) GetMonitorIdOk() (*string, bool)`

GetMonitorIdOk returns a tuple with the MonitorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorId

`func (o *V2MonitorError) SetMonitorId(v string)`

SetMonitorId sets MonitorId field to given value.

### HasMonitorId

`func (o *V2MonitorError) HasMonitorId() bool`

HasMonitorId returns a boolean if a field has been set.

### GetLocations

`func (o *V2MonitorError) GetLocations() []Monitoringv2Location`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *V2MonitorError) GetLocationsOk() (*[]Monitoringv2Location, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *V2MonitorError) SetLocations(v []Monitoringv2Location)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *V2MonitorError) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetError

`func (o *V2MonitorError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V2MonitorError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V2MonitorError) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *V2MonitorError) HasError() bool`

HasError returns a boolean if a field has been set.

### GetCreatedAt

`func (o *V2MonitorError) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *V2MonitorError) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *V2MonitorError) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *V2MonitorError) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


