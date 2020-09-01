# UpdateMonitorRequestPatchHttpConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | The URL of the service that will be monitored.  The URL must begin with an HTTP scheme of &#39;http&#39; or &#39;https&#39;. | [optional] 
**Method** | Pointer to **string** | The HTTP method used when the monitor makes a request to the service. | [optional] 
**Body** | Pointer to **string** | A base64 encoded HTTP request body to use when a monitor checks the service. | [optional] 
**Headers** | Pointer to [**PatchHttpConfigurationHeaderValue**](PatchHttpConfigurationHeaderValue.md) |  | [optional] 
**Basic** | Pointer to [**UpdateMonitorRequestPatchHttpConfigurationBasicAuth**](UpdateMonitorRequestPatchHttpConfigurationBasicAuth.md) |  | [optional] 
**Digest** | Pointer to [**UpdateMonitorRequestPatchHttpConfigurationDigestAuth**](UpdateMonitorRequestPatchHttpConfigurationDigestAuth.md) |  | [optional] 
**ClientCertificate** | Pointer to [**UpdateMonitorRequestPatchHttpConfigurationClientCertificate**](UpdateMonitorRequestPatchHttpConfigurationClientCertificate.md) |  | [optional] 
**ValidateCertificate** | Pointer to **bool** | A flag for validating a service&#39;s certificate. | [optional] 

## Methods

### NewUpdateMonitorRequestPatchHttpConfiguration

`func NewUpdateMonitorRequestPatchHttpConfiguration() *UpdateMonitorRequestPatchHttpConfiguration`

NewUpdateMonitorRequestPatchHttpConfiguration instantiates a new UpdateMonitorRequestPatchHttpConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMonitorRequestPatchHttpConfigurationWithDefaults

`func NewUpdateMonitorRequestPatchHttpConfigurationWithDefaults() *UpdateMonitorRequestPatchHttpConfiguration`

NewUpdateMonitorRequestPatchHttpConfigurationWithDefaults instantiates a new UpdateMonitorRequestPatchHttpConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *UpdateMonitorRequestPatchHttpConfiguration) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UpdateMonitorRequestPatchHttpConfiguration) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UpdateMonitorRequestPatchHttpConfiguration) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UpdateMonitorRequestPatchHttpConfiguration) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetMethod

`func (o *UpdateMonitorRequestPatchHttpConfiguration) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *UpdateMonitorRequestPatchHttpConfiguration) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *UpdateMonitorRequestPatchHttpConfiguration) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *UpdateMonitorRequestPatchHttpConfiguration) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetBody

`func (o *UpdateMonitorRequestPatchHttpConfiguration) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *UpdateMonitorRequestPatchHttpConfiguration) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *UpdateMonitorRequestPatchHttpConfiguration) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *UpdateMonitorRequestPatchHttpConfiguration) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetHeaders

`func (o *UpdateMonitorRequestPatchHttpConfiguration) GetHeaders() PatchHttpConfigurationHeaderValue`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *UpdateMonitorRequestPatchHttpConfiguration) GetHeadersOk() (*PatchHttpConfigurationHeaderValue, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *UpdateMonitorRequestPatchHttpConfiguration) SetHeaders(v PatchHttpConfigurationHeaderValue)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *UpdateMonitorRequestPatchHttpConfiguration) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetBasic

`func (o *UpdateMonitorRequestPatchHttpConfiguration) GetBasic() UpdateMonitorRequestPatchHttpConfigurationBasicAuth`

GetBasic returns the Basic field if non-nil, zero value otherwise.

### GetBasicOk

`func (o *UpdateMonitorRequestPatchHttpConfiguration) GetBasicOk() (*UpdateMonitorRequestPatchHttpConfigurationBasicAuth, bool)`

GetBasicOk returns a tuple with the Basic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasic

`func (o *UpdateMonitorRequestPatchHttpConfiguration) SetBasic(v UpdateMonitorRequestPatchHttpConfigurationBasicAuth)`

SetBasic sets Basic field to given value.

### HasBasic

`func (o *UpdateMonitorRequestPatchHttpConfiguration) HasBasic() bool`

HasBasic returns a boolean if a field has been set.

### GetDigest

`func (o *UpdateMonitorRequestPatchHttpConfiguration) GetDigest() UpdateMonitorRequestPatchHttpConfigurationDigestAuth`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *UpdateMonitorRequestPatchHttpConfiguration) GetDigestOk() (*UpdateMonitorRequestPatchHttpConfigurationDigestAuth, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *UpdateMonitorRequestPatchHttpConfiguration) SetDigest(v UpdateMonitorRequestPatchHttpConfigurationDigestAuth)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *UpdateMonitorRequestPatchHttpConfiguration) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetClientCertificate

`func (o *UpdateMonitorRequestPatchHttpConfiguration) GetClientCertificate() UpdateMonitorRequestPatchHttpConfigurationClientCertificate`

GetClientCertificate returns the ClientCertificate field if non-nil, zero value otherwise.

### GetClientCertificateOk

`func (o *UpdateMonitorRequestPatchHttpConfiguration) GetClientCertificateOk() (*UpdateMonitorRequestPatchHttpConfigurationClientCertificate, bool)`

GetClientCertificateOk returns a tuple with the ClientCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificate

`func (o *UpdateMonitorRequestPatchHttpConfiguration) SetClientCertificate(v UpdateMonitorRequestPatchHttpConfigurationClientCertificate)`

SetClientCertificate sets ClientCertificate field to given value.

### HasClientCertificate

`func (o *UpdateMonitorRequestPatchHttpConfiguration) HasClientCertificate() bool`

HasClientCertificate returns a boolean if a field has been set.

### GetValidateCertificate

`func (o *UpdateMonitorRequestPatchHttpConfiguration) GetValidateCertificate() bool`

GetValidateCertificate returns the ValidateCertificate field if non-nil, zero value otherwise.

### GetValidateCertificateOk

`func (o *UpdateMonitorRequestPatchHttpConfiguration) GetValidateCertificateOk() (*bool, bool)`

GetValidateCertificateOk returns a tuple with the ValidateCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateCertificate

`func (o *UpdateMonitorRequestPatchHttpConfiguration) SetValidateCertificate(v bool)`

SetValidateCertificate sets ValidateCertificate field to given value.

### HasValidateCertificate

`func (o *UpdateMonitorRequestPatchHttpConfiguration) HasValidateCertificate() bool`

HasValidateCertificate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


