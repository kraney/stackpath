# CdnScanOriginResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAddress** | Pointer to **string** | The IP address that was scanned | [optional] 
**Hostname** | Pointer to **string** | The hostname that was scanned | [optional] 
**SslDetails** | Pointer to [**ScanOriginResponseOriginScanSSLDetails**](ScanOriginResponseOriginScanSSLDetails.md) |  | [optional] 
**DomainInUse** | Pointer to **bool** | Whether or not the scanned domain is already in use on the StackPath platform | [optional] 

## Methods

### NewCdnScanOriginResponse

`func NewCdnScanOriginResponse() *CdnScanOriginResponse`

NewCdnScanOriginResponse instantiates a new CdnScanOriginResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdnScanOriginResponseWithDefaults

`func NewCdnScanOriginResponseWithDefaults() *CdnScanOriginResponse`

NewCdnScanOriginResponseWithDefaults instantiates a new CdnScanOriginResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpAddress

`func (o *CdnScanOriginResponse) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *CdnScanOriginResponse) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *CdnScanOriginResponse) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *CdnScanOriginResponse) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetHostname

`func (o *CdnScanOriginResponse) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *CdnScanOriginResponse) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *CdnScanOriginResponse) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *CdnScanOriginResponse) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetSslDetails

`func (o *CdnScanOriginResponse) GetSslDetails() ScanOriginResponseOriginScanSSLDetails`

GetSslDetails returns the SslDetails field if non-nil, zero value otherwise.

### GetSslDetailsOk

`func (o *CdnScanOriginResponse) GetSslDetailsOk() (*ScanOriginResponseOriginScanSSLDetails, bool)`

GetSslDetailsOk returns a tuple with the SslDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslDetails

`func (o *CdnScanOriginResponse) SetSslDetails(v ScanOriginResponseOriginScanSSLDetails)`

SetSslDetails sets SslDetails field to given value.

### HasSslDetails

`func (o *CdnScanOriginResponse) HasSslDetails() bool`

HasSslDetails returns a boolean if a field has been set.

### GetDomainInUse

`func (o *CdnScanOriginResponse) GetDomainInUse() bool`

GetDomainInUse returns the DomainInUse field if non-nil, zero value otherwise.

### GetDomainInUseOk

`func (o *CdnScanOriginResponse) GetDomainInUseOk() (*bool, bool)`

GetDomainInUseOk returns a tuple with the DomainInUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainInUse

`func (o *CdnScanOriginResponse) SetDomainInUse(v bool)`

SetDomainInUse sets DomainInUse field to given value.

### HasDomainInUse

`func (o *CdnScanOriginResponse) HasDomainInUse() bool`

HasDomainInUse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


