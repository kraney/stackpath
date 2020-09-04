# DeliveryOriginAuthentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Basic** | Pointer to [**DeliveryBasicAuthentication**](deliveryBasicAuthentication.md) |  | [optional] 
**S3** | Pointer to [**DeliveryS3Authentication**](deliveryS3Authentication.md) |  | [optional] 

## Methods

### NewDeliveryOriginAuthentication

`func NewDeliveryOriginAuthentication() *DeliveryOriginAuthentication`

NewDeliveryOriginAuthentication instantiates a new DeliveryOriginAuthentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryOriginAuthenticationWithDefaults

`func NewDeliveryOriginAuthenticationWithDefaults() *DeliveryOriginAuthentication`

NewDeliveryOriginAuthenticationWithDefaults instantiates a new DeliveryOriginAuthentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasic

`func (o *DeliveryOriginAuthentication) GetBasic() DeliveryBasicAuthentication`

GetBasic returns the Basic field if non-nil, zero value otherwise.

### GetBasicOk

`func (o *DeliveryOriginAuthentication) GetBasicOk() (*DeliveryBasicAuthentication, bool)`

GetBasicOk returns a tuple with the Basic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasic

`func (o *DeliveryOriginAuthentication) SetBasic(v DeliveryBasicAuthentication)`

SetBasic sets Basic field to given value.

### HasBasic

`func (o *DeliveryOriginAuthentication) HasBasic() bool`

HasBasic returns a boolean if a field has been set.

### GetS3

`func (o *DeliveryOriginAuthentication) GetS3() DeliveryS3Authentication`

GetS3 returns the S3 field if non-nil, zero value otherwise.

### GetS3Ok

`func (o *DeliveryOriginAuthentication) GetS3Ok() (*DeliveryS3Authentication, bool)`

GetS3Ok returns a tuple with the S3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3

`func (o *DeliveryOriginAuthentication) SetS3(v DeliveryS3Authentication)`

SetS3 sets S3 field to given value.

### HasS3

`func (o *DeliveryOriginAuthentication) HasS3() bool`

HasS3 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


