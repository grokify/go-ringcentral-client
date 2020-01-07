# \ReplyAssistantGroupsApi

All URIs are relative to *https://DOMAIN.api.engagement.dimelo.com/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateReplyAssistantGroup**](ReplyAssistantGroupsApi.md#CreateReplyAssistantGroup) | **Post** /reply_assistant/groups | Creating a reply assistant group
[**DeleteReplyAssistantGroup**](ReplyAssistantGroupsApi.md#DeleteReplyAssistantGroup) | **Delete** /reply_assistant/groups/{replyAssistantGroupId} | Deleting a reply assistant group
[**GetAllReplyAssistantGroups**](ReplyAssistantGroupsApi.md#GetAllReplyAssistantGroups) | **Get** /reply_assistant/groups | Getting​ a​ll​ reply assistant groups
[**GetReplyAssistantGroup**](ReplyAssistantGroupsApi.md#GetReplyAssistantGroup) | **Get** /reply_assistant/groups/{replyAssistantGroupId} | Getting a reply assistant group from its id
[**UpdateReplyAssistantGroup**](ReplyAssistantGroupsApi.md#UpdateReplyAssistantGroup) | **Put** /reply_assistant/groups/{replyAssistantGroupId} | Updating a reply assistant group



## CreateReplyAssistantGroup

> ReplyAssistantGroup CreateReplyAssistantGroup(ctx, name, optional)
Creating a reply assistant group

This method creates an entry group. In case of success it renders the group, otherwise, it renders an error (422 HTTP code).  Authorization​: only users that have the right to manage reply assistant.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| The name of the group. | 
 **optional** | ***CreateReplyAssistantGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateReplyAssistantGroupOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **entryIds** | [**optional.Interface of []string**](string.md)| List of the reply assistant entries in this group. | 
 **autocomplete** | **optional.Bool**| Used for autocompletion in chat. | 
 **position** | **optional.Int32**| Used to determine the order of the groups in the interface, in ascending order. | 

### Return type

[**ReplyAssistantGroup**](ReplyAssistantGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteReplyAssistantGroup

> ReplyAssistantGroup DeleteReplyAssistantGroup(ctx, replyAssistantGroupId)
Deleting a reply assistant group

This method destroys an existing group. It renders the group itself. It renders a 404 if id is invalid.  Authorization​: only users that have the right to manage reply assistant.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**replyAssistantGroupId** | **string**|  | 

### Return type

[**ReplyAssistantGroup**](ReplyAssistantGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllReplyAssistantGroups

> GetAllReplyAssistantGroupsResponse GetAllReplyAssistantGroups(ctx, optional)
Getting​ a​ll​ reply assistant groups

This method renders all groups ordered by creation date (ascending).  Authorization​: only users that have the right to manage reply assistant.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllReplyAssistantGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllReplyAssistantGroupsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.Int32**| The record index to start. Default value is 0. | 
 **limit** | **optional.Int32**| The max number of records to return. Default value is 30, max value is 150. | 

### Return type

[**GetAllReplyAssistantGroupsResponse**](GetAllReplyAssistantGroupsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReplyAssistantGroup

> ReplyAssistantGroup GetReplyAssistantGroup(ctx, replyAssistantGroupId)
Getting a reply assistant group from its id

This method renders an entry group from given id.  Authorization​: only users that have the right to manage reply assistant.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**replyAssistantGroupId** | **string**|  | 

### Return type

[**ReplyAssistantGroup**](ReplyAssistantGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateReplyAssistantGroup

> ReplyAssistantGroup UpdateReplyAssistantGroup(ctx, replyAssistantGroupId, optional)
Updating a reply assistant group

This method updates an existing group from given attributes and renders it in case of success.  Authorization​: only users that have the right to manage reply assistant.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**replyAssistantGroupId** | **string**|  | 
 **optional** | ***UpdateReplyAssistantGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateReplyAssistantGroupOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| The name of the group. | 
 **entryIds** | [**optional.Interface of []string**](string.md)| List of the reply assistant entries in this group. | 
 **autocomplete** | **optional.Bool**| Used for autocompletion in chat. | 
 **position** | **optional.Int32**| Used to determine the order of the groups in the interface, in ascending order. | 

### Return type

[**ReplyAssistantGroup**](ReplyAssistantGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

