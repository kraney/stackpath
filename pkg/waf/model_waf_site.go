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

// WafSite A site's WAF service's general settings
type WafSite struct {
	// A WAF site's unique identifier
	Id *string `json:"id,omitempty"`
	// The WAF site's name, used on stand-alone WAF sites
	Name *string `json:"name,omitempty"`
	// Whether or not a site's WAF service is enabled
	Enabled *bool `json:"enabled,omitempty"`
	// A list of API URLs which receive different processing through the WAF than website requests
	ApiUrls *[]string `json:"apiUrls,omitempty"`
	// The date a WAF site was created
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The date a WAF site was last updated
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	Status *WafSiteStatus `json:"status,omitempty"`
	Mode *SiteAttachMode `json:"mode,omitempty"`
	// For an ATTACHED site, the CDN site ID where caching can be enabled
	DeliveryId *string `json:"deliveryId,omitempty"`
	Type *WafSiteType `json:"type,omitempty"`
	// Whether or not a site's WAF service is in a monitor state  WAF sites in monitoring mode accept incoming requests and log but take no action on policy and rule violations.
	Monitoring *bool `json:"monitoring,omitempty"`
}

// NewWafSite instantiates a new WafSite object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafSite() *WafSite {
	this := WafSite{}
	var status WafSiteStatus = "ACTIVE"
	this.Status = &status
	var mode SiteAttachMode = "STANDALONE"
	this.Mode = &mode
	var type_ WafSiteType = "UNKNOWN_TYPE"
	this.Type = &type_
	return &this
}

// NewWafSiteWithDefaults instantiates a new WafSite object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafSiteWithDefaults() *WafSite {
	this := WafSite{}
	var status WafSiteStatus = "ACTIVE"
	this.Status = &status
	var mode SiteAttachMode = "STANDALONE"
	this.Mode = &mode
	var type_ WafSiteType = "UNKNOWN_TYPE"
	this.Type = &type_
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WafSite) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafSite) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WafSite) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WafSite) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WafSite) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafSite) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WafSite) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WafSite) SetName(v string) {
	o.Name = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *WafSite) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafSite) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *WafSite) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *WafSite) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetApiUrls returns the ApiUrls field value if set, zero value otherwise.
func (o *WafSite) GetApiUrls() []string {
	if o == nil || o.ApiUrls == nil {
		var ret []string
		return ret
	}
	return *o.ApiUrls
}

// GetApiUrlsOk returns a tuple with the ApiUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafSite) GetApiUrlsOk() (*[]string, bool) {
	if o == nil || o.ApiUrls == nil {
		return nil, false
	}
	return o.ApiUrls, true
}

// HasApiUrls returns a boolean if a field has been set.
func (o *WafSite) HasApiUrls() bool {
	if o != nil && o.ApiUrls != nil {
		return true
	}

	return false
}

// SetApiUrls gets a reference to the given []string and assigns it to the ApiUrls field.
func (o *WafSite) SetApiUrls(v []string) {
	o.ApiUrls = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *WafSite) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafSite) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *WafSite) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *WafSite) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *WafSite) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafSite) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *WafSite) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *WafSite) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WafSite) GetStatus() WafSiteStatus {
	if o == nil || o.Status == nil {
		var ret WafSiteStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafSite) GetStatusOk() (*WafSiteStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WafSite) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given WafSiteStatus and assigns it to the Status field.
func (o *WafSite) SetStatus(v WafSiteStatus) {
	o.Status = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *WafSite) GetMode() SiteAttachMode {
	if o == nil || o.Mode == nil {
		var ret SiteAttachMode
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafSite) GetModeOk() (*SiteAttachMode, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *WafSite) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given SiteAttachMode and assigns it to the Mode field.
func (o *WafSite) SetMode(v SiteAttachMode) {
	o.Mode = &v
}

// GetDeliveryId returns the DeliveryId field value if set, zero value otherwise.
func (o *WafSite) GetDeliveryId() string {
	if o == nil || o.DeliveryId == nil {
		var ret string
		return ret
	}
	return *o.DeliveryId
}

// GetDeliveryIdOk returns a tuple with the DeliveryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafSite) GetDeliveryIdOk() (*string, bool) {
	if o == nil || o.DeliveryId == nil {
		return nil, false
	}
	return o.DeliveryId, true
}

// HasDeliveryId returns a boolean if a field has been set.
func (o *WafSite) HasDeliveryId() bool {
	if o != nil && o.DeliveryId != nil {
		return true
	}

	return false
}

// SetDeliveryId gets a reference to the given string and assigns it to the DeliveryId field.
func (o *WafSite) SetDeliveryId(v string) {
	o.DeliveryId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WafSite) GetType() WafSiteType {
	if o == nil || o.Type == nil {
		var ret WafSiteType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafSite) GetTypeOk() (*WafSiteType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WafSite) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given WafSiteType and assigns it to the Type field.
func (o *WafSite) SetType(v WafSiteType) {
	o.Type = &v
}

// GetMonitoring returns the Monitoring field value if set, zero value otherwise.
func (o *WafSite) GetMonitoring() bool {
	if o == nil || o.Monitoring == nil {
		var ret bool
		return ret
	}
	return *o.Monitoring
}

// GetMonitoringOk returns a tuple with the Monitoring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafSite) GetMonitoringOk() (*bool, bool) {
	if o == nil || o.Monitoring == nil {
		return nil, false
	}
	return o.Monitoring, true
}

// HasMonitoring returns a boolean if a field has been set.
func (o *WafSite) HasMonitoring() bool {
	if o != nil && o.Monitoring != nil {
		return true
	}

	return false
}

// SetMonitoring gets a reference to the given bool and assigns it to the Monitoring field.
func (o *WafSite) SetMonitoring(v bool) {
	o.Monitoring = &v
}

func (o WafSite) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.ApiUrls != nil {
		toSerialize["apiUrls"] = o.ApiUrls
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	if o.DeliveryId != nil {
		toSerialize["deliveryId"] = o.DeliveryId
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Monitoring != nil {
		toSerialize["monitoring"] = o.Monitoring
	}
	return json.Marshal(toSerialize)
}

type NullableWafSite struct {
	value *WafSite
	isSet bool
}

func (v NullableWafSite) Get() *WafSite {
	return v.value
}

func (v *NullableWafSite) Set(val *WafSite) {
	v.value = val
	v.isSet = true
}

func (v NullableWafSite) IsSet() bool {
	return v.isSet
}

func (v *NullableWafSite) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWafSite(val *WafSite) *NullableWafSite {
	return &NullableWafSite{value: val, isSet: true}
}

func (v NullableWafSite) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWafSite) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
