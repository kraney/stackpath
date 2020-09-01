/*
 * Accounts and Users
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package accounts-and-users

import (
	"encoding/json"
	"time"
)

// IdentityUserCredential User API credentials  API credentials are used to generate authorization tokens for use with the StackPath API. Generating API credentials creates a client ID and secret. Client secrets are exposed to the user only once on creation and are not stored at StackPath. Please record your client secret after generating API credentials.  StackPath recommends using one set of API credentials per application and operating environment.
type IdentityUserCredential struct {
	// An API credential's unique identifier
	Id *string `json:"id,omitempty"`
	// An API credential's name  API credential names are typically associated with the user's application and operating environment
	Name *string `json:"name,omitempty"`
	// An API credential's OAuth2 client ID
	ClientId *string `json:"clientId,omitempty"`
	// The date an API credential was created
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The date an API credential was last updated
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// NewIdentityUserCredential instantiates a new IdentityUserCredential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityUserCredential() *IdentityUserCredential {
	this := IdentityUserCredential{}
	return &this
}

// NewIdentityUserCredentialWithDefaults instantiates a new IdentityUserCredential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityUserCredentialWithDefaults() *IdentityUserCredential {
	this := IdentityUserCredential{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IdentityUserCredential) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityUserCredential) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IdentityUserCredential) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IdentityUserCredential) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IdentityUserCredential) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityUserCredential) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IdentityUserCredential) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IdentityUserCredential) SetName(v string) {
	o.Name = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *IdentityUserCredential) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityUserCredential) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *IdentityUserCredential) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *IdentityUserCredential) SetClientId(v string) {
	o.ClientId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *IdentityUserCredential) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityUserCredential) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *IdentityUserCredential) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *IdentityUserCredential) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *IdentityUserCredential) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityUserCredential) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *IdentityUserCredential) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *IdentityUserCredential) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o IdentityUserCredential) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ClientId != nil {
		toSerialize["clientId"] = o.ClientId
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityUserCredential struct {
	value *IdentityUserCredential
	isSet bool
}

func (v NullableIdentityUserCredential) Get() *IdentityUserCredential {
	return v.value
}

func (v *NullableIdentityUserCredential) Set(val *IdentityUserCredential) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityUserCredential) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityUserCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityUserCredential(val *IdentityUserCredential) *NullableIdentityUserCredential {
	return &NullableIdentityUserCredential{value: val, isSet: true}
}

func (v NullableIdentityUserCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityUserCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}