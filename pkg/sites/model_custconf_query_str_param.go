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

// CustconfQueryStrParam The reserved query string parameters policy describes all the query string parameters the StackPath CDN caching server reserves for special processing of your requests.
type CustconfQueryStrParam struct {
	// This is used by the API to perform conflict checking
	Id *string `json:"id,omitempty"`
	// This is the name of the query string parameter which contains the name of the file to specify in the Content-Disposition HTTP response header. This setting is typically used by customers to configure a \"friendly name\" for URLs that have obfuscated file names. This setting controls the \"filename\" directive that is part of the Content-Disposition HTTP header.
	DispositionName *string `json:"dispositionName,omitempty"`
	// This is the name of the query string parameter which contains the disposition type to use in the Content-Disposition HTTP header. Typically, this value is set to \"attachment,\" but you may supply a custom string using this setting.
	DispositionType *string `json:"dispositionType,omitempty"`
	// This setting allows you to completely override the Content-Disposition HTTP header that the CDN caching servers use on a response.
	DispositionOverride *string `json:"dispositionOverride,omitempty"`
	// This is the name of the query string parameter that indicates to the CDN the start time offset of the video returned. This parameter is part of the jump-to-time feature that is initiated on a per request basis.
	JumpToTimeStart *string `json:"jumpToTimeStart,omitempty"`
	// This is the name of the query string parameter that indicates to the CDN the end time offset of the video that should be returned. This parameter is part of the jump-to-time feature that is initiated on a per request basis.
	JumpToTimeEnd *string `json:"jumpToTimeEnd,omitempty"`
	// This is the  name of the query string parameter that indicates to the CDN the initial bytes of a video that should be returned before sending the requested byte offset. This parameter is part of the jump-to-byte feature that is initiated on a per request basis.
	JumpToByteInitialBytes *string `json:"jumpToByteInitialBytes,omitempty"`
	// This is the name of the query string parameter that indicates to the CDN the specific offset into the video that is being requested. This parameter is part of the jump-to-byte feature that is initiated on a per request basis.
	JumpToByteStartOffset *string `json:"jumpToByteStartOffset,omitempty"`
	// This is the name of the query string parameter that indicates to the CDN an initial burst rate to use when delivering a file. This parameter is part of the bandwidth limiting feature that is initiated on a per request basis.
	RateLimitInitial *string `json:"rateLimitInitial,omitempty"`
	// This is the name of the query string parameter that indicates to the CDN the sustained rate being requested for the delivery of a file. This parameter is part of the bandwidth throttling feature that is initiated on a per request basis.
	RateLimitSustained *string `json:"rateLimitSustained,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
}

// NewCustconfQueryStrParam instantiates a new CustconfQueryStrParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustconfQueryStrParam() *CustconfQueryStrParam {
	this := CustconfQueryStrParam{}
	return &this
}

// NewCustconfQueryStrParamWithDefaults instantiates a new CustconfQueryStrParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustconfQueryStrParamWithDefaults() *CustconfQueryStrParam {
	this := CustconfQueryStrParam{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustconfQueryStrParam) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustconfQueryStrParam) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustconfQueryStrParam) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CustconfQueryStrParam) SetId(v string) {
	o.Id = &v
}

// GetDispositionName returns the DispositionName field value if set, zero value otherwise.
func (o *CustconfQueryStrParam) GetDispositionName() string {
	if o == nil || o.DispositionName == nil {
		var ret string
		return ret
	}
	return *o.DispositionName
}

// GetDispositionNameOk returns a tuple with the DispositionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustconfQueryStrParam) GetDispositionNameOk() (*string, bool) {
	if o == nil || o.DispositionName == nil {
		return nil, false
	}
	return o.DispositionName, true
}

// HasDispositionName returns a boolean if a field has been set.
func (o *CustconfQueryStrParam) HasDispositionName() bool {
	if o != nil && o.DispositionName != nil {
		return true
	}

	return false
}

// SetDispositionName gets a reference to the given string and assigns it to the DispositionName field.
func (o *CustconfQueryStrParam) SetDispositionName(v string) {
	o.DispositionName = &v
}

// GetDispositionType returns the DispositionType field value if set, zero value otherwise.
func (o *CustconfQueryStrParam) GetDispositionType() string {
	if o == nil || o.DispositionType == nil {
		var ret string
		return ret
	}
	return *o.DispositionType
}

// GetDispositionTypeOk returns a tuple with the DispositionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustconfQueryStrParam) GetDispositionTypeOk() (*string, bool) {
	if o == nil || o.DispositionType == nil {
		return nil, false
	}
	return o.DispositionType, true
}

// HasDispositionType returns a boolean if a field has been set.
func (o *CustconfQueryStrParam) HasDispositionType() bool {
	if o != nil && o.DispositionType != nil {
		return true
	}

	return false
}

// SetDispositionType gets a reference to the given string and assigns it to the DispositionType field.
func (o *CustconfQueryStrParam) SetDispositionType(v string) {
	o.DispositionType = &v
}

// GetDispositionOverride returns the DispositionOverride field value if set, zero value otherwise.
func (o *CustconfQueryStrParam) GetDispositionOverride() string {
	if o == nil || o.DispositionOverride == nil {
		var ret string
		return ret
	}
	return *o.DispositionOverride
}

// GetDispositionOverrideOk returns a tuple with the DispositionOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustconfQueryStrParam) GetDispositionOverrideOk() (*string, bool) {
	if o == nil || o.DispositionOverride == nil {
		return nil, false
	}
	return o.DispositionOverride, true
}

// HasDispositionOverride returns a boolean if a field has been set.
func (o *CustconfQueryStrParam) HasDispositionOverride() bool {
	if o != nil && o.DispositionOverride != nil {
		return true
	}

	return false
}

// SetDispositionOverride gets a reference to the given string and assigns it to the DispositionOverride field.
func (o *CustconfQueryStrParam) SetDispositionOverride(v string) {
	o.DispositionOverride = &v
}

// GetJumpToTimeStart returns the JumpToTimeStart field value if set, zero value otherwise.
func (o *CustconfQueryStrParam) GetJumpToTimeStart() string {
	if o == nil || o.JumpToTimeStart == nil {
		var ret string
		return ret
	}
	return *o.JumpToTimeStart
}

// GetJumpToTimeStartOk returns a tuple with the JumpToTimeStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustconfQueryStrParam) GetJumpToTimeStartOk() (*string, bool) {
	if o == nil || o.JumpToTimeStart == nil {
		return nil, false
	}
	return o.JumpToTimeStart, true
}

// HasJumpToTimeStart returns a boolean if a field has been set.
func (o *CustconfQueryStrParam) HasJumpToTimeStart() bool {
	if o != nil && o.JumpToTimeStart != nil {
		return true
	}

	return false
}

// SetJumpToTimeStart gets a reference to the given string and assigns it to the JumpToTimeStart field.
func (o *CustconfQueryStrParam) SetJumpToTimeStart(v string) {
	o.JumpToTimeStart = &v
}

// GetJumpToTimeEnd returns the JumpToTimeEnd field value if set, zero value otherwise.
func (o *CustconfQueryStrParam) GetJumpToTimeEnd() string {
	if o == nil || o.JumpToTimeEnd == nil {
		var ret string
		return ret
	}
	return *o.JumpToTimeEnd
}

// GetJumpToTimeEndOk returns a tuple with the JumpToTimeEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustconfQueryStrParam) GetJumpToTimeEndOk() (*string, bool) {
	if o == nil || o.JumpToTimeEnd == nil {
		return nil, false
	}
	return o.JumpToTimeEnd, true
}

// HasJumpToTimeEnd returns a boolean if a field has been set.
func (o *CustconfQueryStrParam) HasJumpToTimeEnd() bool {
	if o != nil && o.JumpToTimeEnd != nil {
		return true
	}

	return false
}

// SetJumpToTimeEnd gets a reference to the given string and assigns it to the JumpToTimeEnd field.
func (o *CustconfQueryStrParam) SetJumpToTimeEnd(v string) {
	o.JumpToTimeEnd = &v
}

// GetJumpToByteInitialBytes returns the JumpToByteInitialBytes field value if set, zero value otherwise.
func (o *CustconfQueryStrParam) GetJumpToByteInitialBytes() string {
	if o == nil || o.JumpToByteInitialBytes == nil {
		var ret string
		return ret
	}
	return *o.JumpToByteInitialBytes
}

// GetJumpToByteInitialBytesOk returns a tuple with the JumpToByteInitialBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustconfQueryStrParam) GetJumpToByteInitialBytesOk() (*string, bool) {
	if o == nil || o.JumpToByteInitialBytes == nil {
		return nil, false
	}
	return o.JumpToByteInitialBytes, true
}

// HasJumpToByteInitialBytes returns a boolean if a field has been set.
func (o *CustconfQueryStrParam) HasJumpToByteInitialBytes() bool {
	if o != nil && o.JumpToByteInitialBytes != nil {
		return true
	}

	return false
}

// SetJumpToByteInitialBytes gets a reference to the given string and assigns it to the JumpToByteInitialBytes field.
func (o *CustconfQueryStrParam) SetJumpToByteInitialBytes(v string) {
	o.JumpToByteInitialBytes = &v
}

// GetJumpToByteStartOffset returns the JumpToByteStartOffset field value if set, zero value otherwise.
func (o *CustconfQueryStrParam) GetJumpToByteStartOffset() string {
	if o == nil || o.JumpToByteStartOffset == nil {
		var ret string
		return ret
	}
	return *o.JumpToByteStartOffset
}

// GetJumpToByteStartOffsetOk returns a tuple with the JumpToByteStartOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustconfQueryStrParam) GetJumpToByteStartOffsetOk() (*string, bool) {
	if o == nil || o.JumpToByteStartOffset == nil {
		return nil, false
	}
	return o.JumpToByteStartOffset, true
}

// HasJumpToByteStartOffset returns a boolean if a field has been set.
func (o *CustconfQueryStrParam) HasJumpToByteStartOffset() bool {
	if o != nil && o.JumpToByteStartOffset != nil {
		return true
	}

	return false
}

// SetJumpToByteStartOffset gets a reference to the given string and assigns it to the JumpToByteStartOffset field.
func (o *CustconfQueryStrParam) SetJumpToByteStartOffset(v string) {
	o.JumpToByteStartOffset = &v
}

// GetRateLimitInitial returns the RateLimitInitial field value if set, zero value otherwise.
func (o *CustconfQueryStrParam) GetRateLimitInitial() string {
	if o == nil || o.RateLimitInitial == nil {
		var ret string
		return ret
	}
	return *o.RateLimitInitial
}

// GetRateLimitInitialOk returns a tuple with the RateLimitInitial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustconfQueryStrParam) GetRateLimitInitialOk() (*string, bool) {
	if o == nil || o.RateLimitInitial == nil {
		return nil, false
	}
	return o.RateLimitInitial, true
}

// HasRateLimitInitial returns a boolean if a field has been set.
func (o *CustconfQueryStrParam) HasRateLimitInitial() bool {
	if o != nil && o.RateLimitInitial != nil {
		return true
	}

	return false
}

// SetRateLimitInitial gets a reference to the given string and assigns it to the RateLimitInitial field.
func (o *CustconfQueryStrParam) SetRateLimitInitial(v string) {
	o.RateLimitInitial = &v
}

// GetRateLimitSustained returns the RateLimitSustained field value if set, zero value otherwise.
func (o *CustconfQueryStrParam) GetRateLimitSustained() string {
	if o == nil || o.RateLimitSustained == nil {
		var ret string
		return ret
	}
	return *o.RateLimitSustained
}

// GetRateLimitSustainedOk returns a tuple with the RateLimitSustained field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustconfQueryStrParam) GetRateLimitSustainedOk() (*string, bool) {
	if o == nil || o.RateLimitSustained == nil {
		return nil, false
	}
	return o.RateLimitSustained, true
}

// HasRateLimitSustained returns a boolean if a field has been set.
func (o *CustconfQueryStrParam) HasRateLimitSustained() bool {
	if o != nil && o.RateLimitSustained != nil {
		return true
	}

	return false
}

// SetRateLimitSustained gets a reference to the given string and assigns it to the RateLimitSustained field.
func (o *CustconfQueryStrParam) SetRateLimitSustained(v string) {
	o.RateLimitSustained = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CustconfQueryStrParam) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustconfQueryStrParam) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CustconfQueryStrParam) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CustconfQueryStrParam) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o CustconfQueryStrParam) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.DispositionName != nil {
		toSerialize["dispositionName"] = o.DispositionName
	}
	if o.DispositionType != nil {
		toSerialize["dispositionType"] = o.DispositionType
	}
	if o.DispositionOverride != nil {
		toSerialize["dispositionOverride"] = o.DispositionOverride
	}
	if o.JumpToTimeStart != nil {
		toSerialize["jumpToTimeStart"] = o.JumpToTimeStart
	}
	if o.JumpToTimeEnd != nil {
		toSerialize["jumpToTimeEnd"] = o.JumpToTimeEnd
	}
	if o.JumpToByteInitialBytes != nil {
		toSerialize["jumpToByteInitialBytes"] = o.JumpToByteInitialBytes
	}
	if o.JumpToByteStartOffset != nil {
		toSerialize["jumpToByteStartOffset"] = o.JumpToByteStartOffset
	}
	if o.RateLimitInitial != nil {
		toSerialize["rateLimitInitial"] = o.RateLimitInitial
	}
	if o.RateLimitSustained != nil {
		toSerialize["rateLimitSustained"] = o.RateLimitSustained
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableCustconfQueryStrParam struct {
	value *CustconfQueryStrParam
	isSet bool
}

func (v NullableCustconfQueryStrParam) Get() *CustconfQueryStrParam {
	return v.value
}

func (v *NullableCustconfQueryStrParam) Set(val *CustconfQueryStrParam) {
	v.value = val
	v.isSet = true
}

func (v NullableCustconfQueryStrParam) IsSet() bool {
	return v.isSet
}

func (v *NullableCustconfQueryStrParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustconfQueryStrParam(val *CustconfQueryStrParam) *NullableCustconfQueryStrParam {
	return &NullableCustconfQueryStrParam{value: val, isSet: true}
}

func (v NullableCustconfQueryStrParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustconfQueryStrParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}