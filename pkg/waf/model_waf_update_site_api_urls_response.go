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

// WafUpdateSiteApiUrlsResponse A response from a request to update an WAF site's API URLs
type WafUpdateSiteApiUrlsResponse struct {
	// A list of API URLs that will be processed differently by the WAF
	ApiUrls *[]string `json:"apiUrls,omitempty"`
}

// NewWafUpdateSiteApiUrlsResponse instantiates a new WafUpdateSiteApiUrlsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafUpdateSiteApiUrlsResponse() *WafUpdateSiteApiUrlsResponse {
	this := WafUpdateSiteApiUrlsResponse{}
	return &this
}

// NewWafUpdateSiteApiUrlsResponseWithDefaults instantiates a new WafUpdateSiteApiUrlsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafUpdateSiteApiUrlsResponseWithDefaults() *WafUpdateSiteApiUrlsResponse {
	this := WafUpdateSiteApiUrlsResponse{}
	return &this
}

// GetApiUrls returns the ApiUrls field value if set, zero value otherwise.
func (o *WafUpdateSiteApiUrlsResponse) GetApiUrls() []string {
	if o == nil || o.ApiUrls == nil {
		var ret []string
		return ret
	}
	return *o.ApiUrls
}

// GetApiUrlsOk returns a tuple with the ApiUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafUpdateSiteApiUrlsResponse) GetApiUrlsOk() (*[]string, bool) {
	if o == nil || o.ApiUrls == nil {
		return nil, false
	}
	return o.ApiUrls, true
}

// HasApiUrls returns a boolean if a field has been set.
func (o *WafUpdateSiteApiUrlsResponse) HasApiUrls() bool {
	if o != nil && o.ApiUrls != nil {
		return true
	}

	return false
}

// SetApiUrls gets a reference to the given []string and assigns it to the ApiUrls field.
func (o *WafUpdateSiteApiUrlsResponse) SetApiUrls(v []string) {
	o.ApiUrls = &v
}

func (o WafUpdateSiteApiUrlsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiUrls != nil {
		toSerialize["apiUrls"] = o.ApiUrls
	}
	return json.Marshal(toSerialize)
}

type NullableWafUpdateSiteApiUrlsResponse struct {
	value *WafUpdateSiteApiUrlsResponse
	isSet bool
}

func (v NullableWafUpdateSiteApiUrlsResponse) Get() *WafUpdateSiteApiUrlsResponse {
	return v.value
}

func (v *NullableWafUpdateSiteApiUrlsResponse) Set(val *WafUpdateSiteApiUrlsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWafUpdateSiteApiUrlsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWafUpdateSiteApiUrlsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWafUpdateSiteApiUrlsResponse(val *WafUpdateSiteApiUrlsResponse) *NullableWafUpdateSiteApiUrlsResponse {
	return &NullableWafUpdateSiteApiUrlsResponse{value: val, isSet: true}
}

func (v NullableWafUpdateSiteApiUrlsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWafUpdateSiteApiUrlsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}