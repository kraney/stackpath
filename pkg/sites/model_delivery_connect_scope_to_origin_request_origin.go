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

// DeliveryConnectScopeToOriginRequestOrigin A new origin
type DeliveryConnectScopeToOriginRequestOrigin struct {
	Http *DeliveryHTTPOrigin `json:"http,omitempty"`
	StackpathStorage *DeliveryStackPathStorageOrigin `json:"stackpathStorage,omitempty"`
	S3Storage *DeliveryAWSS3Origin `json:"s3Storage,omitempty"`
	GoogleCloudStorage *DeliveryGoogleStorageOrigin `json:"googleCloudStorage,omitempty"`
}

// NewDeliveryConnectScopeToOriginRequestOrigin instantiates a new DeliveryConnectScopeToOriginRequestOrigin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryConnectScopeToOriginRequestOrigin() *DeliveryConnectScopeToOriginRequestOrigin {
	this := DeliveryConnectScopeToOriginRequestOrigin{}
	return &this
}

// NewDeliveryConnectScopeToOriginRequestOriginWithDefaults instantiates a new DeliveryConnectScopeToOriginRequestOrigin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryConnectScopeToOriginRequestOriginWithDefaults() *DeliveryConnectScopeToOriginRequestOrigin {
	this := DeliveryConnectScopeToOriginRequestOrigin{}
	return &this
}

// GetHttp returns the Http field value if set, zero value otherwise.
func (o *DeliveryConnectScopeToOriginRequestOrigin) GetHttp() DeliveryHTTPOrigin {
	if o == nil || o.Http == nil {
		var ret DeliveryHTTPOrigin
		return ret
	}
	return *o.Http
}

// GetHttpOk returns a tuple with the Http field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryConnectScopeToOriginRequestOrigin) GetHttpOk() (*DeliveryHTTPOrigin, bool) {
	if o == nil || o.Http == nil {
		return nil, false
	}
	return o.Http, true
}

// HasHttp returns a boolean if a field has been set.
func (o *DeliveryConnectScopeToOriginRequestOrigin) HasHttp() bool {
	if o != nil && o.Http != nil {
		return true
	}

	return false
}

// SetHttp gets a reference to the given DeliveryHTTPOrigin and assigns it to the Http field.
func (o *DeliveryConnectScopeToOriginRequestOrigin) SetHttp(v DeliveryHTTPOrigin) {
	o.Http = &v
}

// GetStackpathStorage returns the StackpathStorage field value if set, zero value otherwise.
func (o *DeliveryConnectScopeToOriginRequestOrigin) GetStackpathStorage() DeliveryStackPathStorageOrigin {
	if o == nil || o.StackpathStorage == nil {
		var ret DeliveryStackPathStorageOrigin
		return ret
	}
	return *o.StackpathStorage
}

// GetStackpathStorageOk returns a tuple with the StackpathStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryConnectScopeToOriginRequestOrigin) GetStackpathStorageOk() (*DeliveryStackPathStorageOrigin, bool) {
	if o == nil || o.StackpathStorage == nil {
		return nil, false
	}
	return o.StackpathStorage, true
}

// HasStackpathStorage returns a boolean if a field has been set.
func (o *DeliveryConnectScopeToOriginRequestOrigin) HasStackpathStorage() bool {
	if o != nil && o.StackpathStorage != nil {
		return true
	}

	return false
}

// SetStackpathStorage gets a reference to the given DeliveryStackPathStorageOrigin and assigns it to the StackpathStorage field.
func (o *DeliveryConnectScopeToOriginRequestOrigin) SetStackpathStorage(v DeliveryStackPathStorageOrigin) {
	o.StackpathStorage = &v
}

// GetS3Storage returns the S3Storage field value if set, zero value otherwise.
func (o *DeliveryConnectScopeToOriginRequestOrigin) GetS3Storage() DeliveryAWSS3Origin {
	if o == nil || o.S3Storage == nil {
		var ret DeliveryAWSS3Origin
		return ret
	}
	return *o.S3Storage
}

// GetS3StorageOk returns a tuple with the S3Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryConnectScopeToOriginRequestOrigin) GetS3StorageOk() (*DeliveryAWSS3Origin, bool) {
	if o == nil || o.S3Storage == nil {
		return nil, false
	}
	return o.S3Storage, true
}

// HasS3Storage returns a boolean if a field has been set.
func (o *DeliveryConnectScopeToOriginRequestOrigin) HasS3Storage() bool {
	if o != nil && o.S3Storage != nil {
		return true
	}

	return false
}

// SetS3Storage gets a reference to the given DeliveryAWSS3Origin and assigns it to the S3Storage field.
func (o *DeliveryConnectScopeToOriginRequestOrigin) SetS3Storage(v DeliveryAWSS3Origin) {
	o.S3Storage = &v
}

// GetGoogleCloudStorage returns the GoogleCloudStorage field value if set, zero value otherwise.
func (o *DeliveryConnectScopeToOriginRequestOrigin) GetGoogleCloudStorage() DeliveryGoogleStorageOrigin {
	if o == nil || o.GoogleCloudStorage == nil {
		var ret DeliveryGoogleStorageOrigin
		return ret
	}
	return *o.GoogleCloudStorage
}

// GetGoogleCloudStorageOk returns a tuple with the GoogleCloudStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryConnectScopeToOriginRequestOrigin) GetGoogleCloudStorageOk() (*DeliveryGoogleStorageOrigin, bool) {
	if o == nil || o.GoogleCloudStorage == nil {
		return nil, false
	}
	return o.GoogleCloudStorage, true
}

// HasGoogleCloudStorage returns a boolean if a field has been set.
func (o *DeliveryConnectScopeToOriginRequestOrigin) HasGoogleCloudStorage() bool {
	if o != nil && o.GoogleCloudStorage != nil {
		return true
	}

	return false
}

// SetGoogleCloudStorage gets a reference to the given DeliveryGoogleStorageOrigin and assigns it to the GoogleCloudStorage field.
func (o *DeliveryConnectScopeToOriginRequestOrigin) SetGoogleCloudStorage(v DeliveryGoogleStorageOrigin) {
	o.GoogleCloudStorage = &v
}

func (o DeliveryConnectScopeToOriginRequestOrigin) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Http != nil {
		toSerialize["http"] = o.Http
	}
	if o.StackpathStorage != nil {
		toSerialize["stackpathStorage"] = o.StackpathStorage
	}
	if o.S3Storage != nil {
		toSerialize["s3Storage"] = o.S3Storage
	}
	if o.GoogleCloudStorage != nil {
		toSerialize["googleCloudStorage"] = o.GoogleCloudStorage
	}
	return json.Marshal(toSerialize)
}

type NullableDeliveryConnectScopeToOriginRequestOrigin struct {
	value *DeliveryConnectScopeToOriginRequestOrigin
	isSet bool
}

func (v NullableDeliveryConnectScopeToOriginRequestOrigin) Get() *DeliveryConnectScopeToOriginRequestOrigin {
	return v.value
}

func (v *NullableDeliveryConnectScopeToOriginRequestOrigin) Set(val *DeliveryConnectScopeToOriginRequestOrigin) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryConnectScopeToOriginRequestOrigin) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryConnectScopeToOriginRequestOrigin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryConnectScopeToOriginRequestOrigin(val *DeliveryConnectScopeToOriginRequestOrigin) *NullableDeliveryConnectScopeToOriginRequestOrigin {
	return &NullableDeliveryConnectScopeToOriginRequestOrigin{value: val, isSet: true}
}

func (v NullableDeliveryConnectScopeToOriginRequestOrigin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryConnectScopeToOriginRequestOrigin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}