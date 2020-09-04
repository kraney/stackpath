# \RulesApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkDeleteRules**](RulesApi.md#BulkDeleteRules) | **Post** /waf/v1/stacks/{stack_id}/sites/{site_id}/rules/bulk_delete | Delete multiple rules
[**CreateRule**](RulesApi.md#CreateRule) | **Post** /waf/v1/stacks/{stack_id}/sites/{site_id}/rules | Create a rule
[**DeleteRule**](RulesApi.md#DeleteRule) | **Delete** /waf/v1/stacks/{stack_id}/sites/{site_id}/rules/{rule_id} | Delete a rule
[**DisableRule**](RulesApi.md#DisableRule) | **Post** /waf/v1/stacks/{stack_id}/sites/{site_id}/rules/{rule_id}/disable | Disable a rule
[**EnableRule**](RulesApi.md#EnableRule) | **Post** /waf/v1/stacks/{stack_id}/sites/{site_id}/rules/{rule_id}/enable | Enable a rule
[**GetRule**](RulesApi.md#GetRule) | **Get** /waf/v1/stacks/{stack_id}/sites/{site_id}/rules/{rule_id} | Get a rule
[**GetRules**](RulesApi.md#GetRules) | **Get** /waf/v1/stacks/{stack_id}/sites/{site_id}/rules | Get all rules
[**UpdateRule**](RulesApi.md#UpdateRule) | **Patch** /waf/v1/stacks/{stack_id}/sites/{site_id}/rules/{rule_id} | Update a rule



## BulkDeleteRules

> BulkDeleteRules(ctx, stackId, siteId).WafBulkDeleteRulesRequest(wafBulkDeleteRulesRequest).Execute()

Delete multiple rules

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
    wafBulkDeleteRulesRequest := openapiclient.wafBulkDeleteRulesRequest{RuleIds: []string{"RuleIds_example")} // WafBulkDeleteRulesRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RulesApi.BulkDeleteRules(context.Background(), stackId, siteId, wafBulkDeleteRulesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.BulkDeleteRules``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBulkDeleteRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **wafBulkDeleteRulesRequest** | [**WafBulkDeleteRulesRequest**](WafBulkDeleteRulesRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRule

> WafCreateRuleResponse CreateRule(ctx, stackId, siteId).WafCreateRuleRequest(wafCreateRuleRequest).Execute()

Create a rule

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
    wafCreateRuleRequest := openapiclient.wafCreateRuleRequest{Name: "Name_example", Description: "Description_example", Conditions: []RuleCondition{openapiclient.RuleCondition{Ip: openapiclient.ConditionIpCondition{IpAddress: "IpAddress_example"}, IpRange: openapiclient.ConditionIpRangeCondition{LowerBound: "LowerBound_example", UpperBound: "UpperBound_example"}, Url: openapiclient.ConditionUrlCondition{Url: "Url_example", ExactMatch: false}, UserAgent: openapiclient.ConditionUserAgentCondition{UserAgent: "UserAgent_example", ExactMatch: false}, Header: openapiclient.ConditionHeaderCondition{Header: "Header_example", Value: "Value_example", ExactMatch: false}, HeaderExists: openapiclient.ConditionHeaderExistsCondition{Header: "Header_example"}, ResponseHeader: openapiclient.ConditionResponseHeaderCondition{Header: "Header_example", Value: "Value_example", ExactMatch: false}, ResponseHeaderExists: openapiclient.ConditionResponseHeaderExistsCondition{Header: "Header_example"}, HttpMethod: openapiclient.ConditionHttpMethodCondition{HttpMethod: openapiclient.wafHttpMethod{}}, FileExtension: openapiclient.ConditionFileExtensionCondition{FileExtension: "FileExtension_example"}, ContentType: openapiclient.ConditionContentTypeCondition{ContentType: "ContentType_example"}, Country: openapiclient.ConditionCountryCondition{CountryCode: "CountryCode_example"}, Organization: openapiclient.ConditionOrganizationCondition{Organization: "Organization_example"}, RequestRate: openapiclient.ConditionRequestRateCondition{Ips: []string{"Ips_example"), HttpMethods: []WafHttpMethod{openapiclient.wafHttpMethod{}), PathPattern: "PathPattern_example", Requests: "Requests_example", Time: "Time_example"}, OwnerTypes: openapiclient.ConditionOwnerTypeCondition{OwnerTypes: []OwnerTypeConditionOwnerType{openapiclient.OwnerTypeConditionOwnerType{})}, Tags: openapiclient.ConditionTagCondition{Tags: []string{"Tags_example")}, SessionRequestCount: openapiclient.ConditionSessionRequestCountCondition{RequestCount: "RequestCount_example"}, Negation: false}), Action: openapiclient.wafRuleAction{}, Enabled: false, StatusCode: openapiclient.RuleStatusCode{}, ActionDuration: "ActionDuration_example"} // WafCreateRuleRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RulesApi.CreateRule(context.Background(), stackId, siteId, wafCreateRuleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.CreateRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRule`: WafCreateRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `RulesApi.CreateRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**siteId** | **string** | A site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **wafCreateRuleRequest** | [**WafCreateRuleRequest**](WafCreateRuleRequest.md) |  | 

### Return type

[**WafCreateRuleResponse**](wafCreateRuleResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRule

> DeleteRule(ctx, stackId, siteId, ruleId).Execute()

Delete a rule

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
    ruleId := "ruleId_example" // string | A WAF rule ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RulesApi.DeleteRule(context.Background(), stackId, siteId, ruleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.DeleteRule``: %v\n", err)
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
**ruleId** | **string** | A WAF rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRuleRequest struct via the builder pattern


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


## DisableRule

> DisableRule(ctx, stackId, siteId, ruleId).Execute()

Disable a rule

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
    ruleId := "ruleId_example" // string | A WAF rule ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RulesApi.DisableRule(context.Background(), stackId, siteId, ruleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.DisableRule``: %v\n", err)
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
**ruleId** | **string** | A WAF rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableRuleRequest struct via the builder pattern


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


## EnableRule

> EnableRule(ctx, stackId, siteId, ruleId).Execute()

Enable a rule

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
    ruleId := "ruleId_example" // string | A WAF rule ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RulesApi.EnableRule(context.Background(), stackId, siteId, ruleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.EnableRule``: %v\n", err)
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
**ruleId** | **string** | A WAF rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableRuleRequest struct via the builder pattern


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


## GetRule

> WafGetRuleResponse GetRule(ctx, stackId, siteId, ruleId).Execute()

Get a rule

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
    ruleId := "ruleId_example" // string | A WAF rule ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RulesApi.GetRule(context.Background(), stackId, siteId, ruleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.GetRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRule`: WafGetRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `RulesApi.GetRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**siteId** | **string** | A site ID | 
**ruleId** | **string** | A WAF rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**WafGetRuleResponse**](wafGetRuleResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRules

> WafGetRulesResponse GetRules(ctx, stackId, siteId).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).Execute()

Get all rules

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
    pageRequestFirst := "pageRequestFirst_example" // string | The number of items desired. (optional)
    pageRequestAfter := "pageRequestAfter_example" // string | The cursor value after which data will be returned. (optional)
    pageRequestFilter := "pageRequestFilter_example" // string | SQL-style constraint filters. (optional)
    pageRequestSortBy := "pageRequestSortBy_example" // string | Sort the response by the given field. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RulesApi.GetRules(context.Background(), stackId, siteId).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.GetRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRules`: WafGetRulesResponse
    fmt.Fprintf(os.Stdout, "Response from `RulesApi.GetRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**siteId** | **string** | A site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageRequestFirst** | **string** | The number of items desired. | 
 **pageRequestAfter** | **string** | The cursor value after which data will be returned. | 
 **pageRequestFilter** | **string** | SQL-style constraint filters. | 
 **pageRequestSortBy** | **string** | Sort the response by the given field. | 

### Return type

[**WafGetRulesResponse**](wafGetRulesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRule

> WafUpdateRuleResponse UpdateRule(ctx, stackId, siteId, ruleId).WafUpdateRuleRequest(wafUpdateRuleRequest).Execute()

Update a rule



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
    ruleId := "ruleId_example" // string | A WAF rule ID
    wafUpdateRuleRequest := openapiclient.wafUpdateRuleRequest{Name: "Name_example", Description: "Description_example", Conditions: []RuleCondition{openapiclient.RuleCondition{Ip: openapiclient.ConditionIpCondition{IpAddress: "IpAddress_example"}, IpRange: openapiclient.ConditionIpRangeCondition{LowerBound: "LowerBound_example", UpperBound: "UpperBound_example"}, Url: openapiclient.ConditionUrlCondition{Url: "Url_example", ExactMatch: false}, UserAgent: openapiclient.ConditionUserAgentCondition{UserAgent: "UserAgent_example", ExactMatch: false}, Header: openapiclient.ConditionHeaderCondition{Header: "Header_example", Value: "Value_example", ExactMatch: false}, HeaderExists: openapiclient.ConditionHeaderExistsCondition{Header: "Header_example"}, ResponseHeader: openapiclient.ConditionResponseHeaderCondition{Header: "Header_example", Value: "Value_example", ExactMatch: false}, ResponseHeaderExists: openapiclient.ConditionResponseHeaderExistsCondition{Header: "Header_example"}, HttpMethod: openapiclient.ConditionHttpMethodCondition{HttpMethod: }, FileExtension: openapiclient.ConditionFileExtensionCondition{FileExtension: "FileExtension_example"}, ContentType: openapiclient.ConditionContentTypeCondition{ContentType: "ContentType_example"}, Country: openapiclient.ConditionCountryCondition{CountryCode: "CountryCode_example"}, Organization: openapiclient.ConditionOrganizationCondition{Organization: "Organization_example"}, RequestRate: openapiclient.ConditionRequestRateCondition{Ips: []string{"Ips_example"), HttpMethods: []WafHttpMethod{), PathPattern: "PathPattern_example", Requests: "Requests_example", Time: "Time_example"}, OwnerTypes: openapiclient.ConditionOwnerTypeCondition{OwnerTypes: []OwnerTypeConditionOwnerType{openapiclient.OwnerTypeConditionOwnerType{})}, Tags: openapiclient.ConditionTagCondition{Tags: []string{"Tags_example")}, SessionRequestCount: openapiclient.ConditionSessionRequestCountCondition{RequestCount: "RequestCount_example"}, Negation: false}), ActionValue: openapiclient.wafRuleAction{}, Enabled: false, StatusCode: openapiclient.RuleStatusCode{}, ActionDuration: "ActionDuration_example"} // WafUpdateRuleRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RulesApi.UpdateRule(context.Background(), stackId, siteId, ruleId, wafUpdateRuleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.UpdateRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRule`: WafUpdateRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `RulesApi.UpdateRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**siteId** | **string** | A site ID | 
**ruleId** | **string** | A WAF rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **wafUpdateRuleRequest** | [**WafUpdateRuleRequest**](WafUpdateRuleRequest.md) |  | 

### Return type

[**WafUpdateRuleResponse**](wafUpdateRuleResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

