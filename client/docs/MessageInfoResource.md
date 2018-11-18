# MessageInfoResource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uri** | **string** |  | [optional] 
**Id** | **string** |  | [optional] 
**To** | [**[]CallerInfo**](CallerInfo.md) |  | [optional] 
**From** | [**CallerInfo**](CallerInfo.md) |  | [optional] 
**Type** | **string** |  | [optional] 
**CreationTime** | [**time.Time**](time.Time.md) |  | [optional] 
**ReadStatus** | **string** |  | [optional] 
**Priority** | **string** |  | [optional] 
**Attachments** | [**[]MessageAttachmentInfo**](MessageAttachmentInfo.md) |  | [optional] 
**Direction** | **string** |  | [optional] 
**Availability** | **string** |  | [optional] 
**Subject** | **string** |  | [optional] 
**MessageStatus** | **string** |  | [optional] 
**FaxResolution** | **string** |  | [optional] 
**FaxPageCount** | **int32** |  | [optional] 
**DeliveryErrorCode** | **string** |  | [optional] 
**SmsDeliveryTime** | [**time.Time**](time.Time.md) |  | [optional] 
**SmsSendingAttemptsCount** | **int64** |  | [optional] 
**ConversationId** | **int64** |  | [optional] 
**Conversation** | [**ConversationResource**](ConversationResource.md) |  | [optional] 
**LastModifiedTime** | [**time.Time**](time.Time.md) |  | [optional] 
**PgToDepartment** | **bool** |  | [optional] [default to false]
**VmTranscriptionStatus** | **string** |  | [optional] 
**SourceFilePath** | **string** |  | [optional] 
**CoverIndex** | **int32** |  | [optional] 
**CoverPageText** | **string** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


