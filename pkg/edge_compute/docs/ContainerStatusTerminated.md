# ContainerStatusTerminated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExitCode** | Pointer to **int32** | The exit code of the command that started the container | [optional] 
**Reason** | Pointer to **string** | The reason a container terminated | [optional] 
**Message** | Pointer to **string** | A more detailed explanation of a container&#39;s termination | [optional] 
**StartedAt** | Pointer to [**time.Time**](time.Time.md) | The date a container started | [optional] 
**FinishedAt** | Pointer to [**time.Time**](time.Time.md) | The date a container terminated | [optional] 

## Methods

### NewContainerStatusTerminated

`func NewContainerStatusTerminated() *ContainerStatusTerminated`

NewContainerStatusTerminated instantiates a new ContainerStatusTerminated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerStatusTerminatedWithDefaults

`func NewContainerStatusTerminatedWithDefaults() *ContainerStatusTerminated`

NewContainerStatusTerminatedWithDefaults instantiates a new ContainerStatusTerminated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExitCode

`func (o *ContainerStatusTerminated) GetExitCode() int32`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *ContainerStatusTerminated) GetExitCodeOk() (*int32, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *ContainerStatusTerminated) SetExitCode(v int32)`

SetExitCode sets ExitCode field to given value.

### HasExitCode

`func (o *ContainerStatusTerminated) HasExitCode() bool`

HasExitCode returns a boolean if a field has been set.

### GetReason

`func (o *ContainerStatusTerminated) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ContainerStatusTerminated) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ContainerStatusTerminated) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ContainerStatusTerminated) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetMessage

`func (o *ContainerStatusTerminated) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ContainerStatusTerminated) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ContainerStatusTerminated) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ContainerStatusTerminated) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStartedAt

`func (o *ContainerStatusTerminated) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ContainerStatusTerminated) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ContainerStatusTerminated) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *ContainerStatusTerminated) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetFinishedAt

`func (o *ContainerStatusTerminated) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *ContainerStatusTerminated) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *ContainerStatusTerminated) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *ContainerStatusTerminated) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


