# ZoneCreateZoneMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | The zone&#39;s domain name | [optional] 
**Labels** | **map[string]string** | A key/value pair of user-defined labels for a DNS zone  Zone labels are not processed by StackPath and are solely used for users to organize their DNS zones. | [optional] 
**UseApexDomain** | **bool** | Whether or not to create a zone for the apex domain only  If this is true and a domain with subdomains is provided, it will be stripped and only the root domain will be used for the zone. If this is false an error will be returned if it&#39;s not already an apex domain. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


