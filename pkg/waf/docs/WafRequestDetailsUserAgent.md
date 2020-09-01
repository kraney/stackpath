# WafRequestDetailsUserAgent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullString** | Pointer to **string** | The full user agent string for the WAF request | [optional] 
**Client** | Pointer to **string** | The name of the requesting client, typically the name of the requesting web browser | [optional] 
**ClientVersion** | Pointer to **string** | The version of the requesting user agent&#39;s client, typically also the version of the requesting web browser | [optional] 
**BaseBrowser** | Pointer to **string** | The name of the requesting user agent&#39;s browser | [optional] 
**BaseBrowserVersion** | Pointer to **string** | The version of the requesting user agent&#39;s browser | [optional] 
**Os** | Pointer to **string** | The name of the requesting user agent&#39;s operating system | [optional] 
**Device** | Pointer to **string** | The name of the requesting user agent&#39;s device | [optional] 
**DeviceType** | Pointer to **string** | The type of the requesting user agent&#39;s device | [optional] 
**Cpu** | Pointer to **string** | The name of the requesting user agent&#39;s CPU | [optional] 
**RenderingEngine** | Pointer to **string** | The name of the requesting user agent&#39;s engine | [optional] 

## Methods

### NewWafRequestDetailsUserAgent

`func NewWafRequestDetailsUserAgent() *WafRequestDetailsUserAgent`

NewWafRequestDetailsUserAgent instantiates a new WafRequestDetailsUserAgent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafRequestDetailsUserAgentWithDefaults

`func NewWafRequestDetailsUserAgentWithDefaults() *WafRequestDetailsUserAgent`

NewWafRequestDetailsUserAgentWithDefaults instantiates a new WafRequestDetailsUserAgent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullString

`func (o *WafRequestDetailsUserAgent) GetFullString() string`

GetFullString returns the FullString field if non-nil, zero value otherwise.

### GetFullStringOk

`func (o *WafRequestDetailsUserAgent) GetFullStringOk() (*string, bool)`

GetFullStringOk returns a tuple with the FullString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullString

`func (o *WafRequestDetailsUserAgent) SetFullString(v string)`

SetFullString sets FullString field to given value.

### HasFullString

`func (o *WafRequestDetailsUserAgent) HasFullString() bool`

HasFullString returns a boolean if a field has been set.

### GetClient

`func (o *WafRequestDetailsUserAgent) GetClient() string`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *WafRequestDetailsUserAgent) GetClientOk() (*string, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *WafRequestDetailsUserAgent) SetClient(v string)`

SetClient sets Client field to given value.

### HasClient

`func (o *WafRequestDetailsUserAgent) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetClientVersion

`func (o *WafRequestDetailsUserAgent) GetClientVersion() string`

GetClientVersion returns the ClientVersion field if non-nil, zero value otherwise.

### GetClientVersionOk

`func (o *WafRequestDetailsUserAgent) GetClientVersionOk() (*string, bool)`

GetClientVersionOk returns a tuple with the ClientVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientVersion

`func (o *WafRequestDetailsUserAgent) SetClientVersion(v string)`

SetClientVersion sets ClientVersion field to given value.

### HasClientVersion

`func (o *WafRequestDetailsUserAgent) HasClientVersion() bool`

HasClientVersion returns a boolean if a field has been set.

### GetBaseBrowser

`func (o *WafRequestDetailsUserAgent) GetBaseBrowser() string`

GetBaseBrowser returns the BaseBrowser field if non-nil, zero value otherwise.

### GetBaseBrowserOk

`func (o *WafRequestDetailsUserAgent) GetBaseBrowserOk() (*string, bool)`

GetBaseBrowserOk returns a tuple with the BaseBrowser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseBrowser

`func (o *WafRequestDetailsUserAgent) SetBaseBrowser(v string)`

SetBaseBrowser sets BaseBrowser field to given value.

### HasBaseBrowser

`func (o *WafRequestDetailsUserAgent) HasBaseBrowser() bool`

HasBaseBrowser returns a boolean if a field has been set.

### GetBaseBrowserVersion

`func (o *WafRequestDetailsUserAgent) GetBaseBrowserVersion() string`

GetBaseBrowserVersion returns the BaseBrowserVersion field if non-nil, zero value otherwise.

### GetBaseBrowserVersionOk

`func (o *WafRequestDetailsUserAgent) GetBaseBrowserVersionOk() (*string, bool)`

GetBaseBrowserVersionOk returns a tuple with the BaseBrowserVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseBrowserVersion

`func (o *WafRequestDetailsUserAgent) SetBaseBrowserVersion(v string)`

SetBaseBrowserVersion sets BaseBrowserVersion field to given value.

### HasBaseBrowserVersion

`func (o *WafRequestDetailsUserAgent) HasBaseBrowserVersion() bool`

HasBaseBrowserVersion returns a boolean if a field has been set.

### GetOs

`func (o *WafRequestDetailsUserAgent) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *WafRequestDetailsUserAgent) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *WafRequestDetailsUserAgent) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *WafRequestDetailsUserAgent) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetDevice

`func (o *WafRequestDetailsUserAgent) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *WafRequestDetailsUserAgent) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *WafRequestDetailsUserAgent) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *WafRequestDetailsUserAgent) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetDeviceType

`func (o *WafRequestDetailsUserAgent) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *WafRequestDetailsUserAgent) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *WafRequestDetailsUserAgent) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *WafRequestDetailsUserAgent) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetCpu

`func (o *WafRequestDetailsUserAgent) GetCpu() string`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *WafRequestDetailsUserAgent) GetCpuOk() (*string, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *WafRequestDetailsUserAgent) SetCpu(v string)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *WafRequestDetailsUserAgent) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetRenderingEngine

`func (o *WafRequestDetailsUserAgent) GetRenderingEngine() string`

GetRenderingEngine returns the RenderingEngine field if non-nil, zero value otherwise.

### GetRenderingEngineOk

`func (o *WafRequestDetailsUserAgent) GetRenderingEngineOk() (*string, bool)`

GetRenderingEngineOk returns a tuple with the RenderingEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderingEngine

`func (o *WafRequestDetailsUserAgent) SetRenderingEngine(v string)`

SetRenderingEngine sets RenderingEngine field to given value.

### HasRenderingEngine

`func (o *WafRequestDetailsUserAgent) HasRenderingEngine() bool`

HasRenderingEngine returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


