# CdnGetSiteDnsTargetsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | Pointer to **[]string** | The requested DNS CNAME targets  A site&#39;s hostname should point to these CNAME targets in order for traffic to be sent through StackPath&#39;s edge nodes. | [optional] 

## Methods

### NewCdnGetSiteDnsTargetsResponse

`func NewCdnGetSiteDnsTargetsResponse() *CdnGetSiteDnsTargetsResponse`

NewCdnGetSiteDnsTargetsResponse instantiates a new CdnGetSiteDnsTargetsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdnGetSiteDnsTargetsResponseWithDefaults

`func NewCdnGetSiteDnsTargetsResponseWithDefaults() *CdnGetSiteDnsTargetsResponse`

NewCdnGetSiteDnsTargetsResponseWithDefaults instantiates a new CdnGetSiteDnsTargetsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *CdnGetSiteDnsTargetsResponse) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *CdnGetSiteDnsTargetsResponse) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *CdnGetSiteDnsTargetsResponse) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *CdnGetSiteDnsTargetsResponse) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


