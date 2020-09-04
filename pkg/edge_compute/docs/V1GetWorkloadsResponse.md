# V1GetWorkloadsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageInfo** | Pointer to [**PaginationPageInfo**](paginationPageInfo.md) |  | [optional] 
**Results** | Pointer to [**[]V1Workload**](v1Workload.md) | The requested workloads | [optional] 

## Methods

### NewV1GetWorkloadsResponse

`func NewV1GetWorkloadsResponse() *V1GetWorkloadsResponse`

NewV1GetWorkloadsResponse instantiates a new V1GetWorkloadsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1GetWorkloadsResponseWithDefaults

`func NewV1GetWorkloadsResponseWithDefaults() *V1GetWorkloadsResponse`

NewV1GetWorkloadsResponseWithDefaults instantiates a new V1GetWorkloadsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageInfo

`func (o *V1GetWorkloadsResponse) GetPageInfo() PaginationPageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *V1GetWorkloadsResponse) GetPageInfoOk() (*PaginationPageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *V1GetWorkloadsResponse) SetPageInfo(v PaginationPageInfo)`

SetPageInfo sets PageInfo field to given value.

### HasPageInfo

`func (o *V1GetWorkloadsResponse) HasPageInfo() bool`

HasPageInfo returns a boolean if a field has been set.

### GetResults

`func (o *V1GetWorkloadsResponse) GetResults() []V1Workload`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *V1GetWorkloadsResponse) GetResultsOk() (*[]V1Workload, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *V1GetWorkloadsResponse) SetResults(v []V1Workload)`

SetResults sets Results field to given value.

### HasResults

`func (o *V1GetWorkloadsResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


