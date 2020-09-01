# WafGetTrafficV2Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Traffic** | Pointer to [**[]WafTrafficV2**](wafTrafficV2.md) | The requested WAF traffic | [optional] 

## Methods

### NewWafGetTrafficV2Response

`func NewWafGetTrafficV2Response() *WafGetTrafficV2Response`

NewWafGetTrafficV2Response instantiates a new WafGetTrafficV2Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafGetTrafficV2ResponseWithDefaults

`func NewWafGetTrafficV2ResponseWithDefaults() *WafGetTrafficV2Response`

NewWafGetTrafficV2ResponseWithDefaults instantiates a new WafGetTrafficV2Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTraffic

`func (o *WafGetTrafficV2Response) GetTraffic() []WafTrafficV2`

GetTraffic returns the Traffic field if non-nil, zero value otherwise.

### GetTrafficOk

`func (o *WafGetTrafficV2Response) GetTrafficOk() (*[]WafTrafficV2, bool)`

GetTrafficOk returns a tuple with the Traffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraffic

`func (o *WafGetTrafficV2Response) SetTraffic(v []WafTrafficV2)`

SetTraffic sets Traffic field to given value.

### HasTraffic

`func (o *WafGetTrafficV2Response) HasTraffic() bool`

HasTraffic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


