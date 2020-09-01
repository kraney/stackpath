# Workloadv1Location

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A human readable location name | [optional] 
**City** | Pointer to **string** | A location&#39;s city | [optional] 
**CityCode** | Pointer to **string** | A location&#39;s city, expressed as an IATA airport code | [optional] 
**Subdivision** | Pointer to **string** | A location&#39;s subdivision | [optional] 
**SubdivisionCode** | Pointer to **string** | A location&#39;s subdivision code | [optional] 
**Country** | Pointer to **string** | A location&#39;s country | [optional] 
**CountryCode** | Pointer to **string** | A location&#39;s ISO-3166-1 alpha-2 country code | [optional] 
**Region** | Pointer to **string** | A location&#39;s state or province | [optional] 
**RegionCode** | Pointer to **string** | A location&#39;s ISO-3166-2 region code | [optional] 
**Continent** | Pointer to **string** | A location&#39;s continent | [optional] 
**ContinentCode** | Pointer to **string** | A location&#39;s continent code | [optional] 
**Latitude** | Pointer to **float64** | A location&#39;s geographic latitude | [optional] 
**Longitude** | Pointer to **float64** | A location&#39;s geographic longitude | [optional] 

## Methods

### NewWorkloadv1Location

`func NewWorkloadv1Location() *Workloadv1Location`

NewWorkloadv1Location instantiates a new Workloadv1Location object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadv1LocationWithDefaults

`func NewWorkloadv1LocationWithDefaults() *Workloadv1Location`

NewWorkloadv1LocationWithDefaults instantiates a new Workloadv1Location object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Workloadv1Location) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Workloadv1Location) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Workloadv1Location) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Workloadv1Location) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCity

`func (o *Workloadv1Location) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Workloadv1Location) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Workloadv1Location) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *Workloadv1Location) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCityCode

`func (o *Workloadv1Location) GetCityCode() string`

GetCityCode returns the CityCode field if non-nil, zero value otherwise.

### GetCityCodeOk

`func (o *Workloadv1Location) GetCityCodeOk() (*string, bool)`

GetCityCodeOk returns a tuple with the CityCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityCode

`func (o *Workloadv1Location) SetCityCode(v string)`

SetCityCode sets CityCode field to given value.

### HasCityCode

`func (o *Workloadv1Location) HasCityCode() bool`

HasCityCode returns a boolean if a field has been set.

### GetSubdivision

`func (o *Workloadv1Location) GetSubdivision() string`

GetSubdivision returns the Subdivision field if non-nil, zero value otherwise.

### GetSubdivisionOk

`func (o *Workloadv1Location) GetSubdivisionOk() (*string, bool)`

GetSubdivisionOk returns a tuple with the Subdivision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdivision

`func (o *Workloadv1Location) SetSubdivision(v string)`

SetSubdivision sets Subdivision field to given value.

### HasSubdivision

`func (o *Workloadv1Location) HasSubdivision() bool`

HasSubdivision returns a boolean if a field has been set.

### GetSubdivisionCode

`func (o *Workloadv1Location) GetSubdivisionCode() string`

GetSubdivisionCode returns the SubdivisionCode field if non-nil, zero value otherwise.

### GetSubdivisionCodeOk

`func (o *Workloadv1Location) GetSubdivisionCodeOk() (*string, bool)`

GetSubdivisionCodeOk returns a tuple with the SubdivisionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdivisionCode

`func (o *Workloadv1Location) SetSubdivisionCode(v string)`

SetSubdivisionCode sets SubdivisionCode field to given value.

### HasSubdivisionCode

`func (o *Workloadv1Location) HasSubdivisionCode() bool`

HasSubdivisionCode returns a boolean if a field has been set.

### GetCountry

`func (o *Workloadv1Location) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Workloadv1Location) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Workloadv1Location) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Workloadv1Location) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCountryCode

`func (o *Workloadv1Location) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *Workloadv1Location) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *Workloadv1Location) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *Workloadv1Location) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetRegion

`func (o *Workloadv1Location) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Workloadv1Location) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Workloadv1Location) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Workloadv1Location) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetRegionCode

`func (o *Workloadv1Location) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *Workloadv1Location) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *Workloadv1Location) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *Workloadv1Location) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### GetContinent

`func (o *Workloadv1Location) GetContinent() string`

GetContinent returns the Continent field if non-nil, zero value otherwise.

### GetContinentOk

`func (o *Workloadv1Location) GetContinentOk() (*string, bool)`

GetContinentOk returns a tuple with the Continent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinent

`func (o *Workloadv1Location) SetContinent(v string)`

SetContinent sets Continent field to given value.

### HasContinent

`func (o *Workloadv1Location) HasContinent() bool`

HasContinent returns a boolean if a field has been set.

### GetContinentCode

`func (o *Workloadv1Location) GetContinentCode() string`

GetContinentCode returns the ContinentCode field if non-nil, zero value otherwise.

### GetContinentCodeOk

`func (o *Workloadv1Location) GetContinentCodeOk() (*string, bool)`

GetContinentCodeOk returns a tuple with the ContinentCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinentCode

`func (o *Workloadv1Location) SetContinentCode(v string)`

SetContinentCode sets ContinentCode field to given value.

### HasContinentCode

`func (o *Workloadv1Location) HasContinentCode() bool`

HasContinentCode returns a boolean if a field has been set.

### GetLatitude

`func (o *Workloadv1Location) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *Workloadv1Location) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *Workloadv1Location) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *Workloadv1Location) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *Workloadv1Location) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *Workloadv1Location) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *Workloadv1Location) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *Workloadv1Location) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


