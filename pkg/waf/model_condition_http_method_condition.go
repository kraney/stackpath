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

// ConditionHttpMethodCondition Match the incoming HTTP method
type ConditionHttpMethodCondition struct {
	HttpMethod *WafHttpMethod `json:"httpMethod,omitempty"`
}

// NewConditionHttpMethodCondition instantiates a new ConditionHttpMethodCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConditionHttpMethodCondition() *ConditionHttpMethodCondition {
	this := ConditionHttpMethodCondition{}
	var httpMethod WafHttpMethod = "METHOD_UNSPECIFIED"
	this.HttpMethod = &httpMethod
	return &this
}

// NewConditionHttpMethodConditionWithDefaults instantiates a new ConditionHttpMethodCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConditionHttpMethodConditionWithDefaults() *ConditionHttpMethodCondition {
	this := ConditionHttpMethodCondition{}
	var httpMethod WafHttpMethod = "METHOD_UNSPECIFIED"
	this.HttpMethod = &httpMethod
	return &this
}

// GetHttpMethod returns the HttpMethod field value if set, zero value otherwise.
func (o *ConditionHttpMethodCondition) GetHttpMethod() WafHttpMethod {
	if o == nil || o.HttpMethod == nil {
		var ret WafHttpMethod
		return ret
	}
	return *o.HttpMethod
}

// GetHttpMethodOk returns a tuple with the HttpMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionHttpMethodCondition) GetHttpMethodOk() (*WafHttpMethod, bool) {
	if o == nil || o.HttpMethod == nil {
		return nil, false
	}
	return o.HttpMethod, true
}

// HasHttpMethod returns a boolean if a field has been set.
func (o *ConditionHttpMethodCondition) HasHttpMethod() bool {
	if o != nil && o.HttpMethod != nil {
		return true
	}

	return false
}

// SetHttpMethod gets a reference to the given WafHttpMethod and assigns it to the HttpMethod field.
func (o *ConditionHttpMethodCondition) SetHttpMethod(v WafHttpMethod) {
	o.HttpMethod = &v
}

func (o ConditionHttpMethodCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.HttpMethod != nil {
		toSerialize["httpMethod"] = o.HttpMethod
	}
	return json.Marshal(toSerialize)
}

type NullableConditionHttpMethodCondition struct {
	value *ConditionHttpMethodCondition
	isSet bool
}

func (v NullableConditionHttpMethodCondition) Get() *ConditionHttpMethodCondition {
	return v.value
}

func (v *NullableConditionHttpMethodCondition) Set(val *ConditionHttpMethodCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableConditionHttpMethodCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableConditionHttpMethodCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConditionHttpMethodCondition(val *ConditionHttpMethodCondition) *NullableConditionHttpMethodCondition {
	return &NullableConditionHttpMethodCondition{value: val, isSet: true}
}

func (v NullableConditionHttpMethodCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConditionHttpMethodCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}