/*
 * Edge Compute
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package edge-compute

import (
	"encoding/json"
)

// V1NetworkInterface Network interfaces that will be created on instances in the workload
type V1NetworkInterface struct {
	// A network interface's name
	Network *string `json:"network,omitempty"`
}

// NewV1NetworkInterface instantiates a new V1NetworkInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1NetworkInterface() *V1NetworkInterface {
	this := V1NetworkInterface{}
	return &this
}

// NewV1NetworkInterfaceWithDefaults instantiates a new V1NetworkInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1NetworkInterfaceWithDefaults() *V1NetworkInterface {
	this := V1NetworkInterface{}
	return &this
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *V1NetworkInterface) GetNetwork() string {
	if o == nil || o.Network == nil {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1NetworkInterface) GetNetworkOk() (*string, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *V1NetworkInterface) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *V1NetworkInterface) SetNetwork(v string) {
	o.Network = &v
}

func (o V1NetworkInterface) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Network != nil {
		toSerialize["network"] = o.Network
	}
	return json.Marshal(toSerialize)
}

type NullableV1NetworkInterface struct {
	value *V1NetworkInterface
	isSet bool
}

func (v NullableV1NetworkInterface) Get() *V1NetworkInterface {
	return v.value
}

func (v *NullableV1NetworkInterface) Set(val *V1NetworkInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableV1NetworkInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableV1NetworkInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1NetworkInterface(val *V1NetworkInterface) *NullableV1NetworkInterface {
	return &NullableV1NetworkInterface{value: val, isSet: true}
}

func (v NullableV1NetworkInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1NetworkInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}