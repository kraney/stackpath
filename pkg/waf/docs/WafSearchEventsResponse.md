# WafSearchEventsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageInfo** | Pointer to [**PaginationPageInfo**](paginationPageInfo.md) |  | [optional] 
**Results** | Pointer to [**[]WafEvent**](wafEvent.md) | The requested WAF events | [optional] 

## Methods

### NewWafSearchEventsResponse

`func NewWafSearchEventsResponse() *WafSearchEventsResponse`

NewWafSearchEventsResponse instantiates a new WafSearchEventsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafSearchEventsResponseWithDefaults

`func NewWafSearchEventsResponseWithDefaults() *WafSearchEventsResponse`

NewWafSearchEventsResponseWithDefaults instantiates a new WafSearchEventsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageInfo

`func (o *WafSearchEventsResponse) GetPageInfo() PaginationPageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *WafSearchEventsResponse) GetPageInfoOk() (*PaginationPageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *WafSearchEventsResponse) SetPageInfo(v PaginationPageInfo)`

SetPageInfo sets PageInfo field to given value.

### HasPageInfo

`func (o *WafSearchEventsResponse) HasPageInfo() bool`

HasPageInfo returns a boolean if a field has been set.

### GetResults

`func (o *WafSearchEventsResponse) GetResults() []WafEvent`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *WafSearchEventsResponse) GetResultsOk() (*[]WafEvent, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *WafSearchEventsResponse) SetResults(v []WafEvent)`

SetResults sets Results field to given value.

### HasResults

`func (o *WafSearchEventsResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


