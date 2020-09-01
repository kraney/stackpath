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

// V1ProtocolTcpUdp TCP or UDP protocol convenience matching
type V1ProtocolTcpUdp struct {
	// List of destination ports to allow 1-65535
	DestinationPorts *[]string `json:"destinationPorts,omitempty"`
	// List of source ports to allow 1-65535, defaults to 1000-65535
	SourcePorts *[]string `json:"sourcePorts,omitempty"`
}

// NewV1ProtocolTcpUdp instantiates a new V1ProtocolTcpUdp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ProtocolTcpUdp() *V1ProtocolTcpUdp {
	this := V1ProtocolTcpUdp{}
	return &this
}

// NewV1ProtocolTcpUdpWithDefaults instantiates a new V1ProtocolTcpUdp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ProtocolTcpUdpWithDefaults() *V1ProtocolTcpUdp {
	this := V1ProtocolTcpUdp{}
	return &this
}

// GetDestinationPorts returns the DestinationPorts field value if set, zero value otherwise.
func (o *V1ProtocolTcpUdp) GetDestinationPorts() []string {
	if o == nil || o.DestinationPorts == nil {
		var ret []string
		return ret
	}
	return *o.DestinationPorts
}

// GetDestinationPortsOk returns a tuple with the DestinationPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ProtocolTcpUdp) GetDestinationPortsOk() (*[]string, bool) {
	if o == nil || o.DestinationPorts == nil {
		return nil, false
	}
	return o.DestinationPorts, true
}

// HasDestinationPorts returns a boolean if a field has been set.
func (o *V1ProtocolTcpUdp) HasDestinationPorts() bool {
	if o != nil && o.DestinationPorts != nil {
		return true
	}

	return false
}

// SetDestinationPorts gets a reference to the given []string and assigns it to the DestinationPorts field.
func (o *V1ProtocolTcpUdp) SetDestinationPorts(v []string) {
	o.DestinationPorts = &v
}

// GetSourcePorts returns the SourcePorts field value if set, zero value otherwise.
func (o *V1ProtocolTcpUdp) GetSourcePorts() []string {
	if o == nil || o.SourcePorts == nil {
		var ret []string
		return ret
	}
	return *o.SourcePorts
}

// GetSourcePortsOk returns a tuple with the SourcePorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ProtocolTcpUdp) GetSourcePortsOk() (*[]string, bool) {
	if o == nil || o.SourcePorts == nil {
		return nil, false
	}
	return o.SourcePorts, true
}

// HasSourcePorts returns a boolean if a field has been set.
func (o *V1ProtocolTcpUdp) HasSourcePorts() bool {
	if o != nil && o.SourcePorts != nil {
		return true
	}

	return false
}

// SetSourcePorts gets a reference to the given []string and assigns it to the SourcePorts field.
func (o *V1ProtocolTcpUdp) SetSourcePorts(v []string) {
	o.SourcePorts = &v
}

func (o V1ProtocolTcpUdp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DestinationPorts != nil {
		toSerialize["destinationPorts"] = o.DestinationPorts
	}
	if o.SourcePorts != nil {
		toSerialize["sourcePorts"] = o.SourcePorts
	}
	return json.Marshal(toSerialize)
}

type NullableV1ProtocolTcpUdp struct {
	value *V1ProtocolTcpUdp
	isSet bool
}

func (v NullableV1ProtocolTcpUdp) Get() *V1ProtocolTcpUdp {
	return v.value
}

func (v *NullableV1ProtocolTcpUdp) Set(val *V1ProtocolTcpUdp) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ProtocolTcpUdp) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ProtocolTcpUdp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ProtocolTcpUdp(val *V1ProtocolTcpUdp) *NullableV1ProtocolTcpUdp {
	return &NullableV1ProtocolTcpUdp{value: val, isSet: true}
}

func (v NullableV1ProtocolTcpUdp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ProtocolTcpUdp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}