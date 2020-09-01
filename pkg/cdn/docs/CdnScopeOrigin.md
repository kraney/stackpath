# CdnScopeOrigin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Priority** | Pointer to **int32** | An origin&#39;s priority to it&#39;s CDN site scope  If a CDN scope is powered by more than one origin, then the one with the lower priority number takes higher precedence. | [optional] 
**Origin** | Pointer to [**SchemacdnOrigin**](schemacdnOrigin.md) |  | [optional] 

## Methods

### NewCdnScopeOrigin

`func NewCdnScopeOrigin() *CdnScopeOrigin`

NewCdnScopeOrigin instantiates a new CdnScopeOrigin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdnScopeOriginWithDefaults

`func NewCdnScopeOriginWithDefaults() *CdnScopeOrigin`

NewCdnScopeOriginWithDefaults instantiates a new CdnScopeOrigin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriority

`func (o *CdnScopeOrigin) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CdnScopeOrigin) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CdnScopeOrigin) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *CdnScopeOrigin) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetOrigin

`func (o *CdnScopeOrigin) GetOrigin() SchemacdnOrigin`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *CdnScopeOrigin) GetOriginOk() (*SchemacdnOrigin, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *CdnScopeOrigin) SetOrigin(v SchemacdnOrigin)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *CdnScopeOrigin) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


