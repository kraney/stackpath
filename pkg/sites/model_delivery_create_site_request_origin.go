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

// DeliveryCreateSiteRequestOrigin The site's origin
type DeliveryCreateSiteRequestOrigin struct {
	// The path the site should request from the origin
	Path *string `json:"path,omitempty"`
	// The origin's fully-qualified domain name
	Hostname *string `json:"hostname,omitempty"`
	// The TCP port the site should connect to for HTTP requests
	Port *int32 `json:"port,omitempty"`
	// The TCP port the site should connect to for HTTPS requests
	SecurePort *int32 `json:"securePort,omitempty"`
	Http *DeliveryHTTPOrigin `json:"http,omitempty"`
	StackpathStorage *DeliveryStackPathStorageOrigin `json:"stackpathStorage,omitempty"`
	S3Storage *DeliveryAWSS3Origin `json:"s3Storage,omitempty"`
	GoogleCloudStorage *DeliveryGoogleStorageOrigin `json:"googleCloudStorage,omitempty"`
}

// NewDeliveryCreateSiteRequestOrigin instantiates a new DeliveryCreateSiteRequestOrigin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryCreateSiteRequestOrigin() *DeliveryCreateSiteRequestOrigin {
	this := DeliveryCreateSiteRequestOrigin{}
	return &this
}

// NewDeliveryCreateSiteRequestOriginWithDefaults instantiates a new DeliveryCreateSiteRequestOrigin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryCreateSiteRequestOriginWithDefaults() *DeliveryCreateSiteRequestOrigin {
	this := DeliveryCreateSiteRequestOrigin{}
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *DeliveryCreateSiteRequestOrigin) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryCreateSiteRequestOrigin) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *DeliveryCreateSiteRequestOrigin) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *DeliveryCreateSiteRequestOrigin) SetPath(v string) {
	o.Path = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *DeliveryCreateSiteRequestOrigin) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryCreateSiteRequestOrigin) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *DeliveryCreateSiteRequestOrigin) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *DeliveryCreateSiteRequestOrigin) SetHostname(v string) {
	o.Hostname = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *DeliveryCreateSiteRequestOrigin) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryCreateSiteRequestOrigin) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *DeliveryCreateSiteRequestOrigin) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *DeliveryCreateSiteRequestOrigin) SetPort(v int32) {
	o.Port = &v
}

// GetSecurePort returns the SecurePort field value if set, zero value otherwise.
func (o *DeliveryCreateSiteRequestOrigin) GetSecurePort() int32 {
	if o == nil || o.SecurePort == nil {
		var ret int32
		return ret
	}
	return *o.SecurePort
}

// GetSecurePortOk returns a tuple with the SecurePort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryCreateSiteRequestOrigin) GetSecurePortOk() (*int32, bool) {
	if o == nil || o.SecurePort == nil {
		return nil, false
	}
	return o.SecurePort, true
}

// HasSecurePort returns a boolean if a field has been set.
func (o *DeliveryCreateSiteRequestOrigin) HasSecurePort() bool {
	if o != nil && o.SecurePort != nil {
		return true
	}

	return false
}

// SetSecurePort gets a reference to the given int32 and assigns it to the SecurePort field.
func (o *DeliveryCreateSiteRequestOrigin) SetSecurePort(v int32) {
	o.SecurePort = &v
}

// GetHttp returns the Http field value if set, zero value otherwise.
func (o *DeliveryCreateSiteRequestOrigin) GetHttp() DeliveryHTTPOrigin {
	if o == nil || o.Http == nil {
		var ret DeliveryHTTPOrigin
		return ret
	}
	return *o.Http
}

// GetHttpOk returns a tuple with the Http field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryCreateSiteRequestOrigin) GetHttpOk() (*DeliveryHTTPOrigin, bool) {
	if o == nil || o.Http == nil {
		return nil, false
	}
	return o.Http, true
}

// HasHttp returns a boolean if a field has been set.
func (o *DeliveryCreateSiteRequestOrigin) HasHttp() bool {
	if o != nil && o.Http != nil {
		return true
	}

	return false
}

// SetHttp gets a reference to the given DeliveryHTTPOrigin and assigns it to the Http field.
func (o *DeliveryCreateSiteRequestOrigin) SetHttp(v DeliveryHTTPOrigin) {
	o.Http = &v
}

// GetStackpathStorage returns the StackpathStorage field value if set, zero value otherwise.
func (o *DeliveryCreateSiteRequestOrigin) GetStackpathStorage() DeliveryStackPathStorageOrigin {
	if o == nil || o.StackpathStorage == nil {
		var ret DeliveryStackPathStorageOrigin
		return ret
	}
	return *o.StackpathStorage
}

// GetStackpathStorageOk returns a tuple with the StackpathStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryCreateSiteRequestOrigin) GetStackpathStorageOk() (*DeliveryStackPathStorageOrigin, bool) {
	if o == nil || o.StackpathStorage == nil {
		return nil, false
	}
	return o.StackpathStorage, true
}

// HasStackpathStorage returns a boolean if a field has been set.
func (o *DeliveryCreateSiteRequestOrigin) HasStackpathStorage() bool {
	if o != nil && o.StackpathStorage != nil {
		return true
	}

	return false
}

// SetStackpathStorage gets a reference to the given DeliveryStackPathStorageOrigin and assigns it to the StackpathStorage field.
func (o *DeliveryCreateSiteRequestOrigin) SetStackpathStorage(v DeliveryStackPathStorageOrigin) {
	o.StackpathStorage = &v
}

// GetS3Storage returns the S3Storage field value if set, zero value otherwise.
func (o *DeliveryCreateSiteRequestOrigin) GetS3Storage() DeliveryAWSS3Origin {
	if o == nil || o.S3Storage == nil {
		var ret DeliveryAWSS3Origin
		return ret
	}
	return *o.S3Storage
}

// GetS3StorageOk returns a tuple with the S3Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryCreateSiteRequestOrigin) GetS3StorageOk() (*DeliveryAWSS3Origin, bool) {
	if o == nil || o.S3Storage == nil {
		return nil, false
	}
	return o.S3Storage, true
}

// HasS3Storage returns a boolean if a field has been set.
func (o *DeliveryCreateSiteRequestOrigin) HasS3Storage() bool {
	if o != nil && o.S3Storage != nil {
		return true
	}

	return false
}

// SetS3Storage gets a reference to the given DeliveryAWSS3Origin and assigns it to the S3Storage field.
func (o *DeliveryCreateSiteRequestOrigin) SetS3Storage(v DeliveryAWSS3Origin) {
	o.S3Storage = &v
}

// GetGoogleCloudStorage returns the GoogleCloudStorage field value if set, zero value otherwise.
func (o *DeliveryCreateSiteRequestOrigin) GetGoogleCloudStorage() DeliveryGoogleStorageOrigin {
	if o == nil || o.GoogleCloudStorage == nil {
		var ret DeliveryGoogleStorageOrigin
		return ret
	}
	return *o.GoogleCloudStorage
}

// GetGoogleCloudStorageOk returns a tuple with the GoogleCloudStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryCreateSiteRequestOrigin) GetGoogleCloudStorageOk() (*DeliveryGoogleStorageOrigin, bool) {
	if o == nil || o.GoogleCloudStorage == nil {
		return nil, false
	}
	return o.GoogleCloudStorage, true
}

// HasGoogleCloudStorage returns a boolean if a field has been set.
func (o *DeliveryCreateSiteRequestOrigin) HasGoogleCloudStorage() bool {
	if o != nil && o.GoogleCloudStorage != nil {
		return true
	}

	return false
}

// SetGoogleCloudStorage gets a reference to the given DeliveryGoogleStorageOrigin and assigns it to the GoogleCloudStorage field.
func (o *DeliveryCreateSiteRequestOrigin) SetGoogleCloudStorage(v DeliveryGoogleStorageOrigin) {
	o.GoogleCloudStorage = &v
}

func (o DeliveryCreateSiteRequestOrigin) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.Hostname != nil {
		toSerialize["hostname"] = o.Hostname
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.SecurePort != nil {
		toSerialize["securePort"] = o.SecurePort
	}
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

type NullableDeliveryCreateSiteRequestOrigin struct {
	value *DeliveryCreateSiteRequestOrigin
	isSet bool
}

func (v NullableDeliveryCreateSiteRequestOrigin) Get() *DeliveryCreateSiteRequestOrigin {
	return v.value
}

func (v *NullableDeliveryCreateSiteRequestOrigin) Set(val *DeliveryCreateSiteRequestOrigin) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryCreateSiteRequestOrigin) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryCreateSiteRequestOrigin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryCreateSiteRequestOrigin(val *DeliveryCreateSiteRequestOrigin) *NullableDeliveryCreateSiteRequestOrigin {
	return &NullableDeliveryCreateSiteRequestOrigin{value: val, isSet: true}
}

func (v NullableDeliveryCreateSiteRequestOrigin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryCreateSiteRequestOrigin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
