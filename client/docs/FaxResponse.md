# FaxResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Internal identifier of a message | [optional] [default to null]
**Uri** | **string** | Canonical URI of a message | [optional] [default to null]
**Type_** | **string** | Message type - &#39;Fax&#39; | [optional] [default to null]
**From** | [***CallerInfoFrom**](CallerInfoFrom.md) | Sender information | [optional] [default to null]
**To** | [**[]CallerInfoTo**](CallerInfoTo.md) | Recipient information | [optional] [default to null]
**CreationTime** | [**time.Time**](time.Time.md) | Message creation datetime in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z | [optional] [default to null]
**ReadStatus** | **string** | Message read status | [optional] [default to null]
**Priority** | **string** | Message priority | [optional] [default to null]
**Attachments** | [**[]MessageAttachmentInfo**](MessageAttachmentInfo.md) | The list of message attachments | [optional] [default to null]
**Direction** | **string** | Message direction | [optional] [default to null]
**Availability** | **string** | Message availability status. Message in &#39;Deleted&#39; state is still preserved with all its attachments and can be restored. &#39;Purged&#39; means that all attachments are already deleted and the message itself is about to be physically deleted shortly | [optional] [default to null]
**MessageStatus** | **string** | Message status. &#39;Queued&#39; - the message is queued for sending; &#39;Sent&#39; - a message is successfully sent; &#39;SendingFailed&#39; - a message sending attempt has failed; &#39;Received&#39; - a message is received (inbound messages have this status by default) | [optional] [default to null]
**FaxResolution** | **string** | Resolution of a fax message. (&#39;High&#39; for black and white image scanned at 200 dpi, &#39;Low&#39; for black and white image scanned at 100 dpi) | [optional] [default to null]
**FaxPageCount** | **int32** | Page count in a fax message | [optional] [default to null]
**LastModifiedTime** | [**time.Time**](time.Time.md) | Datetime when the message was modified on server in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z | [optional] [default to null]
**CoverIndex** | **int32** | Cover page identifier. For the list of available cover page identifiers please call the method Fax Cover Pages | [optional] [default to null]
**CoverPageText** | **string** | Cover page text, entered by the fax sender and printed on the cover page. Maximum length is limited to 1024 symbols | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


