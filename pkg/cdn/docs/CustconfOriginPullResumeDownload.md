# CustconfOriginPullResumeDownload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | This is used by the API to perform conflict checking | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**OriginalStatusCodeMatch** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. Comma separated list of glob patterns that indicate which origin pulls this policy applies to based on the status code of the original origin response. This feature is limited to 2xx responses, but this policy provides fine control, such as permitting resume of 200 responses by not 206, or for adding a 2xx response code other than 200 or 206. | [optional] 
**MinimumBodyBytesConsumed** | Pointer to **string** | Minimum number of body bytes that CDN must have consumed during an Origin Pull before encountering an error before it is permitted to attempt resuming the download. This value does not include the headers bytes. | [optional] 
**MaximumAttempts** | Pointer to **float32** | Maximum number of times beyond the initial request that the CDN is permitted to attempt resuming an Origin Pull download. | [optional] 
**RequireOriginRangeSupport** | Pointer to **bool** | The CDN resumes an Origin Pull even if the origin does not support range requests. If the origin does not support range requests and/or returns a 200 response for the same version given in the original response, the CDN fast-forwards (reads and discards bytes) until it reaches its previous position in the asset. | [optional] 
**EtagValidation** | Pointer to [**OriginPullResumeDownloadEtagValidationEnumWrapperValue**](OriginPullResumeDownloadEtagValidationEnumWrapperValue.md) |  | [optional] [default to "UNKNOWN"]
**HeaderFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**PathFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 

## Methods

### NewCustconfOriginPullResumeDownload

`func NewCustconfOriginPullResumeDownload() *CustconfOriginPullResumeDownload`

NewCustconfOriginPullResumeDownload instantiates a new CustconfOriginPullResumeDownload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustconfOriginPullResumeDownloadWithDefaults

`func NewCustconfOriginPullResumeDownloadWithDefaults() *CustconfOriginPullResumeDownload`

NewCustconfOriginPullResumeDownloadWithDefaults instantiates a new CustconfOriginPullResumeDownload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustconfOriginPullResumeDownload) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustconfOriginPullResumeDownload) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustconfOriginPullResumeDownload) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustconfOriginPullResumeDownload) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEnabled

`func (o *CustconfOriginPullResumeDownload) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustconfOriginPullResumeDownload) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustconfOriginPullResumeDownload) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CustconfOriginPullResumeDownload) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetOriginalStatusCodeMatch

`func (o *CustconfOriginPullResumeDownload) GetOriginalStatusCodeMatch() string`

GetOriginalStatusCodeMatch returns the OriginalStatusCodeMatch field if non-nil, zero value otherwise.

### GetOriginalStatusCodeMatchOk

`func (o *CustconfOriginPullResumeDownload) GetOriginalStatusCodeMatchOk() (*string, bool)`

GetOriginalStatusCodeMatchOk returns a tuple with the OriginalStatusCodeMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalStatusCodeMatch

`func (o *CustconfOriginPullResumeDownload) SetOriginalStatusCodeMatch(v string)`

SetOriginalStatusCodeMatch sets OriginalStatusCodeMatch field to given value.

### HasOriginalStatusCodeMatch

`func (o *CustconfOriginPullResumeDownload) HasOriginalStatusCodeMatch() bool`

HasOriginalStatusCodeMatch returns a boolean if a field has been set.

### GetMinimumBodyBytesConsumed

`func (o *CustconfOriginPullResumeDownload) GetMinimumBodyBytesConsumed() string`

GetMinimumBodyBytesConsumed returns the MinimumBodyBytesConsumed field if non-nil, zero value otherwise.

### GetMinimumBodyBytesConsumedOk

`func (o *CustconfOriginPullResumeDownload) GetMinimumBodyBytesConsumedOk() (*string, bool)`

GetMinimumBodyBytesConsumedOk returns a tuple with the MinimumBodyBytesConsumed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumBodyBytesConsumed

`func (o *CustconfOriginPullResumeDownload) SetMinimumBodyBytesConsumed(v string)`

SetMinimumBodyBytesConsumed sets MinimumBodyBytesConsumed field to given value.

### HasMinimumBodyBytesConsumed

`func (o *CustconfOriginPullResumeDownload) HasMinimumBodyBytesConsumed() bool`

HasMinimumBodyBytesConsumed returns a boolean if a field has been set.

### GetMaximumAttempts

`func (o *CustconfOriginPullResumeDownload) GetMaximumAttempts() float32`

GetMaximumAttempts returns the MaximumAttempts field if non-nil, zero value otherwise.

### GetMaximumAttemptsOk

`func (o *CustconfOriginPullResumeDownload) GetMaximumAttemptsOk() (*float32, bool)`

GetMaximumAttemptsOk returns a tuple with the MaximumAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumAttempts

`func (o *CustconfOriginPullResumeDownload) SetMaximumAttempts(v float32)`

SetMaximumAttempts sets MaximumAttempts field to given value.

### HasMaximumAttempts

`func (o *CustconfOriginPullResumeDownload) HasMaximumAttempts() bool`

HasMaximumAttempts returns a boolean if a field has been set.

### GetRequireOriginRangeSupport

`func (o *CustconfOriginPullResumeDownload) GetRequireOriginRangeSupport() bool`

GetRequireOriginRangeSupport returns the RequireOriginRangeSupport field if non-nil, zero value otherwise.

### GetRequireOriginRangeSupportOk

`func (o *CustconfOriginPullResumeDownload) GetRequireOriginRangeSupportOk() (*bool, bool)`

GetRequireOriginRangeSupportOk returns a tuple with the RequireOriginRangeSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireOriginRangeSupport

`func (o *CustconfOriginPullResumeDownload) SetRequireOriginRangeSupport(v bool)`

SetRequireOriginRangeSupport sets RequireOriginRangeSupport field to given value.

### HasRequireOriginRangeSupport

`func (o *CustconfOriginPullResumeDownload) HasRequireOriginRangeSupport() bool`

HasRequireOriginRangeSupport returns a boolean if a field has been set.

### GetEtagValidation

`func (o *CustconfOriginPullResumeDownload) GetEtagValidation() OriginPullResumeDownloadEtagValidationEnumWrapperValue`

GetEtagValidation returns the EtagValidation field if non-nil, zero value otherwise.

### GetEtagValidationOk

`func (o *CustconfOriginPullResumeDownload) GetEtagValidationOk() (*OriginPullResumeDownloadEtagValidationEnumWrapperValue, bool)`

GetEtagValidationOk returns a tuple with the EtagValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtagValidation

`func (o *CustconfOriginPullResumeDownload) SetEtagValidation(v OriginPullResumeDownloadEtagValidationEnumWrapperValue)`

SetEtagValidation sets EtagValidation field to given value.

### HasEtagValidation

`func (o *CustconfOriginPullResumeDownload) HasEtagValidation() bool`

HasEtagValidation returns a boolean if a field has been set.

### GetHeaderFilter

`func (o *CustconfOriginPullResumeDownload) GetHeaderFilter() string`

GetHeaderFilter returns the HeaderFilter field if non-nil, zero value otherwise.

### GetHeaderFilterOk

`func (o *CustconfOriginPullResumeDownload) GetHeaderFilterOk() (*string, bool)`

GetHeaderFilterOk returns a tuple with the HeaderFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderFilter

`func (o *CustconfOriginPullResumeDownload) SetHeaderFilter(v string)`

SetHeaderFilter sets HeaderFilter field to given value.

### HasHeaderFilter

`func (o *CustconfOriginPullResumeDownload) HasHeaderFilter() bool`

HasHeaderFilter returns a boolean if a field has been set.

### GetPathFilter

`func (o *CustconfOriginPullResumeDownload) GetPathFilter() string`

GetPathFilter returns the PathFilter field if non-nil, zero value otherwise.

### GetPathFilterOk

`func (o *CustconfOriginPullResumeDownload) GetPathFilterOk() (*string, bool)`

GetPathFilterOk returns a tuple with the PathFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathFilter

`func (o *CustconfOriginPullResumeDownload) SetPathFilter(v string)`

SetPathFilter sets PathFilter field to given value.

### HasPathFilter

`func (o *CustconfOriginPullResumeDownload) HasPathFilter() bool`

HasPathFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


