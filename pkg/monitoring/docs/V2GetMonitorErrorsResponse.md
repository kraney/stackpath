# V2GetMonitorErrorsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]V2MonitorError**](v2MonitorError.md) | The requested monitor errors. | [optional] 

## Methods

### NewV2GetMonitorErrorsResponse

`func NewV2GetMonitorErrorsResponse() *V2GetMonitorErrorsResponse`

NewV2GetMonitorErrorsResponse instantiates a new V2GetMonitorErrorsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2GetMonitorErrorsResponseWithDefaults

`func NewV2GetMonitorErrorsResponseWithDefaults() *V2GetMonitorErrorsResponse`

NewV2GetMonitorErrorsResponseWithDefaults instantiates a new V2GetMonitorErrorsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *V2GetMonitorErrorsResponse) GetResults() []V2MonitorError`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *V2GetMonitorErrorsResponse) GetResultsOk() (*[]V2MonitorError, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *V2GetMonitorErrorsResponse) SetResults(v []V2MonitorError)`

SetResults sets Results field to given value.

### HasResults

`func (o *V2GetMonitorErrorsResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


