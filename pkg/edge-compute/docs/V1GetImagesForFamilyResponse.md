# V1GetImagesForFamilyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageInfo** | Pointer to [**PaginationPageInfo**](paginationPageInfo.md) |  | [optional] 
**Results** | Pointer to [**[]V1Image**](v1Image.md) | The requested images | [optional] 

## Methods

### NewV1GetImagesForFamilyResponse

`func NewV1GetImagesForFamilyResponse() *V1GetImagesForFamilyResponse`

NewV1GetImagesForFamilyResponse instantiates a new V1GetImagesForFamilyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1GetImagesForFamilyResponseWithDefaults

`func NewV1GetImagesForFamilyResponseWithDefaults() *V1GetImagesForFamilyResponse`

NewV1GetImagesForFamilyResponseWithDefaults instantiates a new V1GetImagesForFamilyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageInfo

`func (o *V1GetImagesForFamilyResponse) GetPageInfo() PaginationPageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *V1GetImagesForFamilyResponse) GetPageInfoOk() (*PaginationPageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *V1GetImagesForFamilyResponse) SetPageInfo(v PaginationPageInfo)`

SetPageInfo sets PageInfo field to given value.

### HasPageInfo

`func (o *V1GetImagesForFamilyResponse) HasPageInfo() bool`

HasPageInfo returns a boolean if a field has been set.

### GetResults

`func (o *V1GetImagesForFamilyResponse) GetResults() []V1Image`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *V1GetImagesForFamilyResponse) GetResultsOk() (*[]V1Image, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *V1GetImagesForFamilyResponse) SetResults(v []V1Image)`

SetResults sets Results field to given value.

### HasResults

`func (o *V1GetImagesForFamilyResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


