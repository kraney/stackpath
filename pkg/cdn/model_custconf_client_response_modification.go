/*
 * Content Delivery Network
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package cdn

import (
	"encoding/json"
)

// CustconfClientResponseModification struct for CustconfClientResponseModification
type CustconfClientResponseModification struct {
	// This is used by the API to perform conflict checking
	Id *string `json:"id,omitempty"`
	StatusCodeRewrite *float32 `json:"statusCodeRewrite,omitempty"`
	HeaderPattern *string `json:"headerPattern,omitempty"`
	HeaderRewrite *string `json:"headerRewrite,omitempty"`
	// String of values delimited by a '|' character. This is the static HTTP header you want inserted into the CDN response. Start value with \"append:\", \"replace:\" or \"create:\" which defines if Header will be added, replaced or create if not exists. Default is append.
	AddHeaders *string `json:"addHeaders,omitempty"`
	FlowControl *CustconfClientResponseModificationFlowControlEnumWrapperValue `json:"flowControl,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	// String of values delimited by a ',' character.
	MethodFilter *string `json:"methodFilter,omitempty"`
	// String of values delimited by a ',' character.
	PathFilter *string `json:"pathFilter,omitempty"`
	// String of values delimited by a ',' character.
	HeaderFilter *string `json:"headerFilter,omitempty"`
}

// NewCustconfClientResponseModification instantiates a new CustconfClientResponseModification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustconfClientResponseModification() *CustconfClientResponseModification {
	this := CustconfClientResponseModification{}
	var flowControl CustconfClientResponseModificationFlowControlEnumWrapperValue = "UNKNOWN"
	this.FlowControl = &flowControl
	return &this
}

// NewCustconfClientResponseModificationWithDefaults instantiates a new CustconfClientResponseModification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustconfClientResponseModificationWithDefaults() *CustconfClientResponseModification {
	this := CustconfClientResponseModification{}
	var flowControl CustconfClientResponseModificationFlowControlEnumWrapperValue = "UNKNOWN"
	this.FlowControl = &flowControl
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustconfClientResponseModification) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustconfClientResponseModification) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustconfClientResponseModification) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CustconfClientResponseModification) SetId(v string) {
	o.Id = &v
}

// GetStatusCodeRewrite returns the StatusCodeRewrite field value if set, zero value otherwise.
func (o *CustconfClientResponseModification) GetStatusCodeRewrite() float32 {
	if o == nil || o.StatusCodeRewrite == nil {
		var ret float32
		return ret
	}
	return *o.StatusCodeRewrite
}

// GetStatusCodeRewriteOk returns a tuple with the StatusCodeRewrite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustconfClientResponseModification) GetStatusCodeRewriteOk() (*float32, bool) {
	if o == nil || o.StatusCodeRewrite == nil {
		return nil, false
	}
	return o.StatusCodeRewrite, true
}

// HasStatusCodeRewrite returns a boolean if a field has been set.
func (o *CustconfClientResponseModification) HasStatusCodeRewrite() bool {
	if o != nil && o.StatusCodeRewrite != nil {
		return true
	}

	return false
}

// SetStatusCodeRewrite gets a reference to the given float32 and assigns it to the StatusCodeRewrite field.
func (o *CustconfClientResponseModification) SetStatusCodeRewrite(v float32) {
	o.StatusCodeRewrite = &v
}

// GetHeaderPattern returns the HeaderPattern field value if set, zero value otherwise.
func (o *CustconfClientResponseModification) GetHeaderPattern() string {
	if o == nil || o.HeaderPattern == nil {
		var ret string
		return ret
	}
	return *o.HeaderPattern
}

// GetHeaderPatternOk returns a tuple with the HeaderPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustconfClientResponseModification) GetHeaderPatternOk() (*string, bool) {
	if o == nil || o.HeaderPattern == nil {
		return nil, false
	}
	return o.HeaderPattern, true
}

// HasHeaderPattern returns a boolean if a field has been set.
func (o *CustconfClientResponseModification) HasHeaderPattern() bool {
	if o != nil && o.HeaderPattern != nil {
		return true
	}

	return false
}

// SetHeaderPattern gets a reference to the given string and assigns it to the HeaderPattern field.
func (o *CustconfClientResponseModification) SetHeaderPattern(v string) {
	o.HeaderPattern = &v
}

// GetHeaderRewrite returns the HeaderRewrite field value if set, zero value otherwise.
func (o *CustconfClientResponseModification) GetHeaderRewrite() string {
	if o == nil || o.HeaderRewrite == nil {
		var ret string
		return ret
	}
	return *o.HeaderRewrite
}

// GetHeaderRewriteOk returns a tuple with the HeaderRewrite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustconfClientResponseModification) GetHeaderRewriteOk() (*string, bool) {
	if o == nil || o.HeaderRewrite == nil {
		return nil, false
	}
	return o.HeaderRewrite, true
}

// HasHeaderRewrite returns a boolean if a field has been set.
func (o *CustconfClientResponseModification) HasHeaderRewrite() bool {
	if o != nil && o.HeaderRewrite != nil {
		return true
	}

	return false
}

// SetHeaderRewrite gets a reference to the given string and assigns it to the HeaderRewrite field.
func (o *CustconfClientResponseModification) SetHeaderRewrite(v string) {
	o.HeaderRewrite = &v
}

// GetAddHeaders returns the AddHeaders field value if set, zero value otherwise.
func (o *CustconfClientResponseModification) GetAddHeaders() string {
	if o == nil || o.AddHeaders == nil {
		var ret string
		return ret
	}
	return *o.AddHeaders
}

// GetAddHeadersOk returns a tuple with the AddHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustconfClientResponseModification) GetAddHeadersOk() (*string, bool) {
	if o == nil || o.AddHeaders == nil {
		return nil, false
	}
	return o.AddHeaders, true
}

// HasAddHeaders returns a boolean if a field has been set.
func (o *CustconfClientResponseModification) HasAddHeaders() bool {
	if o != nil && o.AddHeaders != nil {
		return true
	}

	return false
}

// SetAddHeaders gets a reference to the given string and assigns it to the AddHeaders field.
func (o *CustconfClientResponseModification) SetAddHeaders(v string) {
	o.AddHeaders = &v
}

// GetFlowControl returns the FlowControl field value if set, zero value otherwise.
func (o *CustconfClientResponseModification) GetFlowControl() CustconfClientResponseModificationFlowControlEnumWrapperValue {
	if o == nil || o.FlowControl == nil {
		var ret CustconfClientResponseModificationFlowControlEnumWrapperValue
		return ret
	}
	return *o.FlowControl
}

// GetFlowControlOk returns a tuple with the FlowControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustconfClientResponseModification) GetFlowControlOk() (*CustconfClientResponseModificationFlowControlEnumWrapperValue, bool) {
	if o == nil || o.FlowControl == nil {
		return nil, false
	}
	return o.FlowControl, true
}

// HasFlowControl returns a boolean if a field has been set.
func (o *CustconfClientResponseModification) HasFlowControl() bool {
	if o != nil && o.FlowControl != nil {
		return true
	}

	return false
}

// SetFlowControl gets a reference to the given CustconfClientResponseModificationFlowControlEnumWrapperValue and assigns it to the FlowControl field.
func (o *CustconfClientResponseModification) SetFlowControl(v CustconfClientResponseModificationFlowControlEnumWrapperValue) {
	o.FlowControl = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CustconfClientResponseModification) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustconfClientResponseModification) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CustconfClientResponseModification) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CustconfClientResponseModification) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetMethodFilter returns the MethodFilter field value if set, zero value otherwise.
func (o *CustconfClientResponseModification) GetMethodFilter() string {
	if o == nil || o.MethodFilter == nil {
		var ret string
		return ret
	}
	return *o.MethodFilter
}

// GetMethodFilterOk returns a tuple with the MethodFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustconfClientResponseModification) GetMethodFilterOk() (*string, bool) {
	if o == nil || o.MethodFilter == nil {
		return nil, false
	}
	return o.MethodFilter, true
}

// HasMethodFilter returns a boolean if a field has been set.
func (o *CustconfClientResponseModification) HasMethodFilter() bool {
	if o != nil && o.MethodFilter != nil {
		return true
	}

	return false
}

// SetMethodFilter gets a reference to the given string and assigns it to the MethodFilter field.
func (o *CustconfClientResponseModification) SetMethodFilter(v string) {
	o.MethodFilter = &v
}

// GetPathFilter returns the PathFilter field value if set, zero value otherwise.
func (o *CustconfClientResponseModification) GetPathFilter() string {
	if o == nil || o.PathFilter == nil {
		var ret string
		return ret
	}
	return *o.PathFilter
}

// GetPathFilterOk returns a tuple with the PathFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustconfClientResponseModification) GetPathFilterOk() (*string, bool) {
	if o == nil || o.PathFilter == nil {
		return nil, false
	}
	return o.PathFilter, true
}

// HasPathFilter returns a boolean if a field has been set.
func (o *CustconfClientResponseModification) HasPathFilter() bool {
	if o != nil && o.PathFilter != nil {
		return true
	}

	return false
}

// SetPathFilter gets a reference to the given string and assigns it to the PathFilter field.
func (o *CustconfClientResponseModification) SetPathFilter(v string) {
	o.PathFilter = &v
}

// GetHeaderFilter returns the HeaderFilter field value if set, zero value otherwise.
func (o *CustconfClientResponseModification) GetHeaderFilter() string {
	if o == nil || o.HeaderFilter == nil {
		var ret string
		return ret
	}
	return *o.HeaderFilter
}

// GetHeaderFilterOk returns a tuple with the HeaderFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustconfClientResponseModification) GetHeaderFilterOk() (*string, bool) {
	if o == nil || o.HeaderFilter == nil {
		return nil, false
	}
	return o.HeaderFilter, true
}

// HasHeaderFilter returns a boolean if a field has been set.
func (o *CustconfClientResponseModification) HasHeaderFilter() bool {
	if o != nil && o.HeaderFilter != nil {
		return true
	}

	return false
}

// SetHeaderFilter gets a reference to the given string and assigns it to the HeaderFilter field.
func (o *CustconfClientResponseModification) SetHeaderFilter(v string) {
	o.HeaderFilter = &v
}

func (o CustconfClientResponseModification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.StatusCodeRewrite != nil {
		toSerialize["statusCodeRewrite"] = o.StatusCodeRewrite
	}
	if o.HeaderPattern != nil {
		toSerialize["headerPattern"] = o.HeaderPattern
	}
	if o.HeaderRewrite != nil {
		toSerialize["headerRewrite"] = o.HeaderRewrite
	}
	if o.AddHeaders != nil {
		toSerialize["addHeaders"] = o.AddHeaders
	}
	if o.FlowControl != nil {
		toSerialize["flowControl"] = o.FlowControl
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.MethodFilter != nil {
		toSerialize["methodFilter"] = o.MethodFilter
	}
	if o.PathFilter != nil {
		toSerialize["pathFilter"] = o.PathFilter
	}
	if o.HeaderFilter != nil {
		toSerialize["headerFilter"] = o.HeaderFilter
	}
	return json.Marshal(toSerialize)
}

type NullableCustconfClientResponseModification struct {
	value *CustconfClientResponseModification
	isSet bool
}

func (v NullableCustconfClientResponseModification) Get() *CustconfClientResponseModification {
	return v.value
}

func (v *NullableCustconfClientResponseModification) Set(val *CustconfClientResponseModification) {
	v.value = val
	v.isSet = true
}

func (v NullableCustconfClientResponseModification) IsSet() bool {
	return v.isSet
}

func (v *NullableCustconfClientResponseModification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustconfClientResponseModification(val *CustconfClientResponseModification) *NullableCustconfClientResponseModification {
	return &NullableCustconfClientResponseModification{value: val, isSet: true}
}

func (v NullableCustconfClientResponseModification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustconfClientResponseModification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}