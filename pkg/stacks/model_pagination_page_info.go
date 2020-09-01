/*
 * Stacks
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package stacks

import (
	"encoding/json"
)

// PaginationPageInfo Information about a paginated response  This is modeled after the GraphQL Relay spec to support both cursor based pagination and traditional offset based pagination.
type PaginationPageInfo struct {
	// The total number of items in the dataset
	TotalCount *string `json:"totalCount,omitempty"`
	// Whether or not a previous page of data exists
	HasPreviousPage *bool `json:"hasPreviousPage,omitempty"`
	// Whether or not another page of data is available
	HasNextPage *bool `json:"hasNextPage,omitempty"`
	// The cursor for the first item in the set of data returned
	StartCursor *string `json:"startCursor,omitempty"`
	// The cursor for the last item in the set of data returned
	EndCursor *string `json:"endCursor,omitempty"`
}

// NewPaginationPageInfo instantiates a new PaginationPageInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginationPageInfo() *PaginationPageInfo {
	this := PaginationPageInfo{}
	return &this
}

// NewPaginationPageInfoWithDefaults instantiates a new PaginationPageInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginationPageInfoWithDefaults() *PaginationPageInfo {
	this := PaginationPageInfo{}
	return &this
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *PaginationPageInfo) GetTotalCount() string {
	if o == nil || o.TotalCount == nil {
		var ret string
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationPageInfo) GetTotalCountOk() (*string, bool) {
	if o == nil || o.TotalCount == nil {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *PaginationPageInfo) HasTotalCount() bool {
	if o != nil && o.TotalCount != nil {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given string and assigns it to the TotalCount field.
func (o *PaginationPageInfo) SetTotalCount(v string) {
	o.TotalCount = &v
}

// GetHasPreviousPage returns the HasPreviousPage field value if set, zero value otherwise.
func (o *PaginationPageInfo) GetHasPreviousPage() bool {
	if o == nil || o.HasPreviousPage == nil {
		var ret bool
		return ret
	}
	return *o.HasPreviousPage
}

// GetHasPreviousPageOk returns a tuple with the HasPreviousPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationPageInfo) GetHasPreviousPageOk() (*bool, bool) {
	if o == nil || o.HasPreviousPage == nil {
		return nil, false
	}
	return o.HasPreviousPage, true
}

// HasHasPreviousPage returns a boolean if a field has been set.
func (o *PaginationPageInfo) HasHasPreviousPage() bool {
	if o != nil && o.HasPreviousPage != nil {
		return true
	}

	return false
}

// SetHasPreviousPage gets a reference to the given bool and assigns it to the HasPreviousPage field.
func (o *PaginationPageInfo) SetHasPreviousPage(v bool) {
	o.HasPreviousPage = &v
}

// GetHasNextPage returns the HasNextPage field value if set, zero value otherwise.
func (o *PaginationPageInfo) GetHasNextPage() bool {
	if o == nil || o.HasNextPage == nil {
		var ret bool
		return ret
	}
	return *o.HasNextPage
}

// GetHasNextPageOk returns a tuple with the HasNextPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationPageInfo) GetHasNextPageOk() (*bool, bool) {
	if o == nil || o.HasNextPage == nil {
		return nil, false
	}
	return o.HasNextPage, true
}

// HasHasNextPage returns a boolean if a field has been set.
func (o *PaginationPageInfo) HasHasNextPage() bool {
	if o != nil && o.HasNextPage != nil {
		return true
	}

	return false
}

// SetHasNextPage gets a reference to the given bool and assigns it to the HasNextPage field.
func (o *PaginationPageInfo) SetHasNextPage(v bool) {
	o.HasNextPage = &v
}

// GetStartCursor returns the StartCursor field value if set, zero value otherwise.
func (o *PaginationPageInfo) GetStartCursor() string {
	if o == nil || o.StartCursor == nil {
		var ret string
		return ret
	}
	return *o.StartCursor
}

// GetStartCursorOk returns a tuple with the StartCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationPageInfo) GetStartCursorOk() (*string, bool) {
	if o == nil || o.StartCursor == nil {
		return nil, false
	}
	return o.StartCursor, true
}

// HasStartCursor returns a boolean if a field has been set.
func (o *PaginationPageInfo) HasStartCursor() bool {
	if o != nil && o.StartCursor != nil {
		return true
	}

	return false
}

// SetStartCursor gets a reference to the given string and assigns it to the StartCursor field.
func (o *PaginationPageInfo) SetStartCursor(v string) {
	o.StartCursor = &v
}

// GetEndCursor returns the EndCursor field value if set, zero value otherwise.
func (o *PaginationPageInfo) GetEndCursor() string {
	if o == nil || o.EndCursor == nil {
		var ret string
		return ret
	}
	return *o.EndCursor
}

// GetEndCursorOk returns a tuple with the EndCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationPageInfo) GetEndCursorOk() (*string, bool) {
	if o == nil || o.EndCursor == nil {
		return nil, false
	}
	return o.EndCursor, true
}

// HasEndCursor returns a boolean if a field has been set.
func (o *PaginationPageInfo) HasEndCursor() bool {
	if o != nil && o.EndCursor != nil {
		return true
	}

	return false
}

// SetEndCursor gets a reference to the given string and assigns it to the EndCursor field.
func (o *PaginationPageInfo) SetEndCursor(v string) {
	o.EndCursor = &v
}

func (o PaginationPageInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TotalCount != nil {
		toSerialize["totalCount"] = o.TotalCount
	}
	if o.HasPreviousPage != nil {
		toSerialize["hasPreviousPage"] = o.HasPreviousPage
	}
	if o.HasNextPage != nil {
		toSerialize["hasNextPage"] = o.HasNextPage
	}
	if o.StartCursor != nil {
		toSerialize["startCursor"] = o.StartCursor
	}
	if o.EndCursor != nil {
		toSerialize["endCursor"] = o.EndCursor
	}
	return json.Marshal(toSerialize)
}

type NullablePaginationPageInfo struct {
	value *PaginationPageInfo
	isSet bool
}

func (v NullablePaginationPageInfo) Get() *PaginationPageInfo {
	return v.value
}

func (v *NullablePaginationPageInfo) Set(val *PaginationPageInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginationPageInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginationPageInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginationPageInfo(val *PaginationPageInfo) *NullablePaginationPageInfo {
	return &NullablePaginationPageInfo{value: val, isSet: true}
}

func (v NullablePaginationPageInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginationPageInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}