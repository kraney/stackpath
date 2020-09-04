# V2UpdateMonitorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The monitor&#39;s name | [optional] 
**Locations** | Pointer to [**UpdateMonitorRequestLocationsValue**](UpdateMonitorRequestLocationsValue.md) |  | [optional] 
**Interval** | Pointer to **string** | How often a monitor should be checked. | [optional] 
**Timeout** | Pointer to **string** | The length of time a monitor check should wait before timing out. | [optional] 
**IpVersion** | Pointer to [**UpdateMonitorRequestIpVersionValue**](UpdateMonitorRequestIpVersionValue.md) |  | [optional] 
**Http** | Pointer to [**UpdateMonitorRequestPatchHttpConfiguration**](UpdateMonitorRequestPatchHttpConfiguration.md) |  | [optional] 
**Tcp** | Pointer to [**UpdateMonitorRequestPatchTcpConfiguration**](UpdateMonitorRequestPatchTcpConfiguration.md) |  | [optional] 
**Enabled** | Pointer to **bool** | Whether a monitor is enabled or not. | [optional] 

## Methods

### NewV2UpdateMonitorRequest

`func NewV2UpdateMonitorRequest() *V2UpdateMonitorRequest`

NewV2UpdateMonitorRequest instantiates a new V2UpdateMonitorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2UpdateMonitorRequestWithDefaults

`func NewV2UpdateMonitorRequestWithDefaults() *V2UpdateMonitorRequest`

NewV2UpdateMonitorRequestWithDefaults instantiates a new V2UpdateMonitorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *V2UpdateMonitorRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2UpdateMonitorRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2UpdateMonitorRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V2UpdateMonitorRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLocations

`func (o *V2UpdateMonitorRequest) GetLocations() UpdateMonitorRequestLocationsValue`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *V2UpdateMonitorRequest) GetLocationsOk() (*UpdateMonitorRequestLocationsValue, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *V2UpdateMonitorRequest) SetLocations(v UpdateMonitorRequestLocationsValue)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *V2UpdateMonitorRequest) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetInterval

`func (o *V2UpdateMonitorRequest) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *V2UpdateMonitorRequest) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *V2UpdateMonitorRequest) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *V2UpdateMonitorRequest) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetTimeout

`func (o *V2UpdateMonitorRequest) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *V2UpdateMonitorRequest) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *V2UpdateMonitorRequest) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *V2UpdateMonitorRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetIpVersion

`func (o *V2UpdateMonitorRequest) GetIpVersion() UpdateMonitorRequestIpVersionValue`

GetIpVersion returns the IpVersion field if non-nil, zero value otherwise.

### GetIpVersionOk

`func (o *V2UpdateMonitorRequest) GetIpVersionOk() (*UpdateMonitorRequestIpVersionValue, bool)`

GetIpVersionOk returns a tuple with the IpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpVersion

`func (o *V2UpdateMonitorRequest) SetIpVersion(v UpdateMonitorRequestIpVersionValue)`

SetIpVersion sets IpVersion field to given value.

### HasIpVersion

`func (o *V2UpdateMonitorRequest) HasIpVersion() bool`

HasIpVersion returns a boolean if a field has been set.

### GetHttp

`func (o *V2UpdateMonitorRequest) GetHttp() UpdateMonitorRequestPatchHttpConfiguration`

GetHttp returns the Http field if non-nil, zero value otherwise.

### GetHttpOk

`func (o *V2UpdateMonitorRequest) GetHttpOk() (*UpdateMonitorRequestPatchHttpConfiguration, bool)`

GetHttpOk returns a tuple with the Http field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp

`func (o *V2UpdateMonitorRequest) SetHttp(v UpdateMonitorRequestPatchHttpConfiguration)`

SetHttp sets Http field to given value.

### HasHttp

`func (o *V2UpdateMonitorRequest) HasHttp() bool`

HasHttp returns a boolean if a field has been set.

### GetTcp

`func (o *V2UpdateMonitorRequest) GetTcp() UpdateMonitorRequestPatchTcpConfiguration`

GetTcp returns the Tcp field if non-nil, zero value otherwise.

### GetTcpOk

`func (o *V2UpdateMonitorRequest) GetTcpOk() (*UpdateMonitorRequestPatchTcpConfiguration, bool)`

GetTcpOk returns a tuple with the Tcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcp

`func (o *V2UpdateMonitorRequest) SetTcp(v UpdateMonitorRequestPatchTcpConfiguration)`

SetTcp sets Tcp field to given value.

### HasTcp

`func (o *V2UpdateMonitorRequest) HasTcp() bool`

HasTcp returns a boolean if a field has been set.

### GetEnabled

`func (o *V2UpdateMonitorRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *V2UpdateMonitorRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *V2UpdateMonitorRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *V2UpdateMonitorRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


