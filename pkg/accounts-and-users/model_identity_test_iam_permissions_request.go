/*
 * Accounts and Users
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package accounts-and-users

import (
	"encoding/json"
)

// IdentityTestIAMPermissionsRequest struct for IdentityTestIAMPermissionsRequest
type IdentityTestIAMPermissionsRequest struct {
	// The set of permissions to test
	Permissions *[]string `json:"permissions,omitempty"`
}

// NewIdentityTestIAMPermissionsRequest instantiates a new IdentityTestIAMPermissionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityTestIAMPermissionsRequest() *IdentityTestIAMPermissionsRequest {
	this := IdentityTestIAMPermissionsRequest{}
	return &this
}

// NewIdentityTestIAMPermissionsRequestWithDefaults instantiates a new IdentityTestIAMPermissionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityTestIAMPermissionsRequestWithDefaults() *IdentityTestIAMPermissionsRequest {
	this := IdentityTestIAMPermissionsRequest{}
	return &this
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *IdentityTestIAMPermissionsRequest) GetPermissions() []string {
	if o == nil || o.Permissions == nil {
		var ret []string
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityTestIAMPermissionsRequest) GetPermissionsOk() (*[]string, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *IdentityTestIAMPermissionsRequest) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []string and assigns it to the Permissions field.
func (o *IdentityTestIAMPermissionsRequest) SetPermissions(v []string) {
	o.Permissions = &v
}

func (o IdentityTestIAMPermissionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityTestIAMPermissionsRequest struct {
	value *IdentityTestIAMPermissionsRequest
	isSet bool
}

func (v NullableIdentityTestIAMPermissionsRequest) Get() *IdentityTestIAMPermissionsRequest {
	return v.value
}

func (v *NullableIdentityTestIAMPermissionsRequest) Set(val *IdentityTestIAMPermissionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityTestIAMPermissionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityTestIAMPermissionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityTestIAMPermissionsRequest(val *IdentityTestIAMPermissionsRequest) *NullableIdentityTestIAMPermissionsRequest {
	return &NullableIdentityTestIAMPermissionsRequest{value: val, isSet: true}
}

func (v NullableIdentityTestIAMPermissionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityTestIAMPermissionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}