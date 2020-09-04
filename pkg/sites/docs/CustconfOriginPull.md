# CustconfOriginPull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | This is used by the API to perform conflict checking | [optional] 
**RedirectAction** | Pointer to [**OriginPullRedirectActionEnumWrapperValue**](OriginPullRedirectActionEnumWrapperValue.md) |  | [optional] [default to "UNKNOWN"]
**NoQSParams** | Pointer to **bool** | GFS sends a path without any query string parameters when making external origin requests regardless if any parameters were sent by the User-Agent. | [optional] 
**RetryMethods** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. List of HTTP Methods that define types of origin pull requests that can be retried if a failure occurs after sending a previous request. | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewCustconfOriginPull

`func NewCustconfOriginPull() *CustconfOriginPull`

NewCustconfOriginPull instantiates a new CustconfOriginPull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustconfOriginPullWithDefaults

`func NewCustconfOriginPullWithDefaults() *CustconfOriginPull`

NewCustconfOriginPullWithDefaults instantiates a new CustconfOriginPull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustconfOriginPull) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustconfOriginPull) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustconfOriginPull) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustconfOriginPull) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRedirectAction

`func (o *CustconfOriginPull) GetRedirectAction() OriginPullRedirectActionEnumWrapperValue`

GetRedirectAction returns the RedirectAction field if non-nil, zero value otherwise.

### GetRedirectActionOk

`func (o *CustconfOriginPull) GetRedirectActionOk() (*OriginPullRedirectActionEnumWrapperValue, bool)`

GetRedirectActionOk returns a tuple with the RedirectAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectAction

`func (o *CustconfOriginPull) SetRedirectAction(v OriginPullRedirectActionEnumWrapperValue)`

SetRedirectAction sets RedirectAction field to given value.

### HasRedirectAction

`func (o *CustconfOriginPull) HasRedirectAction() bool`

HasRedirectAction returns a boolean if a field has been set.

### GetNoQSParams

`func (o *CustconfOriginPull) GetNoQSParams() bool`

GetNoQSParams returns the NoQSParams field if non-nil, zero value otherwise.

### GetNoQSParamsOk

`func (o *CustconfOriginPull) GetNoQSParamsOk() (*bool, bool)`

GetNoQSParamsOk returns a tuple with the NoQSParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoQSParams

`func (o *CustconfOriginPull) SetNoQSParams(v bool)`

SetNoQSParams sets NoQSParams field to given value.

### HasNoQSParams

`func (o *CustconfOriginPull) HasNoQSParams() bool`

HasNoQSParams returns a boolean if a field has been set.

### GetRetryMethods

`func (o *CustconfOriginPull) GetRetryMethods() string`

GetRetryMethods returns the RetryMethods field if non-nil, zero value otherwise.

### GetRetryMethodsOk

`func (o *CustconfOriginPull) GetRetryMethodsOk() (*string, bool)`

GetRetryMethodsOk returns a tuple with the RetryMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryMethods

`func (o *CustconfOriginPull) SetRetryMethods(v string)`

SetRetryMethods sets RetryMethods field to given value.

### HasRetryMethods

`func (o *CustconfOriginPull) HasRetryMethods() bool`

HasRetryMethods returns a boolean if a field has been set.

### GetEnabled

`func (o *CustconfOriginPull) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustconfOriginPull) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustconfOriginPull) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CustconfOriginPull) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


