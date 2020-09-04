# V2GetLocationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]Monitoringv2Location**](monitoringv2Location.md) | The list of monitoring locations for a stack. | [optional] 

## Methods

### NewV2GetLocationsResponse

`func NewV2GetLocationsResponse() *V2GetLocationsResponse`

NewV2GetLocationsResponse instantiates a new V2GetLocationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2GetLocationsResponseWithDefaults

`func NewV2GetLocationsResponseWithDefaults() *V2GetLocationsResponse`

NewV2GetLocationsResponseWithDefaults instantiates a new V2GetLocationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *V2GetLocationsResponse) GetResults() []Monitoringv2Location`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *V2GetLocationsResponse) GetResultsOk() (*[]Monitoringv2Location, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *V2GetLocationsResponse) SetResults(v []Monitoringv2Location)`

SetResults sets Results field to given value.

### HasResults

`func (o *V2GetLocationsResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


