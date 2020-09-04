# CustconfStaticHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | This is used by the API to perform conflict checking | [optional] 
**ClientRequest** | Pointer to **string** | This is the static HTTP header you want inserted into the CDN request. | [optional] 
**Http** | Pointer to **string** | This is the static HTTP header you want inserted into the CDN response. | [optional] 
**OriginPull** | Pointer to **string** | This is the HTTP header you want inserted into the origin pull request. | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**MethodFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**PathFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**HeaderFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 

## Methods

### NewCustconfStaticHeader

`func NewCustconfStaticHeader() *CustconfStaticHeader`

NewCustconfStaticHeader instantiates a new CustconfStaticHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustconfStaticHeaderWithDefaults

`func NewCustconfStaticHeaderWithDefaults() *CustconfStaticHeader`

NewCustconfStaticHeaderWithDefaults instantiates a new CustconfStaticHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustconfStaticHeader) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustconfStaticHeader) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustconfStaticHeader) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustconfStaticHeader) HasId() bool`

HasId returns a boolean if a field has been set.

### GetClientRequest

`func (o *CustconfStaticHeader) GetClientRequest() string`

GetClientRequest returns the ClientRequest field if non-nil, zero value otherwise.

### GetClientRequestOk

`func (o *CustconfStaticHeader) GetClientRequestOk() (*string, bool)`

GetClientRequestOk returns a tuple with the ClientRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientRequest

`func (o *CustconfStaticHeader) SetClientRequest(v string)`

SetClientRequest sets ClientRequest field to given value.

### HasClientRequest

`func (o *CustconfStaticHeader) HasClientRequest() bool`

HasClientRequest returns a boolean if a field has been set.

### GetHttp

`func (o *CustconfStaticHeader) GetHttp() string`

GetHttp returns the Http field if non-nil, zero value otherwise.

### GetHttpOk

`func (o *CustconfStaticHeader) GetHttpOk() (*string, bool)`

GetHttpOk returns a tuple with the Http field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp

`func (o *CustconfStaticHeader) SetHttp(v string)`

SetHttp sets Http field to given value.

### HasHttp

`func (o *CustconfStaticHeader) HasHttp() bool`

HasHttp returns a boolean if a field has been set.

### GetOriginPull

`func (o *CustconfStaticHeader) GetOriginPull() string`

GetOriginPull returns the OriginPull field if non-nil, zero value otherwise.

### GetOriginPullOk

`func (o *CustconfStaticHeader) GetOriginPullOk() (*string, bool)`

GetOriginPullOk returns a tuple with the OriginPull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginPull

`func (o *CustconfStaticHeader) SetOriginPull(v string)`

SetOriginPull sets OriginPull field to given value.

### HasOriginPull

`func (o *CustconfStaticHeader) HasOriginPull() bool`

HasOriginPull returns a boolean if a field has been set.

### GetEnabled

`func (o *CustconfStaticHeader) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustconfStaticHeader) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustconfStaticHeader) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CustconfStaticHeader) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMethodFilter

`func (o *CustconfStaticHeader) GetMethodFilter() string`

GetMethodFilter returns the MethodFilter field if non-nil, zero value otherwise.

### GetMethodFilterOk

`func (o *CustconfStaticHeader) GetMethodFilterOk() (*string, bool)`

GetMethodFilterOk returns a tuple with the MethodFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodFilter

`func (o *CustconfStaticHeader) SetMethodFilter(v string)`

SetMethodFilter sets MethodFilter field to given value.

### HasMethodFilter

`func (o *CustconfStaticHeader) HasMethodFilter() bool`

HasMethodFilter returns a boolean if a field has been set.

### GetPathFilter

`func (o *CustconfStaticHeader) GetPathFilter() string`

GetPathFilter returns the PathFilter field if non-nil, zero value otherwise.

### GetPathFilterOk

`func (o *CustconfStaticHeader) GetPathFilterOk() (*string, bool)`

GetPathFilterOk returns a tuple with the PathFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathFilter

`func (o *CustconfStaticHeader) SetPathFilter(v string)`

SetPathFilter sets PathFilter field to given value.

### HasPathFilter

`func (o *CustconfStaticHeader) HasPathFilter() bool`

HasPathFilter returns a boolean if a field has been set.

### GetHeaderFilter

`func (o *CustconfStaticHeader) GetHeaderFilter() string`

GetHeaderFilter returns the HeaderFilter field if non-nil, zero value otherwise.

### GetHeaderFilterOk

`func (o *CustconfStaticHeader) GetHeaderFilterOk() (*string, bool)`

GetHeaderFilterOk returns a tuple with the HeaderFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderFilter

`func (o *CustconfStaticHeader) SetHeaderFilter(v string)`

SetHeaderFilter sets HeaderFilter field to given value.

### HasHeaderFilter

`func (o *CustconfStaticHeader) HasHeaderFilter() bool`

HasHeaderFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


