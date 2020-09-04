# CustconfHttpMethods

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | This is used by the API to perform conflict checking | [optional] 
**PassThru** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. This is a comma separated list of HTTP methods you want the CDN to proxy to your origin. A wildcard can be entered to include all methods (excluding HEAD and GET). | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewCustconfHttpMethods

`func NewCustconfHttpMethods() *CustconfHttpMethods`

NewCustconfHttpMethods instantiates a new CustconfHttpMethods object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustconfHttpMethodsWithDefaults

`func NewCustconfHttpMethodsWithDefaults() *CustconfHttpMethods`

NewCustconfHttpMethodsWithDefaults instantiates a new CustconfHttpMethods object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustconfHttpMethods) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustconfHttpMethods) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustconfHttpMethods) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustconfHttpMethods) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPassThru

`func (o *CustconfHttpMethods) GetPassThru() string`

GetPassThru returns the PassThru field if non-nil, zero value otherwise.

### GetPassThruOk

`func (o *CustconfHttpMethods) GetPassThruOk() (*string, bool)`

GetPassThruOk returns a tuple with the PassThru field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassThru

`func (o *CustconfHttpMethods) SetPassThru(v string)`

SetPassThru sets PassThru field to given value.

### HasPassThru

`func (o *CustconfHttpMethods) HasPassThru() bool`

HasPassThru returns a boolean if a field has been set.

### GetEnabled

`func (o *CustconfHttpMethods) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustconfHttpMethods) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustconfHttpMethods) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CustconfHttpMethods) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


