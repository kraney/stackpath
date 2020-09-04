# V1Ingress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Detailed summary of what the ingress rule does | [optional] 
**Action** | Pointer to [**V1Action**](v1Action.md) |  | [optional] [default to "UNDEFINED"]
**From** | Pointer to [**V1HostRule**](v1HostRule.md) |  | [optional] 
**Protocols** | Pointer to [**V1Protocols**](v1Protocols.md) |  | [optional] 

## Methods

### NewV1Ingress

`func NewV1Ingress() *V1Ingress`

NewV1Ingress instantiates a new V1Ingress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1IngressWithDefaults

`func NewV1IngressWithDefaults() *V1Ingress`

NewV1IngressWithDefaults instantiates a new V1Ingress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *V1Ingress) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V1Ingress) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V1Ingress) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V1Ingress) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAction

`func (o *V1Ingress) GetAction() V1Action`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *V1Ingress) GetActionOk() (*V1Action, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *V1Ingress) SetAction(v V1Action)`

SetAction sets Action field to given value.

### HasAction

`func (o *V1Ingress) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetFrom

`func (o *V1Ingress) GetFrom() V1HostRule`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *V1Ingress) GetFromOk() (*V1HostRule, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *V1Ingress) SetFrom(v V1HostRule)`

SetFrom sets From field to given value.

### HasFrom

`func (o *V1Ingress) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetProtocols

`func (o *V1Ingress) GetProtocols() V1Protocols`

GetProtocols returns the Protocols field if non-nil, zero value otherwise.

### GetProtocolsOk

`func (o *V1Ingress) GetProtocolsOk() (*V1Protocols, bool)`

GetProtocolsOk returns a tuple with the Protocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocols

`func (o *V1Ingress) SetProtocols(v V1Protocols)`

SetProtocols sets Protocols field to given value.

### HasProtocols

`func (o *V1Ingress) HasProtocols() bool`

HasProtocols returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


