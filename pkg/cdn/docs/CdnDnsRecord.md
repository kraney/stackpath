# CdnDnsRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the network node to which a zone resource record pertains  Use the value \&quot;@\&quot; to denote current root domain name. | [optional] 
**Type** | **string** | A zone record&#39;s type  Zone record types describe the zone record&#39;s behavior. For instance, a zone record&#39;s type can say that the record is a name to IP address value, a name alias, or which mail exchanger is responsible for the domain. See https://support.stackpath.com/hc/en-us/articles/360001085563-What-DNS-record-types-does-StackPath-support for more information. | [optional] 
**Class** | **string** | A zone record&#39;s class code  This is typically \&quot;IN\&quot; for Internet related resource records. | [optional] 
**Ttl** | **int32** | A zone record&#39;s time to live  A record&#39;s TTL is the number of seconds that the record should be cached by DNS resolvers. Use lower TTL values if you expect zone records to change often. Use higher TTL values for records that won&#39;t change to prevent extra DNS lookups by clients. | [optional] 
**Data** | **string** | A zone record&#39;s value  Expected data formats can vary depending on the zone record&#39;s type. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


