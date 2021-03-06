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

// CustconfAuthUrlAsymmetricSignTluAlgorithmIdMapEnumWrapperValue the model 'CustconfAuthUrlAsymmetricSignTluAlgorithmIdMapEnumWrapperValue'
type CustconfAuthUrlAsymmetricSignTluAlgorithmIdMapEnumWrapperValue string

// List of custconfAuthUrlAsymmetricSignTluAlgorithmIdMapEnumWrapperValue
const (
	CUSTCONFAUTHURLASYMMETRICSIGNTLUALGORITHMIDMAPENUMWRAPPERVALUE_UNKNOWN CustconfAuthUrlAsymmetricSignTluAlgorithmIdMapEnumWrapperValue = "UNKNOWN"
	CUSTCONFAUTHURLASYMMETRICSIGNTLUALGORITHMIDMAPENUMWRAPPERVALUE_HMACSHA1 CustconfAuthUrlAsymmetricSignTluAlgorithmIdMapEnumWrapperValue = "hmacsha1"
	CUSTCONFAUTHURLASYMMETRICSIGNTLUALGORITHMIDMAPENUMWRAPPERVALUE_HMACSHA256 CustconfAuthUrlAsymmetricSignTluAlgorithmIdMapEnumWrapperValue = "hmacsha256"
)

// Ptr returns reference to custconfAuthUrlAsymmetricSignTluAlgorithmIdMapEnumWrapperValue value
func (v CustconfAuthUrlAsymmetricSignTluAlgorithmIdMapEnumWrapperValue) Ptr() *CustconfAuthUrlAsymmetricSignTluAlgorithmIdMapEnumWrapperValue {
	return &v
}


type NullableCustconfAuthUrlAsymmetricSignTluAlgorithmIdMapEnumWrapperValue struct {
	value *CustconfAuthUrlAsymmetricSignTluAlgorithmIdMapEnumWrapperValue
	isSet bool
}

func (v NullableCustconfAuthUrlAsymmetricSignTluAlgorithmIdMapEnumWrapperValue) Get() *CustconfAuthUrlAsymmetricSignTluAlgorithmIdMapEnumWrapperValue {
	return v.value
}

func (v *NullableCustconfAuthUrlAsymmetricSignTluAlgorithmIdMapEnumWrapperValue) Set(val *CustconfAuthUrlAsymmetricSignTluAlgorithmIdMapEnumWrapperValue) {
	v.value = val
	v.isSet = true
}

func (v NullableCustconfAuthUrlAsymmetricSignTluAlgorithmIdMapEnumWrapperValue) IsSet() bool {
	return v.isSet
}

func (v *NullableCustconfAuthUrlAsymmetricSignTluAlgorithmIdMapEnumWrapperValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustconfAuthUrlAsymmetricSignTluAlgorithmIdMapEnumWrapperValue(val *CustconfAuthUrlAsymmetricSignTluAlgorithmIdMapEnumWrapperValue) *NullableCustconfAuthUrlAsymmetricSignTluAlgorithmIdMapEnumWrapperValue {
	return &NullableCustconfAuthUrlAsymmetricSignTluAlgorithmIdMapEnumWrapperValue{value: val, isSet: true}
}

func (v NullableCustconfAuthUrlAsymmetricSignTluAlgorithmIdMapEnumWrapperValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustconfAuthUrlAsymmetricSignTluAlgorithmIdMapEnumWrapperValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
