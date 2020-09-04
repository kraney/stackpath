# PurgeContentRequestPurgeSelector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SelectorType** | Pointer to [**PurgeContentRequestPurgeSelectorType**](PurgeContentRequestPurgeSelectorType.md) |  | [optional] [default to "HEADER"]
**SelectorName** | Pointer to **string** | The name of the type of content to purge  For example, the name of the HTTP response header. Names are case sensitive. | [optional] 
**SelectorValue** | Pointer to **string** | The value of the content to purge  For example, the value of the HTTP response header. Values are case sensitive and may be wild-carded, but cannot match a \&quot;/\&quot;. | [optional] 
**SelectorValueDelimiter** | Pointer to **string** | The delimiter to separate multiple values with  Defaults to \&quot;,\&quot;. | [optional] 

## Methods

### NewPurgeContentRequestPurgeSelector

`func NewPurgeContentRequestPurgeSelector() *PurgeContentRequestPurgeSelector`

NewPurgeContentRequestPurgeSelector instantiates a new PurgeContentRequestPurgeSelector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurgeContentRequestPurgeSelectorWithDefaults

`func NewPurgeContentRequestPurgeSelectorWithDefaults() *PurgeContentRequestPurgeSelector`

NewPurgeContentRequestPurgeSelectorWithDefaults instantiates a new PurgeContentRequestPurgeSelector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelectorType

`func (o *PurgeContentRequestPurgeSelector) GetSelectorType() PurgeContentRequestPurgeSelectorType`

GetSelectorType returns the SelectorType field if non-nil, zero value otherwise.

### GetSelectorTypeOk

`func (o *PurgeContentRequestPurgeSelector) GetSelectorTypeOk() (*PurgeContentRequestPurgeSelectorType, bool)`

GetSelectorTypeOk returns a tuple with the SelectorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectorType

`func (o *PurgeContentRequestPurgeSelector) SetSelectorType(v PurgeContentRequestPurgeSelectorType)`

SetSelectorType sets SelectorType field to given value.

### HasSelectorType

`func (o *PurgeContentRequestPurgeSelector) HasSelectorType() bool`

HasSelectorType returns a boolean if a field has been set.

### GetSelectorName

`func (o *PurgeContentRequestPurgeSelector) GetSelectorName() string`

GetSelectorName returns the SelectorName field if non-nil, zero value otherwise.

### GetSelectorNameOk

`func (o *PurgeContentRequestPurgeSelector) GetSelectorNameOk() (*string, bool)`

GetSelectorNameOk returns a tuple with the SelectorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectorName

`func (o *PurgeContentRequestPurgeSelector) SetSelectorName(v string)`

SetSelectorName sets SelectorName field to given value.

### HasSelectorName

`func (o *PurgeContentRequestPurgeSelector) HasSelectorName() bool`

HasSelectorName returns a boolean if a field has been set.

### GetSelectorValue

`func (o *PurgeContentRequestPurgeSelector) GetSelectorValue() string`

GetSelectorValue returns the SelectorValue field if non-nil, zero value otherwise.

### GetSelectorValueOk

`func (o *PurgeContentRequestPurgeSelector) GetSelectorValueOk() (*string, bool)`

GetSelectorValueOk returns a tuple with the SelectorValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectorValue

`func (o *PurgeContentRequestPurgeSelector) SetSelectorValue(v string)`

SetSelectorValue sets SelectorValue field to given value.

### HasSelectorValue

`func (o *PurgeContentRequestPurgeSelector) HasSelectorValue() bool`

HasSelectorValue returns a boolean if a field has been set.

### GetSelectorValueDelimiter

`func (o *PurgeContentRequestPurgeSelector) GetSelectorValueDelimiter() string`

GetSelectorValueDelimiter returns the SelectorValueDelimiter field if non-nil, zero value otherwise.

### GetSelectorValueDelimiterOk

`func (o *PurgeContentRequestPurgeSelector) GetSelectorValueDelimiterOk() (*string, bool)`

GetSelectorValueDelimiterOk returns a tuple with the SelectorValueDelimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectorValueDelimiter

`func (o *PurgeContentRequestPurgeSelector) SetSelectorValueDelimiter(v string)`

SetSelectorValueDelimiter sets SelectorValueDelimiter field to given value.

### HasSelectorValueDelimiter

`func (o *PurgeContentRequestPurgeSelector) HasSelectorValueDelimiter() bool`

HasSelectorValueDelimiter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


