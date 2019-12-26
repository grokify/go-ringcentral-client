# \ThreadsApi

All URIs are relative to *https://DOMAIN.api.engagement.dimelo.com/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveThread**](ThreadsApi.md#ArchiveThread) | **Put** /content_threads/{threadId}/ignore | Archiving a thread
[**CategorizeThread**](ThreadsApi.md#CategorizeThread) | **Put** /content_threads/{threadId}/update_categories | Categorizing a thread
[**CloseThread**](ThreadsApi.md#CloseThread) | **Put** /content_threads/{threadId}/close | Close a thread
[**GetAllThreads**](ThreadsApi.md#GetAllThreads) | **Get** /content_threads | Getting all threads
[**GetThread**](ThreadsApi.md#GetThread) | **Get** /content_threads/{threadId} | Getting a thread from its id
[**OpenThread**](ThreadsApi.md#OpenThread) | **Get** /content_threads/{threadId}/open | Open a thread



## ArchiveThread

> Thread ArchiveThread(ctx, threadId)
Archiving a thread

Archives the contents of a thread. If token’s user does not have “read” on thread’s source a 404 HTTP response will be returned.  If the thread is already being archived, a 409 HTTP response will be returned.  Authorization​: no.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadId** | **string**|  | 

### Return type

[**Thread**](Thread.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CategorizeThread

> Thread CategorizeThread(ctx, threadId, optional)
Categorizing a thread

This method updates the categories of a thread. If token’s user does not have “read” on thread’s source a 404 HTTP response will be returned.  If the thread is already being categorized, a 409 HTTP response will be returned.  Authorization​: no.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadId** | **string**|  | 
 **optional** | ***CategorizeThreadOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CategorizeThreadOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **threadCategoryIds** | [**optional.Interface of []string**](string.md)| An array containing the new categories to set on the thread. | 

### Return type

[**Thread**](Thread.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloseThread

> Thread CloseThread(ctx, threadId)
Close a thread

Thread closure/opening is only available for the following sources: * Emails * Answers * Ideas * Facebook Messenger * Google+ * Lithium * Mobile Messaging  Starts a job to close a thread. It returns the thread but as the job is asynchronous, the state of the “close” attribute in the returned object do not is the one when the job started.  If token’s user does not have “read” on thread’s source a 404 HTTP response will be returned. Returns a 403 if the thread cannot be closed or if the user does not have the permission to close a thread.  Authorization​: no.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadId** | **string**|  | 

### Return type

[**Thread**](Thread.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllThreads

> GetAllThreadsResponse GetAllThreads(ctx, optional)
Getting all threads

This method renders threads ordered by last content date (descending). Only threads in sources where token’s user has “read” permission are returned.  Authorization​: no.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllThreadsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllThreadsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **optional.Int32**| A search query to filter threads. Please refer to ​Search &amp; filtering parameters​ for more details. | 
 **offset** | **optional.Int32**| The record index to start. Default value is 0. | 
 **limit** | **optional.Int32**| The max number of records to return. Default value is 30, max value is 150. | 

### Return type

[**GetAllThreadsResponse**](GetAllThreadsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetThread

> Thread GetThread(ctx, threadId)
Getting a thread from its id

This method renders a thread from given id. If token’s user does not have “read” on thread’s source a 404 HTTP response will be returned.  Authorization​: no.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadId** | **string**|  | 

### Return type

[**Thread**](Thread.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpenThread

> Thread OpenThread(ctx, threadId)
Open a thread

Thread closure/opening is only available for the following sources:  * Emails * Answers * Ideas * Facebook Messenger * Google+ * Lithium * Mobile Messaging  Starts a job to open a thread. It returns the thread but as the job is asynchronous, the state of the “close” attribute in the returned object is the one when the job started.  If token’s user does not have “read” on thread’s source a 404 HTTP response will be returned. Returns a 403 if the thread cannot be opened or if the user does not have the permission to open a thread.  Authorization​: no.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadId** | **string**|  | 

### Return type

[**Thread**](Thread.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

