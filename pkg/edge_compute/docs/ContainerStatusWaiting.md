# ContainerStatusWaiting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** | The reason that a container is waiting to start | [optional] 
**Message** | Pointer to **string** | A more detailed explanation of a container&#39;s waiting state | [optional] 

## Methods

### NewContainerStatusWaiting

`func NewContainerStatusWaiting() *ContainerStatusWaiting`

NewContainerStatusWaiting instantiates a new ContainerStatusWaiting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerStatusWaitingWithDefaults

`func NewContainerStatusWaitingWithDefaults() *ContainerStatusWaiting`

NewContainerStatusWaitingWithDefaults instantiates a new ContainerStatusWaiting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *ContainerStatusWaiting) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ContainerStatusWaiting) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ContainerStatusWaiting) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ContainerStatusWaiting) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetMessage

`func (o *ContainerStatusWaiting) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ContainerStatusWaiting) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ContainerStatusWaiting) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ContainerStatusWaiting) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


