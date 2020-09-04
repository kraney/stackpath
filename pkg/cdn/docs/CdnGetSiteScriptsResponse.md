# CdnGetSiteScriptsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageInfo** | Pointer to [**PaginationPageInfo**](paginationPageInfo.md) |  | [optional] 
**Results** | Pointer to [**[]CdnSiteScript**](cdnSiteScript.md) | The requested serverless scripts | [optional] 

## Methods

### NewCdnGetSiteScriptsResponse

`func NewCdnGetSiteScriptsResponse() *CdnGetSiteScriptsResponse`

NewCdnGetSiteScriptsResponse instantiates a new CdnGetSiteScriptsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdnGetSiteScriptsResponseWithDefaults

`func NewCdnGetSiteScriptsResponseWithDefaults() *CdnGetSiteScriptsResponse`

NewCdnGetSiteScriptsResponseWithDefaults instantiates a new CdnGetSiteScriptsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageInfo

`func (o *CdnGetSiteScriptsResponse) GetPageInfo() PaginationPageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *CdnGetSiteScriptsResponse) GetPageInfoOk() (*PaginationPageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *CdnGetSiteScriptsResponse) SetPageInfo(v PaginationPageInfo)`

SetPageInfo sets PageInfo field to given value.

### HasPageInfo

`func (o *CdnGetSiteScriptsResponse) HasPageInfo() bool`

HasPageInfo returns a boolean if a field has been set.

### GetResults

`func (o *CdnGetSiteScriptsResponse) GetResults() []CdnSiteScript`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *CdnGetSiteScriptsResponse) GetResultsOk() (*[]CdnSiteScript, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *CdnGetSiteScriptsResponse) SetResults(v []CdnSiteScript)`

SetResults sets Results field to given value.

### HasResults

`func (o *CdnGetSiteScriptsResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


