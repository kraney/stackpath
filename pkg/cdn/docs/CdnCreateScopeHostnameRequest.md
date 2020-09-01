# CdnCreateScopeHostnameRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** | The hostname to add to a scope | [optional] 
**DisableTransparentMode** | Pointer to **bool** | Whether or not to add the hostname to a CDN site&#39;s CDN scope or its WAF scope  When true, this call adds the hostname to a CDN site&#39;s scope instead of loading from a CDN site&#39;s WAF scope, if the site has WAF service. | [optional] 

## Methods

### NewCdnCreateScopeHostnameRequest

`func NewCdnCreateScopeHostnameRequest() *CdnCreateScopeHostnameRequest`

NewCdnCreateScopeHostnameRequest instantiates a new CdnCreateScopeHostnameRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdnCreateScopeHostnameRequestWithDefaults

`func NewCdnCreateScopeHostnameRequestWithDefaults() *CdnCreateScopeHostnameRequest`

NewCdnCreateScopeHostnameRequestWithDefaults instantiates a new CdnCreateScopeHostnameRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *CdnCreateScopeHostnameRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *CdnCreateScopeHostnameRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *CdnCreateScopeHostnameRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *CdnCreateScopeHostnameRequest) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetDisableTransparentMode

`func (o *CdnCreateScopeHostnameRequest) GetDisableTransparentMode() bool`

GetDisableTransparentMode returns the DisableTransparentMode field if non-nil, zero value otherwise.

### GetDisableTransparentModeOk

`func (o *CdnCreateScopeHostnameRequest) GetDisableTransparentModeOk() (*bool, bool)`

GetDisableTransparentModeOk returns a tuple with the DisableTransparentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableTransparentMode

`func (o *CdnCreateScopeHostnameRequest) SetDisableTransparentMode(v bool)`

SetDisableTransparentMode sets DisableTransparentMode field to given value.

### HasDisableTransparentMode

`func (o *CdnCreateScopeHostnameRequest) HasDisableTransparentMode() bool`

HasDisableTransparentMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


