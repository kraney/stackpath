# ContainerStatusRunning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartedAt** | Pointer to [**time.Time**](time.Time.md) | The date a container started | [optional] 

## Methods

### NewContainerStatusRunning

`func NewContainerStatusRunning() *ContainerStatusRunning`

NewContainerStatusRunning instantiates a new ContainerStatusRunning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerStatusRunningWithDefaults

`func NewContainerStatusRunningWithDefaults() *ContainerStatusRunning`

NewContainerStatusRunningWithDefaults instantiates a new ContainerStatusRunning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartedAt

`func (o *ContainerStatusRunning) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ContainerStatusRunning) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ContainerStatusRunning) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *ContainerStatusRunning) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


