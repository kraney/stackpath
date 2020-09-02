# V1ContainerStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A container status&#39; name | [optional] 
**Phase** | [**V1ContainerStatusContainerPhase**](v1ContainerStatusContainerPhase.md) |  | [optional] 
**StartedAt** | [**time.Time**](time.Time.md) | The date a container started | [optional] 
**FinishedAt** | [**time.Time**](time.Time.md) | The date a container terminated | [optional] 
**Waiting** | [**ContainerStatusWaiting**](ContainerStatusWaiting.md) |  | [optional] 
**Running** | [**ContainerStatusRunning**](ContainerStatusRunning.md) |  | [optional] 
**Terminated** | [**ContainerStatusTerminated**](ContainerStatusTerminated.md) |  | [optional] 
**Ready** | **bool** | Whether or not a container is running and ready for service | [optional] 
**RestartCount** | **int32** | How many times a container has restarted since it was first started | [optional] 
**ContainerId** | **string** | A unique value that identifies a container | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


