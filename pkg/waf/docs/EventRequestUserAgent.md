# EventRequestUserAgent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientName** | Pointer to **string** | The name of the requesting client, typically the name of the requesting web browser | [optional] 
**ClientType** | Pointer to **string** | The requesting client&#39;s type, such as \&quot;major browser\&quot; or \&quot;mobile app\&quot; | [optional] 
**Device** | Pointer to **string** | The brand name of the device making the request, such as \&quot;iphone\&quot; or \&quot;playstation\&quot; | [optional] 
**Os** | Pointer to **string** | The operating system of the device making the request | [optional] 
**Raw** | Pointer to **string** | The full user agent string | [optional] 

## Methods

### NewEventRequestUserAgent

`func NewEventRequestUserAgent() *EventRequestUserAgent`

NewEventRequestUserAgent instantiates a new EventRequestUserAgent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventRequestUserAgentWithDefaults

`func NewEventRequestUserAgentWithDefaults() *EventRequestUserAgent`

NewEventRequestUserAgentWithDefaults instantiates a new EventRequestUserAgent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientName

`func (o *EventRequestUserAgent) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *EventRequestUserAgent) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *EventRequestUserAgent) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *EventRequestUserAgent) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetClientType

`func (o *EventRequestUserAgent) GetClientType() string`

GetClientType returns the ClientType field if non-nil, zero value otherwise.

### GetClientTypeOk

`func (o *EventRequestUserAgent) GetClientTypeOk() (*string, bool)`

GetClientTypeOk returns a tuple with the ClientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientType

`func (o *EventRequestUserAgent) SetClientType(v string)`

SetClientType sets ClientType field to given value.

### HasClientType

`func (o *EventRequestUserAgent) HasClientType() bool`

HasClientType returns a boolean if a field has been set.

### GetDevice

`func (o *EventRequestUserAgent) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *EventRequestUserAgent) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *EventRequestUserAgent) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *EventRequestUserAgent) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetOs

`func (o *EventRequestUserAgent) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *EventRequestUserAgent) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *EventRequestUserAgent) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *EventRequestUserAgent) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetRaw

`func (o *EventRequestUserAgent) GetRaw() string`

GetRaw returns the Raw field if non-nil, zero value otherwise.

### GetRawOk

`func (o *EventRequestUserAgent) GetRawOk() (*string, bool)`

GetRawOk returns a tuple with the Raw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw

`func (o *EventRequestUserAgent) SetRaw(v string)`

SetRaw sets Raw field to given value.

### HasRaw

`func (o *EventRequestUserAgent) HasRaw() bool`

HasRaw returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


