# V1GetLocationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageInfo** | Pointer to [**PaginationPageInfo**](paginationPageInfo.md) |  | [optional] 
**Results** | Pointer to [**[]Workloadv1Location**](workloadv1Location.md) | The requested locations | [optional] 

## Methods

### NewV1GetLocationsResponse

`func NewV1GetLocationsResponse() *V1GetLocationsResponse`

NewV1GetLocationsResponse instantiates a new V1GetLocationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1GetLocationsResponseWithDefaults

`func NewV1GetLocationsResponseWithDefaults() *V1GetLocationsResponse`

NewV1GetLocationsResponseWithDefaults instantiates a new V1GetLocationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageInfo

`func (o *V1GetLocationsResponse) GetPageInfo() PaginationPageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *V1GetLocationsResponse) GetPageInfoOk() (*PaginationPageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *V1GetLocationsResponse) SetPageInfo(v PaginationPageInfo)`

SetPageInfo sets PageInfo field to given value.

### HasPageInfo

`func (o *V1GetLocationsResponse) HasPageInfo() bool`

HasPageInfo returns a boolean if a field has been set.

### GetResults

`func (o *V1GetLocationsResponse) GetResults() []Workloadv1Location`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *V1GetLocationsResponse) GetResultsOk() (*[]Workloadv1Location, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *V1GetLocationsResponse) SetResults(v []Workloadv1Location)`

SetResults sets Results field to given value.

### HasResults

`func (o *V1GetLocationsResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


