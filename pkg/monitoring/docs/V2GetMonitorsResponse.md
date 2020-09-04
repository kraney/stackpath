# V2GetMonitorsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageInfo** | Pointer to [**PaginationPageInfo**](paginationPageInfo.md) |  | [optional] 
**Results** | Pointer to [**[]V2Monitor**](v2Monitor.md) | The requested monitors. | [optional] 

## Methods

### NewV2GetMonitorsResponse

`func NewV2GetMonitorsResponse() *V2GetMonitorsResponse`

NewV2GetMonitorsResponse instantiates a new V2GetMonitorsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2GetMonitorsResponseWithDefaults

`func NewV2GetMonitorsResponseWithDefaults() *V2GetMonitorsResponse`

NewV2GetMonitorsResponseWithDefaults instantiates a new V2GetMonitorsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageInfo

`func (o *V2GetMonitorsResponse) GetPageInfo() PaginationPageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *V2GetMonitorsResponse) GetPageInfoOk() (*PaginationPageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *V2GetMonitorsResponse) SetPageInfo(v PaginationPageInfo)`

SetPageInfo sets PageInfo field to given value.

### HasPageInfo

`func (o *V2GetMonitorsResponse) HasPageInfo() bool`

HasPageInfo returns a boolean if a field has been set.

### GetResults

`func (o *V2GetMonitorsResponse) GetResults() []V2Monitor`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *V2GetMonitorsResponse) GetResultsOk() (*[]V2Monitor, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *V2GetMonitorsResponse) SetResults(v []V2Monitor)`

SetResults sets Results field to given value.

### HasResults

`func (o *V2GetMonitorsResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


