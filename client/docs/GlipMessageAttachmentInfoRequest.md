# GlipMessageAttachmentInfoRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** | Type of an attachment | [optional] [default to null]
**Fallback** | **string** | A string of default text that will be rendered in the case that the client does not support Interactive Messages | [optional] [default to null]
**Color** | **string** | A Hex color code that determines the color of the side border of the Interactive Message | [optional] [default to null]
**Intro** | **string** | A pretext to the message | [optional] [default to null]
**Author** | [***GlipMessageAttachmentAuthorInfo**](GlipMessageAttachmentAuthorInfo.md) | Information about the author | [optional] [default to null]
**Title** | **string** | Message title | [optional] [default to null]
**Text** | **string** | A large string field (up to 1000 chars) to be displayed as the body of a message (Supports GlipDown) | [optional] [default to null]
**ImageUri** | **string** | url used to display a single image at the bottom of a message | [optional] [default to null]
**ThumbnailUri** | **string** | url used to display a thumbnail to the right of a message (82x82) | [optional] [default to null]
**Fields** | [**[]GlipMessageAttachmentFieldsInfo**](GlipMessageAttachmentFieldsInfo.md) | Information on navigation | [optional] [default to null]
**Footnote** | [***GlipMessageAttachmentFootnoteInfo**](GlipMessageAttachmentFootnoteInfo.md) | Message Footer | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


