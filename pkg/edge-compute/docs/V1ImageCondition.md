# V1ImageCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Type of the condition | [optional] 
**Status** | Pointer to [**V1ImageConditionStatus**](v1ImageConditionStatus.md) |  | [optional] [default to "IMAGE_CONDITION_STATUS_UNKNOWN"]
**CheckedAt** | Pointer to [**time.Time**](time.Time.md) | The last time the condition was checked | [optional] 
**TransitionedAt** | Pointer to [**time.Time**](time.Time.md) | The last time the condition transitioned status | [optional] 
**Reason** | Pointer to **string** | A stable identifier for the reason the condition is in its current state | [optional] 
**Message** | Pointer to **string** | A human readable message with details regarding the condition | [optional] 

## Methods

### NewV1ImageCondition

`func NewV1ImageCondition() *V1ImageCondition`

NewV1ImageCondition instantiates a new V1ImageCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ImageConditionWithDefaults

`func NewV1ImageConditionWithDefaults() *V1ImageCondition`

NewV1ImageConditionWithDefaults instantiates a new V1ImageCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *V1ImageCondition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1ImageCondition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1ImageCondition) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *V1ImageCondition) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *V1ImageCondition) GetStatus() V1ImageConditionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1ImageCondition) GetStatusOk() (*V1ImageConditionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1ImageCondition) SetStatus(v V1ImageConditionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V1ImageCondition) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCheckedAt

`func (o *V1ImageCondition) GetCheckedAt() time.Time`

GetCheckedAt returns the CheckedAt field if non-nil, zero value otherwise.

### GetCheckedAtOk

`func (o *V1ImageCondition) GetCheckedAtOk() (*time.Time, bool)`

GetCheckedAtOk returns a tuple with the CheckedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckedAt

`func (o *V1ImageCondition) SetCheckedAt(v time.Time)`

SetCheckedAt sets CheckedAt field to given value.

### HasCheckedAt

`func (o *V1ImageCondition) HasCheckedAt() bool`

HasCheckedAt returns a boolean if a field has been set.

### GetTransitionedAt

`func (o *V1ImageCondition) GetTransitionedAt() time.Time`

GetTransitionedAt returns the TransitionedAt field if non-nil, zero value otherwise.

### GetTransitionedAtOk

`func (o *V1ImageCondition) GetTransitionedAtOk() (*time.Time, bool)`

GetTransitionedAtOk returns a tuple with the TransitionedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitionedAt

`func (o *V1ImageCondition) SetTransitionedAt(v time.Time)`

SetTransitionedAt sets TransitionedAt field to given value.

### HasTransitionedAt

`func (o *V1ImageCondition) HasTransitionedAt() bool`

HasTransitionedAt returns a boolean if a field has been set.

### GetReason

`func (o *V1ImageCondition) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *V1ImageCondition) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *V1ImageCondition) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *V1ImageCondition) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetMessage

`func (o *V1ImageCondition) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V1ImageCondition) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V1ImageCondition) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *V1ImageCondition) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


