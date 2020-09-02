# V2CreateMonitorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The monitor&#39;s name | [optional] 
**Locations** | **[]string** | The IDs of locations to monitor from | [optional] 
**Interval** | **string** | How often a monitor should be checked. | [optional] 
**Timeout** | **string** | How long to wait before timing out a monitor check request to a service. | [optional] 
**IpVersion** | [**V2IpVersion**](v2IpVersion.md) |  | [optional] 
**Http** | [**V2HttpConfiguration**](v2HttpConfiguration.md) |  | [optional] 
**Tcp** | [**V2TcpConfiguration**](v2TcpConfiguration.md) |  | [optional] 
**Enabled** | **bool** | Whether a monitor is enabled or not. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


