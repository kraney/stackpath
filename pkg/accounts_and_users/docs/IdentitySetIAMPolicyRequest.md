# IdentitySetIAMPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | Pointer to [**IamPolicy**](iamPolicy.md) |  | [optional] 

## Methods

### NewIdentitySetIAMPolicyRequest

`func NewIdentitySetIAMPolicyRequest() *IdentitySetIAMPolicyRequest`

NewIdentitySetIAMPolicyRequest instantiates a new IdentitySetIAMPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentitySetIAMPolicyRequestWithDefaults

`func NewIdentitySetIAMPolicyRequestWithDefaults() *IdentitySetIAMPolicyRequest`

NewIdentitySetIAMPolicyRequestWithDefaults instantiates a new IdentitySetIAMPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicy

`func (o *IdentitySetIAMPolicyRequest) GetPolicy() IamPolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *IdentitySetIAMPolicyRequest) GetPolicyOk() (*IamPolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *IdentitySetIAMPolicyRequest) SetPolicy(v IamPolicy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *IdentitySetIAMPolicyRequest) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


