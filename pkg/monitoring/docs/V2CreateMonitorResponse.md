# V2CreateMonitorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Record** | Pointer to [**V2Monitor**](v2Monitor.md) |  | [optional] 

## Methods

### NewV2CreateMonitorResponse

`func NewV2CreateMonitorResponse() *V2CreateMonitorResponse`

NewV2CreateMonitorResponse instantiates a new V2CreateMonitorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2CreateMonitorResponseWithDefaults

`func NewV2CreateMonitorResponseWithDefaults() *V2CreateMonitorResponse`

NewV2CreateMonitorResponseWithDefaults instantiates a new V2CreateMonitorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecord

`func (o *V2CreateMonitorResponse) GetRecord() V2Monitor`

GetRecord returns the Record field if non-nil, zero value otherwise.

### GetRecordOk

`func (o *V2CreateMonitorResponse) GetRecordOk() (*V2Monitor, bool)`

GetRecordOk returns a tuple with the Record field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecord

`func (o *V2CreateMonitorResponse) SetRecord(v V2Monitor)`

SetRecord sets Record field to given value.

### HasRecord

`func (o *V2CreateMonitorResponse) HasRecord() bool`

HasRecord returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


