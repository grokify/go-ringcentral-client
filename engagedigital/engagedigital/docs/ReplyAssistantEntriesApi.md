# \ReplyAssistantEntriesApi

All URIs are relative to *https://DOMAIN.api.engagement.dimelo.com/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateReplyAssistantEntry**](ReplyAssistantEntriesApi.md#CreateReplyAssistantEntry) | **Post** /reply_assistant/entries | Creating an entry
[**DeleteReplyAssistantEntry**](ReplyAssistantEntriesApi.md#DeleteReplyAssistantEntry) | **Delete** /reply_assistant/entries/{replyAssistantEntryId} | Deleting a reply assistant entry
[**GetAllReplyAssistantEntries**](ReplyAssistantEntriesApi.md#GetAllReplyAssistantEntries) | **Get** /reply_assistant/entries | Getting​ a​ll​ reply assistant e​ntries
[**GetReplyAssistantEntry**](ReplyAssistantEntriesApi.md#GetReplyAssistantEntry) | **Get** /reply_assistant/entries/{replyAssistantEntryId} | Getting a reply assistant entry from its id
[**UpdateReplyAssistantEntry**](ReplyAssistantEntriesApi.md#UpdateReplyAssistantEntry) | **Put** /reply_assistant/entries/{replyAssistantEntryId} | Updating a reply assistant entry



## CreateReplyAssistantEntry

> ReplyAssistantEntry CreateReplyAssistantEntry(ctx, label)
Creating an entry

This method creates a reply assistant entry. In case of success it renders the entry, otherwise, it renders an error (422 HTTP code).  Authorization​: only users that have the right to manage reply assistant.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**label** | **string**| The name of the entry. | 

### Return type

[**ReplyAssistantEntry**](ReplyAssistantEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteReplyAssistantEntry

> ReplyAssistantEntry DeleteReplyAssistantEntry(ctx, replyAssistantEntryId)
Deleting a reply assistant entry

This method destroys an existing entry. It renders the entry itself. It renders a 404 if id is invalid.  Authorization​: only users that have the right to manage reply assistant.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**replyAssistantEntryId** | **string**|  | 

### Return type

[**ReplyAssistantEntry**](ReplyAssistantEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllReplyAssistantEntries

> GetAllReplyAssistantEntriesResponse GetAllReplyAssistantEntries(ctx, optional)
Getting​ a​ll​ reply assistant e​ntries

This method renders all entries ordered by creation date (ascending).  Authorization​: only users that have the right to manage reply assistant.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllReplyAssistantEntriesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllReplyAssistantEntriesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.Int32**| The record index to start. Default value is 0. | 
 **limit** | **optional.Int32**| The max number of records to return. Default value is 30, max value is 150. | 

### Return type

[**GetAllReplyAssistantEntriesResponse**](GetAllReplyAssistantEntriesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReplyAssistantEntry

> ReplyAssistantEntry GetReplyAssistantEntry(ctx, replyAssistantEntryId)
Getting a reply assistant entry from its id

This method renders an entry from given id.  Authorization​: only users that have the right to manage reply assistant.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**replyAssistantEntryId** | **string**|  | 

### Return type

[**ReplyAssistantEntry**](ReplyAssistantEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateReplyAssistantEntry

> ReplyAssistantEntry UpdateReplyAssistantEntry(ctx, replyAssistantEntryId, optional)
Updating a reply assistant entry

This method updates an existing entry from given attributes and renders it in case of success.  Authorization​: only users that have the right to manage reply assistant.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**replyAssistantEntryId** | **string**|  | 
 **optional** | ***UpdateReplyAssistantEntryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateReplyAssistantEntryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **label** | **optional.String**| The name of the entry. | 
 **foreignId** | **optional.String**| The internal/company id of the entry. This is used to match Engage Digital entry’s id with the company one. Example: KB042. | 
 **categoryIds** | [**optional.Interface of []string**](string.md)| To restrict the entry to a set of Engage Digital categories. Then, KB entries that do not match message’s categories to which you are replying will not be suggested. | 
 **shortcuts** | [**optional.Interface of []string**](string.md)| entry shortcuts | 
 **entryGroupId** | **optional.String**| Entry group id. | 
 **sourceIds** | [**optional.Interface of []string**](string.md)| Source ids (array) | 

### Return type

[**ReplyAssistantEntry**](ReplyAssistantEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

