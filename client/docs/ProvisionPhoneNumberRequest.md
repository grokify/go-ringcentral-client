# ProvisionPhoneNumberRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtensionId** | [**ExtensionInfoProvisionPhoneNumbers**](ExtensionInfoProvisionPhoneNumbers.md) |  | [optional] 
**PhoneNumber** | **string** | Phone number to purchase returned in E.164 (11-digits) format | [optional] 
**ReservationId** | **string** | Internal identifier of phone number reservation; encoded data including reservation type (by brand, by account, by session), particular brand/account/session data, and reservation date and time | [optional] 
**Label** | **string** | Custom user name of a phone number, if any. Supported for numbers assigned to Auto-Receptionist, with usage type &#39;CompanyNumber | [optional] 
**UsageType** | **string** | Usage type of a phone number. The default value is &#39;DirectNumber&#39; | [optional] 
**Type** | **string** | Type of a phone number | [optional] 
**VanityPattern** | **string** |  Vanity pattern that was used to find this number. It should be passed as if it was returned from the Number Lookup call | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


