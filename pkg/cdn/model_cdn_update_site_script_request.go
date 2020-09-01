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

// CdnUpdateSiteScriptRequest struct for CdnUpdateSiteScriptRequest
type CdnUpdateSiteScriptRequest struct {
	// Base64 encoded contents of a serverless script
	Code *string `json:"code,omitempty"`
	// The HTTP request paths that are handled by a serverless script
	Paths *[]string `json:"paths,omitempty"`
}

// NewCdnUpdateSiteScriptRequest instantiates a new CdnUpdateSiteScriptRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCdnUpdateSiteScriptRequest() *CdnUpdateSiteScriptRequest {
	this := CdnUpdateSiteScriptRequest{}
	return &this
}

// NewCdnUpdateSiteScriptRequestWithDefaults instantiates a new CdnUpdateSiteScriptRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCdnUpdateSiteScriptRequestWithDefaults() *CdnUpdateSiteScriptRequest {
	this := CdnUpdateSiteScriptRequest{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *CdnUpdateSiteScriptRequest) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdnUpdateSiteScriptRequest) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *CdnUpdateSiteScriptRequest) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *CdnUpdateSiteScriptRequest) SetCode(v string) {
	o.Code = &v
}

// GetPaths returns the Paths field value if set, zero value otherwise.
func (o *CdnUpdateSiteScriptRequest) GetPaths() []string {
	if o == nil || o.Paths == nil {
		var ret []string
		return ret
	}
	return *o.Paths
}

// GetPathsOk returns a tuple with the Paths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdnUpdateSiteScriptRequest) GetPathsOk() (*[]string, bool) {
	if o == nil || o.Paths == nil {
		return nil, false
	}
	return o.Paths, true
}

// HasPaths returns a boolean if a field has been set.
func (o *CdnUpdateSiteScriptRequest) HasPaths() bool {
	if o != nil && o.Paths != nil {
		return true
	}

	return false
}

// SetPaths gets a reference to the given []string and assigns it to the Paths field.
func (o *CdnUpdateSiteScriptRequest) SetPaths(v []string) {
	o.Paths = &v
}

func (o CdnUpdateSiteScriptRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Paths != nil {
		toSerialize["paths"] = o.Paths
	}
	return json.Marshal(toSerialize)
}

type NullableCdnUpdateSiteScriptRequest struct {
	value *CdnUpdateSiteScriptRequest
	isSet bool
}

func (v NullableCdnUpdateSiteScriptRequest) Get() *CdnUpdateSiteScriptRequest {
	return v.value
}

func (v *NullableCdnUpdateSiteScriptRequest) Set(val *CdnUpdateSiteScriptRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCdnUpdateSiteScriptRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCdnUpdateSiteScriptRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCdnUpdateSiteScriptRequest(val *CdnUpdateSiteScriptRequest) *NullableCdnUpdateSiteScriptRequest {
	return &NullableCdnUpdateSiteScriptRequest{value: val, isSet: true}
}

func (v NullableCdnUpdateSiteScriptRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCdnUpdateSiteScriptRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}