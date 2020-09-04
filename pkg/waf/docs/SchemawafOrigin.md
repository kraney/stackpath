# SchemawafOrigin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** | An origin&#39;s path | [optional] 
**Hostname** | Pointer to **string** | An origin&#39;s fully-qualified domain name | [optional] 
**Port** | Pointer to **int32** | The http port to connect to the origin | [optional] 
**SecurePort** | Pointer to **int32** | The https port to connect to the origin | [optional] 

## Methods

### NewSchemawafOrigin

`func NewSchemawafOrigin() *SchemawafOrigin`

NewSchemawafOrigin instantiates a new SchemawafOrigin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemawafOriginWithDefaults

`func NewSchemawafOriginWithDefaults() *SchemawafOrigin`

NewSchemawafOriginWithDefaults instantiates a new SchemawafOrigin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *SchemawafOrigin) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SchemawafOrigin) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SchemawafOrigin) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *SchemawafOrigin) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetHostname

`func (o *SchemawafOrigin) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *SchemawafOrigin) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *SchemawafOrigin) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *SchemawafOrigin) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetPort

`func (o *SchemawafOrigin) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SchemawafOrigin) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SchemawafOrigin) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *SchemawafOrigin) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSecurePort

`func (o *SchemawafOrigin) GetSecurePort() int32`

GetSecurePort returns the SecurePort field if non-nil, zero value otherwise.

### GetSecurePortOk

`func (o *SchemawafOrigin) GetSecurePortOk() (*int32, bool)`

GetSecurePortOk returns a tuple with the SecurePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurePort

`func (o *SchemawafOrigin) SetSecurePort(v int32)`

SetSecurePort sets SecurePort field to given value.

### HasSecurePort

`func (o *SchemawafOrigin) HasSecurePort() bool`

HasSecurePort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


