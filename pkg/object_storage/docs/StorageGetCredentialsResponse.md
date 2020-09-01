# StorageGetCredentialsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**[]GetCredentialsResponseCredential**](GetCredentialsResponseCredential.md) | The list of active credentials on account | [optional] 

## Methods

### NewStorageGetCredentialsResponse

`func NewStorageGetCredentialsResponse() *StorageGetCredentialsResponse`

NewStorageGetCredentialsResponse instantiates a new StorageGetCredentialsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageGetCredentialsResponseWithDefaults

`func NewStorageGetCredentialsResponseWithDefaults() *StorageGetCredentialsResponse`

NewStorageGetCredentialsResponseWithDefaults instantiates a new StorageGetCredentialsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *StorageGetCredentialsResponse) GetCredentials() []GetCredentialsResponseCredential`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *StorageGetCredentialsResponse) GetCredentialsOk() (*[]GetCredentialsResponseCredential, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *StorageGetCredentialsResponse) SetCredentials(v []GetCredentialsResponseCredential)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *StorageGetCredentialsResponse) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


