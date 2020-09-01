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

// CustconfCustomHeader The X-Forwarded-For header policy allows you to override the name of the X-Forwarded-For header the CDN sends to your origin.
type CustconfCustomHeader struct {
	// This is used by the API to perform conflict checking
	Id *string `json:"id,omitempty"`
	// This is the name of the X-Forwarded-For header you want the CDN to use when making requests to your basic authorization server.
	XForwardedForAuth *string `json:"xForwardedForAuth,omitempty"`
	// This is the name of the X-Forwarded-For header you want the CDN to use when making requests to your origin server.
	XForwardedForOrigin *string `json:"xForwardedForOrigin,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
}

// NewCustconfCustomHeader instantiates a new CustconfCustomHeader object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustconfCustomHeader() *CustconfCustomHeader {
	this := CustconfCustomHeader{}
	return &this
}

// NewCustconfCustomHeaderWithDefaults instantiates a new CustconfCustomHeader object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustconfCustomHeaderWithDefaults() *CustconfCustomHeader {
	this := CustconfCustomHeader{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustconfCustomHeader) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustconfCustomHeader) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustconfCustomHeader) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CustconfCustomHeader) SetId(v string) {
	o.Id = &v
}

// GetXForwardedForAuth returns the XForwardedForAuth field value if set, zero value otherwise.
func (o *CustconfCustomHeader) GetXForwardedForAuth() string {
	if o == nil || o.XForwardedForAuth == nil {
		var ret string
		return ret
	}
	return *o.XForwardedForAuth
}

// GetXForwardedForAuthOk returns a tuple with the XForwardedForAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustconfCustomHeader) GetXForwardedForAuthOk() (*string, bool) {
	if o == nil || o.XForwardedForAuth == nil {
		return nil, false
	}
	return o.XForwardedForAuth, true
}

// HasXForwardedForAuth returns a boolean if a field has been set.
func (o *CustconfCustomHeader) HasXForwardedForAuth() bool {
	if o != nil && o.XForwardedForAuth != nil {
		return true
	}

	return false
}

// SetXForwardedForAuth gets a reference to the given string and assigns it to the XForwardedForAuth field.
func (o *CustconfCustomHeader) SetXForwardedForAuth(v string) {
	o.XForwardedForAuth = &v
}

// GetXForwardedForOrigin returns the XForwardedForOrigin field value if set, zero value otherwise.
func (o *CustconfCustomHeader) GetXForwardedForOrigin() string {
	if o == nil || o.XForwardedForOrigin == nil {
		var ret string
		return ret
	}
	return *o.XForwardedForOrigin
}

// GetXForwardedForOriginOk returns a tuple with the XForwardedForOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustconfCustomHeader) GetXForwardedForOriginOk() (*string, bool) {
	if o == nil || o.XForwardedForOrigin == nil {
		return nil, false
	}
	return o.XForwardedForOrigin, true
}

// HasXForwardedForOrigin returns a boolean if a field has been set.
func (o *CustconfCustomHeader) HasXForwardedForOrigin() bool {
	if o != nil && o.XForwardedForOrigin != nil {
		return true
	}

	return false
}

// SetXForwardedForOrigin gets a reference to the given string and assigns it to the XForwardedForOrigin field.
func (o *CustconfCustomHeader) SetXForwardedForOrigin(v string) {
	o.XForwardedForOrigin = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CustconfCustomHeader) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustconfCustomHeader) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CustconfCustomHeader) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CustconfCustomHeader) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o CustconfCustomHeader) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.XForwardedForAuth != nil {
		toSerialize["xForwardedForAuth"] = o.XForwardedForAuth
	}
	if o.XForwardedForOrigin != nil {
		toSerialize["xForwardedForOrigin"] = o.XForwardedForOrigin
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableCustconfCustomHeader struct {
	value *CustconfCustomHeader
	isSet bool
}

func (v NullableCustconfCustomHeader) Get() *CustconfCustomHeader {
	return v.value
}

func (v *NullableCustconfCustomHeader) Set(val *CustconfCustomHeader) {
	v.value = val
	v.isSet = true
}

func (v NullableCustconfCustomHeader) IsSet() bool {
	return v.isSet
}

func (v *NullableCustconfCustomHeader) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustconfCustomHeader(val *CustconfCustomHeader) *NullableCustconfCustomHeader {
	return &NullableCustconfCustomHeader{value: val, isSet: true}
}

func (v NullableCustconfCustomHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustconfCustomHeader) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}