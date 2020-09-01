# \ServerlessScriptingApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSiteScript**](ServerlessScriptingApi.md#CreateSiteScript) | **Post** /cdn/v1/stacks/{stack_id}/sites/{site_id}/scripts | Create a serverless script
[**DeleteSiteScript**](ServerlessScriptingApi.md#DeleteSiteScript) | **Delete** /cdn/v1/stacks/{stack_id}/sites/{site_id}/scripts/{script_id} | Delete a serverless script
[**GetSiteScript**](ServerlessScriptingApi.md#GetSiteScript) | **Get** /cdn/v1/stacks/{stack_id}/sites/{site_id}/scripts/{script_id} | Get a serverless script
[**GetSiteScript2**](ServerlessScriptingApi.md#GetSiteScript2) | **Get** /cdn/v1/stacks/{stack_id}/sites/{site_id}/scripts/{script_id}/{script_version} | Get a serverless script version
[**GetSiteScripts**](ServerlessScriptingApi.md#GetSiteScripts) | **Get** /cdn/v1/stacks/{stack_id}/sites/{site_id}/scripts | Get all serverless scripts
[**UpdateSiteScript**](ServerlessScriptingApi.md#UpdateSiteScript) | **Patch** /cdn/v1/stacks/{stack_id}/sites/{site_id}/scripts/{script_id} | Update a serverless script



## CreateSiteScript

> CdnCreateSiteScriptResponse CreateSiteScript(ctx, stackId, siteId).CdnCreateSiteScriptRequest(cdnCreateSiteScriptRequest).Execute()

Create a serverless script

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
    cdnCreateSiteScriptRequest := openapiclient.cdnCreateSiteScriptRequest{Name: "Name_example", Code: 123, Paths: []string{"Paths_example")} // CdnCreateSiteScriptRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServerlessScriptingApi.CreateSiteScript(context.Background(), stackId, siteId, cdnCreateSiteScriptRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerlessScriptingApi.CreateSiteScript``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSiteScript`: CdnCreateSiteScriptResponse
    fmt.Fprintf(os.Stdout, "Response from `ServerlessScriptingApi.CreateSiteScript`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**siteId** | **string** | A site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSiteScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cdnCreateSiteScriptRequest** | [**CdnCreateSiteScriptRequest**](CdnCreateSiteScriptRequest.md) |  | 

### Return type

[**CdnCreateSiteScriptResponse**](cdnCreateSiteScriptResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSiteScript

> DeleteSiteScript(ctx, stackId, siteId, scriptId).Execute()

Delete a serverless script

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
    scriptId := "scriptId_example" // string | A serverless script ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServerlessScriptingApi.DeleteSiteScript(context.Background(), stackId, siteId, scriptId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerlessScriptingApi.DeleteSiteScript``: %v\n", err)
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
**scriptId** | **string** | A serverless script ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSiteScriptRequest struct via the builder pattern


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


## GetSiteScript

> CdnGetSiteScriptResponse GetSiteScript(ctx, stackId, siteId, scriptId).ScriptVersion(scriptVersion).Execute()

Get a serverless script

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
    scriptId := "scriptId_example" // string | A serverless script ID
    scriptVersion := "scriptVersion_example" // string | The version of the serverless script to get (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServerlessScriptingApi.GetSiteScript(context.Background(), stackId, siteId, scriptId).ScriptVersion(scriptVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerlessScriptingApi.GetSiteScript``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSiteScript`: CdnGetSiteScriptResponse
    fmt.Fprintf(os.Stdout, "Response from `ServerlessScriptingApi.GetSiteScript`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**siteId** | **string** | A site ID | 
**scriptId** | **string** | A serverless script ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **scriptVersion** | **string** | The version of the serverless script to get | 

### Return type

[**CdnGetSiteScriptResponse**](cdnGetSiteScriptResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteScript2

> CdnGetSiteScriptResponse GetSiteScript2(ctx, stackId, siteId, scriptId, scriptVersion).Execute()

Get a serverless script version

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
    scriptId := "scriptId_example" // string | A serverless script ID
    scriptVersion := "scriptVersion_example" // string | The version of the serverless script to get

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServerlessScriptingApi.GetSiteScript2(context.Background(), stackId, siteId, scriptId, scriptVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerlessScriptingApi.GetSiteScript2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSiteScript2`: CdnGetSiteScriptResponse
    fmt.Fprintf(os.Stdout, "Response from `ServerlessScriptingApi.GetSiteScript2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**siteId** | **string** | A site ID | 
**scriptId** | **string** | A serverless script ID | 
**scriptVersion** | **string** | The version of the serverless script to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteScript2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**CdnGetSiteScriptResponse**](cdnGetSiteScriptResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteScripts

> CdnGetSiteScriptsResponse GetSiteScripts(ctx, stackId, siteId).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).Execute()

Get all serverless scripts

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
    resp, r, err := api_client.ServerlessScriptingApi.GetSiteScripts(context.Background(), stackId, siteId).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerlessScriptingApi.GetSiteScripts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSiteScripts`: CdnGetSiteScriptsResponse
    fmt.Fprintf(os.Stdout, "Response from `ServerlessScriptingApi.GetSiteScripts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**siteId** | **string** | A site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteScriptsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageRequestFirst** | **string** | The number of items desired. | 
 **pageRequestAfter** | **string** | The cursor value after which data will be returned. | 
 **pageRequestFilter** | **string** | SQL-style constraint filters. | 
 **pageRequestSortBy** | **string** | Sort the response by the given field. | 

### Return type

[**CdnGetSiteScriptsResponse**](cdnGetSiteScriptsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSiteScript

> CdnUpdateSiteScriptResponse UpdateSiteScript(ctx, stackId, siteId, scriptId).CdnUpdateSiteScriptRequest(cdnUpdateSiteScriptRequest).Execute()

Update a serverless script

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
    scriptId := "scriptId_example" // string | A serverless script ID
    cdnUpdateSiteScriptRequest := openapiclient.cdnUpdateSiteScriptRequest{Code: 123, Paths: []string{"Paths_example")} // CdnUpdateSiteScriptRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServerlessScriptingApi.UpdateSiteScript(context.Background(), stackId, siteId, scriptId, cdnUpdateSiteScriptRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerlessScriptingApi.UpdateSiteScript``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSiteScript`: CdnUpdateSiteScriptResponse
    fmt.Fprintf(os.Stdout, "Response from `ServerlessScriptingApi.UpdateSiteScript`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**siteId** | **string** | A site ID | 
**scriptId** | **string** | A serverless script ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSiteScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **cdnUpdateSiteScriptRequest** | [**CdnUpdateSiteScriptRequest**](CdnUpdateSiteScriptRequest.md) |  | 

### Return type

[**CdnUpdateSiteScriptResponse**](cdnUpdateSiteScriptResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

