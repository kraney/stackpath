# Monitoringv2Location

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A location&#39;s unique identifier | [optional] 
**Name** | Pointer to **string** | A location&#39;s name | [optional] 
**City** | Pointer to **string** | A location&#39;s city | [optional] 
**CityCode** | Pointer to **string** | A location&#39;s city, expressed as an IATA airport code | [optional] 
**Country** | Pointer to **string** | A location&#39;s country | [optional] 
**CountryCode** | Pointer to **string** | A location&#39;s ISO-3166-1 alpha-2 country code | [optional] 
**Provider** | Pointer to **string** | A location&#39;s network provider | [optional] 
**IpAddresses** | Pointer to **[]string** | The IP addresses of a location  The IP addresses where monitoring checks originate from. | [optional] 
**HasIpv4** | Pointer to **bool** | Whether or not a location supports IPv4 monitoring checks. | [optional] 
**HasIpv6** | Pointer to **bool** | Whether or not a location supports IPv6 monitoring checks. | [optional] 

## Methods

### NewMonitoringv2Location

`func NewMonitoringv2Location() *Monitoringv2Location`

NewMonitoringv2Location instantiates a new Monitoringv2Location object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringv2LocationWithDefaults

`func NewMonitoringv2LocationWithDefaults() *Monitoringv2Location`

NewMonitoringv2LocationWithDefaults instantiates a new Monitoringv2Location object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Monitoringv2Location) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Monitoringv2Location) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Monitoringv2Location) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Monitoringv2Location) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Monitoringv2Location) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Monitoringv2Location) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Monitoringv2Location) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Monitoringv2Location) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCity

`func (o *Monitoringv2Location) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Monitoringv2Location) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Monitoringv2Location) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *Monitoringv2Location) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCityCode

`func (o *Monitoringv2Location) GetCityCode() string`

GetCityCode returns the CityCode field if non-nil, zero value otherwise.

### GetCityCodeOk

`func (o *Monitoringv2Location) GetCityCodeOk() (*string, bool)`

GetCityCodeOk returns a tuple with the CityCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityCode

`func (o *Monitoringv2Location) SetCityCode(v string)`

SetCityCode sets CityCode field to given value.

### HasCityCode

`func (o *Monitoringv2Location) HasCityCode() bool`

HasCityCode returns a boolean if a field has been set.

### GetCountry

`func (o *Monitoringv2Location) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Monitoringv2Location) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Monitoringv2Location) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Monitoringv2Location) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCountryCode

`func (o *Monitoringv2Location) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *Monitoringv2Location) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *Monitoringv2Location) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *Monitoringv2Location) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetProvider

`func (o *Monitoringv2Location) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *Monitoringv2Location) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *Monitoringv2Location) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *Monitoringv2Location) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetIpAddresses

`func (o *Monitoringv2Location) GetIpAddresses() []string`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *Monitoringv2Location) GetIpAddressesOk() (*[]string, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *Monitoringv2Location) SetIpAddresses(v []string)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *Monitoringv2Location) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### GetHasIpv4

`func (o *Monitoringv2Location) GetHasIpv4() bool`

GetHasIpv4 returns the HasIpv4 field if non-nil, zero value otherwise.

### GetHasIpv4Ok

`func (o *Monitoringv2Location) GetHasIpv4Ok() (*bool, bool)`

GetHasIpv4Ok returns a tuple with the HasIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasIpv4

`func (o *Monitoringv2Location) SetHasIpv4(v bool)`

SetHasIpv4 sets HasIpv4 field to given value.

### HasHasIpv4

`func (o *Monitoringv2Location) HasHasIpv4() bool`

HasHasIpv4 returns a boolean if a field has been set.

### GetHasIpv6

`func (o *Monitoringv2Location) GetHasIpv6() bool`

GetHasIpv6 returns the HasIpv6 field if non-nil, zero value otherwise.

### GetHasIpv6Ok

`func (o *Monitoringv2Location) GetHasIpv6Ok() (*bool, bool)`

GetHasIpv6Ok returns a tuple with the HasIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasIpv6

`func (o *Monitoringv2Location) SetHasIpv6(v bool)`

SetHasIpv6 sets HasIpv6 field to given value.

### HasHasIpv6

`func (o *Monitoringv2Location) HasHasIpv6() bool`

HasHasIpv6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


