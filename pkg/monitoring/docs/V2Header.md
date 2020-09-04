# V2Header

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Header** | Pointer to **string** | An HTTP header&#39;s field name | [optional] 
**Value** | Pointer to **string** | An HTTP header&#39;s value | [optional] 

## Methods

### NewV2Header

`func NewV2Header() *V2Header`

NewV2Header instantiates a new V2Header object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2HeaderWithDefaults

`func NewV2HeaderWithDefaults() *V2Header`

NewV2HeaderWithDefaults instantiates a new V2Header object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeader

`func (o *V2Header) GetHeader() string`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *V2Header) GetHeaderOk() (*string, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *V2Header) SetHeader(v string)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *V2Header) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetValue

`func (o *V2Header) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *V2Header) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *V2Header) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *V2Header) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


