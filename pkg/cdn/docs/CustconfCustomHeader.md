# CustconfCustomHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | This is used by the API to perform conflict checking | [optional] 
**XForwardedForAuth** | Pointer to **string** | This is the name of the X-Forwarded-For header you want the CDN to use when making requests to your basic authorization server. | [optional] 
**XForwardedForOrigin** | Pointer to **string** | This is the name of the X-Forwarded-For header you want the CDN to use when making requests to your origin server. | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewCustconfCustomHeader

`func NewCustconfCustomHeader() *CustconfCustomHeader`

NewCustconfCustomHeader instantiates a new CustconfCustomHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustconfCustomHeaderWithDefaults

`func NewCustconfCustomHeaderWithDefaults() *CustconfCustomHeader`

NewCustconfCustomHeaderWithDefaults instantiates a new CustconfCustomHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustconfCustomHeader) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustconfCustomHeader) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustconfCustomHeader) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustconfCustomHeader) HasId() bool`

HasId returns a boolean if a field has been set.

### GetXForwardedForAuth

`func (o *CustconfCustomHeader) GetXForwardedForAuth() string`

GetXForwardedForAuth returns the XForwardedForAuth field if non-nil, zero value otherwise.

### GetXForwardedForAuthOk

`func (o *CustconfCustomHeader) GetXForwardedForAuthOk() (*string, bool)`

GetXForwardedForAuthOk returns a tuple with the XForwardedForAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXForwardedForAuth

`func (o *CustconfCustomHeader) SetXForwardedForAuth(v string)`

SetXForwardedForAuth sets XForwardedForAuth field to given value.

### HasXForwardedForAuth

`func (o *CustconfCustomHeader) HasXForwardedForAuth() bool`

HasXForwardedForAuth returns a boolean if a field has been set.

### GetXForwardedForOrigin

`func (o *CustconfCustomHeader) GetXForwardedForOrigin() string`

GetXForwardedForOrigin returns the XForwardedForOrigin field if non-nil, zero value otherwise.

### GetXForwardedForOriginOk

`func (o *CustconfCustomHeader) GetXForwardedForOriginOk() (*string, bool)`

GetXForwardedForOriginOk returns a tuple with the XForwardedForOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXForwardedForOrigin

`func (o *CustconfCustomHeader) SetXForwardedForOrigin(v string)`

SetXForwardedForOrigin sets XForwardedForOrigin field to given value.

### HasXForwardedForOrigin

`func (o *CustconfCustomHeader) HasXForwardedForOrigin() bool`

HasXForwardedForOrigin returns a boolean if a field has been set.

### GetEnabled

`func (o *CustconfCustomHeader) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustconfCustomHeader) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustconfCustomHeader) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CustconfCustomHeader) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


