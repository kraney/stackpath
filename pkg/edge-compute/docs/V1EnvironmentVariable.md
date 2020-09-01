# V1EnvironmentVariable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** | An environment variable&#39;s value | [optional] 
**SecretValue** | Pointer to **string** | A sensitive environment variable that should not be exposed | [optional] 

## Methods

### NewV1EnvironmentVariable

`func NewV1EnvironmentVariable() *V1EnvironmentVariable`

NewV1EnvironmentVariable instantiates a new V1EnvironmentVariable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1EnvironmentVariableWithDefaults

`func NewV1EnvironmentVariableWithDefaults() *V1EnvironmentVariable`

NewV1EnvironmentVariableWithDefaults instantiates a new V1EnvironmentVariable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *V1EnvironmentVariable) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *V1EnvironmentVariable) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *V1EnvironmentVariable) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *V1EnvironmentVariable) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetSecretValue

`func (o *V1EnvironmentVariable) GetSecretValue() string`

GetSecretValue returns the SecretValue field if non-nil, zero value otherwise.

### GetSecretValueOk

`func (o *V1EnvironmentVariable) GetSecretValueOk() (*string, bool)`

GetSecretValueOk returns a tuple with the SecretValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretValue

`func (o *V1EnvironmentVariable) SetSecretValue(v string)`

SetSecretValue sets SecretValue field to given value.

### HasSecretValue

`func (o *V1EnvironmentVariable) HasSecretValue() bool`

HasSecretValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


