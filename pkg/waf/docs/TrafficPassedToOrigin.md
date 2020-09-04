# TrafficPassedToOrigin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LegitimateRequest** | Pointer to **string** | The number of requests that passed all policies and rules and were sent to the origin | [optional] 
**Whitelisted** | Pointer to **string** | The number of requests allowed by whitelist | [optional] 
**AllowedSearchEngine** | Pointer to **string** | The number of requests allowed by allow search bot policies | [optional] 
**StaticContent** | Pointer to **string** | The number of requests for static content that was not analyzed by the WAF | [optional] 
**Uncategorized** | Pointer to **string** | All uncategorized requests that the WAF passed to the origin | [optional] 

## Methods

### NewTrafficPassedToOrigin

`func NewTrafficPassedToOrigin() *TrafficPassedToOrigin`

NewTrafficPassedToOrigin instantiates a new TrafficPassedToOrigin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrafficPassedToOriginWithDefaults

`func NewTrafficPassedToOriginWithDefaults() *TrafficPassedToOrigin`

NewTrafficPassedToOriginWithDefaults instantiates a new TrafficPassedToOrigin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLegitimateRequest

`func (o *TrafficPassedToOrigin) GetLegitimateRequest() string`

GetLegitimateRequest returns the LegitimateRequest field if non-nil, zero value otherwise.

### GetLegitimateRequestOk

`func (o *TrafficPassedToOrigin) GetLegitimateRequestOk() (*string, bool)`

GetLegitimateRequestOk returns a tuple with the LegitimateRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegitimateRequest

`func (o *TrafficPassedToOrigin) SetLegitimateRequest(v string)`

SetLegitimateRequest sets LegitimateRequest field to given value.

### HasLegitimateRequest

`func (o *TrafficPassedToOrigin) HasLegitimateRequest() bool`

HasLegitimateRequest returns a boolean if a field has been set.

### GetWhitelisted

`func (o *TrafficPassedToOrigin) GetWhitelisted() string`

GetWhitelisted returns the Whitelisted field if non-nil, zero value otherwise.

### GetWhitelistedOk

`func (o *TrafficPassedToOrigin) GetWhitelistedOk() (*string, bool)`

GetWhitelistedOk returns a tuple with the Whitelisted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelisted

`func (o *TrafficPassedToOrigin) SetWhitelisted(v string)`

SetWhitelisted sets Whitelisted field to given value.

### HasWhitelisted

`func (o *TrafficPassedToOrigin) HasWhitelisted() bool`

HasWhitelisted returns a boolean if a field has been set.

### GetAllowedSearchEngine

`func (o *TrafficPassedToOrigin) GetAllowedSearchEngine() string`

GetAllowedSearchEngine returns the AllowedSearchEngine field if non-nil, zero value otherwise.

### GetAllowedSearchEngineOk

`func (o *TrafficPassedToOrigin) GetAllowedSearchEngineOk() (*string, bool)`

GetAllowedSearchEngineOk returns a tuple with the AllowedSearchEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedSearchEngine

`func (o *TrafficPassedToOrigin) SetAllowedSearchEngine(v string)`

SetAllowedSearchEngine sets AllowedSearchEngine field to given value.

### HasAllowedSearchEngine

`func (o *TrafficPassedToOrigin) HasAllowedSearchEngine() bool`

HasAllowedSearchEngine returns a boolean if a field has been set.

### GetStaticContent

`func (o *TrafficPassedToOrigin) GetStaticContent() string`

GetStaticContent returns the StaticContent field if non-nil, zero value otherwise.

### GetStaticContentOk

`func (o *TrafficPassedToOrigin) GetStaticContentOk() (*string, bool)`

GetStaticContentOk returns a tuple with the StaticContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticContent

`func (o *TrafficPassedToOrigin) SetStaticContent(v string)`

SetStaticContent sets StaticContent field to given value.

### HasStaticContent

`func (o *TrafficPassedToOrigin) HasStaticContent() bool`

HasStaticContent returns a boolean if a field has been set.

### GetUncategorized

`func (o *TrafficPassedToOrigin) GetUncategorized() string`

GetUncategorized returns the Uncategorized field if non-nil, zero value otherwise.

### GetUncategorizedOk

`func (o *TrafficPassedToOrigin) GetUncategorizedOk() (*string, bool)`

GetUncategorizedOk returns a tuple with the Uncategorized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncategorized

`func (o *TrafficPassedToOrigin) SetUncategorized(v string)`

SetUncategorized sets Uncategorized field to given value.

### HasUncategorized

`func (o *TrafficPassedToOrigin) HasUncategorized() bool`

HasUncategorized returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


