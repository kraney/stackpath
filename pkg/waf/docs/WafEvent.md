# WafEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A WAF event&#39;s unique identifier | [optional] 
**ReferenceId** | Pointer to **string** | An event&#39;s user-facing identifier  Reference IDs are displayed to the end user when the WAF blocks a request to a site. Please note that an event&#39;s ID and reference ID are different. | [optional] 
**EventDate** | Pointer to [**time.Time**](time.Time.md) | When a WAF event occurred | [optional] 
**Request** | Pointer to [**WafEventRequest**](wafEventRequest.md) |  | [optional] 
**Action** | Pointer to [**EventRuleAction**](EventRuleAction.md) |  | [optional] 
**Client** | Pointer to [**WafEventNetwork**](wafEventNetwork.md) |  | [optional] 
**Count** | Pointer to **string** | Number of events which matched this | [optional] 

## Methods

### NewWafEvent

`func NewWafEvent() *WafEvent`

NewWafEvent instantiates a new WafEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafEventWithDefaults

`func NewWafEventWithDefaults() *WafEvent`

NewWafEventWithDefaults instantiates a new WafEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WafEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WafEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WafEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WafEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReferenceId

`func (o *WafEvent) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *WafEvent) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *WafEvent) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *WafEvent) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetEventDate

`func (o *WafEvent) GetEventDate() time.Time`

GetEventDate returns the EventDate field if non-nil, zero value otherwise.

### GetEventDateOk

`func (o *WafEvent) GetEventDateOk() (*time.Time, bool)`

GetEventDateOk returns a tuple with the EventDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDate

`func (o *WafEvent) SetEventDate(v time.Time)`

SetEventDate sets EventDate field to given value.

### HasEventDate

`func (o *WafEvent) HasEventDate() bool`

HasEventDate returns a boolean if a field has been set.

### GetRequest

`func (o *WafEvent) GetRequest() WafEventRequest`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *WafEvent) GetRequestOk() (*WafEventRequest, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *WafEvent) SetRequest(v WafEventRequest)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *WafEvent) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetAction

`func (o *WafEvent) GetAction() EventRuleAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *WafEvent) GetActionOk() (*EventRuleAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *WafEvent) SetAction(v EventRuleAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *WafEvent) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetClient

`func (o *WafEvent) GetClient() WafEventNetwork`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *WafEvent) GetClientOk() (*WafEventNetwork, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *WafEvent) SetClient(v WafEventNetwork)`

SetClient sets Client field to given value.

### HasClient

`func (o *WafEvent) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetCount

`func (o *WafEvent) GetCount() string`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *WafEvent) GetCountOk() (*string, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *WafEvent) SetCount(v string)`

SetCount sets Count field to given value.

### HasCount

`func (o *WafEvent) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


