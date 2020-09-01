# CdnCreateSiteScriptRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The serverless script&#39;s name | [optional] 
**Code** | Pointer to **string** | The base64 encoded serverless script&#39;s contents | [optional] 
**Paths** | Pointer to **[]string** | The HTTP request paths that are handled by the serverless script | [optional] 

## Methods

### NewCdnCreateSiteScriptRequest

`func NewCdnCreateSiteScriptRequest() *CdnCreateSiteScriptRequest`

NewCdnCreateSiteScriptRequest instantiates a new CdnCreateSiteScriptRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdnCreateSiteScriptRequestWithDefaults

`func NewCdnCreateSiteScriptRequestWithDefaults() *CdnCreateSiteScriptRequest`

NewCdnCreateSiteScriptRequestWithDefaults instantiates a new CdnCreateSiteScriptRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CdnCreateSiteScriptRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CdnCreateSiteScriptRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CdnCreateSiteScriptRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CdnCreateSiteScriptRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *CdnCreateSiteScriptRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CdnCreateSiteScriptRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CdnCreateSiteScriptRequest) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CdnCreateSiteScriptRequest) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetPaths

`func (o *CdnCreateSiteScriptRequest) GetPaths() []string`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *CdnCreateSiteScriptRequest) GetPathsOk() (*[]string, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *CdnCreateSiteScriptRequest) SetPaths(v []string)`

SetPaths sets Paths field to given value.

### HasPaths

`func (o *CdnCreateSiteScriptRequest) HasPaths() bool`

HasPaths returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


