# IdentityGetIAMPolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | The ID of the StackPath account the policy is on | [optional] 
**Policy** | Pointer to [**IamPolicy**](iamPolicy.md) |  | [optional] 

## Methods

### NewIdentityGetIAMPolicyResponse

`func NewIdentityGetIAMPolicyResponse() *IdentityGetIAMPolicyResponse`

NewIdentityGetIAMPolicyResponse instantiates a new IdentityGetIAMPolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityGetIAMPolicyResponseWithDefaults

`func NewIdentityGetIAMPolicyResponseWithDefaults() *IdentityGetIAMPolicyResponse`

NewIdentityGetIAMPolicyResponseWithDefaults instantiates a new IdentityGetIAMPolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *IdentityGetIAMPolicyResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *IdentityGetIAMPolicyResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *IdentityGetIAMPolicyResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *IdentityGetIAMPolicyResponse) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetPolicy

`func (o *IdentityGetIAMPolicyResponse) GetPolicy() IamPolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *IdentityGetIAMPolicyResponse) GetPolicyOk() (*IamPolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *IdentityGetIAMPolicyResponse) SetPolicy(v IamPolicy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *IdentityGetIAMPolicyResponse) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


