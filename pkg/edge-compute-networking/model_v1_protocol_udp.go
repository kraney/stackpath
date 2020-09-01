/*
 * Edge Compute Networking
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package edge-compute-networking

import (
	"encoding/json"
)

// V1ProtocolUdp UDP protocol matching
type V1ProtocolUdp struct {
	// List of destination ports to allow 1-65535
	DestinationPorts *[]string `json:"destinationPorts,omitempty"`
	// List of source ports to allow 1-65535, defaults to 1000-65535
	SourcePorts *[]string `json:"sourcePorts,omitempty"`
}

// NewV1ProtocolUdp instantiates a new V1ProtocolUdp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ProtocolUdp() *V1ProtocolUdp {
	this := V1ProtocolUdp{}
	return &this
}

// NewV1ProtocolUdpWithDefaults instantiates a new V1ProtocolUdp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ProtocolUdpWithDefaults() *V1ProtocolUdp {
	this := V1ProtocolUdp{}
	return &this
}

// GetDestinationPorts returns the DestinationPorts field value if set, zero value otherwise.
func (o *V1ProtocolUdp) GetDestinationPorts() []string {
	if o == nil || o.DestinationPorts == nil {
		var ret []string
		return ret
	}
	return *o.DestinationPorts
}

// GetDestinationPortsOk returns a tuple with the DestinationPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ProtocolUdp) GetDestinationPortsOk() (*[]string, bool) {
	if o == nil || o.DestinationPorts == nil {
		return nil, false
	}
	return o.DestinationPorts, true
}

// HasDestinationPorts returns a boolean if a field has been set.
func (o *V1ProtocolUdp) HasDestinationPorts() bool {
	if o != nil && o.DestinationPorts != nil {
		return true
	}

	return false
}

// SetDestinationPorts gets a reference to the given []string and assigns it to the DestinationPorts field.
func (o *V1ProtocolUdp) SetDestinationPorts(v []string) {
	o.DestinationPorts = &v
}

// GetSourcePorts returns the SourcePorts field value if set, zero value otherwise.
func (o *V1ProtocolUdp) GetSourcePorts() []string {
	if o == nil || o.SourcePorts == nil {
		var ret []string
		return ret
	}
	return *o.SourcePorts
}

// GetSourcePortsOk returns a tuple with the SourcePorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ProtocolUdp) GetSourcePortsOk() (*[]string, bool) {
	if o == nil || o.SourcePorts == nil {
		return nil, false
	}
	return o.SourcePorts, true
}

// HasSourcePorts returns a boolean if a field has been set.
func (o *V1ProtocolUdp) HasSourcePorts() bool {
	if o != nil && o.SourcePorts != nil {
		return true
	}

	return false
}

// SetSourcePorts gets a reference to the given []string and assigns it to the SourcePorts field.
func (o *V1ProtocolUdp) SetSourcePorts(v []string) {
	o.SourcePorts = &v
}

func (o V1ProtocolUdp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DestinationPorts != nil {
		toSerialize["destinationPorts"] = o.DestinationPorts
	}
	if o.SourcePorts != nil {
		toSerialize["sourcePorts"] = o.SourcePorts
	}
	return json.Marshal(toSerialize)
}

type NullableV1ProtocolUdp struct {
	value *V1ProtocolUdp
	isSet bool
}

func (v NullableV1ProtocolUdp) Get() *V1ProtocolUdp {
	return v.value
}

func (v *NullableV1ProtocolUdp) Set(val *V1ProtocolUdp) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ProtocolUdp) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ProtocolUdp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ProtocolUdp(val *V1ProtocolUdp) *NullableV1ProtocolUdp {
	return &NullableV1ProtocolUdp{value: val, isSet: true}
}

func (v NullableV1ProtocolUdp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ProtocolUdp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}