# V1IpBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cidr** | Pointer to **string** | A subnet that will define all the IPs allowed by a rule | [optional] 
**Except** | Pointer to **[]string** | A list of subnets that will be excluded from the above subnet  This allows a convenient way to allow multiple ip ranges in a single expression | [optional] 

## Methods

### NewV1IpBlock

`func NewV1IpBlock() *V1IpBlock`

NewV1IpBlock instantiates a new V1IpBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1IpBlockWithDefaults

`func NewV1IpBlockWithDefaults() *V1IpBlock`

NewV1IpBlockWithDefaults instantiates a new V1IpBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCidr

`func (o *V1IpBlock) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *V1IpBlock) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *V1IpBlock) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *V1IpBlock) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetExcept

`func (o *V1IpBlock) GetExcept() []string`

GetExcept returns the Except field if non-nil, zero value otherwise.

### GetExceptOk

`func (o *V1IpBlock) GetExceptOk() (*[]string, bool)`

GetExceptOk returns a tuple with the Except field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcept

`func (o *V1IpBlock) SetExcept(v []string)`

SetExcept sets Except field to given value.

### HasExcept

`func (o *V1IpBlock) HasExcept() bool`

HasExcept returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


