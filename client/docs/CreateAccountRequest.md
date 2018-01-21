# CreateAccountRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MainNumber** | **string** | Main account VoIP phone number, either Local or Toll-Free. Cannot be Fax Only. Obtained via lookup/reserve API | [optional] [default to null]
**Operator** | [***GetExtensionInfoResponse**](GetExtensionInfoResponse.md) | Operator&#39;s extension information. This extension will receive all calls and messages intended for the operator | [optional] [default to null]
**PartnerId** | **string** | Partner identifier for this account | [optional] [default to null]
**PromotionCode** | **string** | Promotion code to calculate a discount | [optional] [default to null]
**ReservationId** | **string** | Internal identifier of phone number reservation; encoded data including reservation type (by brand, by account, by session), particular brand/account/session data, and reservation date and time | [optional] [default to null]
**ServiceInfo** | [***GetServiceInfoResponse**](GetServiceInfoResponse.md) | Account service information, brand identifier and service plan | [optional] [default to null]
**Status** | **string** | The status with which an account is created. The default value is &#39;Initial&#39; | [optional] [default to null]
**SignupInfo** | [***AccountSignupInfoRequest**](AccountSignupInfoRequest.md) | Account sign up data | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


