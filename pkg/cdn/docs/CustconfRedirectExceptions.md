# CustconfRedirectExceptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | This is used by the API to perform conflict checking | [optional] 
**RedirectAgentCode** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. This is a comma separated list of user agents and redirect code pairs. The user agent and redirect code values are separated by a colon (:), and you may use wildcards in the user agent field. For example, to map assign a 307 status code to all Chrome browsers, you would specify: *chrome*:307. | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewCustconfRedirectExceptions

`func NewCustconfRedirectExceptions() *CustconfRedirectExceptions`

NewCustconfRedirectExceptions instantiates a new CustconfRedirectExceptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustconfRedirectExceptionsWithDefaults

`func NewCustconfRedirectExceptionsWithDefaults() *CustconfRedirectExceptions`

NewCustconfRedirectExceptionsWithDefaults instantiates a new CustconfRedirectExceptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustconfRedirectExceptions) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustconfRedirectExceptions) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustconfRedirectExceptions) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustconfRedirectExceptions) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRedirectAgentCode

`func (o *CustconfRedirectExceptions) GetRedirectAgentCode() string`

GetRedirectAgentCode returns the RedirectAgentCode field if non-nil, zero value otherwise.

### GetRedirectAgentCodeOk

`func (o *CustconfRedirectExceptions) GetRedirectAgentCodeOk() (*string, bool)`

GetRedirectAgentCodeOk returns a tuple with the RedirectAgentCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectAgentCode

`func (o *CustconfRedirectExceptions) SetRedirectAgentCode(v string)`

SetRedirectAgentCode sets RedirectAgentCode field to given value.

### HasRedirectAgentCode

`func (o *CustconfRedirectExceptions) HasRedirectAgentCode() bool`

HasRedirectAgentCode returns a boolean if a field has been set.

### GetEnabled

`func (o *CustconfRedirectExceptions) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustconfRedirectExceptions) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustconfRedirectExceptions) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CustconfRedirectExceptions) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


