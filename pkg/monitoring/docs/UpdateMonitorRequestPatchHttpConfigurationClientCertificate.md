# UpdateMonitorRequestPatchHttpConfigurationClientCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicCertificate** | Pointer to **string** | The X.509 PEM formatted public certificate used for authentication. | [optional] 
**PrivateKey** | Pointer to **string** | The X.509 PEM formatted private key used for authentication. | [optional] 
**CaBundle** | Pointer to **string** | The X.509 PEM formatted certificate authority bundle used for authentication. | [optional] 

## Methods

### NewUpdateMonitorRequestPatchHttpConfigurationClientCertificate

`func NewUpdateMonitorRequestPatchHttpConfigurationClientCertificate() *UpdateMonitorRequestPatchHttpConfigurationClientCertificate`

NewUpdateMonitorRequestPatchHttpConfigurationClientCertificate instantiates a new UpdateMonitorRequestPatchHttpConfigurationClientCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMonitorRequestPatchHttpConfigurationClientCertificateWithDefaults

`func NewUpdateMonitorRequestPatchHttpConfigurationClientCertificateWithDefaults() *UpdateMonitorRequestPatchHttpConfigurationClientCertificate`

NewUpdateMonitorRequestPatchHttpConfigurationClientCertificateWithDefaults instantiates a new UpdateMonitorRequestPatchHttpConfigurationClientCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicCertificate

`func (o *UpdateMonitorRequestPatchHttpConfigurationClientCertificate) GetPublicCertificate() string`

GetPublicCertificate returns the PublicCertificate field if non-nil, zero value otherwise.

### GetPublicCertificateOk

`func (o *UpdateMonitorRequestPatchHttpConfigurationClientCertificate) GetPublicCertificateOk() (*string, bool)`

GetPublicCertificateOk returns a tuple with the PublicCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicCertificate

`func (o *UpdateMonitorRequestPatchHttpConfigurationClientCertificate) SetPublicCertificate(v string)`

SetPublicCertificate sets PublicCertificate field to given value.

### HasPublicCertificate

`func (o *UpdateMonitorRequestPatchHttpConfigurationClientCertificate) HasPublicCertificate() bool`

HasPublicCertificate returns a boolean if a field has been set.

### GetPrivateKey

`func (o *UpdateMonitorRequestPatchHttpConfigurationClientCertificate) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *UpdateMonitorRequestPatchHttpConfigurationClientCertificate) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *UpdateMonitorRequestPatchHttpConfigurationClientCertificate) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *UpdateMonitorRequestPatchHttpConfigurationClientCertificate) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetCaBundle

`func (o *UpdateMonitorRequestPatchHttpConfigurationClientCertificate) GetCaBundle() string`

GetCaBundle returns the CaBundle field if non-nil, zero value otherwise.

### GetCaBundleOk

`func (o *UpdateMonitorRequestPatchHttpConfigurationClientCertificate) GetCaBundleOk() (*string, bool)`

GetCaBundleOk returns a tuple with the CaBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaBundle

`func (o *UpdateMonitorRequestPatchHttpConfigurationClientCertificate) SetCaBundle(v string)`

SetCaBundle sets CaBundle field to given value.

### HasCaBundle

`func (o *UpdateMonitorRequestPatchHttpConfigurationClientCertificate) HasCaBundle() bool`

HasCaBundle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


