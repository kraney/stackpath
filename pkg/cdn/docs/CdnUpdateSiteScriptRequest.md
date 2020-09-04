# CdnUpdateSiteScriptRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Base64 encoded contents of a serverless script | [optional] 
**Paths** | Pointer to **[]string** | The HTTP request paths that are handled by a serverless script | [optional] 

## Methods

### NewCdnUpdateSiteScriptRequest

`func NewCdnUpdateSiteScriptRequest() *CdnUpdateSiteScriptRequest`

NewCdnUpdateSiteScriptRequest instantiates a new CdnUpdateSiteScriptRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdnUpdateSiteScriptRequestWithDefaults

`func NewCdnUpdateSiteScriptRequestWithDefaults() *CdnUpdateSiteScriptRequest`

NewCdnUpdateSiteScriptRequestWithDefaults instantiates a new CdnUpdateSiteScriptRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *CdnUpdateSiteScriptRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CdnUpdateSiteScriptRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CdnUpdateSiteScriptRequest) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CdnUpdateSiteScriptRequest) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetPaths

`func (o *CdnUpdateSiteScriptRequest) GetPaths() []string`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *CdnUpdateSiteScriptRequest) GetPathsOk() (*[]string, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *CdnUpdateSiteScriptRequest) SetPaths(v []string)`

SetPaths sets Paths field to given value.

### HasPaths

`func (o *CdnUpdateSiteScriptRequest) HasPaths() bool`

HasPaths returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


