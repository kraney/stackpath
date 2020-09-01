# DeliveryConnectScopeToOriginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Origin** | Pointer to [**DeliveryConnectScopeToOriginRequestOrigin**](deliveryConnectScopeToOriginRequestOrigin.md) |  | [optional] 
**Priority** | Pointer to **int32** | The origin&#39;s priority to the scope  If more than one origin powers a CDN scope, then the one with the lower priority number takes higher precedence. When there is an origin already in place, the following rules are followed:  - If an origin ID is provided, then the current origin at that priority is disconnected in favor of the new one. - If an origin is provided and the current origin at the given priority is dedicated, then the origin is updated in place. - If an origin is provided and no dedicated origin exists, the origin is created and connected to the scope. | [optional] 
**OriginId** | Pointer to **string** | The ID of an existing origin to associate with a scope  This is useful for connecting to a shared origin. | [optional] 

## Methods

### NewDeliveryConnectScopeToOriginRequest

`func NewDeliveryConnectScopeToOriginRequest() *DeliveryConnectScopeToOriginRequest`

NewDeliveryConnectScopeToOriginRequest instantiates a new DeliveryConnectScopeToOriginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryConnectScopeToOriginRequestWithDefaults

`func NewDeliveryConnectScopeToOriginRequestWithDefaults() *DeliveryConnectScopeToOriginRequest`

NewDeliveryConnectScopeToOriginRequestWithDefaults instantiates a new DeliveryConnectScopeToOriginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrigin

`func (o *DeliveryConnectScopeToOriginRequest) GetOrigin() DeliveryConnectScopeToOriginRequestOrigin`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *DeliveryConnectScopeToOriginRequest) GetOriginOk() (*DeliveryConnectScopeToOriginRequestOrigin, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *DeliveryConnectScopeToOriginRequest) SetOrigin(v DeliveryConnectScopeToOriginRequestOrigin)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *DeliveryConnectScopeToOriginRequest) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetPriority

`func (o *DeliveryConnectScopeToOriginRequest) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *DeliveryConnectScopeToOriginRequest) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *DeliveryConnectScopeToOriginRequest) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *DeliveryConnectScopeToOriginRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetOriginId

`func (o *DeliveryConnectScopeToOriginRequest) GetOriginId() string`

GetOriginId returns the OriginId field if non-nil, zero value otherwise.

### GetOriginIdOk

`func (o *DeliveryConnectScopeToOriginRequest) GetOriginIdOk() (*string, bool)`

GetOriginIdOk returns a tuple with the OriginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginId

`func (o *DeliveryConnectScopeToOriginRequest) SetOriginId(v string)`

SetOriginId sets OriginId field to given value.

### HasOriginId

`func (o *DeliveryConnectScopeToOriginRequest) HasOriginId() bool`

HasOriginId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


