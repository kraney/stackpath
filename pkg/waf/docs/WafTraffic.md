# WafTraffic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to [**time.Time**](time.Time.md) | The date and time traffic metrics were recorded for | [optional] 
**PolicyCount** | Pointer to [**WafTrafficPolicy**](wafTrafficPolicy.md) |  | [optional] 
**PolicyBlocked** | Pointer to [**WafTrafficPolicy**](wafTrafficPolicy.md) |  | [optional] 
**RuleCount** | Pointer to **string** | The number of requests analyzed by rules | [optional] 
**RuleBlocked** | Pointer to **string** | The number of requests blocked by rules | [optional] 
**PassedToOrigin** | Pointer to [**TrafficPassedToOrigin**](TrafficPassedToOrigin.md) |  | [optional] 
**Ddos** | Pointer to **string** | The number of requests handled by the layer 7 DDOS engine | [optional] 
**OriginError4xx** | Pointer to **string** | The number of HTTP 4xx errors received from the origin | [optional] 
**OriginError5xx** | Pointer to **string** | The number of HTTP 5xx errors received from the origin | [optional] 
**OriginTimeout** | Pointer to **string** | The number of request timeouts from the WAF to the origin | [optional] 
**ResponseTime** | Pointer to **string** | Average response time to the origin in milliseconds | [optional] 

## Methods

### NewWafTraffic

`func NewWafTraffic() *WafTraffic`

NewWafTraffic instantiates a new WafTraffic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafTrafficWithDefaults

`func NewWafTrafficWithDefaults() *WafTraffic`

NewWafTrafficWithDefaults instantiates a new WafTraffic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *WafTraffic) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WafTraffic) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WafTraffic) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *WafTraffic) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetPolicyCount

`func (o *WafTraffic) GetPolicyCount() WafTrafficPolicy`

GetPolicyCount returns the PolicyCount field if non-nil, zero value otherwise.

### GetPolicyCountOk

`func (o *WafTraffic) GetPolicyCountOk() (*WafTrafficPolicy, bool)`

GetPolicyCountOk returns a tuple with the PolicyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyCount

`func (o *WafTraffic) SetPolicyCount(v WafTrafficPolicy)`

SetPolicyCount sets PolicyCount field to given value.

### HasPolicyCount

`func (o *WafTraffic) HasPolicyCount() bool`

HasPolicyCount returns a boolean if a field has been set.

### GetPolicyBlocked

`func (o *WafTraffic) GetPolicyBlocked() WafTrafficPolicy`

GetPolicyBlocked returns the PolicyBlocked field if non-nil, zero value otherwise.

### GetPolicyBlockedOk

`func (o *WafTraffic) GetPolicyBlockedOk() (*WafTrafficPolicy, bool)`

GetPolicyBlockedOk returns a tuple with the PolicyBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyBlocked

`func (o *WafTraffic) SetPolicyBlocked(v WafTrafficPolicy)`

SetPolicyBlocked sets PolicyBlocked field to given value.

### HasPolicyBlocked

`func (o *WafTraffic) HasPolicyBlocked() bool`

HasPolicyBlocked returns a boolean if a field has been set.

### GetRuleCount

`func (o *WafTraffic) GetRuleCount() string`

GetRuleCount returns the RuleCount field if non-nil, zero value otherwise.

### GetRuleCountOk

`func (o *WafTraffic) GetRuleCountOk() (*string, bool)`

GetRuleCountOk returns a tuple with the RuleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleCount

`func (o *WafTraffic) SetRuleCount(v string)`

SetRuleCount sets RuleCount field to given value.

### HasRuleCount

`func (o *WafTraffic) HasRuleCount() bool`

HasRuleCount returns a boolean if a field has been set.

### GetRuleBlocked

`func (o *WafTraffic) GetRuleBlocked() string`

GetRuleBlocked returns the RuleBlocked field if non-nil, zero value otherwise.

### GetRuleBlockedOk

`func (o *WafTraffic) GetRuleBlockedOk() (*string, bool)`

GetRuleBlockedOk returns a tuple with the RuleBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleBlocked

`func (o *WafTraffic) SetRuleBlocked(v string)`

SetRuleBlocked sets RuleBlocked field to given value.

### HasRuleBlocked

`func (o *WafTraffic) HasRuleBlocked() bool`

HasRuleBlocked returns a boolean if a field has been set.

### GetPassedToOrigin

`func (o *WafTraffic) GetPassedToOrigin() TrafficPassedToOrigin`

GetPassedToOrigin returns the PassedToOrigin field if non-nil, zero value otherwise.

### GetPassedToOriginOk

`func (o *WafTraffic) GetPassedToOriginOk() (*TrafficPassedToOrigin, bool)`

GetPassedToOriginOk returns a tuple with the PassedToOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassedToOrigin

`func (o *WafTraffic) SetPassedToOrigin(v TrafficPassedToOrigin)`

SetPassedToOrigin sets PassedToOrigin field to given value.

### HasPassedToOrigin

`func (o *WafTraffic) HasPassedToOrigin() bool`

HasPassedToOrigin returns a boolean if a field has been set.

### GetDdos

`func (o *WafTraffic) GetDdos() string`

GetDdos returns the Ddos field if non-nil, zero value otherwise.

### GetDdosOk

`func (o *WafTraffic) GetDdosOk() (*string, bool)`

GetDdosOk returns a tuple with the Ddos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdos

`func (o *WafTraffic) SetDdos(v string)`

SetDdos sets Ddos field to given value.

### HasDdos

`func (o *WafTraffic) HasDdos() bool`

HasDdos returns a boolean if a field has been set.

### GetOriginError4xx

`func (o *WafTraffic) GetOriginError4xx() string`

GetOriginError4xx returns the OriginError4xx field if non-nil, zero value otherwise.

### GetOriginError4xxOk

`func (o *WafTraffic) GetOriginError4xxOk() (*string, bool)`

GetOriginError4xxOk returns a tuple with the OriginError4xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginError4xx

`func (o *WafTraffic) SetOriginError4xx(v string)`

SetOriginError4xx sets OriginError4xx field to given value.

### HasOriginError4xx

`func (o *WafTraffic) HasOriginError4xx() bool`

HasOriginError4xx returns a boolean if a field has been set.

### GetOriginError5xx

`func (o *WafTraffic) GetOriginError5xx() string`

GetOriginError5xx returns the OriginError5xx field if non-nil, zero value otherwise.

### GetOriginError5xxOk

`func (o *WafTraffic) GetOriginError5xxOk() (*string, bool)`

GetOriginError5xxOk returns a tuple with the OriginError5xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginError5xx

`func (o *WafTraffic) SetOriginError5xx(v string)`

SetOriginError5xx sets OriginError5xx field to given value.

### HasOriginError5xx

`func (o *WafTraffic) HasOriginError5xx() bool`

HasOriginError5xx returns a boolean if a field has been set.

### GetOriginTimeout

`func (o *WafTraffic) GetOriginTimeout() string`

GetOriginTimeout returns the OriginTimeout field if non-nil, zero value otherwise.

### GetOriginTimeoutOk

`func (o *WafTraffic) GetOriginTimeoutOk() (*string, bool)`

GetOriginTimeoutOk returns a tuple with the OriginTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginTimeout

`func (o *WafTraffic) SetOriginTimeout(v string)`

SetOriginTimeout sets OriginTimeout field to given value.

### HasOriginTimeout

`func (o *WafTraffic) HasOriginTimeout() bool`

HasOriginTimeout returns a boolean if a field has been set.

### GetResponseTime

`func (o *WafTraffic) GetResponseTime() string`

GetResponseTime returns the ResponseTime field if non-nil, zero value otherwise.

### GetResponseTimeOk

`func (o *WafTraffic) GetResponseTimeOk() (*string, bool)`

GetResponseTimeOk returns a tuple with the ResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTime

`func (o *WafTraffic) SetResponseTime(v string)`

SetResponseTime sets ResponseTime field to given value.

### HasResponseTime

`func (o *WafTraffic) HasResponseTime() bool`

HasResponseTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


