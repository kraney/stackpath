# V1NetworkPolicySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceSelectors** | Pointer to [**[]V1MatchExpression**](v1MatchExpression.md) | A selector to match workload instances | [optional] 
**NetworkSelectors** | Pointer to [**[]V1MatchExpression**](v1MatchExpression.md) | A selector to match networks | [optional] 
**PolicyTypes** | Pointer to [**[]NetworkPolicySpecPolicyType**](NetworkPolicySpecPolicyType.md) | A list of policy types  Policy types are used to specify what rules will be defined. If a policy type is given but not defined it will default. If it is not provided then no action will be used. | [optional] 
**Priority** | Pointer to **int32** | A policy&#39;s priority among other network policies. 1 - 65535  Network policies apply to all compute workloads on the stack. Lower values have a higher priority, and priorities must be unique across the stack. Use the special value 65534 to apply the same priority to different workload instances in the same stack. | [optional] 
**Ingress** | Pointer to [**[]V1Ingress**](v1Ingress.md) | A list of rules for inbound traffic to instances  If the ingress policy type is given but not defined here the default is to block all | [optional] 
**Egress** | Pointer to [**[]V1Egress**](v1Egress.md) | A list of rules for outbound traffic from instances  If the egress policy type is given but not defined here the default is to allow all | [optional] 

## Methods

### NewV1NetworkPolicySpec

`func NewV1NetworkPolicySpec() *V1NetworkPolicySpec`

NewV1NetworkPolicySpec instantiates a new V1NetworkPolicySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1NetworkPolicySpecWithDefaults

`func NewV1NetworkPolicySpecWithDefaults() *V1NetworkPolicySpec`

NewV1NetworkPolicySpecWithDefaults instantiates a new V1NetworkPolicySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceSelectors

`func (o *V1NetworkPolicySpec) GetInstanceSelectors() []V1MatchExpression`

GetInstanceSelectors returns the InstanceSelectors field if non-nil, zero value otherwise.

### GetInstanceSelectorsOk

`func (o *V1NetworkPolicySpec) GetInstanceSelectorsOk() (*[]V1MatchExpression, bool)`

GetInstanceSelectorsOk returns a tuple with the InstanceSelectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSelectors

`func (o *V1NetworkPolicySpec) SetInstanceSelectors(v []V1MatchExpression)`

SetInstanceSelectors sets InstanceSelectors field to given value.

### HasInstanceSelectors

`func (o *V1NetworkPolicySpec) HasInstanceSelectors() bool`

HasInstanceSelectors returns a boolean if a field has been set.

### GetNetworkSelectors

`func (o *V1NetworkPolicySpec) GetNetworkSelectors() []V1MatchExpression`

GetNetworkSelectors returns the NetworkSelectors field if non-nil, zero value otherwise.

### GetNetworkSelectorsOk

`func (o *V1NetworkPolicySpec) GetNetworkSelectorsOk() (*[]V1MatchExpression, bool)`

GetNetworkSelectorsOk returns a tuple with the NetworkSelectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSelectors

`func (o *V1NetworkPolicySpec) SetNetworkSelectors(v []V1MatchExpression)`

SetNetworkSelectors sets NetworkSelectors field to given value.

### HasNetworkSelectors

`func (o *V1NetworkPolicySpec) HasNetworkSelectors() bool`

HasNetworkSelectors returns a boolean if a field has been set.

### GetPolicyTypes

`func (o *V1NetworkPolicySpec) GetPolicyTypes() []NetworkPolicySpecPolicyType`

GetPolicyTypes returns the PolicyTypes field if non-nil, zero value otherwise.

### GetPolicyTypesOk

`func (o *V1NetworkPolicySpec) GetPolicyTypesOk() (*[]NetworkPolicySpecPolicyType, bool)`

GetPolicyTypesOk returns a tuple with the PolicyTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyTypes

`func (o *V1NetworkPolicySpec) SetPolicyTypes(v []NetworkPolicySpecPolicyType)`

SetPolicyTypes sets PolicyTypes field to given value.

### HasPolicyTypes

`func (o *V1NetworkPolicySpec) HasPolicyTypes() bool`

HasPolicyTypes returns a boolean if a field has been set.

### GetPriority

`func (o *V1NetworkPolicySpec) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *V1NetworkPolicySpec) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *V1NetworkPolicySpec) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *V1NetworkPolicySpec) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetIngress

`func (o *V1NetworkPolicySpec) GetIngress() []V1Ingress`

GetIngress returns the Ingress field if non-nil, zero value otherwise.

### GetIngressOk

`func (o *V1NetworkPolicySpec) GetIngressOk() (*[]V1Ingress, bool)`

GetIngressOk returns a tuple with the Ingress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngress

`func (o *V1NetworkPolicySpec) SetIngress(v []V1Ingress)`

SetIngress sets Ingress field to given value.

### HasIngress

`func (o *V1NetworkPolicySpec) HasIngress() bool`

HasIngress returns a boolean if a field has been set.

### GetEgress

`func (o *V1NetworkPolicySpec) GetEgress() []V1Egress`

GetEgress returns the Egress field if non-nil, zero value otherwise.

### GetEgressOk

`func (o *V1NetworkPolicySpec) GetEgressOk() (*[]V1Egress, bool)`

GetEgressOk returns a tuple with the Egress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgress

`func (o *V1NetworkPolicySpec) SetEgress(v []V1Egress)`

SetEgress sets Egress field to given value.

### HasEgress

`func (o *V1NetworkPolicySpec) HasEgress() bool`

HasEgress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


