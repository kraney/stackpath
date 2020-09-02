# V1Probe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HttpGet** | [**V1HttpGetAction**](v1HTTPGetAction.md) |  | [optional] 
**TcpSocket** | [**V1TcpSocketAction**](v1TCPSocketAction.md) |  | [optional] 
**InitialDelaySeconds** | **int32** | The number of seconds after a workload instance has started before liveness probes are initiated | [optional] 
**TimeoutSeconds** | **int32** | The number of seconds after which a probe times out  Defaults to 1 second. Minimum value is 1. | [optional] 
**PeriodSeconds** | **int32** | How often, in seconds, to perform a probe  Default to 10 seconds. Minimum value is 1. | [optional] 
**SuccessThreshold** | **int32** | The minimum consecutive successes for a probe to be considered successful after having failed  Defaults to 1. Must be 1 for liveness. Minimum value is 1. | [optional] 
**FailureThreshold** | **int32** | The minimum consecutive failures for a probe to be considered failed after having succeeded  Defaults to 3. Minimum value is 1. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


