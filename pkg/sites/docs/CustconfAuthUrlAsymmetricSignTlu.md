# CustconfAuthUrlAsymmetricSignTlu

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | This is used by the API to perform conflict checking | [optional] 
**ExpireParameterName** | Pointer to **string** |  | [optional] 
**KeyIdParameterName** | Pointer to **string** |  | [optional] 
**AlgorithmIdParameterName** | Pointer to **string** |  | [optional] 
**DigestParameterName** | Pointer to **string** |  | [optional] 
**PublicKeyIdMap** | Pointer to **map[string]string** |  | [optional] 
**AlgorithmIdMap** | Pointer to [**map[string]CustconfAuthUrlAsymmetricSignTluAlgorithmIdMapEnumWrapperValue**](custconfAuthUrlAsymmetricSignTluAlgorithmIdMapEnumWrapperValue.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**MethodFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**PathFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**HeaderFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 

## Methods

### NewCustconfAuthUrlAsymmetricSignTlu

`func NewCustconfAuthUrlAsymmetricSignTlu() *CustconfAuthUrlAsymmetricSignTlu`

NewCustconfAuthUrlAsymmetricSignTlu instantiates a new CustconfAuthUrlAsymmetricSignTlu object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustconfAuthUrlAsymmetricSignTluWithDefaults

`func NewCustconfAuthUrlAsymmetricSignTluWithDefaults() *CustconfAuthUrlAsymmetricSignTlu`

NewCustconfAuthUrlAsymmetricSignTluWithDefaults instantiates a new CustconfAuthUrlAsymmetricSignTlu object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustconfAuthUrlAsymmetricSignTlu) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustconfAuthUrlAsymmetricSignTlu) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustconfAuthUrlAsymmetricSignTlu) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustconfAuthUrlAsymmetricSignTlu) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExpireParameterName

`func (o *CustconfAuthUrlAsymmetricSignTlu) GetExpireParameterName() string`

GetExpireParameterName returns the ExpireParameterName field if non-nil, zero value otherwise.

### GetExpireParameterNameOk

`func (o *CustconfAuthUrlAsymmetricSignTlu) GetExpireParameterNameOk() (*string, bool)`

GetExpireParameterNameOk returns a tuple with the ExpireParameterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireParameterName

`func (o *CustconfAuthUrlAsymmetricSignTlu) SetExpireParameterName(v string)`

SetExpireParameterName sets ExpireParameterName field to given value.

### HasExpireParameterName

`func (o *CustconfAuthUrlAsymmetricSignTlu) HasExpireParameterName() bool`

HasExpireParameterName returns a boolean if a field has been set.

### GetKeyIdParameterName

`func (o *CustconfAuthUrlAsymmetricSignTlu) GetKeyIdParameterName() string`

GetKeyIdParameterName returns the KeyIdParameterName field if non-nil, zero value otherwise.

### GetKeyIdParameterNameOk

`func (o *CustconfAuthUrlAsymmetricSignTlu) GetKeyIdParameterNameOk() (*string, bool)`

GetKeyIdParameterNameOk returns a tuple with the KeyIdParameterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyIdParameterName

`func (o *CustconfAuthUrlAsymmetricSignTlu) SetKeyIdParameterName(v string)`

SetKeyIdParameterName sets KeyIdParameterName field to given value.

### HasKeyIdParameterName

`func (o *CustconfAuthUrlAsymmetricSignTlu) HasKeyIdParameterName() bool`

HasKeyIdParameterName returns a boolean if a field has been set.

### GetAlgorithmIdParameterName

`func (o *CustconfAuthUrlAsymmetricSignTlu) GetAlgorithmIdParameterName() string`

GetAlgorithmIdParameterName returns the AlgorithmIdParameterName field if non-nil, zero value otherwise.

### GetAlgorithmIdParameterNameOk

`func (o *CustconfAuthUrlAsymmetricSignTlu) GetAlgorithmIdParameterNameOk() (*string, bool)`

GetAlgorithmIdParameterNameOk returns a tuple with the AlgorithmIdParameterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithmIdParameterName

`func (o *CustconfAuthUrlAsymmetricSignTlu) SetAlgorithmIdParameterName(v string)`

SetAlgorithmIdParameterName sets AlgorithmIdParameterName field to given value.

### HasAlgorithmIdParameterName

`func (o *CustconfAuthUrlAsymmetricSignTlu) HasAlgorithmIdParameterName() bool`

HasAlgorithmIdParameterName returns a boolean if a field has been set.

### GetDigestParameterName

`func (o *CustconfAuthUrlAsymmetricSignTlu) GetDigestParameterName() string`

GetDigestParameterName returns the DigestParameterName field if non-nil, zero value otherwise.

### GetDigestParameterNameOk

`func (o *CustconfAuthUrlAsymmetricSignTlu) GetDigestParameterNameOk() (*string, bool)`

GetDigestParameterNameOk returns a tuple with the DigestParameterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestParameterName

`func (o *CustconfAuthUrlAsymmetricSignTlu) SetDigestParameterName(v string)`

SetDigestParameterName sets DigestParameterName field to given value.

### HasDigestParameterName

`func (o *CustconfAuthUrlAsymmetricSignTlu) HasDigestParameterName() bool`

HasDigestParameterName returns a boolean if a field has been set.

### GetPublicKeyIdMap

`func (o *CustconfAuthUrlAsymmetricSignTlu) GetPublicKeyIdMap() map[string]string`

GetPublicKeyIdMap returns the PublicKeyIdMap field if non-nil, zero value otherwise.

### GetPublicKeyIdMapOk

`func (o *CustconfAuthUrlAsymmetricSignTlu) GetPublicKeyIdMapOk() (*map[string]string, bool)`

GetPublicKeyIdMapOk returns a tuple with the PublicKeyIdMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyIdMap

`func (o *CustconfAuthUrlAsymmetricSignTlu) SetPublicKeyIdMap(v map[string]string)`

SetPublicKeyIdMap sets PublicKeyIdMap field to given value.

### HasPublicKeyIdMap

`func (o *CustconfAuthUrlAsymmetricSignTlu) HasPublicKeyIdMap() bool`

HasPublicKeyIdMap returns a boolean if a field has been set.

### GetAlgorithmIdMap

`func (o *CustconfAuthUrlAsymmetricSignTlu) GetAlgorithmIdMap() map[string]CustconfAuthUrlAsymmetricSignTluAlgorithmIdMapEnumWrapperValue`

GetAlgorithmIdMap returns the AlgorithmIdMap field if non-nil, zero value otherwise.

### GetAlgorithmIdMapOk

`func (o *CustconfAuthUrlAsymmetricSignTlu) GetAlgorithmIdMapOk() (*map[string]CustconfAuthUrlAsymmetricSignTluAlgorithmIdMapEnumWrapperValue, bool)`

GetAlgorithmIdMapOk returns a tuple with the AlgorithmIdMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithmIdMap

`func (o *CustconfAuthUrlAsymmetricSignTlu) SetAlgorithmIdMap(v map[string]CustconfAuthUrlAsymmetricSignTluAlgorithmIdMapEnumWrapperValue)`

SetAlgorithmIdMap sets AlgorithmIdMap field to given value.

### HasAlgorithmIdMap

`func (o *CustconfAuthUrlAsymmetricSignTlu) HasAlgorithmIdMap() bool`

HasAlgorithmIdMap returns a boolean if a field has been set.

### GetEnabled

`func (o *CustconfAuthUrlAsymmetricSignTlu) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustconfAuthUrlAsymmetricSignTlu) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustconfAuthUrlAsymmetricSignTlu) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CustconfAuthUrlAsymmetricSignTlu) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMethodFilter

`func (o *CustconfAuthUrlAsymmetricSignTlu) GetMethodFilter() string`

GetMethodFilter returns the MethodFilter field if non-nil, zero value otherwise.

### GetMethodFilterOk

`func (o *CustconfAuthUrlAsymmetricSignTlu) GetMethodFilterOk() (*string, bool)`

GetMethodFilterOk returns a tuple with the MethodFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodFilter

`func (o *CustconfAuthUrlAsymmetricSignTlu) SetMethodFilter(v string)`

SetMethodFilter sets MethodFilter field to given value.

### HasMethodFilter

`func (o *CustconfAuthUrlAsymmetricSignTlu) HasMethodFilter() bool`

HasMethodFilter returns a boolean if a field has been set.

### GetPathFilter

`func (o *CustconfAuthUrlAsymmetricSignTlu) GetPathFilter() string`

GetPathFilter returns the PathFilter field if non-nil, zero value otherwise.

### GetPathFilterOk

`func (o *CustconfAuthUrlAsymmetricSignTlu) GetPathFilterOk() (*string, bool)`

GetPathFilterOk returns a tuple with the PathFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathFilter

`func (o *CustconfAuthUrlAsymmetricSignTlu) SetPathFilter(v string)`

SetPathFilter sets PathFilter field to given value.

### HasPathFilter

`func (o *CustconfAuthUrlAsymmetricSignTlu) HasPathFilter() bool`

HasPathFilter returns a boolean if a field has been set.

### GetHeaderFilter

`func (o *CustconfAuthUrlAsymmetricSignTlu) GetHeaderFilter() string`

GetHeaderFilter returns the HeaderFilter field if non-nil, zero value otherwise.

### GetHeaderFilterOk

`func (o *CustconfAuthUrlAsymmetricSignTlu) GetHeaderFilterOk() (*string, bool)`

GetHeaderFilterOk returns a tuple with the HeaderFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderFilter

`func (o *CustconfAuthUrlAsymmetricSignTlu) SetHeaderFilter(v string)`

SetHeaderFilter sets HeaderFilter field to given value.

### HasHeaderFilter

`func (o *CustconfAuthUrlAsymmetricSignTlu) HasHeaderFilter() bool`

HasHeaderFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


