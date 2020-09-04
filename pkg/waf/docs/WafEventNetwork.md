# WafEventNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** | The originating IP address | [optional] 
**Country** | Pointer to **string** | The long name of the country where the request originated from | [optional] 
**CountryCode** | Pointer to **string** | The ISO 3166-1 alpha-2 code of the country where the request originated from | [optional] 
**Organization** | Pointer to **string** | The organization that owns the originating IP address according to a WHOIS lookup | [optional] 

## Methods

### NewWafEventNetwork

`func NewWafEventNetwork() *WafEventNetwork`

NewWafEventNetwork instantiates a new WafEventNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafEventNetworkWithDefaults

`func NewWafEventNetworkWithDefaults() *WafEventNetwork`

NewWafEventNetworkWithDefaults instantiates a new WafEventNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *WafEventNetwork) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *WafEventNetwork) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *WafEventNetwork) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *WafEventNetwork) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetCountry

`func (o *WafEventNetwork) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *WafEventNetwork) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *WafEventNetwork) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *WafEventNetwork) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCountryCode

`func (o *WafEventNetwork) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *WafEventNetwork) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *WafEventNetwork) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *WafEventNetwork) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetOrganization

`func (o *WafEventNetwork) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *WafEventNetwork) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *WafEventNetwork) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *WafEventNetwork) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


