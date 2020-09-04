# SchemawafRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A request&#39;s unique identifier | [optional] 
**Path** | Pointer to **string** | The path of the requested URL | [optional] 
**ClientIp** | Pointer to **string** | The originating IP address | [optional] 
**Method** | Pointer to [**WafHttpMethod**](wafHttpMethod.md) |  | [optional] [default to "METHOD_UNSPECIFIED"]
**RuleName** | Pointer to **string** | The name of the WAF rule triggered by the request | [optional] 
**Country** | Pointer to **string** | The ISO 3166-1 alpha-2 code of the country where the request originated from | [optional] 
**Action** | Pointer to [**WafRequestAction**](wafRequestAction.md) |  | [optional] [default to "ACTION_UNSPECIFIED"]
**RuleId** | Pointer to **string** | The unique ID of the WAF rule that triggered the action against the request | [optional] 
**UserAgent** | Pointer to **string** | The full user agent string for the request | [optional] 
**UserAgentClient** | Pointer to **string** | The name of the requesting client, typically the name of the requesting web browser | [optional] 
**Organization** | Pointer to **string** | The organization that owns the request&#39;s originating IP address according to a WHOIS lookup | [optional] 
**RequestTime** | Pointer to [**time.Time**](time.Time.md) | When the request occurred | [optional] 
**ReferenceId** | Pointer to **string** | The request&#39;s user-facing identifier  Reference IDs are displayed to the end user when the WAF blocks a request to a site. Please note that a request&#39;s ID and reference ID are different. | [optional] 

## Methods

### NewSchemawafRequest

`func NewSchemawafRequest() *SchemawafRequest`

NewSchemawafRequest instantiates a new SchemawafRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemawafRequestWithDefaults

`func NewSchemawafRequestWithDefaults() *SchemawafRequest`

NewSchemawafRequestWithDefaults instantiates a new SchemawafRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SchemawafRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SchemawafRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SchemawafRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SchemawafRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPath

`func (o *SchemawafRequest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SchemawafRequest) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SchemawafRequest) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *SchemawafRequest) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetClientIp

`func (o *SchemawafRequest) GetClientIp() string`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *SchemawafRequest) GetClientIpOk() (*string, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *SchemawafRequest) SetClientIp(v string)`

SetClientIp sets ClientIp field to given value.

### HasClientIp

`func (o *SchemawafRequest) HasClientIp() bool`

HasClientIp returns a boolean if a field has been set.

### GetMethod

`func (o *SchemawafRequest) GetMethod() WafHttpMethod`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *SchemawafRequest) GetMethodOk() (*WafHttpMethod, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *SchemawafRequest) SetMethod(v WafHttpMethod)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *SchemawafRequest) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetRuleName

`func (o *SchemawafRequest) GetRuleName() string`

GetRuleName returns the RuleName field if non-nil, zero value otherwise.

### GetRuleNameOk

`func (o *SchemawafRequest) GetRuleNameOk() (*string, bool)`

GetRuleNameOk returns a tuple with the RuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleName

`func (o *SchemawafRequest) SetRuleName(v string)`

SetRuleName sets RuleName field to given value.

### HasRuleName

`func (o *SchemawafRequest) HasRuleName() bool`

HasRuleName returns a boolean if a field has been set.

### GetCountry

`func (o *SchemawafRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *SchemawafRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *SchemawafRequest) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *SchemawafRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetAction

`func (o *SchemawafRequest) GetAction() WafRequestAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SchemawafRequest) GetActionOk() (*WafRequestAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SchemawafRequest) SetAction(v WafRequestAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *SchemawafRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetRuleId

`func (o *SchemawafRequest) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *SchemawafRequest) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *SchemawafRequest) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *SchemawafRequest) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetUserAgent

`func (o *SchemawafRequest) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *SchemawafRequest) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *SchemawafRequest) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *SchemawafRequest) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetUserAgentClient

`func (o *SchemawafRequest) GetUserAgentClient() string`

GetUserAgentClient returns the UserAgentClient field if non-nil, zero value otherwise.

### GetUserAgentClientOk

`func (o *SchemawafRequest) GetUserAgentClientOk() (*string, bool)`

GetUserAgentClientOk returns a tuple with the UserAgentClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgentClient

`func (o *SchemawafRequest) SetUserAgentClient(v string)`

SetUserAgentClient sets UserAgentClient field to given value.

### HasUserAgentClient

`func (o *SchemawafRequest) HasUserAgentClient() bool`

HasUserAgentClient returns a boolean if a field has been set.

### GetOrganization

`func (o *SchemawafRequest) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *SchemawafRequest) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *SchemawafRequest) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *SchemawafRequest) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetRequestTime

`func (o *SchemawafRequest) GetRequestTime() time.Time`

GetRequestTime returns the RequestTime field if non-nil, zero value otherwise.

### GetRequestTimeOk

`func (o *SchemawafRequest) GetRequestTimeOk() (*time.Time, bool)`

GetRequestTimeOk returns a tuple with the RequestTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTime

`func (o *SchemawafRequest) SetRequestTime(v time.Time)`

SetRequestTime sets RequestTime field to given value.

### HasRequestTime

`func (o *SchemawafRequest) HasRequestTime() bool`

HasRequestTime returns a boolean if a field has been set.

### GetReferenceId

`func (o *SchemawafRequest) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *SchemawafRequest) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *SchemawafRequest) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *SchemawafRequest) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


