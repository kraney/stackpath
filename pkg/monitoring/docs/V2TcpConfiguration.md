# V2TcpConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** | The host address of the service to monitor | [optional] 
**Port** | Pointer to **int32** | The port of the service to monitor | [optional] 
**Data** | Pointer to **string** | Base64 encoded TCP data to send to the monitored service | [optional] 

## Methods

### NewV2TcpConfiguration

`func NewV2TcpConfiguration() *V2TcpConfiguration`

NewV2TcpConfiguration instantiates a new V2TcpConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2TcpConfigurationWithDefaults

`func NewV2TcpConfigurationWithDefaults() *V2TcpConfiguration`

NewV2TcpConfigurationWithDefaults instantiates a new V2TcpConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *V2TcpConfiguration) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *V2TcpConfiguration) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *V2TcpConfiguration) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *V2TcpConfiguration) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *V2TcpConfiguration) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *V2TcpConfiguration) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *V2TcpConfiguration) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *V2TcpConfiguration) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetData

`func (o *V2TcpConfiguration) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *V2TcpConfiguration) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *V2TcpConfiguration) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *V2TcpConfiguration) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


