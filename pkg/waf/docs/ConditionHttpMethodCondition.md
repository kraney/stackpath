# ConditionHttpMethodCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HttpMethod** | Pointer to [**WafHttpMethod**](wafHttpMethod.md) |  | [optional] [default to "METHOD_UNSPECIFIED"]

## Methods

### NewConditionHttpMethodCondition

`func NewConditionHttpMethodCondition() *ConditionHttpMethodCondition`

NewConditionHttpMethodCondition instantiates a new ConditionHttpMethodCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionHttpMethodConditionWithDefaults

`func NewConditionHttpMethodConditionWithDefaults() *ConditionHttpMethodCondition`

NewConditionHttpMethodConditionWithDefaults instantiates a new ConditionHttpMethodCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttpMethod

`func (o *ConditionHttpMethodCondition) GetHttpMethod() WafHttpMethod`

GetHttpMethod returns the HttpMethod field if non-nil, zero value otherwise.

### GetHttpMethodOk

`func (o *ConditionHttpMethodCondition) GetHttpMethodOk() (*WafHttpMethod, bool)`

GetHttpMethodOk returns a tuple with the HttpMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpMethod

`func (o *ConditionHttpMethodCondition) SetHttpMethod(v WafHttpMethod)`

SetHttpMethod sets HttpMethod field to given value.

### HasHttpMethod

`func (o *ConditionHttpMethodCondition) HasHttpMethod() bool`

HasHttpMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


