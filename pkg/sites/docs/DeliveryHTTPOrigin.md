# DeliveryHTTPOrigin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** | The origin&#39;s HTTP request path | [optional] 
**Hostname** | Pointer to **string** | The origin&#39;s HTTP request hostname | [optional] 
**Port** | Pointer to **int32** | The origin&#39;s HTTP request port  Set this to 0 to remove this value | [optional] 
**SecurePort** | Pointer to **int32** | The origin&#39;s HTTPS request port  Set this to 0 to remove this value | [optional] 
**Authentication** | Pointer to [**DeliveryOriginAuthentication**](deliveryOriginAuthentication.md) |  | [optional] 
**VerifyCertificate** | Pointer to **bool** | Verify the origin&#39;s SSL certificate when requesting from the origin | [optional] 
**CertificateCommonName** | Pointer to **string** | The CommonName to validate SSL origin requests from. This value needs to be provided when enabling &#x60;verify_certificate&#x60;. | [optional] 

## Methods

### NewDeliveryHTTPOrigin

`func NewDeliveryHTTPOrigin() *DeliveryHTTPOrigin`

NewDeliveryHTTPOrigin instantiates a new DeliveryHTTPOrigin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryHTTPOriginWithDefaults

`func NewDeliveryHTTPOriginWithDefaults() *DeliveryHTTPOrigin`

NewDeliveryHTTPOriginWithDefaults instantiates a new DeliveryHTTPOrigin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *DeliveryHTTPOrigin) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DeliveryHTTPOrigin) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DeliveryHTTPOrigin) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *DeliveryHTTPOrigin) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetHostname

`func (o *DeliveryHTTPOrigin) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *DeliveryHTTPOrigin) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *DeliveryHTTPOrigin) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *DeliveryHTTPOrigin) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetPort

`func (o *DeliveryHTTPOrigin) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DeliveryHTTPOrigin) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DeliveryHTTPOrigin) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *DeliveryHTTPOrigin) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSecurePort

`func (o *DeliveryHTTPOrigin) GetSecurePort() int32`

GetSecurePort returns the SecurePort field if non-nil, zero value otherwise.

### GetSecurePortOk

`func (o *DeliveryHTTPOrigin) GetSecurePortOk() (*int32, bool)`

GetSecurePortOk returns a tuple with the SecurePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurePort

`func (o *DeliveryHTTPOrigin) SetSecurePort(v int32)`

SetSecurePort sets SecurePort field to given value.

### HasSecurePort

`func (o *DeliveryHTTPOrigin) HasSecurePort() bool`

HasSecurePort returns a boolean if a field has been set.

### GetAuthentication

`func (o *DeliveryHTTPOrigin) GetAuthentication() DeliveryOriginAuthentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *DeliveryHTTPOrigin) GetAuthenticationOk() (*DeliveryOriginAuthentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *DeliveryHTTPOrigin) SetAuthentication(v DeliveryOriginAuthentication)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *DeliveryHTTPOrigin) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetVerifyCertificate

`func (o *DeliveryHTTPOrigin) GetVerifyCertificate() bool`

GetVerifyCertificate returns the VerifyCertificate field if non-nil, zero value otherwise.

### GetVerifyCertificateOk

`func (o *DeliveryHTTPOrigin) GetVerifyCertificateOk() (*bool, bool)`

GetVerifyCertificateOk returns a tuple with the VerifyCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyCertificate

`func (o *DeliveryHTTPOrigin) SetVerifyCertificate(v bool)`

SetVerifyCertificate sets VerifyCertificate field to given value.

### HasVerifyCertificate

`func (o *DeliveryHTTPOrigin) HasVerifyCertificate() bool`

HasVerifyCertificate returns a boolean if a field has been set.

### GetCertificateCommonName

`func (o *DeliveryHTTPOrigin) GetCertificateCommonName() string`

GetCertificateCommonName returns the CertificateCommonName field if non-nil, zero value otherwise.

### GetCertificateCommonNameOk

`func (o *DeliveryHTTPOrigin) GetCertificateCommonNameOk() (*string, bool)`

GetCertificateCommonNameOk returns a tuple with the CertificateCommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateCommonName

`func (o *DeliveryHTTPOrigin) SetCertificateCommonName(v string)`

SetCertificateCommonName sets CertificateCommonName field to given value.

### HasCertificateCommonName

`func (o *DeliveryHTTPOrigin) HasCertificateCommonName() bool`

HasCertificateCommonName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


