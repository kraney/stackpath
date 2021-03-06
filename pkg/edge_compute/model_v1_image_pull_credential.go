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

// V1ImagePullCredential The credentials that should be used to pull the container image
type V1ImagePullCredential struct {
	DockerRegistry *V1DockerRegistryCredentials `json:"dockerRegistry,omitempty"`
}

// NewV1ImagePullCredential instantiates a new V1ImagePullCredential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ImagePullCredential() *V1ImagePullCredential {
	this := V1ImagePullCredential{}
	return &this
}

// NewV1ImagePullCredentialWithDefaults instantiates a new V1ImagePullCredential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ImagePullCredentialWithDefaults() *V1ImagePullCredential {
	this := V1ImagePullCredential{}
	return &this
}

// GetDockerRegistry returns the DockerRegistry field value if set, zero value otherwise.
func (o *V1ImagePullCredential) GetDockerRegistry() V1DockerRegistryCredentials {
	if o == nil || o.DockerRegistry == nil {
		var ret V1DockerRegistryCredentials
		return ret
	}
	return *o.DockerRegistry
}

// GetDockerRegistryOk returns a tuple with the DockerRegistry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ImagePullCredential) GetDockerRegistryOk() (*V1DockerRegistryCredentials, bool) {
	if o == nil || o.DockerRegistry == nil {
		return nil, false
	}
	return o.DockerRegistry, true
}

// HasDockerRegistry returns a boolean if a field has been set.
func (o *V1ImagePullCredential) HasDockerRegistry() bool {
	if o != nil && o.DockerRegistry != nil {
		return true
	}

	return false
}

// SetDockerRegistry gets a reference to the given V1DockerRegistryCredentials and assigns it to the DockerRegistry field.
func (o *V1ImagePullCredential) SetDockerRegistry(v V1DockerRegistryCredentials) {
	o.DockerRegistry = &v
}

func (o V1ImagePullCredential) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DockerRegistry != nil {
		toSerialize["dockerRegistry"] = o.DockerRegistry
	}
	return json.Marshal(toSerialize)
}

type NullableV1ImagePullCredential struct {
	value *V1ImagePullCredential
	isSet bool
}

func (v NullableV1ImagePullCredential) Get() *V1ImagePullCredential {
	return v.value
}

func (v *NullableV1ImagePullCredential) Set(val *V1ImagePullCredential) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ImagePullCredential) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ImagePullCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ImagePullCredential(val *V1ImagePullCredential) *NullableV1ImagePullCredential {
	return &NullableV1ImagePullCredential{value: val, isSet: true}
}

func (v NullableV1ImagePullCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ImagePullCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
