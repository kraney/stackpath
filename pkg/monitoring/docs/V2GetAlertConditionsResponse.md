# V2GetAlertConditionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageInfo** | Pointer to [**PaginationPageInfo**](paginationPageInfo.md) |  | [optional] 
**Results** | Pointer to [**[]V2AlertCondition**](v2AlertCondition.md) | The requested alert conditions. | [optional] 

## Methods

### NewV2GetAlertConditionsResponse

`func NewV2GetAlertConditionsResponse() *V2GetAlertConditionsResponse`

NewV2GetAlertConditionsResponse instantiates a new V2GetAlertConditionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2GetAlertConditionsResponseWithDefaults

`func NewV2GetAlertConditionsResponseWithDefaults() *V2GetAlertConditionsResponse`

NewV2GetAlertConditionsResponseWithDefaults instantiates a new V2GetAlertConditionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageInfo

`func (o *V2GetAlertConditionsResponse) GetPageInfo() PaginationPageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *V2GetAlertConditionsResponse) GetPageInfoOk() (*PaginationPageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *V2GetAlertConditionsResponse) SetPageInfo(v PaginationPageInfo)`

SetPageInfo sets PageInfo field to given value.

### HasPageInfo

`func (o *V2GetAlertConditionsResponse) HasPageInfo() bool`

HasPageInfo returns a boolean if a field has been set.

### GetResults

`func (o *V2GetAlertConditionsResponse) GetResults() []V2AlertCondition`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *V2GetAlertConditionsResponse) GetResultsOk() (*[]V2AlertCondition, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *V2GetAlertConditionsResponse) SetResults(v []V2AlertCondition)`

SetResults sets Results field to given value.

### HasResults

`func (o *V2GetAlertConditionsResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


