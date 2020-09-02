# V2UpdateAlertConditionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | [**V2AlertConditionMetric**](v2AlertConditionMetric.md) |  | [optional] 
**Comparator** | [**AlertConditionComparator**](AlertConditionComparator.md) |  | [optional] 
**Value** | **string** | The value to check to determine if an alert should be generated. | [optional] 
**AlarmDelay** | **int32** | The amount of time to wait before sending an alert.  This is useful to prevent alerts due to a flapping service. | [optional] 
**AlarmEvery** | **int32** | How often an alert should be generated if an alert condition is still being met. | [optional] 
**AlarmUntilSilenced** | **bool** | A flag that determines if an alert should continue to trigger until it is resolved. | [optional] 
**Enabled** | **bool** | Whether an alert condition is enabled or not. | [optional] 
**MinimumLocations** | **int32** | The minimum number of locations that the alert condition must be triggered for before an alert is generated.  This should always be less than or equal to the number of locations a monitor has. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


