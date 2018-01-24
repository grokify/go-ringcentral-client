# InboundFaxesInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotifyByEmail** | **bool** | Email notification flag | [optional] [default to null]
**NotifyBySms** | **bool** | SMS notification flag | [optional] [default to null]
**AdvancedEmailAddresses** | **[]string** | List of recipient email addresses for inbound fax notifications. Returned if specified, in both modes (advanced/basic). Applied in advanced mode only | [optional] [default to null]
**AdvancedSmsEmailAddresses** | **[]string** | List of recipient phone numbers for inbound fax notifications. Returned if specified, in both modes (advanced/basic). Applied in advanced mode only | [optional] [default to null]
**IncludeAttachment** | **bool** | Indicates whether fax should be attached to email | [optional] [default to null]
**MarkAsRead** | **bool** | Indicates whether email should be automatically marked as read | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


