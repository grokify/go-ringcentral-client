# NotificationSettingsUpdateRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAddresses** | **[]string** | List of notification recipient email addresses | [optional] [default to null]
**SmsEmailAddresses** | **[]string** | List of notification recipient email addresses | [optional] [default to null]
**AdvancedMode** | **bool** | Specifies notifications settings mode. If &#39;True&#39; then advanced mode is on, it allows using different emails and/or phone numbers for each notification type. If &#39;False&#39; then basic mode is on. Advanced mode settings are returned in both modes, if specified once, but if basic mode is switched on, they are not applied | [optional] [default to null]
**Voicemails** | [***VoicemailsInfo**](VoicemailsInfo.md) |  | [optional] [default to null]
**InboundFaxes** | [***InboundFaxesInfo**](InboundFaxesInfo.md) |  | [optional] [default to null]
**OutboundFaxes** | [***OutboundFaxesInfo**](OutboundFaxesInfo.md) |  | [optional] [default to null]
**InboundTexts** | [***InboundTextsInfo**](InboundTextsInfo.md) |  | [optional] [default to null]
**MissedCalls** | [***MissedCallsInfo**](MissedCallsInfo.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


