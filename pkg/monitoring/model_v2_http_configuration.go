/*
 * Monitoring
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package monitoring

import (
	"encoding/json"
)

// V2HttpConfiguration HTTP monitor configuration
type V2HttpConfiguration struct {
	// The URL of the service to monitor  The URL must begin with a case insensitive HTTP scheme of 'http' or 'https'.
	Url *string `json:"url,omitempty"`
	// The HTTP method used when the monitor makes a request to the service.
	Method *string `json:"method,omitempty"`
	// A base64 encoded HTTP request body to use when a monitor checks the service.
	Body *string `json:"body,omitempty"`
	// A list of HTTP headers to add to the request when a monitor checks a service.
	Headers *[]V2Header `json:"headers,omitempty"`
	Basic *V2HttpConfigurationBasicAuth `json:"basic,omitempty"`
	Digest *V2HttpConfigurationDigestAuth `json:"digest,omitempty"`
	ClientCertificate *V2HttpConfigurationClientCertificate `json:"clientCertificate,omitempty"`
	// Whether or not to validate a service's certificate.
	ValidateCertificate *bool `json:"validateCertificate,omitempty"`
}

// NewV2HttpConfiguration instantiates a new V2HttpConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2HttpConfiguration() *V2HttpConfiguration {
	this := V2HttpConfiguration{}
	return &this
}

// NewV2HttpConfigurationWithDefaults instantiates a new V2HttpConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2HttpConfigurationWithDefaults() *V2HttpConfiguration {
	this := V2HttpConfiguration{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *V2HttpConfiguration) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2HttpConfiguration) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *V2HttpConfiguration) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *V2HttpConfiguration) SetUrl(v string) {
	o.Url = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *V2HttpConfiguration) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2HttpConfiguration) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *V2HttpConfiguration) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *V2HttpConfiguration) SetMethod(v string) {
	o.Method = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *V2HttpConfiguration) GetBody() string {
	if o == nil || o.Body == nil {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2HttpConfiguration) GetBodyOk() (*string, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *V2HttpConfiguration) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *V2HttpConfiguration) SetBody(v string) {
	o.Body = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *V2HttpConfiguration) GetHeaders() []V2Header {
	if o == nil || o.Headers == nil {
		var ret []V2Header
		return ret
	}
	return *o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2HttpConfiguration) GetHeadersOk() (*[]V2Header, bool) {
	if o == nil || o.Headers == nil {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *V2HttpConfiguration) HasHeaders() bool {
	if o != nil && o.Headers != nil {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []V2Header and assigns it to the Headers field.
func (o *V2HttpConfiguration) SetHeaders(v []V2Header) {
	o.Headers = &v
}

// GetBasic returns the Basic field value if set, zero value otherwise.
func (o *V2HttpConfiguration) GetBasic() V2HttpConfigurationBasicAuth {
	if o == nil || o.Basic == nil {
		var ret V2HttpConfigurationBasicAuth
		return ret
	}
	return *o.Basic
}

// GetBasicOk returns a tuple with the Basic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2HttpConfiguration) GetBasicOk() (*V2HttpConfigurationBasicAuth, bool) {
	if o == nil || o.Basic == nil {
		return nil, false
	}
	return o.Basic, true
}

// HasBasic returns a boolean if a field has been set.
func (o *V2HttpConfiguration) HasBasic() bool {
	if o != nil && o.Basic != nil {
		return true
	}

	return false
}

// SetBasic gets a reference to the given V2HttpConfigurationBasicAuth and assigns it to the Basic field.
func (o *V2HttpConfiguration) SetBasic(v V2HttpConfigurationBasicAuth) {
	o.Basic = &v
}

// GetDigest returns the Digest field value if set, zero value otherwise.
func (o *V2HttpConfiguration) GetDigest() V2HttpConfigurationDigestAuth {
	if o == nil || o.Digest == nil {
		var ret V2HttpConfigurationDigestAuth
		return ret
	}
	return *o.Digest
}

// GetDigestOk returns a tuple with the Digest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2HttpConfiguration) GetDigestOk() (*V2HttpConfigurationDigestAuth, bool) {
	if o == nil || o.Digest == nil {
		return nil, false
	}
	return o.Digest, true
}

// HasDigest returns a boolean if a field has been set.
func (o *V2HttpConfiguration) HasDigest() bool {
	if o != nil && o.Digest != nil {
		return true
	}

	return false
}

// SetDigest gets a reference to the given V2HttpConfigurationDigestAuth and assigns it to the Digest field.
func (o *V2HttpConfiguration) SetDigest(v V2HttpConfigurationDigestAuth) {
	o.Digest = &v
}

// GetClientCertificate returns the ClientCertificate field value if set, zero value otherwise.
func (o *V2HttpConfiguration) GetClientCertificate() V2HttpConfigurationClientCertificate {
	if o == nil || o.ClientCertificate == nil {
		var ret V2HttpConfigurationClientCertificate
		return ret
	}
	return *o.ClientCertificate
}

// GetClientCertificateOk returns a tuple with the ClientCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2HttpConfiguration) GetClientCertificateOk() (*V2HttpConfigurationClientCertificate, bool) {
	if o == nil || o.ClientCertificate == nil {
		return nil, false
	}
	return o.ClientCertificate, true
}

// HasClientCertificate returns a boolean if a field has been set.
func (o *V2HttpConfiguration) HasClientCertificate() bool {
	if o != nil && o.ClientCertificate != nil {
		return true
	}

	return false
}

// SetClientCertificate gets a reference to the given V2HttpConfigurationClientCertificate and assigns it to the ClientCertificate field.
func (o *V2HttpConfiguration) SetClientCertificate(v V2HttpConfigurationClientCertificate) {
	o.ClientCertificate = &v
}

// GetValidateCertificate returns the ValidateCertificate field value if set, zero value otherwise.
func (o *V2HttpConfiguration) GetValidateCertificate() bool {
	if o == nil || o.ValidateCertificate == nil {
		var ret bool
		return ret
	}
	return *o.ValidateCertificate
}

// GetValidateCertificateOk returns a tuple with the ValidateCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2HttpConfiguration) GetValidateCertificateOk() (*bool, bool) {
	if o == nil || o.ValidateCertificate == nil {
		return nil, false
	}
	return o.ValidateCertificate, true
}

// HasValidateCertificate returns a boolean if a field has been set.
func (o *V2HttpConfiguration) HasValidateCertificate() bool {
	if o != nil && o.ValidateCertificate != nil {
		return true
	}

	return false
}

// SetValidateCertificate gets a reference to the given bool and assigns it to the ValidateCertificate field.
func (o *V2HttpConfiguration) SetValidateCertificate(v bool) {
	o.ValidateCertificate = &v
}

func (o V2HttpConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}
	if o.Body != nil {
		toSerialize["body"] = o.Body
	}
	if o.Headers != nil {
		toSerialize["headers"] = o.Headers
	}
	if o.Basic != nil {
		toSerialize["basic"] = o.Basic
	}
	if o.Digest != nil {
		toSerialize["digest"] = o.Digest
	}
	if o.ClientCertificate != nil {
		toSerialize["clientCertificate"] = o.ClientCertificate
	}
	if o.ValidateCertificate != nil {
		toSerialize["validateCertificate"] = o.ValidateCertificate
	}
	return json.Marshal(toSerialize)
}

type NullableV2HttpConfiguration struct {
	value *V2HttpConfiguration
	isSet bool
}

func (v NullableV2HttpConfiguration) Get() *V2HttpConfiguration {
	return v.value
}

func (v *NullableV2HttpConfiguration) Set(val *V2HttpConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableV2HttpConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableV2HttpConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2HttpConfiguration(val *V2HttpConfiguration) *NullableV2HttpConfiguration {
	return &NullableV2HttpConfiguration{value: val, isSet: true}
}

func (v NullableV2HttpConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2HttpConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
