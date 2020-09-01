# ApiStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** |  | [optional] 
**Details** | Pointer to [**[]ApiStatusDetail**](apiStatusDetail.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewApiStatus

`func NewApiStatus() *ApiStatus`

NewApiStatus instantiates a new ApiStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiStatusWithDefaults

`func NewApiStatusWithDefaults() *ApiStatus`

NewApiStatusWithDefaults instantiates a new ApiStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ApiStatus) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ApiStatus) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ApiStatus) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *ApiStatus) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDetails

`func (o *ApiStatus) GetDetails() []ApiStatusDetail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ApiStatus) GetDetailsOk() (*[]ApiStatusDetail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ApiStatus) SetDetails(v []ApiStatusDetail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ApiStatus) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetMessage

`func (o *ApiStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApiStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


