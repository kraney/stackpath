# CustconfAuthUrlSignAliCloudC

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | This is used by the API to perform conflict checking | [optional] 
**PassPhrase** | Pointer to **string** |  | [optional] 
**ExpirationExtension** | Pointer to **float32** |  | [optional] 
**TokenField** | Pointer to **string** |  | [optional] 
**ExpireField** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**MethodFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**PathFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**HeaderFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 

## Methods

### NewCustconfAuthUrlSignAliCloudC

`func NewCustconfAuthUrlSignAliCloudC() *CustconfAuthUrlSignAliCloudC`

NewCustconfAuthUrlSignAliCloudC instantiates a new CustconfAuthUrlSignAliCloudC object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustconfAuthUrlSignAliCloudCWithDefaults

`func NewCustconfAuthUrlSignAliCloudCWithDefaults() *CustconfAuthUrlSignAliCloudC`

NewCustconfAuthUrlSignAliCloudCWithDefaults instantiates a new CustconfAuthUrlSignAliCloudC object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustconfAuthUrlSignAliCloudC) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustconfAuthUrlSignAliCloudC) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustconfAuthUrlSignAliCloudC) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustconfAuthUrlSignAliCloudC) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPassPhrase

`func (o *CustconfAuthUrlSignAliCloudC) GetPassPhrase() string`

GetPassPhrase returns the PassPhrase field if non-nil, zero value otherwise.

### GetPassPhraseOk

`func (o *CustconfAuthUrlSignAliCloudC) GetPassPhraseOk() (*string, bool)`

GetPassPhraseOk returns a tuple with the PassPhrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassPhrase

`func (o *CustconfAuthUrlSignAliCloudC) SetPassPhrase(v string)`

SetPassPhrase sets PassPhrase field to given value.

### HasPassPhrase

`func (o *CustconfAuthUrlSignAliCloudC) HasPassPhrase() bool`

HasPassPhrase returns a boolean if a field has been set.

### GetExpirationExtension

`func (o *CustconfAuthUrlSignAliCloudC) GetExpirationExtension() float32`

GetExpirationExtension returns the ExpirationExtension field if non-nil, zero value otherwise.

### GetExpirationExtensionOk

`func (o *CustconfAuthUrlSignAliCloudC) GetExpirationExtensionOk() (*float32, bool)`

GetExpirationExtensionOk returns a tuple with the ExpirationExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationExtension

`func (o *CustconfAuthUrlSignAliCloudC) SetExpirationExtension(v float32)`

SetExpirationExtension sets ExpirationExtension field to given value.

### HasExpirationExtension

`func (o *CustconfAuthUrlSignAliCloudC) HasExpirationExtension() bool`

HasExpirationExtension returns a boolean if a field has been set.

### GetTokenField

`func (o *CustconfAuthUrlSignAliCloudC) GetTokenField() string`

GetTokenField returns the TokenField field if non-nil, zero value otherwise.

### GetTokenFieldOk

`func (o *CustconfAuthUrlSignAliCloudC) GetTokenFieldOk() (*string, bool)`

GetTokenFieldOk returns a tuple with the TokenField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenField

`func (o *CustconfAuthUrlSignAliCloudC) SetTokenField(v string)`

SetTokenField sets TokenField field to given value.

### HasTokenField

`func (o *CustconfAuthUrlSignAliCloudC) HasTokenField() bool`

HasTokenField returns a boolean if a field has been set.

### GetExpireField

`func (o *CustconfAuthUrlSignAliCloudC) GetExpireField() string`

GetExpireField returns the ExpireField field if non-nil, zero value otherwise.

### GetExpireFieldOk

`func (o *CustconfAuthUrlSignAliCloudC) GetExpireFieldOk() (*string, bool)`

GetExpireFieldOk returns a tuple with the ExpireField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireField

`func (o *CustconfAuthUrlSignAliCloudC) SetExpireField(v string)`

SetExpireField sets ExpireField field to given value.

### HasExpireField

`func (o *CustconfAuthUrlSignAliCloudC) HasExpireField() bool`

HasExpireField returns a boolean if a field has been set.

### GetEnabled

`func (o *CustconfAuthUrlSignAliCloudC) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustconfAuthUrlSignAliCloudC) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustconfAuthUrlSignAliCloudC) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CustconfAuthUrlSignAliCloudC) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMethodFilter

`func (o *CustconfAuthUrlSignAliCloudC) GetMethodFilter() string`

GetMethodFilter returns the MethodFilter field if non-nil, zero value otherwise.

### GetMethodFilterOk

`func (o *CustconfAuthUrlSignAliCloudC) GetMethodFilterOk() (*string, bool)`

GetMethodFilterOk returns a tuple with the MethodFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodFilter

`func (o *CustconfAuthUrlSignAliCloudC) SetMethodFilter(v string)`

SetMethodFilter sets MethodFilter field to given value.

### HasMethodFilter

`func (o *CustconfAuthUrlSignAliCloudC) HasMethodFilter() bool`

HasMethodFilter returns a boolean if a field has been set.

### GetPathFilter

`func (o *CustconfAuthUrlSignAliCloudC) GetPathFilter() string`

GetPathFilter returns the PathFilter field if non-nil, zero value otherwise.

### GetPathFilterOk

`func (o *CustconfAuthUrlSignAliCloudC) GetPathFilterOk() (*string, bool)`

GetPathFilterOk returns a tuple with the PathFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathFilter

`func (o *CustconfAuthUrlSignAliCloudC) SetPathFilter(v string)`

SetPathFilter sets PathFilter field to given value.

### HasPathFilter

`func (o *CustconfAuthUrlSignAliCloudC) HasPathFilter() bool`

HasPathFilter returns a boolean if a field has been set.

### GetHeaderFilter

`func (o *CustconfAuthUrlSignAliCloudC) GetHeaderFilter() string`

GetHeaderFilter returns the HeaderFilter field if non-nil, zero value otherwise.

### GetHeaderFilterOk

`func (o *CustconfAuthUrlSignAliCloudC) GetHeaderFilterOk() (*string, bool)`

GetHeaderFilterOk returns a tuple with the HeaderFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderFilter

`func (o *CustconfAuthUrlSignAliCloudC) SetHeaderFilter(v string)`

SetHeaderFilter sets HeaderFilter field to given value.

### HasHeaderFilter

`func (o *CustconfAuthUrlSignAliCloudC) HasHeaderFilter() bool`

HasHeaderFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


