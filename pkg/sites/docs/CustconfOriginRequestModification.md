# CustconfOriginRequestModification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | This is used by the API to perform conflict checking | [optional] 
**UrlPattern** | Pointer to **string** |  | [optional] 
**UrlRewrite** | Pointer to **string** |  | [optional] 
**HeaderPattern** | Pointer to **string** |  | [optional] 
**HeaderRewrite** | Pointer to **string** |  | [optional] 
**AddHeaders** | Pointer to **string** | String of values delimited by a &#39;|&#39; character. This is the static HTTP header you want inserted into the CDN request. Start value with \&quot;append:\&quot;, \&quot;replace:\&quot; or \&quot;create:\&quot; which defines if Header will be added, replaced or create if not exists. Default is append. | [optional] 
**FlowControl** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**MethodFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**PathFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**HeaderFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**CookieFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 

## Methods

### NewCustconfOriginRequestModification

`func NewCustconfOriginRequestModification() *CustconfOriginRequestModification`

NewCustconfOriginRequestModification instantiates a new CustconfOriginRequestModification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustconfOriginRequestModificationWithDefaults

`func NewCustconfOriginRequestModificationWithDefaults() *CustconfOriginRequestModification`

NewCustconfOriginRequestModificationWithDefaults instantiates a new CustconfOriginRequestModification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustconfOriginRequestModification) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustconfOriginRequestModification) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustconfOriginRequestModification) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustconfOriginRequestModification) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrlPattern

`func (o *CustconfOriginRequestModification) GetUrlPattern() string`

GetUrlPattern returns the UrlPattern field if non-nil, zero value otherwise.

### GetUrlPatternOk

`func (o *CustconfOriginRequestModification) GetUrlPatternOk() (*string, bool)`

GetUrlPatternOk returns a tuple with the UrlPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlPattern

`func (o *CustconfOriginRequestModification) SetUrlPattern(v string)`

SetUrlPattern sets UrlPattern field to given value.

### HasUrlPattern

`func (o *CustconfOriginRequestModification) HasUrlPattern() bool`

HasUrlPattern returns a boolean if a field has been set.

### GetUrlRewrite

`func (o *CustconfOriginRequestModification) GetUrlRewrite() string`

GetUrlRewrite returns the UrlRewrite field if non-nil, zero value otherwise.

### GetUrlRewriteOk

`func (o *CustconfOriginRequestModification) GetUrlRewriteOk() (*string, bool)`

GetUrlRewriteOk returns a tuple with the UrlRewrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlRewrite

`func (o *CustconfOriginRequestModification) SetUrlRewrite(v string)`

SetUrlRewrite sets UrlRewrite field to given value.

### HasUrlRewrite

`func (o *CustconfOriginRequestModification) HasUrlRewrite() bool`

HasUrlRewrite returns a boolean if a field has been set.

### GetHeaderPattern

`func (o *CustconfOriginRequestModification) GetHeaderPattern() string`

GetHeaderPattern returns the HeaderPattern field if non-nil, zero value otherwise.

### GetHeaderPatternOk

`func (o *CustconfOriginRequestModification) GetHeaderPatternOk() (*string, bool)`

GetHeaderPatternOk returns a tuple with the HeaderPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderPattern

`func (o *CustconfOriginRequestModification) SetHeaderPattern(v string)`

SetHeaderPattern sets HeaderPattern field to given value.

### HasHeaderPattern

`func (o *CustconfOriginRequestModification) HasHeaderPattern() bool`

HasHeaderPattern returns a boolean if a field has been set.

### GetHeaderRewrite

`func (o *CustconfOriginRequestModification) GetHeaderRewrite() string`

GetHeaderRewrite returns the HeaderRewrite field if non-nil, zero value otherwise.

### GetHeaderRewriteOk

`func (o *CustconfOriginRequestModification) GetHeaderRewriteOk() (*string, bool)`

GetHeaderRewriteOk returns a tuple with the HeaderRewrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderRewrite

`func (o *CustconfOriginRequestModification) SetHeaderRewrite(v string)`

SetHeaderRewrite sets HeaderRewrite field to given value.

### HasHeaderRewrite

`func (o *CustconfOriginRequestModification) HasHeaderRewrite() bool`

HasHeaderRewrite returns a boolean if a field has been set.

### GetAddHeaders

`func (o *CustconfOriginRequestModification) GetAddHeaders() string`

GetAddHeaders returns the AddHeaders field if non-nil, zero value otherwise.

### GetAddHeadersOk

`func (o *CustconfOriginRequestModification) GetAddHeadersOk() (*string, bool)`

GetAddHeadersOk returns a tuple with the AddHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddHeaders

`func (o *CustconfOriginRequestModification) SetAddHeaders(v string)`

SetAddHeaders sets AddHeaders field to given value.

### HasAddHeaders

`func (o *CustconfOriginRequestModification) HasAddHeaders() bool`

HasAddHeaders returns a boolean if a field has been set.

### GetFlowControl

`func (o *CustconfOriginRequestModification) GetFlowControl() string`

GetFlowControl returns the FlowControl field if non-nil, zero value otherwise.

### GetFlowControlOk

`func (o *CustconfOriginRequestModification) GetFlowControlOk() (*string, bool)`

GetFlowControlOk returns a tuple with the FlowControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowControl

`func (o *CustconfOriginRequestModification) SetFlowControl(v string)`

SetFlowControl sets FlowControl field to given value.

### HasFlowControl

`func (o *CustconfOriginRequestModification) HasFlowControl() bool`

HasFlowControl returns a boolean if a field has been set.

### GetEnabled

`func (o *CustconfOriginRequestModification) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustconfOriginRequestModification) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustconfOriginRequestModification) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CustconfOriginRequestModification) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMethodFilter

`func (o *CustconfOriginRequestModification) GetMethodFilter() string`

GetMethodFilter returns the MethodFilter field if non-nil, zero value otherwise.

### GetMethodFilterOk

`func (o *CustconfOriginRequestModification) GetMethodFilterOk() (*string, bool)`

GetMethodFilterOk returns a tuple with the MethodFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodFilter

`func (o *CustconfOriginRequestModification) SetMethodFilter(v string)`

SetMethodFilter sets MethodFilter field to given value.

### HasMethodFilter

`func (o *CustconfOriginRequestModification) HasMethodFilter() bool`

HasMethodFilter returns a boolean if a field has been set.

### GetPathFilter

`func (o *CustconfOriginRequestModification) GetPathFilter() string`

GetPathFilter returns the PathFilter field if non-nil, zero value otherwise.

### GetPathFilterOk

`func (o *CustconfOriginRequestModification) GetPathFilterOk() (*string, bool)`

GetPathFilterOk returns a tuple with the PathFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathFilter

`func (o *CustconfOriginRequestModification) SetPathFilter(v string)`

SetPathFilter sets PathFilter field to given value.

### HasPathFilter

`func (o *CustconfOriginRequestModification) HasPathFilter() bool`

HasPathFilter returns a boolean if a field has been set.

### GetHeaderFilter

`func (o *CustconfOriginRequestModification) GetHeaderFilter() string`

GetHeaderFilter returns the HeaderFilter field if non-nil, zero value otherwise.

### GetHeaderFilterOk

`func (o *CustconfOriginRequestModification) GetHeaderFilterOk() (*string, bool)`

GetHeaderFilterOk returns a tuple with the HeaderFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderFilter

`func (o *CustconfOriginRequestModification) SetHeaderFilter(v string)`

SetHeaderFilter sets HeaderFilter field to given value.

### HasHeaderFilter

`func (o *CustconfOriginRequestModification) HasHeaderFilter() bool`

HasHeaderFilter returns a boolean if a field has been set.

### GetCookieFilter

`func (o *CustconfOriginRequestModification) GetCookieFilter() string`

GetCookieFilter returns the CookieFilter field if non-nil, zero value otherwise.

### GetCookieFilterOk

`func (o *CustconfOriginRequestModification) GetCookieFilterOk() (*string, bool)`

GetCookieFilterOk returns a tuple with the CookieFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookieFilter

`func (o *CustconfOriginRequestModification) SetCookieFilter(v string)`

SetCookieFilter sets CookieFilter field to given value.

### HasCookieFilter

`func (o *CustconfOriginRequestModification) HasCookieFilter() bool`

HasCookieFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


