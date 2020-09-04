# CdnSite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A CDN site&#39;s unique identifier | [optional] 
**StackId** | Pointer to **string** | The ID of the stack to which a CDN site belongs | [optional] 
**Label** | Pointer to **string** | A CDN site&#39;s name  Site names correspond to their fully-qualified domain name. | [optional] 
**Status** | Pointer to **string** | A CDN site&#39;s internal state  Site status is controlled by StackPath as sites are provisioned and managed by StackPath&#39;s accounting and security teams. | [optional] 
**CreatedAt** | Pointer to [**time.Time**](time.Time.md) | The date that a CDN site was created | [optional] 
**UpdatedAt** | Pointer to [**time.Time**](time.Time.md) | The date that a CDN site was last updated | [optional] 
**Features** | Pointer to [**[]CdnSiteFeature**](cdnSiteFeature.md) | A CDN site&#39;s associated features  Features control how StackPath provisions and configures a site. | [optional] 
**Enabled** | Pointer to **bool** | Whether or not a site&#39;s CDN service is enabled | [optional] 
**Type** | Pointer to [**SiteTypeValue**](SiteTypeValue.md) |  | [optional] [default to "UNKNOWN"]

## Methods

### NewCdnSite

`func NewCdnSite() *CdnSite`

NewCdnSite instantiates a new CdnSite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdnSiteWithDefaults

`func NewCdnSiteWithDefaults() *CdnSite`

NewCdnSiteWithDefaults instantiates a new CdnSite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CdnSite) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CdnSite) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CdnSite) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CdnSite) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStackId

`func (o *CdnSite) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *CdnSite) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *CdnSite) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *CdnSite) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetLabel

`func (o *CdnSite) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CdnSite) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CdnSite) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CdnSite) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetStatus

`func (o *CdnSite) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CdnSite) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CdnSite) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CdnSite) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CdnSite) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CdnSite) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CdnSite) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CdnSite) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CdnSite) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CdnSite) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CdnSite) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CdnSite) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetFeatures

`func (o *CdnSite) GetFeatures() []CdnSiteFeature`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *CdnSite) GetFeaturesOk() (*[]CdnSiteFeature, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *CdnSite) SetFeatures(v []CdnSiteFeature)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *CdnSite) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetEnabled

`func (o *CdnSite) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CdnSite) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CdnSite) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CdnSite) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetType

`func (o *CdnSite) GetType() SiteTypeValue`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CdnSite) GetTypeOk() (*SiteTypeValue, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CdnSite) SetType(v SiteTypeValue)`

SetType sets Type field to given value.

### HasType

`func (o *CdnSite) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


