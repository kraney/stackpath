# V1VirtualMachineSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | **string** | The image to use for the virtual machine  This is in the format of &lt;stack-slug&gt;/&lt;image-family&gt;[:&lt;image-tag&gt;]. If the image tag portion is omitted, &#39;default&#39; is assumed which is the most recently created, ready, and non-deprecated image of that slug. A set of common images is present on the &#39;stackpath-edge&#39; stack. | [optional] 
**UserData** | **string** | Base64 encoded cloud-init compatible user-data | [optional] 
**Ports** | [**map[string]V1InstancePort**](v1InstancePort.md) | A string to network port key/value pair | [optional] 
**LivenessProbe** | [**V1Probe**](v1Probe.md) |  | [optional] 
**ReadinessProbe** | [**V1Probe**](v1Probe.md) |  | [optional] 
**Resources** | [**V1ResourceRequirements**](v1ResourceRequirements.md) |  | [optional] 
**VolumeMounts** | [**[]V1InstanceVolumeMount**](v1InstanceVolumeMount.md) | Volumes to mount in the virtual machine | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


