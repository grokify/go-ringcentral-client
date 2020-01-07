# \ContentsApi

All URIs are relative to *https://DOMAIN.api.engagement.dimelo.com/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CategorizeContent**](ContentsApi.md#CategorizeContent) | **Put** /contents/{contentId}/update_categories | Categorizing a content
[**CreateContent**](ContentsApi.md#CreateContent) | **Post** /contents | Creating a content
[**GetAllContents**](ContentsApi.md#GetAllContents) | **Get** /contents | Contents
[**GetContent**](ContentsApi.md#GetContent) | **Get** /contents/{contentId} | Getting a content from its id
[**IgnoreContent**](ContentsApi.md#IgnoreContent) | **Put** /contents/{contentId}/ignore | Ignoring a content



## CategorizeContent

> Content CategorizeContent(ctx, contentId, categoryIds)
Categorizing a content

This method updates the categories of a content. If token’s user does not have “read” on this content’s source a 404 HTTP response will be returned.  Authorization​: no.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentId** | **string**|  | 
**categoryIds** | [**[]string**](string.md)| An array containing the new categories to set on the content. | 

### Return type

[**Content**](Content.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContent

> Content CreateContent(ctx, body, optional)
Creating a content

This method allows you to create an new content. It can be a reply to another content or a content that initiate discussion. It use token’s user as content user if he is authorized. Content will be created in Dimelo Digital and pushed asynchronously to the source.  Replying to a customer content is usually possible (unless the source/conversation is read only).  Composing a content on the contrary depend on the source itself: * The source may not support it (and be purely reactive like Instagram, Messenger ...) * Some sources (usually public account) like Twitter or Facebook page allows to publish content without targeting specific individuals. * Some sources (usually non public media) require specific targeting (phone number for SMS, email address for email, customer_id ...) to be able to create a content. This is source specific and detailed under the generic parameters.  Authorization​: only users that can reply or initiate discussion (= compose) on given source. it renders also an error if in_reply_to isn’t synchronized or if in_reply_to content is not under an ​open intervention.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | **string**| The content’s body. This parameter is mandatory. | 
 **optional** | ***CreateContentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateContentOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorId** | **optional.String**| The identity id of content. This parameter is not mandatory, by default it use the token’s user first identity on source. | 
 **inReplyToId** | **optional.String**| The content’s id you want to reply to. If omitted, a new discussion will be created. If source does not support to initiate discussion this parameter is mandatory. | 
 **private** | **optional.Int32**| Created contents are public by default, set this parameter to \&quot;1\&quot; in order to create a private reply. It is NOT supported on every source. | 
 **sourceId** | **optional.String**| The source to create content to. If you specify &#x60;in_reply_to_id&#x60; parameter, source will be determined from it. Otherwise, this parameter is mandatory. | 
 **attachmentIds** | [**optional.Interface of []string**](string.md)| An array containing the attachments’ ids that need to be attached to this content. | 
 **title** | **optional.String**| For an email source. The subject of the email. This parameter is mandatory when initiating a discussion. | 
 **to** | [**optional.Interface of []string**](string.md)| For an email or SMS source. For an email, an array containing the email addresses used in the “To” section of the email. This parameter is mandatory when initiating a discussion. For a SMS, the number the SMS will be sent to. It must start with 00 or +, example: +33634231224 or 0033634231224. This parameter is mandatory when initiating a discussion. | 
 **cc** | [**optional.Interface of []string**](string.md)| For an email source. An array containing the email addresses used in the “Cc” section of the email. | 
 **bcc** | [**optional.Interface of []string**](string.md)| For an email source. An array containing the email addresses used in the “Bcc” section of the email. | 

### Return type

[**Content**](Content.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllContents

> GetAllContentsResponse GetAllContents(ctx, optional)
Contents

This method renders contents ordered by creation date (descending). Only contents in sources where token’s user has “read” permission are returned.  Authorization​: no.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllContentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllContentsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **optional.String**| To filter contents on given query. Query works exactly like threads query but only have those keywords: intervention,identity, identity_group, source, status_in, thread or text. Order can be created_at.desc (default) or created_at.asc. e.g. q&#x3D;intervention:\&quot;7f946431b6eebffafae642cc\&quot;%20source:\&quot;d19c81948c137d86dac77216\&quot; Please refer to ​Search &amp; filtering parameters​ for more details. | 
 **offset** | **optional.Int32**| The record index to start. Default value is 0. | 
 **limit** | **optional.Int32**| The max number of records to return. Default value is 30, max value is 150. | 

### Return type

[**GetAllContentsResponse**](GetAllContentsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContent

> Content GetContent(ctx, contentId)
Getting a content from its id

This method renders a content from given id. If token’s user does not have “read” on content’s source a 404 HTTP response will be returned.  Authorization​: no.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentId** | **string**|  | 

### Return type

[**Content**](Content.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IgnoreContent

> Content IgnoreContent(ctx, contentId)
Ignoring a content

Ignores a content. If token’s user does not have “read” on content’s source a 404 HTTP response will be returned.  Authorization​: no.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentId** | **string**|  | 

### Return type

[**Content**](Content.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

