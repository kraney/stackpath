# UpdateMonitorRequestPatchTcpConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** | The host address of the service to monitor. | [optional] 
**Port** | Pointer to **int32** | The port of the service to monitor. | [optional] 
**Data** | Pointer to **string** | Base64 encoded TCP data to send to the monitored service | [optional] 

## Methods

### NewUpdateMonitorRequestPatchTcpConfiguration

`func NewUpdateMonitorRequestPatchTcpConfiguration() *UpdateMonitorRequestPatchTcpConfiguration`

NewUpdateMonitorRequestPatchTcpConfiguration instantiates a new UpdateMonitorRequestPatchTcpConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMonitorRequestPatchTcpConfigurationWithDefaults

`func NewUpdateMonitorRequestPatchTcpConfigurationWithDefaults() *UpdateMonitorRequestPatchTcpConfiguration`

NewUpdateMonitorRequestPatchTcpConfigurationWithDefaults instantiates a new UpdateMonitorRequestPatchTcpConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *UpdateMonitorRequestPatchTcpConfiguration) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *UpdateMonitorRequestPatchTcpConfiguration) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *UpdateMonitorRequestPatchTcpConfiguration) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *UpdateMonitorRequestPatchTcpConfiguration) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *UpdateMonitorRequestPatchTcpConfiguration) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *UpdateMonitorRequestPatchTcpConfiguration) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *UpdateMonitorRequestPatchTcpConfiguration) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *UpdateMonitorRequestPatchTcpConfiguration) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetData

`func (o *UpdateMonitorRequestPatchTcpConfiguration) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UpdateMonitorRequestPatchTcpConfiguration) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UpdateMonitorRequestPatchTcpConfiguration) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *UpdateMonitorRequestPatchTcpConfiguration) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


