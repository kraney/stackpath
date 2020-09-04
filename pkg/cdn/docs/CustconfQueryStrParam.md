# CustconfQueryStrParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | This is used by the API to perform conflict checking | [optional] 
**DispositionName** | Pointer to **string** | This is the name of the query string parameter which contains the name of the file to specify in the Content-Disposition HTTP response header. This setting is typically used by customers to configure a \&quot;friendly name\&quot; for URLs that have obfuscated file names. This setting controls the \&quot;filename\&quot; directive that is part of the Content-Disposition HTTP header. | [optional] 
**DispositionType** | Pointer to **string** | This is the name of the query string parameter which contains the disposition type to use in the Content-Disposition HTTP header. Typically, this value is set to \&quot;attachment,\&quot; but you may supply a custom string using this setting. | [optional] 
**DispositionOverride** | Pointer to **string** | This setting allows you to completely override the Content-Disposition HTTP header that the CDN caching servers use on a response. | [optional] 
**JumpToTimeStart** | Pointer to **string** | This is the name of the query string parameter that indicates to the CDN the start time offset of the video returned. This parameter is part of the jump-to-time feature that is initiated on a per request basis. | [optional] 
**JumpToTimeEnd** | Pointer to **string** | This is the name of the query string parameter that indicates to the CDN the end time offset of the video that should be returned. This parameter is part of the jump-to-time feature that is initiated on a per request basis. | [optional] 
**JumpToByteInitialBytes** | Pointer to **string** | This is the  name of the query string parameter that indicates to the CDN the initial bytes of a video that should be returned before sending the requested byte offset. This parameter is part of the jump-to-byte feature that is initiated on a per request basis. | [optional] 
**JumpToByteStartOffset** | Pointer to **string** | This is the name of the query string parameter that indicates to the CDN the specific offset into the video that is being requested. This parameter is part of the jump-to-byte feature that is initiated on a per request basis. | [optional] 
**RateLimitInitial** | Pointer to **string** | This is the name of the query string parameter that indicates to the CDN an initial burst rate to use when delivering a file. This parameter is part of the bandwidth limiting feature that is initiated on a per request basis. | [optional] 
**RateLimitSustained** | Pointer to **string** | This is the name of the query string parameter that indicates to the CDN the sustained rate being requested for the delivery of a file. This parameter is part of the bandwidth throttling feature that is initiated on a per request basis. | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewCustconfQueryStrParam

`func NewCustconfQueryStrParam() *CustconfQueryStrParam`

NewCustconfQueryStrParam instantiates a new CustconfQueryStrParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustconfQueryStrParamWithDefaults

`func NewCustconfQueryStrParamWithDefaults() *CustconfQueryStrParam`

NewCustconfQueryStrParamWithDefaults instantiates a new CustconfQueryStrParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustconfQueryStrParam) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustconfQueryStrParam) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustconfQueryStrParam) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustconfQueryStrParam) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDispositionName

`func (o *CustconfQueryStrParam) GetDispositionName() string`

GetDispositionName returns the DispositionName field if non-nil, zero value otherwise.

### GetDispositionNameOk

`func (o *CustconfQueryStrParam) GetDispositionNameOk() (*string, bool)`

GetDispositionNameOk returns a tuple with the DispositionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispositionName

`func (o *CustconfQueryStrParam) SetDispositionName(v string)`

SetDispositionName sets DispositionName field to given value.

### HasDispositionName

`func (o *CustconfQueryStrParam) HasDispositionName() bool`

HasDispositionName returns a boolean if a field has been set.

### GetDispositionType

`func (o *CustconfQueryStrParam) GetDispositionType() string`

GetDispositionType returns the DispositionType field if non-nil, zero value otherwise.

### GetDispositionTypeOk

`func (o *CustconfQueryStrParam) GetDispositionTypeOk() (*string, bool)`

GetDispositionTypeOk returns a tuple with the DispositionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispositionType

`func (o *CustconfQueryStrParam) SetDispositionType(v string)`

SetDispositionType sets DispositionType field to given value.

### HasDispositionType

`func (o *CustconfQueryStrParam) HasDispositionType() bool`

HasDispositionType returns a boolean if a field has been set.

### GetDispositionOverride

`func (o *CustconfQueryStrParam) GetDispositionOverride() string`

GetDispositionOverride returns the DispositionOverride field if non-nil, zero value otherwise.

### GetDispositionOverrideOk

`func (o *CustconfQueryStrParam) GetDispositionOverrideOk() (*string, bool)`

GetDispositionOverrideOk returns a tuple with the DispositionOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispositionOverride

`func (o *CustconfQueryStrParam) SetDispositionOverride(v string)`

SetDispositionOverride sets DispositionOverride field to given value.

### HasDispositionOverride

`func (o *CustconfQueryStrParam) HasDispositionOverride() bool`

HasDispositionOverride returns a boolean if a field has been set.

### GetJumpToTimeStart

`func (o *CustconfQueryStrParam) GetJumpToTimeStart() string`

GetJumpToTimeStart returns the JumpToTimeStart field if non-nil, zero value otherwise.

### GetJumpToTimeStartOk

`func (o *CustconfQueryStrParam) GetJumpToTimeStartOk() (*string, bool)`

GetJumpToTimeStartOk returns a tuple with the JumpToTimeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJumpToTimeStart

`func (o *CustconfQueryStrParam) SetJumpToTimeStart(v string)`

SetJumpToTimeStart sets JumpToTimeStart field to given value.

### HasJumpToTimeStart

`func (o *CustconfQueryStrParam) HasJumpToTimeStart() bool`

HasJumpToTimeStart returns a boolean if a field has been set.

### GetJumpToTimeEnd

`func (o *CustconfQueryStrParam) GetJumpToTimeEnd() string`

GetJumpToTimeEnd returns the JumpToTimeEnd field if non-nil, zero value otherwise.

### GetJumpToTimeEndOk

`func (o *CustconfQueryStrParam) GetJumpToTimeEndOk() (*string, bool)`

GetJumpToTimeEndOk returns a tuple with the JumpToTimeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJumpToTimeEnd

`func (o *CustconfQueryStrParam) SetJumpToTimeEnd(v string)`

SetJumpToTimeEnd sets JumpToTimeEnd field to given value.

### HasJumpToTimeEnd

`func (o *CustconfQueryStrParam) HasJumpToTimeEnd() bool`

HasJumpToTimeEnd returns a boolean if a field has been set.

### GetJumpToByteInitialBytes

`func (o *CustconfQueryStrParam) GetJumpToByteInitialBytes() string`

GetJumpToByteInitialBytes returns the JumpToByteInitialBytes field if non-nil, zero value otherwise.

### GetJumpToByteInitialBytesOk

`func (o *CustconfQueryStrParam) GetJumpToByteInitialBytesOk() (*string, bool)`

GetJumpToByteInitialBytesOk returns a tuple with the JumpToByteInitialBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJumpToByteInitialBytes

`func (o *CustconfQueryStrParam) SetJumpToByteInitialBytes(v string)`

SetJumpToByteInitialBytes sets JumpToByteInitialBytes field to given value.

### HasJumpToByteInitialBytes

`func (o *CustconfQueryStrParam) HasJumpToByteInitialBytes() bool`

HasJumpToByteInitialBytes returns a boolean if a field has been set.

### GetJumpToByteStartOffset

`func (o *CustconfQueryStrParam) GetJumpToByteStartOffset() string`

GetJumpToByteStartOffset returns the JumpToByteStartOffset field if non-nil, zero value otherwise.

### GetJumpToByteStartOffsetOk

`func (o *CustconfQueryStrParam) GetJumpToByteStartOffsetOk() (*string, bool)`

GetJumpToByteStartOffsetOk returns a tuple with the JumpToByteStartOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJumpToByteStartOffset

`func (o *CustconfQueryStrParam) SetJumpToByteStartOffset(v string)`

SetJumpToByteStartOffset sets JumpToByteStartOffset field to given value.

### HasJumpToByteStartOffset

`func (o *CustconfQueryStrParam) HasJumpToByteStartOffset() bool`

HasJumpToByteStartOffset returns a boolean if a field has been set.

### GetRateLimitInitial

`func (o *CustconfQueryStrParam) GetRateLimitInitial() string`

GetRateLimitInitial returns the RateLimitInitial field if non-nil, zero value otherwise.

### GetRateLimitInitialOk

`func (o *CustconfQueryStrParam) GetRateLimitInitialOk() (*string, bool)`

GetRateLimitInitialOk returns a tuple with the RateLimitInitial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitInitial

`func (o *CustconfQueryStrParam) SetRateLimitInitial(v string)`

SetRateLimitInitial sets RateLimitInitial field to given value.

### HasRateLimitInitial

`func (o *CustconfQueryStrParam) HasRateLimitInitial() bool`

HasRateLimitInitial returns a boolean if a field has been set.

### GetRateLimitSustained

`func (o *CustconfQueryStrParam) GetRateLimitSustained() string`

GetRateLimitSustained returns the RateLimitSustained field if non-nil, zero value otherwise.

### GetRateLimitSustainedOk

`func (o *CustconfQueryStrParam) GetRateLimitSustainedOk() (*string, bool)`

GetRateLimitSustainedOk returns a tuple with the RateLimitSustained field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitSustained

`func (o *CustconfQueryStrParam) SetRateLimitSustained(v string)`

SetRateLimitSustained sets RateLimitSustained field to given value.

### HasRateLimitSustained

`func (o *CustconfQueryStrParam) HasRateLimitSustained() bool`

HasRateLimitSustained returns a boolean if a field has been set.

### GetEnabled

`func (o *CustconfQueryStrParam) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustconfQueryStrParam) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustconfQueryStrParam) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CustconfQueryStrParam) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


