# MetricsOrganizations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Organization** | Pointer to **string** | An organization&#39;s name | [optional] 
**Count** | Pointer to **string** | The number of requests from the organization | [optional] 

## Methods

### NewMetricsOrganizations

`func NewMetricsOrganizations() *MetricsOrganizations`

NewMetricsOrganizations instantiates a new MetricsOrganizations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsOrganizationsWithDefaults

`func NewMetricsOrganizationsWithDefaults() *MetricsOrganizations`

NewMetricsOrganizationsWithDefaults instantiates a new MetricsOrganizations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganization

`func (o *MetricsOrganizations) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *MetricsOrganizations) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *MetricsOrganizations) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *MetricsOrganizations) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetCount

`func (o *MetricsOrganizations) GetCount() string`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *MetricsOrganizations) GetCountOk() (*string, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *MetricsOrganizations) SetCount(v string)`

SetCount sets Count field to given value.

### HasCount

`func (o *MetricsOrganizations) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


