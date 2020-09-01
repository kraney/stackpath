# CustconfCacheKeyModification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | This is used by the API to perform conflict checking | [optional] 
**NormalizeKeyPathToLowerCase** | Pointer to **bool** | When set, purges and requests for a file will be case insensitive. This setting is useful if you have a case insensitive origin server and would like to avoid duplicating assets. | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewCustconfCacheKeyModification

`func NewCustconfCacheKeyModification() *CustconfCacheKeyModification`

NewCustconfCacheKeyModification instantiates a new CustconfCacheKeyModification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustconfCacheKeyModificationWithDefaults

`func NewCustconfCacheKeyModificationWithDefaults() *CustconfCacheKeyModification`

NewCustconfCacheKeyModificationWithDefaults instantiates a new CustconfCacheKeyModification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustconfCacheKeyModification) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustconfCacheKeyModification) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustconfCacheKeyModification) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustconfCacheKeyModification) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNormalizeKeyPathToLowerCase

`func (o *CustconfCacheKeyModification) GetNormalizeKeyPathToLowerCase() bool`

GetNormalizeKeyPathToLowerCase returns the NormalizeKeyPathToLowerCase field if non-nil, zero value otherwise.

### GetNormalizeKeyPathToLowerCaseOk

`func (o *CustconfCacheKeyModification) GetNormalizeKeyPathToLowerCaseOk() (*bool, bool)`

GetNormalizeKeyPathToLowerCaseOk returns a tuple with the NormalizeKeyPathToLowerCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizeKeyPathToLowerCase

`func (o *CustconfCacheKeyModification) SetNormalizeKeyPathToLowerCase(v bool)`

SetNormalizeKeyPathToLowerCase sets NormalizeKeyPathToLowerCase field to given value.

### HasNormalizeKeyPathToLowerCase

`func (o *CustconfCacheKeyModification) HasNormalizeKeyPathToLowerCase() bool`

HasNormalizeKeyPathToLowerCase returns a boolean if a field has been set.

### GetEnabled

`func (o *CustconfCacheKeyModification) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustconfCacheKeyModification) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustconfCacheKeyModification) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CustconfCacheKeyModification) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


