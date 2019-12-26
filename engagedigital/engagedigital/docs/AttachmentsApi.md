# \AttachmentsApi

All URIs are relative to *https://DOMAIN.api.engagement.dimelo.com/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAttachment**](AttachmentsApi.md#CreateAttachment) | **Post** /attachments | Creating an attachment
[**GetAllAttachments**](AttachmentsApi.md#GetAllAttachments) | **Get** /attachments | Getting all attachments
[**GetAttachment**](AttachmentsApi.md#GetAttachment) | **Get** /attachments/{attachmentId} | Getting an attachment from its id



## CreateAttachment

> Attachment CreateAttachment(ctx, optional)
Creating an attachment

This method allows you to create an new attachment.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateAttachmentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateAttachmentOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | **optional.Interface of *os.File****optional.*os.File**|  | 
 **private** | **optional.String**|  | 

### Return type

[**Attachment**](Attachment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data:
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllAttachments

> GetAllAttachmentsResponse GetAllAttachments(ctx, optional)
Getting all attachments

This method renders attachments ordered by creation date (descending).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllAttachmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllAttachmentsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.Int32**| The record index to start. Default value is 0. | 
 **limit** | **optional.Int32**| The max number of records to return. Default value is 30, max value is 150. | 

### Return type

[**GetAllAttachmentsResponse**](GetAllAttachmentsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAttachment

> Attachment GetAttachment(ctx, attachmentId)
Getting an attachment from its id

This method renders an attachment from given id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attachmentId** | **string**|  | 

### Return type

[**Attachment**](Attachment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

