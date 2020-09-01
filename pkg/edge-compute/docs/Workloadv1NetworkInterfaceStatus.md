# Workloadv1NetworkInterfaceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | Pointer to **string** | A network interface&#39;s name | [optional] 
**IpAddress** | Pointer to **string** | A network interface&#39;s primary IP address | [optional] 
**IpAddressAliases** | Pointer to **[]string** | Additional IP addresses bound to a network interface | [optional] 
**Gateway** | Pointer to **string** | A network interface&#39;s gateway address | [optional] 

## Methods

### NewWorkloadv1NetworkInterfaceStatus

`func NewWorkloadv1NetworkInterfaceStatus() *Workloadv1NetworkInterfaceStatus`

NewWorkloadv1NetworkInterfaceStatus instantiates a new Workloadv1NetworkInterfaceStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadv1NetworkInterfaceStatusWithDefaults

`func NewWorkloadv1NetworkInterfaceStatusWithDefaults() *Workloadv1NetworkInterfaceStatus`

NewWorkloadv1NetworkInterfaceStatusWithDefaults instantiates a new Workloadv1NetworkInterfaceStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *Workloadv1NetworkInterfaceStatus) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *Workloadv1NetworkInterfaceStatus) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *Workloadv1NetworkInterfaceStatus) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *Workloadv1NetworkInterfaceStatus) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetIpAddress

`func (o *Workloadv1NetworkInterfaceStatus) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *Workloadv1NetworkInterfaceStatus) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *Workloadv1NetworkInterfaceStatus) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *Workloadv1NetworkInterfaceStatus) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetIpAddressAliases

`func (o *Workloadv1NetworkInterfaceStatus) GetIpAddressAliases() []string`

GetIpAddressAliases returns the IpAddressAliases field if non-nil, zero value otherwise.

### GetIpAddressAliasesOk

`func (o *Workloadv1NetworkInterfaceStatus) GetIpAddressAliasesOk() (*[]string, bool)`

GetIpAddressAliasesOk returns a tuple with the IpAddressAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddressAliases

`func (o *Workloadv1NetworkInterfaceStatus) SetIpAddressAliases(v []string)`

SetIpAddressAliases sets IpAddressAliases field to given value.

### HasIpAddressAliases

`func (o *Workloadv1NetworkInterfaceStatus) HasIpAddressAliases() bool`

HasIpAddressAliases returns a boolean if a field has been set.

### GetGateway

`func (o *Workloadv1NetworkInterfaceStatus) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *Workloadv1NetworkInterfaceStatus) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *Workloadv1NetworkInterfaceStatus) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *Workloadv1NetworkInterfaceStatus) HasGateway() bool`

HasGateway returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


