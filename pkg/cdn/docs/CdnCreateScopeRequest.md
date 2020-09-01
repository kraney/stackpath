# CdnCreateScopeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** | The HTTP request path the scope should respond to | [optional] 
**Platform** | Pointer to **string** | The StackPath billing platform to create the scope on | [optional] 

## Methods

### NewCdnCreateScopeRequest

`func NewCdnCreateScopeRequest() *CdnCreateScopeRequest`

NewCdnCreateScopeRequest instantiates a new CdnCreateScopeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdnCreateScopeRequestWithDefaults

`func NewCdnCreateScopeRequestWithDefaults() *CdnCreateScopeRequest`

NewCdnCreateScopeRequestWithDefaults instantiates a new CdnCreateScopeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *CdnCreateScopeRequest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CdnCreateScopeRequest) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CdnCreateScopeRequest) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *CdnCreateScopeRequest) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPlatform

`func (o *CdnCreateScopeRequest) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *CdnCreateScopeRequest) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *CdnCreateScopeRequest) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *CdnCreateScopeRequest) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


