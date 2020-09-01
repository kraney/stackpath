# V1Egress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Detailed summary of what the egress rule does | [optional] 
**Action** | Pointer to [**V1Action**](v1Action.md) |  | [optional] [default to "UNDEFINED"]
**To** | Pointer to [**V1HostRule**](v1HostRule.md) |  | [optional] 
**Protocols** | Pointer to [**V1Protocols**](v1Protocols.md) |  | [optional] 

## Methods

### NewV1Egress

`func NewV1Egress() *V1Egress`

NewV1Egress instantiates a new V1Egress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1EgressWithDefaults

`func NewV1EgressWithDefaults() *V1Egress`

NewV1EgressWithDefaults instantiates a new V1Egress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *V1Egress) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V1Egress) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V1Egress) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V1Egress) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAction

`func (o *V1Egress) GetAction() V1Action`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *V1Egress) GetActionOk() (*V1Action, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *V1Egress) SetAction(v V1Action)`

SetAction sets Action field to given value.

### HasAction

`func (o *V1Egress) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetTo

`func (o *V1Egress) GetTo() V1HostRule`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *V1Egress) GetToOk() (*V1HostRule, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *V1Egress) SetTo(v V1HostRule)`

SetTo sets To field to given value.

### HasTo

`func (o *V1Egress) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetProtocols

`func (o *V1Egress) GetProtocols() V1Protocols`

GetProtocols returns the Protocols field if non-nil, zero value otherwise.

### GetProtocolsOk

`func (o *V1Egress) GetProtocolsOk() (*V1Protocols, bool)`

GetProtocolsOk returns a tuple with the Protocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocols

`func (o *V1Egress) SetProtocols(v V1Protocols)`

SetProtocols sets Protocols field to given value.

### HasProtocols

`func (o *V1Egress) HasProtocols() bool`

HasProtocols returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


