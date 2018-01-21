# ProvisionPhoneNumberRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtensionId** | [***ExtensionInfoProvisionPhoneNumbers**](ExtensionInfoProvisionPhoneNumbers.md) | Information on extension which the phone number is added to, only for provisioning extension-level numbers | [optional] [default to null]
**PhoneNumber** | **string** | Phone number to purchase returned in E.164 (11-digits) format | [optional] [default to null]
**ReservationId** | **string** | Internal identifier of phone number reservation; encoded data including reservation type (by brand, by account, by session), particular brand/account/session data, and reservation date and time | [optional] [default to null]
**Label** | **string** | Custom user name of a phone number, if any. Supported for numbers assigned to Auto-Receptionist, with usage type &#39;CompanyNumber | [optional] [default to null]
**UsageType** | **string** | Usage type of a phone number. The default value is &#39;DirectNumber&#39; | [optional] [default to null]
**Type_** | **string** | Type of a phone number | [optional] [default to null]
**VanityPattern** | **string** |  Vanity pattern that was used to find this number. It should be passed as if it was returned from the Number Lookup call | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


