# V1HTTPGetAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** | The path portion of the URL to request | [optional] 
**Port** | Pointer to **int32** | The TCP port to query in the HTTP request | [optional] 
**Scheme** | Pointer to **string** | HTTP scheme to use in the HTTP request | [optional] 
**HttpHeaders** | Pointer to **map[string]string** | A string to string key/value pair | [optional] 

## Methods

### NewV1HTTPGetAction

`func NewV1HTTPGetAction() *V1HTTPGetAction`

NewV1HTTPGetAction instantiates a new V1HTTPGetAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1HTTPGetActionWithDefaults

`func NewV1HTTPGetActionWithDefaults() *V1HTTPGetAction`

NewV1HTTPGetActionWithDefaults instantiates a new V1HTTPGetAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *V1HTTPGetAction) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *V1HTTPGetAction) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *V1HTTPGetAction) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *V1HTTPGetAction) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPort

`func (o *V1HTTPGetAction) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *V1HTTPGetAction) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *V1HTTPGetAction) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *V1HTTPGetAction) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetScheme

`func (o *V1HTTPGetAction) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *V1HTTPGetAction) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *V1HTTPGetAction) SetScheme(v string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *V1HTTPGetAction) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### GetHttpHeaders

`func (o *V1HTTPGetAction) GetHttpHeaders() map[string]string`

GetHttpHeaders returns the HttpHeaders field if non-nil, zero value otherwise.

### GetHttpHeadersOk

`func (o *V1HTTPGetAction) GetHttpHeadersOk() (*map[string]string, bool)`

GetHttpHeadersOk returns a tuple with the HttpHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpHeaders

`func (o *V1HTTPGetAction) SetHttpHeaders(v map[string]string)`

SetHttpHeaders sets HttpHeaders field to given value.

### HasHttpHeaders

`func (o *V1HTTPGetAction) HasHttpHeaders() bool`

HasHttpHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


