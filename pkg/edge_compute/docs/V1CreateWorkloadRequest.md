# V1CreateWorkloadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Workload** | Pointer to [**V1Workload**](v1Workload.md) |  | [optional] 

## Methods

### NewV1CreateWorkloadRequest

`func NewV1CreateWorkloadRequest() *V1CreateWorkloadRequest`

NewV1CreateWorkloadRequest instantiates a new V1CreateWorkloadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1CreateWorkloadRequestWithDefaults

`func NewV1CreateWorkloadRequestWithDefaults() *V1CreateWorkloadRequest`

NewV1CreateWorkloadRequestWithDefaults instantiates a new V1CreateWorkloadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkload

`func (o *V1CreateWorkloadRequest) GetWorkload() V1Workload`

GetWorkload returns the Workload field if non-nil, zero value otherwise.

### GetWorkloadOk

`func (o *V1CreateWorkloadRequest) GetWorkloadOk() (*V1Workload, bool)`

GetWorkloadOk returns a tuple with the Workload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkload

`func (o *V1CreateWorkloadRequest) SetWorkload(v V1Workload)`

SetWorkload sets Workload field to given value.

### HasWorkload

`func (o *V1CreateWorkloadRequest) HasWorkload() bool`

HasWorkload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


