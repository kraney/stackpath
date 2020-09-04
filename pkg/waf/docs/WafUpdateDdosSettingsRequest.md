# WafUpdateDdosSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GlobalThreshold** | Pointer to **string** | The number of requests per ten seconds that can trigger DDOS protection | [optional] 
**BurstThreshold** | Pointer to **string** | The number of requests per two seconds that can trigger DDOS protection | [optional] 
**SubSecondBurstThreshold** | Pointer to **string** | The number of requests per 0.1 seconds that can trigger DDOS protection | [optional] 

## Methods

### NewWafUpdateDdosSettingsRequest

`func NewWafUpdateDdosSettingsRequest() *WafUpdateDdosSettingsRequest`

NewWafUpdateDdosSettingsRequest instantiates a new WafUpdateDdosSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafUpdateDdosSettingsRequestWithDefaults

`func NewWafUpdateDdosSettingsRequestWithDefaults() *WafUpdateDdosSettingsRequest`

NewWafUpdateDdosSettingsRequestWithDefaults instantiates a new WafUpdateDdosSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGlobalThreshold

`func (o *WafUpdateDdosSettingsRequest) GetGlobalThreshold() string`

GetGlobalThreshold returns the GlobalThreshold field if non-nil, zero value otherwise.

### GetGlobalThresholdOk

`func (o *WafUpdateDdosSettingsRequest) GetGlobalThresholdOk() (*string, bool)`

GetGlobalThresholdOk returns a tuple with the GlobalThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalThreshold

`func (o *WafUpdateDdosSettingsRequest) SetGlobalThreshold(v string)`

SetGlobalThreshold sets GlobalThreshold field to given value.

### HasGlobalThreshold

`func (o *WafUpdateDdosSettingsRequest) HasGlobalThreshold() bool`

HasGlobalThreshold returns a boolean if a field has been set.

### GetBurstThreshold

`func (o *WafUpdateDdosSettingsRequest) GetBurstThreshold() string`

GetBurstThreshold returns the BurstThreshold field if non-nil, zero value otherwise.

### GetBurstThresholdOk

`func (o *WafUpdateDdosSettingsRequest) GetBurstThresholdOk() (*string, bool)`

GetBurstThresholdOk returns a tuple with the BurstThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurstThreshold

`func (o *WafUpdateDdosSettingsRequest) SetBurstThreshold(v string)`

SetBurstThreshold sets BurstThreshold field to given value.

### HasBurstThreshold

`func (o *WafUpdateDdosSettingsRequest) HasBurstThreshold() bool`

HasBurstThreshold returns a boolean if a field has been set.

### GetSubSecondBurstThreshold

`func (o *WafUpdateDdosSettingsRequest) GetSubSecondBurstThreshold() string`

GetSubSecondBurstThreshold returns the SubSecondBurstThreshold field if non-nil, zero value otherwise.

### GetSubSecondBurstThresholdOk

`func (o *WafUpdateDdosSettingsRequest) GetSubSecondBurstThresholdOk() (*string, bool)`

GetSubSecondBurstThresholdOk returns a tuple with the SubSecondBurstThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubSecondBurstThreshold

`func (o *WafUpdateDdosSettingsRequest) SetSubSecondBurstThreshold(v string)`

SetSubSecondBurstThreshold sets SubSecondBurstThreshold field to given value.

### HasSubSecondBurstThreshold

`func (o *WafUpdateDdosSettingsRequest) HasSubSecondBurstThreshold() bool`

HasSubSecondBurstThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


