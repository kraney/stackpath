# V1InstancePort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | Pointer to **int32** | The network port | [optional] 
**Protocol** | Pointer to **string** | The network protocol for the port  Values are either \&quot;TCP\&quot; or \&quot;UDP\&quot;. Defaults to \&quot;TCP\&quot;. | [optional] 
**EnableImplicitNetworkPolicy** | Pointer to **bool** | Allow the internet to connect to the port  When true, this port will be given an implicit network policy of priority 65534 permitting 0.0.0.0/0 access to the port. This can be overridden by creating network policies of a higher priority to block the port. | [optional] 

## Methods

### NewV1InstancePort

`func NewV1InstancePort() *V1InstancePort`

NewV1InstancePort instantiates a new V1InstancePort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1InstancePortWithDefaults

`func NewV1InstancePortWithDefaults() *V1InstancePort`

NewV1InstancePortWithDefaults instantiates a new V1InstancePort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *V1InstancePort) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *V1InstancePort) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *V1InstancePort) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *V1InstancePort) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProtocol

`func (o *V1InstancePort) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *V1InstancePort) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *V1InstancePort) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *V1InstancePort) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetEnableImplicitNetworkPolicy

`func (o *V1InstancePort) GetEnableImplicitNetworkPolicy() bool`

GetEnableImplicitNetworkPolicy returns the EnableImplicitNetworkPolicy field if non-nil, zero value otherwise.

### GetEnableImplicitNetworkPolicyOk

`func (o *V1InstancePort) GetEnableImplicitNetworkPolicyOk() (*bool, bool)`

GetEnableImplicitNetworkPolicyOk returns a tuple with the EnableImplicitNetworkPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableImplicitNetworkPolicy

`func (o *V1InstancePort) SetEnableImplicitNetworkPolicy(v bool)`

SetEnableImplicitNetworkPolicy sets EnableImplicitNetworkPolicy field to given value.

### HasEnableImplicitNetworkPolicy

`func (o *V1InstancePort) HasEnableImplicitNetworkPolicy() bool`

HasEnableImplicitNetworkPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


