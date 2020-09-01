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

// CustconfOriginPullProtocolProtocolEnumWrapperValue the model 'CustconfOriginPullProtocolProtocolEnumWrapperValue'
type CustconfOriginPullProtocolProtocolEnumWrapperValue string

// List of custconfOriginPullProtocolProtocolEnumWrapperValue
const (
	UNKNOWN CustconfOriginPullProtocolProtocolEnumWrapperValue = "UNKNOWN"
	HTTP CustconfOriginPullProtocolProtocolEnumWrapperValue = "http"
	HTTPS CustconfOriginPullProtocolProtocolEnumWrapperValue = "https"
	MATCH CustconfOriginPullProtocolProtocolEnumWrapperValue = "match"
)

// Ptr returns reference to custconfOriginPullProtocolProtocolEnumWrapperValue value
func (v CustconfOriginPullProtocolProtocolEnumWrapperValue) Ptr() *CustconfOriginPullProtocolProtocolEnumWrapperValue {
	return &v
}


type NullableCustconfOriginPullProtocolProtocolEnumWrapperValue struct {
	value *CustconfOriginPullProtocolProtocolEnumWrapperValue
	isSet bool
}

func (v NullableCustconfOriginPullProtocolProtocolEnumWrapperValue) Get() *CustconfOriginPullProtocolProtocolEnumWrapperValue {
	return v.value
}

func (v *NullableCustconfOriginPullProtocolProtocolEnumWrapperValue) Set(val *CustconfOriginPullProtocolProtocolEnumWrapperValue) {
	v.value = val
	v.isSet = true
}

func (v NullableCustconfOriginPullProtocolProtocolEnumWrapperValue) IsSet() bool {
	return v.isSet
}

func (v *NullableCustconfOriginPullProtocolProtocolEnumWrapperValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustconfOriginPullProtocolProtocolEnumWrapperValue(val *CustconfOriginPullProtocolProtocolEnumWrapperValue) *NullableCustconfOriginPullProtocolProtocolEnumWrapperValue {
	return &NullableCustconfOriginPullProtocolProtocolEnumWrapperValue{value: val, isSet: true}
}

func (v NullableCustconfOriginPullProtocolProtocolEnumWrapperValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustconfOriginPullProtocolProtocolEnumWrapperValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}