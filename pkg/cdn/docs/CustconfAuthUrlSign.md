# CustconfAuthUrlSign

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | This is used by the API to perform conflict checking | [optional] 
**TokenField** | Pointer to **string** | This is the name of the query string parameter that will be used by the publisher to specify the signature for the URL. | [optional] 
**IgnoreFieldsAfterToken** | Pointer to **bool** | Select this option if you want StackPath to exclude query string parameters specified after the passphrase in the validation process. | [optional] 
**PassPhraseField** | Pointer to **string** | This is the name of the query string parameter that contains the value of the secret. This query string parameter is only used during the generation and validation of a URL signature and should not be present in the published URL. | [optional] 
**PassPhrase** | Pointer to **string** | The shared secret used during the signing process. This value should only be known by StackPath and systems authorized to sign your content. | [optional] 
**ExpiresField** | Pointer to **string** | This is the name of the query string parameter that contains the Epoch time after which the URL is considered invalid. | [optional] 
**IpAddressField** | Pointer to **string** | This is a query string parameter that contains an IPv4 address to which the published URL will be restricted. | [optional] 
**UriLengthField** | Pointer to **string** | This is a query string parameter used to limit the number of characters in the path that should be considered when validating the URL signature. This feature allows you to reuse the same signature on all assets stored in the same cache path. Setting this value to 0 will strip off the filename in the URL (leaving the trailing slash) when calculating the checksum. | [optional] 
**UserAgentField** | Pointer to **string** | This is a query string parameter used to restrict the published URL to a specific user agent. Publishers can use this feature during the signing process to ensure that only a specific user agent can access the published content. You do not need to specify the user agent on the published URL. StackPath will use the HTTP User-Agent header value during signature validation. | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**MethodFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**PathFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**HeaderFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 

## Methods

### NewCustconfAuthUrlSign

`func NewCustconfAuthUrlSign() *CustconfAuthUrlSign`

NewCustconfAuthUrlSign instantiates a new CustconfAuthUrlSign object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustconfAuthUrlSignWithDefaults

`func NewCustconfAuthUrlSignWithDefaults() *CustconfAuthUrlSign`

NewCustconfAuthUrlSignWithDefaults instantiates a new CustconfAuthUrlSign object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustconfAuthUrlSign) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustconfAuthUrlSign) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustconfAuthUrlSign) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustconfAuthUrlSign) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTokenField

`func (o *CustconfAuthUrlSign) GetTokenField() string`

GetTokenField returns the TokenField field if non-nil, zero value otherwise.

### GetTokenFieldOk

`func (o *CustconfAuthUrlSign) GetTokenFieldOk() (*string, bool)`

GetTokenFieldOk returns a tuple with the TokenField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenField

`func (o *CustconfAuthUrlSign) SetTokenField(v string)`

SetTokenField sets TokenField field to given value.

### HasTokenField

`func (o *CustconfAuthUrlSign) HasTokenField() bool`

HasTokenField returns a boolean if a field has been set.

### GetIgnoreFieldsAfterToken

`func (o *CustconfAuthUrlSign) GetIgnoreFieldsAfterToken() bool`

GetIgnoreFieldsAfterToken returns the IgnoreFieldsAfterToken field if non-nil, zero value otherwise.

### GetIgnoreFieldsAfterTokenOk

`func (o *CustconfAuthUrlSign) GetIgnoreFieldsAfterTokenOk() (*bool, bool)`

GetIgnoreFieldsAfterTokenOk returns a tuple with the IgnoreFieldsAfterToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreFieldsAfterToken

`func (o *CustconfAuthUrlSign) SetIgnoreFieldsAfterToken(v bool)`

SetIgnoreFieldsAfterToken sets IgnoreFieldsAfterToken field to given value.

### HasIgnoreFieldsAfterToken

`func (o *CustconfAuthUrlSign) HasIgnoreFieldsAfterToken() bool`

HasIgnoreFieldsAfterToken returns a boolean if a field has been set.

### GetPassPhraseField

`func (o *CustconfAuthUrlSign) GetPassPhraseField() string`

GetPassPhraseField returns the PassPhraseField field if non-nil, zero value otherwise.

### GetPassPhraseFieldOk

`func (o *CustconfAuthUrlSign) GetPassPhraseFieldOk() (*string, bool)`

GetPassPhraseFieldOk returns a tuple with the PassPhraseField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassPhraseField

`func (o *CustconfAuthUrlSign) SetPassPhraseField(v string)`

SetPassPhraseField sets PassPhraseField field to given value.

### HasPassPhraseField

`func (o *CustconfAuthUrlSign) HasPassPhraseField() bool`

HasPassPhraseField returns a boolean if a field has been set.

### GetPassPhrase

`func (o *CustconfAuthUrlSign) GetPassPhrase() string`

GetPassPhrase returns the PassPhrase field if non-nil, zero value otherwise.

### GetPassPhraseOk

`func (o *CustconfAuthUrlSign) GetPassPhraseOk() (*string, bool)`

GetPassPhraseOk returns a tuple with the PassPhrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassPhrase

`func (o *CustconfAuthUrlSign) SetPassPhrase(v string)`

SetPassPhrase sets PassPhrase field to given value.

### HasPassPhrase

`func (o *CustconfAuthUrlSign) HasPassPhrase() bool`

HasPassPhrase returns a boolean if a field has been set.

### GetExpiresField

`func (o *CustconfAuthUrlSign) GetExpiresField() string`

GetExpiresField returns the ExpiresField field if non-nil, zero value otherwise.

### GetExpiresFieldOk

`func (o *CustconfAuthUrlSign) GetExpiresFieldOk() (*string, bool)`

GetExpiresFieldOk returns a tuple with the ExpiresField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresField

`func (o *CustconfAuthUrlSign) SetExpiresField(v string)`

SetExpiresField sets ExpiresField field to given value.

### HasExpiresField

`func (o *CustconfAuthUrlSign) HasExpiresField() bool`

HasExpiresField returns a boolean if a field has been set.

### GetIpAddressField

`func (o *CustconfAuthUrlSign) GetIpAddressField() string`

GetIpAddressField returns the IpAddressField field if non-nil, zero value otherwise.

### GetIpAddressFieldOk

`func (o *CustconfAuthUrlSign) GetIpAddressFieldOk() (*string, bool)`

GetIpAddressFieldOk returns a tuple with the IpAddressField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddressField

`func (o *CustconfAuthUrlSign) SetIpAddressField(v string)`

SetIpAddressField sets IpAddressField field to given value.

### HasIpAddressField

`func (o *CustconfAuthUrlSign) HasIpAddressField() bool`

HasIpAddressField returns a boolean if a field has been set.

### GetUriLengthField

`func (o *CustconfAuthUrlSign) GetUriLengthField() string`

GetUriLengthField returns the UriLengthField field if non-nil, zero value otherwise.

### GetUriLengthFieldOk

`func (o *CustconfAuthUrlSign) GetUriLengthFieldOk() (*string, bool)`

GetUriLengthFieldOk returns a tuple with the UriLengthField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUriLengthField

`func (o *CustconfAuthUrlSign) SetUriLengthField(v string)`

SetUriLengthField sets UriLengthField field to given value.

### HasUriLengthField

`func (o *CustconfAuthUrlSign) HasUriLengthField() bool`

HasUriLengthField returns a boolean if a field has been set.

### GetUserAgentField

`func (o *CustconfAuthUrlSign) GetUserAgentField() string`

GetUserAgentField returns the UserAgentField field if non-nil, zero value otherwise.

### GetUserAgentFieldOk

`func (o *CustconfAuthUrlSign) GetUserAgentFieldOk() (*string, bool)`

GetUserAgentFieldOk returns a tuple with the UserAgentField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgentField

`func (o *CustconfAuthUrlSign) SetUserAgentField(v string)`

SetUserAgentField sets UserAgentField field to given value.

### HasUserAgentField

`func (o *CustconfAuthUrlSign) HasUserAgentField() bool`

HasUserAgentField returns a boolean if a field has been set.

### GetEnabled

`func (o *CustconfAuthUrlSign) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustconfAuthUrlSign) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustconfAuthUrlSign) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CustconfAuthUrlSign) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMethodFilter

`func (o *CustconfAuthUrlSign) GetMethodFilter() string`

GetMethodFilter returns the MethodFilter field if non-nil, zero value otherwise.

### GetMethodFilterOk

`func (o *CustconfAuthUrlSign) GetMethodFilterOk() (*string, bool)`

GetMethodFilterOk returns a tuple with the MethodFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodFilter

`func (o *CustconfAuthUrlSign) SetMethodFilter(v string)`

SetMethodFilter sets MethodFilter field to given value.

### HasMethodFilter

`func (o *CustconfAuthUrlSign) HasMethodFilter() bool`

HasMethodFilter returns a boolean if a field has been set.

### GetPathFilter

`func (o *CustconfAuthUrlSign) GetPathFilter() string`

GetPathFilter returns the PathFilter field if non-nil, zero value otherwise.

### GetPathFilterOk

`func (o *CustconfAuthUrlSign) GetPathFilterOk() (*string, bool)`

GetPathFilterOk returns a tuple with the PathFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathFilter

`func (o *CustconfAuthUrlSign) SetPathFilter(v string)`

SetPathFilter sets PathFilter field to given value.

### HasPathFilter

`func (o *CustconfAuthUrlSign) HasPathFilter() bool`

HasPathFilter returns a boolean if a field has been set.

### GetHeaderFilter

`func (o *CustconfAuthUrlSign) GetHeaderFilter() string`

GetHeaderFilter returns the HeaderFilter field if non-nil, zero value otherwise.

### GetHeaderFilterOk

`func (o *CustconfAuthUrlSign) GetHeaderFilterOk() (*string, bool)`

GetHeaderFilterOk returns a tuple with the HeaderFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderFilter

`func (o *CustconfAuthUrlSign) SetHeaderFilter(v string)`

SetHeaderFilter sets HeaderFilter field to given value.

### HasHeaderFilter

`func (o *CustconfAuthUrlSign) HasHeaderFilter() bool`

HasHeaderFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


