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

// DeliveryStackPathStorageOrigin A StackPath storage bucket origin
type DeliveryStackPathStorageOrigin struct {
	// The origin bucket's name
	BucketName *string `json:"bucketName,omitempty"`
	// The origin bucket's region
	BucketRegion *string `json:"bucketRegion,omitempty"`
	Authentication *DeliveryOriginAuthentication `json:"authentication,omitempty"`
}

// NewDeliveryStackPathStorageOrigin instantiates a new DeliveryStackPathStorageOrigin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryStackPathStorageOrigin() *DeliveryStackPathStorageOrigin {
	this := DeliveryStackPathStorageOrigin{}
	return &this
}

// NewDeliveryStackPathStorageOriginWithDefaults instantiates a new DeliveryStackPathStorageOrigin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryStackPathStorageOriginWithDefaults() *DeliveryStackPathStorageOrigin {
	this := DeliveryStackPathStorageOrigin{}
	return &this
}

// GetBucketName returns the BucketName field value if set, zero value otherwise.
func (o *DeliveryStackPathStorageOrigin) GetBucketName() string {
	if o == nil || o.BucketName == nil {
		var ret string
		return ret
	}
	return *o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryStackPathStorageOrigin) GetBucketNameOk() (*string, bool) {
	if o == nil || o.BucketName == nil {
		return nil, false
	}
	return o.BucketName, true
}

// HasBucketName returns a boolean if a field has been set.
func (o *DeliveryStackPathStorageOrigin) HasBucketName() bool {
	if o != nil && o.BucketName != nil {
		return true
	}

	return false
}

// SetBucketName gets a reference to the given string and assigns it to the BucketName field.
func (o *DeliveryStackPathStorageOrigin) SetBucketName(v string) {
	o.BucketName = &v
}

// GetBucketRegion returns the BucketRegion field value if set, zero value otherwise.
func (o *DeliveryStackPathStorageOrigin) GetBucketRegion() string {
	if o == nil || o.BucketRegion == nil {
		var ret string
		return ret
	}
	return *o.BucketRegion
}

// GetBucketRegionOk returns a tuple with the BucketRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryStackPathStorageOrigin) GetBucketRegionOk() (*string, bool) {
	if o == nil || o.BucketRegion == nil {
		return nil, false
	}
	return o.BucketRegion, true
}

// HasBucketRegion returns a boolean if a field has been set.
func (o *DeliveryStackPathStorageOrigin) HasBucketRegion() bool {
	if o != nil && o.BucketRegion != nil {
		return true
	}

	return false
}

// SetBucketRegion gets a reference to the given string and assigns it to the BucketRegion field.
func (o *DeliveryStackPathStorageOrigin) SetBucketRegion(v string) {
	o.BucketRegion = &v
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *DeliveryStackPathStorageOrigin) GetAuthentication() DeliveryOriginAuthentication {
	if o == nil || o.Authentication == nil {
		var ret DeliveryOriginAuthentication
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryStackPathStorageOrigin) GetAuthenticationOk() (*DeliveryOriginAuthentication, bool) {
	if o == nil || o.Authentication == nil {
		return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *DeliveryStackPathStorageOrigin) HasAuthentication() bool {
	if o != nil && o.Authentication != nil {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given DeliveryOriginAuthentication and assigns it to the Authentication field.
func (o *DeliveryStackPathStorageOrigin) SetAuthentication(v DeliveryOriginAuthentication) {
	o.Authentication = &v
}

func (o DeliveryStackPathStorageOrigin) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BucketName != nil {
		toSerialize["bucketName"] = o.BucketName
	}
	if o.BucketRegion != nil {
		toSerialize["bucketRegion"] = o.BucketRegion
	}
	if o.Authentication != nil {
		toSerialize["authentication"] = o.Authentication
	}
	return json.Marshal(toSerialize)
}

type NullableDeliveryStackPathStorageOrigin struct {
	value *DeliveryStackPathStorageOrigin
	isSet bool
}

func (v NullableDeliveryStackPathStorageOrigin) Get() *DeliveryStackPathStorageOrigin {
	return v.value
}

func (v *NullableDeliveryStackPathStorageOrigin) Set(val *DeliveryStackPathStorageOrigin) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryStackPathStorageOrigin) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryStackPathStorageOrigin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryStackPathStorageOrigin(val *DeliveryStackPathStorageOrigin) *NullableDeliveryStackPathStorageOrigin {
	return &NullableDeliveryStackPathStorageOrigin{value: val, isSet: true}
}

func (v NullableDeliveryStackPathStorageOrigin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryStackPathStorageOrigin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}