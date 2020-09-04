# CdnConnectScopeToOriginRequestOrigin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** | The origin&#39;s path  Paths default to \&quot;/\&quot; | [optional] 
**Hostname** | Pointer to **string** | The origin&#39;s hostname or IP address | [optional] 
**Port** | Pointer to **int32** | The HTTP port to connect to the origin | [optional] 
**SecurePort** | Pointer to **int32** | The HTTPS port to connect to the origin | [optional] 

## Methods

### NewCdnConnectScopeToOriginRequestOrigin

`func NewCdnConnectScopeToOriginRequestOrigin() *CdnConnectScopeToOriginRequestOrigin`

NewCdnConnectScopeToOriginRequestOrigin instantiates a new CdnConnectScopeToOriginRequestOrigin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdnConnectScopeToOriginRequestOriginWithDefaults

`func NewCdnConnectScopeToOriginRequestOriginWithDefaults() *CdnConnectScopeToOriginRequestOrigin`

NewCdnConnectScopeToOriginRequestOriginWithDefaults instantiates a new CdnConnectScopeToOriginRequestOrigin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *CdnConnectScopeToOriginRequestOrigin) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CdnConnectScopeToOriginRequestOrigin) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CdnConnectScopeToOriginRequestOrigin) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *CdnConnectScopeToOriginRequestOrigin) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetHostname

`func (o *CdnConnectScopeToOriginRequestOrigin) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *CdnConnectScopeToOriginRequestOrigin) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *CdnConnectScopeToOriginRequestOrigin) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *CdnConnectScopeToOriginRequestOrigin) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetPort

`func (o *CdnConnectScopeToOriginRequestOrigin) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CdnConnectScopeToOriginRequestOrigin) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CdnConnectScopeToOriginRequestOrigin) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *CdnConnectScopeToOriginRequestOrigin) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSecurePort

`func (o *CdnConnectScopeToOriginRequestOrigin) GetSecurePort() int32`

GetSecurePort returns the SecurePort field if non-nil, zero value otherwise.

### GetSecurePortOk

`func (o *CdnConnectScopeToOriginRequestOrigin) GetSecurePortOk() (*int32, bool)`

GetSecurePortOk returns a tuple with the SecurePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurePort

`func (o *CdnConnectScopeToOriginRequestOrigin) SetSecurePort(v int32)`

SetSecurePort sets SecurePort field to given value.

### HasSecurePort

`func (o *CdnConnectScopeToOriginRequestOrigin) HasSecurePort() bool`

HasSecurePort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


