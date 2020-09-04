# CustconfDynamicContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | This is used by the API to perform conflict checking | [optional] 
**QueryParams** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. A comma separated list of query string parameters you want to include in the cache key generation. NOTE: This list is ignored when the Key Location is set to header. | [optional] 
**HeaderFields** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. A comma-separated list of glob patterns that represent HTTP request headers you want included in the cache key generation. Via the use of a colon &#39;:&#39;, users can define each glob pattern to match a header name only, or the pattern can be used to match both the header name and value. A pattern without a colon &#39;:&#39; is treated as a header name ONLY match. If the pattern matches the header, then all values are used as a part of the cache key. If the pattern contains a colon, the CDN uses the pattern on the left of the colon to match the header. The pattern to the right of the colon is used to match the value. The CDN only uses the header/value as a part of the cache key if both patterns result in a match. Note, no pattern after a colon indicates an empty header (no value). See the fnmatch manpage for acceptable glob patterns. | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**MethodFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**PathFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**HeaderFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 

## Methods

### NewCustconfDynamicContent

`func NewCustconfDynamicContent() *CustconfDynamicContent`

NewCustconfDynamicContent instantiates a new CustconfDynamicContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustconfDynamicContentWithDefaults

`func NewCustconfDynamicContentWithDefaults() *CustconfDynamicContent`

NewCustconfDynamicContentWithDefaults instantiates a new CustconfDynamicContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustconfDynamicContent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustconfDynamicContent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustconfDynamicContent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustconfDynamicContent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetQueryParams

`func (o *CustconfDynamicContent) GetQueryParams() string`

GetQueryParams returns the QueryParams field if non-nil, zero value otherwise.

### GetQueryParamsOk

`func (o *CustconfDynamicContent) GetQueryParamsOk() (*string, bool)`

GetQueryParamsOk returns a tuple with the QueryParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryParams

`func (o *CustconfDynamicContent) SetQueryParams(v string)`

SetQueryParams sets QueryParams field to given value.

### HasQueryParams

`func (o *CustconfDynamicContent) HasQueryParams() bool`

HasQueryParams returns a boolean if a field has been set.

### GetHeaderFields

`func (o *CustconfDynamicContent) GetHeaderFields() string`

GetHeaderFields returns the HeaderFields field if non-nil, zero value otherwise.

### GetHeaderFieldsOk

`func (o *CustconfDynamicContent) GetHeaderFieldsOk() (*string, bool)`

GetHeaderFieldsOk returns a tuple with the HeaderFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderFields

`func (o *CustconfDynamicContent) SetHeaderFields(v string)`

SetHeaderFields sets HeaderFields field to given value.

### HasHeaderFields

`func (o *CustconfDynamicContent) HasHeaderFields() bool`

HasHeaderFields returns a boolean if a field has been set.

### GetEnabled

`func (o *CustconfDynamicContent) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustconfDynamicContent) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustconfDynamicContent) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CustconfDynamicContent) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMethodFilter

`func (o *CustconfDynamicContent) GetMethodFilter() string`

GetMethodFilter returns the MethodFilter field if non-nil, zero value otherwise.

### GetMethodFilterOk

`func (o *CustconfDynamicContent) GetMethodFilterOk() (*string, bool)`

GetMethodFilterOk returns a tuple with the MethodFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodFilter

`func (o *CustconfDynamicContent) SetMethodFilter(v string)`

SetMethodFilter sets MethodFilter field to given value.

### HasMethodFilter

`func (o *CustconfDynamicContent) HasMethodFilter() bool`

HasMethodFilter returns a boolean if a field has been set.

### GetPathFilter

`func (o *CustconfDynamicContent) GetPathFilter() string`

GetPathFilter returns the PathFilter field if non-nil, zero value otherwise.

### GetPathFilterOk

`func (o *CustconfDynamicContent) GetPathFilterOk() (*string, bool)`

GetPathFilterOk returns a tuple with the PathFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathFilter

`func (o *CustconfDynamicContent) SetPathFilter(v string)`

SetPathFilter sets PathFilter field to given value.

### HasPathFilter

`func (o *CustconfDynamicContent) HasPathFilter() bool`

HasPathFilter returns a boolean if a field has been set.

### GetHeaderFilter

`func (o *CustconfDynamicContent) GetHeaderFilter() string`

GetHeaderFilter returns the HeaderFilter field if non-nil, zero value otherwise.

### GetHeaderFilterOk

`func (o *CustconfDynamicContent) GetHeaderFilterOk() (*string, bool)`

GetHeaderFilterOk returns a tuple with the HeaderFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderFilter

`func (o *CustconfDynamicContent) SetHeaderFilter(v string)`

SetHeaderFilter sets HeaderFilter field to given value.

### HasHeaderFilter

`func (o *CustconfDynamicContent) HasHeaderFilter() bool`

HasHeaderFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


