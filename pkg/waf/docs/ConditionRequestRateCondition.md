# ConditionRequestRateCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ips** | Pointer to **[]string** | A list of source IPs that can trigger a request rate condition | [optional] 
**HttpMethods** | Pointer to [**[]WafHttpMethod**](wafHttpMethod.md) | Possible HTTP request methods that can trigger a request rate condition | [optional] 
**PathPattern** | Pointer to **string** | A regular expression matching the URL path of the incoming request | [optional] 
**Requests** | Pointer to **string** | The number of incoming requests over the given time that can trigger a request rate condition | [optional] 
**Time** | Pointer to **string** | The number of seconds that the WAF measures incoming requests over before triggering a request rate condition | [optional] 

## Methods

### NewConditionRequestRateCondition

`func NewConditionRequestRateCondition() *ConditionRequestRateCondition`

NewConditionRequestRateCondition instantiates a new ConditionRequestRateCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionRequestRateConditionWithDefaults

`func NewConditionRequestRateConditionWithDefaults() *ConditionRequestRateCondition`

NewConditionRequestRateConditionWithDefaults instantiates a new ConditionRequestRateCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIps

`func (o *ConditionRequestRateCondition) GetIps() []string`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *ConditionRequestRateCondition) GetIpsOk() (*[]string, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *ConditionRequestRateCondition) SetIps(v []string)`

SetIps sets Ips field to given value.

### HasIps

`func (o *ConditionRequestRateCondition) HasIps() bool`

HasIps returns a boolean if a field has been set.

### GetHttpMethods

`func (o *ConditionRequestRateCondition) GetHttpMethods() []WafHttpMethod`

GetHttpMethods returns the HttpMethods field if non-nil, zero value otherwise.

### GetHttpMethodsOk

`func (o *ConditionRequestRateCondition) GetHttpMethodsOk() (*[]WafHttpMethod, bool)`

GetHttpMethodsOk returns a tuple with the HttpMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpMethods

`func (o *ConditionRequestRateCondition) SetHttpMethods(v []WafHttpMethod)`

SetHttpMethods sets HttpMethods field to given value.

### HasHttpMethods

`func (o *ConditionRequestRateCondition) HasHttpMethods() bool`

HasHttpMethods returns a boolean if a field has been set.

### GetPathPattern

`func (o *ConditionRequestRateCondition) GetPathPattern() string`

GetPathPattern returns the PathPattern field if non-nil, zero value otherwise.

### GetPathPatternOk

`func (o *ConditionRequestRateCondition) GetPathPatternOk() (*string, bool)`

GetPathPatternOk returns a tuple with the PathPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathPattern

`func (o *ConditionRequestRateCondition) SetPathPattern(v string)`

SetPathPattern sets PathPattern field to given value.

### HasPathPattern

`func (o *ConditionRequestRateCondition) HasPathPattern() bool`

HasPathPattern returns a boolean if a field has been set.

### GetRequests

`func (o *ConditionRequestRateCondition) GetRequests() string`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *ConditionRequestRateCondition) GetRequestsOk() (*string, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *ConditionRequestRateCondition) SetRequests(v string)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *ConditionRequestRateCondition) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetTime

`func (o *ConditionRequestRateCondition) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ConditionRequestRateCondition) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ConditionRequestRateCondition) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *ConditionRequestRateCondition) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


