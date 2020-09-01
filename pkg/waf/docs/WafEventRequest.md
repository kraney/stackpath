# WafEventRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** | The requested domain name | [optional] 
**Method** | Pointer to **string** | The HTTP method that triggered a WAF event | [optional] 
**Scheme** | Pointer to **string** | The URL scheme that triggered a WAF event | [optional] 
**Uri** | Pointer to **string** | The full URL that triggered a WAF event | [optional] 
**QueryString** | Pointer to **string** | The query string portion of a URL that triggered a WAF event | [optional] 
**Headers** | Pointer to **map[string]string** | A key/value pair of the event&#39;s request headers | [optional] 
**UserAgent** | Pointer to [**EventRequestUserAgent**](EventRequestUserAgent.md) |  | [optional] 

## Methods

### NewWafEventRequest

`func NewWafEventRequest() *WafEventRequest`

NewWafEventRequest instantiates a new WafEventRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafEventRequestWithDefaults

`func NewWafEventRequestWithDefaults() *WafEventRequest`

NewWafEventRequestWithDefaults instantiates a new WafEventRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *WafEventRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *WafEventRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *WafEventRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *WafEventRequest) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetMethod

`func (o *WafEventRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *WafEventRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *WafEventRequest) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *WafEventRequest) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetScheme

`func (o *WafEventRequest) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *WafEventRequest) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *WafEventRequest) SetScheme(v string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *WafEventRequest) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### GetUri

`func (o *WafEventRequest) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *WafEventRequest) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *WafEventRequest) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *WafEventRequest) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetQueryString

`func (o *WafEventRequest) GetQueryString() string`

GetQueryString returns the QueryString field if non-nil, zero value otherwise.

### GetQueryStringOk

`func (o *WafEventRequest) GetQueryStringOk() (*string, bool)`

GetQueryStringOk returns a tuple with the QueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryString

`func (o *WafEventRequest) SetQueryString(v string)`

SetQueryString sets QueryString field to given value.

### HasQueryString

`func (o *WafEventRequest) HasQueryString() bool`

HasQueryString returns a boolean if a field has been set.

### GetHeaders

`func (o *WafEventRequest) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *WafEventRequest) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *WafEventRequest) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *WafEventRequest) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetUserAgent

`func (o *WafEventRequest) GetUserAgent() EventRequestUserAgent`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *WafEventRequest) GetUserAgentOk() (*EventRequestUserAgent, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *WafEventRequest) SetUserAgent(v EventRequestUserAgent)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *WafEventRequest) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


