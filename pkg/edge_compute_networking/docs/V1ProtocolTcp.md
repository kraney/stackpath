# V1ProtocolTcp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationPorts** | Pointer to **[]string** | List of destination ports to allow 1-65535 | [optional] 
**SourcePorts** | Pointer to **[]string** | List of source ports to allow 1-65535, defaults to 1000-65535 | [optional] 

## Methods

### NewV1ProtocolTcp

`func NewV1ProtocolTcp() *V1ProtocolTcp`

NewV1ProtocolTcp instantiates a new V1ProtocolTcp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ProtocolTcpWithDefaults

`func NewV1ProtocolTcpWithDefaults() *V1ProtocolTcp`

NewV1ProtocolTcpWithDefaults instantiates a new V1ProtocolTcp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationPorts

`func (o *V1ProtocolTcp) GetDestinationPorts() []string`

GetDestinationPorts returns the DestinationPorts field if non-nil, zero value otherwise.

### GetDestinationPortsOk

`func (o *V1ProtocolTcp) GetDestinationPortsOk() (*[]string, bool)`

GetDestinationPortsOk returns a tuple with the DestinationPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPorts

`func (o *V1ProtocolTcp) SetDestinationPorts(v []string)`

SetDestinationPorts sets DestinationPorts field to given value.

### HasDestinationPorts

`func (o *V1ProtocolTcp) HasDestinationPorts() bool`

HasDestinationPorts returns a boolean if a field has been set.

### GetSourcePorts

`func (o *V1ProtocolTcp) GetSourcePorts() []string`

GetSourcePorts returns the SourcePorts field if non-nil, zero value otherwise.

### GetSourcePortsOk

`func (o *V1ProtocolTcp) GetSourcePortsOk() (*[]string, bool)`

GetSourcePortsOk returns a tuple with the SourcePorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePorts

`func (o *V1ProtocolTcp) SetSourcePorts(v []string)`

SetSourcePorts sets SourcePorts field to given value.

### HasSourcePorts

`func (o *V1ProtocolTcp) HasSourcePorts() bool`

HasSourcePorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


