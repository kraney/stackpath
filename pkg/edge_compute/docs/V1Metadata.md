# V1Metadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotations** | **map[string]string** | A string to string key/value pair | [optional] 
**Labels** | **map[string]string** | A string to string key/value pair | [optional] 
**CreatedAt** | Pointer to [**time.Time**](time.Time.md) | The date that a metadata entry was created | [optional] [readonly] 
**UpdatedAt** | Pointer to [**time.Time**](time.Time.md) | The date that a metadata entry was last updated | [optional] [readonly] 
**DeleteRequestedAt** | Pointer to [**time.Time**](time.Time.md) | The date an entity was requested to be deleted | [optional] [readonly] 
**Version** | **string** | A metadata entry&#39;s version number  Metadata versions start at 1 when they are created and increment by 1 every time they are updated. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


