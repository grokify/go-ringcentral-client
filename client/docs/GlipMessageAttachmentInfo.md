# GlipMessageAttachmentInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Internal identifier of an attachment | [optional] 
**Type** | **string** | Type of an attachment | [optional] [default to Card]
**Fallback** | **string** | A string of default text that will be rendered in the case that the client does not support Interactive Messages | [optional] 
**Color** | **string** | A Hex color code that determines the color of the side border of the Interactive Message | [optional] 
**Intro** | **string** | A pretext to the message | [optional] 
**Author** | [**GlipMessageAttachmentAuthorInfo**](GlipMessageAttachmentAuthorInfo.md) |  | [optional] 
**Title** | **string** | Message title | [optional] 
**Text** | **string** | A large string field (up to 1000 chars) to be displayed as the body of a message (Supports GlipDown) | [optional] 
**ImageUri** | **string** | url used to display a single image at the bottom of a message | [optional] 
**ThumbnailUri** | **string** | url used to display a thumbnail to the right of a message (82x82) | [optional] 
**Fields** | [**[]GlipMessageAttachmentFieldsInfo**](GlipMessageAttachmentFieldsInfo.md) | Information on navigation | [optional] 
**Footnote** | [**GlipMessageAttachmentFootnoteInfo**](GlipMessageAttachmentFootnoteInfo.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


