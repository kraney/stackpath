# WafSite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A WAF site&#39;s unique identifier | [optional] 
**Name** | **string** | The WAF site&#39;s name, used on stand-alone WAF sites | [optional] 
**Enabled** | **bool** | Whether or not a site&#39;s WAF service is enabled | [optional] 
**ApiUrls** | **[]string** | A list of API URLs which receive different processing through the WAF than website requests | [optional] 
**CreatedAt** | [**time.Time**](time.Time.md) | The date a WAF site was created | [optional] 
**UpdatedAt** | [**time.Time**](time.Time.md) | The date a WAF site was last updated | [optional] 
**Status** | [**WafSiteStatus**](wafSiteStatus.md) |  | [optional] 
**Mode** | [**SiteAttachMode**](SiteAttachMode.md) |  | [optional] 
**DeliveryId** | **string** | For an ATTACHED site, the CDN site ID where caching can be enabled | [optional] 
**Type** | [**WafSiteType**](wafSiteType.md) |  | [optional] 
**Monitoring** | **bool** | Whether or not a site&#39;s WAF service is in a monitor state  WAF sites in monitoring mode accept incoming requests and log but take no action on policy and rule violations. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


