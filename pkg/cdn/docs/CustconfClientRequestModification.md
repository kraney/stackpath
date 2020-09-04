# CustconfClientRequestModification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | This is used by the API to perform conflict checking | [optional] 
**UrlPattern** | Pointer to **string** |  | [optional] 
**UrlRewrite** | Pointer to **string** |  | [optional] 
**HeaderPattern** | Pointer to **string** |  | [optional] 
**HeaderRewrite** | Pointer to **string** |  | [optional] 
**AddHeaders** | Pointer to **string** | String of values delimited by a &#39;|&#39; character. This is the static HTTP header you want inserted into the CDN request. Start value with \&quot;append:\&quot;, \&quot;replace:\&quot; or \&quot;create:\&quot; which defines if Header will be added, replaced or create if not exists. Default is append. | [optional] 
**FlowControl** | Pointer to [**CustconfClientRequestModificationFlowControlEnumWrapperValue**](custconfClientRequestModificationFlowControlEnumWrapperValue.md) |  | [optional] [default to "UNKNOWN"]
**Enabled** | Pointer to **bool** |  | [optional] 
**MethodFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**PathFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**HeaderFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**CookieFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 

## Methods

### NewCustconfClientRequestModification

`func NewCustconfClientRequestModification() *CustconfClientRequestModification`

NewCustconfClientRequestModification instantiates a new CustconfClientRequestModification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustconfClientRequestModificationWithDefaults

`func NewCustconfClientRequestModificationWithDefaults() *CustconfClientRequestModification`

NewCustconfClientRequestModificationWithDefaults instantiates a new CustconfClientRequestModification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustconfClientRequestModification) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustconfClientRequestModification) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustconfClientRequestModification) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustconfClientRequestModification) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrlPattern

`func (o *CustconfClientRequestModification) GetUrlPattern() string`

GetUrlPattern returns the UrlPattern field if non-nil, zero value otherwise.

### GetUrlPatternOk

`func (o *CustconfClientRequestModification) GetUrlPatternOk() (*string, bool)`

GetUrlPatternOk returns a tuple with the UrlPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlPattern

`func (o *CustconfClientRequestModification) SetUrlPattern(v string)`

SetUrlPattern sets UrlPattern field to given value.

### HasUrlPattern

`func (o *CustconfClientRequestModification) HasUrlPattern() bool`

HasUrlPattern returns a boolean if a field has been set.

### GetUrlRewrite

`func (o *CustconfClientRequestModification) GetUrlRewrite() string`

GetUrlRewrite returns the UrlRewrite field if non-nil, zero value otherwise.

### GetUrlRewriteOk

`func (o *CustconfClientRequestModification) GetUrlRewriteOk() (*string, bool)`

GetUrlRewriteOk returns a tuple with the UrlRewrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlRewrite

`func (o *CustconfClientRequestModification) SetUrlRewrite(v string)`

SetUrlRewrite sets UrlRewrite field to given value.

### HasUrlRewrite

`func (o *CustconfClientRequestModification) HasUrlRewrite() bool`

HasUrlRewrite returns a boolean if a field has been set.

### GetHeaderPattern

`func (o *CustconfClientRequestModification) GetHeaderPattern() string`

GetHeaderPattern returns the HeaderPattern field if non-nil, zero value otherwise.

### GetHeaderPatternOk

`func (o *CustconfClientRequestModification) GetHeaderPatternOk() (*string, bool)`

GetHeaderPatternOk returns a tuple with the HeaderPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderPattern

`func (o *CustconfClientRequestModification) SetHeaderPattern(v string)`

SetHeaderPattern sets HeaderPattern field to given value.

### HasHeaderPattern

`func (o *CustconfClientRequestModification) HasHeaderPattern() bool`

HasHeaderPattern returns a boolean if a field has been set.

### GetHeaderRewrite

`func (o *CustconfClientRequestModification) GetHeaderRewrite() string`

GetHeaderRewrite returns the HeaderRewrite field if non-nil, zero value otherwise.

### GetHeaderRewriteOk

`func (o *CustconfClientRequestModification) GetHeaderRewriteOk() (*string, bool)`

GetHeaderRewriteOk returns a tuple with the HeaderRewrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderRewrite

`func (o *CustconfClientRequestModification) SetHeaderRewrite(v string)`

SetHeaderRewrite sets HeaderRewrite field to given value.

### HasHeaderRewrite

`func (o *CustconfClientRequestModification) HasHeaderRewrite() bool`

HasHeaderRewrite returns a boolean if a field has been set.

### GetAddHeaders

`func (o *CustconfClientRequestModification) GetAddHeaders() string`

GetAddHeaders returns the AddHeaders field if non-nil, zero value otherwise.

### GetAddHeadersOk

`func (o *CustconfClientRequestModification) GetAddHeadersOk() (*string, bool)`

GetAddHeadersOk returns a tuple with the AddHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddHeaders

`func (o *CustconfClientRequestModification) SetAddHeaders(v string)`

SetAddHeaders sets AddHeaders field to given value.

### HasAddHeaders

`func (o *CustconfClientRequestModification) HasAddHeaders() bool`

HasAddHeaders returns a boolean if a field has been set.

### GetFlowControl

`func (o *CustconfClientRequestModification) GetFlowControl() CustconfClientRequestModificationFlowControlEnumWrapperValue`

GetFlowControl returns the FlowControl field if non-nil, zero value otherwise.

### GetFlowControlOk

`func (o *CustconfClientRequestModification) GetFlowControlOk() (*CustconfClientRequestModificationFlowControlEnumWrapperValue, bool)`

GetFlowControlOk returns a tuple with the FlowControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowControl

`func (o *CustconfClientRequestModification) SetFlowControl(v CustconfClientRequestModificationFlowControlEnumWrapperValue)`

SetFlowControl sets FlowControl field to given value.

### HasFlowControl

`func (o *CustconfClientRequestModification) HasFlowControl() bool`

HasFlowControl returns a boolean if a field has been set.

### GetEnabled

`func (o *CustconfClientRequestModification) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustconfClientRequestModification) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustconfClientRequestModification) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CustconfClientRequestModification) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMethodFilter

`func (o *CustconfClientRequestModification) GetMethodFilter() string`

GetMethodFilter returns the MethodFilter field if non-nil, zero value otherwise.

### GetMethodFilterOk

`func (o *CustconfClientRequestModification) GetMethodFilterOk() (*string, bool)`

GetMethodFilterOk returns a tuple with the MethodFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodFilter

`func (o *CustconfClientRequestModification) SetMethodFilter(v string)`

SetMethodFilter sets MethodFilter field to given value.

### HasMethodFilter

`func (o *CustconfClientRequestModification) HasMethodFilter() bool`

HasMethodFilter returns a boolean if a field has been set.

### GetPathFilter

`func (o *CustconfClientRequestModification) GetPathFilter() string`

GetPathFilter returns the PathFilter field if non-nil, zero value otherwise.

### GetPathFilterOk

`func (o *CustconfClientRequestModification) GetPathFilterOk() (*string, bool)`

GetPathFilterOk returns a tuple with the PathFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathFilter

`func (o *CustconfClientRequestModification) SetPathFilter(v string)`

SetPathFilter sets PathFilter field to given value.

### HasPathFilter

`func (o *CustconfClientRequestModification) HasPathFilter() bool`

HasPathFilter returns a boolean if a field has been set.

### GetHeaderFilter

`func (o *CustconfClientRequestModification) GetHeaderFilter() string`

GetHeaderFilter returns the HeaderFilter field if non-nil, zero value otherwise.

### GetHeaderFilterOk

`func (o *CustconfClientRequestModification) GetHeaderFilterOk() (*string, bool)`

GetHeaderFilterOk returns a tuple with the HeaderFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderFilter

`func (o *CustconfClientRequestModification) SetHeaderFilter(v string)`

SetHeaderFilter sets HeaderFilter field to given value.

### HasHeaderFilter

`func (o *CustconfClientRequestModification) HasHeaderFilter() bool`

HasHeaderFilter returns a boolean if a field has been set.

### GetCookieFilter

`func (o *CustconfClientRequestModification) GetCookieFilter() string`

GetCookieFilter returns the CookieFilter field if non-nil, zero value otherwise.

### GetCookieFilterOk

`func (o *CustconfClientRequestModification) GetCookieFilterOk() (*string, bool)`

GetCookieFilterOk returns a tuple with the CookieFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookieFilter

`func (o *CustconfClientRequestModification) SetCookieFilter(v string)`

SetCookieFilter sets CookieFilter field to given value.

### HasCookieFilter

`func (o *CustconfClientRequestModification) HasCookieFilter() bool`

HasCookieFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


