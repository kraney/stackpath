# CustconfBandwidthRateLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | This is used by the API to perform conflict checking | [optional] 
**InitialBurstName** | Pointer to **string** |  | [optional] 
**SustainedRateName** | Pointer to **string** |  | [optional] 
**InitialBurstUnits** | Pointer to [**BandwidthRateLimitInitialBurstUnitsEnumWrapperValue**](BandwidthRateLimitInitialBurstUnitsEnumWrapperValue.md) |  | [optional] [default to "UNKNOWN"]
**SustainedRateUnits** | Pointer to [**BandwidthRateLimitSustainedRateUnitsEnumWrapperValue**](BandwidthRateLimitSustainedRateUnitsEnumWrapperValue.md) |  | [optional] [default to "UNKNOWN"]
**Enabled** | Pointer to **bool** |  | [optional] 
**MethodFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**PathFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**HeaderFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 

## Methods

### NewCustconfBandwidthRateLimit

`func NewCustconfBandwidthRateLimit() *CustconfBandwidthRateLimit`

NewCustconfBandwidthRateLimit instantiates a new CustconfBandwidthRateLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustconfBandwidthRateLimitWithDefaults

`func NewCustconfBandwidthRateLimitWithDefaults() *CustconfBandwidthRateLimit`

NewCustconfBandwidthRateLimitWithDefaults instantiates a new CustconfBandwidthRateLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustconfBandwidthRateLimit) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustconfBandwidthRateLimit) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustconfBandwidthRateLimit) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustconfBandwidthRateLimit) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInitialBurstName

`func (o *CustconfBandwidthRateLimit) GetInitialBurstName() string`

GetInitialBurstName returns the InitialBurstName field if non-nil, zero value otherwise.

### GetInitialBurstNameOk

`func (o *CustconfBandwidthRateLimit) GetInitialBurstNameOk() (*string, bool)`

GetInitialBurstNameOk returns a tuple with the InitialBurstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialBurstName

`func (o *CustconfBandwidthRateLimit) SetInitialBurstName(v string)`

SetInitialBurstName sets InitialBurstName field to given value.

### HasInitialBurstName

`func (o *CustconfBandwidthRateLimit) HasInitialBurstName() bool`

HasInitialBurstName returns a boolean if a field has been set.

### GetSustainedRateName

`func (o *CustconfBandwidthRateLimit) GetSustainedRateName() string`

GetSustainedRateName returns the SustainedRateName field if non-nil, zero value otherwise.

### GetSustainedRateNameOk

`func (o *CustconfBandwidthRateLimit) GetSustainedRateNameOk() (*string, bool)`

GetSustainedRateNameOk returns a tuple with the SustainedRateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSustainedRateName

`func (o *CustconfBandwidthRateLimit) SetSustainedRateName(v string)`

SetSustainedRateName sets SustainedRateName field to given value.

### HasSustainedRateName

`func (o *CustconfBandwidthRateLimit) HasSustainedRateName() bool`

HasSustainedRateName returns a boolean if a field has been set.

### GetInitialBurstUnits

`func (o *CustconfBandwidthRateLimit) GetInitialBurstUnits() BandwidthRateLimitInitialBurstUnitsEnumWrapperValue`

GetInitialBurstUnits returns the InitialBurstUnits field if non-nil, zero value otherwise.

### GetInitialBurstUnitsOk

`func (o *CustconfBandwidthRateLimit) GetInitialBurstUnitsOk() (*BandwidthRateLimitInitialBurstUnitsEnumWrapperValue, bool)`

GetInitialBurstUnitsOk returns a tuple with the InitialBurstUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialBurstUnits

`func (o *CustconfBandwidthRateLimit) SetInitialBurstUnits(v BandwidthRateLimitInitialBurstUnitsEnumWrapperValue)`

SetInitialBurstUnits sets InitialBurstUnits field to given value.

### HasInitialBurstUnits

`func (o *CustconfBandwidthRateLimit) HasInitialBurstUnits() bool`

HasInitialBurstUnits returns a boolean if a field has been set.

### GetSustainedRateUnits

`func (o *CustconfBandwidthRateLimit) GetSustainedRateUnits() BandwidthRateLimitSustainedRateUnitsEnumWrapperValue`

GetSustainedRateUnits returns the SustainedRateUnits field if non-nil, zero value otherwise.

### GetSustainedRateUnitsOk

`func (o *CustconfBandwidthRateLimit) GetSustainedRateUnitsOk() (*BandwidthRateLimitSustainedRateUnitsEnumWrapperValue, bool)`

GetSustainedRateUnitsOk returns a tuple with the SustainedRateUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSustainedRateUnits

`func (o *CustconfBandwidthRateLimit) SetSustainedRateUnits(v BandwidthRateLimitSustainedRateUnitsEnumWrapperValue)`

SetSustainedRateUnits sets SustainedRateUnits field to given value.

### HasSustainedRateUnits

`func (o *CustconfBandwidthRateLimit) HasSustainedRateUnits() bool`

HasSustainedRateUnits returns a boolean if a field has been set.

### GetEnabled

`func (o *CustconfBandwidthRateLimit) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustconfBandwidthRateLimit) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustconfBandwidthRateLimit) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CustconfBandwidthRateLimit) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMethodFilter

`func (o *CustconfBandwidthRateLimit) GetMethodFilter() string`

GetMethodFilter returns the MethodFilter field if non-nil, zero value otherwise.

### GetMethodFilterOk

`func (o *CustconfBandwidthRateLimit) GetMethodFilterOk() (*string, bool)`

GetMethodFilterOk returns a tuple with the MethodFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodFilter

`func (o *CustconfBandwidthRateLimit) SetMethodFilter(v string)`

SetMethodFilter sets MethodFilter field to given value.

### HasMethodFilter

`func (o *CustconfBandwidthRateLimit) HasMethodFilter() bool`

HasMethodFilter returns a boolean if a field has been set.

### GetPathFilter

`func (o *CustconfBandwidthRateLimit) GetPathFilter() string`

GetPathFilter returns the PathFilter field if non-nil, zero value otherwise.

### GetPathFilterOk

`func (o *CustconfBandwidthRateLimit) GetPathFilterOk() (*string, bool)`

GetPathFilterOk returns a tuple with the PathFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathFilter

`func (o *CustconfBandwidthRateLimit) SetPathFilter(v string)`

SetPathFilter sets PathFilter field to given value.

### HasPathFilter

`func (o *CustconfBandwidthRateLimit) HasPathFilter() bool`

HasPathFilter returns a boolean if a field has been set.

### GetHeaderFilter

`func (o *CustconfBandwidthRateLimit) GetHeaderFilter() string`

GetHeaderFilter returns the HeaderFilter field if non-nil, zero value otherwise.

### GetHeaderFilterOk

`func (o *CustconfBandwidthRateLimit) GetHeaderFilterOk() (*string, bool)`

GetHeaderFilterOk returns a tuple with the HeaderFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderFilter

`func (o *CustconfBandwidthRateLimit) SetHeaderFilter(v string)`

SetHeaderFilter sets HeaderFilter field to given value.

### HasHeaderFilter

`func (o *CustconfBandwidthRateLimit) HasHeaderFilter() bool`

HasHeaderFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


