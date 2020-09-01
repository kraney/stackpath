# WafRequestDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The WAF request&#39;s unique identifier | [optional] 
**Path** | Pointer to **string** | The path of the requested URL | [optional] 
**Method** | Pointer to [**WafHttpMethod**](wafHttpMethod.md) |  | [optional] [default to "METHOD_UNSPECIFIED"]
**Action** | Pointer to [**WafRequestAction**](wafRequestAction.md) |  | [optional] [default to "ACTION_UNSPECIFIED"]
**RuleId** | Pointer to **string** | The unique ID of the WAF rule that triggered the action against the WAF request | [optional] 
**RuleName** | Pointer to **string** | The name of the WAF rule triggered by the request | [optional] 
**UserAgent** | Pointer to [**WafRequestDetailsUserAgent**](wafRequestDetailsUserAgent.md) |  | [optional] 
**Network** | Pointer to [**WafRequestDetailsNetwork**](wafRequestDetailsNetwork.md) |  | [optional] 
**RequestTime** | Pointer to [**time.Time**](time.Time.md) | When the WAF request occurred | [optional] 
**ReferenceId** | Pointer to **string** | The WAF request&#39;s user-facing identifier  Reference IDs are displayed to the end user when the WAF blocks a request to a site. Please note that a request&#39;s ID and reference ID are different. | [optional] 
**ContentType** | Pointer to **string** | The content type of the WAF request | [optional] 
**Scheme** | Pointer to **string** | The HTTP scheme of the WAF request | [optional] 
**HttpStatusCode** | Pointer to **string** | The HTTP status code of the WAF request | [optional] 
**HttpVersion** | Pointer to **string** | The HTTP version of the WAF request | [optional] 
**Tags** | Pointer to [**map[string]RequestDetailsTags**](RequestDetailsTags.md) | Aspects of a WAF request | [optional] 
**ResponseTime** | Pointer to **string** | The time it took the WAF to process the WAF request | [optional] 
**RequestHeaders** | Pointer to **map[string]string** | The HTTP request headers of the WAF request | [optional] 
**IncidentId** | Pointer to **string** | The ID of the of challenge that was generated for the WAF request | [optional] 
**RequestType** | Pointer to [**RequestDetailsRequestType**](RequestDetailsRequestType.md) |  | [optional] [default to "API"]
**SessionRequestCount** | Pointer to **string** | The number of WAF requests in the session | [optional] 

## Methods

### NewWafRequestDetails

`func NewWafRequestDetails() *WafRequestDetails`

NewWafRequestDetails instantiates a new WafRequestDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafRequestDetailsWithDefaults

`func NewWafRequestDetailsWithDefaults() *WafRequestDetails`

NewWafRequestDetailsWithDefaults instantiates a new WafRequestDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WafRequestDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WafRequestDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WafRequestDetails) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WafRequestDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPath

`func (o *WafRequestDetails) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *WafRequestDetails) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *WafRequestDetails) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *WafRequestDetails) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetMethod

`func (o *WafRequestDetails) GetMethod() WafHttpMethod`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *WafRequestDetails) GetMethodOk() (*WafHttpMethod, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *WafRequestDetails) SetMethod(v WafHttpMethod)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *WafRequestDetails) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetAction

`func (o *WafRequestDetails) GetAction() WafRequestAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *WafRequestDetails) GetActionOk() (*WafRequestAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *WafRequestDetails) SetAction(v WafRequestAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *WafRequestDetails) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetRuleId

`func (o *WafRequestDetails) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *WafRequestDetails) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *WafRequestDetails) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *WafRequestDetails) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetRuleName

`func (o *WafRequestDetails) GetRuleName() string`

GetRuleName returns the RuleName field if non-nil, zero value otherwise.

### GetRuleNameOk

`func (o *WafRequestDetails) GetRuleNameOk() (*string, bool)`

GetRuleNameOk returns a tuple with the RuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleName

`func (o *WafRequestDetails) SetRuleName(v string)`

SetRuleName sets RuleName field to given value.

### HasRuleName

`func (o *WafRequestDetails) HasRuleName() bool`

HasRuleName returns a boolean if a field has been set.

### GetUserAgent

`func (o *WafRequestDetails) GetUserAgent() WafRequestDetailsUserAgent`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *WafRequestDetails) GetUserAgentOk() (*WafRequestDetailsUserAgent, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *WafRequestDetails) SetUserAgent(v WafRequestDetailsUserAgent)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *WafRequestDetails) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetNetwork

`func (o *WafRequestDetails) GetNetwork() WafRequestDetailsNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *WafRequestDetails) GetNetworkOk() (*WafRequestDetailsNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *WafRequestDetails) SetNetwork(v WafRequestDetailsNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *WafRequestDetails) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetRequestTime

`func (o *WafRequestDetails) GetRequestTime() time.Time`

GetRequestTime returns the RequestTime field if non-nil, zero value otherwise.

### GetRequestTimeOk

`func (o *WafRequestDetails) GetRequestTimeOk() (*time.Time, bool)`

GetRequestTimeOk returns a tuple with the RequestTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTime

`func (o *WafRequestDetails) SetRequestTime(v time.Time)`

SetRequestTime sets RequestTime field to given value.

### HasRequestTime

`func (o *WafRequestDetails) HasRequestTime() bool`

HasRequestTime returns a boolean if a field has been set.

### GetReferenceId

`func (o *WafRequestDetails) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *WafRequestDetails) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *WafRequestDetails) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *WafRequestDetails) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetContentType

`func (o *WafRequestDetails) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *WafRequestDetails) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *WafRequestDetails) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *WafRequestDetails) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetScheme

`func (o *WafRequestDetails) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *WafRequestDetails) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *WafRequestDetails) SetScheme(v string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *WafRequestDetails) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### GetHttpStatusCode

`func (o *WafRequestDetails) GetHttpStatusCode() string`

GetHttpStatusCode returns the HttpStatusCode field if non-nil, zero value otherwise.

### GetHttpStatusCodeOk

`func (o *WafRequestDetails) GetHttpStatusCodeOk() (*string, bool)`

GetHttpStatusCodeOk returns a tuple with the HttpStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpStatusCode

`func (o *WafRequestDetails) SetHttpStatusCode(v string)`

SetHttpStatusCode sets HttpStatusCode field to given value.

### HasHttpStatusCode

`func (o *WafRequestDetails) HasHttpStatusCode() bool`

HasHttpStatusCode returns a boolean if a field has been set.

### GetHttpVersion

`func (o *WafRequestDetails) GetHttpVersion() string`

GetHttpVersion returns the HttpVersion field if non-nil, zero value otherwise.

### GetHttpVersionOk

`func (o *WafRequestDetails) GetHttpVersionOk() (*string, bool)`

GetHttpVersionOk returns a tuple with the HttpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpVersion

`func (o *WafRequestDetails) SetHttpVersion(v string)`

SetHttpVersion sets HttpVersion field to given value.

### HasHttpVersion

`func (o *WafRequestDetails) HasHttpVersion() bool`

HasHttpVersion returns a boolean if a field has been set.

### GetTags

`func (o *WafRequestDetails) GetTags() map[string]RequestDetailsTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WafRequestDetails) GetTagsOk() (*map[string]RequestDetailsTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WafRequestDetails) SetTags(v map[string]RequestDetailsTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WafRequestDetails) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetResponseTime

`func (o *WafRequestDetails) GetResponseTime() string`

GetResponseTime returns the ResponseTime field if non-nil, zero value otherwise.

### GetResponseTimeOk

`func (o *WafRequestDetails) GetResponseTimeOk() (*string, bool)`

GetResponseTimeOk returns a tuple with the ResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTime

`func (o *WafRequestDetails) SetResponseTime(v string)`

SetResponseTime sets ResponseTime field to given value.

### HasResponseTime

`func (o *WafRequestDetails) HasResponseTime() bool`

HasResponseTime returns a boolean if a field has been set.

### GetRequestHeaders

`func (o *WafRequestDetails) GetRequestHeaders() map[string]string`

GetRequestHeaders returns the RequestHeaders field if non-nil, zero value otherwise.

### GetRequestHeadersOk

`func (o *WafRequestDetails) GetRequestHeadersOk() (*map[string]string, bool)`

GetRequestHeadersOk returns a tuple with the RequestHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestHeaders

`func (o *WafRequestDetails) SetRequestHeaders(v map[string]string)`

SetRequestHeaders sets RequestHeaders field to given value.

### HasRequestHeaders

`func (o *WafRequestDetails) HasRequestHeaders() bool`

HasRequestHeaders returns a boolean if a field has been set.

### GetIncidentId

`func (o *WafRequestDetails) GetIncidentId() string`

GetIncidentId returns the IncidentId field if non-nil, zero value otherwise.

### GetIncidentIdOk

`func (o *WafRequestDetails) GetIncidentIdOk() (*string, bool)`

GetIncidentIdOk returns a tuple with the IncidentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentId

`func (o *WafRequestDetails) SetIncidentId(v string)`

SetIncidentId sets IncidentId field to given value.

### HasIncidentId

`func (o *WafRequestDetails) HasIncidentId() bool`

HasIncidentId returns a boolean if a field has been set.

### GetRequestType

`func (o *WafRequestDetails) GetRequestType() RequestDetailsRequestType`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *WafRequestDetails) GetRequestTypeOk() (*RequestDetailsRequestType, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *WafRequestDetails) SetRequestType(v RequestDetailsRequestType)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *WafRequestDetails) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetSessionRequestCount

`func (o *WafRequestDetails) GetSessionRequestCount() string`

GetSessionRequestCount returns the SessionRequestCount field if non-nil, zero value otherwise.

### GetSessionRequestCountOk

`func (o *WafRequestDetails) GetSessionRequestCountOk() (*string, bool)`

GetSessionRequestCountOk returns a tuple with the SessionRequestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRequestCount

`func (o *WafRequestDetails) SetSessionRequestCount(v string)`

SetSessionRequestCount sets SessionRequestCount field to given value.

### HasSessionRequestCount

`func (o *WafRequestDetails) HasSessionRequestCount() bool`

HasSessionRequestCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


