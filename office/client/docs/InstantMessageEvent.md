# InstantMessageEvent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Internal identifier of an message | [optional] 
**To** | [**[]NotificationRecipientInfo**](NotificationRecipientInfo.md) | Message receiver(s) information | [optional] 
**From** | [**SenderInfo**](SenderInfo.md) |  | [optional] 
**Type** | **string** | Type of a message. The default value is &#39;SMS&#39; | [optional] 
**CreationTime** | [**time.Time**](time.Time.md) | Message creation datetime in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z | [optional] 
**LastModifiedTime** | [**time.Time**](time.Time.md) | Datetime when the message was modified in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z | [optional] 
**ReadStatus** | **string** | Status of a message. The default value is &#39;Unread&#39; | [optional] 
**Priority** | **string** | The default value is &#39;Normal&#39; | [optional] 
**Attachments** | [**[]MessageAttachmentInfo**](MessageAttachmentInfo.md) | Message attachment data | [optional] 
**Direction** | **string** | Message direction. The default value is &#39;Inbound&#39; | [optional] 
**Availability** | **string** | Message availability status. The default value is &#39;Alive&#39; | [optional] 
**Subject** | **string** | Message subject. It replicates message text which is also returned as an attachment | [optional] 
**MessageStatus** | **string** | Status of a message. The default value is &#39;Received&#39; | [optional] 
**ConversationId** | **string** | Identifier of the conversation the message belongs to | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


