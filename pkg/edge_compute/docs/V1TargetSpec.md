# V1TargetSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentScope** | Pointer to **string** | The scope at which a deployment should be created. Valid values are: \&quot;cityCode\&quot; | [optional] 
**Deployments** | Pointer to [**V1DeploymentSpec**](v1DeploymentSpec.md) |  | [optional] 

## Methods

### NewV1TargetSpec

`func NewV1TargetSpec() *V1TargetSpec`

NewV1TargetSpec instantiates a new V1TargetSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TargetSpecWithDefaults

`func NewV1TargetSpecWithDefaults() *V1TargetSpec`

NewV1TargetSpecWithDefaults instantiates a new V1TargetSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentScope

`func (o *V1TargetSpec) GetDeploymentScope() string`

GetDeploymentScope returns the DeploymentScope field if non-nil, zero value otherwise.

### GetDeploymentScopeOk

`func (o *V1TargetSpec) GetDeploymentScopeOk() (*string, bool)`

GetDeploymentScopeOk returns a tuple with the DeploymentScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentScope

`func (o *V1TargetSpec) SetDeploymentScope(v string)`

SetDeploymentScope sets DeploymentScope field to given value.

### HasDeploymentScope

`func (o *V1TargetSpec) HasDeploymentScope() bool`

HasDeploymentScope returns a boolean if a field has been set.

### GetDeployments

`func (o *V1TargetSpec) GetDeployments() V1DeploymentSpec`

GetDeployments returns the Deployments field if non-nil, zero value otherwise.

### GetDeploymentsOk

`func (o *V1TargetSpec) GetDeploymentsOk() (*V1DeploymentSpec, bool)`

GetDeploymentsOk returns a tuple with the Deployments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployments

`func (o *V1TargetSpec) SetDeployments(v V1DeploymentSpec)`

SetDeployments sets Deployments field to given value.

### HasDeployments

`func (o *V1TargetSpec) HasDeployments() bool`

HasDeployments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


