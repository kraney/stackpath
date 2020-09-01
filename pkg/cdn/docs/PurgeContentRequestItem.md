# PurgeContentRequestItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | The URL at which to delete content | [optional] 
**Recursive** | Pointer to **bool** | Whether or not to recursively delete content from the CDN | [optional] 
**InvalidateOnly** | Pointer to **bool** | Whether or not to mark the asset as expired and re-validate instead of deleting | [optional] 
**PurgeAllDynamic** | Pointer to **bool** | Whether or not to purge dynamic versions of assets  This is ignored if recursive is true. | [optional] 
**Headers** | Pointer to **[]string** | A list of HTTP request headers used to construct a cache key to purge content by. These headers must be configured in the site configuration&#39;s _DynamicContent.headerFields_ property. | [optional] 
**PurgeSelector** | Pointer to [**[]PurgeContentRequestPurgeSelector**](PurgeContentRequestPurgeSelector.md) | A key/value pair definition of content to purge from the CDN | [optional] 

## Methods

### NewPurgeContentRequestItem

`func NewPurgeContentRequestItem() *PurgeContentRequestItem`

NewPurgeContentRequestItem instantiates a new PurgeContentRequestItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurgeContentRequestItemWithDefaults

`func NewPurgeContentRequestItemWithDefaults() *PurgeContentRequestItem`

NewPurgeContentRequestItemWithDefaults instantiates a new PurgeContentRequestItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *PurgeContentRequestItem) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PurgeContentRequestItem) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PurgeContentRequestItem) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PurgeContentRequestItem) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetRecursive

`func (o *PurgeContentRequestItem) GetRecursive() bool`

GetRecursive returns the Recursive field if non-nil, zero value otherwise.

### GetRecursiveOk

`func (o *PurgeContentRequestItem) GetRecursiveOk() (*bool, bool)`

GetRecursiveOk returns a tuple with the Recursive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursive

`func (o *PurgeContentRequestItem) SetRecursive(v bool)`

SetRecursive sets Recursive field to given value.

### HasRecursive

`func (o *PurgeContentRequestItem) HasRecursive() bool`

HasRecursive returns a boolean if a field has been set.

### GetInvalidateOnly

`func (o *PurgeContentRequestItem) GetInvalidateOnly() bool`

GetInvalidateOnly returns the InvalidateOnly field if non-nil, zero value otherwise.

### GetInvalidateOnlyOk

`func (o *PurgeContentRequestItem) GetInvalidateOnlyOk() (*bool, bool)`

GetInvalidateOnlyOk returns a tuple with the InvalidateOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidateOnly

`func (o *PurgeContentRequestItem) SetInvalidateOnly(v bool)`

SetInvalidateOnly sets InvalidateOnly field to given value.

### HasInvalidateOnly

`func (o *PurgeContentRequestItem) HasInvalidateOnly() bool`

HasInvalidateOnly returns a boolean if a field has been set.

### GetPurgeAllDynamic

`func (o *PurgeContentRequestItem) GetPurgeAllDynamic() bool`

GetPurgeAllDynamic returns the PurgeAllDynamic field if non-nil, zero value otherwise.

### GetPurgeAllDynamicOk

`func (o *PurgeContentRequestItem) GetPurgeAllDynamicOk() (*bool, bool)`

GetPurgeAllDynamicOk returns a tuple with the PurgeAllDynamic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurgeAllDynamic

`func (o *PurgeContentRequestItem) SetPurgeAllDynamic(v bool)`

SetPurgeAllDynamic sets PurgeAllDynamic field to given value.

### HasPurgeAllDynamic

`func (o *PurgeContentRequestItem) HasPurgeAllDynamic() bool`

HasPurgeAllDynamic returns a boolean if a field has been set.

### GetHeaders

`func (o *PurgeContentRequestItem) GetHeaders() []string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *PurgeContentRequestItem) GetHeadersOk() (*[]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *PurgeContentRequestItem) SetHeaders(v []string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *PurgeContentRequestItem) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetPurgeSelector

`func (o *PurgeContentRequestItem) GetPurgeSelector() []PurgeContentRequestPurgeSelector`

GetPurgeSelector returns the PurgeSelector field if non-nil, zero value otherwise.

### GetPurgeSelectorOk

`func (o *PurgeContentRequestItem) GetPurgeSelectorOk() (*[]PurgeContentRequestPurgeSelector, bool)`

GetPurgeSelectorOk returns a tuple with the PurgeSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurgeSelector

`func (o *PurgeContentRequestItem) SetPurgeSelector(v []PurgeContentRequestPurgeSelector)`

SetPurgeSelector sets PurgeSelector field to given value.

### HasPurgeSelector

`func (o *PurgeContentRequestItem) HasPurgeSelector() bool`

HasPurgeSelector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


