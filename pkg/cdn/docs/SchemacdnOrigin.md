# SchemacdnOrigin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | An origin&#39;s unique identifier | [optional] 
**Path** | Pointer to **string** | An origin&#39;s path  Paths default to \&quot;/\&quot; | [optional] 
**Hostname** | Pointer to **string** | An origin&#39;s hostname or IP address | [optional] 
**Port** | Pointer to **int32** | The HTTP port to connect to the origin | [optional] 
**SecurePort** | Pointer to **int32** | The HTTPS port to connect to the origin | [optional] 
**Dedicated** | Pointer to **bool** | Whether or not an origin is dedicated to a CDN site  Dedicated origins cannot be used by any site other than that to which it is dedicated. | [optional] 
**SiteId** | Pointer to **string** | The ID of the CDN site to which an origin is dedicated | [optional] 

## Methods

### NewSchemacdnOrigin

`func NewSchemacdnOrigin() *SchemacdnOrigin`

NewSchemacdnOrigin instantiates a new SchemacdnOrigin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemacdnOriginWithDefaults

`func NewSchemacdnOriginWithDefaults() *SchemacdnOrigin`

NewSchemacdnOriginWithDefaults instantiates a new SchemacdnOrigin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SchemacdnOrigin) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SchemacdnOrigin) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SchemacdnOrigin) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SchemacdnOrigin) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPath

`func (o *SchemacdnOrigin) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SchemacdnOrigin) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SchemacdnOrigin) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *SchemacdnOrigin) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetHostname

`func (o *SchemacdnOrigin) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *SchemacdnOrigin) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *SchemacdnOrigin) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *SchemacdnOrigin) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetPort

`func (o *SchemacdnOrigin) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SchemacdnOrigin) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SchemacdnOrigin) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *SchemacdnOrigin) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSecurePort

`func (o *SchemacdnOrigin) GetSecurePort() int32`

GetSecurePort returns the SecurePort field if non-nil, zero value otherwise.

### GetSecurePortOk

`func (o *SchemacdnOrigin) GetSecurePortOk() (*int32, bool)`

GetSecurePortOk returns a tuple with the SecurePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurePort

`func (o *SchemacdnOrigin) SetSecurePort(v int32)`

SetSecurePort sets SecurePort field to given value.

### HasSecurePort

`func (o *SchemacdnOrigin) HasSecurePort() bool`

HasSecurePort returns a boolean if a field has been set.

### GetDedicated

`func (o *SchemacdnOrigin) GetDedicated() bool`

GetDedicated returns the Dedicated field if non-nil, zero value otherwise.

### GetDedicatedOk

`func (o *SchemacdnOrigin) GetDedicatedOk() (*bool, bool)`

GetDedicatedOk returns a tuple with the Dedicated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicated

`func (o *SchemacdnOrigin) SetDedicated(v bool)`

SetDedicated sets Dedicated field to given value.

### HasDedicated

`func (o *SchemacdnOrigin) HasDedicated() bool`

HasDedicated returns a boolean if a field has been set.

### GetSiteId

`func (o *SchemacdnOrigin) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *SchemacdnOrigin) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *SchemacdnOrigin) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *SchemacdnOrigin) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


