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

// CustconfAwsSignedOriginPullV4 struct for CustconfAwsSignedOriginPullV4
type CustconfAwsSignedOriginPullV4 struct {
	// This is used by the API to perform conflict checking
	Id *string `json:"id,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	BucketIdentifier *string `json:"bucketIdentifier,omitempty"`
	AccessKeyId *string `json:"accessKeyId,omitempty"`
	SecretAccessKey *string `json:"secretAccessKey,omitempty"`
	AwsRegion *string `json:"awsRegion,omitempty"`
	AwsService *string `json:"awsService,omitempty"`
	ExpireTimeSeconds *float32 `json:"expireTimeSeconds,omitempty"`
	// String of values delimited by a ',' character.
	SignedHeaders *string `json:"signedHeaders,omitempty"`
	AuthenticationType *CustconfAwsSignedOriginPullV4AuthenticationTypeEnumWrapperValue `json:"authenticationType,omitempty"`
	// String of values delimited by a ',' character.
	HeaderFilter *string `json:"headerFilter,omitempty"`
	// String of values delimited by a ',' character.
	MethodFilter *string `json:"methodFilter,omitempty"`
	// String of values delimited by a ',' character.
	PathFilter *string `json:"pathFilter,omitempty"`
}

// NewCustconfAwsSignedOriginPullV4 instantiates a new CustconfAwsSignedOriginPullV4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustconfAwsSignedOriginPullV4() *CustconfAwsSignedOriginPullV4 {
	this := CustconfAwsSignedOriginPullV4{}
	var authenticationType CustconfAwsSignedOriginPullV4AuthenticationTypeEnumWrapperValue = "UNKNOWN"
	this.AuthenticationType = &authenticationType
	return &this
}

// NewCustconfAwsSignedOriginPullV4WithDefaults instantiates a new CustconfAwsSignedOriginPullV4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustconfAwsSignedOriginPullV4WithDefaults() *CustconfAwsSignedOriginPullV4 {
	this := CustconfAwsSignedOriginPullV4{}
	var authenticationType CustconfAwsSignedOriginPullV4AuthenticationTypeEnumWrapperValue = "UNKNOWN"
	this.AuthenticationType = &authenticationType
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustconfAwsSignedOriginPullV4) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustconfAwsSignedOriginPullV4) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustconfAwsSignedOriginPullV4) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CustconfAwsSignedOriginPullV4) SetId(v string) {
	o.Id = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CustconfAwsSignedOriginPullV4) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustconfAwsSignedOriginPullV4) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CustconfAwsSignedOriginPullV4) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CustconfAwsSignedOriginPullV4) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetBucketIdentifier returns the BucketIdentifier field value if set, zero value otherwise.
func (o *CustconfAwsSignedOriginPullV4) GetBucketIdentifier() string {
	if o == nil || o.BucketIdentifier == nil {
		var ret string
		return ret
	}
	return *o.BucketIdentifier
}

// GetBucketIdentifierOk returns a tuple with the BucketIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustconfAwsSignedOriginPullV4) GetBucketIdentifierOk() (*string, bool) {
	if o == nil || o.BucketIdentifier == nil {
		return nil, false
	}
	return o.BucketIdentifier, true
}

// HasBucketIdentifier returns a boolean if a field has been set.
func (o *CustconfAwsSignedOriginPullV4) HasBucketIdentifier() bool {
	if o != nil && o.BucketIdentifier != nil {
		return true
	}

	return false
}

// SetBucketIdentifier gets a reference to the given string and assigns it to the BucketIdentifier field.
func (o *CustconfAwsSignedOriginPullV4) SetBucketIdentifier(v string) {
	o.BucketIdentifier = &v
}

// GetAccessKeyId returns the AccessKeyId field value if set, zero value otherwise.
func (o *CustconfAwsSignedOriginPullV4) GetAccessKeyId() string {
	if o == nil || o.AccessKeyId == nil {
		var ret string
		return ret
	}
	return *o.AccessKeyId
}

// GetAccessKeyIdOk returns a tuple with the AccessKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustconfAwsSignedOriginPullV4) GetAccessKeyIdOk() (*string, bool) {
	if o == nil || o.AccessKeyId == nil {
		return nil, false
	}
	return o.AccessKeyId, true
}

// HasAccessKeyId returns a boolean if a field has been set.
func (o *CustconfAwsSignedOriginPullV4) HasAccessKeyId() bool {
	if o != nil && o.AccessKeyId != nil {
		return true
	}

	return false
}

// SetAccessKeyId gets a reference to the given string and assigns it to the AccessKeyId field.
func (o *CustconfAwsSignedOriginPullV4) SetAccessKeyId(v string) {
	o.AccessKeyId = &v
}

// GetSecretAccessKey returns the SecretAccessKey field value if set, zero value otherwise.
func (o *CustconfAwsSignedOriginPullV4) GetSecretAccessKey() string {
	if o == nil || o.SecretAccessKey == nil {
		var ret string
		return ret
	}
	return *o.SecretAccessKey
}

// GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustconfAwsSignedOriginPullV4) GetSecretAccessKeyOk() (*string, bool) {
	if o == nil || o.SecretAccessKey == nil {
		return nil, false
	}
	return o.SecretAccessKey, true
}

// HasSecretAccessKey returns a boolean if a field has been set.
func (o *CustconfAwsSignedOriginPullV4) HasSecretAccessKey() bool {
	if o != nil && o.SecretAccessKey != nil {
		return true
	}

	return false
}

// SetSecretAccessKey gets a reference to the given string and assigns it to the SecretAccessKey field.
func (o *CustconfAwsSignedOriginPullV4) SetSecretAccessKey(v string) {
	o.SecretAccessKey = &v
}

// GetAwsRegion returns the AwsRegion field value if set, zero value otherwise.
func (o *CustconfAwsSignedOriginPullV4) GetAwsRegion() string {
	if o == nil || o.AwsRegion == nil {
		var ret string
		return ret
	}
	return *o.AwsRegion
}

// GetAwsRegionOk returns a tuple with the AwsRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustconfAwsSignedOriginPullV4) GetAwsRegionOk() (*string, bool) {
	if o == nil || o.AwsRegion == nil {
		return nil, false
	}
	return o.AwsRegion, true
}

// HasAwsRegion returns a boolean if a field has been set.
func (o *CustconfAwsSignedOriginPullV4) HasAwsRegion() bool {
	if o != nil && o.AwsRegion != nil {
		return true
	}

	return false
}

// SetAwsRegion gets a reference to the given string and assigns it to the AwsRegion field.
func (o *CustconfAwsSignedOriginPullV4) SetAwsRegion(v string) {
	o.AwsRegion = &v
}

// GetAwsService returns the AwsService field value if set, zero value otherwise.
func (o *CustconfAwsSignedOriginPullV4) GetAwsService() string {
	if o == nil || o.AwsService == nil {
		var ret string
		return ret
	}
	return *o.AwsService
}

// GetAwsServiceOk returns a tuple with the AwsService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustconfAwsSignedOriginPullV4) GetAwsServiceOk() (*string, bool) {
	if o == nil || o.AwsService == nil {
		return nil, false
	}
	return o.AwsService, true
}

// HasAwsService returns a boolean if a field has been set.
func (o *CustconfAwsSignedOriginPullV4) HasAwsService() bool {
	if o != nil && o.AwsService != nil {
		return true
	}

	return false
}

// SetAwsService gets a reference to the given string and assigns it to the AwsService field.
func (o *CustconfAwsSignedOriginPullV4) SetAwsService(v string) {
	o.AwsService = &v
}

// GetExpireTimeSeconds returns the ExpireTimeSeconds field value if set, zero value otherwise.
func (o *CustconfAwsSignedOriginPullV4) GetExpireTimeSeconds() float32 {
	if o == nil || o.ExpireTimeSeconds == nil {
		var ret float32
		return ret
	}
	return *o.ExpireTimeSeconds
}

// GetExpireTimeSecondsOk returns a tuple with the ExpireTimeSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustconfAwsSignedOriginPullV4) GetExpireTimeSecondsOk() (*float32, bool) {
	if o == nil || o.ExpireTimeSeconds == nil {
		return nil, false
	}
	return o.ExpireTimeSeconds, true
}

// HasExpireTimeSeconds returns a boolean if a field has been set.
func (o *CustconfAwsSignedOriginPullV4) HasExpireTimeSeconds() bool {
	if o != nil && o.ExpireTimeSeconds != nil {
		return true
	}

	return false
}

// SetExpireTimeSeconds gets a reference to the given float32 and assigns it to the ExpireTimeSeconds field.
func (o *CustconfAwsSignedOriginPullV4) SetExpireTimeSeconds(v float32) {
	o.ExpireTimeSeconds = &v
}

// GetSignedHeaders returns the SignedHeaders field value if set, zero value otherwise.
func (o *CustconfAwsSignedOriginPullV4) GetSignedHeaders() string {
	if o == nil || o.SignedHeaders == nil {
		var ret string
		return ret
	}
	return *o.SignedHeaders
}

// GetSignedHeadersOk returns a tuple with the SignedHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustconfAwsSignedOriginPullV4) GetSignedHeadersOk() (*string, bool) {
	if o == nil || o.SignedHeaders == nil {
		return nil, false
	}
	return o.SignedHeaders, true
}

// HasSignedHeaders returns a boolean if a field has been set.
func (o *CustconfAwsSignedOriginPullV4) HasSignedHeaders() bool {
	if o != nil && o.SignedHeaders != nil {
		return true
	}

	return false
}

// SetSignedHeaders gets a reference to the given string and assigns it to the SignedHeaders field.
func (o *CustconfAwsSignedOriginPullV4) SetSignedHeaders(v string) {
	o.SignedHeaders = &v
}

// GetAuthenticationType returns the AuthenticationType field value if set, zero value otherwise.
func (o *CustconfAwsSignedOriginPullV4) GetAuthenticationType() CustconfAwsSignedOriginPullV4AuthenticationTypeEnumWrapperValue {
	if o == nil || o.AuthenticationType == nil {
		var ret CustconfAwsSignedOriginPullV4AuthenticationTypeEnumWrapperValue
		return ret
	}
	return *o.AuthenticationType
}

// GetAuthenticationTypeOk returns a tuple with the AuthenticationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustconfAwsSignedOriginPullV4) GetAuthenticationTypeOk() (*CustconfAwsSignedOriginPullV4AuthenticationTypeEnumWrapperValue, bool) {
	if o == nil || o.AuthenticationType == nil {
		return nil, false
	}
	return o.AuthenticationType, true
}

// HasAuthenticationType returns a boolean if a field has been set.
func (o *CustconfAwsSignedOriginPullV4) HasAuthenticationType() bool {
	if o != nil && o.AuthenticationType != nil {
		return true
	}

	return false
}

// SetAuthenticationType gets a reference to the given CustconfAwsSignedOriginPullV4AuthenticationTypeEnumWrapperValue and assigns it to the AuthenticationType field.
func (o *CustconfAwsSignedOriginPullV4) SetAuthenticationType(v CustconfAwsSignedOriginPullV4AuthenticationTypeEnumWrapperValue) {
	o.AuthenticationType = &v
}

// GetHeaderFilter returns the HeaderFilter field value if set, zero value otherwise.
func (o *CustconfAwsSignedOriginPullV4) GetHeaderFilter() string {
	if o == nil || o.HeaderFilter == nil {
		var ret string
		return ret
	}
	return *o.HeaderFilter
}

// GetHeaderFilterOk returns a tuple with the HeaderFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustconfAwsSignedOriginPullV4) GetHeaderFilterOk() (*string, bool) {
	if o == nil || o.HeaderFilter == nil {
		return nil, false
	}
	return o.HeaderFilter, true
}

// HasHeaderFilter returns a boolean if a field has been set.
func (o *CustconfAwsSignedOriginPullV4) HasHeaderFilter() bool {
	if o != nil && o.HeaderFilter != nil {
		return true
	}

	return false
}

// SetHeaderFilter gets a reference to the given string and assigns it to the HeaderFilter field.
func (o *CustconfAwsSignedOriginPullV4) SetHeaderFilter(v string) {
	o.HeaderFilter = &v
}

// GetMethodFilter returns the MethodFilter field value if set, zero value otherwise.
func (o *CustconfAwsSignedOriginPullV4) GetMethodFilter() string {
	if o == nil || o.MethodFilter == nil {
		var ret string
		return ret
	}
	return *o.MethodFilter
}

// GetMethodFilterOk returns a tuple with the MethodFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustconfAwsSignedOriginPullV4) GetMethodFilterOk() (*string, bool) {
	if o == nil || o.MethodFilter == nil {
		return nil, false
	}
	return o.MethodFilter, true
}

// HasMethodFilter returns a boolean if a field has been set.
func (o *CustconfAwsSignedOriginPullV4) HasMethodFilter() bool {
	if o != nil && o.MethodFilter != nil {
		return true
	}

	return false
}

// SetMethodFilter gets a reference to the given string and assigns it to the MethodFilter field.
func (o *CustconfAwsSignedOriginPullV4) SetMethodFilter(v string) {
	o.MethodFilter = &v
}

// GetPathFilter returns the PathFilter field value if set, zero value otherwise.
func (o *CustconfAwsSignedOriginPullV4) GetPathFilter() string {
	if o == nil || o.PathFilter == nil {
		var ret string
		return ret
	}
	return *o.PathFilter
}

// GetPathFilterOk returns a tuple with the PathFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustconfAwsSignedOriginPullV4) GetPathFilterOk() (*string, bool) {
	if o == nil || o.PathFilter == nil {
		return nil, false
	}
	return o.PathFilter, true
}

// HasPathFilter returns a boolean if a field has been set.
func (o *CustconfAwsSignedOriginPullV4) HasPathFilter() bool {
	if o != nil && o.PathFilter != nil {
		return true
	}

	return false
}

// SetPathFilter gets a reference to the given string and assigns it to the PathFilter field.
func (o *CustconfAwsSignedOriginPullV4) SetPathFilter(v string) {
	o.PathFilter = &v
}

func (o CustconfAwsSignedOriginPullV4) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.BucketIdentifier != nil {
		toSerialize["bucketIdentifier"] = o.BucketIdentifier
	}
	if o.AccessKeyId != nil {
		toSerialize["accessKeyId"] = o.AccessKeyId
	}
	if o.SecretAccessKey != nil {
		toSerialize["secretAccessKey"] = o.SecretAccessKey
	}
	if o.AwsRegion != nil {
		toSerialize["awsRegion"] = o.AwsRegion
	}
	if o.AwsService != nil {
		toSerialize["awsService"] = o.AwsService
	}
	if o.ExpireTimeSeconds != nil {
		toSerialize["expireTimeSeconds"] = o.ExpireTimeSeconds
	}
	if o.SignedHeaders != nil {
		toSerialize["signedHeaders"] = o.SignedHeaders
	}
	if o.AuthenticationType != nil {
		toSerialize["authenticationType"] = o.AuthenticationType
	}
	if o.HeaderFilter != nil {
		toSerialize["headerFilter"] = o.HeaderFilter
	}
	if o.MethodFilter != nil {
		toSerialize["methodFilter"] = o.MethodFilter
	}
	if o.PathFilter != nil {
		toSerialize["pathFilter"] = o.PathFilter
	}
	return json.Marshal(toSerialize)
}

type NullableCustconfAwsSignedOriginPullV4 struct {
	value *CustconfAwsSignedOriginPullV4
	isSet bool
}

func (v NullableCustconfAwsSignedOriginPullV4) Get() *CustconfAwsSignedOriginPullV4 {
	return v.value
}

func (v *NullableCustconfAwsSignedOriginPullV4) Set(val *CustconfAwsSignedOriginPullV4) {
	v.value = val
	v.isSet = true
}

func (v NullableCustconfAwsSignedOriginPullV4) IsSet() bool {
	return v.isSet
}

func (v *NullableCustconfAwsSignedOriginPullV4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustconfAwsSignedOriginPullV4(val *CustconfAwsSignedOriginPullV4) *NullableCustconfAwsSignedOriginPullV4 {
	return &NullableCustconfAwsSignedOriginPullV4{value: val, isSet: true}
}

func (v NullableCustconfAwsSignedOriginPullV4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustconfAwsSignedOriginPullV4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
