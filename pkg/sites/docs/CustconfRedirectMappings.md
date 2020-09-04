# CustconfRedirectMappings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | This is used by the API to perform conflict checking | [optional] 
**Code** | Pointer to **float32** | The HTTP cache response code that applies to this policy. | [optional] 
**RedirectURL** | Pointer to **string** | The URL that clients would be redirected to when applying this policy. | [optional] 
**ReplacementToken** | Pointer to **string** | A token that will be replaced by the caching server with the URL of the request that triggered the policy. This token can be positioned anywhere in the redirect URL. | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**MethodFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**PathFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**HeaderFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 

## Methods

### NewCustconfRedirectMappings

`func NewCustconfRedirectMappings() *CustconfRedirectMappings`

NewCustconfRedirectMappings instantiates a new CustconfRedirectMappings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustconfRedirectMappingsWithDefaults

`func NewCustconfRedirectMappingsWithDefaults() *CustconfRedirectMappings`

NewCustconfRedirectMappingsWithDefaults instantiates a new CustconfRedirectMappings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustconfRedirectMappings) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustconfRedirectMappings) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustconfRedirectMappings) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustconfRedirectMappings) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *CustconfRedirectMappings) GetCode() float32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CustconfRedirectMappings) GetCodeOk() (*float32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CustconfRedirectMappings) SetCode(v float32)`

SetCode sets Code field to given value.

### HasCode

`func (o *CustconfRedirectMappings) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetRedirectURL

`func (o *CustconfRedirectMappings) GetRedirectURL() string`

GetRedirectURL returns the RedirectURL field if non-nil, zero value otherwise.

### GetRedirectURLOk

`func (o *CustconfRedirectMappings) GetRedirectURLOk() (*string, bool)`

GetRedirectURLOk returns a tuple with the RedirectURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectURL

`func (o *CustconfRedirectMappings) SetRedirectURL(v string)`

SetRedirectURL sets RedirectURL field to given value.

### HasRedirectURL

`func (o *CustconfRedirectMappings) HasRedirectURL() bool`

HasRedirectURL returns a boolean if a field has been set.

### GetReplacementToken

`func (o *CustconfRedirectMappings) GetReplacementToken() string`

GetReplacementToken returns the ReplacementToken field if non-nil, zero value otherwise.

### GetReplacementTokenOk

`func (o *CustconfRedirectMappings) GetReplacementTokenOk() (*string, bool)`

GetReplacementTokenOk returns a tuple with the ReplacementToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacementToken

`func (o *CustconfRedirectMappings) SetReplacementToken(v string)`

SetReplacementToken sets ReplacementToken field to given value.

### HasReplacementToken

`func (o *CustconfRedirectMappings) HasReplacementToken() bool`

HasReplacementToken returns a boolean if a field has been set.

### GetEnabled

`func (o *CustconfRedirectMappings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustconfRedirectMappings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustconfRedirectMappings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CustconfRedirectMappings) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMethodFilter

`func (o *CustconfRedirectMappings) GetMethodFilter() string`

GetMethodFilter returns the MethodFilter field if non-nil, zero value otherwise.

### GetMethodFilterOk

`func (o *CustconfRedirectMappings) GetMethodFilterOk() (*string, bool)`

GetMethodFilterOk returns a tuple with the MethodFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodFilter

`func (o *CustconfRedirectMappings) SetMethodFilter(v string)`

SetMethodFilter sets MethodFilter field to given value.

### HasMethodFilter

`func (o *CustconfRedirectMappings) HasMethodFilter() bool`

HasMethodFilter returns a boolean if a field has been set.

### GetPathFilter

`func (o *CustconfRedirectMappings) GetPathFilter() string`

GetPathFilter returns the PathFilter field if non-nil, zero value otherwise.

### GetPathFilterOk

`func (o *CustconfRedirectMappings) GetPathFilterOk() (*string, bool)`

GetPathFilterOk returns a tuple with the PathFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathFilter

`func (o *CustconfRedirectMappings) SetPathFilter(v string)`

SetPathFilter sets PathFilter field to given value.

### HasPathFilter

`func (o *CustconfRedirectMappings) HasPathFilter() bool`

HasPathFilter returns a boolean if a field has been set.

### GetHeaderFilter

`func (o *CustconfRedirectMappings) GetHeaderFilter() string`

GetHeaderFilter returns the HeaderFilter field if non-nil, zero value otherwise.

### GetHeaderFilterOk

`func (o *CustconfRedirectMappings) GetHeaderFilterOk() (*string, bool)`

GetHeaderFilterOk returns a tuple with the HeaderFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderFilter

`func (o *CustconfRedirectMappings) SetHeaderFilter(v string)`

SetHeaderFilter sets HeaderFilter field to given value.

### HasHeaderFilter

`func (o *CustconfRedirectMappings) HasHeaderFilter() bool`

HasHeaderFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


