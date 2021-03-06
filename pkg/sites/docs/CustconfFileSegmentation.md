# CustconfFileSegmentation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | This is used by the API to perform conflict checking | [optional] 
**Enabled** | Pointer to **bool** | Flag for enabling the File Segmentation Feature. | [optional] 

## Methods

### NewCustconfFileSegmentation

`func NewCustconfFileSegmentation() *CustconfFileSegmentation`

NewCustconfFileSegmentation instantiates a new CustconfFileSegmentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustconfFileSegmentationWithDefaults

`func NewCustconfFileSegmentationWithDefaults() *CustconfFileSegmentation`

NewCustconfFileSegmentationWithDefaults instantiates a new CustconfFileSegmentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustconfFileSegmentation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustconfFileSegmentation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustconfFileSegmentation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustconfFileSegmentation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEnabled

`func (o *CustconfFileSegmentation) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustconfFileSegmentation) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustconfFileSegmentation) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CustconfFileSegmentation) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


