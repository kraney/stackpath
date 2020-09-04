# WafCreateSiteRequestOrigin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** | The path the WAF should request from the origin | [optional] 
**Hostname** | Pointer to **string** | The origin&#39;s fully-qualified domain name | [optional] 
**Port** | Pointer to **int32** | The TCP port the WAF should connect to for http requests | [optional] 
**SecurePort** | Pointer to **int32** | The TCP port the WAF should connect to for https requests | [optional] 
**HostHeader** | Pointer to **string** | The value of the Host header that the WAF should set when requesting from the origin. This field is deprecated and should not be used. | [optional] 

## Methods

### NewWafCreateSiteRequestOrigin

`func NewWafCreateSiteRequestOrigin() *WafCreateSiteRequestOrigin`

NewWafCreateSiteRequestOrigin instantiates a new WafCreateSiteRequestOrigin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafCreateSiteRequestOriginWithDefaults

`func NewWafCreateSiteRequestOriginWithDefaults() *WafCreateSiteRequestOrigin`

NewWafCreateSiteRequestOriginWithDefaults instantiates a new WafCreateSiteRequestOrigin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *WafCreateSiteRequestOrigin) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *WafCreateSiteRequestOrigin) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *WafCreateSiteRequestOrigin) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *WafCreateSiteRequestOrigin) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetHostname

`func (o *WafCreateSiteRequestOrigin) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *WafCreateSiteRequestOrigin) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *WafCreateSiteRequestOrigin) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *WafCreateSiteRequestOrigin) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetPort

`func (o *WafCreateSiteRequestOrigin) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *WafCreateSiteRequestOrigin) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *WafCreateSiteRequestOrigin) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *WafCreateSiteRequestOrigin) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSecurePort

`func (o *WafCreateSiteRequestOrigin) GetSecurePort() int32`

GetSecurePort returns the SecurePort field if non-nil, zero value otherwise.

### GetSecurePortOk

`func (o *WafCreateSiteRequestOrigin) GetSecurePortOk() (*int32, bool)`

GetSecurePortOk returns a tuple with the SecurePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurePort

`func (o *WafCreateSiteRequestOrigin) SetSecurePort(v int32)`

SetSecurePort sets SecurePort field to given value.

### HasSecurePort

`func (o *WafCreateSiteRequestOrigin) HasSecurePort() bool`

HasSecurePort returns a boolean if a field has been set.

### GetHostHeader

`func (o *WafCreateSiteRequestOrigin) GetHostHeader() string`

GetHostHeader returns the HostHeader field if non-nil, zero value otherwise.

### GetHostHeaderOk

`func (o *WafCreateSiteRequestOrigin) GetHostHeaderOk() (*string, bool)`

GetHostHeaderOk returns a tuple with the HostHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostHeader

`func (o *WafCreateSiteRequestOrigin) SetHostHeader(v string)`

SetHostHeader sets HostHeader field to given value.

### HasHostHeader

`func (o *WafCreateSiteRequestOrigin) HasHostHeader() bool`

HasHostHeader returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


