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
	"time"
)

// WafEvent A record of an action taken by the WAF  Events are generated by the WAF when an incoming request to a WAF-enabled site triggers a policy or custom rule. Events contain information about the request, the rule that was triggered, and the action taken by the WAF as a result (block the request, allow the request, present the user a captcha, etc).
type WafEvent struct {
	// A WAF event's unique identifier
	Id *string `json:"id,omitempty"`
	// An event's user-facing identifier  Reference IDs are displayed to the end user when the WAF blocks a request to a site. Please note that an event's ID and reference ID are different.
	ReferenceId *string `json:"referenceId,omitempty"`
	// When a WAF event occurred
	EventDate *time.Time `json:"eventDate,omitempty"`
	Request *WafEventRequest `json:"request,omitempty"`
	Action *EventRuleAction `json:"action,omitempty"`
	Client *WafEventNetwork `json:"client,omitempty"`
	// Number of events which matched this
	Count *string `json:"count,omitempty"`
}

// NewWafEvent instantiates a new WafEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafEvent() *WafEvent {
	this := WafEvent{}
	return &this
}

// NewWafEventWithDefaults instantiates a new WafEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafEventWithDefaults() *WafEvent {
	this := WafEvent{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WafEvent) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafEvent) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WafEvent) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WafEvent) SetId(v string) {
	o.Id = &v
}

// GetReferenceId returns the ReferenceId field value if set, zero value otherwise.
func (o *WafEvent) GetReferenceId() string {
	if o == nil || o.ReferenceId == nil {
		var ret string
		return ret
	}
	return *o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafEvent) GetReferenceIdOk() (*string, bool) {
	if o == nil || o.ReferenceId == nil {
		return nil, false
	}
	return o.ReferenceId, true
}

// HasReferenceId returns a boolean if a field has been set.
func (o *WafEvent) HasReferenceId() bool {
	if o != nil && o.ReferenceId != nil {
		return true
	}

	return false
}

// SetReferenceId gets a reference to the given string and assigns it to the ReferenceId field.
func (o *WafEvent) SetReferenceId(v string) {
	o.ReferenceId = &v
}

// GetEventDate returns the EventDate field value if set, zero value otherwise.
func (o *WafEvent) GetEventDate() time.Time {
	if o == nil || o.EventDate == nil {
		var ret time.Time
		return ret
	}
	return *o.EventDate
}

// GetEventDateOk returns a tuple with the EventDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafEvent) GetEventDateOk() (*time.Time, bool) {
	if o == nil || o.EventDate == nil {
		return nil, false
	}
	return o.EventDate, true
}

// HasEventDate returns a boolean if a field has been set.
func (o *WafEvent) HasEventDate() bool {
	if o != nil && o.EventDate != nil {
		return true
	}

	return false
}

// SetEventDate gets a reference to the given time.Time and assigns it to the EventDate field.
func (o *WafEvent) SetEventDate(v time.Time) {
	o.EventDate = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *WafEvent) GetRequest() WafEventRequest {
	if o == nil || o.Request == nil {
		var ret WafEventRequest
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafEvent) GetRequestOk() (*WafEventRequest, bool) {
	if o == nil || o.Request == nil {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *WafEvent) HasRequest() bool {
	if o != nil && o.Request != nil {
		return true
	}

	return false
}

// SetRequest gets a reference to the given WafEventRequest and assigns it to the Request field.
func (o *WafEvent) SetRequest(v WafEventRequest) {
	o.Request = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *WafEvent) GetAction() EventRuleAction {
	if o == nil || o.Action == nil {
		var ret EventRuleAction
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafEvent) GetActionOk() (*EventRuleAction, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *WafEvent) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given EventRuleAction and assigns it to the Action field.
func (o *WafEvent) SetAction(v EventRuleAction) {
	o.Action = &v
}

// GetClient returns the Client field value if set, zero value otherwise.
func (o *WafEvent) GetClient() WafEventNetwork {
	if o == nil || o.Client == nil {
		var ret WafEventNetwork
		return ret
	}
	return *o.Client
}

// GetClientOk returns a tuple with the Client field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafEvent) GetClientOk() (*WafEventNetwork, bool) {
	if o == nil || o.Client == nil {
		return nil, false
	}
	return o.Client, true
}

// HasClient returns a boolean if a field has been set.
func (o *WafEvent) HasClient() bool {
	if o != nil && o.Client != nil {
		return true
	}

	return false
}

// SetClient gets a reference to the given WafEventNetwork and assigns it to the Client field.
func (o *WafEvent) SetClient(v WafEventNetwork) {
	o.Client = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *WafEvent) GetCount() string {
	if o == nil || o.Count == nil {
		var ret string
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafEvent) GetCountOk() (*string, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *WafEvent) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given string and assigns it to the Count field.
func (o *WafEvent) SetCount(v string) {
	o.Count = &v
}

func (o WafEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ReferenceId != nil {
		toSerialize["referenceId"] = o.ReferenceId
	}
	if o.EventDate != nil {
		toSerialize["eventDate"] = o.EventDate
	}
	if o.Request != nil {
		toSerialize["request"] = o.Request
	}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.Client != nil {
		toSerialize["client"] = o.Client
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableWafEvent struct {
	value *WafEvent
	isSet bool
}

func (v NullableWafEvent) Get() *WafEvent {
	return v.value
}

func (v *NullableWafEvent) Set(val *WafEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableWafEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableWafEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWafEvent(val *WafEvent) *NullableWafEvent {
	return &NullableWafEvent{value: val, isSet: true}
}

func (v NullableWafEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWafEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}