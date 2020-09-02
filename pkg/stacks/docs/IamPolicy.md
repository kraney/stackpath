# IamPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bindings** | [**[]PolicyBinding**](PolicyBinding.md) | Bindings to assign members to roles | [optional] 
**Version** | **int64** | The current update number used to ensure updates happen to the expected version.  On first update this must be 0 and on each successive update this must be the last known version. When getting or as the result of a set, this will be the current version. | [optional] 
**CreatedAt** | [**time.Time**](time.Time.md) | When this policy was created.  Always present on get, ignored on set. | [optional] [readonly] 
**UpdatedAt** | [**time.Time**](time.Time.md) | When this policy was last updated.  Could be absent on get if not updated since initial creation. Ignored on set. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


