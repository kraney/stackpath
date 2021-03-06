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

// StackCreateStackRequest struct for StackCreateStackRequest
type StackCreateStackRequest struct {
	// The ID of the account to create the stack for
	AccountId *string `json:"accountId,omitempty"`
	// The StackPath stack's friendly, text-based unique identifier  This can be used in place of a stack's ID in most situations to identify a stack. It can be no larger than 30 characters, each being a lowercase letter, number, or dash. It also cannot start with a dash, end with a dash, or have consecutive dashes.If this value is not present, it is derived from the name. This value cannot be updated.
	Slug *string `json:"slug,omitempty"`
	// The stack's name
	Name *string `json:"name,omitempty"`
}

// NewStackCreateStackRequest instantiates a new StackCreateStackRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStackCreateStackRequest() *StackCreateStackRequest {
	this := StackCreateStackRequest{}
	return &this
}

// NewStackCreateStackRequestWithDefaults instantiates a new StackCreateStackRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStackCreateStackRequestWithDefaults() *StackCreateStackRequest {
	this := StackCreateStackRequest{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *StackCreateStackRequest) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StackCreateStackRequest) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *StackCreateStackRequest) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *StackCreateStackRequest) SetAccountId(v string) {
	o.AccountId = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *StackCreateStackRequest) GetSlug() string {
	if o == nil || o.Slug == nil {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StackCreateStackRequest) GetSlugOk() (*string, bool) {
	if o == nil || o.Slug == nil {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *StackCreateStackRequest) HasSlug() bool {
	if o != nil && o.Slug != nil {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *StackCreateStackRequest) SetSlug(v string) {
	o.Slug = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StackCreateStackRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StackCreateStackRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StackCreateStackRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StackCreateStackRequest) SetName(v string) {
	o.Name = &v
}

func (o StackCreateStackRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountId != nil {
		toSerialize["accountId"] = o.AccountId
	}
	if o.Slug != nil {
		toSerialize["slug"] = o.Slug
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableStackCreateStackRequest struct {
	value *StackCreateStackRequest
	isSet bool
}

func (v NullableStackCreateStackRequest) Get() *StackCreateStackRequest {
	return v.value
}

func (v *NullableStackCreateStackRequest) Set(val *StackCreateStackRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableStackCreateStackRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableStackCreateStackRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStackCreateStackRequest(val *StackCreateStackRequest) *NullableStackCreateStackRequest {
	return &NullableStackCreateStackRequest{value: val, isSet: true}
}

func (v NullableStackCreateStackRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStackCreateStackRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
