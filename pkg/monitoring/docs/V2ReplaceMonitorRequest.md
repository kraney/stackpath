# V2ReplaceMonitorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the monitor. | [optional] 
**Locations** | **[]string** | The list of locations a monitor is checked from. | [optional] 
**Interval** | **string** | How often a monitor is checked. | [optional] 
**Timeout** | **string** | The amount of time to wait before a monitor check times out. | [optional] 
**IpVersion** | [**V2IpVersion**](v2IpVersion.md) |  | [optional] 
**Http** | [**V2HttpConfiguration**](v2HttpConfiguration.md) |  | [optional] 
**Tcp** | [**V2TcpConfiguration**](v2TcpConfiguration.md) |  | [optional] 
**Enabled** | **bool** | Whether a monitor is enabled or not. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


