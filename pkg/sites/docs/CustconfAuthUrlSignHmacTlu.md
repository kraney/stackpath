# CustconfAuthUrlSignHmacTlu

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | This is used by the API to perform conflict checking | [optional] 
**ExpireParameterName** | Pointer to **string** |  | [optional] 
**KeyIdParameterName** | Pointer to **string** |  | [optional] 
**AlgorithmIdParameterName** | Pointer to **string** |  | [optional] 
**DigestParameterName** | Pointer to **string** |  | [optional] 
**SymmetricKeyIdMap** | Pointer to **map[string]string** |  | [optional] 
**AlgorithmIdMap** | Pointer to [**map[string]CustconfAuthUrlSignHmacTluAlgorithmIdMapEnumWrapperValue**](custconfAuthUrlSignHmacTluAlgorithmIdMapEnumWrapperValue.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**MethodFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**PathFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**HeaderFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 

## Methods

### NewCustconfAuthUrlSignHmacTlu

`func NewCustconfAuthUrlSignHmacTlu() *CustconfAuthUrlSignHmacTlu`

NewCustconfAuthUrlSignHmacTlu instantiates a new CustconfAuthUrlSignHmacTlu object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustconfAuthUrlSignHmacTluWithDefaults

`func NewCustconfAuthUrlSignHmacTluWithDefaults() *CustconfAuthUrlSignHmacTlu`

NewCustconfAuthUrlSignHmacTluWithDefaults instantiates a new CustconfAuthUrlSignHmacTlu object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustconfAuthUrlSignHmacTlu) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustconfAuthUrlSignHmacTlu) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustconfAuthUrlSignHmacTlu) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustconfAuthUrlSignHmacTlu) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExpireParameterName

`func (o *CustconfAuthUrlSignHmacTlu) GetExpireParameterName() string`

GetExpireParameterName returns the ExpireParameterName field if non-nil, zero value otherwise.

### GetExpireParameterNameOk

`func (o *CustconfAuthUrlSignHmacTlu) GetExpireParameterNameOk() (*string, bool)`

GetExpireParameterNameOk returns a tuple with the ExpireParameterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireParameterName

`func (o *CustconfAuthUrlSignHmacTlu) SetExpireParameterName(v string)`

SetExpireParameterName sets ExpireParameterName field to given value.

### HasExpireParameterName

`func (o *CustconfAuthUrlSignHmacTlu) HasExpireParameterName() bool`

HasExpireParameterName returns a boolean if a field has been set.

### GetKeyIdParameterName

`func (o *CustconfAuthUrlSignHmacTlu) GetKeyIdParameterName() string`

GetKeyIdParameterName returns the KeyIdParameterName field if non-nil, zero value otherwise.

### GetKeyIdParameterNameOk

`func (o *CustconfAuthUrlSignHmacTlu) GetKeyIdParameterNameOk() (*string, bool)`

GetKeyIdParameterNameOk returns a tuple with the KeyIdParameterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyIdParameterName

`func (o *CustconfAuthUrlSignHmacTlu) SetKeyIdParameterName(v string)`

SetKeyIdParameterName sets KeyIdParameterName field to given value.

### HasKeyIdParameterName

`func (o *CustconfAuthUrlSignHmacTlu) HasKeyIdParameterName() bool`

HasKeyIdParameterName returns a boolean if a field has been set.

### GetAlgorithmIdParameterName

`func (o *CustconfAuthUrlSignHmacTlu) GetAlgorithmIdParameterName() string`

GetAlgorithmIdParameterName returns the AlgorithmIdParameterName field if non-nil, zero value otherwise.

### GetAlgorithmIdParameterNameOk

`func (o *CustconfAuthUrlSignHmacTlu) GetAlgorithmIdParameterNameOk() (*string, bool)`

GetAlgorithmIdParameterNameOk returns a tuple with the AlgorithmIdParameterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithmIdParameterName

`func (o *CustconfAuthUrlSignHmacTlu) SetAlgorithmIdParameterName(v string)`

SetAlgorithmIdParameterName sets AlgorithmIdParameterName field to given value.

### HasAlgorithmIdParameterName

`func (o *CustconfAuthUrlSignHmacTlu) HasAlgorithmIdParameterName() bool`

HasAlgorithmIdParameterName returns a boolean if a field has been set.

### GetDigestParameterName

`func (o *CustconfAuthUrlSignHmacTlu) GetDigestParameterName() string`

GetDigestParameterName returns the DigestParameterName field if non-nil, zero value otherwise.

### GetDigestParameterNameOk

`func (o *CustconfAuthUrlSignHmacTlu) GetDigestParameterNameOk() (*string, bool)`

GetDigestParameterNameOk returns a tuple with the DigestParameterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestParameterName

`func (o *CustconfAuthUrlSignHmacTlu) SetDigestParameterName(v string)`

SetDigestParameterName sets DigestParameterName field to given value.

### HasDigestParameterName

`func (o *CustconfAuthUrlSignHmacTlu) HasDigestParameterName() bool`

HasDigestParameterName returns a boolean if a field has been set.

### GetSymmetricKeyIdMap

`func (o *CustconfAuthUrlSignHmacTlu) GetSymmetricKeyIdMap() map[string]string`

GetSymmetricKeyIdMap returns the SymmetricKeyIdMap field if non-nil, zero value otherwise.

### GetSymmetricKeyIdMapOk

`func (o *CustconfAuthUrlSignHmacTlu) GetSymmetricKeyIdMapOk() (*map[string]string, bool)`

GetSymmetricKeyIdMapOk returns a tuple with the SymmetricKeyIdMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymmetricKeyIdMap

`func (o *CustconfAuthUrlSignHmacTlu) SetSymmetricKeyIdMap(v map[string]string)`

SetSymmetricKeyIdMap sets SymmetricKeyIdMap field to given value.

### HasSymmetricKeyIdMap

`func (o *CustconfAuthUrlSignHmacTlu) HasSymmetricKeyIdMap() bool`

HasSymmetricKeyIdMap returns a boolean if a field has been set.

### GetAlgorithmIdMap

`func (o *CustconfAuthUrlSignHmacTlu) GetAlgorithmIdMap() map[string]CustconfAuthUrlSignHmacTluAlgorithmIdMapEnumWrapperValue`

GetAlgorithmIdMap returns the AlgorithmIdMap field if non-nil, zero value otherwise.

### GetAlgorithmIdMapOk

`func (o *CustconfAuthUrlSignHmacTlu) GetAlgorithmIdMapOk() (*map[string]CustconfAuthUrlSignHmacTluAlgorithmIdMapEnumWrapperValue, bool)`

GetAlgorithmIdMapOk returns a tuple with the AlgorithmIdMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithmIdMap

`func (o *CustconfAuthUrlSignHmacTlu) SetAlgorithmIdMap(v map[string]CustconfAuthUrlSignHmacTluAlgorithmIdMapEnumWrapperValue)`

SetAlgorithmIdMap sets AlgorithmIdMap field to given value.

### HasAlgorithmIdMap

`func (o *CustconfAuthUrlSignHmacTlu) HasAlgorithmIdMap() bool`

HasAlgorithmIdMap returns a boolean if a field has been set.

### GetEnabled

`func (o *CustconfAuthUrlSignHmacTlu) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustconfAuthUrlSignHmacTlu) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustconfAuthUrlSignHmacTlu) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CustconfAuthUrlSignHmacTlu) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMethodFilter

`func (o *CustconfAuthUrlSignHmacTlu) GetMethodFilter() string`

GetMethodFilter returns the MethodFilter field if non-nil, zero value otherwise.

### GetMethodFilterOk

`func (o *CustconfAuthUrlSignHmacTlu) GetMethodFilterOk() (*string, bool)`

GetMethodFilterOk returns a tuple with the MethodFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodFilter

`func (o *CustconfAuthUrlSignHmacTlu) SetMethodFilter(v string)`

SetMethodFilter sets MethodFilter field to given value.

### HasMethodFilter

`func (o *CustconfAuthUrlSignHmacTlu) HasMethodFilter() bool`

HasMethodFilter returns a boolean if a field has been set.

### GetPathFilter

`func (o *CustconfAuthUrlSignHmacTlu) GetPathFilter() string`

GetPathFilter returns the PathFilter field if non-nil, zero value otherwise.

### GetPathFilterOk

`func (o *CustconfAuthUrlSignHmacTlu) GetPathFilterOk() (*string, bool)`

GetPathFilterOk returns a tuple with the PathFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathFilter

`func (o *CustconfAuthUrlSignHmacTlu) SetPathFilter(v string)`

SetPathFilter sets PathFilter field to given value.

### HasPathFilter

`func (o *CustconfAuthUrlSignHmacTlu) HasPathFilter() bool`

HasPathFilter returns a boolean if a field has been set.

### GetHeaderFilter

`func (o *CustconfAuthUrlSignHmacTlu) GetHeaderFilter() string`

GetHeaderFilter returns the HeaderFilter field if non-nil, zero value otherwise.

### GetHeaderFilterOk

`func (o *CustconfAuthUrlSignHmacTlu) GetHeaderFilterOk() (*string, bool)`

GetHeaderFilterOk returns a tuple with the HeaderFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderFilter

`func (o *CustconfAuthUrlSignHmacTlu) SetHeaderFilter(v string)`

SetHeaderFilter sets HeaderFilter field to given value.

### HasHeaderFilter

`func (o *CustconfAuthUrlSignHmacTlu) HasHeaderFilter() bool`

HasHeaderFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


