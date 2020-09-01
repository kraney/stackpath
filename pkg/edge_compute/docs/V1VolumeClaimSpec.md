# V1VolumeClaimSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resources** | Pointer to [**V1ResourceRequirements**](v1ResourceRequirements.md) |  | [optional] 

## Methods

### NewV1VolumeClaimSpec

`func NewV1VolumeClaimSpec() *V1VolumeClaimSpec`

NewV1VolumeClaimSpec instantiates a new V1VolumeClaimSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1VolumeClaimSpecWithDefaults

`func NewV1VolumeClaimSpecWithDefaults() *V1VolumeClaimSpec`

NewV1VolumeClaimSpecWithDefaults instantiates a new V1VolumeClaimSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResources

`func (o *V1VolumeClaimSpec) GetResources() V1ResourceRequirements`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *V1VolumeClaimSpec) GetResourcesOk() (*V1ResourceRequirements, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *V1VolumeClaimSpec) SetResources(v V1ResourceRequirements)`

SetResources sets Resources field to given value.

### HasResources

`func (o *V1VolumeClaimSpec) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


