# NotificationSettings

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uri** | **string** | Canonical URI of notifications settings resource | [optional] 
**EmailAddresses** | **[]string** | List of notification recipient email addresses | [optional] 
**SmsEmailAddresses** | **[]string** | List of notification recipient email addresses | [optional] 
**AdvancedMode** | **bool** | Specifies notifications settings mode. If &#39;True&#39; then advanced mode is on, it allows using different emails and/or phone numbers for each notification type. If &#39;False&#39; then basic mode is on. Advanced mode settings are returned in both modes, if specified once, but if basic mode is switched on, they are not applied | [optional] 
**Voicemails** | [**VoicemailsInfo**](VoicemailsInfo.md) |  | [optional] 
**InboundFaxes** | [**InboundFaxesInfo**](InboundFaxesInfo.md) |  | [optional] 
**OutboundFaxes** | [**OutboundFaxesInfo**](OutboundFaxesInfo.md) |  | [optional] 
**InboundTexts** | [**InboundTextsInfo**](InboundTextsInfo.md) |  | [optional] 
**MissedCalls** | [**MissedCallsInfo**](MissedCallsInfo.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


