# StackSetIAMPolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StackId** | Pointer to **string** | The ID of the StackPath stack the policy is on | [optional] 
**Policy** | Pointer to [**IamPolicy**](iamPolicy.md) |  | [optional] 

## Methods

### NewStackSetIAMPolicyResponse

`func NewStackSetIAMPolicyResponse() *StackSetIAMPolicyResponse`

NewStackSetIAMPolicyResponse instantiates a new StackSetIAMPolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackSetIAMPolicyResponseWithDefaults

`func NewStackSetIAMPolicyResponseWithDefaults() *StackSetIAMPolicyResponse`

NewStackSetIAMPolicyResponseWithDefaults instantiates a new StackSetIAMPolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStackId

`func (o *StackSetIAMPolicyResponse) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *StackSetIAMPolicyResponse) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *StackSetIAMPolicyResponse) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *StackSetIAMPolicyResponse) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetPolicy

`func (o *StackSetIAMPolicyResponse) GetPolicy() IamPolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *StackSetIAMPolicyResponse) GetPolicyOk() (*IamPolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *StackSetIAMPolicyResponse) SetPolicy(v IamPolicy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *StackSetIAMPolicyResponse) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


