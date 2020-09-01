# CustconfCompression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | This is used by the API to perform conflict checking | [optional] 
**Gzip** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. The list of file extensions you want the caching servers to use to identify the content you want compressed before delivering it to end users | [optional] 
**Mime** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. A list of rules based on MIME types you want the caching servers to use to identify content you want compressed before delivering it to end users. | [optional] 
**Level** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewCustconfCompression

`func NewCustconfCompression() *CustconfCompression`

NewCustconfCompression instantiates a new CustconfCompression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustconfCompressionWithDefaults

`func NewCustconfCompressionWithDefaults() *CustconfCompression`

NewCustconfCompressionWithDefaults instantiates a new CustconfCompression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustconfCompression) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustconfCompression) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustconfCompression) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustconfCompression) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGzip

`func (o *CustconfCompression) GetGzip() string`

GetGzip returns the Gzip field if non-nil, zero value otherwise.

### GetGzipOk

`func (o *CustconfCompression) GetGzipOk() (*string, bool)`

GetGzipOk returns a tuple with the Gzip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGzip

`func (o *CustconfCompression) SetGzip(v string)`

SetGzip sets Gzip field to given value.

### HasGzip

`func (o *CustconfCompression) HasGzip() bool`

HasGzip returns a boolean if a field has been set.

### GetMime

`func (o *CustconfCompression) GetMime() string`

GetMime returns the Mime field if non-nil, zero value otherwise.

### GetMimeOk

`func (o *CustconfCompression) GetMimeOk() (*string, bool)`

GetMimeOk returns a tuple with the Mime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMime

`func (o *CustconfCompression) SetMime(v string)`

SetMime sets Mime field to given value.

### HasMime

`func (o *CustconfCompression) HasMime() bool`

HasMime returns a boolean if a field has been set.

### GetLevel

`func (o *CustconfCompression) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *CustconfCompression) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *CustconfCompression) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *CustconfCompression) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetEnabled

`func (o *CustconfCompression) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustconfCompression) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustconfCompression) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CustconfCompression) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


