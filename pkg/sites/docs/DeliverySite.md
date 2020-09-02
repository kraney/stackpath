# DeliverySite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A site&#39;s unique identifier | [optional] 
**StackId** | **string** | The ID of the stack to which a site belongs | [optional] 
**Label** | **string** | The site&#39;s name | [optional] 
**Status** | **string** | The site&#39;s status | [optional] 
**Features** | [**[]DeliverySiteFeature**](deliverySiteFeature.md) | A site&#39;s features  Site features control how its content is delivered to end users. | [optional] 
**CreatedAt** | [**time.Time**](time.Time.md) | The date a site was created | [optional] 
**UpdatedAt** | [**time.Time**](time.Time.md) | The date a site was last updated | [optional] 
**ApiUrls** | **[]string** | A list of API URLs which receive different processing through the WAF than website requests | [optional] 
**Monitoring** | **bool** | Whether or not a site&#39;s WAF service is in a monitor state  Sites in WAF monitoring mode accept incoming requests and log but take no action on policy and rule violations. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


