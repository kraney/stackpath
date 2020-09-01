# V2HttpConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | The URL of the service to monitor  The URL must begin with a case insensitive HTTP scheme of &#39;http&#39; or &#39;https&#39;. | [optional] 
**Method** | Pointer to **string** | The HTTP method used when the monitor makes a request to the service. | [optional] 
**Body** | Pointer to **string** | A base64 encoded HTTP request body to use when a monitor checks the service. | [optional] 
**Headers** | Pointer to [**[]V2Header**](v2Header.md) | A list of HTTP headers to add to the request when a monitor checks a service. | [optional] 
**Basic** | Pointer to [**V2HttpConfigurationBasicAuth**](v2HttpConfigurationBasicAuth.md) |  | [optional] 
**Digest** | Pointer to [**V2HttpConfigurationDigestAuth**](v2HttpConfigurationDigestAuth.md) |  | [optional] 
**ClientCertificate** | Pointer to [**V2HttpConfigurationClientCertificate**](v2HttpConfigurationClientCertificate.md) |  | [optional] 
**ValidateCertificate** | Pointer to **bool** | Whether or not to validate a service&#39;s certificate. | [optional] 

## Methods

### NewV2HttpConfiguration

`func NewV2HttpConfiguration() *V2HttpConfiguration`

NewV2HttpConfiguration instantiates a new V2HttpConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2HttpConfigurationWithDefaults

`func NewV2HttpConfigurationWithDefaults() *V2HttpConfiguration`

NewV2HttpConfigurationWithDefaults instantiates a new V2HttpConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *V2HttpConfiguration) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *V2HttpConfiguration) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *V2HttpConfiguration) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *V2HttpConfiguration) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetMethod

`func (o *V2HttpConfiguration) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *V2HttpConfiguration) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *V2HttpConfiguration) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *V2HttpConfiguration) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetBody

`func (o *V2HttpConfiguration) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *V2HttpConfiguration) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *V2HttpConfiguration) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *V2HttpConfiguration) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetHeaders

`func (o *V2HttpConfiguration) GetHeaders() []V2Header`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *V2HttpConfiguration) GetHeadersOk() (*[]V2Header, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *V2HttpConfiguration) SetHeaders(v []V2Header)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *V2HttpConfiguration) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetBasic

`func (o *V2HttpConfiguration) GetBasic() V2HttpConfigurationBasicAuth`

GetBasic returns the Basic field if non-nil, zero value otherwise.

### GetBasicOk

`func (o *V2HttpConfiguration) GetBasicOk() (*V2HttpConfigurationBasicAuth, bool)`

GetBasicOk returns a tuple with the Basic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasic

`func (o *V2HttpConfiguration) SetBasic(v V2HttpConfigurationBasicAuth)`

SetBasic sets Basic field to given value.

### HasBasic

`func (o *V2HttpConfiguration) HasBasic() bool`

HasBasic returns a boolean if a field has been set.

### GetDigest

`func (o *V2HttpConfiguration) GetDigest() V2HttpConfigurationDigestAuth`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *V2HttpConfiguration) GetDigestOk() (*V2HttpConfigurationDigestAuth, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *V2HttpConfiguration) SetDigest(v V2HttpConfigurationDigestAuth)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *V2HttpConfiguration) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetClientCertificate

`func (o *V2HttpConfiguration) GetClientCertificate() V2HttpConfigurationClientCertificate`

GetClientCertificate returns the ClientCertificate field if non-nil, zero value otherwise.

### GetClientCertificateOk

`func (o *V2HttpConfiguration) GetClientCertificateOk() (*V2HttpConfigurationClientCertificate, bool)`

GetClientCertificateOk returns a tuple with the ClientCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificate

`func (o *V2HttpConfiguration) SetClientCertificate(v V2HttpConfigurationClientCertificate)`

SetClientCertificate sets ClientCertificate field to given value.

### HasClientCertificate

`func (o *V2HttpConfiguration) HasClientCertificate() bool`

HasClientCertificate returns a boolean if a field has been set.

### GetValidateCertificate

`func (o *V2HttpConfiguration) GetValidateCertificate() bool`

GetValidateCertificate returns the ValidateCertificate field if non-nil, zero value otherwise.

### GetValidateCertificateOk

`func (o *V2HttpConfiguration) GetValidateCertificateOk() (*bool, bool)`

GetValidateCertificateOk returns a tuple with the ValidateCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateCertificate

`func (o *V2HttpConfiguration) SetValidateCertificate(v bool)`

SetValidateCertificate sets ValidateCertificate field to given value.

### HasValidateCertificate

`func (o *V2HttpConfiguration) HasValidateCertificate() bool`

HasValidateCertificate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


