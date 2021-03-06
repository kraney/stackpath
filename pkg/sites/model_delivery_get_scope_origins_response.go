/*
 * Sites
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package sites

import (
	"encoding/json"
)

// DeliveryGetScopeOriginsResponse The response from a request to retrieve a site scope's origins
type DeliveryGetScopeOriginsResponse struct {
	PageInfo *PaginationPageInfo `json:"pageInfo,omitempty"`
	// The requested scope's origins
	Results *[]DeliveryScopeOrigin `json:"results,omitempty"`
}

// NewDeliveryGetScopeOriginsResponse instantiates a new DeliveryGetScopeOriginsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryGetScopeOriginsResponse() *DeliveryGetScopeOriginsResponse {
	this := DeliveryGetScopeOriginsResponse{}
	return &this
}

// NewDeliveryGetScopeOriginsResponseWithDefaults instantiates a new DeliveryGetScopeOriginsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryGetScopeOriginsResponseWithDefaults() *DeliveryGetScopeOriginsResponse {
	this := DeliveryGetScopeOriginsResponse{}
	return &this
}

// GetPageInfo returns the PageInfo field value if set, zero value otherwise.
func (o *DeliveryGetScopeOriginsResponse) GetPageInfo() PaginationPageInfo {
	if o == nil || o.PageInfo == nil {
		var ret PaginationPageInfo
		return ret
	}
	return *o.PageInfo
}

// GetPageInfoOk returns a tuple with the PageInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryGetScopeOriginsResponse) GetPageInfoOk() (*PaginationPageInfo, bool) {
	if o == nil || o.PageInfo == nil {
		return nil, false
	}
	return o.PageInfo, true
}

// HasPageInfo returns a boolean if a field has been set.
func (o *DeliveryGetScopeOriginsResponse) HasPageInfo() bool {
	if o != nil && o.PageInfo != nil {
		return true
	}

	return false
}

// SetPageInfo gets a reference to the given PaginationPageInfo and assigns it to the PageInfo field.
func (o *DeliveryGetScopeOriginsResponse) SetPageInfo(v PaginationPageInfo) {
	o.PageInfo = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *DeliveryGetScopeOriginsResponse) GetResults() []DeliveryScopeOrigin {
	if o == nil || o.Results == nil {
		var ret []DeliveryScopeOrigin
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryGetScopeOriginsResponse) GetResultsOk() (*[]DeliveryScopeOrigin, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *DeliveryGetScopeOriginsResponse) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []DeliveryScopeOrigin and assigns it to the Results field.
func (o *DeliveryGetScopeOriginsResponse) SetResults(v []DeliveryScopeOrigin) {
	o.Results = &v
}

func (o DeliveryGetScopeOriginsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PageInfo != nil {
		toSerialize["pageInfo"] = o.PageInfo
	}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableDeliveryGetScopeOriginsResponse struct {
	value *DeliveryGetScopeOriginsResponse
	isSet bool
}

func (v NullableDeliveryGetScopeOriginsResponse) Get() *DeliveryGetScopeOriginsResponse {
	return v.value
}

func (v *NullableDeliveryGetScopeOriginsResponse) Set(val *DeliveryGetScopeOriginsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryGetScopeOriginsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryGetScopeOriginsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryGetScopeOriginsResponse(val *DeliveryGetScopeOriginsResponse) *NullableDeliveryGetScopeOriginsResponse {
	return &NullableDeliveryGetScopeOriginsResponse{value: val, isSet: true}
}

func (v NullableDeliveryGetScopeOriginsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryGetScopeOriginsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
