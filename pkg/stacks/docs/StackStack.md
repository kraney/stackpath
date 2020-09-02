# StackStack

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A stack&#39;s unique identifier | [optional] 
**AccountId** | **string** | The ID of the account that a stack belongs to | [optional] 
**Slug** | **string** | A stack&#39;s human-friendly identifier  This can be used in place of a stack&#39;s ID in most situations to identify a stack. It can be no longer than 30 characters, each being a lowercase letter, number, or dash. It also cannot start with a dash, end with a dash, or have consecutive dashes. This value can be provided on creation or is derived from the stack&#39;s name if a slug isn&#39;t provided. This value cannot be updated once it&#39;s created. | [optional] 
**Name** | **string** | A stack&#39;s name | [optional] 
**CreatedAt** | [**time.Time**](time.Time.md) | The date a stack was created | [optional] [readonly] 
**UpdatedAt** | [**time.Time**](time.Time.md) | The date a stack was last updated | [optional] [readonly] 
**Status** | [**StackStackStatus**](stackStackStatus.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


