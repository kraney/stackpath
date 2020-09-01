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

// CustconfAwsSignedS3PostV4AuthenticationTypeEnumWrapperValue the model 'CustconfAwsSignedS3PostV4AuthenticationTypeEnumWrapperValue'
type CustconfAwsSignedS3PostV4AuthenticationTypeEnumWrapperValue string

// List of custconfAwsSignedS3PostV4AuthenticationTypeEnumWrapperValue
const (
	UNKNOWN CustconfAwsSignedS3PostV4AuthenticationTypeEnumWrapperValue = "UNKNOWN"
	QUERY CustconfAwsSignedS3PostV4AuthenticationTypeEnumWrapperValue = "query"
	HEADER CustconfAwsSignedS3PostV4AuthenticationTypeEnumWrapperValue = "header"
)

// Ptr returns reference to custconfAwsSignedS3PostV4AuthenticationTypeEnumWrapperValue value
func (v CustconfAwsSignedS3PostV4AuthenticationTypeEnumWrapperValue) Ptr() *CustconfAwsSignedS3PostV4AuthenticationTypeEnumWrapperValue {
	return &v
}


type NullableCustconfAwsSignedS3PostV4AuthenticationTypeEnumWrapperValue struct {
	value *CustconfAwsSignedS3PostV4AuthenticationTypeEnumWrapperValue
	isSet bool
}

func (v NullableCustconfAwsSignedS3PostV4AuthenticationTypeEnumWrapperValue) Get() *CustconfAwsSignedS3PostV4AuthenticationTypeEnumWrapperValue {
	return v.value
}

func (v *NullableCustconfAwsSignedS3PostV4AuthenticationTypeEnumWrapperValue) Set(val *CustconfAwsSignedS3PostV4AuthenticationTypeEnumWrapperValue) {
	v.value = val
	v.isSet = true
}

func (v NullableCustconfAwsSignedS3PostV4AuthenticationTypeEnumWrapperValue) IsSet() bool {
	return v.isSet
}

func (v *NullableCustconfAwsSignedS3PostV4AuthenticationTypeEnumWrapperValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustconfAwsSignedS3PostV4AuthenticationTypeEnumWrapperValue(val *CustconfAwsSignedS3PostV4AuthenticationTypeEnumWrapperValue) *NullableCustconfAwsSignedS3PostV4AuthenticationTypeEnumWrapperValue {
	return &NullableCustconfAwsSignedS3PostV4AuthenticationTypeEnumWrapperValue{value: val, isSet: true}
}

func (v NullableCustconfAwsSignedS3PostV4AuthenticationTypeEnumWrapperValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustconfAwsSignedS3PostV4AuthenticationTypeEnumWrapperValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}