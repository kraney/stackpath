# CustconfVHost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | This is used by the API to perform conflict checking | [optional] 
**Domain** | Pointer to **string** | This is the hostname you want to enable in this policy. Note: You must configure your container&#39;s CNAME record with your DNS provider to enable this hostname to deliver content. | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewCustconfVHost

`func NewCustconfVHost() *CustconfVHost`

NewCustconfVHost instantiates a new CustconfVHost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustconfVHostWithDefaults

`func NewCustconfVHostWithDefaults() *CustconfVHost`

NewCustconfVHostWithDefaults instantiates a new CustconfVHost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustconfVHost) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustconfVHost) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustconfVHost) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustconfVHost) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDomain

`func (o *CustconfVHost) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *CustconfVHost) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *CustconfVHost) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *CustconfVHost) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetEnabled

`func (o *CustconfVHost) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustconfVHost) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustconfVHost) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CustconfVHost) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


