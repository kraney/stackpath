# StatisticsMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to [**[]MetricsRules**](MetricsRules.md) | The \&quot;top threats\&quot; metric, the number of events encountered per rule | [optional] 
**Countries** | Pointer to [**[]MetricsCountries**](MetricsCountries.md) | The number of events per country of origin | [optional] 
**Organizations** | Pointer to [**[]MetricsOrganizations**](MetricsOrganizations.md) | The number of events per requesting organization, as determined by WHOIS lookup on the requesting IP | [optional] 
**Actions** | Pointer to [**[]MetricsActions**](MetricsActions.md) | The number of events per action taken by the WAF | [optional] 

## Methods

### NewStatisticsMetrics

`func NewStatisticsMetrics() *StatisticsMetrics`

NewStatisticsMetrics instantiates a new StatisticsMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatisticsMetricsWithDefaults

`func NewStatisticsMetricsWithDefaults() *StatisticsMetrics`

NewStatisticsMetricsWithDefaults instantiates a new StatisticsMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *StatisticsMetrics) GetRules() []MetricsRules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *StatisticsMetrics) GetRulesOk() (*[]MetricsRules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *StatisticsMetrics) SetRules(v []MetricsRules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *StatisticsMetrics) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetCountries

`func (o *StatisticsMetrics) GetCountries() []MetricsCountries`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *StatisticsMetrics) GetCountriesOk() (*[]MetricsCountries, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *StatisticsMetrics) SetCountries(v []MetricsCountries)`

SetCountries sets Countries field to given value.

### HasCountries

`func (o *StatisticsMetrics) HasCountries() bool`

HasCountries returns a boolean if a field has been set.

### GetOrganizations

`func (o *StatisticsMetrics) GetOrganizations() []MetricsOrganizations`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *StatisticsMetrics) GetOrganizationsOk() (*[]MetricsOrganizations, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *StatisticsMetrics) SetOrganizations(v []MetricsOrganizations)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *StatisticsMetrics) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.

### GetActions

`func (o *StatisticsMetrics) GetActions() []MetricsActions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *StatisticsMetrics) GetActionsOk() (*[]MetricsActions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *StatisticsMetrics) SetActions(v []MetricsActions)`

SetActions sets Actions field to given value.

### HasActions

`func (o *StatisticsMetrics) HasActions() bool`

HasActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


