# CustconfOriginPullProtocol

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | This is used by the API to perform conflict checking | [optional] 
**Protocol** | Pointer to [**CustconfOriginPullProtocolProtocolEnumWrapperValue**](custconfOriginPullProtocolProtocolEnumWrapperValue.md) |  | [optional] [default to "UNKNOWN"]
**EnableSNI** | Pointer to **bool** | This key allows you to configure the CDN caching servers to use SNI while making Secured Connection to Origin. | [optional] 

## Methods

### NewCustconfOriginPullProtocol

`func NewCustconfOriginPullProtocol() *CustconfOriginPullProtocol`

NewCustconfOriginPullProtocol instantiates a new CustconfOriginPullProtocol object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustconfOriginPullProtocolWithDefaults

`func NewCustconfOriginPullProtocolWithDefaults() *CustconfOriginPullProtocol`

NewCustconfOriginPullProtocolWithDefaults instantiates a new CustconfOriginPullProtocol object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustconfOriginPullProtocol) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustconfOriginPullProtocol) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustconfOriginPullProtocol) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustconfOriginPullProtocol) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProtocol

`func (o *CustconfOriginPullProtocol) GetProtocol() CustconfOriginPullProtocolProtocolEnumWrapperValue`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *CustconfOriginPullProtocol) GetProtocolOk() (*CustconfOriginPullProtocolProtocolEnumWrapperValue, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *CustconfOriginPullProtocol) SetProtocol(v CustconfOriginPullProtocolProtocolEnumWrapperValue)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *CustconfOriginPullProtocol) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetEnableSNI

`func (o *CustconfOriginPullProtocol) GetEnableSNI() bool`

GetEnableSNI returns the EnableSNI field if non-nil, zero value otherwise.

### GetEnableSNIOk

`func (o *CustconfOriginPullProtocol) GetEnableSNIOk() (*bool, bool)`

GetEnableSNIOk returns a tuple with the EnableSNI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSNI

`func (o *CustconfOriginPullProtocol) SetEnableSNI(v bool)`

SetEnableSNI sets EnableSNI field to given value.

### HasEnableSNI

`func (o *CustconfOriginPullProtocol) HasEnableSNI() bool`

HasEnableSNI returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


