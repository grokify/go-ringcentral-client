# CreateAccountRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MainNumber** | **string** | Main account VoIP phone number, either Local or Toll-Free. Cannot be Fax Only. Obtained via lookup/reserve API | [optional] 
**Operator** | [**GetExtensionInfoResponse**](GetExtensionInfoResponse.md) |  | [optional] 
**PartnerId** | **string** | Partner identifier for this account | [optional] 
**PromotionCode** | **string** | Promotion code to calculate a discount | [optional] 
**ReservationId** | **string** | Internal identifier of phone number reservation; encoded data including reservation type (by brand, by account, by session), particular brand/account/session data, and reservation date and time | [optional] 
**ServiceInfo** | [**GetServiceInfoResponse**](GetServiceInfoResponse.md) |  | [optional] 
**Status** | **string** | The status with which an account is created. The default value is &#39;Initial&#39; | [optional] 
**SignupInfo** | [**AccountSignupInfoRequest**](AccountSignupInfoRequest.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


