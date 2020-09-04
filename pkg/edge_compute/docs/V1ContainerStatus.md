# V1ContainerStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A container status&#39; name | [optional] 
**Phase** | Pointer to [**V1ContainerStatusContainerPhase**](v1ContainerStatusContainerPhase.md) |  | [optional] [default to "CONTAINER_PHASE_UNSPECIFIED"]
**StartedAt** | Pointer to [**time.Time**](time.Time.md) | The date a container started | [optional] 
**FinishedAt** | Pointer to [**time.Time**](time.Time.md) | The date a container terminated | [optional] 
**Waiting** | Pointer to [**ContainerStatusWaiting**](ContainerStatusWaiting.md) |  | [optional] 
**Running** | Pointer to [**ContainerStatusRunning**](ContainerStatusRunning.md) |  | [optional] 
**Terminated** | Pointer to [**ContainerStatusTerminated**](ContainerStatusTerminated.md) |  | [optional] 
**Ready** | Pointer to **bool** | Whether or not a container is running and ready for service | [optional] 
**RestartCount** | Pointer to **int32** | How many times a container has restarted since it was first started | [optional] 
**ContainerId** | Pointer to **string** | A unique value that identifies a container | [optional] 

## Methods

### NewV1ContainerStatus

`func NewV1ContainerStatus() *V1ContainerStatus`

NewV1ContainerStatus instantiates a new V1ContainerStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ContainerStatusWithDefaults

`func NewV1ContainerStatusWithDefaults() *V1ContainerStatus`

NewV1ContainerStatusWithDefaults instantiates a new V1ContainerStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *V1ContainerStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1ContainerStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1ContainerStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1ContainerStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPhase

`func (o *V1ContainerStatus) GetPhase() V1ContainerStatusContainerPhase`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *V1ContainerStatus) GetPhaseOk() (*V1ContainerStatusContainerPhase, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *V1ContainerStatus) SetPhase(v V1ContainerStatusContainerPhase)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *V1ContainerStatus) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetStartedAt

`func (o *V1ContainerStatus) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *V1ContainerStatus) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *V1ContainerStatus) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *V1ContainerStatus) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetFinishedAt

`func (o *V1ContainerStatus) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *V1ContainerStatus) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *V1ContainerStatus) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *V1ContainerStatus) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetWaiting

`func (o *V1ContainerStatus) GetWaiting() ContainerStatusWaiting`

GetWaiting returns the Waiting field if non-nil, zero value otherwise.

### GetWaitingOk

`func (o *V1ContainerStatus) GetWaitingOk() (*ContainerStatusWaiting, bool)`

GetWaitingOk returns a tuple with the Waiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaiting

`func (o *V1ContainerStatus) SetWaiting(v ContainerStatusWaiting)`

SetWaiting sets Waiting field to given value.

### HasWaiting

`func (o *V1ContainerStatus) HasWaiting() bool`

HasWaiting returns a boolean if a field has been set.

### GetRunning

`func (o *V1ContainerStatus) GetRunning() ContainerStatusRunning`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *V1ContainerStatus) GetRunningOk() (*ContainerStatusRunning, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *V1ContainerStatus) SetRunning(v ContainerStatusRunning)`

SetRunning sets Running field to given value.

### HasRunning

`func (o *V1ContainerStatus) HasRunning() bool`

HasRunning returns a boolean if a field has been set.

### GetTerminated

`func (o *V1ContainerStatus) GetTerminated() ContainerStatusTerminated`

GetTerminated returns the Terminated field if non-nil, zero value otherwise.

### GetTerminatedOk

`func (o *V1ContainerStatus) GetTerminatedOk() (*ContainerStatusTerminated, bool)`

GetTerminatedOk returns a tuple with the Terminated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminated

`func (o *V1ContainerStatus) SetTerminated(v ContainerStatusTerminated)`

SetTerminated sets Terminated field to given value.

### HasTerminated

`func (o *V1ContainerStatus) HasTerminated() bool`

HasTerminated returns a boolean if a field has been set.

### GetReady

`func (o *V1ContainerStatus) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *V1ContainerStatus) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *V1ContainerStatus) SetReady(v bool)`

SetReady sets Ready field to given value.

### HasReady

`func (o *V1ContainerStatus) HasReady() bool`

HasReady returns a boolean if a field has been set.

### GetRestartCount

`func (o *V1ContainerStatus) GetRestartCount() int32`

GetRestartCount returns the RestartCount field if non-nil, zero value otherwise.

### GetRestartCountOk

`func (o *V1ContainerStatus) GetRestartCountOk() (*int32, bool)`

GetRestartCountOk returns a tuple with the RestartCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartCount

`func (o *V1ContainerStatus) SetRestartCount(v int32)`

SetRestartCount sets RestartCount field to given value.

### HasRestartCount

`func (o *V1ContainerStatus) HasRestartCount() bool`

HasRestartCount returns a boolean if a field has been set.

### GetContainerId

`func (o *V1ContainerStatus) GetContainerId() string`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *V1ContainerStatus) GetContainerIdOk() (*string, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *V1ContainerStatus) SetContainerId(v string)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *V1ContainerStatus) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


