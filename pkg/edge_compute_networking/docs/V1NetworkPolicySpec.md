# V1NetworkPolicySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceSelectors** | [**[]V1MatchExpression**](v1MatchExpression.md) | A selector to match workload instances | [optional] 
**NetworkSelectors** | [**[]V1MatchExpression**](v1MatchExpression.md) | A selector to match networks | [optional] 
**PolicyTypes** | [**[]NetworkPolicySpecPolicyType**](NetworkPolicySpecPolicyType.md) | A list of policy types  Policy types are used to specify what rules will be defined. If a policy type is given but not defined it will default. If it is not provided then no action will be used. | [optional] 
**Priority** | **int32** | A policy&#39;s priority among other network policies. 1 - 65535  Network policies apply to all compute workloads on the stack. Lower values have a higher priority, and priorities must be unique across the stack. Use the special value 65534 to apply the same priority to different workload instances in the same stack. | [optional] 
**Ingress** | [**[]V1Ingress**](v1Ingress.md) | A list of rules for inbound traffic to instances  If the ingress policy type is given but not defined here the default is to block all | [optional] 
**Egress** | [**[]V1Egress**](v1Egress.md) | A list of rules for outbound traffic from instances  If the egress policy type is given but not defined here the default is to allow all | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


