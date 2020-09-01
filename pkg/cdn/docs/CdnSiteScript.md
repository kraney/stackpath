# CdnSiteScript

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A serverless script&#39;s unique identifier | [optional] 
**StackId** | Pointer to **string** | The ID of the stack to which a serverless script&#39;s site belongs | [optional] 
**SiteId** | Pointer to **string** | The ID of the site to which a serverless script belongs | [optional] 
**Name** | Pointer to **string** | A serverless script&#39;s name | [optional] 
**Version** | Pointer to **string** | A serverless script&#39;s version number  Version numbers start at 1 and are incremented every time the script is updated. | [optional] 
**Code** | Pointer to **string** | Base64 encoded contents of a serverless script | [optional] 
**Paths** | Pointer to **[]string** | The URL paths that incoming requests should answer with a serverless script  Every serverless script needs at least one path, and paths must be unique. | [optional] 
**CreatedAt** | Pointer to [**time.Time**](time.Time.md) | The date a serverless script was created | [optional] 
**UpdatedAt** | Pointer to [**time.Time**](time.Time.md) | The date a serverless script was last updated | [optional] 

## Methods

### NewCdnSiteScript

`func NewCdnSiteScript() *CdnSiteScript`

NewCdnSiteScript instantiates a new CdnSiteScript object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdnSiteScriptWithDefaults

`func NewCdnSiteScriptWithDefaults() *CdnSiteScript`

NewCdnSiteScriptWithDefaults instantiates a new CdnSiteScript object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CdnSiteScript) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CdnSiteScript) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CdnSiteScript) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CdnSiteScript) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStackId

`func (o *CdnSiteScript) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *CdnSiteScript) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *CdnSiteScript) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *CdnSiteScript) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetSiteId

`func (o *CdnSiteScript) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *CdnSiteScript) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *CdnSiteScript) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *CdnSiteScript) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetName

`func (o *CdnSiteScript) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CdnSiteScript) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CdnSiteScript) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CdnSiteScript) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *CdnSiteScript) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CdnSiteScript) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CdnSiteScript) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CdnSiteScript) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCode

`func (o *CdnSiteScript) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CdnSiteScript) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CdnSiteScript) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CdnSiteScript) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetPaths

`func (o *CdnSiteScript) GetPaths() []string`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *CdnSiteScript) GetPathsOk() (*[]string, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *CdnSiteScript) SetPaths(v []string)`

SetPaths sets Paths field to given value.

### HasPaths

`func (o *CdnSiteScript) HasPaths() bool`

HasPaths returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CdnSiteScript) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CdnSiteScript) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CdnSiteScript) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CdnSiteScript) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CdnSiteScript) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CdnSiteScript) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CdnSiteScript) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CdnSiteScript) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


