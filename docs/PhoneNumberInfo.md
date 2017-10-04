# PhoneNumberInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Internal identifier of a phone number | [optional] [default to null]
**Country** | [**CountryInfo**](CountryInfo.md) | Brief information on a phone number country | [optional] [default to null]
**Extension** | [**PhoneNumberInfoExtensionInfo**](PhoneNumberInfo.ExtensionInfo.md) | Information on the extension, to which the phone number is assigned. Returned only for the request of Account phone number list | [optional] [default to null]
**Features** | **[]string** | Indicates if this phone number is enabled to appear as CallerId and/or to send outbound SMS from it. Returned only for the request of Extension phone number list | [optional] [default to null]
**Location** | **string** | Location (City, State). Filled for local US numbers | [optional] [default to null]
**PaymentType** | **string** | Payment type. &#39;External&#39; is returned for forwarded numbers which are not terminated in the RingCentral phone system | [optional] [default to null]
**PhoneNumber** | **string** | Phone number | [optional] [default to null]
**Status** | **string** | Status of a phone number. If the value is &#39;Normal&#39;, the phone number is ready to be used. Otherwise it is an external number not yet ported to RingCentral | [optional] [default to null]
**Type_** | **string** | Phone number type | [optional] [default to null]
**UsageType** | **string** | Usage type of the phone number | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


