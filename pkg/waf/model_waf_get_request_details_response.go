/*
 * Web Application Firewall
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package waf

import (
	"encoding/json"
)

// WafGetRequestDetailsResponse A response from a request to retrieve all available information about a request to a WAF site
type WafGetRequestDetailsResponse struct {
	RequestDetails *WafRequestDetails `json:"requestDetails,omitempty"`
}

// NewWafGetRequestDetailsResponse instantiates a new WafGetRequestDetailsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafGetRequestDetailsResponse() *WafGetRequestDetailsResponse {
	this := WafGetRequestDetailsResponse{}
	return &this
}

// NewWafGetRequestDetailsResponseWithDefaults instantiates a new WafGetRequestDetailsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafGetRequestDetailsResponseWithDefaults() *WafGetRequestDetailsResponse {
	this := WafGetRequestDetailsResponse{}
	return &this
}

// GetRequestDetails returns the RequestDetails field value if set, zero value otherwise.
func (o *WafGetRequestDetailsResponse) GetRequestDetails() WafRequestDetails {
	if o == nil || o.RequestDetails == nil {
		var ret WafRequestDetails
		return ret
	}
	return *o.RequestDetails
}

// GetRequestDetailsOk returns a tuple with the RequestDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafGetRequestDetailsResponse) GetRequestDetailsOk() (*WafRequestDetails, bool) {
	if o == nil || o.RequestDetails == nil {
		return nil, false
	}
	return o.RequestDetails, true
}

// HasRequestDetails returns a boolean if a field has been set.
func (o *WafGetRequestDetailsResponse) HasRequestDetails() bool {
	if o != nil && o.RequestDetails != nil {
		return true
	}

	return false
}

// SetRequestDetails gets a reference to the given WafRequestDetails and assigns it to the RequestDetails field.
func (o *WafGetRequestDetailsResponse) SetRequestDetails(v WafRequestDetails) {
	o.RequestDetails = &v
}

func (o WafGetRequestDetailsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RequestDetails != nil {
		toSerialize["requestDetails"] = o.RequestDetails
	}
	return json.Marshal(toSerialize)
}

type NullableWafGetRequestDetailsResponse struct {
	value *WafGetRequestDetailsResponse
	isSet bool
}

func (v NullableWafGetRequestDetailsResponse) Get() *WafGetRequestDetailsResponse {
	return v.value
}

func (v *NullableWafGetRequestDetailsResponse) Set(val *WafGetRequestDetailsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWafGetRequestDetailsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWafGetRequestDetailsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWafGetRequestDetailsResponse(val *WafGetRequestDetailsResponse) *NullableWafGetRequestDetailsResponse {
	return &NullableWafGetRequestDetailsResponse{value: val, isSet: true}
}

func (v NullableWafGetRequestDetailsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWafGetRequestDetailsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
