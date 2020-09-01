# CdnCreateScopeRuleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rule** | Pointer to [**CdnScopeRule**](cdnScopeRule.md) |  | [optional] 
**Configuration** | Pointer to [**CustconfConfiguration**](custconfConfiguration.md) |  | [optional] 

## Methods

### NewCdnCreateScopeRuleResponse

`func NewCdnCreateScopeRuleResponse() *CdnCreateScopeRuleResponse`

NewCdnCreateScopeRuleResponse instantiates a new CdnCreateScopeRuleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdnCreateScopeRuleResponseWithDefaults

`func NewCdnCreateScopeRuleResponseWithDefaults() *CdnCreateScopeRuleResponse`

NewCdnCreateScopeRuleResponseWithDefaults instantiates a new CdnCreateScopeRuleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRule

`func (o *CdnCreateScopeRuleResponse) GetRule() CdnScopeRule`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *CdnCreateScopeRuleResponse) GetRuleOk() (*CdnScopeRule, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *CdnCreateScopeRuleResponse) SetRule(v CdnScopeRule)`

SetRule sets Rule field to given value.

### HasRule

`func (o *CdnCreateScopeRuleResponse) HasRule() bool`

HasRule returns a boolean if a field has been set.

### GetConfiguration

`func (o *CdnCreateScopeRuleResponse) GetConfiguration() CustconfConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *CdnCreateScopeRuleResponse) GetConfigurationOk() (*CustconfConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *CdnCreateScopeRuleResponse) SetConfiguration(v CustconfConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *CdnCreateScopeRuleResponse) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


