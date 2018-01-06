# NumberPoolPhoneNumberRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumber** | **string** | Phone number in E.164 format | [optional] [default to null]
**ReservedTill** | [**time.Time**](time.Time.md) | Datetime up to which the number is reserved in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z. If it is omitted or explicitly set to &#39;null&#39;, the number will be un-reserved if it was reserved previously by the same session. Also the values &#39;Min&#39; and &#39;Max&#39; are supported. &#39;Min&#39; is 30 seconds; and &#39;Max&#39; is 30 days (for reservation by brand) and 20 minutes (for reservation by account/session) | [optional] [default to null]
**ReservationId** | **string** | Internal identifier of a phone number reservation; encoded data including reservation type (by brand, by account, by session), particular brand/account/session data, and reservation date and time | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


