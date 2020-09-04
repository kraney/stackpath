# CustconfAwsSignedS3PostV4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | This is used by the API to perform conflict checking | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**BucketIdentifier** | Pointer to **string** |  | [optional] 
**AccessKeyId** | Pointer to **string** |  | [optional] 
**SecretAccessKey** | Pointer to **string** |  | [optional] 
**AwsRegion** | Pointer to **string** |  | [optional] 
**AwsService** | Pointer to **string** |  | [optional] 
**ExpireTimeSeconds** | Pointer to **float32** |  | [optional] 
**SignedHeaders** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**AuthenticationType** | Pointer to [**CustconfAwsSignedS3PostV4AuthenticationTypeEnumWrapperValue**](custconfAwsSignedS3PostV4AuthenticationTypeEnumWrapperValue.md) |  | [optional] [default to "UNKNOWN"]
**HeaderFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**MethodFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 
**PathFilter** | Pointer to **string** | String of values delimited by a &#39;,&#39; character. | [optional] 

## Methods

### NewCustconfAwsSignedS3PostV4

`func NewCustconfAwsSignedS3PostV4() *CustconfAwsSignedS3PostV4`

NewCustconfAwsSignedS3PostV4 instantiates a new CustconfAwsSignedS3PostV4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustconfAwsSignedS3PostV4WithDefaults

`func NewCustconfAwsSignedS3PostV4WithDefaults() *CustconfAwsSignedS3PostV4`

NewCustconfAwsSignedS3PostV4WithDefaults instantiates a new CustconfAwsSignedS3PostV4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustconfAwsSignedS3PostV4) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustconfAwsSignedS3PostV4) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustconfAwsSignedS3PostV4) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustconfAwsSignedS3PostV4) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEnabled

`func (o *CustconfAwsSignedS3PostV4) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustconfAwsSignedS3PostV4) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustconfAwsSignedS3PostV4) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CustconfAwsSignedS3PostV4) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetBucketIdentifier

`func (o *CustconfAwsSignedS3PostV4) GetBucketIdentifier() string`

GetBucketIdentifier returns the BucketIdentifier field if non-nil, zero value otherwise.

### GetBucketIdentifierOk

`func (o *CustconfAwsSignedS3PostV4) GetBucketIdentifierOk() (*string, bool)`

GetBucketIdentifierOk returns a tuple with the BucketIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketIdentifier

`func (o *CustconfAwsSignedS3PostV4) SetBucketIdentifier(v string)`

SetBucketIdentifier sets BucketIdentifier field to given value.

### HasBucketIdentifier

`func (o *CustconfAwsSignedS3PostV4) HasBucketIdentifier() bool`

HasBucketIdentifier returns a boolean if a field has been set.

### GetAccessKeyId

`func (o *CustconfAwsSignedS3PostV4) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *CustconfAwsSignedS3PostV4) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *CustconfAwsSignedS3PostV4) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.

### HasAccessKeyId

`func (o *CustconfAwsSignedS3PostV4) HasAccessKeyId() bool`

HasAccessKeyId returns a boolean if a field has been set.

### GetSecretAccessKey

`func (o *CustconfAwsSignedS3PostV4) GetSecretAccessKey() string`

GetSecretAccessKey returns the SecretAccessKey field if non-nil, zero value otherwise.

### GetSecretAccessKeyOk

`func (o *CustconfAwsSignedS3PostV4) GetSecretAccessKeyOk() (*string, bool)`

GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAccessKey

`func (o *CustconfAwsSignedS3PostV4) SetSecretAccessKey(v string)`

SetSecretAccessKey sets SecretAccessKey field to given value.

### HasSecretAccessKey

`func (o *CustconfAwsSignedS3PostV4) HasSecretAccessKey() bool`

HasSecretAccessKey returns a boolean if a field has been set.

### GetAwsRegion

`func (o *CustconfAwsSignedS3PostV4) GetAwsRegion() string`

GetAwsRegion returns the AwsRegion field if non-nil, zero value otherwise.

### GetAwsRegionOk

`func (o *CustconfAwsSignedS3PostV4) GetAwsRegionOk() (*string, bool)`

GetAwsRegionOk returns a tuple with the AwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegion

`func (o *CustconfAwsSignedS3PostV4) SetAwsRegion(v string)`

SetAwsRegion sets AwsRegion field to given value.

### HasAwsRegion

`func (o *CustconfAwsSignedS3PostV4) HasAwsRegion() bool`

HasAwsRegion returns a boolean if a field has been set.

### GetAwsService

`func (o *CustconfAwsSignedS3PostV4) GetAwsService() string`

GetAwsService returns the AwsService field if non-nil, zero value otherwise.

### GetAwsServiceOk

`func (o *CustconfAwsSignedS3PostV4) GetAwsServiceOk() (*string, bool)`

GetAwsServiceOk returns a tuple with the AwsService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsService

`func (o *CustconfAwsSignedS3PostV4) SetAwsService(v string)`

SetAwsService sets AwsService field to given value.

### HasAwsService

`func (o *CustconfAwsSignedS3PostV4) HasAwsService() bool`

HasAwsService returns a boolean if a field has been set.

### GetExpireTimeSeconds

`func (o *CustconfAwsSignedS3PostV4) GetExpireTimeSeconds() float32`

GetExpireTimeSeconds returns the ExpireTimeSeconds field if non-nil, zero value otherwise.

### GetExpireTimeSecondsOk

`func (o *CustconfAwsSignedS3PostV4) GetExpireTimeSecondsOk() (*float32, bool)`

GetExpireTimeSecondsOk returns a tuple with the ExpireTimeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireTimeSeconds

`func (o *CustconfAwsSignedS3PostV4) SetExpireTimeSeconds(v float32)`

SetExpireTimeSeconds sets ExpireTimeSeconds field to given value.

### HasExpireTimeSeconds

`func (o *CustconfAwsSignedS3PostV4) HasExpireTimeSeconds() bool`

HasExpireTimeSeconds returns a boolean if a field has been set.

### GetSignedHeaders

`func (o *CustconfAwsSignedS3PostV4) GetSignedHeaders() string`

GetSignedHeaders returns the SignedHeaders field if non-nil, zero value otherwise.

### GetSignedHeadersOk

`func (o *CustconfAwsSignedS3PostV4) GetSignedHeadersOk() (*string, bool)`

GetSignedHeadersOk returns a tuple with the SignedHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedHeaders

`func (o *CustconfAwsSignedS3PostV4) SetSignedHeaders(v string)`

SetSignedHeaders sets SignedHeaders field to given value.

### HasSignedHeaders

`func (o *CustconfAwsSignedS3PostV4) HasSignedHeaders() bool`

HasSignedHeaders returns a boolean if a field has been set.

### GetAuthenticationType

`func (o *CustconfAwsSignedS3PostV4) GetAuthenticationType() CustconfAwsSignedS3PostV4AuthenticationTypeEnumWrapperValue`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *CustconfAwsSignedS3PostV4) GetAuthenticationTypeOk() (*CustconfAwsSignedS3PostV4AuthenticationTypeEnumWrapperValue, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *CustconfAwsSignedS3PostV4) SetAuthenticationType(v CustconfAwsSignedS3PostV4AuthenticationTypeEnumWrapperValue)`

SetAuthenticationType sets AuthenticationType field to given value.

### HasAuthenticationType

`func (o *CustconfAwsSignedS3PostV4) HasAuthenticationType() bool`

HasAuthenticationType returns a boolean if a field has been set.

### GetHeaderFilter

`func (o *CustconfAwsSignedS3PostV4) GetHeaderFilter() string`

GetHeaderFilter returns the HeaderFilter field if non-nil, zero value otherwise.

### GetHeaderFilterOk

`func (o *CustconfAwsSignedS3PostV4) GetHeaderFilterOk() (*string, bool)`

GetHeaderFilterOk returns a tuple with the HeaderFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderFilter

`func (o *CustconfAwsSignedS3PostV4) SetHeaderFilter(v string)`

SetHeaderFilter sets HeaderFilter field to given value.

### HasHeaderFilter

`func (o *CustconfAwsSignedS3PostV4) HasHeaderFilter() bool`

HasHeaderFilter returns a boolean if a field has been set.

### GetMethodFilter

`func (o *CustconfAwsSignedS3PostV4) GetMethodFilter() string`

GetMethodFilter returns the MethodFilter field if non-nil, zero value otherwise.

### GetMethodFilterOk

`func (o *CustconfAwsSignedS3PostV4) GetMethodFilterOk() (*string, bool)`

GetMethodFilterOk returns a tuple with the MethodFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodFilter

`func (o *CustconfAwsSignedS3PostV4) SetMethodFilter(v string)`

SetMethodFilter sets MethodFilter field to given value.

### HasMethodFilter

`func (o *CustconfAwsSignedS3PostV4) HasMethodFilter() bool`

HasMethodFilter returns a boolean if a field has been set.

### GetPathFilter

`func (o *CustconfAwsSignedS3PostV4) GetPathFilter() string`

GetPathFilter returns the PathFilter field if non-nil, zero value otherwise.

### GetPathFilterOk

`func (o *CustconfAwsSignedS3PostV4) GetPathFilterOk() (*string, bool)`

GetPathFilterOk returns a tuple with the PathFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathFilter

`func (o *CustconfAwsSignedS3PostV4) SetPathFilter(v string)`

SetPathFilter sets PathFilter field to given value.

### HasPathFilter

`func (o *CustconfAwsSignedS3PostV4) HasPathFilter() bool`

HasPathFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


