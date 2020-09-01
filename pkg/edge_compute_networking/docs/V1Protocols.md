# V1Protocols

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tcp** | Pointer to [**V1ProtocolTcp**](v1ProtocolTcp.md) |  | [optional] 
**Udp** | Pointer to [**V1ProtocolUdp**](v1ProtocolUdp.md) |  | [optional] 
**TcpUdp** | Pointer to [**V1ProtocolTcpUdp**](v1ProtocolTcpUdp.md) |  | [optional] 
**Icmp** | Pointer to [**map[string]interface{}**](.md) | ICMP protocol matching | [optional] 
**Ah** | Pointer to [**map[string]interface{}**](.md) | Authentication Header (AH) protocol matching | [optional] 
**Esp** | Pointer to [**map[string]interface{}**](.md) | Encapsulating Security Payload (ESP) protocol matching | [optional] 
**Gre** | Pointer to [**map[string]interface{}**](.md) | Generic Routing Encapsulation (GRE) protocol matching | [optional] 

## Methods

### NewV1Protocols

`func NewV1Protocols() *V1Protocols`

NewV1Protocols instantiates a new V1Protocols object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ProtocolsWithDefaults

`func NewV1ProtocolsWithDefaults() *V1Protocols`

NewV1ProtocolsWithDefaults instantiates a new V1Protocols object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTcp

`func (o *V1Protocols) GetTcp() V1ProtocolTcp`

GetTcp returns the Tcp field if non-nil, zero value otherwise.

### GetTcpOk

`func (o *V1Protocols) GetTcpOk() (*V1ProtocolTcp, bool)`

GetTcpOk returns a tuple with the Tcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcp

`func (o *V1Protocols) SetTcp(v V1ProtocolTcp)`

SetTcp sets Tcp field to given value.

### HasTcp

`func (o *V1Protocols) HasTcp() bool`

HasTcp returns a boolean if a field has been set.

### GetUdp

`func (o *V1Protocols) GetUdp() V1ProtocolUdp`

GetUdp returns the Udp field if non-nil, zero value otherwise.

### GetUdpOk

`func (o *V1Protocols) GetUdpOk() (*V1ProtocolUdp, bool)`

GetUdpOk returns a tuple with the Udp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdp

`func (o *V1Protocols) SetUdp(v V1ProtocolUdp)`

SetUdp sets Udp field to given value.

### HasUdp

`func (o *V1Protocols) HasUdp() bool`

HasUdp returns a boolean if a field has been set.

### GetTcpUdp

`func (o *V1Protocols) GetTcpUdp() V1ProtocolTcpUdp`

GetTcpUdp returns the TcpUdp field if non-nil, zero value otherwise.

### GetTcpUdpOk

`func (o *V1Protocols) GetTcpUdpOk() (*V1ProtocolTcpUdp, bool)`

GetTcpUdpOk returns a tuple with the TcpUdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpUdp

`func (o *V1Protocols) SetTcpUdp(v V1ProtocolTcpUdp)`

SetTcpUdp sets TcpUdp field to given value.

### HasTcpUdp

`func (o *V1Protocols) HasTcpUdp() bool`

HasTcpUdp returns a boolean if a field has been set.

### GetIcmp

`func (o *V1Protocols) GetIcmp() map[string]interface{}`

GetIcmp returns the Icmp field if non-nil, zero value otherwise.

### GetIcmpOk

`func (o *V1Protocols) GetIcmpOk() (*map[string]interface{}, bool)`

GetIcmpOk returns a tuple with the Icmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcmp

`func (o *V1Protocols) SetIcmp(v map[string]interface{})`

SetIcmp sets Icmp field to given value.

### HasIcmp

`func (o *V1Protocols) HasIcmp() bool`

HasIcmp returns a boolean if a field has been set.

### GetAh

`func (o *V1Protocols) GetAh() map[string]interface{}`

GetAh returns the Ah field if non-nil, zero value otherwise.

### GetAhOk

`func (o *V1Protocols) GetAhOk() (*map[string]interface{}, bool)`

GetAhOk returns a tuple with the Ah field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAh

`func (o *V1Protocols) SetAh(v map[string]interface{})`

SetAh sets Ah field to given value.

### HasAh

`func (o *V1Protocols) HasAh() bool`

HasAh returns a boolean if a field has been set.

### GetEsp

`func (o *V1Protocols) GetEsp() map[string]interface{}`

GetEsp returns the Esp field if non-nil, zero value otherwise.

### GetEspOk

`func (o *V1Protocols) GetEspOk() (*map[string]interface{}, bool)`

GetEspOk returns a tuple with the Esp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsp

`func (o *V1Protocols) SetEsp(v map[string]interface{})`

SetEsp sets Esp field to given value.

### HasEsp

`func (o *V1Protocols) HasEsp() bool`

HasEsp returns a boolean if a field has been set.

### GetGre

`func (o *V1Protocols) GetGre() map[string]interface{}`

GetGre returns the Gre field if non-nil, zero value otherwise.

### GetGreOk

`func (o *V1Protocols) GetGreOk() (*map[string]interface{}, bool)`

GetGreOk returns a tuple with the Gre field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGre

`func (o *V1Protocols) SetGre(v map[string]interface{})`

SetGre sets Gre field to given value.

### HasGre

`func (o *V1Protocols) HasGre() bool`

HasGre returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


