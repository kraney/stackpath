# V1ProtocolTcpUdp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationPorts** | Pointer to **[]string** | List of destination ports to allow 1-65535 | [optional] 
**SourcePorts** | Pointer to **[]string** | List of source ports to allow 1-65535, defaults to 1000-65535 | [optional] 

## Methods

### NewV1ProtocolTcpUdp

`func NewV1ProtocolTcpUdp() *V1ProtocolTcpUdp`

NewV1ProtocolTcpUdp instantiates a new V1ProtocolTcpUdp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ProtocolTcpUdpWithDefaults

`func NewV1ProtocolTcpUdpWithDefaults() *V1ProtocolTcpUdp`

NewV1ProtocolTcpUdpWithDefaults instantiates a new V1ProtocolTcpUdp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationPorts

`func (o *V1ProtocolTcpUdp) GetDestinationPorts() []string`

GetDestinationPorts returns the DestinationPorts field if non-nil, zero value otherwise.

### GetDestinationPortsOk

`func (o *V1ProtocolTcpUdp) GetDestinationPortsOk() (*[]string, bool)`

GetDestinationPortsOk returns a tuple with the DestinationPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPorts

`func (o *V1ProtocolTcpUdp) SetDestinationPorts(v []string)`

SetDestinationPorts sets DestinationPorts field to given value.

### HasDestinationPorts

`func (o *V1ProtocolTcpUdp) HasDestinationPorts() bool`

HasDestinationPorts returns a boolean if a field has been set.

### GetSourcePorts

`func (o *V1ProtocolTcpUdp) GetSourcePorts() []string`

GetSourcePorts returns the SourcePorts field if non-nil, zero value otherwise.

### GetSourcePortsOk

`func (o *V1ProtocolTcpUdp) GetSourcePortsOk() (*[]string, bool)`

GetSourcePortsOk returns a tuple with the SourcePorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePorts

`func (o *V1ProtocolTcpUdp) SetSourcePorts(v []string)`

SetSourcePorts sets SourcePorts field to given value.

### HasSourcePorts

`func (o *V1ProtocolTcpUdp) HasSourcePorts() bool`

HasSourcePorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


