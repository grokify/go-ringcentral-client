# \PresenceStatusApi

All URIs are relative to *https://DOMAIN.api.engagement.dimelo.com/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePresenceStatus**](PresenceStatusApi.md#CreatePresenceStatus) | **Post** /presence_status | Creating a presence status
[**DeletePresenceStatus**](PresenceStatusApi.md#DeletePresenceStatus) | **Delete** /presence_status/{presenceStatusId} | Deleting a presence status
[**GetAllPresenceStatus**](PresenceStatusApi.md#GetAllPresenceStatus) | **Get** /presence_status | Getting all presence statuses
[**GetPresenceStatus**](PresenceStatusApi.md#GetPresenceStatus) | **Get** /presence_status/{presenceStatusId} | Getting a presence status from its id
[**UpdatePresenceStatus**](PresenceStatusApi.md#UpdatePresenceStatus) | **Put** /presence_status/{presenceStatusId} | Updating a presence status



## CreatePresenceStatus

> PresenceStatus CreatePresenceStatus(ctx, name)
Creating a presence status

This method creates a presence status. In case of success it renders the presence status, otherwise, it renders an error (422 HTTP code).  Authorization​: only users that have the right to monitor the task view.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| The name of the presence status. | 

### Return type

[**PresenceStatus**](PresenceStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePresenceStatus

> PresenceStatus DeletePresenceStatus(ctx, presenceStatusId)
Deleting a presence status

This method destroys an existing presence status. It renders presence status itself. It renders a 404 if id is invalid.  Authorization​: only users that have the right to monitor the task view.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**presenceStatusId** | **string**|  | 

### Return type

[**PresenceStatus**](PresenceStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllPresenceStatus

> GetAllPresenceStatusResponse GetAllPresenceStatus(ctx, optional)
Getting all presence statuses

This method renders all presence statuses ordered by name (in alphabetical order).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllPresenceStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllPresenceStatusOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.Int32**| The record index to start. Default value is 0. | 
 **limit** | **optional.Int32**| The max number of records to return. Default value is 30, max value is 150. | 

### Return type

[**GetAllPresenceStatusResponse**](GetAllPresenceStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPresenceStatus

> PresenceStatus GetPresenceStatus(ctx, presenceStatusId)
Getting a presence status from its id

This method renders a presence status from given id.  Authorization​: only users that have the right to monitor the task view.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**presenceStatusId** | **string**|  | 

### Return type

[**PresenceStatus**](PresenceStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePresenceStatus

> PresenceStatus UpdatePresenceStatus(ctx, presenceStatusId, name)
Updating a presence status

This method updates an existing presence status from given attributes and renders it in case of success.  Authorization​: only users that have the right to monitor the task view.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**presenceStatusId** | **string**|  | 
**name** | **string**| The name of the presence status. | 

### Return type

[**PresenceStatus**](PresenceStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

