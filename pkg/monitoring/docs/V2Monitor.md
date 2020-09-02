# V2Monitor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A monitor&#39;s unique identifier | [optional] 
**Name** | **string** | The name of a monitor  A monitor&#39;s name | [optional] 
**CurrentStatus** | [**V2MonitorStatus**](v2MonitorStatus.md) |  | [optional] 
**UptimePercentage** | **float64** | The percent of time a monitored service has had a status of UP or SLOW in the last hour | [optional] 
**AvgResponseTime** | **string** | The average response time for a monitor | [optional] 
**Locations** | [**[]Monitoringv2Location**](monitoringv2Location.md) | The locations a service is monitored from. | [optional] 
**Interval** | **string** | The amount of time between each monitor check.  By default each service monitor will be checked every five minutes. | [optional] 
**Timeout** | **string** | Then length of time to wait until the monitor check times out. | [optional] 
**IpVersion** | [**V2IpVersion**](v2IpVersion.md) |  | [optional] 
**Http** | [**V2HttpConfiguration**](v2HttpConfiguration.md) |  | [optional] 
**Tcp** | [**V2TcpConfiguration**](v2TcpConfiguration.md) |  | [optional] 
**CreatedAt** | [**time.Time**](time.Time.md) | The date a monitor was created. | [optional] 
**UpdatedAt** | [**time.Time**](time.Time.md) | The date a monitor was last updated. | [optional] 
**Enabled** | **bool** | Whether or not a monitor is enabled. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


