# CdnCreateSiteRequestOrigin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** | The origin&#39;s path  Paths default to \&quot;/\&quot; | [optional] 
**Hostname** | Pointer to **string** | The origin&#39;s hostname or IP address | [optional] 
**Port** | Pointer to **int32** | The HTTP port to connect to the origin | [optional] 
**SecurePort** | Pointer to **int32** | The HTTPS port to connect to the origin | [optional] 

## Methods

### NewCdnCreateSiteRequestOrigin

`func NewCdnCreateSiteRequestOrigin() *CdnCreateSiteRequestOrigin`

NewCdnCreateSiteRequestOrigin instantiates a new CdnCreateSiteRequestOrigin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdnCreateSiteRequestOriginWithDefaults

`func NewCdnCreateSiteRequestOriginWithDefaults() *CdnCreateSiteRequestOrigin`

NewCdnCreateSiteRequestOriginWithDefaults instantiates a new CdnCreateSiteRequestOrigin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *CdnCreateSiteRequestOrigin) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CdnCreateSiteRequestOrigin) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CdnCreateSiteRequestOrigin) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *CdnCreateSiteRequestOrigin) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetHostname

`func (o *CdnCreateSiteRequestOrigin) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *CdnCreateSiteRequestOrigin) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *CdnCreateSiteRequestOrigin) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *CdnCreateSiteRequestOrigin) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetPort

`func (o *CdnCreateSiteRequestOrigin) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CdnCreateSiteRequestOrigin) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CdnCreateSiteRequestOrigin) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *CdnCreateSiteRequestOrigin) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSecurePort

`func (o *CdnCreateSiteRequestOrigin) GetSecurePort() int32`

GetSecurePort returns the SecurePort field if non-nil, zero value otherwise.

### GetSecurePortOk

`func (o *CdnCreateSiteRequestOrigin) GetSecurePortOk() (*int32, bool)`

GetSecurePortOk returns a tuple with the SecurePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurePort

`func (o *CdnCreateSiteRequestOrigin) SetSecurePort(v int32)`

SetSecurePort sets SecurePort field to given value.

### HasSecurePort

`func (o *CdnCreateSiteRequestOrigin) HasSecurePort() bool`

HasSecurePort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


