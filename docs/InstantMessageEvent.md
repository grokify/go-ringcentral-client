# InstantMessageEvent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Internal identifier of a message | [optional] [default to null]
**To** | [**[]InstantMessageEventCallerInfo**](InstantMessageEvent.CallerInfo.md) | Message receiver(s) information | [optional] [default to null]
**From** | [**InstantMessageEventCallerInfo**](InstantMessageEvent.CallerInfo.md) | Message sender information | [optional] [default to null]
**Type_** | **string** | Type of a message. The default value is &#39;SMS&#39; | [optional] [default to null]
**CreationTime** | [**time.Time**](time.Time.md) | Message creation datetime in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z | [optional] [default to null]
**LastModifiedTime** | [**time.Time**](time.Time.md) | The datetime when the message was modified in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z | [optional] [default to null]
**ReadStatus** | **string** | Status of a message. The default value is &#39;Unread&#39; | [optional] [default to null]
**Priority** | **string** | The default value is &#39;Normal&#39; | [optional] [default to null]
**Attachments** | [**[]InstantMessageAttachmentInfo**](InstantMessageAttachmentInfo.md) | Message attachment data | [optional] [default to null]
**Direction** | **string** | Message direction. The default value is &#39;Inbound&#39; | [optional] [default to null]
**Availability** | **string** | Message availability status. The default value is &#39;Alive&#39; | [optional] [default to null]
**Subject** | **string** | Message subject. It replicates message text which is also returned as an attachment | [optional] [default to null]
**MessageStatus** | **string** | Status of a message. The default value is &#39;Received&#39; | [optional] [default to null]
**ConversationId** | **string** | Identifier of the conversation the message belongs to | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


