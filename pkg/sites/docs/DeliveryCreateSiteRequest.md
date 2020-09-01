# DeliveryCreateSiteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** | The site&#39;s domain name | [optional] 
**Origin** | Pointer to [**DeliveryCreateSiteRequestOrigin**](deliveryCreateSiteRequestOrigin.md) |  | [optional] 
**Features** | Pointer to [**[]DeliveryCreateSiteRequestFeature**](deliveryCreateSiteRequestFeature.md) | A list of features desired on the site | [optional] 
**Configuration** | Pointer to [**CustconfConfiguration**](custconfConfiguration.md) |  | [optional] 

## Methods

### NewDeliveryCreateSiteRequest

`func NewDeliveryCreateSiteRequest() *DeliveryCreateSiteRequest`

NewDeliveryCreateSiteRequest instantiates a new DeliveryCreateSiteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryCreateSiteRequestWithDefaults

`func NewDeliveryCreateSiteRequestWithDefaults() *DeliveryCreateSiteRequest`

NewDeliveryCreateSiteRequestWithDefaults instantiates a new DeliveryCreateSiteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *DeliveryCreateSiteRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DeliveryCreateSiteRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DeliveryCreateSiteRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *DeliveryCreateSiteRequest) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetOrigin

`func (o *DeliveryCreateSiteRequest) GetOrigin() DeliveryCreateSiteRequestOrigin`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *DeliveryCreateSiteRequest) GetOriginOk() (*DeliveryCreateSiteRequestOrigin, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *DeliveryCreateSiteRequest) SetOrigin(v DeliveryCreateSiteRequestOrigin)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *DeliveryCreateSiteRequest) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetFeatures

`func (o *DeliveryCreateSiteRequest) GetFeatures() []DeliveryCreateSiteRequestFeature`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *DeliveryCreateSiteRequest) GetFeaturesOk() (*[]DeliveryCreateSiteRequestFeature, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *DeliveryCreateSiteRequest) SetFeatures(v []DeliveryCreateSiteRequestFeature)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *DeliveryCreateSiteRequest) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetConfiguration

`func (o *DeliveryCreateSiteRequest) GetConfiguration() CustconfConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *DeliveryCreateSiteRequest) GetConfigurationOk() (*CustconfConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *DeliveryCreateSiteRequest) SetConfiguration(v CustconfConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *DeliveryCreateSiteRequest) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


