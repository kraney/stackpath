# CdnDnsRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the network node to which a zone resource record pertains  Use the value \&quot;@\&quot; to denote current root domain name. | [optional] 
**Type** | Pointer to **string** | A zone record&#39;s type  Zone record types describe the zone record&#39;s behavior. For instance, a zone record&#39;s type can say that the record is a name to IP address value, a name alias, or which mail exchanger is responsible for the domain. See https://support.stackpath.com/hc/en-us/articles/360001085563-What-DNS-record-types-does-StackPath-support for more information. | [optional] 
**Class** | Pointer to **string** | A zone record&#39;s class code  This is typically \&quot;IN\&quot; for Internet related resource records. | [optional] 
**Ttl** | Pointer to **int32** | A zone record&#39;s time to live  A record&#39;s TTL is the number of seconds that the record should be cached by DNS resolvers. Use lower TTL values if you expect zone records to change often. Use higher TTL values for records that won&#39;t change to prevent extra DNS lookups by clients. | [optional] 
**Data** | Pointer to **string** | A zone record&#39;s value  Expected data formats can vary depending on the zone record&#39;s type. | [optional] 

## Methods

### NewCdnDnsRecord

`func NewCdnDnsRecord() *CdnDnsRecord`

NewCdnDnsRecord instantiates a new CdnDnsRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdnDnsRecordWithDefaults

`func NewCdnDnsRecordWithDefaults() *CdnDnsRecord`

NewCdnDnsRecordWithDefaults instantiates a new CdnDnsRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CdnDnsRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CdnDnsRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CdnDnsRecord) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CdnDnsRecord) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *CdnDnsRecord) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CdnDnsRecord) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CdnDnsRecord) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CdnDnsRecord) HasType() bool`

HasType returns a boolean if a field has been set.

### GetClass

`func (o *CdnDnsRecord) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *CdnDnsRecord) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *CdnDnsRecord) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *CdnDnsRecord) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetTtl

`func (o *CdnDnsRecord) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *CdnDnsRecord) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *CdnDnsRecord) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *CdnDnsRecord) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetData

`func (o *CdnDnsRecord) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CdnDnsRecord) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CdnDnsRecord) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *CdnDnsRecord) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


