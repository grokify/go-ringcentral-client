# Content

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApprovalRequired** | **bool** |  | [optional] 
**Attachments** | [**[]ContentAttachment**](ContentAttachment.md) |  | [optional] 
**AttachmentsCount** | **int32** |  | [optional] 
**AuthorId** | **string** |  | [optional] 
**Body** | **string** |  | [optional] 
**BodyFormatted** | [**ContentBodyFormatted**](ContentBodyFormatted.md) |  | [optional] 
**BodyInputFormat** | **string** | values can be: text or html. | [optional] 
**CategoryIds** | **[]string** | are content categories if none, they are default to intervention categories or thread categories | [optional] 
**ContextData** | [**map[string]interface{}**](.md) | is present only if the content has context_data associated. The context_data hash keys are the custom fields keys. | [optional] 
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] 
**CreatedFrom** | **string** |  | [optional] 
**CreatorId** | **string** |  | [optional] 
**ForeignId** | **string** |  | [optional] 
**Id** | **string** |  | 
**InReplyToAuthorId** | **string** |  | [optional] 
**InReplyToId** | **string** |  | [optional] 
**InterventionId** | **string** |  | [optional] 
**Language** | **string** |  | [optional] 
**PrivateMessage** | **bool** |  | [optional] 
**Published** | **bool** |  | [optional] 
**RemotelyDeleted** | **bool** |  | [optional] 
**SourceId** | **string** |  | [optional] 
**SourceType** | **string** |  | [optional] 
**SourceUrl** | **string** |  | [optional] 
**Status** | **string** |  | [optional] 
**SynchronizationStatus** | **string** |  | [optional] 
**ThreadId** | **string** |  | [optional] 
**Title** | **string** |  | [optional] 
**Type** | **string** |  | [optional] 
**UpdatedAt** | [**time.Time**](time.Time.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


