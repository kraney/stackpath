# \SitesApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSite**](SitesApi.md#CreateSite) | **Post** /delivery/v1/stacks/{stack_id}/sites | Create a site
[**DeleteSite**](SitesApi.md#DeleteSite) | **Delete** /delivery/v1/stacks/{stack_id}/sites/{site_id} | Delete a site
[**GetSite**](SitesApi.md#GetSite) | **Get** /delivery/v1/stacks/{stack_id}/sites/{site_id} | Get a site
[**GetSites**](SitesApi.md#GetSites) | **Get** /delivery/v1/stacks/{stack_id}/sites | Get all sites



## CreateSite

> DeliveryCreateSiteResponse CreateSite(ctx, stackId).DeliveryCreateSiteRequest(deliveryCreateSiteRequest).Execute()

Create a site

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    stackId := "stackId_example" // string | A stack ID or slug
    deliveryCreateSiteRequest := openapiclient.deliveryCreateSiteRequest{Domain: "Domain_example", Origin: openapiclient.deliveryCreateSiteRequestOrigin{Path: "Path_example", Hostname: "Hostname_example", Port: 123, SecurePort: 123, Http: openapiclient.deliveryHTTPOrigin{Path: "Path_example", Hostname: "Hostname_example", Port: 123, SecurePort: 123, Authentication: openapiclient.deliveryOriginAuthentication{Basic: openapiclient.deliveryBasicAuthentication{Username: "Username_example", Password: "Password_example"}, S3: openapiclient.deliveryS3Authentication{AccessKey: "AccessKey_example", SecretKey: "SecretKey_example"}}, VerifyCertificate: false, CertificateCommonName: "CertificateCommonName_example"}, StackpathStorage: openapiclient.deliveryStackPathStorageOrigin{BucketName: "BucketName_example", BucketRegion: "BucketRegion_example", Authentication: openapiclient.deliveryOriginAuthentication{Basic: openapiclient.deliveryBasicAuthentication{Username: "Username_example", Password: "Password_example"}, S3: openapiclient.deliveryS3Authentication{AccessKey: "AccessKey_example", SecretKey: "SecretKey_example"}}}, S3Storage: openapiclient.deliveryAWSS3Origin{BucketName: "BucketName_example", BucketRegion: "BucketRegion_example", Authentication: }, GoogleCloudStorage: openapiclient.deliveryGoogleStorageOrigin{BucketName: "BucketName_example"}}, Features: []DeliveryCreateSiteRequestFeature{openapiclient.deliveryCreateSiteRequestFeature{}), Configuration: openapiclient.custconfConfiguration{AccessLogs: openapiclient.custconfAccessLogs{Id: "Id_example", Enabled: false}, AccessLogsConfig: openapiclient.custconfAccessLogsConfig{Id: "Id_example", ExtraLogFields: "ExtraLogFields_example", Enabled: false}, AuthACL: []CustconfAuthACL{openapiclient.custconfAuthACL{Id: "Id_example", AccessCode: openapiclient.AuthACLAccessCodeEnumWrapperValue{}, IpList: "IpList_example", Protocol: openapiclient.custconfAuthACLProtocolEnumWrapperValue{}, ClientIPSrc: openapiclient.AuthACLClientIPSrcEnumWrapperValue{}, Header: "Header_example", Enabled: false}), AuthGeo: []CustconfAuthGeo{openapiclient.custconfAuthGeo{Id: "Id_example", Code: openapiclient.AuthGeoCodeEnumWrapperValue{}, Values: "Values_example", Enabled: false}), AuthHttpBasic: openapiclient.custconfAuthHttpBasic{Id: "Id_example", BindingPoint: "BindingPoint_example", Realm: "Realm_example", Ttl: 123, Enabled: false}, AuthReferer: openapiclient.custconfAuthReferer{Id: "Id_example", Referer: "Referer_example", Enabled: false}, AuthUrlSign: []CustconfAuthUrlSign{openapiclient.custconfAuthUrlSign{Id: "Id_example", TokenField: "TokenField_example", IgnoreFieldsAfterToken: false, PassPhraseField: "PassPhraseField_example", PassPhrase: "PassPhrase_example", ExpiresField: "ExpiresField_example", IpAddressField: "IpAddressField_example", UriLengthField: "UriLengthField_example", UserAgentField: "UserAgentField_example", Enabled: false, MethodFilter: "MethodFilter_example", PathFilter: "PathFilter_example", HeaderFilter: "HeaderFilter_example"}), AuthUrlSignAliCloudA: []CustconfAuthUrlSignAliCloudA{openapiclient.custconfAuthUrlSignAliCloudA{Id: "Id_example", PassPhrase: "PassPhrase_example", TokenField: "TokenField_example", IncludeParamsBeforeToken: false, ExpirationExtension: 123, Enabled: false, MethodFilter: "MethodFilter_example", PathFilter: "PathFilter_example", HeaderFilter: "HeaderFilter_example"}), AuthUrlSignAliCloudB: []CustconfAuthUrlSignAliCloudB{openapiclient.custconfAuthUrlSignAliCloudB{Id: "Id_example", PassPhrase: "PassPhrase_example", ExpirationExtension: 123, Enabled: false, MethodFilter: "MethodFilter_example", PathFilter: "PathFilter_example", HeaderFilter: "HeaderFilter_example"}), AuthUrlSignAliCloudC: []CustconfAuthUrlSignAliCloudC{openapiclient.custconfAuthUrlSignAliCloudC{Id: "Id_example", PassPhrase: "PassPhrase_example", ExpirationExtension: 123, TokenField: "TokenField_example", ExpireField: "ExpireField_example", Enabled: false, MethodFilter: "MethodFilter_example", PathFilter: "PathFilter_example", HeaderFilter: "HeaderFilter_example"}), AuthUrlSignHmacTlu: []CustconfAuthUrlSignHmacTlu{openapiclient.custconfAuthUrlSignHmacTlu{Id: "Id_example", ExpireParameterName: "ExpireParameterName_example", KeyIdParameterName: "KeyIdParameterName_example", AlgorithmIdParameterName: "AlgorithmIdParameterName_example", DigestParameterName: "DigestParameterName_example", SymmetricKeyIdMap: map[string]string{ "Key" = "Value" }, AlgorithmIdMap: map[string]string{ "Key" = "Value" }, Enabled: false, MethodFilter: "MethodFilter_example", PathFilter: "PathFilter_example", HeaderFilter: "HeaderFilter_example"}), AuthUrlSignIq: []CustconfAuthUrlSignIq{openapiclient.custconfAuthUrlSignIq{Id: "Id_example", SecretKey: "SecretKey_example", Enabled: false, MethodFilter: "MethodFilter_example", PathFilter: "PathFilter_example", HeaderFilter: "HeaderFilter_example"}), AuthUrlAsymmetricSignTlu: []CustconfAuthUrlAsymmetricSignTlu{openapiclient.custconfAuthUrlAsymmetricSignTlu{Id: "Id_example", ExpireParameterName: "ExpireParameterName_example", KeyIdParameterName: "KeyIdParameterName_example", AlgorithmIdParameterName: "AlgorithmIdParameterName_example", DigestParameterName: "DigestParameterName_example", PublicKeyIdMap: map[string]string{ "Key" = "Value" }, AlgorithmIdMap: map[string]string{ "Key" = "Value" }, Enabled: false, MethodFilter: "MethodFilter_example", PathFilter: "PathFilter_example", HeaderFilter: "HeaderFilter_example"}), BandWidthLimit: openapiclient.custconfBandWidthLimit{Id: "Id_example", Rule: "Rule_example", Values: "Values_example", Enabled: false}, BandwidthRateLimit: openapiclient.custconfBandwidthRateLimit{Id: "Id_example", InitialBurstName: "InitialBurstName_example", SustainedRateName: "SustainedRateName_example", InitialBurstUnits: openapiclient.BandwidthRateLimitInitialBurstUnitsEnumWrapperValue{}, SustainedRateUnits: openapiclient.BandwidthRateLimitSustainedRateUnitsEnumWrapperValue{}, Enabled: false, MethodFilter: "MethodFilter_example", PathFilter: "PathFilter_example", HeaderFilter: "HeaderFilter_example"}, BypassCache: []CustconfBypassCache{openapiclient.custconfBypassCache{Id: "Id_example", Enabled: false, MethodFilter: "MethodFilter_example", PathFilter: "PathFilter_example", HeaderFilter: "HeaderFilter_example", CookieFilter: "CookieFilter_example"}), CacheControl: []CustconfCacheControl{openapiclient.custconfCacheControl{Id: "Id_example", StatusCodeMatch: "StatusCodeMatch_example", MaxAge: 123, MustRevalidate: false, SynchronizeMaxAge: false, Override: "Override_example", Enabled: false}), CacheKeyModification: openapiclient.custconfCacheKeyModification{Id: "Id_example", NormalizeKeyPathToLowerCase: false, Enabled: false}, ClientRequestModification: []CustconfClientRequestModification{openapiclient.custconfClientRequestModification{Id: "Id_example", UrlPattern: "UrlPattern_example", UrlRewrite: "UrlRewrite_example", HeaderPattern: "HeaderPattern_example", HeaderRewrite: "HeaderRewrite_example", AddHeaders: "AddHeaders_example", FlowControl: openapiclient.custconfClientRequestModificationFlowControlEnumWrapperValue{}, Enabled: false, MethodFilter: "MethodFilter_example", PathFilter: "PathFilter_example", HeaderFilter: "HeaderFilter_example", CookieFilter: "CookieFilter_example"}), ClientResponseModification: []CustconfClientResponseModification{openapiclient.custconfClientResponseModification{Id: "Id_example", StatusCodeRewrite: 123, HeaderPattern: "HeaderPattern_example", HeaderRewrite: "HeaderRewrite_example", AddHeaders: "AddHeaders_example", FlowControl: openapiclient.custconfClientResponseModificationFlowControlEnumWrapperValue{}, Enabled: false, MethodFilter: "MethodFilter_example", PathFilter: "PathFilter_example", HeaderFilter: "HeaderFilter_example"}), Compression: openapiclient.custconfCompression{Id: "Id_example", Gzip: "Gzip_example", Mime: "Mime_example", Level: "Level_example", Enabled: false}, ContentDispositionByURL: openapiclient.custconfContentDispositionByURL{Id: "Id_example", DispositionNameQSParam: "DispositionNameQSParam_example", DispositionTypeQSParam: "DispositionTypeQSParam_example", DispositionOverrideQSParam: "DispositionOverrideQSParam_example", OverrideOriginHeader: false, Enabled: false}, ContentDispositionByHeader: []CustconfContentDispositionByHeader{openapiclient.custconfContentDispositionByHeader{Id: "Id_example", HeaderFieldName: "HeaderFieldName_example", HeaderValueMatch: "HeaderValueMatch_example", DefaultType: openapiclient.ContentDispositionByHeaderDefaultTypeEnumWrapperValue{}, OverrideOriginHeader: false, Enabled: false, MethodFilter: "MethodFilter_example", PathFilter: "PathFilter_example", HeaderFilter: "HeaderFilter_example"}), Customer: openapiclient.custconfCustomer{Id: "Id_example", AccessLogFields: "AccessLogFields_example", OpLogFields: "OpLogFields_example", ReceiptLogFields: "ReceiptLogFields_example"}, CustomHeader: openapiclient.custconfCustomHeader{Id: "Id_example", XForwardedForAuth: "XForwardedForAuth_example", XForwardedForOrigin: "XForwardedForOrigin_example", Enabled: false}, CustomMimeType: []CustconfCustomMimeType{openapiclient.custconfCustomMimeType{Id: "Id_example", Code: "Code_example", ExtensionMap: "ExtensionMap_example", Enabled: false, MethodFilter: "MethodFilter_example", PathFilter: "PathFilter_example", HeaderFilter: "HeaderFilter_example"}), DynamicCacheRule: []CustconfDynamicCacheRule{openapiclient.custconfDynamicCacheRule{Id: "Id_example", StatusCode: 123, Headers: "Headers_example", Enabled: false, MethodFilter: "MethodFilter_example", PathFilter: "PathFilter_example", HeaderFilter: "HeaderFilter_example"}), DynamicContent: []CustconfDynamicContent{openapiclient.custconfDynamicContent{Id: "Id_example", QueryParams: "QueryParams_example", HeaderFields: "HeaderFields_example", Enabled: false, MethodFilter: "MethodFilter_example", PathFilter: "PathFilter_example", HeaderFilter: "HeaderFilter_example"}), FailSafeOriginPull: openapiclient.custconfFailSafeOriginPull{Id: "Id_example", Enabled: false, StatusCodeMatch: "StatusCodeMatch_example", PathFilter: "PathFilter_example"}, FarAheadRangeProxy: openapiclient.custconfFarAheadRangeProxy{Id: "Id_example", Enabled: false, ThresholdBytes: 123}, FileSegmentation: openapiclient.custconfFileSegmentation{Id: "Id_example", Enabled: false}, FlvPseudoStreaming: openapiclient.custconfFlvPseudoStreaming{Id: "Id_example", JumpToByteInitialBytesParam: "JumpToByteInitialBytesParam_example", JumpToByteStartOffsetParam: "JumpToByteStartOffsetParam_example", Enabled: false}, GzipOriginPull: openapiclient.custconfGzipOriginPull{Id: "Id_example", Enabled: false}, HttpMethods: openapiclient.custconfHttpMethods{Id: "Id_example", PassThru: "PassThru_example", Enabled: false}, Origin: []CdncustconfOrigin{openapiclient.cdncustconfOrigin{Id: "Id_example", Host: "Host_example", OriginPullHeaders: "OriginPullHeaders_example", OriginCacheHeaders: "OriginCacheHeaders_example"}), OriginPersistentConnections: openapiclient.custconfOriginPersistentConnections{Id: "Id_example", Enabled: false}, OriginPull: openapiclient.custconfOriginPull{Id: "Id_example", RedirectAction: openapiclient.OriginPullRedirectActionEnumWrapperValue{}, NoQSParams: false, RetryMethods: "RetryMethods_example", Enabled: false}, OriginPullCacheExtension: openapiclient.custconfOriginPullCacheExtension{Id: "Id_example", ExpiredCacheExtension: 123, OriginUnreachableCacheExtension: 123, Enabled: false}, OriginPullHost: []CustconfOriginPullHost{openapiclient.custconfOriginPullHost{Id: "Id_example", OriginUrl: "OriginUrl_example", UserName: "UserName_example", Password: "Password_example", MethodFilter: "MethodFilter_example", PathFilter: "PathFilter_example", HeaderFilter: "HeaderFilter_example"}), OriginPullProtocol: openapiclient.custconfOriginPullProtocol{Id: "Id_example", Protocol: openapiclient.custconfOriginPullProtocolProtocolEnumWrapperValue{}, EnableSNI: false}, OriginPullLogs: openapiclient.custconfOriginPullLogs{Id: "Id_example", Enabled: false}, OriginPullLogsConfig: openapiclient.custconfOriginPullLogsConfig{Id: "Id_example", ExtraLogFields: "ExtraLogFields_example", Enabled: false}, OriginPullPolicy: []CustconfOriginPullPolicy{openapiclient.custconfOriginPullPolicy{Id: "Id_example", StatusCodeMatch: "StatusCodeMatch_example", ExpirePolicy: openapiclient.OriginPullPolicyExpirePolicyEnumWrapperValue{}, ExpireSeconds: 123, HonorNoStore: false, HonorNoCache: false, HonorMustRevalidate: false, NoCacheBehavior: openapiclient.OriginPullPolicyNoCacheBehaviorEnumWrapperValue{}, MaxAgeZeroToNoCache: false, MustRevalidateToNoCache: false, BypassCacheIdentifier: "BypassCacheIdentifier_example", ForceBypassCache: false, HttpHeaders: "HttpHeaders_example", HonorPrivate: false, HonorSMaxAge: false, UpdateHttpHeadersOn304Response: false, DefaultCacheBehavior: openapiclient.OriginPullPolicyDefaultCacheBehaviorEnumWrapperValue{}, Enabled: false, MethodFilter: "MethodFilter_example", PathFilter: "PathFilter_example", HeaderFilter: "HeaderFilter_example", ContentTypeFilter: "ContentTypeFilter_example"}), OriginRequestModification: []CustconfOriginRequestModification{openapiclient.custconfOriginRequestModification{Id: "Id_example", UrlPattern: "UrlPattern_example", UrlRewrite: "UrlRewrite_example", HeaderPattern: "HeaderPattern_example", HeaderRewrite: "HeaderRewrite_example", AddHeaders: "AddHeaders_example", FlowControl: "FlowControl_example", Enabled: false, MethodFilter: "MethodFilter_example", PathFilter: "PathFilter_example", HeaderFilter: "HeaderFilter_example", CookieFilter: "CookieFilter_example"}), OriginResponseModification: []CustconfOriginResponseModification{openapiclient.custconfOriginResponseModification{Id: "Id_example", StatusCodeRewrite: 123, HeaderPattern: "HeaderPattern_example", HeaderRewrite: "HeaderRewrite_example", AddHeaders: "AddHeaders_example", FlowControl: openapiclient.custconfOriginResponseModificationFlowControlEnumWrapperValue{}, Enabled: false, MethodFilter: "MethodFilter_example", PathFilter: "PathFilter_example", HeaderFilter: "HeaderFilter_example"}), QueryStrParam: openapiclient.custconfQueryStrParam{Id: "Id_example", DispositionName: "DispositionName_example", DispositionType: "DispositionType_example", DispositionOverride: "DispositionOverride_example", JumpToTimeStart: "JumpToTimeStart_example", JumpToTimeEnd: "JumpToTimeEnd_example", JumpToByteInitialBytes: "JumpToByteInitialBytes_example", JumpToByteStartOffset: "JumpToByteStartOffset_example", RateLimitInitial: "RateLimitInitial_example", RateLimitSustained: "RateLimitSustained_example", Enabled: false}, ReceiptLogsConfig: openapiclient.custconfReceiptLogsConfig{Id: "Id_example", ExtraLogFields: "ExtraLogFields_example", Enabled: false}, RedirectExceptions: openapiclient.custconfRedirectExceptions{Id: "Id_example", RedirectAgentCode: "RedirectAgentCode_example", Enabled: false}, RedirectMappings: []CustconfRedirectMappings{openapiclient.custconfRedirectMappings{Id: "Id_example", Code: 123, RedirectURL: "RedirectURL_example", ReplacementToken: "ReplacementToken_example", Enabled: false, MethodFilter: "MethodFilter_example", PathFilter: "PathFilter_example", HeaderFilter: "HeaderFilter_example"}), ResponseHeader: openapiclient.custconfResponseHeader{Id: "Id_example", Http: "Http_example", EnableETag: false, Enabled: false}, OriginPullResumeDownload: []CustconfOriginPullResumeDownload{openapiclient.custconfOriginPullResumeDownload{Id: "Id_example", Enabled: false, OriginalStatusCodeMatch: "OriginalStatusCodeMatch_example", MinimumBodyBytesConsumed: "MinimumBodyBytesConsumed_example", MaximumAttempts: 123, RequireOriginRangeSupport: false, EtagValidation: openapiclient.OriginPullResumeDownloadEtagValidationEnumWrapperValue{}, HeaderFilter: "HeaderFilter_example", PathFilter: "PathFilter_example"}), RobotsTxt: []CustconfRobotsTxt{openapiclient.custconfRobotsTxt{Id: "Id_example", File: "File_example", CacheControlHeader: "CacheControlHeader_example", Enabled: false, MethodFilter: "MethodFilter_example", PathFilter: "PathFilter_example", HeaderFilter: "HeaderFilter_example"}), AwsSignedOriginPullV4: []CustconfAwsSignedOriginPullV4{openapiclient.custconfAwsSignedOriginPullV4{Id: "Id_example", Enabled: false, BucketIdentifier: "BucketIdentifier_example", AccessKeyId: "AccessKeyId_example", SecretAccessKey: "SecretAccessKey_example", AwsRegion: "AwsRegion_example", AwsService: "AwsService_example", ExpireTimeSeconds: 123, SignedHeaders: "SignedHeaders_example", AuthenticationType: openapiclient.custconfAwsSignedOriginPullV4AuthenticationTypeEnumWrapperValue{}, HeaderFilter: "HeaderFilter_example", MethodFilter: "MethodFilter_example", PathFilter: "PathFilter_example"}), AwsSignedS3PostV4: []CustconfAwsSignedS3PostV4{openapiclient.custconfAwsSignedS3PostV4{Id: "Id_example", Enabled: false, BucketIdentifier: "BucketIdentifier_example", AccessKeyId: "AccessKeyId_example", SecretAccessKey: "SecretAccessKey_example", AwsRegion: "AwsRegion_example", AwsService: "AwsService_example", ExpireTimeSeconds: 123, SignedHeaders: "SignedHeaders_example", AuthenticationType: openapiclient.custconfAwsSignedS3PostV4AuthenticationTypeEnumWrapperValue{}, HeaderFilter: "HeaderFilter_example", MethodFilter: "MethodFilter_example", PathFilter: "PathFilter_example"}), StaticHeader: []CustconfStaticHeader{openapiclient.custconfStaticHeader{Id: "Id_example", ClientRequest: "ClientRequest_example", Http: "Http_example", OriginPull: "OriginPull_example", Enabled: false, MethodFilter: "MethodFilter_example", PathFilter: "PathFilter_example", HeaderFilter: "HeaderFilter_example"}), TimePseudoStreaming: openapiclient.custconfTimePseudoStreaming{Id: "Id_example", JumpToTimeStartParam: "JumpToTimeStartParam_example", JumpToTimeEndParam: "JumpToTimeEndParam_example", Enabled: false}, Http2Support: openapiclient.custconfHttp2Support{Id: "Id_example", Enabled: false}, VHost: []CustconfVHost{openapiclient.custconfVHost{Id: "Id_example", Domain: "Domain_example", Enabled: false})}} // DeliveryCreateSiteRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesApi.CreateSite(context.Background(), stackId, deliveryCreateSiteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesApi.CreateSite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSite`: DeliveryCreateSiteResponse
    fmt.Fprintf(os.Stdout, "Response from `SitesApi.CreateSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deliveryCreateSiteRequest** | [**DeliveryCreateSiteRequest**](DeliveryCreateSiteRequest.md) |  | 

### Return type

[**DeliveryCreateSiteResponse**](deliveryCreateSiteResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSite

> DeleteSite(ctx, stackId, siteId).Execute()

Delete a site

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    stackId := "stackId_example" // string | A stack ID or slug
    siteId := "siteId_example" // string | A site ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesApi.DeleteSite(context.Background(), stackId, siteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesApi.DeleteSite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**siteId** | **string** | A site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSite

> DeliveryGetSiteResponse GetSite(ctx, stackId, siteId).Execute()

Get a site

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    stackId := "stackId_example" // string | A stack ID or slug
    siteId := "siteId_example" // string | A site ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesApi.GetSite(context.Background(), stackId, siteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesApi.GetSite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSite`: DeliveryGetSiteResponse
    fmt.Fprintf(os.Stdout, "Response from `SitesApi.GetSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**siteId** | **string** | A site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeliveryGetSiteResponse**](deliveryGetSiteResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSites

> DeliveryGetSitesResponse GetSites(ctx, stackId).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).Execute()

Get all sites

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    stackId := "stackId_example" // string | A stack ID or slug
    pageRequestFirst := "pageRequestFirst_example" // string | The number of items desired. (optional)
    pageRequestAfter := "pageRequestAfter_example" // string | The cursor value after which data will be returned. (optional)
    pageRequestFilter := "pageRequestFilter_example" // string | SQL-style constraint filters. (optional)
    pageRequestSortBy := "pageRequestSortBy_example" // string | Sort the response by the given field. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesApi.GetSites(context.Background(), stackId).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesApi.GetSites``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSites`: DeliveryGetSitesResponse
    fmt.Fprintf(os.Stdout, "Response from `SitesApi.GetSites`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageRequestFirst** | **string** | The number of items desired. | 
 **pageRequestAfter** | **string** | The cursor value after which data will be returned. | 
 **pageRequestFilter** | **string** | SQL-style constraint filters. | 
 **pageRequestSortBy** | **string** | Sort the response by the given field. | 

### Return type

[**DeliveryGetSitesResponse**](deliveryGetSitesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

