# CdnPop

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | A StackPath POP&#39;s IATA formatted location code | [optional] 
**Name** | Pointer to **string** | A StackPath POP&#39;s name | [optional] 
**Latitude** | Pointer to **float32** | A StackPath POP&#39;s latitude coordinates | [optional] 
**Longitude** | Pointer to **float32** | A StackPath POP&#39;s longitude coordinates | [optional] 

## Methods

### NewCdnPop

`func NewCdnPop() *CdnPop`

NewCdnPop instantiates a new CdnPop object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdnPopWithDefaults

`func NewCdnPopWithDefaults() *CdnPop`

NewCdnPopWithDefaults instantiates a new CdnPop object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *CdnPop) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CdnPop) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CdnPop) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CdnPop) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *CdnPop) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CdnPop) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CdnPop) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CdnPop) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLatitude

`func (o *CdnPop) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *CdnPop) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *CdnPop) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *CdnPop) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *CdnPop) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *CdnPop) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *CdnPop) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *CdnPop) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


