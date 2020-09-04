# DeliverySite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A site&#39;s unique identifier | [optional] 
**StackId** | Pointer to **string** | The ID of the stack to which a site belongs | [optional] 
**Label** | Pointer to **string** | The site&#39;s name | [optional] 
**Status** | Pointer to **string** | The site&#39;s status | [optional] 
**Features** | Pointer to [**[]DeliverySiteFeature**](deliverySiteFeature.md) | A site&#39;s features  Site features control how its content is delivered to end users. | [optional] 
**CreatedAt** | Pointer to [**time.Time**](time.Time.md) | The date a site was created | [optional] 
**UpdatedAt** | Pointer to [**time.Time**](time.Time.md) | The date a site was last updated | [optional] 
**ApiUrls** | Pointer to **[]string** | A list of API URLs which receive different processing through the WAF than website requests | [optional] 
**Monitoring** | Pointer to **bool** | Whether or not a site&#39;s WAF service is in a monitor state  Sites in WAF monitoring mode accept incoming requests and log but take no action on policy and rule violations. | [optional] 

## Methods

### NewDeliverySite

`func NewDeliverySite() *DeliverySite`

NewDeliverySite instantiates a new DeliverySite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliverySiteWithDefaults

`func NewDeliverySiteWithDefaults() *DeliverySite`

NewDeliverySiteWithDefaults instantiates a new DeliverySite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeliverySite) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeliverySite) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeliverySite) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeliverySite) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStackId

`func (o *DeliverySite) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *DeliverySite) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *DeliverySite) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *DeliverySite) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetLabel

`func (o *DeliverySite) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *DeliverySite) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *DeliverySite) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *DeliverySite) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetStatus

`func (o *DeliverySite) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeliverySite) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeliverySite) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeliverySite) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetFeatures

`func (o *DeliverySite) GetFeatures() []DeliverySiteFeature`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *DeliverySite) GetFeaturesOk() (*[]DeliverySiteFeature, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *DeliverySite) SetFeatures(v []DeliverySiteFeature)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *DeliverySite) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DeliverySite) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DeliverySite) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DeliverySite) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DeliverySite) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DeliverySite) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DeliverySite) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DeliverySite) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DeliverySite) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetApiUrls

`func (o *DeliverySite) GetApiUrls() []string`

GetApiUrls returns the ApiUrls field if non-nil, zero value otherwise.

### GetApiUrlsOk

`func (o *DeliverySite) GetApiUrlsOk() (*[]string, bool)`

GetApiUrlsOk returns a tuple with the ApiUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUrls

`func (o *DeliverySite) SetApiUrls(v []string)`

SetApiUrls sets ApiUrls field to given value.

### HasApiUrls

`func (o *DeliverySite) HasApiUrls() bool`

HasApiUrls returns a boolean if a field has been set.

### GetMonitoring

`func (o *DeliverySite) GetMonitoring() bool`

GetMonitoring returns the Monitoring field if non-nil, zero value otherwise.

### GetMonitoringOk

`func (o *DeliverySite) GetMonitoringOk() (*bool, bool)`

GetMonitoringOk returns a tuple with the Monitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoring

`func (o *DeliverySite) SetMonitoring(v bool)`

SetMonitoring sets Monitoring field to given value.

### HasMonitoring

`func (o *DeliverySite) HasMonitoring() bool`

HasMonitoring returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


