# V1ImagePullCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DockerRegistry** | Pointer to [**V1DockerRegistryCredentials**](v1DockerRegistryCredentials.md) |  | [optional] 

## Methods

### NewV1ImagePullCredential

`func NewV1ImagePullCredential() *V1ImagePullCredential`

NewV1ImagePullCredential instantiates a new V1ImagePullCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ImagePullCredentialWithDefaults

`func NewV1ImagePullCredentialWithDefaults() *V1ImagePullCredential`

NewV1ImagePullCredentialWithDefaults instantiates a new V1ImagePullCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDockerRegistry

`func (o *V1ImagePullCredential) GetDockerRegistry() V1DockerRegistryCredentials`

GetDockerRegistry returns the DockerRegistry field if non-nil, zero value otherwise.

### GetDockerRegistryOk

`func (o *V1ImagePullCredential) GetDockerRegistryOk() (*V1DockerRegistryCredentials, bool)`

GetDockerRegistryOk returns a tuple with the DockerRegistry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerRegistry

`func (o *V1ImagePullCredential) SetDockerRegistry(v V1DockerRegistryCredentials)`

SetDockerRegistry sets DockerRegistry field to given value.

### HasDockerRegistry

`func (o *V1ImagePullCredential) HasDockerRegistry() bool`

HasDockerRegistry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


