# \MessagesApi

All URIs are relative to *https://platform.devtest.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMessage**](MessagesApi.md#DeleteMessage) | **Delete** /restapi/v1.0/account/{accountId}/extension/{extensionId}/message-store/{messageId} | Delete Message(s) by ID
[**DeleteMessagesByFilter**](MessagesApi.md#DeleteMessagesByFilter) | **Delete** /restapi/v1.0/account/{accountId}/extension/{extensionId}/message-store | Delete Conversations by ID&#39;s
[**GetFaxCoverPages**](MessagesApi.md#GetFaxCoverPages) | **Get** /restapi/v1.0/dictionary/fax-cover-page | Get Fax Cover Pages
[**ListMessages**](MessagesApi.md#ListMessages) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/message-store | Get Message List
[**LoadMessage**](MessagesApi.md#LoadMessage) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/message-store/{messageId} | Get Message(s) by ID
[**LoadMessageAttachment**](MessagesApi.md#LoadMessageAttachment) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/message-store/{messageId}/content/{attachmentId} | Get Message Attachment
[**SendFaxMessage**](MessagesApi.md#SendFaxMessage) | **Post** /restapi/v1.0/account/{accountId}/extension/{extensionId}/fax | Create Fax Message
[**SendInternalMessage**](MessagesApi.md#SendInternalMessage) | **Post** /restapi/v1.0/account/{accountId}/extension/{extensionId}/company-pager | Create Pager Message
[**SendSMS**](MessagesApi.md#SendSMS) | **Post** /restapi/v1.0/account/{accountId}/extension/{extensionId}/sms | Create SMS Message
[**SyncMessages**](MessagesApi.md#SyncMessages) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/message-sync | Get Message Sync
[**UpdateMessage**](MessagesApi.md#UpdateMessage) | **Put** /restapi/v1.0/account/{accountId}/extension/{extensionId}/message-store/{messageId} | Update Message(s) by ID


# **DeleteMessage**
> DeleteMessage(ctx, accountId, extensionId, messageId, optional)
Delete Message(s) by ID

<p style='font-style:italic;'></p><p>Deletes message(s) by the given message ID(s). The first call of this method transfers the message to the 'Delete' status. The second call transfers the deleted message to the 'Purged' status. If it is required to make the message 'Purged' immediately (from the first call), then set the query parameter purge to 'True'.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>EditMessages</td><td>Viewing and updating user messages</td></tr><tr><td class='code'>ReadMessages</td><td>Viewing user messages</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
  **messageId** | **int32**| Internal identifier of a message | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **messageId** | **int32**| Internal identifier of a message | 
 **purge** | **bool**| If the value is &#39;True&#39;, then the message is purged immediately with all the attachments. The default value is &#39;False&#39; | 
 **conversationId** | **int32**| Internal identifier of a message thread | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMessagesByFilter**
> DeleteMessagesByFilter(ctx, extensionId, accountId, optional)
Delete Conversations by ID's

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **extensionId** | **string**|  | 
  **accountId** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **extensionId** | **string**|  | 
 **accountId** | **string**|  | 
 **conversationId** | [**[]string**](string.md)|  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFaxCoverPages**
> GetFaxCoverPages(ctx, optional)
Get Fax Cover Pages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32**| Indicates the page number to retrieve. Only positive number values are allowed. Default value is &#39;1&#39; | 
 **perPage** | **int32**| Indicates the page size (number of items). If not specified, the value is &#39;100&#39; by default | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMessages**
> GetMessageList ListMessages(ctx, accountId, extensionId, optional)
Get Message List

<p style='font-style:italic;'>Since 1.0.2</p><p>Returns the list of messages from an extension mailbox.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadMessages</td><td>Viewing user messages</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **availability** | [**[]string**](string.md)| Specifies the availability status for the resulting messages. Default value is &#39;Alive&#39;. Multiple values are accepted | 
 **conversationId** | **int32**| Specifies the conversation identifier for the resulting messages | 
 **dateFrom** | **time.Time**| The start datetime for resulting messages in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z. The default value is dateTo minus 24 hours | 
 **dateTo** | **time.Time**| The end datetime for resulting messages in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z. The default value is current time | 
 **direction** | [**[]string**](string.md)| The direction for the resulting messages. If not specified, both inbound and outbound messages are returned. Multiple values are accepted | 
 **distinctConversations** | **bool**| If &#39;True&#39;, then the latest messages per every conversation ID are returned | 
 **messageType** | [**[]string**](string.md)| The type of the resulting messages. If not specified, all messages without message type filtering are returned. Multiple values are accepted | 
 **readStatus** | [**[]string**](string.md)| The read status for the resulting messages. Multiple values are accepted | 
 **page** | **int32**| Indicates the page number to retrieve. Only positive number values are allowed. Default value is &#39;1&#39; | 
 **perPage** | **int32**| Indicates the page size (number of items). If not specified, the value is &#39;100&#39; by default | 
 **phoneNumber** | **string**| The phone number. If specified, messages are returned for this particular phone number only | 

### Return type

[**GetMessageList**](GetMessageList.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadMessage**
> GetMessageInfoResponse LoadMessage(ctx, accountId, extensionId, messageId)
Get Message(s) by ID

<p style='font-style:italic;'>Since 1.0.2</p><p>Returns individual message record(s) by the given message ID(s). The length of inbound messages is unlimited. Batch request is supported, see Batch Requests for details.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadMessages</td><td>Viewing user messages</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
  **messageId** | **int32**| Internal identifier of a message | 

### Return type

[**GetMessageInfoResponse**](GetMessageInfoResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadMessageAttachment**
> LoadMessageAttachment(ctx, accountId, extensionId, attachmentId, messageId, optional)
Get Message Attachment

<p style='font-style:italic;'>Since 1.0.4 (Release 5.13)</p><p>Returns particular message attachment data as a media stream.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadMessages</td><td>Viewing user messages</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
  **attachmentId** | **int32**| Internal identifier of a message attachment | 
  **messageId** | **int32**| Internal identifier of a message | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **attachmentId** | **int32**| Internal identifier of a message attachment | 
 **messageId** | **int32**| Internal identifier of a message | 
 **range_** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendFaxMessage**
> FaxResponse SendFaxMessage(ctx, accountId, extensionId)
Create Fax Message

<p style='font-style:italic;'>Since 1.0.2</p><p>Creates and sends/resends new fax message. Resend can be done if sending failed.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>Faxes</td><td>Sending and receiving faxes</td></tr><tr><td class='code'>ReadMessages</td><td>Viewing user messages</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Heavy</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account (integer) or tilde (~) to indicate the account which was logged-in within the current session. | 
  **extensionId** | **string**| Internal identifier of an extension (integer) or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 

### Return type

[**FaxResponse**](FaxResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: multipart/mixed; boundary=Boundary_1_14413901_1361871080888
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendInternalMessage**
> GetMessageInfoResponse SendInternalMessage(ctx, accountId, extensionId, body)
Create Pager Message

<p style='font-style:italic;'>Since 1.0.2</p><p>Creates and sends a pager message.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>InternalMessages</td><td>Sending and receiving intra-company text messages</td></tr><tr><td class='code'>ReadMessages</td><td>Viewing user messages</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
  **body** | [**CreatePagerMessageRequest**](CreatePagerMessageRequest.md)| JSON body | 

### Return type

[**GetMessageInfoResponse**](GetMessageInfoResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendSMS**
> GetMessageInfoResponse SendSMS(ctx, accountId, extensionId, body)
Create SMS Message

<p style='font-style:italic;'>Since 1.0.2</p><p>Creates and sends new SMS message. Sending SMS messages simultaneously to different recipients is limited up to 50 requests per minute; relevant for all client applications.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>SMS</td><td>Sending and receiving SMS (text) messages</td></tr><tr><td class='code'>ReadMessages</td><td>Viewing user messages</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
  **body** | [**CreateSmsMessage**](CreateSmsMessage.md)| JSON body | 

### Return type

[**GetMessageInfoResponse**](GetMessageInfoResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncMessages**
> GetMessageSyncResponse SyncMessages(ctx, accountId, extensionId, optional)
Get Message Sync

<p style='font-style:italic;'>Since 1.0.4 (Release 5.13)</p><p>Provides facilities to synchronize mailbox content stored externally with server state.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadMessages</td><td>Viewing user messages</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **conversationId** | **int32**| Conversation identifier for the resulting messages. Meaningful for SMS and Pager messages only. | 
 **dateFrom** | **string**| The start datetime for resulting messages in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z. The default value is dateTo minus 24 hours | 
 **dateTo** | **string**| The end datetime for resulting messages in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z. The default value is current time | 
 **direction** | [**[]string**](string.md)| Direction for the resulting messages. If not specified, both inbound and outbound messages are returned. Multiple values are accepted | 
 **distinctConversations** | **bool**| If &#39;True&#39;, then the latest messages per every conversation ID are returned | 
 **messageType** | [**[]string**](string.md)| Type for the resulting messages. If not specified, all types of messages are returned. Multiple values are accepted | 
 **recordCount** | **int32**| Limits the number of records to be returned (works in combination with dateFrom and dateTo if specified) | 
 **syncToken** | **string**| Value of syncToken property of last sync request response | 
 **syncType** | [**[]string**](string.md)| Type of message synchronization | 

### Return type

[**GetMessageSyncResponse**](GetMessageSyncResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMessage**
> GetMessageInfoResponse UpdateMessage(ctx, accountId, extensionId, messageId, body)
Update Message(s) by ID

<p style='font-style:italic;'>Since 1.0.2</p><p>Updates message(s) by ID(s). Batch request is supported, see Batch Requests for details. Currently, only the message read status updating is supported.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>EditMessages</td><td>Viewing and updating user messages</td></tr><tr><td class='code'>ReadMessages</td><td>Viewing user messages</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
  **messageId** | **int32**| Internal identifier of a message | 
  **body** | [**UpdateMessageRequest**](UpdateMessageRequest.md)| JSON body | 

### Return type

[**GetMessageInfoResponse**](GetMessageInfoResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

