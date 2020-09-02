# CustconfAuthHttpBasic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | This is used by the API to perform conflict checking | [optional] 
**BindingPoint** | **string** | This is a URL to a resource on the authentication server responsible for authentication of users. | [optional] 
**Realm** | **string** | This is the authentication realm given back to the user on requests which do not contain the authentication credentials. Browsers typically display this value to the user when the login credentials are requested. | [optional] 
**Ttl** | **float32** | This is the number of seconds the authentication session will be cached by the browsers. After this time, browsers will be asked to present the user credentials again for re-authentication. | [optional] 
**Enabled** | **bool** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


