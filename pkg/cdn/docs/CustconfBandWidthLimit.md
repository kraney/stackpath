# CustconfBandWidthLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | This is used by the API to perform conflict checking | [optional] 
**Rule** | Pointer to **string** | String of values delimited by a &#39;|&#39; character. These are pattern match rules to use for applying rate limiting on requests. | [optional] 
**Values** | Pointer to **string** | These are the initial bytes (ri) and the sustained rate (rs) query string parameters to use for this rule. Example: ri&#x3D;100,rs&#x3D;1000 | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewCustconfBandWidthLimit

`func NewCustconfBandWidthLimit() *CustconfBandWidthLimit`

NewCustconfBandWidthLimit instantiates a new CustconfBandWidthLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustconfBandWidthLimitWithDefaults

`func NewCustconfBandWidthLimitWithDefaults() *CustconfBandWidthLimit`

NewCustconfBandWidthLimitWithDefaults instantiates a new CustconfBandWidthLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustconfBandWidthLimit) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustconfBandWidthLimit) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustconfBandWidthLimit) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustconfBandWidthLimit) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRule

`func (o *CustconfBandWidthLimit) GetRule() string`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *CustconfBandWidthLimit) GetRuleOk() (*string, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *CustconfBandWidthLimit) SetRule(v string)`

SetRule sets Rule field to given value.

### HasRule

`func (o *CustconfBandWidthLimit) HasRule() bool`

HasRule returns a boolean if a field has been set.

### GetValues

`func (o *CustconfBandWidthLimit) GetValues() string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *CustconfBandWidthLimit) GetValuesOk() (*string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *CustconfBandWidthLimit) SetValues(v string)`

SetValues sets Values field to given value.

### HasValues

`func (o *CustconfBandWidthLimit) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetEnabled

`func (o *CustconfBandWidthLimit) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustconfBandWidthLimit) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustconfBandWidthLimit) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CustconfBandWidthLimit) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


