# V1ImageDeprecation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeprecationDate** | [**time.Time**](time.Time.md) | The date when an image is considered deprecated  Deprecated images may be used in new virtual machine based workloads, but are no longer considered for the \&quot;default\&quot; tag, nor are they returned in image lists by default. | [optional] 
**ObsoletionDate** | [**time.Time**](time.Time.md) | The date when an image is considered obsolete  Obsolete images may not be used in new virtual machine based workloads. If an obsoletion date is present then the deprecation date must also be present with a same or earlier date. | [optional] 
**DeletionDate** | [**time.Time**](time.Time.md) | The date when an image is deleted  Deletion dates cannot be before an image&#39;s deprecation or obsoletion dates. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


