/*
 * Content Delivery Network
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package cdn

import (
	"encoding/json"
)

// CustconfOriginPersistentConnections The origin persistent connection settings are used to enable persistent connections to origins
type CustconfOriginPersistentConnections struct {
	// This is used by the API to perform conflict checking
	Id *string `json:"id,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
}

// NewCustconfOriginPersistentConnections instantiates a new CustconfOriginPersistentConnections object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustconfOriginPersistentConnections() *CustconfOriginPersistentConnections {
	this := CustconfOriginPersistentConnections{}
	return &this
}

// NewCustconfOriginPersistentConnectionsWithDefaults instantiates a new CustconfOriginPersistentConnections object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustconfOriginPersistentConnectionsWithDefaults() *CustconfOriginPersistentConnections {
	this := CustconfOriginPersistentConnections{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustconfOriginPersistentConnections) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustconfOriginPersistentConnections) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustconfOriginPersistentConnections) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CustconfOriginPersistentConnections) SetId(v string) {
	o.Id = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CustconfOriginPersistentConnections) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustconfOriginPersistentConnections) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CustconfOriginPersistentConnections) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CustconfOriginPersistentConnections) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o CustconfOriginPersistentConnections) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableCustconfOriginPersistentConnections struct {
	value *CustconfOriginPersistentConnections
	isSet bool
}

func (v NullableCustconfOriginPersistentConnections) Get() *CustconfOriginPersistentConnections {
	return v.value
}

func (v *NullableCustconfOriginPersistentConnections) Set(val *CustconfOriginPersistentConnections) {
	v.value = val
	v.isSet = true
}

func (v NullableCustconfOriginPersistentConnections) IsSet() bool {
	return v.isSet
}

func (v *NullableCustconfOriginPersistentConnections) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustconfOriginPersistentConnections(val *CustconfOriginPersistentConnections) *NullableCustconfOriginPersistentConnections {
	return &NullableCustconfOriginPersistentConnections{value: val, isSet: true}
}

func (v NullableCustconfOriginPersistentConnections) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustconfOriginPersistentConnections) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
