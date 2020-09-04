# CdnScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A CDN site scope&#39;s unique identifier | [optional] 
**Platform** | Pointer to **string** | A CDN site scope&#39;s platform  Scope platforms are used internally by StackPath for metrics collection and billing purposes. Typically, most site scope platforms have the value \&quot;CDS\&quot;. | [optional] 
**Path** | Pointer to **string** | The HTTP request path that is handled by a scope | [optional] 

## Methods

### NewCdnScope

`func NewCdnScope() *CdnScope`

NewCdnScope instantiates a new CdnScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdnScopeWithDefaults

`func NewCdnScopeWithDefaults() *CdnScope`

NewCdnScopeWithDefaults instantiates a new CdnScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CdnScope) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CdnScope) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CdnScope) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CdnScope) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPlatform

`func (o *CdnScope) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *CdnScope) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *CdnScope) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *CdnScope) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetPath

`func (o *CdnScope) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CdnScope) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CdnScope) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *CdnScope) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


