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

// CdnGetSiteScriptsResponse The result from a request to retrieve all EdgeEngine scripts on a CDN site
type CdnGetSiteScriptsResponse struct {
	PageInfo *PaginationPageInfo `json:"pageInfo,omitempty"`
	// The requested serverless scripts
	Results *[]CdnSiteScript `json:"results,omitempty"`
}

// NewCdnGetSiteScriptsResponse instantiates a new CdnGetSiteScriptsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCdnGetSiteScriptsResponse() *CdnGetSiteScriptsResponse {
	this := CdnGetSiteScriptsResponse{}
	return &this
}

// NewCdnGetSiteScriptsResponseWithDefaults instantiates a new CdnGetSiteScriptsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCdnGetSiteScriptsResponseWithDefaults() *CdnGetSiteScriptsResponse {
	this := CdnGetSiteScriptsResponse{}
	return &this
}

// GetPageInfo returns the PageInfo field value if set, zero value otherwise.
func (o *CdnGetSiteScriptsResponse) GetPageInfo() PaginationPageInfo {
	if o == nil || o.PageInfo == nil {
		var ret PaginationPageInfo
		return ret
	}
	return *o.PageInfo
}

// GetPageInfoOk returns a tuple with the PageInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdnGetSiteScriptsResponse) GetPageInfoOk() (*PaginationPageInfo, bool) {
	if o == nil || o.PageInfo == nil {
		return nil, false
	}
	return o.PageInfo, true
}

// HasPageInfo returns a boolean if a field has been set.
func (o *CdnGetSiteScriptsResponse) HasPageInfo() bool {
	if o != nil && o.PageInfo != nil {
		return true
	}

	return false
}

// SetPageInfo gets a reference to the given PaginationPageInfo and assigns it to the PageInfo field.
func (o *CdnGetSiteScriptsResponse) SetPageInfo(v PaginationPageInfo) {
	o.PageInfo = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *CdnGetSiteScriptsResponse) GetResults() []CdnSiteScript {
	if o == nil || o.Results == nil {
		var ret []CdnSiteScript
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdnGetSiteScriptsResponse) GetResultsOk() (*[]CdnSiteScript, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *CdnGetSiteScriptsResponse) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []CdnSiteScript and assigns it to the Results field.
func (o *CdnGetSiteScriptsResponse) SetResults(v []CdnSiteScript) {
	o.Results = &v
}

func (o CdnGetSiteScriptsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PageInfo != nil {
		toSerialize["pageInfo"] = o.PageInfo
	}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableCdnGetSiteScriptsResponse struct {
	value *CdnGetSiteScriptsResponse
	isSet bool
}

func (v NullableCdnGetSiteScriptsResponse) Get() *CdnGetSiteScriptsResponse {
	return v.value
}

func (v *NullableCdnGetSiteScriptsResponse) Set(val *CdnGetSiteScriptsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCdnGetSiteScriptsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCdnGetSiteScriptsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCdnGetSiteScriptsResponse(val *CdnGetSiteScriptsResponse) *NullableCdnGetSiteScriptsResponse {
	return &NullableCdnGetSiteScriptsResponse{value: val, isSet: true}
}

func (v NullableCdnGetSiteScriptsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCdnGetSiteScriptsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}