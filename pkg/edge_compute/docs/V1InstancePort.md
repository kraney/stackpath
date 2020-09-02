# V1InstancePort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | **int32** | The network port | [optional] 
**Protocol** | **string** | The network protocol for the port  Values are either \&quot;TCP\&quot; or \&quot;UDP\&quot;. Defaults to \&quot;TCP\&quot;. | [optional] 
**EnableImplicitNetworkPolicy** | **bool** | Allow the internet to connect to the port  When true, this port will be given an implicit network policy of priority 65534 permitting 0.0.0.0/0 access to the port. This can be overridden by creating network policies of a higher priority to block the port. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


