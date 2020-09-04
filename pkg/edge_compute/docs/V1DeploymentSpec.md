# V1DeploymentSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinReplicas** | Pointer to **int32** | The minimum number of instances in a deployment | [optional] 
**MaxReplicas** | Pointer to **int32** | The maximum number of instances in a deployment | [optional] 
**ScaleSettings** | Pointer to [**V1ScaleSettings**](v1ScaleSettings.md) |  | [optional] 
**Selectors** | Pointer to [**[]V1MatchExpression**](v1MatchExpression.md) | A collection of filters that match the deployment&#39;s scope | [optional] 

## Methods

### NewV1DeploymentSpec

`func NewV1DeploymentSpec() *V1DeploymentSpec`

NewV1DeploymentSpec instantiates a new V1DeploymentSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DeploymentSpecWithDefaults

`func NewV1DeploymentSpecWithDefaults() *V1DeploymentSpec`

NewV1DeploymentSpecWithDefaults instantiates a new V1DeploymentSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinReplicas

`func (o *V1DeploymentSpec) GetMinReplicas() int32`

GetMinReplicas returns the MinReplicas field if non-nil, zero value otherwise.

### GetMinReplicasOk

`func (o *V1DeploymentSpec) GetMinReplicasOk() (*int32, bool)`

GetMinReplicasOk returns a tuple with the MinReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReplicas

`func (o *V1DeploymentSpec) SetMinReplicas(v int32)`

SetMinReplicas sets MinReplicas field to given value.

### HasMinReplicas

`func (o *V1DeploymentSpec) HasMinReplicas() bool`

HasMinReplicas returns a boolean if a field has been set.

### GetMaxReplicas

`func (o *V1DeploymentSpec) GetMaxReplicas() int32`

GetMaxReplicas returns the MaxReplicas field if non-nil, zero value otherwise.

### GetMaxReplicasOk

`func (o *V1DeploymentSpec) GetMaxReplicasOk() (*int32, bool)`

GetMaxReplicasOk returns a tuple with the MaxReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReplicas

`func (o *V1DeploymentSpec) SetMaxReplicas(v int32)`

SetMaxReplicas sets MaxReplicas field to given value.

### HasMaxReplicas

`func (o *V1DeploymentSpec) HasMaxReplicas() bool`

HasMaxReplicas returns a boolean if a field has been set.

### GetScaleSettings

`func (o *V1DeploymentSpec) GetScaleSettings() V1ScaleSettings`

GetScaleSettings returns the ScaleSettings field if non-nil, zero value otherwise.

### GetScaleSettingsOk

`func (o *V1DeploymentSpec) GetScaleSettingsOk() (*V1ScaleSettings, bool)`

GetScaleSettingsOk returns a tuple with the ScaleSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleSettings

`func (o *V1DeploymentSpec) SetScaleSettings(v V1ScaleSettings)`

SetScaleSettings sets ScaleSettings field to given value.

### HasScaleSettings

`func (o *V1DeploymentSpec) HasScaleSettings() bool`

HasScaleSettings returns a boolean if a field has been set.

### GetSelectors

`func (o *V1DeploymentSpec) GetSelectors() []V1MatchExpression`

GetSelectors returns the Selectors field if non-nil, zero value otherwise.

### GetSelectorsOk

`func (o *V1DeploymentSpec) GetSelectorsOk() (*[]V1MatchExpression, bool)`

GetSelectorsOk returns a tuple with the Selectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectors

`func (o *V1DeploymentSpec) SetSelectors(v []V1MatchExpression)`

SetSelectors sets Selectors field to given value.

### HasSelectors

`func (o *V1DeploymentSpec) HasSelectors() bool`

HasSelectors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


