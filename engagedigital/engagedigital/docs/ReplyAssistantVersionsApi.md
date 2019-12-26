# \ReplyAssistantVersionsApi

All URIs are relative to *https://DOMAIN.api.engagement.dimelo.com/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateReplyAssistantVersion**](ReplyAssistantVersionsApi.md#CreateReplyAssistantVersion) | **Post** /reply_assistant/versions | Creating a reply assistant version
[**DeleteReplyAssistantVersion**](ReplyAssistantVersionsApi.md#DeleteReplyAssistantVersion) | **Delete** /reply_assistant/versions/{replyAssistantVersionId} | Deleting a reply assistant version
[**GetAllReplyAssistantVersions**](ReplyAssistantVersionsApi.md#GetAllReplyAssistantVersions) | **Get** /reply_assistant/versions | Getting​ a​ll​ reply assistant versions
[**GetReplyAssistantVersion**](ReplyAssistantVersionsApi.md#GetReplyAssistantVersion) | **Get** /reply_assistant/versions/{replyAssistantVersionId} | Getting a reply assistant version from its id
[**UpdateReplyAssistantVersion**](ReplyAssistantVersionsApi.md#UpdateReplyAssistantVersion) | **Put** /reply_assistant/versions/{replyAssistantVersionId} | Updating a reply assistant version



## CreateReplyAssistantVersion

> ReplyAssistantVersion CreateReplyAssistantVersion(ctx, body, entryId, optional)
Creating a reply assistant version

This method creates a reply assistant version. In case of success it renders the version, otherwise, it renders an error (422 HTTP code, 404 if the entry_id is invalid).  Authorization​: only users that have the right to manage reply assistant.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | **string**| Body of the version | 
**entryId** | **string**| Reply assistant entry id (mandatory) | 
 **optional** | ***CreateReplyAssistantVersionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateReplyAssistantVersionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sourceIds** | [**optional.Interface of []string**](string.md)| Source ids (array) | 
 **format** | **optional.String**| Either “text” or “html” | 
 **language** | **optional.String**| Language (ex: “fr”) | 

### Return type

[**ReplyAssistantVersion**](ReplyAssistantVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteReplyAssistantVersion

> ReplyAssistantVersion DeleteReplyAssistantVersion(ctx, replyAssistantVersionId)
Deleting a reply assistant version

This method destroys an existing version. It renders the version itself. It renders a 404 if id is invalid.  Authorization​: only users that have the right to manage reply assistant.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**replyAssistantVersionId** | **string**|  | 

### Return type

[**ReplyAssistantVersion**](ReplyAssistantVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllReplyAssistantVersions

> GetAllReplyAssistantVersionsResponse GetAllReplyAssistantVersions(ctx, optional)
Getting​ a​ll​ reply assistant versions

This method renders all reply assistant versions ordered by creation date (ascending).  Authorization​: only users that have the right to manage reply assistant.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllReplyAssistantVersionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllReplyAssistantVersionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.Int32**| The record index to start. Default value is 0. | 
 **limit** | **optional.Int32**| The max number of records to return. Default value is 30, max value is 150. | 

### Return type

[**GetAllReplyAssistantVersionsResponse**](GetAllReplyAssistantVersionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReplyAssistantVersion

> ReplyAssistantVersion GetReplyAssistantVersion(ctx, replyAssistantVersionId)
Getting a reply assistant version from its id

This method renders a version from given id.  Authorization​: only users that have the right to manage reply assistant.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**replyAssistantVersionId** | **string**|  | 

### Return type

[**ReplyAssistantVersion**](ReplyAssistantVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateReplyAssistantVersion

> ReplyAssistantVersion UpdateReplyAssistantVersion(ctx, replyAssistantVersionId, optional)
Updating a reply assistant version

This method updates an existing version from given attributes and renders it in case of success.  Authorization​: only users that have the right to manage reply assistant.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**replyAssistantVersionId** | **string**|  | 
 **optional** | ***UpdateReplyAssistantVersionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateReplyAssistantVersionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **optional.String**| Body of the version | 
 **entryId** | **optional.String**| Reply assistant entry id. | 
 **sourceIds** | [**optional.Interface of []string**](string.md)| Source ids (array) | 
 **format** | **optional.String**| Either “text” or “html” | 
 **language** | **optional.String**| Language (ex: “fr”) | 

### Return type

[**ReplyAssistantVersion**](ReplyAssistantVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

