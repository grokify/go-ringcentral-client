# UserUpdateRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** | Status of a user | [optional] [default to null]
**Addresses** | [**[]AddressInfo**](AddressInfo.md) | User addresses | [optional] [default to null]
**Emails** | [**[]EmailInfo**](EmailInfo.md) | User email addresses | [default to null]
**ExternalId** | **string** | External identifier of a user | [optional] [default to null]
**Id** | **string** | Internal identifier of a user | [optional] [default to null]
**Name** | [***NameInfo**](NameInfo.md) | User name | [default to null]
**PhoneNumbers** | [**[]PhoneNumberInfoRequest**](PhoneNumberInfoRequest.md) | User phone numbers | [optional] [default to null]
**Photos** | [**[]PhotoInfo**](PhotoInfo.md) |  | [optional] [default to null]
**Schemas** | **[]string** | Specification links | [default to null]
**Urnietfparamsscimschemasextensionenterprise20User** | [***EnterpriseUser**](EnterpriseUser.md) |  | [optional] [default to null]
**UserName** | **string** | User mailbox. Must be same as work type email address | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


