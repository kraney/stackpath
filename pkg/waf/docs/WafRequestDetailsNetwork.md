# WafRequestDetailsNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientIp** | Pointer to **string** | The originating IP address | [optional] 
**Country** | Pointer to **string** | The ISO 3166-1 alpha-2 code of the country where the request originated from | [optional] 
**Organization** | Pointer to [**NetworkOrganization**](NetworkOrganization.md) |  | [optional] 

## Methods

### NewWafRequestDetailsNetwork

`func NewWafRequestDetailsNetwork() *WafRequestDetailsNetwork`

NewWafRequestDetailsNetwork instantiates a new WafRequestDetailsNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafRequestDetailsNetworkWithDefaults

`func NewWafRequestDetailsNetworkWithDefaults() *WafRequestDetailsNetwork`

NewWafRequestDetailsNetworkWithDefaults instantiates a new WafRequestDetailsNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientIp

`func (o *WafRequestDetailsNetwork) GetClientIp() string`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *WafRequestDetailsNetwork) GetClientIpOk() (*string, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *WafRequestDetailsNetwork) SetClientIp(v string)`

SetClientIp sets ClientIp field to given value.

### HasClientIp

`func (o *WafRequestDetailsNetwork) HasClientIp() bool`

HasClientIp returns a boolean if a field has been set.

### GetCountry

`func (o *WafRequestDetailsNetwork) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *WafRequestDetailsNetwork) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *WafRequestDetailsNetwork) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *WafRequestDetailsNetwork) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetOrganization

`func (o *WafRequestDetailsNetwork) GetOrganization() NetworkOrganization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *WafRequestDetailsNetwork) GetOrganizationOk() (*NetworkOrganization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *WafRequestDetailsNetwork) SetOrganization(v NetworkOrganization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *WafRequestDetailsNetwork) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


