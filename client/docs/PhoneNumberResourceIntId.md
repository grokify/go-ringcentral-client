# PhoneNumberResourceIntId

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Internal identifier of a phone number | [optional] [default to null]
**Country** | [***CountryResource**](CountryResource.md) | Brief information on a phone number country | [optional] [default to null]
**Extension** | [***ExtensionResource**](ExtensionResource.md) | Information on an extension to which the phone number is assigned | [optional] [default to null]
**Label** | **string** | Custom user name of a phone number, if any. Supported for numbers assigned to Auto-Receptionist, with usage type &#39;CompanyNumber&#39; | [optional] [default to null]
**Location** | **string** | Location (City, State). Filled for local US numbers | [optional] [default to null]
**PaymentType** | **string** | Payment type. &#39;External&#39; is returned for forwarded numbers which are not terminated in the RingCentral phone system &#x3D; [&#39;External&#39;, &#39;TollFree&#39;, &#39;Local&#39;], | [optional] [default to null]
**PhoneNumber** | **string** | Phone number | [optional] [default to null]
**Status** | **string** | Status of a phone number. If the value is &#39;Normal&#39;, the phone number is ready to be used. Otherwise it is an external number not yet ported to RingCentral , | [optional] [default to null]
**UsageType** | **string** |  | [optional] [default to null]
**Type_** | **string** | Type of a phone number | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


