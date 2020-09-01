# WafSite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A WAF site&#39;s unique identifier | [optional] 
**Name** | Pointer to **string** | The WAF site&#39;s name, used on stand-alone WAF sites | [optional] 
**Enabled** | Pointer to **bool** | Whether or not a site&#39;s WAF service is enabled | [optional] 
**ApiUrls** | Pointer to **[]string** | A list of API URLs which receive different processing through the WAF than website requests | [optional] 
**CreatedAt** | Pointer to [**time.Time**](time.Time.md) | The date a WAF site was created | [optional] 
**UpdatedAt** | Pointer to [**time.Time**](time.Time.md) | The date a WAF site was last updated | [optional] 
**Status** | Pointer to [**WafSiteStatus**](wafSiteStatus.md) |  | [optional] [default to "ACTIVE"]
**Mode** | Pointer to [**SiteAttachMode**](SiteAttachMode.md) |  | [optional] [default to "STANDALONE"]
**DeliveryId** | Pointer to **string** | For an ATTACHED site, the CDN site ID where caching can be enabled | [optional] 
**Type** | Pointer to [**WafSiteType**](wafSiteType.md) |  | [optional] [default to "UNKNOWN_TYPE"]
**Monitoring** | Pointer to **bool** | Whether or not a site&#39;s WAF service is in a monitor state  WAF sites in monitoring mode accept incoming requests and log but take no action on policy and rule violations. | [optional] 

## Methods

### NewWafSite

`func NewWafSite() *WafSite`

NewWafSite instantiates a new WafSite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafSiteWithDefaults

`func NewWafSiteWithDefaults() *WafSite`

NewWafSiteWithDefaults instantiates a new WafSite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WafSite) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WafSite) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WafSite) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WafSite) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *WafSite) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WafSite) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WafSite) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WafSite) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *WafSite) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WafSite) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WafSite) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WafSite) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetApiUrls

`func (o *WafSite) GetApiUrls() []string`

GetApiUrls returns the ApiUrls field if non-nil, zero value otherwise.

### GetApiUrlsOk

`func (o *WafSite) GetApiUrlsOk() (*[]string, bool)`

GetApiUrlsOk returns a tuple with the ApiUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUrls

`func (o *WafSite) SetApiUrls(v []string)`

SetApiUrls sets ApiUrls field to given value.

### HasApiUrls

`func (o *WafSite) HasApiUrls() bool`

HasApiUrls returns a boolean if a field has been set.

### GetCreatedAt

`func (o *WafSite) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WafSite) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WafSite) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *WafSite) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *WafSite) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WafSite) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WafSite) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *WafSite) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetStatus

`func (o *WafSite) GetStatus() WafSiteStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WafSite) GetStatusOk() (*WafSiteStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WafSite) SetStatus(v WafSiteStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WafSite) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMode

`func (o *WafSite) GetMode() SiteAttachMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *WafSite) GetModeOk() (*SiteAttachMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *WafSite) SetMode(v SiteAttachMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *WafSite) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetDeliveryId

`func (o *WafSite) GetDeliveryId() string`

GetDeliveryId returns the DeliveryId field if non-nil, zero value otherwise.

### GetDeliveryIdOk

`func (o *WafSite) GetDeliveryIdOk() (*string, bool)`

GetDeliveryIdOk returns a tuple with the DeliveryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryId

`func (o *WafSite) SetDeliveryId(v string)`

SetDeliveryId sets DeliveryId field to given value.

### HasDeliveryId

`func (o *WafSite) HasDeliveryId() bool`

HasDeliveryId returns a boolean if a field has been set.

### GetType

`func (o *WafSite) GetType() WafSiteType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WafSite) GetTypeOk() (*WafSiteType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WafSite) SetType(v WafSiteType)`

SetType sets Type field to given value.

### HasType

`func (o *WafSite) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMonitoring

`func (o *WafSite) GetMonitoring() bool`

GetMonitoring returns the Monitoring field if non-nil, zero value otherwise.

### GetMonitoringOk

`func (o *WafSite) GetMonitoringOk() (*bool, bool)`

GetMonitoringOk returns a tuple with the Monitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoring

`func (o *WafSite) SetMonitoring(v bool)`

SetMonitoring sets Monitoring field to given value.

### HasMonitoring

`func (o *WafSite) HasMonitoring() bool`

HasMonitoring returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


