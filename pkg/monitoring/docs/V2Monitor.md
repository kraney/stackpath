# V2Monitor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A monitor&#39;s unique identifier | [optional] 
**Name** | Pointer to **string** | The name of a monitor  A monitor&#39;s name | [optional] 
**CurrentStatus** | Pointer to [**V2MonitorStatus**](v2MonitorStatus.md) |  | [optional] [default to "UNKNOWN"]
**UptimePercentage** | Pointer to **float64** | The percent of time a monitored service has had a status of UP or SLOW in the last hour | [optional] 
**AvgResponseTime** | Pointer to **string** | The average response time for a monitor | [optional] 
**Locations** | Pointer to [**[]Monitoringv2Location**](monitoringv2Location.md) | The locations a service is monitored from. | [optional] 
**Interval** | Pointer to **string** | The amount of time between each monitor check.  By default each service monitor will be checked every five minutes. | [optional] 
**Timeout** | Pointer to **string** | Then length of time to wait until the monitor check times out. | [optional] 
**IpVersion** | Pointer to [**V2IpVersion**](v2IpVersion.md) |  | [optional] [default to "IPV4"]
**Http** | Pointer to [**V2HttpConfiguration**](v2HttpConfiguration.md) |  | [optional] 
**Tcp** | Pointer to [**V2TcpConfiguration**](v2TcpConfiguration.md) |  | [optional] 
**CreatedAt** | Pointer to [**time.Time**](time.Time.md) | The date a monitor was created. | [optional] 
**UpdatedAt** | Pointer to [**time.Time**](time.Time.md) | The date a monitor was last updated. | [optional] 
**Enabled** | Pointer to **bool** | Whether or not a monitor is enabled. | [optional] 

## Methods

### NewV2Monitor

`func NewV2Monitor() *V2Monitor`

NewV2Monitor instantiates a new V2Monitor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2MonitorWithDefaults

`func NewV2MonitorWithDefaults() *V2Monitor`

NewV2MonitorWithDefaults instantiates a new V2Monitor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V2Monitor) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V2Monitor) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V2Monitor) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *V2Monitor) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *V2Monitor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2Monitor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2Monitor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V2Monitor) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCurrentStatus

`func (o *V2Monitor) GetCurrentStatus() V2MonitorStatus`

GetCurrentStatus returns the CurrentStatus field if non-nil, zero value otherwise.

### GetCurrentStatusOk

`func (o *V2Monitor) GetCurrentStatusOk() (*V2MonitorStatus, bool)`

GetCurrentStatusOk returns a tuple with the CurrentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStatus

`func (o *V2Monitor) SetCurrentStatus(v V2MonitorStatus)`

SetCurrentStatus sets CurrentStatus field to given value.

### HasCurrentStatus

`func (o *V2Monitor) HasCurrentStatus() bool`

HasCurrentStatus returns a boolean if a field has been set.

### GetUptimePercentage

`func (o *V2Monitor) GetUptimePercentage() float64`

GetUptimePercentage returns the UptimePercentage field if non-nil, zero value otherwise.

### GetUptimePercentageOk

`func (o *V2Monitor) GetUptimePercentageOk() (*float64, bool)`

GetUptimePercentageOk returns a tuple with the UptimePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptimePercentage

`func (o *V2Monitor) SetUptimePercentage(v float64)`

SetUptimePercentage sets UptimePercentage field to given value.

### HasUptimePercentage

`func (o *V2Monitor) HasUptimePercentage() bool`

HasUptimePercentage returns a boolean if a field has been set.

### GetAvgResponseTime

`func (o *V2Monitor) GetAvgResponseTime() string`

GetAvgResponseTime returns the AvgResponseTime field if non-nil, zero value otherwise.

### GetAvgResponseTimeOk

`func (o *V2Monitor) GetAvgResponseTimeOk() (*string, bool)`

GetAvgResponseTimeOk returns a tuple with the AvgResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgResponseTime

`func (o *V2Monitor) SetAvgResponseTime(v string)`

SetAvgResponseTime sets AvgResponseTime field to given value.

### HasAvgResponseTime

`func (o *V2Monitor) HasAvgResponseTime() bool`

HasAvgResponseTime returns a boolean if a field has been set.

### GetLocations

`func (o *V2Monitor) GetLocations() []Monitoringv2Location`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *V2Monitor) GetLocationsOk() (*[]Monitoringv2Location, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *V2Monitor) SetLocations(v []Monitoringv2Location)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *V2Monitor) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetInterval

`func (o *V2Monitor) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *V2Monitor) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *V2Monitor) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *V2Monitor) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetTimeout

`func (o *V2Monitor) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *V2Monitor) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *V2Monitor) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *V2Monitor) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetIpVersion

`func (o *V2Monitor) GetIpVersion() V2IpVersion`

GetIpVersion returns the IpVersion field if non-nil, zero value otherwise.

### GetIpVersionOk

`func (o *V2Monitor) GetIpVersionOk() (*V2IpVersion, bool)`

GetIpVersionOk returns a tuple with the IpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpVersion

`func (o *V2Monitor) SetIpVersion(v V2IpVersion)`

SetIpVersion sets IpVersion field to given value.

### HasIpVersion

`func (o *V2Monitor) HasIpVersion() bool`

HasIpVersion returns a boolean if a field has been set.

### GetHttp

`func (o *V2Monitor) GetHttp() V2HttpConfiguration`

GetHttp returns the Http field if non-nil, zero value otherwise.

### GetHttpOk

`func (o *V2Monitor) GetHttpOk() (*V2HttpConfiguration, bool)`

GetHttpOk returns a tuple with the Http field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp

`func (o *V2Monitor) SetHttp(v V2HttpConfiguration)`

SetHttp sets Http field to given value.

### HasHttp

`func (o *V2Monitor) HasHttp() bool`

HasHttp returns a boolean if a field has been set.

### GetTcp

`func (o *V2Monitor) GetTcp() V2TcpConfiguration`

GetTcp returns the Tcp field if non-nil, zero value otherwise.

### GetTcpOk

`func (o *V2Monitor) GetTcpOk() (*V2TcpConfiguration, bool)`

GetTcpOk returns a tuple with the Tcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcp

`func (o *V2Monitor) SetTcp(v V2TcpConfiguration)`

SetTcp sets Tcp field to given value.

### HasTcp

`func (o *V2Monitor) HasTcp() bool`

HasTcp returns a boolean if a field has been set.

### GetCreatedAt

`func (o *V2Monitor) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *V2Monitor) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *V2Monitor) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *V2Monitor) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *V2Monitor) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *V2Monitor) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *V2Monitor) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *V2Monitor) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetEnabled

`func (o *V2Monitor) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *V2Monitor) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *V2Monitor) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *V2Monitor) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


