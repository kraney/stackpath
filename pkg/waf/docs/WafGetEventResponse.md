# WafGetEventResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | Pointer to [**WafEvent**](wafEvent.md) |  | [optional] 

## Methods

### NewWafGetEventResponse

`func NewWafGetEventResponse() *WafGetEventResponse`

NewWafGetEventResponse instantiates a new WafGetEventResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafGetEventResponseWithDefaults

`func NewWafGetEventResponseWithDefaults() *WafGetEventResponse`

NewWafGetEventResponseWithDefaults instantiates a new WafGetEventResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *WafGetEventResponse) GetEvent() WafEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *WafGetEventResponse) GetEventOk() (*WafEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *WafGetEventResponse) SetEvent(v WafEvent)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *WafGetEventResponse) HasEvent() bool`

HasEvent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


