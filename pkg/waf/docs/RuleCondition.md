# RuleCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to [**ConditionIpCondition**](ConditionIpCondition.md) |  | [optional] 
**IpRange** | Pointer to [**ConditionIpRangeCondition**](ConditionIpRangeCondition.md) |  | [optional] 
**Url** | Pointer to [**ConditionUrlCondition**](ConditionUrlCondition.md) |  | [optional] 
**UserAgent** | Pointer to [**ConditionUserAgentCondition**](ConditionUserAgentCondition.md) |  | [optional] 
**Header** | Pointer to [**ConditionHeaderCondition**](ConditionHeaderCondition.md) |  | [optional] 
**HeaderExists** | Pointer to [**ConditionHeaderExistsCondition**](ConditionHeaderExistsCondition.md) |  | [optional] 
**ResponseHeader** | Pointer to [**ConditionResponseHeaderCondition**](ConditionResponseHeaderCondition.md) |  | [optional] 
**ResponseHeaderExists** | Pointer to [**ConditionResponseHeaderExistsCondition**](ConditionResponseHeaderExistsCondition.md) |  | [optional] 
**HttpMethod** | Pointer to [**ConditionHttpMethodCondition**](ConditionHttpMethodCondition.md) |  | [optional] 
**FileExtension** | Pointer to [**ConditionFileExtensionCondition**](ConditionFileExtensionCondition.md) |  | [optional] 
**ContentType** | Pointer to [**ConditionContentTypeCondition**](ConditionContentTypeCondition.md) |  | [optional] 
**Country** | Pointer to [**ConditionCountryCondition**](ConditionCountryCondition.md) |  | [optional] 
**Organization** | Pointer to [**ConditionOrganizationCondition**](ConditionOrganizationCondition.md) |  | [optional] 
**RequestRate** | Pointer to [**ConditionRequestRateCondition**](ConditionRequestRateCondition.md) |  | [optional] 
**OwnerTypes** | Pointer to [**ConditionOwnerTypeCondition**](ConditionOwnerTypeCondition.md) |  | [optional] 
**Tags** | Pointer to [**ConditionTagCondition**](ConditionTagCondition.md) |  | [optional] 
**SessionRequestCount** | Pointer to [**ConditionSessionRequestCountCondition**](ConditionSessionRequestCountCondition.md) |  | [optional] 
**Negation** | Pointer to **bool** | Whether or not to apply a boolean NOT operation to the rule&#39;s condition | [optional] 

## Methods

### NewRuleCondition

`func NewRuleCondition() *RuleCondition`

NewRuleCondition instantiates a new RuleCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleConditionWithDefaults

`func NewRuleConditionWithDefaults() *RuleCondition`

NewRuleConditionWithDefaults instantiates a new RuleCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *RuleCondition) GetIp() ConditionIpCondition`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *RuleCondition) GetIpOk() (*ConditionIpCondition, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *RuleCondition) SetIp(v ConditionIpCondition)`

SetIp sets Ip field to given value.

### HasIp

`func (o *RuleCondition) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpRange

`func (o *RuleCondition) GetIpRange() ConditionIpRangeCondition`

GetIpRange returns the IpRange field if non-nil, zero value otherwise.

### GetIpRangeOk

`func (o *RuleCondition) GetIpRangeOk() (*ConditionIpRangeCondition, bool)`

GetIpRangeOk returns a tuple with the IpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRange

`func (o *RuleCondition) SetIpRange(v ConditionIpRangeCondition)`

SetIpRange sets IpRange field to given value.

### HasIpRange

`func (o *RuleCondition) HasIpRange() bool`

HasIpRange returns a boolean if a field has been set.

### GetUrl

`func (o *RuleCondition) GetUrl() ConditionUrlCondition`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RuleCondition) GetUrlOk() (*ConditionUrlCondition, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RuleCondition) SetUrl(v ConditionUrlCondition)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *RuleCondition) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUserAgent

`func (o *RuleCondition) GetUserAgent() ConditionUserAgentCondition`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *RuleCondition) GetUserAgentOk() (*ConditionUserAgentCondition, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *RuleCondition) SetUserAgent(v ConditionUserAgentCondition)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *RuleCondition) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetHeader

`func (o *RuleCondition) GetHeader() ConditionHeaderCondition`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *RuleCondition) GetHeaderOk() (*ConditionHeaderCondition, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *RuleCondition) SetHeader(v ConditionHeaderCondition)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *RuleCondition) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetHeaderExists

`func (o *RuleCondition) GetHeaderExists() ConditionHeaderExistsCondition`

GetHeaderExists returns the HeaderExists field if non-nil, zero value otherwise.

### GetHeaderExistsOk

`func (o *RuleCondition) GetHeaderExistsOk() (*ConditionHeaderExistsCondition, bool)`

GetHeaderExistsOk returns a tuple with the HeaderExists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderExists

`func (o *RuleCondition) SetHeaderExists(v ConditionHeaderExistsCondition)`

SetHeaderExists sets HeaderExists field to given value.

### HasHeaderExists

`func (o *RuleCondition) HasHeaderExists() bool`

HasHeaderExists returns a boolean if a field has been set.

### GetResponseHeader

`func (o *RuleCondition) GetResponseHeader() ConditionResponseHeaderCondition`

GetResponseHeader returns the ResponseHeader field if non-nil, zero value otherwise.

### GetResponseHeaderOk

`func (o *RuleCondition) GetResponseHeaderOk() (*ConditionResponseHeaderCondition, bool)`

GetResponseHeaderOk returns a tuple with the ResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeader

`func (o *RuleCondition) SetResponseHeader(v ConditionResponseHeaderCondition)`

SetResponseHeader sets ResponseHeader field to given value.

### HasResponseHeader

`func (o *RuleCondition) HasResponseHeader() bool`

HasResponseHeader returns a boolean if a field has been set.

### GetResponseHeaderExists

`func (o *RuleCondition) GetResponseHeaderExists() ConditionResponseHeaderExistsCondition`

GetResponseHeaderExists returns the ResponseHeaderExists field if non-nil, zero value otherwise.

### GetResponseHeaderExistsOk

`func (o *RuleCondition) GetResponseHeaderExistsOk() (*ConditionResponseHeaderExistsCondition, bool)`

GetResponseHeaderExistsOk returns a tuple with the ResponseHeaderExists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeaderExists

`func (o *RuleCondition) SetResponseHeaderExists(v ConditionResponseHeaderExistsCondition)`

SetResponseHeaderExists sets ResponseHeaderExists field to given value.

### HasResponseHeaderExists

`func (o *RuleCondition) HasResponseHeaderExists() bool`

HasResponseHeaderExists returns a boolean if a field has been set.

### GetHttpMethod

`func (o *RuleCondition) GetHttpMethod() ConditionHttpMethodCondition`

GetHttpMethod returns the HttpMethod field if non-nil, zero value otherwise.

### GetHttpMethodOk

`func (o *RuleCondition) GetHttpMethodOk() (*ConditionHttpMethodCondition, bool)`

GetHttpMethodOk returns a tuple with the HttpMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpMethod

`func (o *RuleCondition) SetHttpMethod(v ConditionHttpMethodCondition)`

SetHttpMethod sets HttpMethod field to given value.

### HasHttpMethod

`func (o *RuleCondition) HasHttpMethod() bool`

HasHttpMethod returns a boolean if a field has been set.

### GetFileExtension

`func (o *RuleCondition) GetFileExtension() ConditionFileExtensionCondition`

GetFileExtension returns the FileExtension field if non-nil, zero value otherwise.

### GetFileExtensionOk

`func (o *RuleCondition) GetFileExtensionOk() (*ConditionFileExtensionCondition, bool)`

GetFileExtensionOk returns a tuple with the FileExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileExtension

`func (o *RuleCondition) SetFileExtension(v ConditionFileExtensionCondition)`

SetFileExtension sets FileExtension field to given value.

### HasFileExtension

`func (o *RuleCondition) HasFileExtension() bool`

HasFileExtension returns a boolean if a field has been set.

### GetContentType

`func (o *RuleCondition) GetContentType() ConditionContentTypeCondition`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *RuleCondition) GetContentTypeOk() (*ConditionContentTypeCondition, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *RuleCondition) SetContentType(v ConditionContentTypeCondition)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *RuleCondition) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetCountry

`func (o *RuleCondition) GetCountry() ConditionCountryCondition`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *RuleCondition) GetCountryOk() (*ConditionCountryCondition, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *RuleCondition) SetCountry(v ConditionCountryCondition)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *RuleCondition) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetOrganization

`func (o *RuleCondition) GetOrganization() ConditionOrganizationCondition`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *RuleCondition) GetOrganizationOk() (*ConditionOrganizationCondition, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *RuleCondition) SetOrganization(v ConditionOrganizationCondition)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *RuleCondition) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetRequestRate

`func (o *RuleCondition) GetRequestRate() ConditionRequestRateCondition`

GetRequestRate returns the RequestRate field if non-nil, zero value otherwise.

### GetRequestRateOk

`func (o *RuleCondition) GetRequestRateOk() (*ConditionRequestRateCondition, bool)`

GetRequestRateOk returns a tuple with the RequestRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestRate

`func (o *RuleCondition) SetRequestRate(v ConditionRequestRateCondition)`

SetRequestRate sets RequestRate field to given value.

### HasRequestRate

`func (o *RuleCondition) HasRequestRate() bool`

HasRequestRate returns a boolean if a field has been set.

### GetOwnerTypes

`func (o *RuleCondition) GetOwnerTypes() ConditionOwnerTypeCondition`

GetOwnerTypes returns the OwnerTypes field if non-nil, zero value otherwise.

### GetOwnerTypesOk

`func (o *RuleCondition) GetOwnerTypesOk() (*ConditionOwnerTypeCondition, bool)`

GetOwnerTypesOk returns a tuple with the OwnerTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerTypes

`func (o *RuleCondition) SetOwnerTypes(v ConditionOwnerTypeCondition)`

SetOwnerTypes sets OwnerTypes field to given value.

### HasOwnerTypes

`func (o *RuleCondition) HasOwnerTypes() bool`

HasOwnerTypes returns a boolean if a field has been set.

### GetTags

`func (o *RuleCondition) GetTags() ConditionTagCondition`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RuleCondition) GetTagsOk() (*ConditionTagCondition, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RuleCondition) SetTags(v ConditionTagCondition)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RuleCondition) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetSessionRequestCount

`func (o *RuleCondition) GetSessionRequestCount() ConditionSessionRequestCountCondition`

GetSessionRequestCount returns the SessionRequestCount field if non-nil, zero value otherwise.

### GetSessionRequestCountOk

`func (o *RuleCondition) GetSessionRequestCountOk() (*ConditionSessionRequestCountCondition, bool)`

GetSessionRequestCountOk returns a tuple with the SessionRequestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRequestCount

`func (o *RuleCondition) SetSessionRequestCount(v ConditionSessionRequestCountCondition)`

SetSessionRequestCount sets SessionRequestCount field to given value.

### HasSessionRequestCount

`func (o *RuleCondition) HasSessionRequestCount() bool`

HasSessionRequestCount returns a boolean if a field has been set.

### GetNegation

`func (o *RuleCondition) GetNegation() bool`

GetNegation returns the Negation field if non-nil, zero value otherwise.

### GetNegationOk

`func (o *RuleCondition) GetNegationOk() (*bool, bool)`

GetNegationOk returns a tuple with the Negation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegation

`func (o *RuleCondition) SetNegation(v bool)`

SetNegation sets Negation field to given value.

### HasNegation

`func (o *RuleCondition) HasNegation() bool`

HasNegation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


