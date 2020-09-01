# CustconfAuthUrlSignAliCloudA

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | This is used by the API to perform conflict checking | [optional] 
**PassPhrase** | Pointer to **string** |  | [optional] 
**TokenField** | Pointer to **string** |  | [optional] 
**IncludeParamsBeforeToken** | Pointer to **bool** |  | [optional] 
**ExpirationExtension** | Pointer to **float32** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**MethodFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**PathFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**HeaderFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 

## Methods

### NewCustconfAuthUrlSignAliCloudA

`func NewCustconfAuthUrlSignAliCloudA() *CustconfAuthUrlSignAliCloudA`

NewCustconfAuthUrlSignAliCloudA instantiates a new CustconfAuthUrlSignAliCloudA object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustconfAuthUrlSignAliCloudAWithDefaults

`func NewCustconfAuthUrlSignAliCloudAWithDefaults() *CustconfAuthUrlSignAliCloudA`

NewCustconfAuthUrlSignAliCloudAWithDefaults instantiates a new CustconfAuthUrlSignAliCloudA object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustconfAuthUrlSignAliCloudA) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustconfAuthUrlSignAliCloudA) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustconfAuthUrlSignAliCloudA) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustconfAuthUrlSignAliCloudA) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPassPhrase

`func (o *CustconfAuthUrlSignAliCloudA) GetPassPhrase() string`

GetPassPhrase returns the PassPhrase field if non-nil, zero value otherwise.

### GetPassPhraseOk

`func (o *CustconfAuthUrlSignAliCloudA) GetPassPhraseOk() (*string, bool)`

GetPassPhraseOk returns a tuple with the PassPhrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassPhrase

`func (o *CustconfAuthUrlSignAliCloudA) SetPassPhrase(v string)`

SetPassPhrase sets PassPhrase field to given value.

### HasPassPhrase

`func (o *CustconfAuthUrlSignAliCloudA) HasPassPhrase() bool`

HasPassPhrase returns a boolean if a field has been set.

### GetTokenField

`func (o *CustconfAuthUrlSignAliCloudA) GetTokenField() string`

GetTokenField returns the TokenField field if non-nil, zero value otherwise.

### GetTokenFieldOk

`func (o *CustconfAuthUrlSignAliCloudA) GetTokenFieldOk() (*string, bool)`

GetTokenFieldOk returns a tuple with the TokenField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenField

`func (o *CustconfAuthUrlSignAliCloudA) SetTokenField(v string)`

SetTokenField sets TokenField field to given value.

### HasTokenField

`func (o *CustconfAuthUrlSignAliCloudA) HasTokenField() bool`

HasTokenField returns a boolean if a field has been set.

### GetIncludeParamsBeforeToken

`func (o *CustconfAuthUrlSignAliCloudA) GetIncludeParamsBeforeToken() bool`

GetIncludeParamsBeforeToken returns the IncludeParamsBeforeToken field if non-nil, zero value otherwise.

### GetIncludeParamsBeforeTokenOk

`func (o *CustconfAuthUrlSignAliCloudA) GetIncludeParamsBeforeTokenOk() (*bool, bool)`

GetIncludeParamsBeforeTokenOk returns a tuple with the IncludeParamsBeforeToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeParamsBeforeToken

`func (o *CustconfAuthUrlSignAliCloudA) SetIncludeParamsBeforeToken(v bool)`

SetIncludeParamsBeforeToken sets IncludeParamsBeforeToken field to given value.

### HasIncludeParamsBeforeToken

`func (o *CustconfAuthUrlSignAliCloudA) HasIncludeParamsBeforeToken() bool`

HasIncludeParamsBeforeToken returns a boolean if a field has been set.

### GetExpirationExtension

`func (o *CustconfAuthUrlSignAliCloudA) GetExpirationExtension() float32`

GetExpirationExtension returns the ExpirationExtension field if non-nil, zero value otherwise.

### GetExpirationExtensionOk

`func (o *CustconfAuthUrlSignAliCloudA) GetExpirationExtensionOk() (*float32, bool)`

GetExpirationExtensionOk returns a tuple with the ExpirationExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationExtension

`func (o *CustconfAuthUrlSignAliCloudA) SetExpirationExtension(v float32)`

SetExpirationExtension sets ExpirationExtension field to given value.

### HasExpirationExtension

`func (o *CustconfAuthUrlSignAliCloudA) HasExpirationExtension() bool`

HasExpirationExtension returns a boolean if a field has been set.

### GetEnabled

`func (o *CustconfAuthUrlSignAliCloudA) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustconfAuthUrlSignAliCloudA) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustconfAuthUrlSignAliCloudA) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CustconfAuthUrlSignAliCloudA) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMethodFilter

`func (o *CustconfAuthUrlSignAliCloudA) GetMethodFilter() string`

GetMethodFilter returns the MethodFilter field if non-nil, zero value otherwise.

### GetMethodFilterOk

`func (o *CustconfAuthUrlSignAliCloudA) GetMethodFilterOk() (*string, bool)`

GetMethodFilterOk returns a tuple with the MethodFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodFilter

`func (o *CustconfAuthUrlSignAliCloudA) SetMethodFilter(v string)`

SetMethodFilter sets MethodFilter field to given value.

### HasMethodFilter

`func (o *CustconfAuthUrlSignAliCloudA) HasMethodFilter() bool`

HasMethodFilter returns a boolean if a field has been set.

### GetPathFilter

`func (o *CustconfAuthUrlSignAliCloudA) GetPathFilter() string`

GetPathFilter returns the PathFilter field if non-nil, zero value otherwise.

### GetPathFilterOk

`func (o *CustconfAuthUrlSignAliCloudA) GetPathFilterOk() (*string, bool)`

GetPathFilterOk returns a tuple with the PathFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathFilter

`func (o *CustconfAuthUrlSignAliCloudA) SetPathFilter(v string)`

SetPathFilter sets PathFilter field to given value.

### HasPathFilter

`func (o *CustconfAuthUrlSignAliCloudA) HasPathFilter() bool`

HasPathFilter returns a boolean if a field has been set.

### GetHeaderFilter

`func (o *CustconfAuthUrlSignAliCloudA) GetHeaderFilter() string`

GetHeaderFilter returns the HeaderFilter field if non-nil, zero value otherwise.

### GetHeaderFilterOk

`func (o *CustconfAuthUrlSignAliCloudA) GetHeaderFilterOk() (*string, bool)`

GetHeaderFilterOk returns a tuple with the HeaderFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderFilter

`func (o *CustconfAuthUrlSignAliCloudA) SetHeaderFilter(v string)`

SetHeaderFilter sets HeaderFilter field to given value.

### HasHeaderFilter

`func (o *CustconfAuthUrlSignAliCloudA) HasHeaderFilter() bool`

HasHeaderFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


