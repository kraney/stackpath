# CustconfDynamicCacheRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | This is used by the API to perform conflict checking | [optional] 
**StatusCode** | Pointer to **float32** | Status code to return | [optional] 
**Headers** | Pointer to **string** | String of values delimited by a &#39;|&#39; character. Pipe delimited (&#39;|&#39;) list of headers to add to response | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**MethodFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**PathFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**HeaderFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 

## Methods

### NewCustconfDynamicCacheRule

`func NewCustconfDynamicCacheRule() *CustconfDynamicCacheRule`

NewCustconfDynamicCacheRule instantiates a new CustconfDynamicCacheRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustconfDynamicCacheRuleWithDefaults

`func NewCustconfDynamicCacheRuleWithDefaults() *CustconfDynamicCacheRule`

NewCustconfDynamicCacheRuleWithDefaults instantiates a new CustconfDynamicCacheRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustconfDynamicCacheRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustconfDynamicCacheRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustconfDynamicCacheRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustconfDynamicCacheRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatusCode

`func (o *CustconfDynamicCacheRule) GetStatusCode() float32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *CustconfDynamicCacheRule) GetStatusCodeOk() (*float32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *CustconfDynamicCacheRule) SetStatusCode(v float32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *CustconfDynamicCacheRule) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetHeaders

`func (o *CustconfDynamicCacheRule) GetHeaders() string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *CustconfDynamicCacheRule) GetHeadersOk() (*string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *CustconfDynamicCacheRule) SetHeaders(v string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *CustconfDynamicCacheRule) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetEnabled

`func (o *CustconfDynamicCacheRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustconfDynamicCacheRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustconfDynamicCacheRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CustconfDynamicCacheRule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMethodFilter

`func (o *CustconfDynamicCacheRule) GetMethodFilter() string`

GetMethodFilter returns the MethodFilter field if non-nil, zero value otherwise.

### GetMethodFilterOk

`func (o *CustconfDynamicCacheRule) GetMethodFilterOk() (*string, bool)`

GetMethodFilterOk returns a tuple with the MethodFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodFilter

`func (o *CustconfDynamicCacheRule) SetMethodFilter(v string)`

SetMethodFilter sets MethodFilter field to given value.

### HasMethodFilter

`func (o *CustconfDynamicCacheRule) HasMethodFilter() bool`

HasMethodFilter returns a boolean if a field has been set.

### GetPathFilter

`func (o *CustconfDynamicCacheRule) GetPathFilter() string`

GetPathFilter returns the PathFilter field if non-nil, zero value otherwise.

### GetPathFilterOk

`func (o *CustconfDynamicCacheRule) GetPathFilterOk() (*string, bool)`

GetPathFilterOk returns a tuple with the PathFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathFilter

`func (o *CustconfDynamicCacheRule) SetPathFilter(v string)`

SetPathFilter sets PathFilter field to given value.

### HasPathFilter

`func (o *CustconfDynamicCacheRule) HasPathFilter() bool`

HasPathFilter returns a boolean if a field has been set.

### GetHeaderFilter

`func (o *CustconfDynamicCacheRule) GetHeaderFilter() string`

GetHeaderFilter returns the HeaderFilter field if non-nil, zero value otherwise.

### GetHeaderFilterOk

`func (o *CustconfDynamicCacheRule) GetHeaderFilterOk() (*string, bool)`

GetHeaderFilterOk returns a tuple with the HeaderFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderFilter

`func (o *CustconfDynamicCacheRule) SetHeaderFilter(v string)`

SetHeaderFilter sets HeaderFilter field to given value.

### HasHeaderFilter

`func (o *CustconfDynamicCacheRule) HasHeaderFilter() bool`

HasHeaderFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


