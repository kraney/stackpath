# V1Workload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StackId** | **string** | The ID of the stack that a workload belongs to | [optional] [readonly] 
**Id** | **string** | A workload&#39;s unique identifier | [optional] [readonly] 
**Name** | **string** | A workload&#39;s name as it appears in the StackPath portal | [optional] 
**Slug** | **string** | A workload&#39;s programmatic name  Workload slugs are used to build its instances names | [optional] 
**Metadata** | [**V1Metadata**](v1Metadata.md) |  | [optional] 
**Spec** | [**V1WorkloadSpec**](v1WorkloadSpec.md) |  | [optional] 
**Targets** | [**map[string]V1Target**](v1Target.md) | A string to deployment target key/value pair | [optional] 
**Status** | [**V1WorkloadStatus**](v1WorkloadStatus.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


