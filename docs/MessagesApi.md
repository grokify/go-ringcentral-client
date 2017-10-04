# \MessagesApi

All URIs are relative to *https://platform.devtest.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RestapiV10AccountAccountIdExtensionExtensionIdCompanyPagerPost**](MessagesApi.md#RestapiV10AccountAccountIdExtensionExtensionIdCompanyPagerPost) | **Post** /restapi/v1.0/account/{accountId}/extension/{extensionId}/company-pager | 
[**RestapiV10AccountAccountIdExtensionExtensionIdFaxPost**](MessagesApi.md#RestapiV10AccountAccountIdExtensionExtensionIdFaxPost) | **Post** /restapi/v1.0/account/{accountId}/extension/{extensionId}/fax | 
[**RestapiV10AccountAccountIdExtensionExtensionIdMessageStoreGet**](MessagesApi.md#RestapiV10AccountAccountIdExtensionExtensionIdMessageStoreGet) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/message-store | 
[**RestapiV10AccountAccountIdExtensionExtensionIdMessageStoreMessageIdContentAttachmentIdGet**](MessagesApi.md#RestapiV10AccountAccountIdExtensionExtensionIdMessageStoreMessageIdContentAttachmentIdGet) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/message-store/{messageId}/content/{attachmentId} | 
[**RestapiV10AccountAccountIdExtensionExtensionIdMessageStoreMessageIdDelete**](MessagesApi.md#RestapiV10AccountAccountIdExtensionExtensionIdMessageStoreMessageIdDelete) | **Delete** /restapi/v1.0/account/{accountId}/extension/{extensionId}/message-store/{messageId} | 
[**RestapiV10AccountAccountIdExtensionExtensionIdMessageStoreMessageIdGet**](MessagesApi.md#RestapiV10AccountAccountIdExtensionExtensionIdMessageStoreMessageIdGet) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/message-store/{messageId} | 
[**RestapiV10AccountAccountIdExtensionExtensionIdMessageStoreMessageIdPut**](MessagesApi.md#RestapiV10AccountAccountIdExtensionExtensionIdMessageStoreMessageIdPut) | **Put** /restapi/v1.0/account/{accountId}/extension/{extensionId}/message-store/{messageId} | 
[**RestapiV10AccountAccountIdExtensionExtensionIdMessageSyncGet**](MessagesApi.md#RestapiV10AccountAccountIdExtensionExtensionIdMessageSyncGet) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/message-sync | 
[**RestapiV10AccountAccountIdExtensionExtensionIdSmsPost**](MessagesApi.md#RestapiV10AccountAccountIdExtensionExtensionIdSmsPost) | **Post** /restapi/v1.0/account/{accountId}/extension/{extensionId}/sms | 


# **RestapiV10AccountAccountIdExtensionExtensionIdCompanyPagerPost**
> MessageInfo RestapiV10AccountAccountIdExtensionExtensionIdCompanyPagerPost($accountId, $extensionId, $body)



Create and Send Pager Message


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **body** | [**Body6**](Body6.md)|  | [optional] 

### Return type

[**MessageInfo**](MessageInfo.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdExtensionExtensionIdFaxPost**
> MessageInfo RestapiV10AccountAccountIdExtensionExtensionIdFaxPost($accountId, $extensionId, $body)



Create and Send Fax Message


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **body** | [**Body8**](Body8.md)|  | [optional] 

### Return type

[**MessageInfo**](MessageInfo.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdExtensionExtensionIdMessageStoreGet**
> InlineResponseDefault21 RestapiV10AccountAccountIdExtensionExtensionIdMessageStoreGet($accountId, $extensionId, $availability, $conversationId, $dateFrom, $dateTo, $direction, $distinctConversations, $messageType, $readStatus, $page, $perPage, $phoneNumber)



Get Message List


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **availability** | **string**| Specifies the availability status for the resulting messages. Default value is &#39;Alive&#39;. Multiple values are accepted | [optional] 
 **conversationId** | **int64**| Specifies the conversation identifier for the resulting messages | [optional] 
 **dateFrom** | **time.Time**| The start datetime for resulting messages in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z. The default value is dateTo minus 24 hours | [optional] 
 **dateTo** | **time.Time**| The end datetime for resulting messages in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z. The default value is current time | [optional] 
 **direction** | **string**| The direction for the resulting messages. If not specified, both inbound and outbound messages are returned. Multiple values are accepted | [optional] 
 **distinctConversations** | **bool**| If &#39;True&#39;, then the latest messages per every conversation ID are returned | [optional] 
 **messageType** | **string**| The type of the resulting messages. If not specified, all messages without message type filtering are returned. Multiple values are accepted | [optional] 
 **readStatus** | **string**| The read status for the resulting messages. Multiple values are accepted | [optional] 
 **page** | **int32**| Indicates the page number to retrieve. Only positive number values are allowed. Default value is &#39;1&#39; | [optional] 
 **perPage** | **int32**| Indicates the page size (number of items). If not specified, the value is &#39;100&#39; by default | [optional] 
 **phoneNumber** | **string**| The phone number. If specified, messages are returned for this particular phone number only | [optional] 

### Return type

[**InlineResponseDefault21**](inline_response_default_21.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdExtensionExtensionIdMessageStoreMessageIdContentAttachmentIdGet**
> Binary RestapiV10AccountAccountIdExtensionExtensionIdMessageStoreMessageIdContentAttachmentIdGet($accountId, $extensionId, $messageId, $attachmentId)



Get Message Attachment


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **messageId** | **string**| Internal identifier of a message | 
 **attachmentId** | **string**| Internal identifier of a message attachment | 

### Return type

[**Binary**](Binary.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdExtensionExtensionIdMessageStoreMessageIdDelete**
> RestapiV10AccountAccountIdExtensionExtensionIdMessageStoreMessageIdDelete($accountId, $extensionId, $messageId, $purge, $conversationId)



Delete Message by ID


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **messageId** | **string**| Internal identifier of a message | 
 **purge** | **bool**| If the value is &#39;True&#39;, then the message is purged immediately with all the attachments. The default value is &#39;False&#39; | [optional] 
 **conversationId** | **int64**| Internal identifier of a message thread | [optional] 

### Return type

void (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdExtensionExtensionIdMessageStoreMessageIdGet**
> MessageInfo RestapiV10AccountAccountIdExtensionExtensionIdMessageStoreMessageIdGet($accountId, $extensionId, $messageId)



Get Message by ID


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **messageId** | **string**| Internal identifier of a message | 

### Return type

[**MessageInfo**](MessageInfo.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdExtensionExtensionIdMessageStoreMessageIdPut**
> MessageInfo RestapiV10AccountAccountIdExtensionExtensionIdMessageStoreMessageIdPut($accountId, $extensionId, $messageId, $body)



Update Message by ID


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **messageId** | **string**| Internal identifier of a message | 
 **body** | [**Body13**](Body13.md)|  | [optional] 

### Return type

[**MessageInfo**](MessageInfo.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdExtensionExtensionIdMessageSyncGet**
> InlineResponseDefault22 RestapiV10AccountAccountIdExtensionExtensionIdMessageSyncGet($accountId, $extensionId, $conversationId, $dateFrom, $dateTo, $direction, $distinctConversations, $messageType, $recordCount, $syncToken, $syncType)



Message Synchronization


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **conversationId** | **int64**| Conversation identifier for the resulting messages. Meaningful for SMS and Pager messages only. | [optional] 
 **dateFrom** | **time.Time**| The start datetime for resulting messages in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z. The default value is dateTo minus 24 hours | [optional] 
 **dateTo** | **time.Time**| The end datetime for resulting messages in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z. The default value is current time | [optional] 
 **direction** | **string**| Direction for the resulting messages. If not specified, both inbound and outbound messages are returned. Multiple values are accepted | [optional] 
 **distinctConversations** | **bool**| If &#39;True&#39;, then the latest messages per every conversation ID are returned | [optional] 
 **messageType** | **string**| Type for the resulting messages. If not specified, all types of messages are returned. Multiple values are accepted | [optional] 
 **recordCount** | **int32**| Limits the number of records to be returned (works in combination with dateFrom and dateTo if specified) | [optional] 
 **syncToken** | **string**| Value of syncToken property of last sync request response | [optional] 
 **syncType** | **string**| Type of message synchronization | [optional] 

### Return type

[**InlineResponseDefault22**](inline_response_default_22.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdExtensionExtensionIdSmsPost**
> MessageInfo RestapiV10AccountAccountIdExtensionExtensionIdSmsPost($accountId, $extensionId, $body)



Create and Send SMS Message


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **body** | [**Body15**](Body15.md)|  | [optional] 

### Return type

[**MessageInfo**](MessageInfo.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

