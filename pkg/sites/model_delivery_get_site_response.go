/*
 * Sites
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package sites

import (
	"encoding/json"
)

// DeliveryGetSiteResponse The response from a request to retrieve a single site
type DeliveryGetSiteResponse struct {
	Site *DeliverySite `json:"site,omitempty"`
}

// NewDeliveryGetSiteResponse instantiates a new DeliveryGetSiteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryGetSiteResponse() *DeliveryGetSiteResponse {
	this := DeliveryGetSiteResponse{}
	return &this
}

// NewDeliveryGetSiteResponseWithDefaults instantiates a new DeliveryGetSiteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryGetSiteResponseWithDefaults() *DeliveryGetSiteResponse {
	this := DeliveryGetSiteResponse{}
	return &this
}

// GetSite returns the Site field value if set, zero value otherwise.
func (o *DeliveryGetSiteResponse) GetSite() DeliverySite {
	if o == nil || o.Site == nil {
		var ret DeliverySite
		return ret
	}
	return *o.Site
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryGetSiteResponse) GetSiteOk() (*DeliverySite, bool) {
	if o == nil || o.Site == nil {
		return nil, false
	}
	return o.Site, true
}

// HasSite returns a boolean if a field has been set.
func (o *DeliveryGetSiteResponse) HasSite() bool {
	if o != nil && o.Site != nil {
		return true
	}

	return false
}

// SetSite gets a reference to the given DeliverySite and assigns it to the Site field.
func (o *DeliveryGetSiteResponse) SetSite(v DeliverySite) {
	o.Site = &v
}

func (o DeliveryGetSiteResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Site != nil {
		toSerialize["site"] = o.Site
	}
	return json.Marshal(toSerialize)
}

type NullableDeliveryGetSiteResponse struct {
	value *DeliveryGetSiteResponse
	isSet bool
}

func (v NullableDeliveryGetSiteResponse) Get() *DeliveryGetSiteResponse {
	return v.value
}

func (v *NullableDeliveryGetSiteResponse) Set(val *DeliveryGetSiteResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryGetSiteResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryGetSiteResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryGetSiteResponse(val *DeliveryGetSiteResponse) *NullableDeliveryGetSiteResponse {
	return &NullableDeliveryGetSiteResponse{value: val, isSet: true}
}

func (v NullableDeliveryGetSiteResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryGetSiteResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}