# V1VirtualMachineStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the virtual machine the status applies to | [optional] 
**Phase** | Pointer to [**VirtualMachineStatusPhase**](VirtualMachineStatusPhase.md) |  | [optional] [default to "UNKNOWN"]
**Reason** | Pointer to **string** | A short reason why the virtual machine is in its current phase | [optional] 
**Message** | Pointer to **string** | A longer message with details about why the virtual machine is in its current phase | [optional] 

## Methods

### NewV1VirtualMachineStatus

`func NewV1VirtualMachineStatus() *V1VirtualMachineStatus`

NewV1VirtualMachineStatus instantiates a new V1VirtualMachineStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1VirtualMachineStatusWithDefaults

`func NewV1VirtualMachineStatusWithDefaults() *V1VirtualMachineStatus`

NewV1VirtualMachineStatusWithDefaults instantiates a new V1VirtualMachineStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *V1VirtualMachineStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1VirtualMachineStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1VirtualMachineStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1VirtualMachineStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPhase

`func (o *V1VirtualMachineStatus) GetPhase() VirtualMachineStatusPhase`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *V1VirtualMachineStatus) GetPhaseOk() (*VirtualMachineStatusPhase, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *V1VirtualMachineStatus) SetPhase(v VirtualMachineStatusPhase)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *V1VirtualMachineStatus) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetReason

`func (o *V1VirtualMachineStatus) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *V1VirtualMachineStatus) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *V1VirtualMachineStatus) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *V1VirtualMachineStatus) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetMessage

`func (o *V1VirtualMachineStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V1VirtualMachineStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V1VirtualMachineStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *V1VirtualMachineStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


