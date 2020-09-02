# CdnPopScanReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PopCode** | **string** | The IATA formatted location code of the StackPath POP that produced a scan report | [optional] 
**ConnectMs** | **float32** | The amount of time in milliseconds that a POP scan took to perform an initial connection handshake | [optional] 
**DnsMs** | **float32** | The amount of time in milliseconds that a POP scan took to resolve the target&#39;s DNS entry | [optional] 
**DownloadMs** | **float32** | The amount of time in milliseconds that a POP scan took to download the target&#39;s contents | [optional] 
**FirstByteMs** | **float32** | The amount of time in milliseconds that a POP scan took from initial connection to starting to receive the target&#39;s contents | [optional] 
**SslMs** | **float32** | The amount of time in milliseconds that a POP scan took to perform an SSL handshake | [optional] 
**TotalMs** | **float32** | The total amount of time in milliseconds that a POP scan took to complete | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


