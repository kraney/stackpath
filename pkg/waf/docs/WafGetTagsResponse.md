# WafGetTagsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]WafTag**](wafTag.md) | The requested WAF tags | [optional] 
**PageInfo** | Pointer to [**PaginationPageInfo**](paginationPageInfo.md) |  | [optional] 

## Methods

### NewWafGetTagsResponse

`func NewWafGetTagsResponse() *WafGetTagsResponse`

NewWafGetTagsResponse instantiates a new WafGetTagsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafGetTagsResponseWithDefaults

`func NewWafGetTagsResponseWithDefaults() *WafGetTagsResponse`

NewWafGetTagsResponseWithDefaults instantiates a new WafGetTagsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *WafGetTagsResponse) GetResults() []WafTag`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *WafGetTagsResponse) GetResultsOk() (*[]WafTag, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *WafGetTagsResponse) SetResults(v []WafTag)`

SetResults sets Results field to given value.

### HasResults

`func (o *WafGetTagsResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetPageInfo

`func (o *WafGetTagsResponse) GetPageInfo() PaginationPageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *WafGetTagsResponse) GetPageInfoOk() (*PaginationPageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *WafGetTagsResponse) SetPageInfo(v PaginationPageInfo)`

SetPageInfo sets PageInfo field to given value.

### HasPageInfo

`func (o *WafGetTagsResponse) HasPageInfo() bool`

HasPageInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


