# V1GetImagesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageInfo** | Pointer to [**PaginationPageInfo**](paginationPageInfo.md) |  | [optional] 
**Results** | Pointer to [**[]V1Image**](v1Image.md) | The requested images | [optional] 

## Methods

### NewV1GetImagesResponse

`func NewV1GetImagesResponse() *V1GetImagesResponse`

NewV1GetImagesResponse instantiates a new V1GetImagesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1GetImagesResponseWithDefaults

`func NewV1GetImagesResponseWithDefaults() *V1GetImagesResponse`

NewV1GetImagesResponseWithDefaults instantiates a new V1GetImagesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageInfo

`func (o *V1GetImagesResponse) GetPageInfo() PaginationPageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *V1GetImagesResponse) GetPageInfoOk() (*PaginationPageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *V1GetImagesResponse) SetPageInfo(v PaginationPageInfo)`

SetPageInfo sets PageInfo field to given value.

### HasPageInfo

`func (o *V1GetImagesResponse) HasPageInfo() bool`

HasPageInfo returns a boolean if a field has been set.

### GetResults

`func (o *V1GetImagesResponse) GetResults() []V1Image`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *V1GetImagesResponse) GetResultsOk() (*[]V1Image, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *V1GetImagesResponse) SetResults(v []V1Image)`

SetResults sets Results field to given value.

### HasResults

`func (o *V1GetImagesResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


