/*
 * Content Delivery Network
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package cdn

import (
	"encoding/json"
)

// CdnGetSiteDnsTargetsResponse The response from a request to retrieve a CDN site's DNS CNAME targets
type CdnGetSiteDnsTargetsResponse struct {
	// The requested DNS CNAME targets  A site's hostname should point to these CNAME targets in order for traffic to be sent through StackPath's edge nodes.
	Addresses *[]string `json:"addresses,omitempty"`
}

// NewCdnGetSiteDnsTargetsResponse instantiates a new CdnGetSiteDnsTargetsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCdnGetSiteDnsTargetsResponse() *CdnGetSiteDnsTargetsResponse {
	this := CdnGetSiteDnsTargetsResponse{}
	return &this
}

// NewCdnGetSiteDnsTargetsResponseWithDefaults instantiates a new CdnGetSiteDnsTargetsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCdnGetSiteDnsTargetsResponseWithDefaults() *CdnGetSiteDnsTargetsResponse {
	this := CdnGetSiteDnsTargetsResponse{}
	return &this
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *CdnGetSiteDnsTargetsResponse) GetAddresses() []string {
	if o == nil || o.Addresses == nil {
		var ret []string
		return ret
	}
	return *o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdnGetSiteDnsTargetsResponse) GetAddressesOk() (*[]string, bool) {
	if o == nil || o.Addresses == nil {
		return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *CdnGetSiteDnsTargetsResponse) HasAddresses() bool {
	if o != nil && o.Addresses != nil {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given []string and assigns it to the Addresses field.
func (o *CdnGetSiteDnsTargetsResponse) SetAddresses(v []string) {
	o.Addresses = &v
}

func (o CdnGetSiteDnsTargetsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Addresses != nil {
		toSerialize["addresses"] = o.Addresses
	}
	return json.Marshal(toSerialize)
}

type NullableCdnGetSiteDnsTargetsResponse struct {
	value *CdnGetSiteDnsTargetsResponse
	isSet bool
}

func (v NullableCdnGetSiteDnsTargetsResponse) Get() *CdnGetSiteDnsTargetsResponse {
	return v.value
}

func (v *NullableCdnGetSiteDnsTargetsResponse) Set(val *CdnGetSiteDnsTargetsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCdnGetSiteDnsTargetsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCdnGetSiteDnsTargetsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCdnGetSiteDnsTargetsResponse(val *CdnGetSiteDnsTargetsResponse) *NullableCdnGetSiteDnsTargetsResponse {
	return &NullableCdnGetSiteDnsTargetsResponse{value: val, isSet: true}
}

func (v NullableCdnGetSiteDnsTargetsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCdnGetSiteDnsTargetsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}