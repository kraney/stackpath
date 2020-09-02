# ZoneZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StackId** | **string** | The ID of the stack that a zone belongs to | [optional] 
**AccountId** | **string** | The ID of the StackPath account that owns a zone | [optional] 
**Id** | **string** | A zone&#39;s unique ID | [optional] 
**Domain** | **string** | A zone&#39;s name | [optional] 
**Version** | **string** | A zone&#39;s version number  Version numbers are incremented automatically when a zone is updated | [optional] 
**Labels** | **map[string]string** | A key/value pair of user-defined labels for a zone  Zone labels are not processed by StackPath and are solely used for users to organize their DNS zones. | [optional] 
**Created** | [**time.Time**](time.Time.md) | The date a zone was created | [optional] 
**Updated** | [**time.Time**](time.Time.md) | The date a zone was last updated | [optional] 
**Nameservers** | **[]string** | The hostnames of the StackPath resolvers that host a zone  Every zone has multiple name servers assigned by StackPath upon creation for redundancy purposes. | [optional] 
**Verified** | [**time.Time**](time.Time.md) | The date a zone&#39;s nameservers were last audited by StackPath | [optional] 
**Status** | [**ZoneZoneStatus**](zoneZoneStatus.md) |  | [optional] 
**Disabled** | **bool** | Whether or not a zone has been disabled by the user | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


