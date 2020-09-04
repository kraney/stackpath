# CdnCreateScopeRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The EdgeRule&#39;s name | [optional] 
**Slug** | Pointer to **string** | A programmatic name for the EdgeRule | [optional] 
**Configuration** | Pointer to [**CustconfConfiguration**](custconfConfiguration.md) |  | [optional] 

## Methods

### NewCdnCreateScopeRuleRequest

`func NewCdnCreateScopeRuleRequest() *CdnCreateScopeRuleRequest`

NewCdnCreateScopeRuleRequest instantiates a new CdnCreateScopeRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdnCreateScopeRuleRequestWithDefaults

`func NewCdnCreateScopeRuleRequestWithDefaults() *CdnCreateScopeRuleRequest`

NewCdnCreateScopeRuleRequestWithDefaults instantiates a new CdnCreateScopeRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CdnCreateScopeRuleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CdnCreateScopeRuleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CdnCreateScopeRuleRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CdnCreateScopeRuleRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *CdnCreateScopeRuleRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *CdnCreateScopeRuleRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *CdnCreateScopeRuleRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *CdnCreateScopeRuleRequest) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetConfiguration

`func (o *CdnCreateScopeRuleRequest) GetConfiguration() CustconfConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *CdnCreateScopeRuleRequest) GetConfigurationOk() (*CustconfConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *CdnCreateScopeRuleRequest) SetConfiguration(v CustconfConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *CdnCreateScopeRuleRequest) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


