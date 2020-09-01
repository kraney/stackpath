# StackGetIAMPolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StackId** | Pointer to **string** | The ID of the StackPath stack the policy is on | [optional] 
**Policy** | Pointer to [**IamPolicy**](iamPolicy.md) |  | [optional] 

## Methods

### NewStackGetIAMPolicyResponse

`func NewStackGetIAMPolicyResponse() *StackGetIAMPolicyResponse`

NewStackGetIAMPolicyResponse instantiates a new StackGetIAMPolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackGetIAMPolicyResponseWithDefaults

`func NewStackGetIAMPolicyResponseWithDefaults() *StackGetIAMPolicyResponse`

NewStackGetIAMPolicyResponseWithDefaults instantiates a new StackGetIAMPolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStackId

`func (o *StackGetIAMPolicyResponse) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *StackGetIAMPolicyResponse) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *StackGetIAMPolicyResponse) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *StackGetIAMPolicyResponse) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetPolicy

`func (o *StackGetIAMPolicyResponse) GetPolicy() IamPolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *StackGetIAMPolicyResponse) GetPolicyOk() (*IamPolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *StackGetIAMPolicyResponse) SetPolicy(v IamPolicy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *StackGetIAMPolicyResponse) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


