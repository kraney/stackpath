# ScanOriginResponseOriginScanSSLDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Valid** | Pointer to **bool** | Whether or not the scanned domain has a valid SSL certificate | [optional] 
**Error** | Pointer to **string** | Errors encountered processing HTTPS requests to the scanned domain | [optional] 

## Methods

### NewScanOriginResponseOriginScanSSLDetails

`func NewScanOriginResponseOriginScanSSLDetails() *ScanOriginResponseOriginScanSSLDetails`

NewScanOriginResponseOriginScanSSLDetails instantiates a new ScanOriginResponseOriginScanSSLDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScanOriginResponseOriginScanSSLDetailsWithDefaults

`func NewScanOriginResponseOriginScanSSLDetailsWithDefaults() *ScanOriginResponseOriginScanSSLDetails`

NewScanOriginResponseOriginScanSSLDetailsWithDefaults instantiates a new ScanOriginResponseOriginScanSSLDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValid

`func (o *ScanOriginResponseOriginScanSSLDetails) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *ScanOriginResponseOriginScanSSLDetails) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *ScanOriginResponseOriginScanSSLDetails) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *ScanOriginResponseOriginScanSSLDetails) HasValid() bool`

HasValid returns a boolean if a field has been set.

### GetError

`func (o *ScanOriginResponseOriginScanSSLDetails) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ScanOriginResponseOriginScanSSLDetails) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ScanOriginResponseOriginScanSSLDetails) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ScanOriginResponseOriginScanSSLDetails) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


