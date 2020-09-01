# CdnConnectScopeToOriginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Origin** | Pointer to [**CdnConnectScopeToOriginRequestOrigin**](cdnConnectScopeToOriginRequestOrigin.md) |  | [optional] 
**Priority** | Pointer to **int32** | The origin&#39;s priority to the scope  If a CDN scope is powered by more than one origin, then the one with the lower priority number takes higher precedence. | [optional] 
**OriginId** | Pointer to **string** | The ID of an existing origin to associate with a scope  This is useful for connecting to a shared origin. | [optional] 

## Methods

### NewCdnConnectScopeToOriginRequest

`func NewCdnConnectScopeToOriginRequest() *CdnConnectScopeToOriginRequest`

NewCdnConnectScopeToOriginRequest instantiates a new CdnConnectScopeToOriginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdnConnectScopeToOriginRequestWithDefaults

`func NewCdnConnectScopeToOriginRequestWithDefaults() *CdnConnectScopeToOriginRequest`

NewCdnConnectScopeToOriginRequestWithDefaults instantiates a new CdnConnectScopeToOriginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrigin

`func (o *CdnConnectScopeToOriginRequest) GetOrigin() CdnConnectScopeToOriginRequestOrigin`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *CdnConnectScopeToOriginRequest) GetOriginOk() (*CdnConnectScopeToOriginRequestOrigin, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *CdnConnectScopeToOriginRequest) SetOrigin(v CdnConnectScopeToOriginRequestOrigin)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *CdnConnectScopeToOriginRequest) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetPriority

`func (o *CdnConnectScopeToOriginRequest) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CdnConnectScopeToOriginRequest) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CdnConnectScopeToOriginRequest) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *CdnConnectScopeToOriginRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetOriginId

`func (o *CdnConnectScopeToOriginRequest) GetOriginId() string`

GetOriginId returns the OriginId field if non-nil, zero value otherwise.

### GetOriginIdOk

`func (o *CdnConnectScopeToOriginRequest) GetOriginIdOk() (*string, bool)`

GetOriginIdOk returns a tuple with the OriginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginId

`func (o *CdnConnectScopeToOriginRequest) SetOriginId(v string)`

SetOriginId sets OriginId field to given value.

### HasOriginId

`func (o *CdnConnectScopeToOriginRequest) HasOriginId() bool`

HasOriginId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


