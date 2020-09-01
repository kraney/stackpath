# CustconfAccessLogsConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | This is used by the API to perform conflict checking | [optional] 
**ExtraLogFields** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewCustconfAccessLogsConfig

`func NewCustconfAccessLogsConfig() *CustconfAccessLogsConfig`

NewCustconfAccessLogsConfig instantiates a new CustconfAccessLogsConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustconfAccessLogsConfigWithDefaults

`func NewCustconfAccessLogsConfigWithDefaults() *CustconfAccessLogsConfig`

NewCustconfAccessLogsConfigWithDefaults instantiates a new CustconfAccessLogsConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustconfAccessLogsConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustconfAccessLogsConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustconfAccessLogsConfig) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustconfAccessLogsConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExtraLogFields

`func (o *CustconfAccessLogsConfig) GetExtraLogFields() string`

GetExtraLogFields returns the ExtraLogFields field if non-nil, zero value otherwise.

### GetExtraLogFieldsOk

`func (o *CustconfAccessLogsConfig) GetExtraLogFieldsOk() (*string, bool)`

GetExtraLogFieldsOk returns a tuple with the ExtraLogFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraLogFields

`func (o *CustconfAccessLogsConfig) SetExtraLogFields(v string)`

SetExtraLogFields sets ExtraLogFields field to given value.

### HasExtraLogFields

`func (o *CustconfAccessLogsConfig) HasExtraLogFields() bool`

HasExtraLogFields returns a boolean if a field has been set.

### GetEnabled

`func (o *CustconfAccessLogsConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustconfAccessLogsConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustconfAccessLogsConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CustconfAccessLogsConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


