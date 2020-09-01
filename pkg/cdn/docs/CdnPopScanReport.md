# CdnPopScanReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PopCode** | Pointer to **string** | The IATA formatted location code of the StackPath POP that produced a scan report | [optional] 
**ConnectMs** | Pointer to **float32** | The amount of time in milliseconds that a POP scan took to perform an initial connection handshake | [optional] 
**DnsMs** | Pointer to **float32** | The amount of time in milliseconds that a POP scan took to resolve the target&#39;s DNS entry | [optional] 
**DownloadMs** | Pointer to **float32** | The amount of time in milliseconds that a POP scan took to download the target&#39;s contents | [optional] 
**FirstByteMs** | Pointer to **float32** | The amount of time in milliseconds that a POP scan took from initial connection to starting to receive the target&#39;s contents | [optional] 
**SslMs** | Pointer to **float32** | The amount of time in milliseconds that a POP scan took to perform an SSL handshake | [optional] 
**TotalMs** | Pointer to **float32** | The total amount of time in milliseconds that a POP scan took to complete | [optional] 

## Methods

### NewCdnPopScanReport

`func NewCdnPopScanReport() *CdnPopScanReport`

NewCdnPopScanReport instantiates a new CdnPopScanReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdnPopScanReportWithDefaults

`func NewCdnPopScanReportWithDefaults() *CdnPopScanReport`

NewCdnPopScanReportWithDefaults instantiates a new CdnPopScanReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPopCode

`func (o *CdnPopScanReport) GetPopCode() string`

GetPopCode returns the PopCode field if non-nil, zero value otherwise.

### GetPopCodeOk

`func (o *CdnPopScanReport) GetPopCodeOk() (*string, bool)`

GetPopCodeOk returns a tuple with the PopCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopCode

`func (o *CdnPopScanReport) SetPopCode(v string)`

SetPopCode sets PopCode field to given value.

### HasPopCode

`func (o *CdnPopScanReport) HasPopCode() bool`

HasPopCode returns a boolean if a field has been set.

### GetConnectMs

`func (o *CdnPopScanReport) GetConnectMs() float32`

GetConnectMs returns the ConnectMs field if non-nil, zero value otherwise.

### GetConnectMsOk

`func (o *CdnPopScanReport) GetConnectMsOk() (*float32, bool)`

GetConnectMsOk returns a tuple with the ConnectMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectMs

`func (o *CdnPopScanReport) SetConnectMs(v float32)`

SetConnectMs sets ConnectMs field to given value.

### HasConnectMs

`func (o *CdnPopScanReport) HasConnectMs() bool`

HasConnectMs returns a boolean if a field has been set.

### GetDnsMs

`func (o *CdnPopScanReport) GetDnsMs() float32`

GetDnsMs returns the DnsMs field if non-nil, zero value otherwise.

### GetDnsMsOk

`func (o *CdnPopScanReport) GetDnsMsOk() (*float32, bool)`

GetDnsMsOk returns a tuple with the DnsMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsMs

`func (o *CdnPopScanReport) SetDnsMs(v float32)`

SetDnsMs sets DnsMs field to given value.

### HasDnsMs

`func (o *CdnPopScanReport) HasDnsMs() bool`

HasDnsMs returns a boolean if a field has been set.

### GetDownloadMs

`func (o *CdnPopScanReport) GetDownloadMs() float32`

GetDownloadMs returns the DownloadMs field if non-nil, zero value otherwise.

### GetDownloadMsOk

`func (o *CdnPopScanReport) GetDownloadMsOk() (*float32, bool)`

GetDownloadMsOk returns a tuple with the DownloadMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadMs

`func (o *CdnPopScanReport) SetDownloadMs(v float32)`

SetDownloadMs sets DownloadMs field to given value.

### HasDownloadMs

`func (o *CdnPopScanReport) HasDownloadMs() bool`

HasDownloadMs returns a boolean if a field has been set.

### GetFirstByteMs

`func (o *CdnPopScanReport) GetFirstByteMs() float32`

GetFirstByteMs returns the FirstByteMs field if non-nil, zero value otherwise.

### GetFirstByteMsOk

`func (o *CdnPopScanReport) GetFirstByteMsOk() (*float32, bool)`

GetFirstByteMsOk returns a tuple with the FirstByteMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstByteMs

`func (o *CdnPopScanReport) SetFirstByteMs(v float32)`

SetFirstByteMs sets FirstByteMs field to given value.

### HasFirstByteMs

`func (o *CdnPopScanReport) HasFirstByteMs() bool`

HasFirstByteMs returns a boolean if a field has been set.

### GetSslMs

`func (o *CdnPopScanReport) GetSslMs() float32`

GetSslMs returns the SslMs field if non-nil, zero value otherwise.

### GetSslMsOk

`func (o *CdnPopScanReport) GetSslMsOk() (*float32, bool)`

GetSslMsOk returns a tuple with the SslMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslMs

`func (o *CdnPopScanReport) SetSslMs(v float32)`

SetSslMs sets SslMs field to given value.

### HasSslMs

`func (o *CdnPopScanReport) HasSslMs() bool`

HasSslMs returns a boolean if a field has been set.

### GetTotalMs

`func (o *CdnPopScanReport) GetTotalMs() float32`

GetTotalMs returns the TotalMs field if non-nil, zero value otherwise.

### GetTotalMsOk

`func (o *CdnPopScanReport) GetTotalMsOk() (*float32, bool)`

GetTotalMsOk returns a tuple with the TotalMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMs

`func (o *CdnPopScanReport) SetTotalMs(v float32)`

SetTotalMs sets TotalMs field to given value.

### HasTotalMs

`func (o *CdnPopScanReport) HasTotalMs() bool`

HasTotalMs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


