# CustconfTimePseudoStreaming

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | This is used by the API to perform conflict checking | [optional] 
**JumpToTimeStartParam** | Pointer to **string** | The name of the query string parameter that indicates to the CDN the specific time interval of the video that is being requested. | [optional] 
**JumpToTimeEndParam** | Pointer to **string** | The name of the query string parameter that indicates to the CDN the end time of the video that should be returned. | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewCustconfTimePseudoStreaming

`func NewCustconfTimePseudoStreaming() *CustconfTimePseudoStreaming`

NewCustconfTimePseudoStreaming instantiates a new CustconfTimePseudoStreaming object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustconfTimePseudoStreamingWithDefaults

`func NewCustconfTimePseudoStreamingWithDefaults() *CustconfTimePseudoStreaming`

NewCustconfTimePseudoStreamingWithDefaults instantiates a new CustconfTimePseudoStreaming object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustconfTimePseudoStreaming) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustconfTimePseudoStreaming) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustconfTimePseudoStreaming) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustconfTimePseudoStreaming) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJumpToTimeStartParam

`func (o *CustconfTimePseudoStreaming) GetJumpToTimeStartParam() string`

GetJumpToTimeStartParam returns the JumpToTimeStartParam field if non-nil, zero value otherwise.

### GetJumpToTimeStartParamOk

`func (o *CustconfTimePseudoStreaming) GetJumpToTimeStartParamOk() (*string, bool)`

GetJumpToTimeStartParamOk returns a tuple with the JumpToTimeStartParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJumpToTimeStartParam

`func (o *CustconfTimePseudoStreaming) SetJumpToTimeStartParam(v string)`

SetJumpToTimeStartParam sets JumpToTimeStartParam field to given value.

### HasJumpToTimeStartParam

`func (o *CustconfTimePseudoStreaming) HasJumpToTimeStartParam() bool`

HasJumpToTimeStartParam returns a boolean if a field has been set.

### GetJumpToTimeEndParam

`func (o *CustconfTimePseudoStreaming) GetJumpToTimeEndParam() string`

GetJumpToTimeEndParam returns the JumpToTimeEndParam field if non-nil, zero value otherwise.

### GetJumpToTimeEndParamOk

`func (o *CustconfTimePseudoStreaming) GetJumpToTimeEndParamOk() (*string, bool)`

GetJumpToTimeEndParamOk returns a tuple with the JumpToTimeEndParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJumpToTimeEndParam

`func (o *CustconfTimePseudoStreaming) SetJumpToTimeEndParam(v string)`

SetJumpToTimeEndParam sets JumpToTimeEndParam field to given value.

### HasJumpToTimeEndParam

`func (o *CustconfTimePseudoStreaming) HasJumpToTimeEndParam() bool`

HasJumpToTimeEndParam returns a boolean if a field has been set.

### GetEnabled

`func (o *CustconfTimePseudoStreaming) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustconfTimePseudoStreaming) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustconfTimePseudoStreaming) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CustconfTimePseudoStreaming) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


