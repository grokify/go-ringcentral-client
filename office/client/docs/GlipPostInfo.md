# GlipPostInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Internal identifier of a post | [optional] 
**GroupId** | **string** | Internal identifier of a group a post belongs to | [optional] 
**Type** | **string** | Type of a post | [optional] 
**Text** | **string** | For &#39;TextMessage&#39; post type only. Message text | [optional] 
**CreatorId** | **string** | Internal identifier of a user - author of a post | [optional] 
**AddedPersonIds** | **[]string** | For PersonsAdded post type only. Identifiers of persons added to a group | [optional] 
**CreationTime** | [**time.Time**](time.Time.md) | Post creation datetime in ISO 8601 format | [optional] 
**LastModifiedTime** | [**time.Time**](time.Time.md) | Post last modification datetime in ISO 8601 format | [optional] 
**Attachments** | [**[]GlipMessageAttachmentInfo**](GlipMessageAttachmentInfo.md) | List of posted attachments | [optional] 
**Mentions** | [**[]GlipMentionsInfo**](GlipMentionsInfo.md) | List of posted attachments | [optional] 
**Activity** | **string** | activity type | [optional] 
**Title** | **string** | Title of the message. Can be set for bot messages only | [optional] 
**IconUri** | **string** | URI to an image to use as the icon for this message. | [optional] 
**IconEmoji** | **string** | Emoji to use as the icon for a message | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


