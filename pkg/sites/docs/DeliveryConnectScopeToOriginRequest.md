# DeliveryConnectScopeToOriginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Origin** | [**DeliveryConnectScopeToOriginRequestOrigin**](deliveryConnectScopeToOriginRequestOrigin.md) |  | [optional] 
**Priority** | **int32** | The origin&#39;s priority to the scope  If more than one origin powers a CDN scope, then the one with the lower priority number takes higher precedence. When there is an origin already in place, the following rules are followed:  - If an origin ID is provided, then the current origin at that priority is disconnected in favor of the new one. - If an origin is provided and the current origin at the given priority is dedicated, then the origin is updated in place. - If an origin is provided and no dedicated origin exists, the origin is created and connected to the scope. | [optional] 
**OriginId** | **string** | The ID of an existing origin to associate with a scope  This is useful for connecting to a shared origin. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


