# CustconfAuthACL

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | This is used by the API to perform conflict checking | [optional] 
**AccessCode** | Pointer to [**AuthACLAccessCodeEnumWrapperValue**](AuthACLAccessCodeEnumWrapperValue.md) |  | [optional] [default to "UNKNOWN"]
**IpList** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. This is a list of IP addresses considered for this policy. | [optional] 
**Protocol** | Pointer to [**CustconfAuthACLProtocolEnumWrapperValue**](custconfAuthACLProtocolEnumWrapperValue.md) |  | [optional] [default to "UNKNOWN"]
**ClientIPSrc** | Pointer to [**AuthACLClientIPSrcEnumWrapperValue**](AuthACLClientIPSrcEnumWrapperValue.md) |  | [optional] [default to "UNKNOWN"]
**Header** | Pointer to **string** | This allows you to specify the name of a HTTP request header which will contain the client IP address to use for this policy. | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewCustconfAuthACL

`func NewCustconfAuthACL() *CustconfAuthACL`

NewCustconfAuthACL instantiates a new CustconfAuthACL object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustconfAuthACLWithDefaults

`func NewCustconfAuthACLWithDefaults() *CustconfAuthACL`

NewCustconfAuthACLWithDefaults instantiates a new CustconfAuthACL object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustconfAuthACL) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustconfAuthACL) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustconfAuthACL) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustconfAuthACL) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccessCode

`func (o *CustconfAuthACL) GetAccessCode() AuthACLAccessCodeEnumWrapperValue`

GetAccessCode returns the AccessCode field if non-nil, zero value otherwise.

### GetAccessCodeOk

`func (o *CustconfAuthACL) GetAccessCodeOk() (*AuthACLAccessCodeEnumWrapperValue, bool)`

GetAccessCodeOk returns a tuple with the AccessCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessCode

`func (o *CustconfAuthACL) SetAccessCode(v AuthACLAccessCodeEnumWrapperValue)`

SetAccessCode sets AccessCode field to given value.

### HasAccessCode

`func (o *CustconfAuthACL) HasAccessCode() bool`

HasAccessCode returns a boolean if a field has been set.

### GetIpList

`func (o *CustconfAuthACL) GetIpList() string`

GetIpList returns the IpList field if non-nil, zero value otherwise.

### GetIpListOk

`func (o *CustconfAuthACL) GetIpListOk() (*string, bool)`

GetIpListOk returns a tuple with the IpList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpList

`func (o *CustconfAuthACL) SetIpList(v string)`

SetIpList sets IpList field to given value.

### HasIpList

`func (o *CustconfAuthACL) HasIpList() bool`

HasIpList returns a boolean if a field has been set.

### GetProtocol

`func (o *CustconfAuthACL) GetProtocol() CustconfAuthACLProtocolEnumWrapperValue`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *CustconfAuthACL) GetProtocolOk() (*CustconfAuthACLProtocolEnumWrapperValue, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *CustconfAuthACL) SetProtocol(v CustconfAuthACLProtocolEnumWrapperValue)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *CustconfAuthACL) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetClientIPSrc

`func (o *CustconfAuthACL) GetClientIPSrc() AuthACLClientIPSrcEnumWrapperValue`

GetClientIPSrc returns the ClientIPSrc field if non-nil, zero value otherwise.

### GetClientIPSrcOk

`func (o *CustconfAuthACL) GetClientIPSrcOk() (*AuthACLClientIPSrcEnumWrapperValue, bool)`

GetClientIPSrcOk returns a tuple with the ClientIPSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIPSrc

`func (o *CustconfAuthACL) SetClientIPSrc(v AuthACLClientIPSrcEnumWrapperValue)`

SetClientIPSrc sets ClientIPSrc field to given value.

### HasClientIPSrc

`func (o *CustconfAuthACL) HasClientIPSrc() bool`

HasClientIPSrc returns a boolean if a field has been set.

### GetHeader

`func (o *CustconfAuthACL) GetHeader() string`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *CustconfAuthACL) GetHeaderOk() (*string, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *CustconfAuthACL) SetHeader(v string)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *CustconfAuthACL) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetEnabled

`func (o *CustconfAuthACL) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustconfAuthACL) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustconfAuthACL) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CustconfAuthACL) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


