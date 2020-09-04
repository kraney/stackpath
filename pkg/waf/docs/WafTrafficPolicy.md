# WafTrafficPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AntiautomationAndBehavioral** | Pointer to **string** | The number of requests either handled or blocked by anti_automation and behavioral templates | [optional] 
**Csrf** | Pointer to **string** | The number of requests either handled or blocked by the CSRF template | [optional] 
**IpReputation** | Pointer to **string** | The number of requests either handled or blocked by the IP reputation template | [optional] 
**SpamAndAbuse** | Pointer to **string** | The number of requests either handled or blocked by the spam and abuse template | [optional] 
**Waf** | Pointer to **string** | The number of requests either handled or blocked by REM templates | [optional] 

## Methods

### NewWafTrafficPolicy

`func NewWafTrafficPolicy() *WafTrafficPolicy`

NewWafTrafficPolicy instantiates a new WafTrafficPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafTrafficPolicyWithDefaults

`func NewWafTrafficPolicyWithDefaults() *WafTrafficPolicy`

NewWafTrafficPolicyWithDefaults instantiates a new WafTrafficPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAntiautomationAndBehavioral

`func (o *WafTrafficPolicy) GetAntiautomationAndBehavioral() string`

GetAntiautomationAndBehavioral returns the AntiautomationAndBehavioral field if non-nil, zero value otherwise.

### GetAntiautomationAndBehavioralOk

`func (o *WafTrafficPolicy) GetAntiautomationAndBehavioralOk() (*string, bool)`

GetAntiautomationAndBehavioralOk returns a tuple with the AntiautomationAndBehavioral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntiautomationAndBehavioral

`func (o *WafTrafficPolicy) SetAntiautomationAndBehavioral(v string)`

SetAntiautomationAndBehavioral sets AntiautomationAndBehavioral field to given value.

### HasAntiautomationAndBehavioral

`func (o *WafTrafficPolicy) HasAntiautomationAndBehavioral() bool`

HasAntiautomationAndBehavioral returns a boolean if a field has been set.

### GetCsrf

`func (o *WafTrafficPolicy) GetCsrf() string`

GetCsrf returns the Csrf field if non-nil, zero value otherwise.

### GetCsrfOk

`func (o *WafTrafficPolicy) GetCsrfOk() (*string, bool)`

GetCsrfOk returns a tuple with the Csrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsrf

`func (o *WafTrafficPolicy) SetCsrf(v string)`

SetCsrf sets Csrf field to given value.

### HasCsrf

`func (o *WafTrafficPolicy) HasCsrf() bool`

HasCsrf returns a boolean if a field has been set.

### GetIpReputation

`func (o *WafTrafficPolicy) GetIpReputation() string`

GetIpReputation returns the IpReputation field if non-nil, zero value otherwise.

### GetIpReputationOk

`func (o *WafTrafficPolicy) GetIpReputationOk() (*string, bool)`

GetIpReputationOk returns a tuple with the IpReputation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpReputation

`func (o *WafTrafficPolicy) SetIpReputation(v string)`

SetIpReputation sets IpReputation field to given value.

### HasIpReputation

`func (o *WafTrafficPolicy) HasIpReputation() bool`

HasIpReputation returns a boolean if a field has been set.

### GetSpamAndAbuse

`func (o *WafTrafficPolicy) GetSpamAndAbuse() string`

GetSpamAndAbuse returns the SpamAndAbuse field if non-nil, zero value otherwise.

### GetSpamAndAbuseOk

`func (o *WafTrafficPolicy) GetSpamAndAbuseOk() (*string, bool)`

GetSpamAndAbuseOk returns a tuple with the SpamAndAbuse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpamAndAbuse

`func (o *WafTrafficPolicy) SetSpamAndAbuse(v string)`

SetSpamAndAbuse sets SpamAndAbuse field to given value.

### HasSpamAndAbuse

`func (o *WafTrafficPolicy) HasSpamAndAbuse() bool`

HasSpamAndAbuse returns a boolean if a field has been set.

### GetWaf

`func (o *WafTrafficPolicy) GetWaf() string`

GetWaf returns the Waf field if non-nil, zero value otherwise.

### GetWafOk

`func (o *WafTrafficPolicy) GetWafOk() (*string, bool)`

GetWafOk returns a tuple with the Waf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaf

`func (o *WafTrafficPolicy) SetWaf(v string)`

SetWaf sets Waf field to given value.

### HasWaf

`func (o *WafTrafficPolicy) HasWaf() bool`

HasWaf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


