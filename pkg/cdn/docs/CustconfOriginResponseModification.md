# CustconfOriginResponseModification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | This is used by the API to perform conflict checking | [optional] 
**StatusCodeRewrite** | Pointer to **float32** |  | [optional] 
**HeaderPattern** | Pointer to **string** |  | [optional] 
**HeaderRewrite** | Pointer to **string** |  | [optional] 
**AddHeaders** | Pointer to **string** | String of values delimited by a &#39;|&#39; character. This is the static HTTP header you want inserted into the CDN response. Start value with \&quot;append:\&quot;, \&quot;replace:\&quot; or \&quot;create:\&quot; which defines if Header will be added, replaced or create if not exists. Default is append. | [optional] 
**FlowControl** | Pointer to [**CustconfOriginResponseModificationFlowControlEnumWrapperValue**](custconfOriginResponseModificationFlowControlEnumWrapperValue.md) |  | [optional] [default to "UNKNOWN"]
**Enabled** | Pointer to **bool** |  | [optional] 
**MethodFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**PathFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**HeaderFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 

## Methods

### NewCustconfOriginResponseModification

`func NewCustconfOriginResponseModification() *CustconfOriginResponseModification`

NewCustconfOriginResponseModification instantiates a new CustconfOriginResponseModification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustconfOriginResponseModificationWithDefaults

`func NewCustconfOriginResponseModificationWithDefaults() *CustconfOriginResponseModification`

NewCustconfOriginResponseModificationWithDefaults instantiates a new CustconfOriginResponseModification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustconfOriginResponseModification) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustconfOriginResponseModification) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustconfOriginResponseModification) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustconfOriginResponseModification) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatusCodeRewrite

`func (o *CustconfOriginResponseModification) GetStatusCodeRewrite() float32`

GetStatusCodeRewrite returns the StatusCodeRewrite field if non-nil, zero value otherwise.

### GetStatusCodeRewriteOk

`func (o *CustconfOriginResponseModification) GetStatusCodeRewriteOk() (*float32, bool)`

GetStatusCodeRewriteOk returns a tuple with the StatusCodeRewrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCodeRewrite

`func (o *CustconfOriginResponseModification) SetStatusCodeRewrite(v float32)`

SetStatusCodeRewrite sets StatusCodeRewrite field to given value.

### HasStatusCodeRewrite

`func (o *CustconfOriginResponseModification) HasStatusCodeRewrite() bool`

HasStatusCodeRewrite returns a boolean if a field has been set.

### GetHeaderPattern

`func (o *CustconfOriginResponseModification) GetHeaderPattern() string`

GetHeaderPattern returns the HeaderPattern field if non-nil, zero value otherwise.

### GetHeaderPatternOk

`func (o *CustconfOriginResponseModification) GetHeaderPatternOk() (*string, bool)`

GetHeaderPatternOk returns a tuple with the HeaderPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderPattern

`func (o *CustconfOriginResponseModification) SetHeaderPattern(v string)`

SetHeaderPattern sets HeaderPattern field to given value.

### HasHeaderPattern

`func (o *CustconfOriginResponseModification) HasHeaderPattern() bool`

HasHeaderPattern returns a boolean if a field has been set.

### GetHeaderRewrite

`func (o *CustconfOriginResponseModification) GetHeaderRewrite() string`

GetHeaderRewrite returns the HeaderRewrite field if non-nil, zero value otherwise.

### GetHeaderRewriteOk

`func (o *CustconfOriginResponseModification) GetHeaderRewriteOk() (*string, bool)`

GetHeaderRewriteOk returns a tuple with the HeaderRewrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderRewrite

`func (o *CustconfOriginResponseModification) SetHeaderRewrite(v string)`

SetHeaderRewrite sets HeaderRewrite field to given value.

### HasHeaderRewrite

`func (o *CustconfOriginResponseModification) HasHeaderRewrite() bool`

HasHeaderRewrite returns a boolean if a field has been set.

### GetAddHeaders

`func (o *CustconfOriginResponseModification) GetAddHeaders() string`

GetAddHeaders returns the AddHeaders field if non-nil, zero value otherwise.

### GetAddHeadersOk

`func (o *CustconfOriginResponseModification) GetAddHeadersOk() (*string, bool)`

GetAddHeadersOk returns a tuple with the AddHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddHeaders

`func (o *CustconfOriginResponseModification) SetAddHeaders(v string)`

SetAddHeaders sets AddHeaders field to given value.

### HasAddHeaders

`func (o *CustconfOriginResponseModification) HasAddHeaders() bool`

HasAddHeaders returns a boolean if a field has been set.

### GetFlowControl

`func (o *CustconfOriginResponseModification) GetFlowControl() CustconfOriginResponseModificationFlowControlEnumWrapperValue`

GetFlowControl returns the FlowControl field if non-nil, zero value otherwise.

### GetFlowControlOk

`func (o *CustconfOriginResponseModification) GetFlowControlOk() (*CustconfOriginResponseModificationFlowControlEnumWrapperValue, bool)`

GetFlowControlOk returns a tuple with the FlowControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowControl

`func (o *CustconfOriginResponseModification) SetFlowControl(v CustconfOriginResponseModificationFlowControlEnumWrapperValue)`

SetFlowControl sets FlowControl field to given value.

### HasFlowControl

`func (o *CustconfOriginResponseModification) HasFlowControl() bool`

HasFlowControl returns a boolean if a field has been set.

### GetEnabled

`func (o *CustconfOriginResponseModification) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustconfOriginResponseModification) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustconfOriginResponseModification) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CustconfOriginResponseModification) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMethodFilter

`func (o *CustconfOriginResponseModification) GetMethodFilter() string`

GetMethodFilter returns the MethodFilter field if non-nil, zero value otherwise.

### GetMethodFilterOk

`func (o *CustconfOriginResponseModification) GetMethodFilterOk() (*string, bool)`

GetMethodFilterOk returns a tuple with the MethodFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodFilter

`func (o *CustconfOriginResponseModification) SetMethodFilter(v string)`

SetMethodFilter sets MethodFilter field to given value.

### HasMethodFilter

`func (o *CustconfOriginResponseModification) HasMethodFilter() bool`

HasMethodFilter returns a boolean if a field has been set.

### GetPathFilter

`func (o *CustconfOriginResponseModification) GetPathFilter() string`

GetPathFilter returns the PathFilter field if non-nil, zero value otherwise.

### GetPathFilterOk

`func (o *CustconfOriginResponseModification) GetPathFilterOk() (*string, bool)`

GetPathFilterOk returns a tuple with the PathFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathFilter

`func (o *CustconfOriginResponseModification) SetPathFilter(v string)`

SetPathFilter sets PathFilter field to given value.

### HasPathFilter

`func (o *CustconfOriginResponseModification) HasPathFilter() bool`

HasPathFilter returns a boolean if a field has been set.

### GetHeaderFilter

`func (o *CustconfOriginResponseModification) GetHeaderFilter() string`

GetHeaderFilter returns the HeaderFilter field if non-nil, zero value otherwise.

### GetHeaderFilterOk

`func (o *CustconfOriginResponseModification) GetHeaderFilterOk() (*string, bool)`

GetHeaderFilterOk returns a tuple with the HeaderFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderFilter

`func (o *CustconfOriginResponseModification) SetHeaderFilter(v string)`

SetHeaderFilter sets HeaderFilter field to given value.

### HasHeaderFilter

`func (o *CustconfOriginResponseModification) HasHeaderFilter() bool`

HasHeaderFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


