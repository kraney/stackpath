# WafTrafficV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to [**time.Time**](time.Time.md) | The date and time traffic metrics were recorded for | [optional] 
**Api** | Pointer to **string** | The number of requests to URLs designated as API URLs | [optional] 
**Ajax** | Pointer to **string** | The number of Ajax requests | [optional] 
**Static** | Pointer to **string** | The number of requests to static assets | [optional] 
**CustomBlocked** | Pointer to **string** | The number of requests blocked by custom rules | [optional] 
**PolicyBlocked** | Pointer to **string** | The number of requests blocked by policies | [optional] 
**DdosBlocked** | Pointer to **string** | The number of requests blocked by layer 7 DDOS protection | [optional] 
**Monitored** | Pointer to **string** | The number of monitored requests | [optional] 
**CustomWhitelisted** | Pointer to **string** | The number of requests allowed by custom rules | [optional] 
**PolicyWhitelisted** | Pointer to **string** | The number of requests from allowed search engines and other allowed bots | [optional] 
**OriginError4xx** | Pointer to **string** | The number of HTTP 4xx errors received from the origin | [optional] 
**OriginError5xx** | Pointer to **string** | The number of HTTP 5xx errors received from the origin | [optional] 
**OriginTimeout** | Pointer to **string** | The number of request timeouts from the WAF to the origin | [optional] 
**Uncategorized** | Pointer to **string** | The number of uncategorized requests | [optional] 
**PassedToOrigin** | Pointer to **string** | The number of requests that were passed to the origin | [optional] 
**ResponseTime** | Pointer to **string** | The average response time to the origin in milliseconds | [optional] 
**Total** | Pointer to **string** | The total number of requests for the given time period | [optional] 

## Methods

### NewWafTrafficV2

`func NewWafTrafficV2() *WafTrafficV2`

NewWafTrafficV2 instantiates a new WafTrafficV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafTrafficV2WithDefaults

`func NewWafTrafficV2WithDefaults() *WafTrafficV2`

NewWafTrafficV2WithDefaults instantiates a new WafTrafficV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *WafTrafficV2) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WafTrafficV2) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WafTrafficV2) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *WafTrafficV2) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetApi

`func (o *WafTrafficV2) GetApi() string`

GetApi returns the Api field if non-nil, zero value otherwise.

### GetApiOk

`func (o *WafTrafficV2) GetApiOk() (*string, bool)`

GetApiOk returns a tuple with the Api field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApi

`func (o *WafTrafficV2) SetApi(v string)`

SetApi sets Api field to given value.

### HasApi

`func (o *WafTrafficV2) HasApi() bool`

HasApi returns a boolean if a field has been set.

### GetAjax

`func (o *WafTrafficV2) GetAjax() string`

GetAjax returns the Ajax field if non-nil, zero value otherwise.

### GetAjaxOk

`func (o *WafTrafficV2) GetAjaxOk() (*string, bool)`

GetAjaxOk returns a tuple with the Ajax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAjax

`func (o *WafTrafficV2) SetAjax(v string)`

SetAjax sets Ajax field to given value.

### HasAjax

`func (o *WafTrafficV2) HasAjax() bool`

HasAjax returns a boolean if a field has been set.

### GetStatic

`func (o *WafTrafficV2) GetStatic() string`

GetStatic returns the Static field if non-nil, zero value otherwise.

### GetStaticOk

`func (o *WafTrafficV2) GetStaticOk() (*string, bool)`

GetStaticOk returns a tuple with the Static field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatic

`func (o *WafTrafficV2) SetStatic(v string)`

SetStatic sets Static field to given value.

### HasStatic

`func (o *WafTrafficV2) HasStatic() bool`

HasStatic returns a boolean if a field has been set.

### GetCustomBlocked

`func (o *WafTrafficV2) GetCustomBlocked() string`

GetCustomBlocked returns the CustomBlocked field if non-nil, zero value otherwise.

### GetCustomBlockedOk

`func (o *WafTrafficV2) GetCustomBlockedOk() (*string, bool)`

GetCustomBlockedOk returns a tuple with the CustomBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomBlocked

`func (o *WafTrafficV2) SetCustomBlocked(v string)`

SetCustomBlocked sets CustomBlocked field to given value.

### HasCustomBlocked

`func (o *WafTrafficV2) HasCustomBlocked() bool`

HasCustomBlocked returns a boolean if a field has been set.

### GetPolicyBlocked

`func (o *WafTrafficV2) GetPolicyBlocked() string`

GetPolicyBlocked returns the PolicyBlocked field if non-nil, zero value otherwise.

### GetPolicyBlockedOk

`func (o *WafTrafficV2) GetPolicyBlockedOk() (*string, bool)`

GetPolicyBlockedOk returns a tuple with the PolicyBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyBlocked

`func (o *WafTrafficV2) SetPolicyBlocked(v string)`

SetPolicyBlocked sets PolicyBlocked field to given value.

### HasPolicyBlocked

`func (o *WafTrafficV2) HasPolicyBlocked() bool`

HasPolicyBlocked returns a boolean if a field has been set.

### GetDdosBlocked

`func (o *WafTrafficV2) GetDdosBlocked() string`

GetDdosBlocked returns the DdosBlocked field if non-nil, zero value otherwise.

### GetDdosBlockedOk

`func (o *WafTrafficV2) GetDdosBlockedOk() (*string, bool)`

GetDdosBlockedOk returns a tuple with the DdosBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdosBlocked

`func (o *WafTrafficV2) SetDdosBlocked(v string)`

SetDdosBlocked sets DdosBlocked field to given value.

### HasDdosBlocked

`func (o *WafTrafficV2) HasDdosBlocked() bool`

HasDdosBlocked returns a boolean if a field has been set.

### GetMonitored

`func (o *WafTrafficV2) GetMonitored() string`

GetMonitored returns the Monitored field if non-nil, zero value otherwise.

### GetMonitoredOk

`func (o *WafTrafficV2) GetMonitoredOk() (*string, bool)`

GetMonitoredOk returns a tuple with the Monitored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitored

`func (o *WafTrafficV2) SetMonitored(v string)`

SetMonitored sets Monitored field to given value.

### HasMonitored

`func (o *WafTrafficV2) HasMonitored() bool`

HasMonitored returns a boolean if a field has been set.

### GetCustomWhitelisted

`func (o *WafTrafficV2) GetCustomWhitelisted() string`

GetCustomWhitelisted returns the CustomWhitelisted field if non-nil, zero value otherwise.

### GetCustomWhitelistedOk

`func (o *WafTrafficV2) GetCustomWhitelistedOk() (*string, bool)`

GetCustomWhitelistedOk returns a tuple with the CustomWhitelisted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomWhitelisted

`func (o *WafTrafficV2) SetCustomWhitelisted(v string)`

SetCustomWhitelisted sets CustomWhitelisted field to given value.

### HasCustomWhitelisted

`func (o *WafTrafficV2) HasCustomWhitelisted() bool`

HasCustomWhitelisted returns a boolean if a field has been set.

### GetPolicyWhitelisted

`func (o *WafTrafficV2) GetPolicyWhitelisted() string`

GetPolicyWhitelisted returns the PolicyWhitelisted field if non-nil, zero value otherwise.

### GetPolicyWhitelistedOk

`func (o *WafTrafficV2) GetPolicyWhitelistedOk() (*string, bool)`

GetPolicyWhitelistedOk returns a tuple with the PolicyWhitelisted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyWhitelisted

`func (o *WafTrafficV2) SetPolicyWhitelisted(v string)`

SetPolicyWhitelisted sets PolicyWhitelisted field to given value.

### HasPolicyWhitelisted

`func (o *WafTrafficV2) HasPolicyWhitelisted() bool`

HasPolicyWhitelisted returns a boolean if a field has been set.

### GetOriginError4xx

`func (o *WafTrafficV2) GetOriginError4xx() string`

GetOriginError4xx returns the OriginError4xx field if non-nil, zero value otherwise.

### GetOriginError4xxOk

`func (o *WafTrafficV2) GetOriginError4xxOk() (*string, bool)`

GetOriginError4xxOk returns a tuple with the OriginError4xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginError4xx

`func (o *WafTrafficV2) SetOriginError4xx(v string)`

SetOriginError4xx sets OriginError4xx field to given value.

### HasOriginError4xx

`func (o *WafTrafficV2) HasOriginError4xx() bool`

HasOriginError4xx returns a boolean if a field has been set.

### GetOriginError5xx

`func (o *WafTrafficV2) GetOriginError5xx() string`

GetOriginError5xx returns the OriginError5xx field if non-nil, zero value otherwise.

### GetOriginError5xxOk

`func (o *WafTrafficV2) GetOriginError5xxOk() (*string, bool)`

GetOriginError5xxOk returns a tuple with the OriginError5xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginError5xx

`func (o *WafTrafficV2) SetOriginError5xx(v string)`

SetOriginError5xx sets OriginError5xx field to given value.

### HasOriginError5xx

`func (o *WafTrafficV2) HasOriginError5xx() bool`

HasOriginError5xx returns a boolean if a field has been set.

### GetOriginTimeout

`func (o *WafTrafficV2) GetOriginTimeout() string`

GetOriginTimeout returns the OriginTimeout field if non-nil, zero value otherwise.

### GetOriginTimeoutOk

`func (o *WafTrafficV2) GetOriginTimeoutOk() (*string, bool)`

GetOriginTimeoutOk returns a tuple with the OriginTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginTimeout

`func (o *WafTrafficV2) SetOriginTimeout(v string)`

SetOriginTimeout sets OriginTimeout field to given value.

### HasOriginTimeout

`func (o *WafTrafficV2) HasOriginTimeout() bool`

HasOriginTimeout returns a boolean if a field has been set.

### GetUncategorized

`func (o *WafTrafficV2) GetUncategorized() string`

GetUncategorized returns the Uncategorized field if non-nil, zero value otherwise.

### GetUncategorizedOk

`func (o *WafTrafficV2) GetUncategorizedOk() (*string, bool)`

GetUncategorizedOk returns a tuple with the Uncategorized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncategorized

`func (o *WafTrafficV2) SetUncategorized(v string)`

SetUncategorized sets Uncategorized field to given value.

### HasUncategorized

`func (o *WafTrafficV2) HasUncategorized() bool`

HasUncategorized returns a boolean if a field has been set.

### GetPassedToOrigin

`func (o *WafTrafficV2) GetPassedToOrigin() string`

GetPassedToOrigin returns the PassedToOrigin field if non-nil, zero value otherwise.

### GetPassedToOriginOk

`func (o *WafTrafficV2) GetPassedToOriginOk() (*string, bool)`

GetPassedToOriginOk returns a tuple with the PassedToOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassedToOrigin

`func (o *WafTrafficV2) SetPassedToOrigin(v string)`

SetPassedToOrigin sets PassedToOrigin field to given value.

### HasPassedToOrigin

`func (o *WafTrafficV2) HasPassedToOrigin() bool`

HasPassedToOrigin returns a boolean if a field has been set.

### GetResponseTime

`func (o *WafTrafficV2) GetResponseTime() string`

GetResponseTime returns the ResponseTime field if non-nil, zero value otherwise.

### GetResponseTimeOk

`func (o *WafTrafficV2) GetResponseTimeOk() (*string, bool)`

GetResponseTimeOk returns a tuple with the ResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTime

`func (o *WafTrafficV2) SetResponseTime(v string)`

SetResponseTime sets ResponseTime field to given value.

### HasResponseTime

`func (o *WafTrafficV2) HasResponseTime() bool`

HasResponseTime returns a boolean if a field has been set.

### GetTotal

`func (o *WafTrafficV2) GetTotal() string`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *WafTrafficV2) GetTotalOk() (*string, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *WafTrafficV2) SetTotal(v string)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *WafTrafficV2) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


