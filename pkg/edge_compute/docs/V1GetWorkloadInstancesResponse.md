# V1GetWorkloadInstancesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageInfo** | Pointer to [**PaginationPageInfo**](paginationPageInfo.md) |  | [optional] 
**Results** | Pointer to [**[]Workloadv1Instance**](workloadv1Instance.md) | The requested workload instances | [optional] 

## Methods

### NewV1GetWorkloadInstancesResponse

`func NewV1GetWorkloadInstancesResponse() *V1GetWorkloadInstancesResponse`

NewV1GetWorkloadInstancesResponse instantiates a new V1GetWorkloadInstancesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1GetWorkloadInstancesResponseWithDefaults

`func NewV1GetWorkloadInstancesResponseWithDefaults() *V1GetWorkloadInstancesResponse`

NewV1GetWorkloadInstancesResponseWithDefaults instantiates a new V1GetWorkloadInstancesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageInfo

`func (o *V1GetWorkloadInstancesResponse) GetPageInfo() PaginationPageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *V1GetWorkloadInstancesResponse) GetPageInfoOk() (*PaginationPageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *V1GetWorkloadInstancesResponse) SetPageInfo(v PaginationPageInfo)`

SetPageInfo sets PageInfo field to given value.

### HasPageInfo

`func (o *V1GetWorkloadInstancesResponse) HasPageInfo() bool`

HasPageInfo returns a boolean if a field has been set.

### GetResults

`func (o *V1GetWorkloadInstancesResponse) GetResults() []Workloadv1Instance`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *V1GetWorkloadInstancesResponse) GetResultsOk() (*[]Workloadv1Instance, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *V1GetWorkloadInstancesResponse) SetResults(v []Workloadv1Instance)`

SetResults sets Results field to given value.

### HasResults

`func (o *V1GetWorkloadInstancesResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


