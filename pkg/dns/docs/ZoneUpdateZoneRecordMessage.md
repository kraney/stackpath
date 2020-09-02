# ZoneUpdateZoneRecordMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A zone record&#39;s name | [optional] 
**Type** | [**ZoneRecordType**](zoneRecordType.md) |  | [optional] 
**Ttl** | **int32** | A zone record&#39;s time to live  A record&#39;s TTL is the number of seconds that the record should be cached by DNS resolvers. Use lower TTL values if you expect zone records to change often. Use higher TTL values for records that won&#39;t change to prevent extra DNS lookups by clients. | [optional] 
**Data** | **string** | A zone record&#39;s value  Expected data formats can vary depending on the zone record&#39;s type. | [optional] 
**Weight** | **int32** | A zone record&#39;s priority  A resource record is replicated in StackPath&#39;s DNS infrastructure the number of times of the record&#39;s weight, giving it a more likely response to queries if a zone has records with the same name and type. | [optional] 
**Labels** | **map[string]string** | A key/value pair of user-defined labels for a DNS zone record  Zone record labels are not processed by StackPath and are solely used for users to organize their DNS zones. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


