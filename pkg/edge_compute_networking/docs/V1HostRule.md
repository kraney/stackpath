# V1HostRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpBlock** | Pointer to [**[]V1IpBlock**](v1IpBlock.md) | List of ip blocks | [optional] 
**InstanceSelectors** | Pointer to [**[]V1MatchExpression**](v1MatchExpression.md) | List of instance selectors | [optional] 
**NetworkSelectors** | Pointer to [**[]V1MatchExpression**](v1MatchExpression.md) | List of network selectors | [optional] 

## Methods

### NewV1HostRule

`func NewV1HostRule() *V1HostRule`

NewV1HostRule instantiates a new V1HostRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1HostRuleWithDefaults

`func NewV1HostRuleWithDefaults() *V1HostRule`

NewV1HostRuleWithDefaults instantiates a new V1HostRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpBlock

`func (o *V1HostRule) GetIpBlock() []V1IpBlock`

GetIpBlock returns the IpBlock field if non-nil, zero value otherwise.

### GetIpBlockOk

`func (o *V1HostRule) GetIpBlockOk() (*[]V1IpBlock, bool)`

GetIpBlockOk returns a tuple with the IpBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpBlock

`func (o *V1HostRule) SetIpBlock(v []V1IpBlock)`

SetIpBlock sets IpBlock field to given value.

### HasIpBlock

`func (o *V1HostRule) HasIpBlock() bool`

HasIpBlock returns a boolean if a field has been set.

### GetInstanceSelectors

`func (o *V1HostRule) GetInstanceSelectors() []V1MatchExpression`

GetInstanceSelectors returns the InstanceSelectors field if non-nil, zero value otherwise.

### GetInstanceSelectorsOk

`func (o *V1HostRule) GetInstanceSelectorsOk() (*[]V1MatchExpression, bool)`

GetInstanceSelectorsOk returns a tuple with the InstanceSelectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSelectors

`func (o *V1HostRule) SetInstanceSelectors(v []V1MatchExpression)`

SetInstanceSelectors sets InstanceSelectors field to given value.

### HasInstanceSelectors

`func (o *V1HostRule) HasInstanceSelectors() bool`

HasInstanceSelectors returns a boolean if a field has been set.

### GetNetworkSelectors

`func (o *V1HostRule) GetNetworkSelectors() []V1MatchExpression`

GetNetworkSelectors returns the NetworkSelectors field if non-nil, zero value otherwise.

### GetNetworkSelectorsOk

`func (o *V1HostRule) GetNetworkSelectorsOk() (*[]V1MatchExpression, bool)`

GetNetworkSelectorsOk returns a tuple with the NetworkSelectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSelectors

`func (o *V1HostRule) SetNetworkSelectors(v []V1MatchExpression)`

SetNetworkSelectors sets NetworkSelectors field to given value.

### HasNetworkSelectors

`func (o *V1HostRule) HasNetworkSelectors() bool`

HasNetworkSelectors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


