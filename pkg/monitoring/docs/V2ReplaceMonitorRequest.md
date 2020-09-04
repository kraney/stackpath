# V2ReplaceMonitorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the monitor. | [optional] 
**Locations** | Pointer to **[]string** | The list of locations a monitor is checked from. | [optional] 
**Interval** | Pointer to **string** | How often a monitor is checked. | [optional] 
**Timeout** | Pointer to **string** | The amount of time to wait before a monitor check times out. | [optional] 
**IpVersion** | Pointer to [**V2IpVersion**](v2IpVersion.md) |  | [optional] [default to "IPV4"]
**Http** | Pointer to [**V2HttpConfiguration**](v2HttpConfiguration.md) |  | [optional] 
**Tcp** | Pointer to [**V2TcpConfiguration**](v2TcpConfiguration.md) |  | [optional] 
**Enabled** | Pointer to **bool** | Whether a monitor is enabled or not. | [optional] 

## Methods

### NewV2ReplaceMonitorRequest

`func NewV2ReplaceMonitorRequest() *V2ReplaceMonitorRequest`

NewV2ReplaceMonitorRequest instantiates a new V2ReplaceMonitorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2ReplaceMonitorRequestWithDefaults

`func NewV2ReplaceMonitorRequestWithDefaults() *V2ReplaceMonitorRequest`

NewV2ReplaceMonitorRequestWithDefaults instantiates a new V2ReplaceMonitorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *V2ReplaceMonitorRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2ReplaceMonitorRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2ReplaceMonitorRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V2ReplaceMonitorRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLocations

`func (o *V2ReplaceMonitorRequest) GetLocations() []string`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *V2ReplaceMonitorRequest) GetLocationsOk() (*[]string, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *V2ReplaceMonitorRequest) SetLocations(v []string)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *V2ReplaceMonitorRequest) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetInterval

`func (o *V2ReplaceMonitorRequest) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *V2ReplaceMonitorRequest) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *V2ReplaceMonitorRequest) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *V2ReplaceMonitorRequest) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetTimeout

`func (o *V2ReplaceMonitorRequest) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *V2ReplaceMonitorRequest) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *V2ReplaceMonitorRequest) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *V2ReplaceMonitorRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetIpVersion

`func (o *V2ReplaceMonitorRequest) GetIpVersion() V2IpVersion`

GetIpVersion returns the IpVersion field if non-nil, zero value otherwise.

### GetIpVersionOk

`func (o *V2ReplaceMonitorRequest) GetIpVersionOk() (*V2IpVersion, bool)`

GetIpVersionOk returns a tuple with the IpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpVersion

`func (o *V2ReplaceMonitorRequest) SetIpVersion(v V2IpVersion)`

SetIpVersion sets IpVersion field to given value.

### HasIpVersion

`func (o *V2ReplaceMonitorRequest) HasIpVersion() bool`

HasIpVersion returns a boolean if a field has been set.

### GetHttp

`func (o *V2ReplaceMonitorRequest) GetHttp() V2HttpConfiguration`

GetHttp returns the Http field if non-nil, zero value otherwise.

### GetHttpOk

`func (o *V2ReplaceMonitorRequest) GetHttpOk() (*V2HttpConfiguration, bool)`

GetHttpOk returns a tuple with the Http field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp

`func (o *V2ReplaceMonitorRequest) SetHttp(v V2HttpConfiguration)`

SetHttp sets Http field to given value.

### HasHttp

`func (o *V2ReplaceMonitorRequest) HasHttp() bool`

HasHttp returns a boolean if a field has been set.

### GetTcp

`func (o *V2ReplaceMonitorRequest) GetTcp() V2TcpConfiguration`

GetTcp returns the Tcp field if non-nil, zero value otherwise.

### GetTcpOk

`func (o *V2ReplaceMonitorRequest) GetTcpOk() (*V2TcpConfiguration, bool)`

GetTcpOk returns a tuple with the Tcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcp

`func (o *V2ReplaceMonitorRequest) SetTcp(v V2TcpConfiguration)`

SetTcp sets Tcp field to given value.

### HasTcp

`func (o *V2ReplaceMonitorRequest) HasTcp() bool`

HasTcp returns a boolean if a field has been set.

### GetEnabled

`func (o *V2ReplaceMonitorRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *V2ReplaceMonitorRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *V2ReplaceMonitorRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *V2ReplaceMonitorRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


