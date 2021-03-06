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

// CustconfOriginPullResumeDownload The CDN attempts to resume downloading Origin Pulls once it successfully reads all the headers of a positive response under certain conditions by sending subsequent origin range requests for the remaining body. The response must be a 2xx to a GET request for the full file or a single range (not multi-range). The response must contain the Last-Modified and ETag headers.
type CustconfOriginPullResumeDownload struct {
	// This is used by the API to perform conflict checking
	Id *string `json:"id,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	// String of values delimited by a ',' character. Comma separated list of glob patterns that indicate which origin pulls this policy applies to based on the status code of the original origin response. This feature is limited to 2xx responses, but this policy provides fine control, such as permitting resume of 200 responses by not 206, or for adding a 2xx response code other than 200 or 206.
	OriginalStatusCodeMatch *string `json:"originalStatusCodeMatch,omitempty"`
	// Minimum number of body bytes that CDN must have consumed during an Origin Pull before encountering an error before it is permitted to attempt resuming the download. This value does not include the headers bytes.
	MinimumBodyBytesConsumed *string `json:"minimumBodyBytesConsumed,omitempty"`
	// Maximum number of times beyond the initial request that the CDN is permitted to attempt resuming an Origin Pull download.
	MaximumAttempts *float32 `json:"maximumAttempts,omitempty"`
	// The CDN resumes an Origin Pull even if the origin does not support range requests. If the origin does not support range requests and/or returns a 200 response for the same version given in the original response, the CDN fast-forwards (reads and discards bytes) until it reaches its previous position in the asset.
	RequireOriginRangeSupport *bool `json:"requireOriginRangeSupport,omitempty"`
	EtagValidation *OriginPullResumeDownloadEtagValidationEnumWrapperValue `json:"etagValidation,omitempty"`
	// String of values delimited by a ',' character.
	HeaderFilter *string `json:"headerFilter,omitempty"`
	// String of values delimited by a ',' character.
	PathFilter *string `json:"pathFilter,omitempty"`
}

// NewCustconfOriginPullResumeDownload instantiates a new CustconfOriginPullResumeDownload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustconfOriginPullResumeDownload() *CustconfOriginPullResumeDownload {
	this := CustconfOriginPullResumeDownload{}
	var etagValidation OriginPullResumeDownloadEtagValidationEnumWrapperValue = "UNKNOWN"
	this.EtagValidation = &etagValidation
	return &this
}

// NewCustconfOriginPullResumeDownloadWithDefaults instantiates a new CustconfOriginPullResumeDownload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustconfOriginPullResumeDownloadWithDefaults() *CustconfOriginPullResumeDownload {
	this := CustconfOriginPullResumeDownload{}
	var etagValidation OriginPullResumeDownloadEtagValidationEnumWrapperValue = "UNKNOWN"
	this.EtagValidation = &etagValidation
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustconfOriginPullResumeDownload) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustconfOriginPullResumeDownload) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustconfOriginPullResumeDownload) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CustconfOriginPullResumeDownload) SetId(v string) {
	o.Id = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CustconfOriginPullResumeDownload) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustconfOriginPullResumeDownload) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CustconfOriginPullResumeDownload) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CustconfOriginPullResumeDownload) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetOriginalStatusCodeMatch returns the OriginalStatusCodeMatch field value if set, zero value otherwise.
func (o *CustconfOriginPullResumeDownload) GetOriginalStatusCodeMatch() string {
	if o == nil || o.OriginalStatusCodeMatch == nil {
		var ret string
		return ret
	}
	return *o.OriginalStatusCodeMatch
}

// GetOriginalStatusCodeMatchOk returns a tuple with the OriginalStatusCodeMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustconfOriginPullResumeDownload) GetOriginalStatusCodeMatchOk() (*string, bool) {
	if o == nil || o.OriginalStatusCodeMatch == nil {
		return nil, false
	}
	return o.OriginalStatusCodeMatch, true
}

// HasOriginalStatusCodeMatch returns a boolean if a field has been set.
func (o *CustconfOriginPullResumeDownload) HasOriginalStatusCodeMatch() bool {
	if o != nil && o.OriginalStatusCodeMatch != nil {
		return true
	}

	return false
}

// SetOriginalStatusCodeMatch gets a reference to the given string and assigns it to the OriginalStatusCodeMatch field.
func (o *CustconfOriginPullResumeDownload) SetOriginalStatusCodeMatch(v string) {
	o.OriginalStatusCodeMatch = &v
}

// GetMinimumBodyBytesConsumed returns the MinimumBodyBytesConsumed field value if set, zero value otherwise.
func (o *CustconfOriginPullResumeDownload) GetMinimumBodyBytesConsumed() string {
	if o == nil || o.MinimumBodyBytesConsumed == nil {
		var ret string
		return ret
	}
	return *o.MinimumBodyBytesConsumed
}

// GetMinimumBodyBytesConsumedOk returns a tuple with the MinimumBodyBytesConsumed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustconfOriginPullResumeDownload) GetMinimumBodyBytesConsumedOk() (*string, bool) {
	if o == nil || o.MinimumBodyBytesConsumed == nil {
		return nil, false
	}
	return o.MinimumBodyBytesConsumed, true
}

// HasMinimumBodyBytesConsumed returns a boolean if a field has been set.
func (o *CustconfOriginPullResumeDownload) HasMinimumBodyBytesConsumed() bool {
	if o != nil && o.MinimumBodyBytesConsumed != nil {
		return true
	}

	return false
}

// SetMinimumBodyBytesConsumed gets a reference to the given string and assigns it to the MinimumBodyBytesConsumed field.
func (o *CustconfOriginPullResumeDownload) SetMinimumBodyBytesConsumed(v string) {
	o.MinimumBodyBytesConsumed = &v
}

// GetMaximumAttempts returns the MaximumAttempts field value if set, zero value otherwise.
func (o *CustconfOriginPullResumeDownload) GetMaximumAttempts() float32 {
	if o == nil || o.MaximumAttempts == nil {
		var ret float32
		return ret
	}
	return *o.MaximumAttempts
}

// GetMaximumAttemptsOk returns a tuple with the MaximumAttempts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustconfOriginPullResumeDownload) GetMaximumAttemptsOk() (*float32, bool) {
	if o == nil || o.MaximumAttempts == nil {
		return nil, false
	}
	return o.MaximumAttempts, true
}

// HasMaximumAttempts returns a boolean if a field has been set.
func (o *CustconfOriginPullResumeDownload) HasMaximumAttempts() bool {
	if o != nil && o.MaximumAttempts != nil {
		return true
	}

	return false
}

// SetMaximumAttempts gets a reference to the given float32 and assigns it to the MaximumAttempts field.
func (o *CustconfOriginPullResumeDownload) SetMaximumAttempts(v float32) {
	o.MaximumAttempts = &v
}

// GetRequireOriginRangeSupport returns the RequireOriginRangeSupport field value if set, zero value otherwise.
func (o *CustconfOriginPullResumeDownload) GetRequireOriginRangeSupport() bool {
	if o == nil || o.RequireOriginRangeSupport == nil {
		var ret bool
		return ret
	}
	return *o.RequireOriginRangeSupport
}

// GetRequireOriginRangeSupportOk returns a tuple with the RequireOriginRangeSupport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustconfOriginPullResumeDownload) GetRequireOriginRangeSupportOk() (*bool, bool) {
	if o == nil || o.RequireOriginRangeSupport == nil {
		return nil, false
	}
	return o.RequireOriginRangeSupport, true
}

// HasRequireOriginRangeSupport returns a boolean if a field has been set.
func (o *CustconfOriginPullResumeDownload) HasRequireOriginRangeSupport() bool {
	if o != nil && o.RequireOriginRangeSupport != nil {
		return true
	}

	return false
}

// SetRequireOriginRangeSupport gets a reference to the given bool and assigns it to the RequireOriginRangeSupport field.
func (o *CustconfOriginPullResumeDownload) SetRequireOriginRangeSupport(v bool) {
	o.RequireOriginRangeSupport = &v
}

// GetEtagValidation returns the EtagValidation field value if set, zero value otherwise.
func (o *CustconfOriginPullResumeDownload) GetEtagValidation() OriginPullResumeDownloadEtagValidationEnumWrapperValue {
	if o == nil || o.EtagValidation == nil {
		var ret OriginPullResumeDownloadEtagValidationEnumWrapperValue
		return ret
	}
	return *o.EtagValidation
}

// GetEtagValidationOk returns a tuple with the EtagValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustconfOriginPullResumeDownload) GetEtagValidationOk() (*OriginPullResumeDownloadEtagValidationEnumWrapperValue, bool) {
	if o == nil || o.EtagValidation == nil {
		return nil, false
	}
	return o.EtagValidation, true
}

// HasEtagValidation returns a boolean if a field has been set.
func (o *CustconfOriginPullResumeDownload) HasEtagValidation() bool {
	if o != nil && o.EtagValidation != nil {
		return true
	}

	return false
}

// SetEtagValidation gets a reference to the given OriginPullResumeDownloadEtagValidationEnumWrapperValue and assigns it to the EtagValidation field.
func (o *CustconfOriginPullResumeDownload) SetEtagValidation(v OriginPullResumeDownloadEtagValidationEnumWrapperValue) {
	o.EtagValidation = &v
}

// GetHeaderFilter returns the HeaderFilter field value if set, zero value otherwise.
func (o *CustconfOriginPullResumeDownload) GetHeaderFilter() string {
	if o == nil || o.HeaderFilter == nil {
		var ret string
		return ret
	}
	return *o.HeaderFilter
}

// GetHeaderFilterOk returns a tuple with the HeaderFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustconfOriginPullResumeDownload) GetHeaderFilterOk() (*string, bool) {
	if o == nil || o.HeaderFilter == nil {
		return nil, false
	}
	return o.HeaderFilter, true
}

// HasHeaderFilter returns a boolean if a field has been set.
func (o *CustconfOriginPullResumeDownload) HasHeaderFilter() bool {
	if o != nil && o.HeaderFilter != nil {
		return true
	}

	return false
}

// SetHeaderFilter gets a reference to the given string and assigns it to the HeaderFilter field.
func (o *CustconfOriginPullResumeDownload) SetHeaderFilter(v string) {
	o.HeaderFilter = &v
}

// GetPathFilter returns the PathFilter field value if set, zero value otherwise.
func (o *CustconfOriginPullResumeDownload) GetPathFilter() string {
	if o == nil || o.PathFilter == nil {
		var ret string
		return ret
	}
	return *o.PathFilter
}

// GetPathFilterOk returns a tuple with the PathFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustconfOriginPullResumeDownload) GetPathFilterOk() (*string, bool) {
	if o == nil || o.PathFilter == nil {
		return nil, false
	}
	return o.PathFilter, true
}

// HasPathFilter returns a boolean if a field has been set.
func (o *CustconfOriginPullResumeDownload) HasPathFilter() bool {
	if o != nil && o.PathFilter != nil {
		return true
	}

	return false
}

// SetPathFilter gets a reference to the given string and assigns it to the PathFilter field.
func (o *CustconfOriginPullResumeDownload) SetPathFilter(v string) {
	o.PathFilter = &v
}

func (o CustconfOriginPullResumeDownload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.OriginalStatusCodeMatch != nil {
		toSerialize["originalStatusCodeMatch"] = o.OriginalStatusCodeMatch
	}
	if o.MinimumBodyBytesConsumed != nil {
		toSerialize["minimumBodyBytesConsumed"] = o.MinimumBodyBytesConsumed
	}
	if o.MaximumAttempts != nil {
		toSerialize["maximumAttempts"] = o.MaximumAttempts
	}
	if o.RequireOriginRangeSupport != nil {
		toSerialize["requireOriginRangeSupport"] = o.RequireOriginRangeSupport
	}
	if o.EtagValidation != nil {
		toSerialize["etagValidation"] = o.EtagValidation
	}
	if o.HeaderFilter != nil {
		toSerialize["headerFilter"] = o.HeaderFilter
	}
	if o.PathFilter != nil {
		toSerialize["pathFilter"] = o.PathFilter
	}
	return json.Marshal(toSerialize)
}

type NullableCustconfOriginPullResumeDownload struct {
	value *CustconfOriginPullResumeDownload
	isSet bool
}

func (v NullableCustconfOriginPullResumeDownload) Get() *CustconfOriginPullResumeDownload {
	return v.value
}

func (v *NullableCustconfOriginPullResumeDownload) Set(val *CustconfOriginPullResumeDownload) {
	v.value = val
	v.isSet = true
}

func (v NullableCustconfOriginPullResumeDownload) IsSet() bool {
	return v.isSet
}

func (v *NullableCustconfOriginPullResumeDownload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustconfOriginPullResumeDownload(val *CustconfOriginPullResumeDownload) *NullableCustconfOriginPullResumeDownload {
	return &NullableCustconfOriginPullResumeDownload{value: val, isSet: true}
}

func (v NullableCustconfOriginPullResumeDownload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustconfOriginPullResumeDownload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
