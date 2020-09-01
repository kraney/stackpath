# WafGetRequestDetailsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestDetails** | Pointer to [**WafRequestDetails**](wafRequestDetails.md) |  | [optional] 

## Methods

### NewWafGetRequestDetailsResponse

`func NewWafGetRequestDetailsResponse() *WafGetRequestDetailsResponse`

NewWafGetRequestDetailsResponse instantiates a new WafGetRequestDetailsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafGetRequestDetailsResponseWithDefaults

`func NewWafGetRequestDetailsResponseWithDefaults() *WafGetRequestDetailsResponse`

NewWafGetRequestDetailsResponseWithDefaults instantiates a new WafGetRequestDetailsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestDetails

`func (o *WafGetRequestDetailsResponse) GetRequestDetails() WafRequestDetails`

GetRequestDetails returns the RequestDetails field if non-nil, zero value otherwise.

### GetRequestDetailsOk

`func (o *WafGetRequestDetailsResponse) GetRequestDetailsOk() (*WafRequestDetails, bool)`

GetRequestDetailsOk returns a tuple with the RequestDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestDetails

`func (o *WafGetRequestDetailsResponse) SetRequestDetails(v WafRequestDetails)`

SetRequestDetails sets RequestDetails field to given value.

### HasRequestDetails

`func (o *WafGetRequestDetailsResponse) HasRequestDetails() bool`

HasRequestDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


