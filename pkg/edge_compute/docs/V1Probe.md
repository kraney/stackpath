# V1Probe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HttpGet** | Pointer to [**V1HTTPGetAction**](v1HTTPGetAction.md) |  | [optional] 
**TcpSocket** | Pointer to [**V1TCPSocketAction**](v1TCPSocketAction.md) |  | [optional] 
**InitialDelaySeconds** | Pointer to **int32** | The number of seconds after a workload instance has started before liveness probes are initiated | [optional] 
**TimeoutSeconds** | Pointer to **int32** | The number of seconds after which a probe times out  Defaults to 1 second. Minimum value is 1. | [optional] 
**PeriodSeconds** | Pointer to **int32** | How often, in seconds, to perform a probe  Default to 10 seconds. Minimum value is 1. | [optional] 
**SuccessThreshold** | Pointer to **int32** | The minimum consecutive successes for a probe to be considered successful after having failed  Defaults to 1. Must be 1 for liveness. Minimum value is 1. | [optional] 
**FailureThreshold** | Pointer to **int32** | The minimum consecutive failures for a probe to be considered failed after having succeeded  Defaults to 3. Minimum value is 1. | [optional] 

## Methods

### NewV1Probe

`func NewV1Probe() *V1Probe`

NewV1Probe instantiates a new V1Probe object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ProbeWithDefaults

`func NewV1ProbeWithDefaults() *V1Probe`

NewV1ProbeWithDefaults instantiates a new V1Probe object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttpGet

`func (o *V1Probe) GetHttpGet() V1HTTPGetAction`

GetHttpGet returns the HttpGet field if non-nil, zero value otherwise.

### GetHttpGetOk

`func (o *V1Probe) GetHttpGetOk() (*V1HTTPGetAction, bool)`

GetHttpGetOk returns a tuple with the HttpGet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpGet

`func (o *V1Probe) SetHttpGet(v V1HTTPGetAction)`

SetHttpGet sets HttpGet field to given value.

### HasHttpGet

`func (o *V1Probe) HasHttpGet() bool`

HasHttpGet returns a boolean if a field has been set.

### GetTcpSocket

`func (o *V1Probe) GetTcpSocket() V1TCPSocketAction`

GetTcpSocket returns the TcpSocket field if non-nil, zero value otherwise.

### GetTcpSocketOk

`func (o *V1Probe) GetTcpSocketOk() (*V1TCPSocketAction, bool)`

GetTcpSocketOk returns a tuple with the TcpSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpSocket

`func (o *V1Probe) SetTcpSocket(v V1TCPSocketAction)`

SetTcpSocket sets TcpSocket field to given value.

### HasTcpSocket

`func (o *V1Probe) HasTcpSocket() bool`

HasTcpSocket returns a boolean if a field has been set.

### GetInitialDelaySeconds

`func (o *V1Probe) GetInitialDelaySeconds() int32`

GetInitialDelaySeconds returns the InitialDelaySeconds field if non-nil, zero value otherwise.

### GetInitialDelaySecondsOk

`func (o *V1Probe) GetInitialDelaySecondsOk() (*int32, bool)`

GetInitialDelaySecondsOk returns a tuple with the InitialDelaySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialDelaySeconds

`func (o *V1Probe) SetInitialDelaySeconds(v int32)`

SetInitialDelaySeconds sets InitialDelaySeconds field to given value.

### HasInitialDelaySeconds

`func (o *V1Probe) HasInitialDelaySeconds() bool`

HasInitialDelaySeconds returns a boolean if a field has been set.

### GetTimeoutSeconds

`func (o *V1Probe) GetTimeoutSeconds() int32`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *V1Probe) GetTimeoutSecondsOk() (*int32, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *V1Probe) SetTimeoutSeconds(v int32)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.

### HasTimeoutSeconds

`func (o *V1Probe) HasTimeoutSeconds() bool`

HasTimeoutSeconds returns a boolean if a field has been set.

### GetPeriodSeconds

`func (o *V1Probe) GetPeriodSeconds() int32`

GetPeriodSeconds returns the PeriodSeconds field if non-nil, zero value otherwise.

### GetPeriodSecondsOk

`func (o *V1Probe) GetPeriodSecondsOk() (*int32, bool)`

GetPeriodSecondsOk returns a tuple with the PeriodSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodSeconds

`func (o *V1Probe) SetPeriodSeconds(v int32)`

SetPeriodSeconds sets PeriodSeconds field to given value.

### HasPeriodSeconds

`func (o *V1Probe) HasPeriodSeconds() bool`

HasPeriodSeconds returns a boolean if a field has been set.

### GetSuccessThreshold

`func (o *V1Probe) GetSuccessThreshold() int32`

GetSuccessThreshold returns the SuccessThreshold field if non-nil, zero value otherwise.

### GetSuccessThresholdOk

`func (o *V1Probe) GetSuccessThresholdOk() (*int32, bool)`

GetSuccessThresholdOk returns a tuple with the SuccessThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessThreshold

`func (o *V1Probe) SetSuccessThreshold(v int32)`

SetSuccessThreshold sets SuccessThreshold field to given value.

### HasSuccessThreshold

`func (o *V1Probe) HasSuccessThreshold() bool`

HasSuccessThreshold returns a boolean if a field has been set.

### GetFailureThreshold

`func (o *V1Probe) GetFailureThreshold() int32`

GetFailureThreshold returns the FailureThreshold field if non-nil, zero value otherwise.

### GetFailureThresholdOk

`func (o *V1Probe) GetFailureThresholdOk() (*int32, bool)`

GetFailureThresholdOk returns a tuple with the FailureThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureThreshold

`func (o *V1Probe) SetFailureThreshold(v int32)`

SetFailureThreshold sets FailureThreshold field to given value.

### HasFailureThreshold

`func (o *V1Probe) HasFailureThreshold() bool`

HasFailureThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


