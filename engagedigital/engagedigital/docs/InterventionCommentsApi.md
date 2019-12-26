# \InterventionCommentsApi

All URIs are relative to *https://DOMAIN.api.engagement.dimelo.com/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInterventionComment**](InterventionCommentsApi.md#CreateInterventionComment) | **Post** /intervention_comments | Creating an intervention comment
[**DeleteInterventionComment**](InterventionCommentsApi.md#DeleteInterventionComment) | **Delete** /intervention_comments/{interventionCommentId} | Deleting an intervention comment
[**GetAllInterventionComments**](InterventionCommentsApi.md#GetAllInterventionComments) | **Get** /intervention_comments | Getting all intervention comments
[**GetInterventionComment**](InterventionCommentsApi.md#GetInterventionComment) | **Get** /intervention_comments/{interventionCommentId} | Getting an intervention comment from its id



## CreateInterventionComment

> InterventionComment CreateInterventionComment(ctx, body, interventionId, userId)
Creating an intervention comment

This method creates a new intervention comment. In case of success it renders the created comment, otherwise, it renders an error (422 HTTP code). It creates comment as token’s user. If token’s user does not have “read” on given intervention’s source a 404 HTTP response will be returned.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | **string**| The comment body (mandatory). | 
**interventionId** | **string**| The comment intervention id (mandatory). | 
**userId** | **string**| The comment user id (mandatory). | 

### Return type

[**InterventionComment**](InterventionComment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInterventionComment

> InterventionComment DeleteInterventionComment(ctx, interventionCommentId)
Deleting an intervention comment

This method destroys an intervention comment. It renders comment itself. If token’s user does not have “read” on comment’s source a 404 HTTP response will be returned.  Authorization​: no.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**interventionCommentId** | **string**|  | 

### Return type

[**InterventionComment**](InterventionComment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllInterventionComments

> GetAllInterventionCommentsResponse GetAllInterventionComments(ctx, optional)
Getting all intervention comments

This method renders interventions comments ordered by creation date (descending). Only comments in sources where token’s user has “read” permission are returned.  Authorization​: no.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllInterventionCommentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllInterventionCommentsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **interventionId** | **optional.String**| To filter comments on given intervention id. | 
 **threadId** | **optional.String**| To filter comments on given thread id. | 
 **userId** | **optional.String**| To filter comments on given user id. | 
 **identityId** | **optional.String**| To filter comments on given identity id. | 
 **offset** | **optional.Int32**| The record index to start. Default value is 0. | 
 **limit** | **optional.Int32**| The max number of records to return. Default value is 30, max value is 150. | 

### Return type

[**GetAllInterventionCommentsResponse**](GetAllInterventionCommentsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInterventionComment

> InterventionComment GetInterventionComment(ctx, interventionCommentId)
Getting an intervention comment from its id

This method renders an intervention comment from given id. If token’s user does not have “read” on comment’s source a 404 HTTP response will be returned.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**interventionCommentId** | **string**|  | 

### Return type

[**InterventionComment**](InterventionComment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

