# WafGetRequestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Request** | Pointer to [**SchemawafRequest**](schemawafRequest.md) |  | [optional] 

## Methods

### NewWafGetRequestResponse

`func NewWafGetRequestResponse() *WafGetRequestResponse`

NewWafGetRequestResponse instantiates a new WafGetRequestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafGetRequestResponseWithDefaults

`func NewWafGetRequestResponseWithDefaults() *WafGetRequestResponse`

NewWafGetRequestResponseWithDefaults instantiates a new WafGetRequestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequest

`func (o *WafGetRequestResponse) GetRequest() SchemawafRequest`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *WafGetRequestResponse) GetRequestOk() (*SchemawafRequest, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *WafGetRequestResponse) SetRequest(v SchemawafRequest)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *WafGetRequestResponse) HasRequest() bool`

HasRequest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


