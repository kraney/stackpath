# WafGetSiteCertificatesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageInfo** | Pointer to [**PaginationPageInfo**](paginationPageInfo.md) |  | [optional] 
**Results** | Pointer to [**[]WafSiteCertificate**](wafSiteCertificate.md) | The requested SSL certificates | [optional] 

## Methods

### NewWafGetSiteCertificatesResponse

`func NewWafGetSiteCertificatesResponse() *WafGetSiteCertificatesResponse`

NewWafGetSiteCertificatesResponse instantiates a new WafGetSiteCertificatesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafGetSiteCertificatesResponseWithDefaults

`func NewWafGetSiteCertificatesResponseWithDefaults() *WafGetSiteCertificatesResponse`

NewWafGetSiteCertificatesResponseWithDefaults instantiates a new WafGetSiteCertificatesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageInfo

`func (o *WafGetSiteCertificatesResponse) GetPageInfo() PaginationPageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *WafGetSiteCertificatesResponse) GetPageInfoOk() (*PaginationPageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *WafGetSiteCertificatesResponse) SetPageInfo(v PaginationPageInfo)`

SetPageInfo sets PageInfo field to given value.

### HasPageInfo

`func (o *WafGetSiteCertificatesResponse) HasPageInfo() bool`

HasPageInfo returns a boolean if a field has been set.

### GetResults

`func (o *WafGetSiteCertificatesResponse) GetResults() []WafSiteCertificate`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *WafGetSiteCertificatesResponse) GetResultsOk() (*[]WafSiteCertificate, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *WafGetSiteCertificatesResponse) SetResults(v []WafSiteCertificate)`

SetResults sets Results field to given value.

### HasResults

`func (o *WafGetSiteCertificatesResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


