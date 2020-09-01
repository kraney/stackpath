# NetworkOrganization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of organization | [optional] 
**Subnet** | Pointer to **string** | The IP bock of the organization | [optional] 

## Methods

### NewNetworkOrganization

`func NewNetworkOrganization() *NetworkOrganization`

NewNetworkOrganization instantiates a new NetworkOrganization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkOrganizationWithDefaults

`func NewNetworkOrganizationWithDefaults() *NetworkOrganization`

NewNetworkOrganizationWithDefaults instantiates a new NetworkOrganization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NetworkOrganization) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkOrganization) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkOrganization) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkOrganization) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubnet

`func (o *NetworkOrganization) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *NetworkOrganization) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *NetworkOrganization) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *NetworkOrganization) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


