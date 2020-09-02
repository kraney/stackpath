# V1ContainerSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | **string** | The location of a Docker image to run as a container | [optional] 
**Command** | **[]string** | The commands that start a container | [optional] 
**Env** | [**map[string]V1EnvironmentVariable**](v1EnvironmentVariable.md) | A string to environment variable key/value pair | [optional] 
**Ports** | [**map[string]V1InstancePort**](v1InstancePort.md) | A string to network port key/value pair | [optional] 
**LivenessProbe** | [**V1Probe**](v1Probe.md) |  | [optional] 
**ReadinessProbe** | [**V1Probe**](v1Probe.md) |  | [optional] 
**Resources** | [**V1ResourceRequirements**](v1ResourceRequirements.md) |  | [optional] 
**VolumeMounts** | [**[]V1InstanceVolumeMount**](v1InstanceVolumeMount.md) | Volumes to mount in the container | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


