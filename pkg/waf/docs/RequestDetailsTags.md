# RequestDetailsTags

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **[]string** | The tags associated with the request | [optional] 
**Hash** | Pointer to **string** | The tag&#39;s hash | [optional] 
**Cached** | Pointer to **bool** | Whether or not the tag was cached | [optional] 

## Methods

### NewRequestDetailsTags

`func NewRequestDetailsTags() *RequestDetailsTags`

NewRequestDetailsTags instantiates a new RequestDetailsTags object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestDetailsTagsWithDefaults

`func NewRequestDetailsTagsWithDefaults() *RequestDetailsTags`

NewRequestDetailsTagsWithDefaults instantiates a new RequestDetailsTags object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *RequestDetailsTags) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RequestDetailsTags) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RequestDetailsTags) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RequestDetailsTags) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetHash

`func (o *RequestDetailsTags) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *RequestDetailsTags) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *RequestDetailsTags) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *RequestDetailsTags) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetCached

`func (o *RequestDetailsTags) GetCached() bool`

GetCached returns the Cached field if non-nil, zero value otherwise.

### GetCachedOk

`func (o *RequestDetailsTags) GetCachedOk() (*bool, bool)`

GetCachedOk returns a tuple with the Cached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCached

`func (o *RequestDetailsTags) SetCached(v bool)`

SetCached sets Cached field to given value.

### HasCached

`func (o *RequestDetailsTags) HasCached() bool`

HasCached returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


