# DeliveryBasicAuthentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** | The username to be used when authenticating with the origin. | [optional] 
**Password** | Pointer to **string** | The password to be used when authenticating with the origin. | [optional] 

## Methods

### NewDeliveryBasicAuthentication

`func NewDeliveryBasicAuthentication() *DeliveryBasicAuthentication`

NewDeliveryBasicAuthentication instantiates a new DeliveryBasicAuthentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryBasicAuthenticationWithDefaults

`func NewDeliveryBasicAuthenticationWithDefaults() *DeliveryBasicAuthentication`

NewDeliveryBasicAuthenticationWithDefaults instantiates a new DeliveryBasicAuthentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *DeliveryBasicAuthentication) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *DeliveryBasicAuthentication) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *DeliveryBasicAuthentication) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *DeliveryBasicAuthentication) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *DeliveryBasicAuthentication) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *DeliveryBasicAuthentication) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *DeliveryBasicAuthentication) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *DeliveryBasicAuthentication) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


