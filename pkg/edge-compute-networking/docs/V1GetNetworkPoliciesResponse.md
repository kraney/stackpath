# V1GetNetworkPoliciesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageInfo** | Pointer to [**PaginationPageInfo**](paginationPageInfo.md) |  | [optional] 
**Results** | Pointer to [**[]V1NetworkPolicy**](v1NetworkPolicy.md) | The requested network policies | [optional] 

## Methods

### NewV1GetNetworkPoliciesResponse

`func NewV1GetNetworkPoliciesResponse() *V1GetNetworkPoliciesResponse`

NewV1GetNetworkPoliciesResponse instantiates a new V1GetNetworkPoliciesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1GetNetworkPoliciesResponseWithDefaults

`func NewV1GetNetworkPoliciesResponseWithDefaults() *V1GetNetworkPoliciesResponse`

NewV1GetNetworkPoliciesResponseWithDefaults instantiates a new V1GetNetworkPoliciesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageInfo

`func (o *V1GetNetworkPoliciesResponse) GetPageInfo() PaginationPageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *V1GetNetworkPoliciesResponse) GetPageInfoOk() (*PaginationPageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *V1GetNetworkPoliciesResponse) SetPageInfo(v PaginationPageInfo)`

SetPageInfo sets PageInfo field to given value.

### HasPageInfo

`func (o *V1GetNetworkPoliciesResponse) HasPageInfo() bool`

HasPageInfo returns a boolean if a field has been set.

### GetResults

`func (o *V1GetNetworkPoliciesResponse) GetResults() []V1NetworkPolicy`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *V1GetNetworkPoliciesResponse) GetResultsOk() (*[]V1NetworkPolicy, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *V1GetNetworkPoliciesResponse) SetResults(v []V1NetworkPolicy)`

SetResults sets Results field to given value.

### HasResults

`func (o *V1GetNetworkPoliciesResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


