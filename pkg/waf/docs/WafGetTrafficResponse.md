# WafGetTrafficResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Traffic** | Pointer to [**[]WafTraffic**](wafTraffic.md) | The requested WAF traffic | [optional] 

## Methods

### NewWafGetTrafficResponse

`func NewWafGetTrafficResponse() *WafGetTrafficResponse`

NewWafGetTrafficResponse instantiates a new WafGetTrafficResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafGetTrafficResponseWithDefaults

`func NewWafGetTrafficResponseWithDefaults() *WafGetTrafficResponse`

NewWafGetTrafficResponseWithDefaults instantiates a new WafGetTrafficResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTraffic

`func (o *WafGetTrafficResponse) GetTraffic() []WafTraffic`

GetTraffic returns the Traffic field if non-nil, zero value otherwise.

### GetTrafficOk

`func (o *WafGetTrafficResponse) GetTrafficOk() (*[]WafTraffic, bool)`

GetTrafficOk returns a tuple with the Traffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraffic

`func (o *WafGetTrafficResponse) SetTraffic(v []WafTraffic)`

SetTraffic sets Traffic field to given value.

### HasTraffic

`func (o *WafGetTrafficResponse) HasTraffic() bool`

HasTraffic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


