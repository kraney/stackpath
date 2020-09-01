# CustconfResponseHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | This is used by the API to perform conflict checking | [optional] 
**Http** | Pointer to **string** | A pipe delimited list of rules that instructs the CDN caching servers to include a content-disposition header. The rules included in this setting must be entered in the following format: Content-Disposition:&lt;User Agent&gt;:&lt;file extension 1&gt;, &lt;file extension 2&gt;. For example, to send the Content-Disposition header for all Mozilla browsers on the delivery of mp3, exe, tar, zip, gz and rar files, you would enter the following in the field: Content-Disposition:Mozilla*:mp3,exe,tar,zip,gz,rar | [optional] 
**EnableETag** | Pointer to **bool** | This gives the ability to disable the ETag header on the response. | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewCustconfResponseHeader

`func NewCustconfResponseHeader() *CustconfResponseHeader`

NewCustconfResponseHeader instantiates a new CustconfResponseHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustconfResponseHeaderWithDefaults

`func NewCustconfResponseHeaderWithDefaults() *CustconfResponseHeader`

NewCustconfResponseHeaderWithDefaults instantiates a new CustconfResponseHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustconfResponseHeader) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustconfResponseHeader) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustconfResponseHeader) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustconfResponseHeader) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHttp

`func (o *CustconfResponseHeader) GetHttp() string`

GetHttp returns the Http field if non-nil, zero value otherwise.

### GetHttpOk

`func (o *CustconfResponseHeader) GetHttpOk() (*string, bool)`

GetHttpOk returns a tuple with the Http field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp

`func (o *CustconfResponseHeader) SetHttp(v string)`

SetHttp sets Http field to given value.

### HasHttp

`func (o *CustconfResponseHeader) HasHttp() bool`

HasHttp returns a boolean if a field has been set.

### GetEnableETag

`func (o *CustconfResponseHeader) GetEnableETag() bool`

GetEnableETag returns the EnableETag field if non-nil, zero value otherwise.

### GetEnableETagOk

`func (o *CustconfResponseHeader) GetEnableETagOk() (*bool, bool)`

GetEnableETagOk returns a tuple with the EnableETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableETag

`func (o *CustconfResponseHeader) SetEnableETag(v bool)`

SetEnableETag sets EnableETag field to given value.

### HasEnableETag

`func (o *CustconfResponseHeader) HasEnableETag() bool`

HasEnableETag returns a boolean if a field has been set.

### GetEnabled

`func (o *CustconfResponseHeader) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustconfResponseHeader) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustconfResponseHeader) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CustconfResponseHeader) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


