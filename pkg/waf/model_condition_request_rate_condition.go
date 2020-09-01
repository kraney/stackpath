/*
 * Web Application Firewall
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package waf

import (
	"encoding/json"
)

// ConditionRequestRateCondition Match the rate at which requests come in that match certain conditions
type ConditionRequestRateCondition struct {
	// A list of source IPs that can trigger a request rate condition
	Ips *[]string `json:"ips,omitempty"`
	// Possible HTTP request methods that can trigger a request rate condition
	HttpMethods *[]WafHttpMethod `json:"httpMethods,omitempty"`
	// A regular expression matching the URL path of the incoming request
	PathPattern *string `json:"pathPattern,omitempty"`
	// The number of incoming requests over the given time that can trigger a request rate condition
	Requests *string `json:"requests,omitempty"`
	// The number of seconds that the WAF measures incoming requests over before triggering a request rate condition
	Time *string `json:"time,omitempty"`
}

// NewConditionRequestRateCondition instantiates a new ConditionRequestRateCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConditionRequestRateCondition() *ConditionRequestRateCondition {
	this := ConditionRequestRateCondition{}
	return &this
}

// NewConditionRequestRateConditionWithDefaults instantiates a new ConditionRequestRateCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConditionRequestRateConditionWithDefaults() *ConditionRequestRateCondition {
	this := ConditionRequestRateCondition{}
	return &this
}

// GetIps returns the Ips field value if set, zero value otherwise.
func (o *ConditionRequestRateCondition) GetIps() []string {
	if o == nil || o.Ips == nil {
		var ret []string
		return ret
	}
	return *o.Ips
}

// GetIpsOk returns a tuple with the Ips field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionRequestRateCondition) GetIpsOk() (*[]string, bool) {
	if o == nil || o.Ips == nil {
		return nil, false
	}
	return o.Ips, true
}

// HasIps returns a boolean if a field has been set.
func (o *ConditionRequestRateCondition) HasIps() bool {
	if o != nil && o.Ips != nil {
		return true
	}

	return false
}

// SetIps gets a reference to the given []string and assigns it to the Ips field.
func (o *ConditionRequestRateCondition) SetIps(v []string) {
	o.Ips = &v
}

// GetHttpMethods returns the HttpMethods field value if set, zero value otherwise.
func (o *ConditionRequestRateCondition) GetHttpMethods() []WafHttpMethod {
	if o == nil || o.HttpMethods == nil {
		var ret []WafHttpMethod
		return ret
	}
	return *o.HttpMethods
}

// GetHttpMethodsOk returns a tuple with the HttpMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionRequestRateCondition) GetHttpMethodsOk() (*[]WafHttpMethod, bool) {
	if o == nil || o.HttpMethods == nil {
		return nil, false
	}
	return o.HttpMethods, true
}

// HasHttpMethods returns a boolean if a field has been set.
func (o *ConditionRequestRateCondition) HasHttpMethods() bool {
	if o != nil && o.HttpMethods != nil {
		return true
	}

	return false
}

// SetHttpMethods gets a reference to the given []WafHttpMethod and assigns it to the HttpMethods field.
func (o *ConditionRequestRateCondition) SetHttpMethods(v []WafHttpMethod) {
	o.HttpMethods = &v
}

// GetPathPattern returns the PathPattern field value if set, zero value otherwise.
func (o *ConditionRequestRateCondition) GetPathPattern() string {
	if o == nil || o.PathPattern == nil {
		var ret string
		return ret
	}
	return *o.PathPattern
}

// GetPathPatternOk returns a tuple with the PathPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionRequestRateCondition) GetPathPatternOk() (*string, bool) {
	if o == nil || o.PathPattern == nil {
		return nil, false
	}
	return o.PathPattern, true
}

// HasPathPattern returns a boolean if a field has been set.
func (o *ConditionRequestRateCondition) HasPathPattern() bool {
	if o != nil && o.PathPattern != nil {
		return true
	}

	return false
}

// SetPathPattern gets a reference to the given string and assigns it to the PathPattern field.
func (o *ConditionRequestRateCondition) SetPathPattern(v string) {
	o.PathPattern = &v
}

// GetRequests returns the Requests field value if set, zero value otherwise.
func (o *ConditionRequestRateCondition) GetRequests() string {
	if o == nil || o.Requests == nil {
		var ret string
		return ret
	}
	return *o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionRequestRateCondition) GetRequestsOk() (*string, bool) {
	if o == nil || o.Requests == nil {
		return nil, false
	}
	return o.Requests, true
}

// HasRequests returns a boolean if a field has been set.
func (o *ConditionRequestRateCondition) HasRequests() bool {
	if o != nil && o.Requests != nil {
		return true
	}

	return false
}

// SetRequests gets a reference to the given string and assigns it to the Requests field.
func (o *ConditionRequestRateCondition) SetRequests(v string) {
	o.Requests = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *ConditionRequestRateCondition) GetTime() string {
	if o == nil || o.Time == nil {
		var ret string
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionRequestRateCondition) GetTimeOk() (*string, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *ConditionRequestRateCondition) HasTime() bool {
	if o != nil && o.Time != nil {
		return true
	}

	return false
}

// SetTime gets a reference to the given string and assigns it to the Time field.
func (o *ConditionRequestRateCondition) SetTime(v string) {
	o.Time = &v
}

func (o ConditionRequestRateCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ips != nil {
		toSerialize["ips"] = o.Ips
	}
	if o.HttpMethods != nil {
		toSerialize["httpMethods"] = o.HttpMethods
	}
	if o.PathPattern != nil {
		toSerialize["pathPattern"] = o.PathPattern
	}
	if o.Requests != nil {
		toSerialize["requests"] = o.Requests
	}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	return json.Marshal(toSerialize)
}

type NullableConditionRequestRateCondition struct {
	value *ConditionRequestRateCondition
	isSet bool
}

func (v NullableConditionRequestRateCondition) Get() *ConditionRequestRateCondition {
	return v.value
}

func (v *NullableConditionRequestRateCondition) Set(val *ConditionRequestRateCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableConditionRequestRateCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableConditionRequestRateCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConditionRequestRateCondition(val *ConditionRequestRateCondition) *NullableConditionRequestRateCondition {
	return &NullableConditionRequestRateCondition{value: val, isSet: true}
}

func (v NullableConditionRequestRateCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConditionRequestRateCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}