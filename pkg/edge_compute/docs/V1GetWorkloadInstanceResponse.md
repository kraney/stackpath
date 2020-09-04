# V1GetWorkloadInstanceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instance** | Pointer to [**Workloadv1Instance**](workloadv1Instance.md) |  | [optional] 

## Methods

### NewV1GetWorkloadInstanceResponse

`func NewV1GetWorkloadInstanceResponse() *V1GetWorkloadInstanceResponse`

NewV1GetWorkloadInstanceResponse instantiates a new V1GetWorkloadInstanceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1GetWorkloadInstanceResponseWithDefaults

`func NewV1GetWorkloadInstanceResponseWithDefaults() *V1GetWorkloadInstanceResponse`

NewV1GetWorkloadInstanceResponseWithDefaults instantiates a new V1GetWorkloadInstanceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstance

`func (o *V1GetWorkloadInstanceResponse) GetInstance() Workloadv1Instance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *V1GetWorkloadInstanceResponse) GetInstanceOk() (*Workloadv1Instance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *V1GetWorkloadInstanceResponse) SetInstance(v Workloadv1Instance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *V1GetWorkloadInstanceResponse) HasInstance() bool`

HasInstance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


