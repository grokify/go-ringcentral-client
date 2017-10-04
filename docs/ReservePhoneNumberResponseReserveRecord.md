# ReservePhoneNumberResponseReserveRecord

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumber** | **string** | Phone number in E.164 format without a &#39;+&#39; | [optional] [default to null]
**FormattedNumber** | **string** | Domestic format of a phone number | [optional] [default to null]
**ReservedTill** | [**time.Time**](time.Time.md) | The datetime up to which the number is reserved in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z. No value means that number is not reserved anymore | [optional] [default to null]
**ReservationId** | **string** | Internal identifier of phone number reservation; encoded data including reservation type (by brand, by account, by session), particular brand/account/session data, and reservation date and time | [optional] [default to null]
**Status** | **string** | Phone number status | [optional] [default to null]
**Error_** | **string** | The error code in case of reservation/un-reservation failure | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


