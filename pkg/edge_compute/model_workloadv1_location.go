/*
 * Edge Compute
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package edge_compute

import (
	"encoding/json"
)

// Workloadv1Location Geographic location information
type Workloadv1Location struct {
	// A human readable location name
	Name *string `json:"name,omitempty"`
	// A location's city
	City *string `json:"city,omitempty"`
	// A location's city, expressed as an IATA airport code
	CityCode *string `json:"cityCode,omitempty"`
	// A location's subdivision
	Subdivision *string `json:"subdivision,omitempty"`
	// A location's subdivision code
	SubdivisionCode *string `json:"subdivisionCode,omitempty"`
	// A location's country
	Country *string `json:"country,omitempty"`
	// A location's ISO-3166-1 alpha-2 country code
	CountryCode *string `json:"countryCode,omitempty"`
	// A location's state or province
	Region *string `json:"region,omitempty"`
	// A location's ISO-3166-2 region code
	RegionCode *string `json:"regionCode,omitempty"`
	// A location's continent
	Continent *string `json:"continent,omitempty"`
	// A location's continent code
	ContinentCode *string `json:"continentCode,omitempty"`
	// A location's geographic latitude
	Latitude *float64 `json:"latitude,omitempty"`
	// A location's geographic longitude
	Longitude *float64 `json:"longitude,omitempty"`
}

// NewWorkloadv1Location instantiates a new Workloadv1Location object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkloadv1Location() *Workloadv1Location {
	this := Workloadv1Location{}
	return &this
}

// NewWorkloadv1LocationWithDefaults instantiates a new Workloadv1Location object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkloadv1LocationWithDefaults() *Workloadv1Location {
	this := Workloadv1Location{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Workloadv1Location) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workloadv1Location) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Workloadv1Location) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Workloadv1Location) SetName(v string) {
	o.Name = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *Workloadv1Location) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workloadv1Location) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *Workloadv1Location) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *Workloadv1Location) SetCity(v string) {
	o.City = &v
}

// GetCityCode returns the CityCode field value if set, zero value otherwise.
func (o *Workloadv1Location) GetCityCode() string {
	if o == nil || o.CityCode == nil {
		var ret string
		return ret
	}
	return *o.CityCode
}

// GetCityCodeOk returns a tuple with the CityCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workloadv1Location) GetCityCodeOk() (*string, bool) {
	if o == nil || o.CityCode == nil {
		return nil, false
	}
	return o.CityCode, true
}

// HasCityCode returns a boolean if a field has been set.
func (o *Workloadv1Location) HasCityCode() bool {
	if o != nil && o.CityCode != nil {
		return true
	}

	return false
}

// SetCityCode gets a reference to the given string and assigns it to the CityCode field.
func (o *Workloadv1Location) SetCityCode(v string) {
	o.CityCode = &v
}

// GetSubdivision returns the Subdivision field value if set, zero value otherwise.
func (o *Workloadv1Location) GetSubdivision() string {
	if o == nil || o.Subdivision == nil {
		var ret string
		return ret
	}
	return *o.Subdivision
}

// GetSubdivisionOk returns a tuple with the Subdivision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workloadv1Location) GetSubdivisionOk() (*string, bool) {
	if o == nil || o.Subdivision == nil {
		return nil, false
	}
	return o.Subdivision, true
}

// HasSubdivision returns a boolean if a field has been set.
func (o *Workloadv1Location) HasSubdivision() bool {
	if o != nil && o.Subdivision != nil {
		return true
	}

	return false
}

// SetSubdivision gets a reference to the given string and assigns it to the Subdivision field.
func (o *Workloadv1Location) SetSubdivision(v string) {
	o.Subdivision = &v
}

// GetSubdivisionCode returns the SubdivisionCode field value if set, zero value otherwise.
func (o *Workloadv1Location) GetSubdivisionCode() string {
	if o == nil || o.SubdivisionCode == nil {
		var ret string
		return ret
	}
	return *o.SubdivisionCode
}

// GetSubdivisionCodeOk returns a tuple with the SubdivisionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workloadv1Location) GetSubdivisionCodeOk() (*string, bool) {
	if o == nil || o.SubdivisionCode == nil {
		return nil, false
	}
	return o.SubdivisionCode, true
}

// HasSubdivisionCode returns a boolean if a field has been set.
func (o *Workloadv1Location) HasSubdivisionCode() bool {
	if o != nil && o.SubdivisionCode != nil {
		return true
	}

	return false
}

// SetSubdivisionCode gets a reference to the given string and assigns it to the SubdivisionCode field.
func (o *Workloadv1Location) SetSubdivisionCode(v string) {
	o.SubdivisionCode = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *Workloadv1Location) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workloadv1Location) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *Workloadv1Location) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *Workloadv1Location) SetCountry(v string) {
	o.Country = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *Workloadv1Location) GetCountryCode() string {
	if o == nil || o.CountryCode == nil {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workloadv1Location) GetCountryCodeOk() (*string, bool) {
	if o == nil || o.CountryCode == nil {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *Workloadv1Location) HasCountryCode() bool {
	if o != nil && o.CountryCode != nil {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *Workloadv1Location) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *Workloadv1Location) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workloadv1Location) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *Workloadv1Location) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *Workloadv1Location) SetRegion(v string) {
	o.Region = &v
}

// GetRegionCode returns the RegionCode field value if set, zero value otherwise.
func (o *Workloadv1Location) GetRegionCode() string {
	if o == nil || o.RegionCode == nil {
		var ret string
		return ret
	}
	return *o.RegionCode
}

// GetRegionCodeOk returns a tuple with the RegionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workloadv1Location) GetRegionCodeOk() (*string, bool) {
	if o == nil || o.RegionCode == nil {
		return nil, false
	}
	return o.RegionCode, true
}

// HasRegionCode returns a boolean if a field has been set.
func (o *Workloadv1Location) HasRegionCode() bool {
	if o != nil && o.RegionCode != nil {
		return true
	}

	return false
}

// SetRegionCode gets a reference to the given string and assigns it to the RegionCode field.
func (o *Workloadv1Location) SetRegionCode(v string) {
	o.RegionCode = &v
}

// GetContinent returns the Continent field value if set, zero value otherwise.
func (o *Workloadv1Location) GetContinent() string {
	if o == nil || o.Continent == nil {
		var ret string
		return ret
	}
	return *o.Continent
}

// GetContinentOk returns a tuple with the Continent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workloadv1Location) GetContinentOk() (*string, bool) {
	if o == nil || o.Continent == nil {
		return nil, false
	}
	return o.Continent, true
}

// HasContinent returns a boolean if a field has been set.
func (o *Workloadv1Location) HasContinent() bool {
	if o != nil && o.Continent != nil {
		return true
	}

	return false
}

// SetContinent gets a reference to the given string and assigns it to the Continent field.
func (o *Workloadv1Location) SetContinent(v string) {
	o.Continent = &v
}

// GetContinentCode returns the ContinentCode field value if set, zero value otherwise.
func (o *Workloadv1Location) GetContinentCode() string {
	if o == nil || o.ContinentCode == nil {
		var ret string
		return ret
	}
	return *o.ContinentCode
}

// GetContinentCodeOk returns a tuple with the ContinentCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workloadv1Location) GetContinentCodeOk() (*string, bool) {
	if o == nil || o.ContinentCode == nil {
		return nil, false
	}
	return o.ContinentCode, true
}

// HasContinentCode returns a boolean if a field has been set.
func (o *Workloadv1Location) HasContinentCode() bool {
	if o != nil && o.ContinentCode != nil {
		return true
	}

	return false
}

// SetContinentCode gets a reference to the given string and assigns it to the ContinentCode field.
func (o *Workloadv1Location) SetContinentCode(v string) {
	o.ContinentCode = &v
}

// GetLatitude returns the Latitude field value if set, zero value otherwise.
func (o *Workloadv1Location) GetLatitude() float64 {
	if o == nil || o.Latitude == nil {
		var ret float64
		return ret
	}
	return *o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workloadv1Location) GetLatitudeOk() (*float64, bool) {
	if o == nil || o.Latitude == nil {
		return nil, false
	}
	return o.Latitude, true
}

// HasLatitude returns a boolean if a field has been set.
func (o *Workloadv1Location) HasLatitude() bool {
	if o != nil && o.Latitude != nil {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given float64 and assigns it to the Latitude field.
func (o *Workloadv1Location) SetLatitude(v float64) {
	o.Latitude = &v
}

// GetLongitude returns the Longitude field value if set, zero value otherwise.
func (o *Workloadv1Location) GetLongitude() float64 {
	if o == nil || o.Longitude == nil {
		var ret float64
		return ret
	}
	return *o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workloadv1Location) GetLongitudeOk() (*float64, bool) {
	if o == nil || o.Longitude == nil {
		return nil, false
	}
	return o.Longitude, true
}

// HasLongitude returns a boolean if a field has been set.
func (o *Workloadv1Location) HasLongitude() bool {
	if o != nil && o.Longitude != nil {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given float64 and assigns it to the Longitude field.
func (o *Workloadv1Location) SetLongitude(v float64) {
	o.Longitude = &v
}

func (o Workloadv1Location) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.City != nil {
		toSerialize["city"] = o.City
	}
	if o.CityCode != nil {
		toSerialize["cityCode"] = o.CityCode
	}
	if o.Subdivision != nil {
		toSerialize["subdivision"] = o.Subdivision
	}
	if o.SubdivisionCode != nil {
		toSerialize["subdivisionCode"] = o.SubdivisionCode
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	if o.CountryCode != nil {
		toSerialize["countryCode"] = o.CountryCode
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.RegionCode != nil {
		toSerialize["regionCode"] = o.RegionCode
	}
	if o.Continent != nil {
		toSerialize["continent"] = o.Continent
	}
	if o.ContinentCode != nil {
		toSerialize["continentCode"] = o.ContinentCode
	}
	if o.Latitude != nil {
		toSerialize["latitude"] = o.Latitude
	}
	if o.Longitude != nil {
		toSerialize["longitude"] = o.Longitude
	}
	return json.Marshal(toSerialize)
}

type NullableWorkloadv1Location struct {
	value *Workloadv1Location
	isSet bool
}

func (v NullableWorkloadv1Location) Get() *Workloadv1Location {
	return v.value
}

func (v *NullableWorkloadv1Location) Set(val *Workloadv1Location) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkloadv1Location) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkloadv1Location) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkloadv1Location(val *Workloadv1Location) *NullableWorkloadv1Location {
	return &NullableWorkloadv1Location{value: val, isSet: true}
}

func (v NullableWorkloadv1Location) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkloadv1Location) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
