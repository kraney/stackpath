# CdnSite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A CDN site&#39;s unique identifier | [optional] 
**StackId** | **string** | The ID of the stack to which a CDN site belongs | [optional] 
**Label** | **string** | A CDN site&#39;s name  Site names correspond to their fully-qualified domain name. | [optional] 
**Status** | **string** | A CDN site&#39;s internal state  Site status is controlled by StackPath as sites are provisioned and managed by StackPath&#39;s accounting and security teams. | [optional] 
**CreatedAt** | [**time.Time**](time.Time.md) | The date that a CDN site was created | [optional] 
**UpdatedAt** | [**time.Time**](time.Time.md) | The date that a CDN site was last updated | [optional] 
**Features** | [**[]CdnSiteFeature**](cdnSiteFeature.md) | A CDN site&#39;s associated features  Features control how StackPath provisions and configures a site. | [optional] 
**Enabled** | **bool** | Whether or not a site&#39;s CDN service is enabled | [optional] 
**Type** | [**SiteTypeValue**](SiteTypeValue.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


