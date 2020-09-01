# CustconfCustomMimeType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | This is used by the API to perform conflict checking | [optional] 
**Code** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. A comma separated list of status codes that apply to this policy | [optional] 
**ExtensionMap** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. This is a comma separated list of file extension and mime type pairs that describe the mime type mapping for the CDN caching servers. The file extension and mime type values should be delimited by a colon (:). For example, to map files ending with jpg to the image/jpeg mime type and files ending with png to image/png, set this value to: jpg:image/jpeg,png:image/png. | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**MethodFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**PathFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**HeaderFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 

## Methods

### NewCustconfCustomMimeType

`func NewCustconfCustomMimeType() *CustconfCustomMimeType`

NewCustconfCustomMimeType instantiates a new CustconfCustomMimeType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustconfCustomMimeTypeWithDefaults

`func NewCustconfCustomMimeTypeWithDefaults() *CustconfCustomMimeType`

NewCustconfCustomMimeTypeWithDefaults instantiates a new CustconfCustomMimeType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustconfCustomMimeType) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustconfCustomMimeType) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustconfCustomMimeType) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustconfCustomMimeType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *CustconfCustomMimeType) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CustconfCustomMimeType) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CustconfCustomMimeType) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CustconfCustomMimeType) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetExtensionMap

`func (o *CustconfCustomMimeType) GetExtensionMap() string`

GetExtensionMap returns the ExtensionMap field if non-nil, zero value otherwise.

### GetExtensionMapOk

`func (o *CustconfCustomMimeType) GetExtensionMapOk() (*string, bool)`

GetExtensionMapOk returns a tuple with the ExtensionMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionMap

`func (o *CustconfCustomMimeType) SetExtensionMap(v string)`

SetExtensionMap sets ExtensionMap field to given value.

### HasExtensionMap

`func (o *CustconfCustomMimeType) HasExtensionMap() bool`

HasExtensionMap returns a boolean if a field has been set.

### GetEnabled

`func (o *CustconfCustomMimeType) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustconfCustomMimeType) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustconfCustomMimeType) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CustconfCustomMimeType) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMethodFilter

`func (o *CustconfCustomMimeType) GetMethodFilter() string`

GetMethodFilter returns the MethodFilter field if non-nil, zero value otherwise.

### GetMethodFilterOk

`func (o *CustconfCustomMimeType) GetMethodFilterOk() (*string, bool)`

GetMethodFilterOk returns a tuple with the MethodFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodFilter

`func (o *CustconfCustomMimeType) SetMethodFilter(v string)`

SetMethodFilter sets MethodFilter field to given value.

### HasMethodFilter

`func (o *CustconfCustomMimeType) HasMethodFilter() bool`

HasMethodFilter returns a boolean if a field has been set.

### GetPathFilter

`func (o *CustconfCustomMimeType) GetPathFilter() string`

GetPathFilter returns the PathFilter field if non-nil, zero value otherwise.

### GetPathFilterOk

`func (o *CustconfCustomMimeType) GetPathFilterOk() (*string, bool)`

GetPathFilterOk returns a tuple with the PathFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathFilter

`func (o *CustconfCustomMimeType) SetPathFilter(v string)`

SetPathFilter sets PathFilter field to given value.

### HasPathFilter

`func (o *CustconfCustomMimeType) HasPathFilter() bool`

HasPathFilter returns a boolean if a field has been set.

### GetHeaderFilter

`func (o *CustconfCustomMimeType) GetHeaderFilter() string`

GetHeaderFilter returns the HeaderFilter field if non-nil, zero value otherwise.

### GetHeaderFilterOk

`func (o *CustconfCustomMimeType) GetHeaderFilterOk() (*string, bool)`

GetHeaderFilterOk returns a tuple with the HeaderFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderFilter

`func (o *CustconfCustomMimeType) SetHeaderFilter(v string)`

SetHeaderFilter sets HeaderFilter field to given value.

### HasHeaderFilter

`func (o *CustconfCustomMimeType) HasHeaderFilter() bool`

HasHeaderFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


